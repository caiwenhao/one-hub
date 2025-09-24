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

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

type ViduTask struct {
	base.TaskBase
	Action   string
	Request  *viduProvider.ViduTaskRequest
	Provider *viduProvider.ViduProvider
}

func (t *ViduTask) HandleError(err *base.TaskError) {
	StringError(t.C, err.StatusCode, err.Code, err.Message)
}

func (t *ViduTask) Init() *base.TaskError {
	t.Action = t.C.Param("action")

	// 解析请求体
	if err := common.UnmarshalBodyReusable(t.C, &t.Request); err != nil {
		return base.StringTaskError(http.StatusBadRequest, "invalid_request", err.Error(), true)
	}

	err := t.actionValidate()
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
		if len(t.Request.Images) == 0 {
			return fmt.Errorf("images is required for img2video")
		}
	case viduProvider.ViduActionReference2Video:
		if len(t.Request.ReferenceVideos) == 0 {
			return fmt.Errorf("reference_videos is required for reference2video")
		}
	case viduProvider.ViduActionStartEnd2Video:
		if t.Request.StartImage == "" || t.Request.EndImage == "" {
			return fmt.Errorf("start_image and end_image are required for start-end2video")
		}
	case viduProvider.ViduActionText2Video:
		if t.Request.Prompt == "" {
			return fmt.Errorf("prompt is required for text2video")
		}
	default:
		return fmt.Errorf("unsupported action: %s", t.Action)
	}

	// 设置默认模型
	if t.Request.Model == "" {
		t.Request.Model = "vidu1.5"
	}

	// 构造原始模型名称用于计费
	t.OriginalModel = fmt.Sprintf("vidu-%s-%s", t.Action, t.Request.Model)
	if t.Request.Duration != nil {
		t.OriginalModel = fmt.Sprintf("%s-%ds", t.OriginalModel, *t.Request.Duration)
	}

	return nil
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

    // 查询任务状态
    resp, errWithCode := vidu.Query(task.TaskID)
    if errWithCode != nil {
        return errWithCode
    }

    // 映射并更新任务状态
    task.Progress = 0
    switch resp.Status {
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
        task.FailReason = resp.Message
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
    if resp.Data != nil {
        taskData, _ := json.Marshal(resp)
        task.Data = taskData
    }

    return task.Update()
}
