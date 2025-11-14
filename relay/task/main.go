package task

import (
    "fmt"
    "math"
    "net/http"
    "one-api/common/config"
    "one-api/common/logger"
    "one-api/metrics"
    "one-api/model"
    "one-api/common/utils"
    "one-api/relay/relay_util"
    "one-api/relay/task/base"
    "one-api/types"
    "strings"
    "reflect"

	"github.com/gin-gonic/gin"
)

func RelayTaskSubmit(c *gin.Context) {
	var taskErr *base.TaskError
	taskAdaptor, err := GetTaskAdaptor(GetRelayMode(c), c)
	if err != nil {
		taskErr = base.StringTaskError(http.StatusBadRequest, "adaptor_not_found", "adaptor not found", true)
		c.JSON(http.StatusBadRequest, taskErr)
		return
	}

	taskErr = taskAdaptor.Init()
	if taskErr != nil {
		taskAdaptor.HandleError(taskErr)
		return
	}

	taskErr = taskAdaptor.SetProvider()
	if taskErr != nil {
		taskAdaptor.HandleError(taskErr)
		return
	}

	estimatedPromptTokens := 1000
	if raw, exists := c.Get("async_estimated_prompt_tokens"); exists {
		switch value := raw.(type) {
		case int:
			if value > 0 {
				estimatedPromptTokens = value
			}
		case int64:
			if value > 0 && value < math.MaxInt32 {
				estimatedPromptTokens = int(value)
			}
		case float64:
			if value > 0 && value < math.MaxInt32 {
				estimatedPromptTokens = int(value)
			}
		}
	}

	quotaInstance := relay_util.NewQuota(c, taskAdaptor.GetModelName(), estimatedPromptTokens)
	if errWithOA := quotaInstance.PreQuotaConsumption(); errWithOA != nil {
		taskAdaptor.HandleError(base.OpenAIErrToTaskErr(errWithOA))
		return
	}

	taskErr = taskAdaptor.Relay()
	if taskErr == nil {
		CompletedTask(quotaInstance, taskAdaptor, c)
		// 返回结果
		taskAdaptor.GinResponse()
		metrics.RecordProvider(c, 200)
		return
	}

	quotaInstance.Undo(c)

	retryTimes := config.RetryTimes

	if !taskAdaptor.ShouldRetry(c, taskErr) {
		logger.LogError(c.Request.Context(), fmt.Sprintf("relay error happen, status code is %d, won't retry in this case", taskErr.StatusCode))
		retryTimes = 0
	}

	channel := taskAdaptor.GetProvider().GetChannel()
	for i := retryTimes; i > 0; i-- {
		model.ChannelGroup.SetCooldowns(channel.Id, taskAdaptor.GetModelName())
		taskErr = taskAdaptor.SetProvider()
		if taskErr != nil {
			continue
		}

		channel = taskAdaptor.GetProvider().GetChannel()
		logger.LogError(c.Request.Context(), fmt.Sprintf("using channel #%d(%s) to retry (remain times %d)", channel.Id, channel.Name, i))

		taskErr = taskAdaptor.Relay()
		if taskErr == nil {
			go CompletedTask(quotaInstance, taskAdaptor, c)
			return
		}

		quotaInstance.Undo(c)
		if !taskAdaptor.ShouldRetry(c, taskErr) {
			break
		}

	}

	if taskErr != nil {
		taskAdaptor.HandleError(taskErr)
	}

}

func CompletedTask(quotaInstance *relay_util.Quota, taskAdaptor base.TaskInterface, c *gin.Context) {
    // 先入库拿到自增主键，生成平台任务ID；再消费配额并回填日志中 task_id
    task := taskAdaptor.GetTask()
    task.Quota = int(quotaInstance.GetInputRatio() * 1000)
    task.BillingGroupRatio = quotaInstance.GetGroupRatio()
    task.BillingModel = taskAdaptor.GetModelName()

    // 设置平台任务ID（ULID）
    if strings.TrimSpace(task.PlatformTaskID) == "" {
        task.PlatformTaskID = utils.NewPlatformULID()
    }

    if err := task.Insert(); err != nil {
        logger.SysError(fmt.Sprintf("task insert error: %s", err.Error()))
    }

    // 在日志中写入 platform_task_id / upstream_task_id，并覆盖返回体
    if task.ID > 0 {
        platformID := utils.AddTaskPrefix(task.PlatformTaskID)
        quotaInstance.SetTaskIDs(platformID, task.TaskID)
        // 尝试覆盖响应体中的 ID/task_id 字段为平台ID
        if resp := taskAdaptor.GetResponse(); resp != nil {
            if m, ok := resp.(map[string]any); ok {
                if _, exists := m["task_id"]; exists { m["task_id"] = platformID }
                if _, exists := m["id"]; exists { m["id"] = platformID }
                taskAdaptor.SetResponse(m)
            } else {
                rv := reflect.ValueOf(resp)
                if rv.Kind() == reflect.Pointer { rv = rv.Elem() }
                if rv.IsValid() && rv.CanAddr() {
                    if f := rv.FieldByName("TaskID"); f.IsValid() && f.CanSet() && f.Kind() == reflect.String {
                        f.SetString(platformID)
                    }
                    if f := rv.FieldByName("ID"); f.IsValid() && f.CanSet() && f.Kind() == reflect.String {
                        f.SetString(platformID)
                    }
                }
                taskAdaptor.SetResponse(resp)
            }
        }
    }

    // 在拿到平台ID后再进行消费，以便日志带上两个ID
    quotaInstance.Consume(c, &types.Usage{CompletionTokens: 0, PromptTokens: 1, TotalTokens: 1}, false)

    // 激活任务
    ActivateUpdateTaskBulk()
}

func GetRelayMode(c *gin.Context) int {
	relayMode := config.RelayModeUnknown
	path := c.Request.URL.Path
	if strings.HasPrefix(path, "/suno") {
		relayMode = config.RelayModeSuno
	} else if strings.HasPrefix(path, "/kling") {
		relayMode = config.RelayModeKling
	} else if strings.HasPrefix(path, "/vidu") {
		relayMode = config.RelayModeVidu
	} else if strings.HasPrefix(path, "/volcark") {
		relayMode = config.RelayModeVolcArkVideo
    } else if strings.HasPrefix(path, "/minimaxi") || strings.HasPrefix(path, "/minimax") || strings.HasPrefix(path, "/v1/video_generation") || strings.HasPrefix(path, "/v1/query/video_generation") {
        relayMode = config.RelayModeMiniMaxVideo
    }

	return relayMode
}
