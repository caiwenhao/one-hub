package relay

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"one-api/common"
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
}

func VideoCreate(c *gin.Context) {
	var req types.VideoCreateRequest
	if err := common.UnmarshalBodyReusable(c, &req); err != nil {
		common.AbortWithMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	originalModel := strings.TrimSpace(req.Model)
	if originalModel == "" {
		common.AbortWithMessage(c, http.StatusBadRequest, "model is required")
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

	billingModel := buildSoraBillingModel(mappedModel, normalize.Resolution)
	quota := relay_util.NewQuota(c, billingModel, req.Seconds)
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
	if job.Seconds == 0 {
		job.Seconds = req.Seconds
	}
	if job.Size == "" {
		job.Size = normalize.Resolution
	}

	usage := &types.Usage{
		PromptTokens: req.Seconds,
		TotalTokens:  req.Seconds,
	}
	quota.Consume(c, usage, false)

	props := soraTaskProperties{
		Model:        originalModel,
		Resolution:   normalize.Resolution,
		Seconds:      req.Seconds,
		Orientation:  normalize.Orientation,
		SizeLabel:    normalize.SizeLabel,
		BillingModel: billingModel,
	}
	saveSoraTask(c, provider.GetChannel().Id, job, props)

	c.JSON(http.StatusOK, job)
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
	if task == nil {
		common.AbortWithMessage(c, http.StatusNotFound, "video task not found")
		return
	}

	props := parseSoraTaskProperties(task.Properties)

	c.Set("specific_channel_id", task.ChannelId)
	provider, mappedModel, err := GetProvider(c, props.Model)
	if err != nil {
		common.AbortWithMessage(c, http.StatusServiceUnavailable, err.Error())
		return
	}
	if mappedModel == "" {
		mappedModel = props.Model
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
	if job.Model == "" {
		job.Model = props.Model
	}
	if job.Size == "" && props.Resolution != "" {
		job.Size = props.Resolution
	}
	if job.Seconds == 0 && props.Seconds > 0 {
		job.Seconds = props.Seconds
	}

	updateSoraTask(task, job)

	c.JSON(http.StatusOK, job)
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
	if task == nil {
		common.AbortWithMessage(c, http.StatusNotFound, "video task not found")
		return
	}

	props := parseSoraTaskProperties(task.Properties)
	c.Set("specific_channel_id", task.ChannelId)
	provider, _, err := GetProvider(c, props.Model)
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

func saveSoraTask(c *gin.Context, channelID int, job *types.VideoJob, props soraTaskProperties) {
	if job == nil || job.ID == "" {
		return
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
			Resolution:  "1280x720",
			Orientation: "landscape",
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
			Resolution:  "1280x720",
			Orientation: "landscape",
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
	return 8
}
