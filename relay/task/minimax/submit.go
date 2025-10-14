package minimax

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"strings"
	"time"

	"one-api/common"
	"one-api/common/logger"
	"one-api/model"
	"one-api/providers"
	miniProvider "one-api/providers/minimax"
	"one-api/relay/task/base"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

const (
	miniMaxActionText2Video     = "text2video"
	miniMaxActionImage2Video    = "image2video"
	miniMaxActionStartEnd2Video = "start-end2video"
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
	case miniMaxActionText2Video, miniMaxActionImage2Video, miniMaxActionStartEnd2Video:
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
	case miniMaxActionImage2Video:
		if req.FirstFrameImage == "" {
			return fmt.Errorf("first_frame_image is required for image2video")
		}
	case miniMaxActionStartEnd2Video:
		if req.FirstFrameImage == "" || req.LastFrameImage == "" {
			return fmt.Errorf("first_frame_image and last_frame_image are required for start-end2video")
		}
	}

	if len(req.Prompt) > 0 && len([]rune(req.Prompt)) > 2000 {
		return fmt.Errorf("prompt length cannot exceed 2000 characters")
	}

	return nil
}

func (t *MiniMaxTask) SetProvider() *base.TaskError {
	provider, err := t.GetProviderByModel()
	if err != nil {
		return base.StringTaskError(http.StatusServiceUnavailable, "provider_not_found", err.Error(), true)
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

	if data, err := json.Marshal(resp); err == nil {
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

	return task.Update()
}

func defaultMiniMaxModel(action string) string {
	switch action {
	case miniMaxActionText2Video, miniMaxActionImage2Video, miniMaxActionStartEnd2Video:
		return "MiniMax-Hailuo-02"
	default:
		return "MiniMax-Hailuo-02"
	}
}

func inferMiniMaxAction(req *miniProvider.MiniMaxVideoCreateRequest) string {
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

func defaultMiniMaxDuration(model string) int {
	if strings.Contains(strings.ToLower(model), "02") {
		return 6
	}
	return 6
}

func defaultMiniMaxResolution(model string, duration int) string {
	modelLower := strings.ToLower(model)
	if strings.Contains(modelLower, "hailuo-02") {
		if duration >= 10 {
			return "768P"
		}
		return "1080P"
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
