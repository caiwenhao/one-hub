package minimaxi

import (
    "bytes"
    "encoding/json"
    "io"
    "net/http"
    "one-api/common/requester"
    "one-api/model"
    "one-api/providers/base"
    "one-api/providers/openai"
    "one-api/types"
    "strings"

    "github.com/gin-gonic/gin"
)

type MiniMaxProviderFactory struct{}

// 创建 MiniMaxProvider
func (f MiniMaxProviderFactory) Create(channel *model.Channel) base.ProviderInterface {
	proxy := ""
	if channel.Proxy != nil {
		proxy = *channel.Proxy
	}

    provider := &MiniMaxProvider{
        OpenAIProvider: openai.OpenAIProvider{
            BaseProvider: base.BaseProvider{
                Config:    getConfig(channel),
                Channel:   channel,
                Requester: requester.NewHTTPRequester(proxy, requestErrorHandle),
            },
        },
    }

	provider.VideoClient = newMiniMaxVideoClient(channel)

	return provider
}

type MiniMaxProvider struct {
	openai.OpenAIProvider
	VideoClient *MiniMaxVideoClient
}

func (p *MiniMaxProvider) SetContext(c *gin.Context) {
	p.OpenAIProvider.SetContext(c)
	if p.VideoClient != nil {
		p.VideoClient.SetContext(c)
		if p.VideoClient.Requester != nil && c != nil {
			p.VideoClient.Requester.Context = c.Request.Context()
		}
	}
}

func (p *MiniMaxProvider) GetVideoClient() *MiniMaxVideoClient {
	return p.VideoClient
}

func getConfig(channel *model.Channel) base.ProviderConfig {
    // 默认使用 MiniMax 官方配置
    cfg := base.ProviderConfig{
        BaseURL:               "https://api.minimaxi.com",
        ChatCompletions:       "/v1/chat/completions",
        AudioSpeech:           "/v1/t2a_v2",
        AudioSpeechAsync:      "/v1/t2a_async_v2",
        AudioSpeechAsyncQuery: "/v1/query/t2a_async_query_v2",
        ModelList:             "/v1/models",
    }

    // 读取 custom_parameter 顶层或 audio 块的 upstream，ppinfra 时切换到 PPInfra 聚合
    if channel != nil && channel.CustomParameter != nil && *channel.CustomParameter != "" {
        var payload map[string]any
        if err := json.Unmarshal([]byte(*channel.CustomParameter), &payload); err == nil {
            upstream := ""
            if v, ok := payload["audio"]; ok {
                if audioMap, ok2 := v.(map[string]any); ok2 {
                    if up, ok3 := audioMap["upstream"].(string); ok3 {
                        upstream = up
                    }
                }
            }
            if upstream == "" {
                if up, ok := payload["upstream"].(string); ok {
                    upstream = up
                }
            }
            if strings.ToLower(strings.TrimSpace(upstream)) == "ppinfra" {
                cfg.BaseURL = "https://api.ppinfra.com"
                // PPInfra 同步 TTS 路径
                cfg.AudioSpeech = "/v3/minimax-speech-02-hd"
                // PPInfra 异步 TTS 路径
                cfg.AudioSpeechAsync = "/v3/async/minimax-speech-02-hd"
                // PPInfra 异步查询统一任务查询接口
                cfg.AudioSpeechAsyncQuery = "/v3/async/task-result"
            }
        }
    }

    // 若未声明 upstream，但渠道 BaseURL 指向 PPInfra，也默认切换语音路径
    if channel != nil {
        base := strings.ToLower(strings.TrimSpace(channel.GetBaseURL()))
        if base != "" && strings.Contains(base, "ppinfra") {
            cfg.BaseURL = channel.GetBaseURL()
            cfg.AudioSpeech = "/v3/minimax-speech-02-hd"
            cfg.AudioSpeechAsync = "/v3/async/minimax-speech-02-hd"
            cfg.AudioSpeechAsyncQuery = "/v3/async/task-result"
        }
    }

    return cfg
}

// 请求错误处理
func requestErrorHandle(resp *http.Response) *types.OpenAIError {
	// 读取响应体，支持多次反序列化尝试
	bodyBytes, err := io.ReadAll(resp.Body)
	if err == nil {
		resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	// 先尝试 OpenAI 标准错误（提升与 OpenAI 兼容度）
	var openaiErrResp types.OpenAIErrorResponse
	if err := json.Unmarshal(bodyBytes, &openaiErrResp); err == nil && openaiErrResp.Error.Message != "" {
		return &openaiErrResp.Error
	}

	// 再尝试 minimaxi base_resp 错误
	var wrap MiniMaxBaseResp
	if err := json.Unmarshal(bodyBytes, &wrap); err == nil {
		return errorHandle(&wrap.BaseResp)
	}

	return nil
}

// 错误处理
func errorHandle(minimaxError *BaseResp) *types.OpenAIError {
	if minimaxError.StatusCode == 0 {
		return nil
	}
	return &types.OpenAIError{
		Message: minimaxError.StatusMsg,
		Type:    "minimax_error",
		Code:    minimaxError.StatusCode,
	}
}
