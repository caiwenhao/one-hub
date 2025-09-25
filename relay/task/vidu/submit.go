package vidu

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"one-api/common"
	"one-api/common/logger"
	"one-api/model"
	"one-api/providers"
	viduProvider "one-api/providers/vidu"
	"one-api/relay/task/base"
	"one-api/types"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

// ViduTask 任务结构体 - 支持多种请求类型
type ViduTask struct {
	base.TaskBase
	Action   string
	Request  interface{} // 支持不同类型的请求
	Provider *viduProvider.ViduProvider
}

func (t *ViduTask) HandleError(err *base.TaskError) {
	StringError(t.C, err.StatusCode, err.Code, err.Message)
}

func (t *ViduTask) Init() *base.TaskError {
	t.Action = t.C.Param("action")

	// 根据不同action解析不同的请求体
	var err error
	switch t.Action {
	case viduProvider.ViduActionReference2Image:
		var req viduProvider.ViduReference2ImageRequest
		if err = common.UnmarshalBodyReusable(t.C, &req); err != nil {
			return base.StringTaskError(http.StatusBadRequest, "invalid_request", err.Error(), true)
		}
		t.Request = &req
	default:
		// 通用视频任务请求
		var req viduProvider.ViduTaskRequest
		if err = common.UnmarshalBodyReusable(t.C, &req); err != nil {
			return base.StringTaskError(http.StatusBadRequest, "invalid_request", err.Error(), true)
		}
		t.Request = &req
	}

	err = t.actionValidate()
	if err != nil {
		return base.StringTaskError(http.StatusBadRequest, "invalid_request", err.Error(), true)
	}

	err = t.HandleOriginTaskID()
	if err != nil {
		return base.StringTaskError(http.StatusInternalServerError, "get_origin_task_failed", err.Error(), true)
	}

	return nil
}

func (t *ViduTask) SetProvider() *base.TaskError {
	// 通过模型查询渠道
	provider, err := t.GetProviderByModel()
	if err != nil {
		return base.StringTaskError(http.StatusServiceUnavailable, "provider_not_found", err.Error(), true)
	}

	viduProvider, ok := provider.(*viduProvider.ViduProvider)
	if !ok {
		return base.StringTaskError(http.StatusServiceUnavailable, "provider_not_found", "provider not found", true)
	}

	t.Provider = viduProvider
	t.BaseProvider = provider

	return nil
}

func (t *ViduTask) Relay() *base.TaskError {
	resp, err := t.Provider.Submit(t.Action, t.Request)
	if err != nil {
		// 将 OpenAIError 转换为 OpenAIErrorWithStatusCode
		errWithCode := &types.OpenAIErrorWithStatusCode{
			OpenAIError: *err,
			StatusCode:  500,
		}
		return base.OpenAIErrToTaskErr(errWithCode)
	}

	t.Response = resp

	t.InitTask()
	t.Task.TaskID = resp.TaskID
	t.Task.ChannelId = t.Provider.Channel.Id
	t.Task.Action = t.Action

	return nil
}

func (t *ViduTask) actionValidate() error {
	// 验证action是否为支持的类型
	switch t.Action {
	case viduProvider.ViduActionImg2Video:
		req := t.Request.(*viduProvider.ViduTaskRequest)
		if len(req.Images) != 1 {
			return fmt.Errorf("images must contain exactly 1 image for img2video")
		}
		// 设置默认模型和参数
		if req.Model == "" {
			req.Model = viduProvider.ViduModel15
		}
		if req.Duration == nil {
			defaultDuration := getDefaultDuration(req.Model)
			req.Duration = &defaultDuration
		}
		if req.Resolution == "" {
			req.Resolution = getDefaultResolution(req.Model, *req.Duration)
		}
		// 构造详细的模型名称用于计费
		t.OriginalModel = buildDetailedModelName(t.Action, req.Model, *req.Duration, req.Resolution, "")

	case viduProvider.ViduActionReference2Video:
		req := t.Request.(*viduProvider.ViduTaskRequest)
		if len(req.Images) == 0 {
			return fmt.Errorf("images is required for reference2video")
		}
		if len(req.Images) > 7 {
			return fmt.Errorf("reference2video supports up to 7 images")
		}
		if strings.TrimSpace(req.Prompt) == "" {
			return fmt.Errorf("prompt is required for reference2video")
		}
		// 设置默认模型和参数
		if req.Model == "" {
			req.Model = viduProvider.ViduModelQ1
		}
		if req.Duration == nil {
			defaultDuration := getDefaultDuration(req.Model)
			req.Duration = &defaultDuration
		}
		if req.Resolution == "" {
			req.Resolution = getDefaultResolution(req.Model, *req.Duration)
		}
		// 构造详细的模型名称用于计费
		t.OriginalModel = buildDetailedModelName(t.Action, req.Model, *req.Duration, req.Resolution, "")

	case viduProvider.ViduActionStartEnd2Video:
		req := t.Request.(*viduProvider.ViduTaskRequest)
		if len(req.Images) != 2 {
			return fmt.Errorf("exactly 2 images are required for start-end2video (start and end frames)")
		}
		// 设置默认模型和参数
		if req.Model == "" {
			req.Model = viduProvider.ViduModelQ2Pro
		}
		if req.Duration == nil {
			defaultDuration := getDefaultDuration(req.Model)
			req.Duration = &defaultDuration
		}
		if req.Resolution == "" {
			req.Resolution = getDefaultResolution(req.Model, *req.Duration)
		}
		// 构造详细的模型名称用于计费
		t.OriginalModel = buildDetailedModelName(t.Action, req.Model, *req.Duration, req.Resolution, "")

	case viduProvider.ViduActionText2Video:
		req := t.Request.(*viduProvider.ViduTaskRequest)
		if strings.TrimSpace(req.Prompt) == "" {
			return fmt.Errorf("prompt is required for text2video")
		}
		// 设置默认模型和参数
		if req.Model == "" {
			req.Model = viduProvider.ViduModelQ1
		}
		if req.Duration == nil {
			defaultDuration := getDefaultDuration(req.Model)
			req.Duration = &defaultDuration
		}
		if req.Resolution == "" {
			req.Resolution = getDefaultResolution(req.Model, *req.Duration)
		}
		if req.Style == "" {
			req.Style = "general"
		}
		// 构造详细的模型名称用于计费（包含风格）
		t.OriginalModel = buildDetailedModelName(t.Action, req.Model, *req.Duration, req.Resolution, req.Style)

	case viduProvider.ViduActionReference2Image:
		req := t.Request.(*viduProvider.ViduReference2ImageRequest)
		if len(req.Images) == 0 {
			return fmt.Errorf("images is required for reference2image")
		}
		if strings.TrimSpace(req.Prompt) == "" {
			return fmt.Errorf("prompt is required for reference2image")
		}
		// 设置默认模型
		if req.Model == "" {
			req.Model = viduProvider.ViduModelQ1
		}
		// 构造简化的模型名称用于计费
		t.OriginalModel = fmt.Sprintf("vidu-%s-%s", t.Action, req.Model)

	default:
		return fmt.Errorf("unsupported action: %s", t.Action)
	}

	return nil
}

// 获取模型默认时长
func getDefaultDuration(model string) int {
	switch model {
	case viduProvider.ViduModelQ2Pro, viduProvider.ViduModelQ2Turbo:
		return 5
	case viduProvider.ViduModelQ1, viduProvider.ViduModelQ1Classic:
		return 5
	case viduProvider.ViduModel20:
		return 4
	case viduProvider.ViduModel15:
		return 4
	default:
		return 5
	}
}

// 获取模型默认分辨率
func getDefaultResolution(model string, duration int) string {
	switch model {
	case viduProvider.ViduModelQ2Pro, viduProvider.ViduModelQ2Turbo:
		return "720p"
	case viduProvider.ViduModelQ1, viduProvider.ViduModelQ1Classic:
		return "1080p"
	case viduProvider.ViduModel20:
		if duration == 4 {
			return "360p"
		}
		return "720p"
	case viduProvider.ViduModel15:
		if duration == 4 {
			return "360p"
		}
		return "720p"
	default:
		return "720p"
	}
}

// 构造详细的模型名称用于计费
func buildDetailedModelName(action, model string, duration int, resolution, style string) string {
	baseName := fmt.Sprintf("vidu-%s-%s-%ds-%s", action, model, duration, resolution)

	// 如果是文生视频且有风格参数，加上风格
	if action == viduProvider.ViduActionText2Video && style != "" && style != "general" {
		baseName = fmt.Sprintf("%s-%s", baseName, style)
	}

	// 先尝试详细格式，如果不存在则退回简化格式
	return baseName
}

func (t *ViduTask) ShouldRetry(c *gin.Context, err *base.TaskError) bool {
	return false
}

func (t *ViduTask) GinResponse() {
	t.C.JSON(http.StatusOK, t.Response)
}

func (t *ViduTask) UpdateTaskStatus(ctx context.Context, taskChannelM map[int][]string, taskM map[string]*model.Task) error {
	logger.LogInfo(ctx, "vidu update task status")

	taskIds := make([]string, 0, len(taskM))
	for taskId := range taskM {
		taskIds = append(taskIds, taskId)
	}
	sort.Strings(taskIds)

	taskIdMaps := lo.Chunk(taskIds, 100)

	for _, taskIds := range taskIdMaps {
		for _, taskId := range taskIds {
			task := taskM[taskId]
			err := t.updateSingleTaskStatus(ctx, task)
			if err != nil {
				logger.LogError(ctx, fmt.Sprintf("update task %s error: %s", taskId, err.Error()))
			}
		}
	}

	return nil
}

func (t *ViduTask) updateSingleTaskStatus(ctx context.Context, task *model.Task) error {
	// 按渠道 ID 获取渠道，避免在后台轮询时依赖 Gin Context
	channel := model.ChannelGroup.GetChannel(task.ChannelId)
	if channel == nil {
		// 渠道缺失直接标记失败，避免死循环
		task.Status = model.TaskStatusFailure
		task.Progress = 100
		task.FailReason = fmt.Sprintf("获取渠道信息失败，请联系管理员，渠道ID：%d", task.ChannelId)
		if err := task.Update(); err != nil {
			logger.SysError(fmt.Sprintf("UpdateTask error: %v", err))
		}
		return fmt.Errorf("channel not found")
	}

	// 构造 Provider（后台场景不传 Gin 上下文）
	provider := providers.GetProvider(channel, nil)
	vidu, ok := provider.(*viduProvider.ViduProvider)
	if !ok {
		task.Status = model.TaskStatusFailure
		task.Progress = 100
		task.FailReason = "获取供应商失败，请联系管理员"
		if err := task.Update(); err != nil {
			logger.SysError(fmt.Sprintf("UpdateTask error: %v", err))
		}
		return fmt.Errorf("provider not found")
	}

	// 查询任务状态 - 使用新的官方接口
	resp, errWithCode := vidu.QueryCreations(task.TaskID)
	if errWithCode != nil {
		return errWithCode
	}

	// 映射并更新任务状态
	task.Progress = 0
	switch resp.State {
	case viduProvider.ViduStatusCreated:
		task.Status = model.TaskStatusSubmitted
		task.Progress = 10
	case viduProvider.ViduStatusQueueing:
		task.Status = model.TaskStatusQueued
		task.Progress = 20
	case viduProvider.ViduStatusProcessing:
		task.Status = model.TaskStatusInProgress
		task.Progress = 50
	case viduProvider.ViduStatusSuccess:
		task.Status = model.TaskStatusSuccess
		task.Progress = 100
	case viduProvider.ViduStatusFailed:
		task.Status = model.TaskStatusFailure
		task.Progress = 100
		task.FailReason = resp.ErrCode
		// 失败补偿配额
		quota := task.Quota
		if quota > 0 {
			if err := model.IncreaseUserQuota(task.UserId, quota); err != nil {
				logger.LogError(ctx, "fail to increase user quota: "+err.Error())
			}
			logContent := fmt.Sprintf("异步任务执行失败 %s，补偿 %s", task.TaskID, common.LogQuota(quota))
			model.RecordLog(task.UserId, model.LogTypeSystem, logContent)
		}
	default:
		task.Status = model.TaskStatusUnknown
	}

	// 保存响应数据
	if len(resp.Creations) > 0 {
		taskData, _ := json.Marshal(resp)
		task.Data = taskData
	}

	return task.Update()
}
