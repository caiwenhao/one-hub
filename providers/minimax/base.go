package minimax

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
				Config:    getConfig(),
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

func getConfig() base.ProviderConfig {
    return base.ProviderConfig{
        // 对齐 minimaxi 官方 OpenAI 兼容入口
        BaseURL:         "https://api.minimaxi.com",
        ChatCompletions: "/v1/chat/completions",
        // 同步 TTS：/v1/t2a_v2；异步长文本：/v1/t2a_async_v2（如需可在渠道自定义路径覆盖）
        AudioSpeech:     "/v1/t2a_v2",
        // Embeddings:      "/v1/embeddings",
        // ModelList:       "/v1/models",
    }
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
