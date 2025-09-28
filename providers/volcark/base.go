package volcark

import (
	"one-api/common/requester"
	"one-api/model"
	"one-api/providers/base"
	"one-api/providers/openai"
)

// VolcArkProviderFactory 用于创建火山方舟渠道的 Provider
// 该渠道兼容 OpenAI 协议，通过调整基础路径即可对接
// https://ark.cn-beijing.volces.com/api/v3/
type VolcArkProviderFactory struct{}

// Create 构建具体 Provider 实例
func (f VolcArkProviderFactory) Create(channel *model.Channel) base.ProviderInterface {
	proxy := ""
	if channel.Proxy != nil {
		proxy = *channel.Proxy
	}

	return &VolcArkProvider{
		OpenAIProvider: openai.OpenAIProvider{
			BaseProvider: base.BaseProvider{
				Config:          getVolcArkConfig(channel),
				Channel:         channel,
				Requester:       requester.NewHTTPRequester(proxy, openai.RequestErrorHandle),
				SupportResponse: true,
			},
			SupportStreamOptions: true,
			BalanceAction:        false,
		},
	}
}

type VolcArkProvider struct {
	openai.OpenAIProvider
}

func getVolcArkConfig(_ *model.Channel) base.ProviderConfig {
	config := base.ProviderConfig{
		BaseURL:             "https://ark.cn-beijing.volces.com",
		Completions:         "/api/v3/completions",
		ChatCompletions:     "/api/v3/chat/completions",
		Embeddings:          "/api/v3/embeddings",
		AudioSpeech:         "/api/v3/audio/speech",
		AudioTranscriptions: "/api/v3/audio/transcriptions",
		AudioTranslations:   "/api/v3/audio/translations",
		ImagesGenerations:   "/api/v3/images/generations",
		ImagesEdit:          "/api/v3/images/edits",
		ImagesVariations:    "/api/v3/images/variations",
		Moderation:          "/api/v3/moderations",
		ModelList:           "/api/v3/models",
		Responses:           "/api/v3/responses",
	}

	return config
}
