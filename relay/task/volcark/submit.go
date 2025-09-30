package volcark

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"one-api/common"
	"one-api/common/logger"
	"one-api/model"
	"one-api/providers"
	volcarkProvider "one-api/providers/volcark"
	"one-api/relay/task/base"
	"strings"

	"github.com/gin-gonic/gin"
)

type VolcArkTask struct {
	base.TaskBase
	Request  map[string]any
	Provider *volcarkProvider.VolcArkProvider
	estimate videoEstimate
}

func (t *VolcArkTask) HandleError(err *base.TaskError) {
	StringError(t.C, err.StatusCode, err.Code, err.Message)
}

func (t *VolcArkTask) Init() *base.TaskError {
	var payload map[string]any
	if err := common.UnmarshalBodyReusable(t.C, &payload); err != nil {
		return base.StringTaskError(http.StatusBadRequest, "invalid_request", err.Error(), true)
	}

	modelName, _ := payload["model"].(string)
	modelName = strings.TrimSpace(modelName)
	if modelName == "" {
		return base.StringTaskError(http.StatusBadRequest, "invalid_request", "model is required", true)
	}

	t.Request = payload
	t.OriginalModel = modelName
	t.estimate = estimateFromPayload(modelName, payload)

	if t.estimate.Tokens > 0 {
		estimatedPromptTokens := int(math.Round(float64(t.estimate.Tokens) / 1000.0))
		if estimatedPromptTokens <= 0 {
			estimatedPromptTokens = 1
		}
		t.C.Set("async_estimated_prompt_tokens", estimatedPromptTokens)
	}

	if originTaskID, ok := payload["origin_task_id"].(string); ok {
		t.OriginTaskID = originTaskID
	}

	if err := t.HandleOriginTaskID(); err != nil {
		return base.StringTaskError(http.StatusInternalServerError, "get_origin_task_failed", err.Error(), true)
	}

	return nil
}

func (t *VolcArkTask) SetProvider() *base.TaskError {
	provider, err := t.GetProviderByModel()
	if err != nil {
		return base.StringTaskError(http.StatusServiceUnavailable, "provider_not_found", err.Error(), true)
	}

	volcProvider, ok := provider.(*volcarkProvider.VolcArkProvider)
	if !ok {
		return base.StringTaskError(http.StatusServiceUnavailable, "provider_not_found", "provider not found", true)
	}

	t.Provider = volcProvider
	t.BaseProvider = provider

	return nil
}

func (t *VolcArkTask) Relay() *base.TaskError {
	resp, err := t.Provider.CreateVideoTask(t.Request)
	if err != nil {
		return base.StringTaskError(http.StatusInternalServerError, "submit_failed", err.Message, false)
	}

	t.Response = resp
	t.InitTask()
	t.Task.TaskID = resp.ID
	t.Task.ChannelId = t.Provider.Channel.Id
	t.Task.Action = "contents/generations"
	t.Task.BillingModel = t.GetModelName()
	if t.estimate.Tokens > 0 {
		t.Task.VideoEstimatedTokens = t.estimate.Tokens
	}
	if t.estimate.Width > 0 {
		t.Task.VideoWidth = t.estimate.Width
	}
	if t.estimate.Height > 0 {
		t.Task.VideoHeight = t.estimate.Height
	}
	if t.estimate.FPS > 0 {
		t.Task.VideoFps = t.estimate.FPS
	}
	if t.estimate.Duration > 0 {
		t.Task.VideoDuration = t.estimate.Duration
	}
	if data, marshalErr := json.Marshal(resp); marshalErr == nil {
		t.Task.Data = data
	}

	return nil
}

func (t *VolcArkTask) ShouldRetry(_ *gin.Context, _ *base.TaskError) bool {
	return false
}

func (t *VolcArkTask) GinResponse() {
	t.C.JSON(http.StatusOK, t.Response)
}

func (t *VolcArkTask) UpdateTaskStatus(ctx context.Context, taskChannelM map[int][]string, taskM map[string]*model.Task) error {
	for channelID, taskIds := range taskChannelM {
		channel := model.ChannelGroup.GetChannel(channelID)
		if channel == nil {
			err := model.TaskBulkUpdate(taskIds, map[string]any{
				"fail_reason": fmt.Sprintf("获取渠道信息失败，请联系管理员，渠道ID：%d", channelID),
				"status":      model.TaskStatusFailure,
				"progress":    100,
			})
			if err != nil {
				logger.SysError(fmt.Sprintf("UpdateTask error: %v", err))
			}
			continue
		}

		provider := providers.GetProvider(channel, nil)
		volcProvider, ok := provider.(*volcarkProvider.VolcArkProvider)
		if !ok {
			err := model.TaskBulkUpdate(taskIds, map[string]any{
				"fail_reason": "获取供应商失败，请联系管理员",
				"status":      model.TaskStatusFailure,
				"progress":    100,
			})
			if err != nil {
				logger.SysError(fmt.Sprintf("UpdateTask error: %v", err))
			}
			continue
		}

		for _, taskID := range taskIds {
			task := taskM[taskID]
			if task == nil {
				continue
			}
			if err := t.updateSingleTask(ctx, volcProvider, task); err != nil {
				logger.LogError(ctx, fmt.Sprintf("update task %s error: %s", taskID, err.Error()))
			}
		}
	}

	return nil
}

func (t *VolcArkTask) updateSingleTask(ctx context.Context, provider *volcarkProvider.VolcArkProvider, task *model.Task) error {
	resp, err := provider.GetVideoTask(task.TaskID)
	if err != nil {
		return fmt.Errorf("get task failed: %s", err.Message)
	}

	if resp == nil {
		return fmt.Errorf("empty response")
	}

	applyResponseToTask(task, resp)

	if err := task.Update(); err != nil {
		return err
	}

	return nil
}

func applyResponseToTask(task *model.Task, resp *volcarkProvider.VolcArkVideoTask) {
	status, progress, failReason := mapVolcStatus(resp.Status, resp.Error)
	task.Status = status
	task.Progress = progress
	if resp.CreatedAt != 0 {
		task.SubmitTime = resp.CreatedAt
	}
	if resp.UpdatedAt != 0 {
		task.FinishTime = resp.UpdatedAt
	}

	if failReason != "" {
		task.FailReason = failReason
	} else if status == model.TaskStatusSuccess {
		task.FailReason = ""
	}

	finalizeVolcArkBilling(task, resp)

	if data, err := json.Marshal(resp); err == nil {
		task.Data = data
	}
}

func mapVolcStatus(status string, errInfo *volcarkProvider.VolcArkVideoError) (model.TaskStatus, int, string) {
	switch strings.ToLower(status) {
	case "queued":
		return model.TaskStatusQueued, 20, ""
	case "running":
		return model.TaskStatusInProgress, 60, ""
	case "succeeded":
		return model.TaskStatusSuccess, 100, ""
	case "cancelled":
		reason := "cancelled"
		if errInfo != nil && errInfo.Message != "" {
			reason = errInfo.Message
		}
		return model.TaskStatusFailure, 100, reason
	case "failed":
		reason := ""
		if errInfo != nil {
			reason = errInfo.Message
		}
		return model.TaskStatusFailure, 100, reason
	default:
		return model.TaskStatusUnknown, 0, ""
	}
}
