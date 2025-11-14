package relay

import (
    "encoding/json"
    "io"
    "net/http"
    "one-api/common"
    "one-api/common/config"
    "one-api/common/utils"
    "one-api/model"
    "one-api/providers"
    miniProvider "one-api/providers/minimaxi"
    "one-api/types"
    "strings"

	"github.com/gin-gonic/gin"
)

// MiniMaxAsyncQuery 查询 MiniMax 同步语音异步任务状态，仅作用于 /minimaxi 官方路由
func MiniMaxAsyncQuery(c *gin.Context) {
    taskID := strings.TrimSpace(c.Query("task_id"))
    if taskID == "" {
        err := common.StringErrorWrapperLocal("task_id is required", "invalid_request", http.StatusBadRequest)
        relayResponseWithOpenAIErr(c, err)
        return
    }

    // 支持平台任务ID：task_<ULID>/裸ULID/旧base36
    if strings.HasPrefix(strings.ToLower(taskID), utils.PlatformTaskPrefix) {
        if t, _ := model.GetTaskByPlatformTaskID(model.TaskPlatformMiniMax, c.GetInt("id"), utils.StripTaskPrefix(taskID)); t != nil {
            taskID = t.TaskID
        }
    } else if id, ok := utils.DecodePlatformTaskID(taskID); ok {
        if t, _ := model.GetTaskByID(id); t != nil && t.Platform == model.TaskPlatformMiniMax && t.UserId == c.GetInt("id") {
            taskID = t.TaskID
        }
    } else if utils.IsULID(taskID) {
        if t, _ := model.GetTaskByPlatformTaskID(model.TaskPlatformMiniMax, c.GetInt("id"), taskID); t != nil {
            taskID = t.TaskID
        }
    }

	// 该查询接口不需要 channel_id，默认按分组择优遍历可用的 MiniMax 渠道
	var channels []*model.Channel
	allChannels, dbErr := model.GetAllChannels()
	if dbErr != nil {
		errWithCode := common.ErrorWrapper(dbErr, "fetch_channels_failed", http.StatusInternalServerError)
		relayResponseWithOpenAIErr(c, errWithCode)
		return
	}
	group := strings.TrimSpace(c.GetString("token_group"))
	for _, ch := range allChannels {
		if ch == nil || ch.Status != config.ChannelStatusEnabled || ch.Type != config.ChannelTypeMiniMax {
			continue
		}
		if group != "" && !channelInGroup(ch, group) {
			continue
		}
		channels = append(channels, ch)
	}

	if len(channels) == 0 {
		err := common.StringErrorWrapperLocal("no available MiniMax channel", "channel_not_found", http.StatusServiceUnavailable)
		relayResponseWithOpenAIErr(c, err)
		return
	}

	var lastErr *types.OpenAIErrorWithStatusCode
	for _, channel := range channels {
		provider := providers.GetProvider(channel, c)
		if provider == nil {
			continue
		}
		mini, ok := provider.(*miniProvider.MiniMaxProvider)
		if !ok {
			continue
		}
		mini.SetContext(c)
        resp, errWithCode := mini.QuerySpeechAsync(taskID)
        if errWithCode != nil {
            lastErr = errWithCode
            continue
        }
        if resp == nil {
            continue
        }
        // 尝试拦截 JSON：将返回中的 task_id 统一改为平台ID（兼容 {task:{task_id}} 结构）
        if resp.Body != nil && strings.Contains(strings.ToLower(resp.Header.Get("Content-Type")), "application/json") {
            body, _ := io.ReadAll(resp.Body)
            _ = resp.Body.Close()
            var m map[string]any
            if json.Unmarshal(body, &m) == nil {
                // 提取上游ID并替换为平台ID
                upID := ""
                if v, ok := m["task_id"].(string); ok { upID = strings.TrimSpace(v) }
                if upID == "" {
                    if t, ok := m["task"].(map[string]any); ok {
                        if v, ok2 := t["task_id"].(string); ok2 { upID = strings.TrimSpace(v) }
                    }
                }
                if upID != "" {
                    // 查本地映射
                    if t, _ := model.GetTaskByTaskId(model.TaskPlatformMiniMax, c.GetInt("id"), upID); t != nil {
                        pid := utils.AddTaskPrefix(t.PlatformTaskID)
                        if _, ok := m["task_id"]; ok {
                            m["task_id"] = pid
                        }
                        if taskObj, ok := m["task"].(map[string]any); ok {
                            taskObj["task_id"] = pid
                            m["task"] = taskObj
                        }
                        patched, _ := json.Marshal(m)
                        c.Data(http.StatusOK, "application/json", patched)
                        return
                    }
                }
            }
            // 回退：原样
            c.Data(resp.StatusCode, "application/json", body)
            return
        }
        if err := responseMultipart(c, resp); err != nil {
            relayResponseWithOpenAIErr(c, err)
            return
        }
        return
	}

	if lastErr != nil {
		relayResponseWithOpenAIErr(c, lastErr)
		return
	}

	err := common.StringErrorWrapperLocal("query failed", "upstream_error", http.StatusBadGateway)
	relayResponseWithOpenAIErr(c, err)
}

func channelInGroup(channel *model.Channel, group string) bool {
	if channel == nil {
		return false
	}
	if strings.TrimSpace(channel.Group) == "" {
		return true
	}
	for _, item := range strings.Split(channel.Group, ",") {
		if strings.TrimSpace(item) == group {
			return true
		}
	}
	return false
}
