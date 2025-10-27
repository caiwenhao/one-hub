package gemini

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "one-api/common"
    "one-api/common/requester"
    "one-api/model"
    "one-api/providers/base"
    "one-api/providers/openai"
    "one-api/types"
    "strings"
)

type GeminiProviderFactory struct{}

// 创建 GeminiProvider
func (f GeminiProviderFactory) Create(channel *model.Channel) base.ProviderInterface {
    useOpenaiAPI := false
    useCodeExecution := false

	if channel.Plugin != nil {
		plugin := channel.Plugin.Data()
		if pWeb, ok := plugin["code_execution"]; ok {
			if enable, ok := pWeb["enable"].(bool); ok && enable {
				useCodeExecution = true
			}
		}

		if pWeb, ok := plugin["use_openai_api"]; ok {
			if enable, ok := pWeb["enable"].(bool); ok && enable {
				useOpenaiAPI = true
			}
		}
	}

    version := "v1beta"
    if channel.Other != "" {
        version = channel.Other
    }

    // 根据是否启用 OpenAI 兼容模式与 BaseURL 判断：
    // - 默认（未设置 BaseURL 或指向 Google 官方域名）：使用 Google 的 OpenAI 兼容路径 `/{version}/chat/completions`，错误解析保留 Gemini 形态
    // - 若 BaseURL 指向非 Google 域名：使用标准 OpenAI 路径 `/v1/chat/completions`，错误解析采用 OpenAI 形态
    baseURL := channel.GetBaseURL()
    useGoogleCompat := (baseURL == "" || strings.Contains(baseURL, "generativelanguage.googleapis.com"))

    cfg := getConfig(useOpenaiAPI, version, baseURL)

    // 选择错误处理器
    errHandler := RequestErrorHandle(channel.Key)
    if useOpenaiAPI && !useGoogleCompat {
        // 非 Google 官方域名时，假定为标准 OpenAI 风格错误
        errHandler = openai.RequestErrorHandle
    }

    provider := &GeminiProvider{
        OpenAIProvider: openai.OpenAIProvider{
            BaseProvider: base.BaseProvider{
                Config:    cfg,
                Channel:   channel,
                Requester: requester.NewHTTPRequester(*channel.Proxy, errHandler),
            },
            SupportStreamOptions: true,
        },
        UseOpenaiAPI:     useOpenaiAPI,
        UseCodeExecution: useCodeExecution,
    }

    return provider
}

type GeminiProvider struct {
    openai.OpenAIProvider
    UseOpenaiAPI     bool
    UseCodeExecution bool
}

func getConfig(useOpenaiAPI bool, version string, baseURL string) base.ProviderConfig {
    // 默认使用 Google 官方兼容端点
    cfg := base.ProviderConfig{
        BaseURL:           "https://generativelanguage.googleapis.com",
        ChatCompletions:   fmt.Sprintf("/%s/chat/completions", version),
        ModelList:         "/models",
        ImagesGenerations: "1",
    }

    // 若用户设置了自定义 BaseURL，则覆盖
    if baseURL != "" {
        cfg.BaseURL = baseURL
    }

    // 当启用 OpenAI API 且 BaseURL 指向非 Google 官方域名时，改为标准 OpenAI 路径
    if useOpenaiAPI && baseURL != "" && !strings.Contains(baseURL, "generativelanguage.googleapis.com") {
        cfg.ChatCompletions = "/v1/chat/completions"
    }

    return cfg
}

// 请求错误处理
func RequestErrorHandle(key string) requester.HttpErrorHandler {
	return func(resp *http.Response) *types.OpenAIError {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil
		}
		resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		geminiError := &GeminiErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(geminiError); err == nil {
			return errorHandle(geminiError, key)
		} else {
			geminiErrors := &GeminiErrors{}
			if err := json.Unmarshal(bodyBytes, geminiErrors); err == nil {
				return errorHandle(geminiErrors.Error(), key)
			}
		}

		return nil
	}
}

// 错误处理
func errorHandle(geminiError *GeminiErrorResponse, key string) *types.OpenAIError {
	if geminiError.ErrorInfo == nil || geminiError.ErrorInfo.Message == "" {
		return nil
	}

	cleaningError(geminiError.ErrorInfo, key)

	return &types.OpenAIError{
		Message: geminiError.ErrorInfo.Message,
		Type:    "gemini_error",
		Param:   geminiError.ErrorInfo.Status,
		Code:    geminiError.ErrorInfo.Code,
	}
}

func cleaningError(errorInfo *GeminiError, key string) {
	if key == "" {
		return
	}
	message := strings.Replace(errorInfo.Message, key, "xxxxx", 1)
	errorInfo.Message = message
}

func (p *GeminiProvider) GetFullRequestURL(requestURL string, modelName string) string {
	baseURL := strings.TrimSuffix(p.GetBaseURL(), "/")
	version := "v1beta"

	if p.Channel.Other != "" {
		version = p.Channel.Other
	}

	inputVersion := p.Context.Param("version")
	if inputVersion != "" {
		version = inputVersion
	}

	return fmt.Sprintf("%s/%s/models/%s:%s", baseURL, version, modelName, requestURL)

}

// 获取请求头
func (p *GeminiProvider) GetRequestHeaders() (headers map[string]string) {
    headers = make(map[string]string)
    p.CommonRequestHeaders(headers)
    headers["x-goog-api-key"] = p.Channel.Key

    return headers
}

// RelayModelAction 透传任意 models/<model>:<action> 原生请求（非流式）。
// 适配 Imagen 的 :predict 与 Veo 的 :predictLongRunning 初始化调用。
func (p *GeminiProvider) RelayModelAction(modelName, action string) (any, *types.OpenAIErrorWithStatusCode) {
    baseURL := strings.TrimSuffix(p.GetBaseURL(), "/")
    version := "v1beta"
    if p.Channel.Other != "" { version = p.Channel.Other }
    if v := p.Context.Param("version"); v != "" { version = v }

    fullRequestURL := fmt.Sprintf("%s/%s/models/%s:%s", baseURL, version, modelName, action)
    headers := p.GetRequestHeaders()

    body, ok := p.GetRawBody()
    if !ok {
        return nil, common.StringErrorWrapperLocal("request body not found", "request_body_not_found", http.StatusInternalServerError)
    }

    req, err := p.Requester.NewRequest(http.MethodPost, fullRequestURL, p.Requester.WithBody(body), p.Requester.WithHeader(headers))
    if err != nil {
        return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError)
    }
    defer req.Body.Close()

    // 动态解析为通用 map（避免定义所有变体的结构体）
    var resp any
    _, errWithCode := p.Requester.SendRequest(req, &resp, false)
    if errWithCode != nil {
        return nil, errWithCode
    }
    return resp, nil
}
