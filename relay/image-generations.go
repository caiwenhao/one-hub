package relay

import (
    "io"
    "encoding/json"
    "net/http"
    "one-api/common"
    "one-api/common/config"
    // "one-api/common/logger"
    "one-api/common/utils"
    "one-api/model"
    providersBase "one-api/providers/base"
    "one-api/types"
    "strings"
    "github.com/gin-gonic/gin"
)

type relayImageGenerations struct {
	relayBase
	request types.ImageRequest
}

func NewRelayImageGenerations(c *gin.Context) *relayImageGenerations {
	relay := &relayImageGenerations{}
	relay.c = c
	return relay
}

func (r *relayImageGenerations) setRequest() error {
	if err := common.UnmarshalBodyReusable(r.c, &r.request); err != nil {
		return err
	}

	if r.request.Model == "" {
		r.request.Model = "dall-e-2"
	}

	if r.request.N == 0 {
		r.request.N = 1
	}

	if strings.HasPrefix(r.request.Model, "dall-e") {
		if r.request.Size == "" {
			r.request.Size = "1024x1024"
		}

		if r.request.Quality == "" {
			r.request.Quality = "standard"
		}
	}

	r.setOriginalModel(r.request.Model)

	return nil
}

func (r *relayImageGenerations) getPromptTokens() (int, error) {
	return common.CountTokenImage(r.request)
}

func (r *relayImageGenerations) send() (err *types.OpenAIErrorWithStatusCode, done bool) {
    // 若为 NewAPI 渠道：直透上游 JSON（供应商风格：{code,data:[{status,task_id}]})
    if r.provider != nil && r.provider.GetChannel() != nil && r.provider.GetChannel().Type == config.ChannelTypeNewAPI {
        // 与官方计费对齐：若该模型在价格表为 tokens 计费，则按每张 1290 image tokens 计入输出
        // 说明：NewAPI 路由为异步任务风格，无法在响应体拿到最终图片张数，此处按请求的 n 补记
        if price := model.PricingInstance.GetPrice(r.modelName); price != nil && price.Type == model.TokensPriceType {
            if u := r.provider.GetUsage(); u != nil {
                u.CompletionTokens += r.request.N * 1290
                u.TotalTokens = u.PromptTokens + u.CompletionTokens
            }
        }

        base := strings.TrimSuffix(r.provider.GetChannel().GetBaseURL(), "/")
        if base == "" {
            base = "https://api.openai.com" // 兜底，通常 NewAPI 会配置第三方域名
        }
        fullURL := base + "/v1/images/generations"

        headers := r.provider.GetRequestHeaders()
        if headers == nil {
            headers = map[string]string{}
        }
        headers["Content-Type"] = "application/json"

        // 使用结构体请求体（与上游字段一致）
        req, e := r.provider.GetRequester().NewRequest(http.MethodPost, fullURL,
            r.provider.GetRequester().WithBody(r.request),
            r.provider.GetRequester().WithHeader(headers),
        )
        if e != nil {
            err = common.ErrorWrapper(e, "new_request_failed", http.StatusInternalServerError)
            done = true
            return
        }
        if req.Body != nil {
            defer req.Body.Close()
        }

    // 直透响应
    resp, errWith := r.provider.GetRequester().SendRequestRaw(req)
    if errWith != nil {
            err = errWith
            done = true
            return
        }
    defer resp.Body.Close()

    // 尝试包装 JSON 的 task_id → 平台ID
    body, _ := io.ReadAll(resp.Body)
    var vendor map[string]any
    if json.Unmarshal(body, &vendor) == nil {
        if arr, ok := vendor["data"].([]any); ok && len(arr) > 0 {
            if first, okm := arr[0].(map[string]any); okm {
                if upID, ok2 := first["task_id"].(string); ok2 && strings.TrimSpace(upID) != "" {
                    // 仅做映射，平台统一归到 Sora 平台下（图像）
                    task := &model.Task{ PlatformTaskID: utils.NewPlatformULID(), TaskID: upID, Platform: model.TaskPlatformSora, UserId: r.c.GetInt("id"), TokenID: r.c.GetInt("token_id"), ChannelId: r.provider.GetChannel().Id, Action: "image.generate.vendor", Status: model.TaskStatusSubmitted, SubmitTime: r.c.GetTime("requestStartTime").Unix(), CreatedAt: r.c.GetTime("requestStartTime").Unix(), UpdatedAt: r.c.GetTime("requestStartTime").Unix(), }
                    _ = model.CreateTask(task)
                    first["task_id"] = utils.AddTaskPrefix(task.PlatformTaskID)
                    vendor["data"] = arr
                    patched, _ := json.Marshal(vendor)
                    r.c.Data(resp.StatusCode, "application/json", patched)
                    done = true
                    return
                }
            }
        }
    }

    for key, values := range resp.Header { for _, v := range values { r.c.Writer.Header().Add(key, v) } }
    r.c.Status(resp.StatusCode)
    r.c.Writer.Write(body)
    done = true
    return
    }

    // 其他渠道：按 OpenAI 标准处理
    provider, ok := r.provider.(providersBase.ImageGenerationsInterface)
    if !ok {
        err = common.StringErrorWrapperLocal("channel not implemented", "channel_error", http.StatusServiceUnavailable)
        done = true
        return
    }

    r.request.Model = r.modelName

    response, err := provider.CreateImageGenerations(&r.request)
    if err != nil {
        return
    }
    err = responseJsonClient(r.c, response)

    if err != nil {
        done = true
    }

    return
}
