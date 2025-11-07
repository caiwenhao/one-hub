package relay

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"one-api/common"
	"one-api/common/config"
	"one-api/common/logger"
	"one-api/model"
	providersBase "one-api/providers/base"
	"one-api/relay/relay_util"
	"one-api/types"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

type soraTaskProperties struct {
	Model        string `json:"model"`
	Resolution   string `json:"resolution"`
	Seconds      int    `json:"seconds"`
	Orientation  string `json:"orientation"`
	SizeLabel    string `json:"size_label"`
	BillingModel string `json:"billing_model"`
	Prompt       string `json:"prompt,omitempty"`
}

type soraVideoResponse struct {
	ID                 string          `json:"id"`
	Object             string          `json:"object"`
	CreatedAt          int64           `json:"created_at,omitempty"`
	CompletedAt        int64           `json:"completed_at,omitempty"`
	ExpiresAt          int64           `json:"expires_at,omitempty"`
	Status             string          `json:"status"`
	Model              string          `json:"model,omitempty"`
	Prompt             string          `json:"prompt,omitempty"`
	Progress           int             `json:"progress"`
	Seconds            string          `json:"seconds,omitempty"`
	Size               string          `json:"size,omitempty"`
	RemixedFromVideoID string          `json:"remixed_from_video_id,omitempty"`
	Error              *soraVideoError `json:"error,omitempty"`
}

type soraVideoError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type soraVideoList struct {
	Object string               `json:"object"`
	Data   []*soraVideoResponse `json:"data"`
}

func VideoCreate(c *gin.Context) {
    var req types.VideoCreateRequest
    if err := common.UnmarshalBodyReusable(c, &req); err != nil {
        common.AbortWithMessage(c, http.StatusBadRequest, err.Error())
        return
    }

    originalModel := strings.TrimSpace(req.Model)
    if originalModel == "" {
        // 官方默认模型 sora-2
        originalModel = "sora-2"
    }

    // 仅对外暴露官方模型：sora-2 / sora-2-pro
    if m := strings.ToLower(originalModel); m != "sora-2" && m != "sora-2-pro" {
        common.AbortWithMessage(c, http.StatusBadRequest, "invalid model: only 'sora-2' or 'sora-2-pro' are allowed")
        return
    }

	normalize := normalizeSoraSizeInfo(req.Size)
	if req.Seconds <= 0 {
		req.Seconds = defaultSoraDuration()
	}

	provider, mappedModel, err := GetProvider(c, originalModel)
	if err != nil {
		common.AbortWithMessage(c, http.StatusServiceUnavailable, err.Error())
		return
	}
	if mappedModel == "" {
		mappedModel = originalModel
	}
	req.Model = mappedModel

	videoProvider, ok := provider.(providersBase.VideoInterface)
	if !ok {
		common.AbortWithMessage(c, http.StatusNotImplemented, "video interface not implemented for channel")
		return
	}

	// 针对 NewAPI 渠道：按次计费，不依赖 seconds
	isNewAPI := c.GetInt("channel_type") == config.ChannelTypeNewAPI
	var quota *relay_util.Quota
	var billingModel string
	if isNewAPI {
		billingModel = originalModel
		quota = relay_util.NewQuota(c, billingModel, 0)
	} else {
		billingModel = buildSoraBillingModel(mappedModel, normalize.Resolution)
		quota = relay_util.NewQuota(c, billingModel, req.Seconds)
	}
	if errWithCode := quota.PreQuotaConsumption(); errWithCode != nil {
		newErr := FilterOpenAIErr(c, errWithCode)
		relayResponseWithOpenAIErr(c, &newErr)
		return
	}

	var job *types.VideoJob
	job, errWithCode := videoProvider.CreateVideo(&req)
	if errWithCode != nil {
		quota.Undo(c)
		newErr := FilterOpenAIErr(c, errWithCode)
		relayResponseWithOpenAIErr(c, &newErr)
		return
	}

	if job.Object == "" {
		job.Object = "video"
	}
	job.Model = originalModel
	if job.Prompt == "" {
		job.Prompt = req.Prompt
	}
	if job.Seconds == 0 {
		job.Seconds = req.Seconds
	}
	if job.Size == "" {
		job.Size = normalize.Resolution
	}

	usage := &types.Usage{}
	if !isNewAPI {
		usage.PromptTokens = req.Seconds
		usage.TotalTokens = req.Seconds
	}
	quota.Consume(c, usage, false)

	props := soraTaskProperties{
		Model:        originalModel,
		Resolution:   normalize.Resolution,
		Seconds:      req.Seconds,
		Orientation:  normalize.Orientation,
		SizeLabel:    normalize.SizeLabel,
		BillingModel: billingModel,
		Prompt:       req.Prompt,
	}
	saveSoraTask(c, provider.GetChannel().Id, job, props)

	c.JSON(http.StatusOK, newSoraVideoResponse(job))
}

func VideoRetrieve(c *gin.Context) {
	videoID := strings.TrimSpace(c.Param("id"))
	if videoID == "" {
		common.AbortWithMessage(c, http.StatusBadRequest, "video id is required")
		return
	}

	task, err := model.GetTaskByTaskId(model.TaskPlatformSora, c.GetInt("id"), videoID)
	if err != nil {
		common.AbortWithMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	var (
		props       soraTaskProperties
		provider    providersBase.ProviderInterface
		mappedModel string
	)

	modelName := ""
	if task != nil {
		props = parseSoraTaskProperties(task.Properties)
		c.Set("specific_channel_id", task.ChannelId)
		modelName = strings.TrimSpace(props.Model)
	}
	if modelName == "" {
		modelName = "sora-2"
	}

	provider, mappedModel, err = GetProvider(c, modelName)
	if err != nil {
		common.AbortWithMessage(c, http.StatusServiceUnavailable, err.Error())
		return
	}
	if mappedModel == "" {
		mappedModel = modelName
	}

	videoProvider, ok := provider.(providersBase.VideoInterface)
	if !ok {
		common.AbortWithMessage(c, http.StatusNotImplemented, "video interface not implemented for channel")
		return
	}

	job, errWithCode := videoProvider.RetrieveVideo(videoID)
	if errWithCode != nil {
		newErr := FilterOpenAIErr(c, errWithCode)
		relayResponseWithOpenAIErr(c, &newErr)
		return
	}

    if job.Object == "" {
        job.Object = "video"
    }
    // 始终优先使用本地任务保存的官方模型名，避免上游内部 SKU 外泄
    if strings.TrimSpace(props.Model) != "" {
        job.Model = props.Model
    } else if strings.TrimSpace(job.Model) == "" {
        if mappedModel != "" {
            job.Model = mappedModel
        } else {
            job.Model = "sora-2"
        }
    }
	if job.Size == "" {
		if props.Resolution != "" {
			job.Size = props.Resolution
		} else {
			job.Size = normalizeSoraSizeInfo("").Resolution
		}
	}
	if job.Seconds == 0 {
		if props.Seconds > 0 {
			job.Seconds = props.Seconds
		} else {
			job.Seconds = defaultSoraDuration()
		}
	}
	if strings.TrimSpace(job.Prompt) == "" && strings.TrimSpace(props.Prompt) != "" {
		job.Prompt = props.Prompt
	}

	if task != nil {
		updateSoraTask(task, job)
	}

	c.JSON(http.StatusOK, newSoraVideoResponse(job))
}

func VideoDownload(c *gin.Context) {
	videoID := strings.TrimSpace(c.Param("id"))
	if videoID == "" {
		common.AbortWithMessage(c, http.StatusBadRequest, "video id is required")
		return
	}

	task, err := model.GetTaskByTaskId(model.TaskPlatformSora, c.GetInt("id"), videoID)
	if err != nil {
		common.AbortWithMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	modelName := ""
	if task != nil {
		props := parseSoraTaskProperties(task.Properties)
		modelName = strings.TrimSpace(props.Model)
		c.Set("specific_channel_id", task.ChannelId)
	}
	if modelName == "" {
		modelName = "sora-2"
	}

	provider, _, err := GetProvider(c, modelName)
	if err != nil {
		common.AbortWithMessage(c, http.StatusServiceUnavailable, err.Error())
		return
	}
	videoProvider, ok := provider.(providersBase.VideoInterface)
	if !ok {
		common.AbortWithMessage(c, http.StatusNotImplemented, "video interface not implemented for channel")
		return
	}

	variant := strings.TrimSpace(c.Query("variant"))
	resp, errWithCode := videoProvider.DownloadVideo(videoID, variant)
	if errWithCode != nil {
		newErr := FilterOpenAIErr(c, errWithCode)
		relayResponseWithOpenAIErr(c, &newErr)
		return
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			c.Writer.Header().Add(key, value)
		}
	}

	c.Status(resp.StatusCode)
	if _, err := io.Copy(c.Writer, resp.Body); err != nil {
		logger.LogError(c.Request.Context(), "copy_video_content_failed:"+err.Error())
	}
}

// VideoRemix 实现 /v1/videos/{id}/remix
func VideoRemix(c *gin.Context) {
	videoID := strings.TrimSpace(c.Param("id"))
	if videoID == "" {
		common.AbortWithMessage(c, http.StatusBadRequest, "video id is required")
		return
	}

	var req types.VideoRemixRequest
	if err := common.UnmarshalBodyReusable(c, &req); err != nil {
		common.AbortWithMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	if strings.TrimSpace(req.Prompt) == "" {
		common.AbortWithMessage(c, http.StatusBadRequest, "prompt is required")
		return
	}

	// 优先根据历史任务锁定渠道与属性
	var props soraTaskProperties
	task, _ := model.GetTaskByTaskId(model.TaskPlatformSora, c.GetInt("id"), videoID)
	if task != nil {
		props = parseSoraTaskProperties(task.Properties)
		c.Set("specific_channel_id", task.ChannelId)
	}

	// 选择供应商（若无历史任务，则默认以 sora-2 选择）
	modelName := props.Model
	if strings.TrimSpace(modelName) == "" {
		modelName = "sora-2"
	}

	provider, mappedModel, err := GetProvider(c, modelName)
	if err != nil {
		common.AbortWithMessage(c, http.StatusServiceUnavailable, err.Error())
		return
	}
	if mappedModel == "" {
		mappedModel = modelName
	}

	videoProvider, ok := provider.(providersBase.VideoInterface)
	if !ok {
		common.AbortWithMessage(c, http.StatusNotImplemented, "video interface not implemented for channel")
		return
	}

	// 计费：使用原视频的时长/分辨率作为预估
	seconds := props.Seconds
	if seconds <= 0 {
		seconds = defaultSoraDuration()
	}
	billingModel := buildSoraBillingModel(mappedModel, props.Resolution)
	quota := relay_util.NewQuota(c, billingModel, seconds)
	if errWithCode := quota.PreQuotaConsumption(); errWithCode != nil {
		newErr := FilterOpenAIErr(c, errWithCode)
		relayResponseWithOpenAIErr(c, &newErr)
		return
	}

	job, errWithCode := videoProvider.RemixVideo(videoID, req.Prompt)
	if errWithCode != nil {
		quota.Undo(c)
		newErr := FilterOpenAIErr(c, errWithCode)
		relayResponseWithOpenAIErr(c, &newErr)
		return
	}

	if job.Object == "" {
		job.Object = "video"
	}
	if job.Model == "" {
		job.Model = modelName
	}
	if job.RemixedFromVideoID == "" {
		job.RemixedFromVideoID = videoID
	}
	if job.Quality == "" {
		job.Quality = "standard"
	}
	if job.CreatedAt == 0 {
		job.CreatedAt = time.Now().Unix()
	}

	usage := &types.Usage{PromptTokens: seconds, TotalTokens: seconds}
	quota.Consume(c, usage, false)

	// 若历史 props 缺失，构建默认属性保存
	if props.Model == "" {
		normalize := normalizeSoraSizeInfo("")
		props = soraTaskProperties{
			Model: mappedModel, Resolution: normalize.Resolution, Seconds: seconds,
			Orientation: normalize.Orientation, SizeLabel: normalize.SizeLabel,
			BillingModel: billingModel, Prompt: req.Prompt,
		}
	} else if strings.TrimSpace(props.Prompt) == "" {
		props.Prompt = req.Prompt
	}

	if props.Prompt == "" {
		props.Prompt = req.Prompt
	}

	saveSoraTask(c, provider.GetChannel().Id, job, props)
	c.JSON(http.StatusOK, newSoraVideoResponse(job))
}

// VideoList 实现 /v1/videos（官方上游列表）。
// 注意：官方为组织级别列表。此处采用上游直透策略。
func VideoList(c *gin.Context) {
	after := strings.TrimSpace(c.Query("after"))
	order := strings.TrimSpace(c.Query("order"))
	limitVal := strings.TrimSpace(c.Query("limit"))
	limit := 0
	if limitVal != "" {
		if v, err := strconv.Atoi(limitVal); err == nil {
			limit = v
		}
	}

	// 选择一个 OpenAI 视频渠道（按默认模型 sora-2 选路）
	provider, _, err := GetProvider(c, "sora-2")
	if err != nil {
		common.AbortWithMessage(c, http.StatusServiceUnavailable, err.Error())
		return
	}
	videoProvider, ok := provider.(providersBase.VideoInterface)
	if !ok {
		common.AbortWithMessage(c, http.StatusNotImplemented, "video interface not implemented for channel")
		return
	}

	list, errWithCode := videoProvider.ListVideos(after, limit, order)
	if errWithCode != nil {
		newErr := FilterOpenAIErr(c, errWithCode)
		relayResponseWithOpenAIErr(c, &newErr)
		return
	}
	resp := &soraVideoList{
		Object: "list",
		Data:   []*soraVideoResponse{},
	}
	if list != nil {
		for i := range list.Data {
			job := list.Data[i]
			resp.Data = append(resp.Data, newSoraVideoResponse(&job))
		}
	}
	c.JSON(http.StatusOK, resp)
}

// VideoDelete 实现 DELETE /v1/videos/{id}
func VideoDelete(c *gin.Context) {
	videoID := strings.TrimSpace(c.Param("id"))
	if videoID == "" {
		common.AbortWithMessage(c, http.StatusBadRequest, "video id is required")
		return
	}

	// 优先根据历史任务锁定渠道
	task, _ := model.GetTaskByTaskId(model.TaskPlatformSora, c.GetInt("id"), videoID)
	if task != nil {
		c.Set("specific_channel_id", task.ChannelId)
	}

	provider, _, err := GetProvider(c, "sora-2")
	if err != nil {
		common.AbortWithMessage(c, http.StatusServiceUnavailable, err.Error())
		return
	}
	videoProvider, ok := provider.(providersBase.VideoInterface)
	if !ok {
		common.AbortWithMessage(c, http.StatusNotImplemented, "video interface not implemented for channel")
		return
	}

	job, errWithCode := videoProvider.DeleteVideo(videoID)
	if errWithCode != nil {
		newErr := FilterOpenAIErr(c, errWithCode)
		relayResponseWithOpenAIErr(c, &newErr)
		return
	}
	if job == nil {
		job = &types.VideoJob{ID: videoID, Object: "video"}
	}
	c.JSON(http.StatusOK, newSoraVideoResponse(job))
}

func saveSoraTask(c *gin.Context, channelID int, job *types.VideoJob, props soraTaskProperties) {
	if job == nil || job.ID == "" {
		return
	}
	if props.Prompt == "" && job.Prompt != "" {
		props.Prompt = job.Prompt
	}

	task := &model.Task{
		TaskID:        job.ID,
		Platform:      model.TaskPlatformSora,
		UserId:        c.GetInt("id"),
		TokenID:       c.GetInt("token_id"),
		ChannelId:     channelID,
		SubmitTime:    submitTimeFromJob(job),
		Status:        mapVideoStatus(job.Status),
		Progress:      int(math.Round(job.Progress)),
		Action:        "video.generate",
		BillingModel:  props.BillingModel,
		VideoDuration: float64(props.Seconds),
	}

	if propsJSON, err := json.Marshal(props); err == nil {
		task.Properties = datatypes.JSON(propsJSON)
	}
	if job.Result != nil {
		if dataJSON, err := json.Marshal(job.Result); err == nil {
			task.Data = datatypes.JSON(dataJSON)
		}
	}

	if job.Status == "completed" || job.Status == "failed" {
		task.FinishTime = time.Now().Unix()
	}

	if err := task.Insert(); err != nil {
		logger.LogError(c.Request.Context(), "save_sora_task_failed:"+err.Error())
	}
}

func updateSoraTask(task *model.Task, job *types.VideoJob) {
	if task == nil || job == nil {
		return
	}

	task.Status = mapVideoStatus(job.Status)
	task.Progress = int(math.Round(job.Progress))
	if job.Seconds > 0 {
		task.VideoDuration = float64(job.Seconds)
	}
	if job.Status == "completed" || job.Status == "failed" {
		task.FinishTime = time.Now().Unix()
	}
	if job.Result != nil {
		if dataJSON, err := json.Marshal(job.Result); err == nil {
			task.Data = datatypes.JSON(dataJSON)
		}
	}

	if err := task.Update(); err != nil {
		logger.LogError(context.Background(), "update_sora_task_failed:"+err.Error())
	}
}

func parseSoraTaskProperties(data datatypes.JSON) soraTaskProperties {
	var props soraTaskProperties
	if len(data) == 0 {
		return props
	}
	if err := json.Unmarshal(data, &props); err != nil {
		return props
	}
	return props
}

func mapVideoStatus(status string) model.TaskStatus {
	switch strings.ToLower(status) {
	case "queued":
		return model.TaskStatusQueued
	case "in_progress", "processing":
		return model.TaskStatusInProgress
	case "completed", "success":
		return model.TaskStatusSuccess
	case "failed", "error":
		return model.TaskStatusFailure
	default:
		return model.TaskStatusUnknown
	}
}

func normalizeSoraSizeInfo(size string) soraSizeInfoHelper {
	value := strings.ToLower(strings.TrimSpace(size))
	value = strings.ReplaceAll(value, " ", "")
	if value == "" {
		return soraSizeInfoHelper{
			Resolution:  "720x1280",
			Orientation: "portrait",
			SizeLabel:   "small",
		}
	}

	if strings.Contains(value, "x") {
		parts := strings.Split(value, "x")
		if len(parts) == 2 {
			w, _ := strconv.Atoi(parts[0])
			h, _ := strconv.Atoi(parts[1])
			orientation := "landscape"
			if h > w {
				orientation = "portrait"
			}
			resolution := fmt.Sprintf("%dx%d", w, h)
			return soraSizeInfoHelper{
				Resolution:  resolution,
				Orientation: orientation,
				SizeLabel:   mapResolutionSizeLabel(resolution),
			}
		}
	}

	switch value {
	case "landscape":
		return soraSizeInfoHelper{
			Resolution:  "1280x720",
			Orientation: "landscape",
			SizeLabel:   "small",
		}
	case "portrait":
		return soraSizeInfoHelper{
			Resolution:  "720x1280",
			Orientation: "portrait",
			SizeLabel:   "small",
		}
	default:
		return soraSizeInfoHelper{
			Resolution:  "720x1280",
			Orientation: "portrait",
			SizeLabel:   "small",
		}
	}
}

type soraSizeInfoHelper struct {
	Resolution  string
	Orientation string
	SizeLabel   string
}

func mapResolutionSizeLabel(resolution string) string {
	switch resolution {
	case "1280x720", "720x1280":
		return "small"
	case "1792x1024", "1024x1792":
		return "large"
	default:
		return "small"
	}
}

func buildSoraBillingModel(modelName, resolution string) string {
	modelName = strings.TrimSpace(modelName)
	if resolution == "" {
		return modelName
	}
	// Sutui 上游模型（sora_video2-*）已经在模型名中体现了方向/清晰度/时长等，无需追加分辨率，保持与定价键一致
	lower := strings.ToLower(modelName)
	if strings.HasPrefix(lower, "sora_video2") {
		return modelName
	}
	return fmt.Sprintf("%s-%s", modelName, resolution)
}

func submitTimeFromJob(job *types.VideoJob) int64 {
	if job == nil {
		return time.Now().Unix()
	}
	if job.CreatedAt > 0 {
		return job.CreatedAt
	}
	return time.Now().Unix()
}

func defaultSoraDuration() int {
	// 官方默认 4 秒
	return 4
}

func newSoraVideoResponse(job *types.VideoJob) *soraVideoResponse {
	if job == nil {
		return nil
	}
	resp := &soraVideoResponse{
		ID:                 strings.TrimSpace(job.ID),
		Object:             strings.TrimSpace(job.Object),
		CreatedAt:          job.CreatedAt,
		CompletedAt:        job.CompletedAt,
		ExpiresAt:          job.ExpiresAt,
		Status:             strings.TrimSpace(job.Status),
		Model:              strings.TrimSpace(job.Model),
		Prompt:             job.Prompt,
		RemixedFromVideoID: strings.TrimSpace(job.RemixedFromVideoID),
	}
	if resp.ID == "" {
		resp.ID = job.ID
	}
	if resp.Object == "" {
		resp.Object = "video"
	}
	if resp.Status == "" {
		resp.Status = "queued"
	}
	resp.Progress = clampSoraProgress(int(math.Round(job.Progress)))

	seconds := job.Seconds
	if seconds <= 0 {
		seconds = defaultSoraDuration()
	}
	resp.Seconds = strconv.Itoa(seconds)

	size := strings.TrimSpace(job.Size)
	if size == "" {
		size = normalizeSoraSizeInfo("").Resolution
	}
	resp.Size = size

	resp.Error = convertSoraVideoError(job.Error)
	return resp
}

func convertSoraVideoError(err *types.VideoJobError) *soraVideoError {
	if err == nil {
		return nil
	}
	code := stringifySoraErrorCode(err.Code)
	if code == "" && strings.TrimSpace(err.Message) == "" {
		return nil
	}
	return &soraVideoError{
		Code:    code,
		Message: err.Message,
	}
}

func stringifySoraErrorCode(code any) string {
	switch v := code.(type) {
	case nil:
		return ""
	case string:
		return v
	case fmt.Stringer:
		return v.String()
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	default:
		return fmt.Sprintf("%v", v)
	}
}

func clampSoraProgress(val int) int {
	switch {
	case val < 0:
		return 0
	case val > 100:
		return 100
	default:
		return val
	}
}
