package relay

import (
    "encoding/json"
    "fmt"
    "io"
    "strings"

    "one-api/common"
    "one-api/common/config"
    "one-api/common/utils"
    "one-api/model"
    miniProvider "one-api/providers/minimaxi"
    "one-api/types"

    "github.com/gin-gonic/gin"
)

type relayMiniMaxAsync struct {
	relayBase
	request types.MiniMaxAsyncSpeechRequest
	rawBody []byte
}

func NewRelayMiniMaxAsync(c *gin.Context) *relayMiniMaxAsync {
	relay := &relayMiniMaxAsync{}
	relay.c = c
	return relay
}

func (r *relayMiniMaxAsync) setRequest() error {
	if err := common.UnmarshalBodyReusable(r.c, &r.request); err != nil {
		return err
	}

	if raw, exists := r.c.Get(config.GinRequestBodyKey); exists {
		if bytes, ok := raw.([]byte); ok {
			r.rawBody = bytes
		}
	}

	r.c.Set("allow_channel_type", []int{config.ChannelTypeMiniMax})

	if strings.TrimSpace(r.request.Model) == "" {
		return fmt.Errorf("model is required")
	}

	textProvided := r.request.Text != nil && strings.TrimSpace(*r.request.Text) != ""
	fileProvided := r.request.TextFileID != nil && len(*r.request.TextFileID) > 0
	if !textProvided && !fileProvided {
		return fmt.Errorf("either text or text_file_id is required")
	}

	if r.request.VoiceSetting == nil || strings.TrimSpace(r.request.VoiceSetting.VoiceID) == "" {
		return fmt.Errorf("voice_setting.voice_id is required")
	}

	r.setOriginalModel(r.request.Model)

	return nil
}

func (r *relayMiniMaxAsync) getPromptTokens() (int, error) {
	if r.request.Text != nil {
		return len(*r.request.Text), nil
	}
	return 0, nil
}

func (r *relayMiniMaxAsync) send() (err *types.OpenAIErrorWithStatusCode, done bool) {
    mini, ok := r.provider.(*miniProvider.MiniMaxProvider)
	if !ok {
		err = common.StringErrorWrapperLocal("invalid minimaxi provider", "channel_error", 503)
		done = true
		return
	}

	r.request.Model = r.modelName

    resp, errWithCode := mini.CreateSpeechAsync(&r.request, r.rawBody)
    if errWithCode != nil {
        return errWithCode, true
    }

    // 尝试拦截 JSON：将上游 task_id → 平台 task_<ULID>，并落库
    if resp != nil && resp.Body != nil && strings.Contains(strings.ToLower(resp.Header.Get("Content-Type")), "application/json") {
        body, _ := io.ReadAll(resp.Body)
        _ = resp.Body.Close()

        // 尝试两种结构：1) 官方 base_resp；2) 供应商 {task:{task_id,...}}
        // 提取上游 task_id
        extractUpstream := func(m map[string]any) (path string, upID string) {
            if v, ok := m["task_id"].(string); ok && strings.TrimSpace(v) != "" { return "task_id", strings.TrimSpace(v) }
            if t, ok := m["task"].(map[string]any); ok {
                if v, ok2 := t["task_id"].(string); ok2 && strings.TrimSpace(v) != "" { return "task.task_id", strings.TrimSpace(v) }
            }
            return "", ""
        }

        var generic map[string]any
        if json.Unmarshal(body, &generic) == nil {
            path, upstreamID := extractUpstream(generic)
            if upstreamID == "" {
                // 兼容官方结构（types 反序列化）
                var payload types.MiniMaxAsyncSpeechResponse
                if json.Unmarshal(body, &payload) == nil && payload.BaseResp.StatusCode == 0 {
                    upstreamID = payload.TaskID.String()
                    if upstreamID != "" {
                        // 落库/改写
                        task := &model.Task{ PlatformTaskID: utils.NewPlatformULID(), TaskID: upstreamID, Platform: model.TaskPlatformMiniMax, UserId: r.c.GetInt("id"), TokenID: r.c.GetInt("token_id"), ChannelId: r.c.GetInt("channel_id"), Action: "audio.t2a_async", Status: model.TaskStatusSubmitted, SubmitTime: r.c.GetTime("requestStartTime").Unix(), CreatedAt: r.c.GetTime("requestStartTime").Unix(), UpdatedAt: r.c.GetTime("requestStartTime").Unix(), }
                        _ = task.Insert()
                        payload.TaskID = types.StringOrNumber(utils.AddTaskPrefix(task.PlatformTaskID))
                        wrapped, _ := json.Marshal(payload)
                        r.c.Data(resp.StatusCode, "application/json", wrapped)
                        return nil, true
                    }
                }
            } else {
                // 供应商结构，落库并改写
                // 若已存在映射就复用
                task, _ := model.GetTaskByTaskId(model.TaskPlatformMiniMax, r.c.GetInt("id"), upstreamID)
                if task == nil {
                    task = &model.Task{ PlatformTaskID: utils.NewPlatformULID(), TaskID: upstreamID, Platform: model.TaskPlatformMiniMax, UserId: r.c.GetInt("id"), TokenID: r.c.GetInt("token_id"), ChannelId: r.c.GetInt("channel_id"), Action: "audio.t2a_async", Status: model.TaskStatusSubmitted, SubmitTime: r.c.GetTime("requestStartTime").Unix(), CreatedAt: r.c.GetTime("requestStartTime").Unix(), UpdatedAt: r.c.GetTime("requestStartTime").Unix(), }
                    _ = task.Insert()
                }
                platformID := utils.AddTaskPrefix(task.PlatformTaskID)
                if path == "task_id" {
                    generic["task_id"] = platformID
                } else {
                    t := generic["task"].(map[string]any)
                    t["task_id"] = platformID
                    generic["task"] = t
                }
                wrapped, _ := json.Marshal(generic)
                r.c.Data(resp.StatusCode, "application/json", wrapped)
                return nil, true
            }
        }
        // 回退：原样返回
        r.c.Data(resp.StatusCode, "application/json", body)
        return nil, true
    }

    if err = responseMultipart(r.c, resp); err != nil {
        done = true
    }
    return nil, done
}

func (r *relayMiniMaxAsync) getRequest() any {
	return &r.request
}
