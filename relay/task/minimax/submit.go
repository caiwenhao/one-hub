package minimax

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"hash/crc32"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"one-api/common"
	"one-api/common/config"
	"one-api/common/logger"
	"one-api/model"
	"one-api/providers"
	pbase "one-api/providers/base"
	miniProvider "one-api/providers/minimaxi"
	"one-api/relay"
	"one-api/relay/task/base"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

const (
	miniMaxActionText2Video     = "text2video"
	miniMaxActionImage2Video    = "image2video"
	miniMaxActionStartEnd2Video = "start-end2video"
	miniMaxActionSubject2Video  = "subject2video"
)

type MiniMaxTask struct {
	base.TaskBase
	Action      string
	Request     *miniProvider.MiniMaxVideoCreateRequest
	Provider    *miniProvider.MiniMaxProvider
	VideoClient *miniProvider.MiniMaxVideoClient
}

type taskProperties struct {
	Model      string `json:"model,omitempty"`
	Action     string `json:"action,omitempty"`
	Duration   int    `json:"duration,omitempty"`
	Resolution string `json:"resolution,omitempty"`
}

var miniMaxAllowedMatrix = map[string]map[string]map[string][]int{
    miniMaxActionText2Video: {
        "minimax-hailuo-02": {
            "512p":  {6, 10},
            "768p":  {6, 10},
            "1080p": {6, 10},
        },
        // MiniMax-Hailuo-2.3：768P(6/10s)、1080P(6s)
        "minimax-hailuo-2.3": {
            "768p":  {6, 10},
            "1080p": {6},
        },
        "t2v-01": {
            "720p": {6},
        },
        "t2v-01-director": {
            "720p": {6},
        },
    },
    miniMaxActionImage2Video: {
        "minimax-hailuo-02": {
            "512p":  {6, 10},
            "768p":  {6, 10},
            "1080p": {6, 10},
        },
        // MiniMax-Hailuo-2.3：768P(6/10s)、1080P(6s)
        "minimax-hailuo-2.3": {
            "768p":  {6, 10},
            "1080p": {6},
        },
        // MiniMax-Hailuo-2.3-Fast（图生）：768P(6/10s)、1080P(6s)
        "minimax-hailuo-2.3-fast": {
            "768p":  {6, 10},
            "1080p": {6},
        },
        "i2v-01": {
            "720p": {6},
        },
        "i2v-01-director": {
            "720p": {6},
        },
        "i2v-01-live": {
            "720p": {6},
        },
    },
	miniMaxActionStartEnd2Video: {
		"minimax-hailuo-02": {
			"512p":  {6, 10},
			"768p":  {6, 10},
			"1080p": {6, 10},
		},
	},
	miniMaxActionSubject2Video: {
		"s2v-01": {
			"720p": {6},
		},
	},
}

func (t *MiniMaxTask) HandleError(err *base.TaskError) {
	StringError(t.C, err.StatusCode, err.Code, err.Message)
}

func (t *MiniMaxTask) Init() *base.TaskError {
	rawAction := strings.TrimSpace(strings.ToLower(t.C.Param("action")))

	var req miniProvider.MiniMaxVideoCreateRequest
	if err := common.UnmarshalBodyReusable(t.C, &req); err != nil {
		return base.StringTaskError(http.StatusBadRequest, "invalid_request", err.Error(), true)
	}

	action := rawAction
	if action == "" || action == "video_generation" {
		action = inferMiniMaxAction(&req)
	}

	switch action {
	case miniMaxActionText2Video, miniMaxActionImage2Video, miniMaxActionStartEnd2Video, miniMaxActionSubject2Video:
		// supported actions
	default:
		return base.StringTaskError(http.StatusBadRequest, "invalid_action", fmt.Sprintf("unsupported action %s", action), true)
	}

	if err := t.normalizeRequest(action, &req); err != nil {
		return base.StringTaskError(http.StatusBadRequest, "invalid_request", err.Error(), true)
	}

	t.Action = action
	t.Request = &req
	t.OriginalModel = buildMiniMaxBillingModel(t.Action, req.Model, req.Resolution, req.Duration)

	if originTaskID := t.C.Query("origin_task_id"); originTaskID != "" {
		t.OriginTaskID = originTaskID
	}

	if err := t.HandleOriginTaskID(); err != nil {
		return base.StringTaskError(http.StatusInternalServerError, "get_origin_task_failed", err.Error(), true)
	}

	return nil
}

func (t *MiniMaxTask) normalizeRequest(action string, req *miniProvider.MiniMaxVideoCreateRequest) error {
	req.Model = strings.TrimSpace(req.Model)
	if req.Model == "" {
		req.Model = defaultMiniMaxModel(action)
	}

	req.Prompt = strings.TrimSpace(req.Prompt)
	req.FirstFrameImage = strings.TrimSpace(req.FirstFrameImage)
	req.LastFrameImage = strings.TrimSpace(req.LastFrameImage)
	req.CallbackURL = strings.TrimSpace(req.CallbackURL)
	req.ExternalTaskID = strings.TrimSpace(req.ExternalTaskID)
	req.Resolution = normalizeResolution(req.Resolution)
	req.SubjectReference = sanitizeSubjectReference(req.SubjectReference)

	if req.Duration <= 0 {
		req.Duration = defaultMiniMaxDuration(req.Model)
	}

	if req.Resolution == "" {
		req.Resolution = defaultMiniMaxResolution(req.Model, req.Duration)
	}

	switch action {
	case miniMaxActionText2Video:
		if req.Prompt == "" {
			return fmt.Errorf("prompt is required for text2video")
		}
		if len(req.SubjectReference) > 0 {
			return fmt.Errorf("subject_reference is only supported for subject2video")
		}
	case miniMaxActionImage2Video:
		if req.FirstFrameImage == "" {
			return fmt.Errorf("first_frame_image is required for image2video")
		}
	case miniMaxActionStartEnd2Video:
		if req.FirstFrameImage == "" || req.LastFrameImage == "" {
			return fmt.Errorf("first_frame_image and last_frame_image are required for start-end2video")
		}
		if len(req.SubjectReference) > 0 {
			return fmt.Errorf("subject_reference is only supported for subject2video")
		}
	case miniMaxActionSubject2Video:
		if err := validateSubjectReference(req.SubjectReference); err != nil {
			return err
		}
		if req.Prompt == "" {
			return fmt.Errorf("prompt is required for subject2video")
		}
	}

	if len(req.Prompt) > 0 && len([]rune(req.Prompt)) > 2000 {
		return fmt.Errorf("prompt length cannot exceed 2000 characters")
	}

	if err := validateMiniMaxResolutionDuration(action, req.Model, req.Resolution, req.Duration); err != nil {
		return err
	}

	return nil
}

func (t *MiniMaxTask) SetProvider() *base.TaskError {
	// 限定渠道类型为 minimaxi，避免误选其他视频渠道
	if t.C != nil {
		t.C.Set("allow_channel_type", []int{config.ChannelTypeMiniMax})
	}

	// 优先按“请求模型”选渠道（例如 MiniMax-Hailuo-02）
	var provider pbase.ProviderInterface
	var mapped string
	var fail error
	if t.C != nil {
		provider, mapped, fail = relay.GetProvider(t.C, t.Request.Model)
	}
	if fail != nil || provider == nil {
		// 回退：按常见 minimaxi 文本/语音模型尝试，保证老渠道（未添加视频模型名）也可复用
		fallbacks := []string{"MiniMax-M1", "MiniMax-Text-01", "speech-02-turbo"}
		for _, m := range fallbacks {
			if t.C != nil {
				provider, mapped, fail = relay.GetProvider(t.C, m)
				if fail == nil && provider != nil {
					break
				}
			}
		}
	}
	if provider == nil || fail != nil {
		return base.StringTaskError(http.StatusServiceUnavailable, "provider_not_found", "no available minimaxi channel for video", true)
	}

	mini, ok := provider.(*miniProvider.MiniMaxProvider)
	if !ok {
		return base.StringTaskError(http.StatusServiceUnavailable, "provider_not_found", "provider not found", true)
	}

	videoClient := mini.GetVideoClient()
	if videoClient == nil {
		return base.StringTaskError(http.StatusServiceUnavailable, "provider_not_found", "video capability not configured", true)
	}

	t.Provider = mini
	t.VideoClient = videoClient
	t.BaseProvider = provider
	t.ModelName = mapped
	if t.C != nil {
		t.C.Set("billing_original_model", true)
	}

	return nil
}

func (t *MiniMaxTask) Relay() *base.TaskError {
	resp, err := t.VideoClient.SubmitVideoTask(t.Request)
	if err != nil {
		return base.StringTaskError(http.StatusInternalServerError, "submit_failed", err.Message, false)
	}

	t.Response = resp
	t.InitTask()
	t.Task.TaskID = resp.TaskID
	t.Task.ChannelId = t.Provider.GetChannel().Id
	t.Task.Action = t.Action
	t.Task.VideoDuration = float64(t.Request.Duration)

	props := taskProperties{
		Model:      t.ModelName,
		Action:     t.Action,
		Duration:   t.Request.Duration,
		Resolution: t.Request.Resolution,
	}
	if bytes, marshalErr := json.Marshal(props); marshalErr == nil {
		t.Task.Properties = datatypes.JSON(bytes)
	}

	return nil
}

func (t *MiniMaxTask) ShouldRetry(_ *gin.Context, _ *base.TaskError) bool {
	return false
}

func (t *MiniMaxTask) UpdateTaskStatus(ctx context.Context, taskChannelM map[int][]string, taskM map[string]*model.Task) error {
	for channelID, taskIDs := range taskChannelM {
		if err := updateMiniMaxChannelTasks(ctx, channelID, taskIDs, taskM); err != nil {
			logger.LogError(ctx, fmt.Sprintf("更新 MiniMax 渠道 #%d 任务失败: %v", channelID, err))
		}
	}
	return nil
}

func updateMiniMaxChannelTasks(ctx context.Context, channelID int, taskIDs []string, taskM map[string]*model.Task) error {
	channel := model.ChannelGroup.GetChannel(channelID)
	if channel == nil {
		return markTasksFailed(taskIDs, taskM, fmt.Sprintf("获取渠道信息失败，请联系管理员，渠道ID：%d", channelID))
	}

	provider := providers.GetProvider(channel, nil)
	mini, ok := provider.(*miniProvider.MiniMaxProvider)
	if !ok || mini.GetVideoClient() == nil {
		return markTasksFailed(taskIDs, taskM, "供应商未配置视频能力，请联系管理员")
	}

	videoClient := mini.GetVideoClient()
	for _, id := range taskIDs {
		task := taskM[id]
		modelName := extractModelFromTask(task)
		resp, queryErr := videoClient.QueryVideoTask(task.TaskID, modelName)
		if queryErr != nil {
			logger.LogError(ctx, fmt.Sprintf("查询任务 %s 失败: %s", task.TaskID, queryErr.Message))
			continue
		}

		if err := applyQueryResult(task, resp); err != nil {
			logger.LogError(ctx, fmt.Sprintf("更新任务 %s 状态失败: %v", task.TaskID, err))
		}
	}

	return nil
}

func markTasksFailed(taskIDs []string, taskM map[string]*model.Task, reason string) error {
	for _, id := range taskIDs {
		task := taskM[id]
		task.Status = model.TaskStatusFailure
		task.Progress = 100
		task.FailReason = reason
		task.FinishTime = time.Now().Unix()
		if err := task.Update(); err != nil {
			return err
		}
	}
	return errors.New(reason)
}

func extractModelFromTask(task *model.Task) string {
	if len(task.Properties) == 0 {
		return ""
	}
	var props taskProperties
	if err := json.Unmarshal(task.Properties, &props); err != nil {
		return ""
	}
	return props.Model
}

func applyQueryResult(task *model.Task, resp *miniProvider.MiniMaxVideoQueryResponse) error {
	if resp == nil {
		return nil
	}

	if resp.TaskID == "" {
		resp.TaskID = task.TaskID
	}

	if data, err := marshalNoEscape(resp); err == nil {
		task.Data = data
	}

	status := strings.ToLower(strings.TrimSpace(resp.Status))
	if status == "" && resp.TaskDetail != nil {
		status = strings.ToLower(strings.TrimSpace(resp.TaskDetail.Status))
	}

	progress := 0
	if resp.ProgressPercent > 0 {
		progress = int(math.Max(0, math.Min(resp.ProgressPercent, 100)))
	}

	failureReason := strings.TrimSpace(resp.ErrorMessage)
	if failureReason == "" && resp.TaskDetail != nil {
		failureReason = strings.TrimSpace(resp.TaskDetail.Reason)
	}

	switch status {
	case "success", "succeeded", "completed", "finish", "finished", "done":
		task.Status = model.TaskStatusSuccess
		task.Progress = 100
		task.FinishTime = time.Now().Unix()
		task.FailReason = ""
	case "fail", "failed", "error", "exception":
		task.Status = model.TaskStatusFailure
		task.Progress = 100
		task.FinishTime = time.Now().Unix()
		task.FailReason = failureReason
	default:
		if status == "processing" || status == "pending" || status == "waiting" || status == "queueing" || status == "queued" || status == "running" || status == "in_progress" {
			task.Status = model.TaskStatusInProgress
			if progress == 0 {
				progress = 50
			}
			task.FailReason = ""
		} else if status == "" {
			task.Status = model.TaskStatusInProgress
			if progress == 0 {
				progress = 50
			}
			task.FailReason = ""
		} else {
			task.Status = model.TaskStatusUnknown
			task.FailReason = failureReason
		}
	}

	if task.Status == model.TaskStatusSuccess {
		progress = 100
	}
	if task.Status == model.TaskStatusFailure {
		progress = 100
	}
	if progress > 100 {
		progress = 100
	}
	if progress < 0 {
		progress = 0
	}
	task.Progress = progress
	if task.Status == model.TaskStatusFailure && task.FailReason == "" {
		task.FailReason = failureReason
	}

	// 在任务更新前，尝试写入/更新 artifact（若包含 file_id 或视频直链）
	_ = upsertMiniMaxArtifacts(task, resp)

	return task.Update()
}

// upsertMiniMaxArtifacts 将查询结果中的文件信息入库，建立 file_id → channel_id 映射
func upsertMiniMaxArtifacts(task *model.Task, resp *miniProvider.MiniMaxVideoQueryResponse) error {
	if task == nil || resp == nil {
		return nil
	}
	userID := task.UserId
	channelID := task.ChannelId
	taskID := task.TaskID

	// 提取代表性直链（用于 PPInfra 等上游无 file_id 的场景）
	download := strings.TrimSpace(resp.VideoURL)
	if download == "" {
		download = strings.TrimSpace(resp.WatermarkedURL)
	}
	if download == "" && len(resp.Videos) > 0 {
		download = strings.TrimSpace(resp.Videos[0].VideoURL)
	}

	// 如果 response 未提供 file_id，但有直链，则生成一个稳定的伪 file_id（便于 /v1/files/retrieve 对齐）
	fileID := strings.TrimSpace(resp.FileID)
	if fileID == "" && download != "" {
		sum := crc32.ChecksumIEEE([]byte(taskID + "|" + download))
		// 加上偏移，避免过短
		fileID = strconv.FormatUint(uint64(sum)+1000000000, 10)
		resp.FileID = fileID
	}
	if fileID == "" {
		return nil
	}

	// TTL 解析（如果上游有 TTL 字段）
	var ttlAt int64 = 0
	if len(resp.Videos) > 0 {
		// VideoURLTTL 可能是字符串秒数；忽略解析错误
		if t := strings.TrimSpace(resp.Videos[0].VideoURLTTL); t != "" {
			if secs, err := strconv.ParseInt(t, 10, 64); err == nil && secs > 0 {
				ttlAt = time.Now().Unix() + secs
			}
		}
	}

	artifact := &model.TaskArtifact{
		Platform:     model.TaskPlatformMiniMax,
		UserId:       userID,
		ChannelId:    channelID,
		TaskID:       taskID,
		FileID:       fileID,
		ArtifactType: "video",
		DownloadURL:  download,
		TTLAt:        ttlAt,
	}

	return model.UpsertTaskArtifact(model.DB, artifact)
}

func defaultMiniMaxModel(action string) string {
	switch action {
	case miniMaxActionText2Video, miniMaxActionImage2Video, miniMaxActionStartEnd2Video:
		return "MiniMax-Hailuo-02"
	case miniMaxActionSubject2Video:
		return "S2V-01"
	default:
		return "MiniMax-Hailuo-02"
	}
}

func inferMiniMaxAction(req *miniProvider.MiniMaxVideoCreateRequest) string {
	modelLower := strings.ToLower(strings.TrimSpace(req.Model))
	if len(req.SubjectReference) > 0 || strings.Contains(modelLower, "s2v-01") {
		return miniMaxActionSubject2Video
	}
	first := strings.TrimSpace(req.FirstFrameImage)
	last := strings.TrimSpace(req.LastFrameImage)

	if first != "" && last != "" {
		return miniMaxActionStartEnd2Video
	}
	if first != "" {
		return miniMaxActionImage2Video
	}
	return miniMaxActionText2Video
}

func sanitizeSubjectReference(refs []miniProvider.MiniMaxSubjectReference) []miniProvider.MiniMaxSubjectReference {
	if len(refs) == 0 {
		return nil
	}
	sanitized := make([]miniProvider.MiniMaxSubjectReference, 0, len(refs))
	for _, item := range refs {
		typ := strings.ToLower(strings.TrimSpace(item.Type))
		images := make([]string, 0, len(item.Image))
		for _, img := range item.Image {
			trimmed := strings.TrimSpace(img)
			if trimmed != "" {
				images = append(images, trimmed)
			}
		}
		sanitized = append(sanitized, miniProvider.MiniMaxSubjectReference{
			Type:  typ,
			Image: images,
		})
	}
	return sanitized
}

func validateSubjectReference(refs []miniProvider.MiniMaxSubjectReference) error {
	if len(refs) == 0 {
		return fmt.Errorf("subject_reference is required for subject2video")
	}
	if len(refs) != 1 {
		return fmt.Errorf("subject_reference currently supports exactly one subject")
	}
	ref := refs[0]
	if ref.Type == "" {
		return fmt.Errorf("subject_reference[0].type must be set to character")
	}
	if ref.Type != "character" {
		return fmt.Errorf("subject_reference[0].type must be character")
	}
	if len(ref.Image) == 0 {
		return fmt.Errorf("subject_reference[0].image requires one image URL")
	}
	if len(ref.Image) > 1 {
		return fmt.Errorf("subject_reference[0].image currently supports only one image")
	}
	return nil
}

func validateMiniMaxResolutionDuration(action, model, resolution string, duration int) error {
	if resolution == "" || duration <= 0 {
		return nil
	}
	allowedByAction, ok := miniMaxAllowedMatrix[action]
	if !ok {
		return nil
	}
	modelKey := formatBillingSegment(model)
	resKey := strings.ToLower(strings.TrimSpace(resolution))
	allowedRes := allowedByAction[modelKey]
	if allowedRes == nil {
		if fallback, ok := allowedByAction["*"]; ok {
			allowedRes = fallback
		} else {
			return nil
		}
	}
	allowedDurations, ok := allowedRes[resKey]
	if !ok {
		return fmt.Errorf("model %s does not support resolution %s for %s", model, resolution, action)
	}
	if !containsInt(allowedDurations, duration) {
		return fmt.Errorf("model %s does not support duration %ds at %s for %s", model, duration, resolution, action)
	}
	return nil
}

func containsInt(list []int, target int) bool {
	for _, v := range list {
		if v == target {
			return true
		}
	}
	return false
}

func defaultMiniMaxDuration(model string) int {
    lower := strings.ToLower(model)
    if strings.Contains(lower, "hailuo-02") || strings.Contains(lower, "hailuo-2.3") {
        return 6
    }
    return 6
}

func defaultMiniMaxResolution(model string, duration int) string {
    modelLower := strings.ToLower(model)
    if strings.Contains(modelLower, "hailuo-02") || strings.Contains(modelLower, "hailuo-2.3") {
        // 对齐官方文档：
        // - MiniMax-Hailuo-02 在 6s 与 10s 的默认分辨率均为 768P
        // - 6s 也支持 1080P，但默认值为 768P
        return "768P"
    }
	return "720P"
}

func normalizeResolution(resolution string) string {
	res := strings.TrimSpace(resolution)
	if res == "" {
		return ""
	}
	res = strings.ToUpper(res)
	res = strings.ReplaceAll(res, " ", "")
	if !strings.HasSuffix(res, "P") && res != "" {
		res = res + "P"
	}
	return res
}

func buildMiniMaxBillingModel(action, model, resolution string, duration int) string {
	modelKey := formatBillingSegment(model)
	resolutionKey := strings.ToLower(resolution)
	if resolutionKey == "" {
		resolutionKey = "default"
	}
	return fmt.Sprintf("minimax-%s-%s-%s-%ds", action, modelKey, resolutionKey, duration)
}

func formatBillingSegment(input string) string {
	segment := strings.ToLower(strings.TrimSpace(input))
	segment = strings.ReplaceAll(segment, " ", "-")
	segment = strings.ReplaceAll(segment, "_", "-")
	segment = strings.ReplaceAll(segment, "/", "-")
	return segment
}
