package minimaxi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"one-api/common/logger"
	"one-api/common/requester"
	"one-api/model"
	"one-api/providers/base"
	"one-api/types"
	"strings"
)

const (
    MiniMaxVideoUpstreamOfficial = "official"
    MiniMaxVideoUpstreamPPInfra  = "ppinfra"
    MiniMaxVideoUpstreamPolloi   = "polloi"
)

// MiniMaxVideoConfig 描述视频能力所需的配置项
// 通过渠道自定义参数中的 video 字段进行覆盖。
type MiniMaxVideoConfig struct {
	Upstream              string            `json:"upstream,omitempty"`
	APIKey                string            `json:"api_key,omitempty"`
	BaseURL               string            `json:"base_url,omitempty"`
	SubmitPath            string            `json:"submit_path,omitempty"`
	SubmitPathTemplate    string            `json:"submit_path_template,omitempty"`
	QueryPath             string            `json:"query_path,omitempty"`
	QueryPathTemplate     string            `json:"query_path_template,omitempty"`
	TemplatePath          string            `json:"template_path,omitempty"`
	FileRetrievePath      string            `json:"file_retrieve_path,omitempty"`
	AuthHeader            string            `json:"auth_header,omitempty"`
	AuthScheme            string            `json:"auth_scheme,omitempty"`
	DefaultCallbackURL    string            `json:"callback_url,omitempty"`
	EnablePromptExpansion *bool             `json:"enable_prompt_expansion,omitempty"`
	ExtraHeaders          map[string]string `json:"extra_headers,omitempty"`
}

// MiniMaxVideoClient 负责同 MiniMax 官方或第三方上游交互
// 复用 BaseProvider 的能力，用于构造请求、注入上下文等。
type MiniMaxVideoClient struct {
	base.BaseProvider
	config MiniMaxVideoConfig
}

// MiniMaxSubjectReference 描述主体参考配置
type MiniMaxSubjectReference struct {
	Type  string   `json:"type,omitempty"`
	Image []string `json:"image,omitempty"`
}

func newMiniMaxVideoClient(channel *model.Channel) *MiniMaxVideoClient {
	cfg := loadMiniMaxVideoConfig(channel)

	proxy := ""
	if channel.Proxy != nil {
		proxy = *channel.Proxy
	}

	requester := requester.NewHTTPRequester(proxy, miniMaxVideoErrorHandle)
	requester.IsOpenAI = false

	client := &MiniMaxVideoClient{
		BaseProvider: base.BaseProvider{
			Config: base.ProviderConfig{
				BaseURL: cfg.BaseURL,
			},
			Channel:   channel,
			Requester: requester,
		},
		config: cfg,
	}

	return client
}

// loadMiniMaxVideoConfig 读取渠道配置，合并默认值
func loadMiniMaxVideoConfig(channel *model.Channel) MiniMaxVideoConfig {
	cfg := MiniMaxVideoConfig{
		Upstream:           MiniMaxVideoUpstreamOfficial,
		BaseURL:            "https://api.minimaxi.com",
		SubmitPath:         "/v1/video_generation",
		QueryPath:          "/v1/query/video_generation",
		TemplatePath:       "/v1/video_template_generation",
		FileRetrievePath:   "/v1/files/retrieve",
		AuthHeader:         "Authorization",
		AuthScheme:         "Bearer",
		ExtraHeaders:       map[string]string{},
		DefaultCallbackURL: "",
	}

	cfg.APIKey = channel.Key

	raw := channel.GetCustomParameter()
	if raw == "" {
		return cfg
	}

	var payload map[string]json.RawMessage
	if err := json.Unmarshal([]byte(raw), &payload); err != nil {
		logger.SysError(fmt.Sprintf("MiniMax video config decode failed: %v", err))
		return cfg
	}

	// 顶层 upstream 兜底：如果 video 块未声明 upstream，则读取顶层 upstream
	if topUpRaw, ok := payload["upstream"]; ok && len(topUpRaw) > 0 {
		var topUp string
		if err := json.Unmarshal(topUpRaw, &topUp); err == nil && strings.TrimSpace(topUp) != "" {
			cfg.Upstream = strings.ToLower(strings.TrimSpace(topUp))
		}
	}

	if apiKeyRaw, ok := payload["api_key"]; ok && len(apiKeyRaw) > 0 {
		var topKey string
		if err := json.Unmarshal(apiKeyRaw, &topKey); err == nil && strings.TrimSpace(topKey) != "" {
			cfg.APIKey = strings.TrimSpace(topKey)
		}
	}

	if baseURLRaw, ok := payload["base_url"]; ok && len(baseURLRaw) > 0 {
		var topBase string
		if err := json.Unmarshal(baseURLRaw, &topBase); err == nil && strings.TrimSpace(topBase) != "" {
			cfg.BaseURL = strings.TrimSpace(topBase)
		}
	}

	if authHeaderRaw, ok := payload["auth_header"]; ok && len(authHeaderRaw) > 0 {
		var topHeader string
		if err := json.Unmarshal(authHeaderRaw, &topHeader); err == nil && strings.TrimSpace(topHeader) != "" {
			cfg.AuthHeader = strings.TrimSpace(topHeader)
		}
	}

	if authSchemeRaw, ok := payload["auth_scheme"]; ok && len(authSchemeRaw) > 0 {
		var topScheme string
		if err := json.Unmarshal(authSchemeRaw, &topScheme); err == nil && strings.TrimSpace(topScheme) != "" {
			cfg.AuthScheme = strings.TrimSpace(topScheme)
		}
	}

	if extraHeadersRaw, ok := payload["extra_headers"]; ok && len(extraHeadersRaw) > 0 {
		var topHeaders map[string]string
		if err := json.Unmarshal(extraHeadersRaw, &topHeaders); err == nil && len(topHeaders) > 0 {
			if cfg.ExtraHeaders == nil {
				cfg.ExtraHeaders = map[string]string{}
			}
			for k, v := range topHeaders {
				cfg.ExtraHeaders[k] = v
			}
		}
	}

	if videoRaw, ok := payload["video"]; ok {
		var custom MiniMaxVideoConfig
		if err := json.Unmarshal(videoRaw, &custom); err != nil {
			logger.SysError(fmt.Sprintf("MiniMax video custom config decode failed: %v", err))
		} else {
			mergeMiniMaxVideoConfig(&cfg, &custom)
		}
	}

	return cfg
}

// mergeMiniMaxVideoConfig 将用户配置合并到默认配置中
func mergeMiniMaxVideoConfig(dst, src *MiniMaxVideoConfig) {
	if src.Upstream != "" {
		dst.Upstream = strings.ToLower(src.Upstream)
	}
	if src.APIKey != "" {
		dst.APIKey = src.APIKey
	}
	if src.BaseURL != "" {
		dst.BaseURL = src.BaseURL
	}
	if src.SubmitPath != "" {
		dst.SubmitPath = src.SubmitPath
	}
	if src.SubmitPathTemplate != "" {
		dst.SubmitPathTemplate = src.SubmitPathTemplate
	}
	if src.QueryPath != "" {
		dst.QueryPath = src.QueryPath
	}
	if src.QueryPathTemplate != "" {
		dst.QueryPathTemplate = src.QueryPathTemplate
	}
	if src.TemplatePath != "" {
		dst.TemplatePath = src.TemplatePath
	}
	if src.FileRetrievePath != "" {
		dst.FileRetrievePath = src.FileRetrievePath
	}
	if src.AuthHeader != "" {
		dst.AuthHeader = src.AuthHeader
	}
	if src.AuthScheme != "" {
		dst.AuthScheme = src.AuthScheme
	}
	if src.DefaultCallbackURL != "" {
		dst.DefaultCallbackURL = src.DefaultCallbackURL
	}
	if src.EnablePromptExpansion != nil {
		dst.EnablePromptExpansion = src.EnablePromptExpansion
	}
	if len(src.ExtraHeaders) > 0 {
		if dst.ExtraHeaders == nil {
			dst.ExtraHeaders = map[string]string{}
		}
		for k, v := range src.ExtraHeaders {
			dst.ExtraHeaders[k] = v
		}
	}

	if dst.Upstream == MiniMaxVideoUpstreamPPInfra && dst.BaseURL == "https://api.minimaxi.com" {
		dst.BaseURL = "https://api.ppinfra.com"
	}
	if dst.Upstream == MiniMaxVideoUpstreamPPInfra && dst.SubmitPathTemplate == "" && dst.SubmitPath == "/v1/video_generation" {
		dst.SubmitPathTemplate = "/v3/async/%s"
		dst.SubmitPath = ""
	}
	if dst.Upstream == MiniMaxVideoUpstreamPPInfra {
		if dst.QueryPathTemplate == "" && dst.QueryPath == "" {
			dst.QueryPath = "/v3/async/task-result"
		} else if dst.QueryPathTemplate == "" && dst.QueryPath == "/v1/query/video_generation" {
			dst.QueryPath = "/v3/async/task-result"
		}
	}
    if dst.Upstream == MiniMaxVideoUpstreamPPInfra && dst.AuthScheme == "Bearer" {
        dst.AuthScheme = "Bearer"
    }

    // Polloi 上游：设置合理默认值（仅当仍为官方默认时生效）
    if dst.Upstream == MiniMaxVideoUpstreamPolloi {
        // 规范域名：Polloi 实际生产域为 https://pollo.ai/api/platform
        if dst.BaseURL == "https://api.minimaxi.com" || strings.TrimSpace(dst.BaseURL) == "" || strings.Contains(strings.ToLower(dst.BaseURL), "api.polloi.ai") {
            dst.BaseURL = "https://pollo.ai/api/platform"
        }
        // Polloi 提交：/generation/minimax/{model-segment}
        if dst.SubmitPathTemplate == "" && (dst.SubmitPath == "/v1/video_generation" || strings.TrimSpace(dst.SubmitPath) == "") {
            dst.SubmitPathTemplate = "/generation/minimax/%s"
            dst.SubmitPath = ""
        }
        // Polloi 查询：/generation/{taskId}/status （无需 model 段）
        if dst.QueryPathTemplate == "" && (dst.QueryPath == "/v1/query/video_generation" || strings.TrimSpace(dst.QueryPath) == "") {
            dst.QueryPathTemplate = "/generation/%s/status"
            dst.QueryPath = ""
        }
        // Polloi 使用 x-api-key 直传
        if strings.TrimSpace(dst.AuthHeader) == "" || strings.EqualFold(dst.AuthHeader, "authorization") {
            dst.AuthHeader = "x-api-key"
        }
        if strings.EqualFold(strings.TrimSpace(dst.AuthScheme), "") || strings.EqualFold(dst.AuthScheme, "bearer") {
            dst.AuthScheme = "none"
        }
    }
}

func (c *MiniMaxVideoClient) buildHeaders() map[string]string {
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	if authValue, ok := c.buildAuthorizationValue(); ok {
		headerName := c.config.AuthHeader
		if headerName == "" {
			headerName = "Authorization"
		}
		headers[headerName] = authValue
	}

	for k, v := range c.config.ExtraHeaders {
		headers[k] = v
	}

	return headers
}

func (c *MiniMaxVideoClient) buildAuthorizationValue() (string, bool) {
	apiKey := strings.TrimSpace(c.config.APIKey)
	if apiKey == "" {
		return "", false
	}

	scheme := strings.TrimSpace(c.config.AuthScheme)
	lowerKey := strings.ToLower(apiKey)
	if strings.HasPrefix(lowerKey, "bearer ") || strings.HasPrefix(lowerKey, "basic ") || strings.HasPrefix(lowerKey, "token ") {
		return apiKey, true
	}
	if strings.EqualFold(scheme, "none") {
		return apiKey, true
	}
	if scheme == "" {
		scheme = "Bearer"
	}
	return fmt.Sprintf("%s %s", scheme, apiKey), true
}

func formatMiniMaxModelSegment(model string) string {
	if model == "" {
		return ""
	}
	lower := strings.ToLower(model)
	lower = strings.ReplaceAll(lower, "_", "-")
	lower = strings.ReplaceAll(lower, " ", "-")
	return lower
}

func (c *MiniMaxVideoClient) buildSubmitPath(model string) string {
    if c.config.SubmitPathTemplate != "" && strings.Contains(c.config.SubmitPathTemplate, "%s") {
        segment := formatMiniMaxModelSegment(model)
        return fmt.Sprintf(c.config.SubmitPathTemplate, segment)
    }
    if c.config.SubmitPath != "" {
        return c.config.SubmitPath
    }
    return "/v1/video_generation"
}

// buildSubmitPathForReq 针对 PPInfra 补齐 2.3/2.3-fast 的 action 路径后缀（-t2v / -i2v）
func (c *MiniMaxVideoClient) buildSubmitPathForReq(req *MiniMaxVideoCreateRequest) string {
    // 官方上游或未使用模板：沿用默认逻辑
    if c.config.Upstream != MiniMaxVideoUpstreamPPInfra {
        return c.buildSubmitPath(req.Model)
    }
    // PPInfra: /v3/async/{segment}
    if c.config.SubmitPathTemplate == "" || !strings.Contains(c.config.SubmitPathTemplate, "%s") {
        return c.buildSubmitPath(req.Model)
    }
    base := formatMiniMaxModelSegment(req.Model)
    // 仅对 2.3/2.3-fast 追加动作后缀，其他保持原样
    if strings.Contains(base, "hailuo-2.3-fast") {
        // 2.3-fast 目前仅图生（i2v）
        base = base + "-i2v"
    } else if strings.Contains(base, "hailuo-2.3") {
        // 简易动作判断：存在首帧/参考图 → i2v，否则 t2v
        if strings.TrimSpace(req.FirstFrameImage) != "" || strings.TrimSpace(req.ReferenceImage) != "" {
            base = base + "-i2v"
        } else {
            base = base + "-t2v"
        }
    }
    return fmt.Sprintf(c.config.SubmitPathTemplate, base)
}

func (c *MiniMaxVideoClient) buildQueryPath(model, taskID string) string {
	if c.config.QueryPathTemplate != "" && strings.Count(c.config.QueryPathTemplate, "%s") > 0 {
		segment := formatMiniMaxModelSegment(model)
		if strings.Count(c.config.QueryPathTemplate, "%s") == 2 {
			return fmt.Sprintf(c.config.QueryPathTemplate, segment, url.PathEscape(taskID))
		}
		return fmt.Sprintf(c.config.QueryPathTemplate, url.PathEscape(taskID))
	}
	if c.config.QueryPath != "" {
		return c.config.QueryPath
	}
	return "/v1/query/video_generation"
}

func (c *MiniMaxVideoClient) buildFileRetrievePath() string {
	if c.config.FileRetrievePath != "" {
		return c.config.FileRetrievePath
	}
	return "/v1/files/retrieve"
}

// SubmitVideoTask 创建视频生成任务
func (c *MiniMaxVideoClient) SubmitVideoTask(req *MiniMaxVideoCreateRequest) (*MiniMaxVideoCreateResponse, *types.OpenAIError) {
	// 前置校验：必须配置有效的上游鉴权密钥（来自 channel.Key 或 video.api_key）
	if v, ok := c.buildAuthorizationValue(); !ok || strings.TrimSpace(v) == "" {
		return nil, &types.OpenAIError{
			Message: "MiniMax video channel key is empty. Please set channel key or custom_parameter.video.api_key",
			Type:    "channel_config_missing",
			Code:    "invalid_config",
		}
	}
    payload := c.prepareSubmitPayload(req)
    submitPath := c.buildSubmitPathForReq(req)
    fullURL := c.GetFullRequestURL(submitPath, "")

	headers := c.buildHeaders()

	httpReq, err := c.Requester.NewRequest(http.MethodPost, fullURL, c.Requester.WithBody(payload), c.Requester.WithHeader(headers))
	if err != nil {
		return nil, &types.OpenAIError{Message: fmt.Sprintf("create request failed: %s", err.Error()), Type: "minimax_video_request_error"}
	}

    // Polloi 上游：响应为 camelCase（taskId 等），需单独处理
    if c.config.Upstream == MiniMaxVideoUpstreamPolloi {
        var raw string
        if _, errWithCode := c.Requester.SendRequest(httpReq, &raw, false); errWithCode != nil {
            return nil, &errWithCode.OpenAIError
        }
        type polloiCreateResp struct {
            TaskID string `json:"taskId"`
            Status string `json:"status"`
        }
        var pr polloiCreateResp
        if err := json.Unmarshal([]byte(raw), &pr); err != nil {
            return nil, &types.OpenAIError{Message: fmt.Sprintf("decode polloi create response failed: %v", err), Type: "minimax_video_decode_error"}
        }
        return &MiniMaxVideoCreateResponse{TaskID: pr.TaskID, BaseResp: BaseResp{StatusCode: 0, StatusMsg: "success"}}, nil
    }

    var resp MiniMaxVideoCreateResponse
    _, errWithCode := c.Requester.SendRequest(httpReq, &resp, false)
    if errWithCode != nil {
        return nil, &errWithCode.OpenAIError
    }

    return &resp, nil
}

// QueryVideoTask 查询视频任务状态
func (c *MiniMaxVideoClient) QueryVideoTask(taskID, model string) (*MiniMaxVideoQueryResponse, *types.OpenAIError) {
	if v, ok := c.buildAuthorizationValue(); !ok || strings.TrimSpace(v) == "" {
		return nil, &types.OpenAIError{
			Message: "MiniMax video channel key is empty. Please set channel key or custom_parameter.video.api_key",
			Type:    "channel_config_missing",
			Code:    "invalid_config",
		}
	}
	queryPath := c.buildQueryPath(model, taskID)
	var fullURL string
	if c.config.QueryPathTemplate != "" && strings.Count(c.config.QueryPathTemplate, "%s") > 0 {
		fullURL = c.GetFullRequestURL(queryPath, "")
	} else {
		values := url.Values{}
		values.Set("task_id", taskID)
		fullURL = c.GetFullRequestURL(fmt.Sprintf("%s?%s", queryPath, values.Encode()), "")
	}

	headers := c.buildHeaders()

	httpReq, err := c.Requester.NewRequest(http.MethodGet, fullURL, c.Requester.WithHeader(headers))
	if err != nil {
		return nil, &types.OpenAIError{Message: fmt.Sprintf("create request failed: %s", err.Error()), Type: "minimax_video_request_error"}
	}

    // Polloi 上游：响应结构为 {taskId, generations: [{id,status,failMsg,url,mediaType,...}]}
    if c.config.Upstream == MiniMaxVideoUpstreamPolloi {
        var raw string
        if _, errWithCode := c.Requester.SendRequest(httpReq, &raw, false); errWithCode != nil {
            return nil, &errWithCode.OpenAIError
        }
        type polloiGen struct {
            ID        string `json:"id"`
            Status    string `json:"status"`
            FailMsg   string `json:"failMsg"`
            URL       string `json:"url"`
            MediaType string `json:"mediaType"`
            CreatedAt string `json:"createdDate"`
            UpdatedAt string `json:"updatedDate"`
        }
        type polloiQueryResp struct {
            TaskID      string       `json:"taskId"`
            Generations []polloiGen  `json:"generations"`
        }
        var pq polloiQueryResp
        if err := json.Unmarshal([]byte(raw), &pq); err != nil {
            return nil, &types.OpenAIError{Message: fmt.Sprintf("decode polloi query response failed: %v", err), Type: "minimax_video_decode_error"}
        }

        out := &MiniMaxVideoQueryResponse{TaskID: pq.TaskID, BaseResp: BaseResp{StatusCode: 0, StatusMsg: "success"}}
        // 选择首个 video 或首个条目
        var chosen *polloiGen
        for i := range pq.Generations {
            if strings.EqualFold(pq.Generations[i].MediaType, "video") {
                chosen = &pq.Generations[i]
                break
            }
        }
        if chosen == nil && len(pq.Generations) > 0 {
            chosen = &pq.Generations[0]
        }
        if chosen != nil {
            st := strings.ToLower(strings.TrimSpace(chosen.Status))
            switch st {
            case "waiting":
                out.Status = "queueing"
            case "processing":
                out.Status = "processing"
            case "succeed", "success", "completed":
                out.Status = "success"
            case "failed", "fail", "error":
                out.Status = "failed"
            default:
                out.Status = st
            }
            out.VideoURL = strings.TrimSpace(chosen.URL)
            if out.Status == "failed" && strings.TrimSpace(chosen.FailMsg) != "" {
                out.ErrorMessage = strings.TrimSpace(chosen.FailMsg)
            }
        }

        out.Normalize()
        return out, nil
    }

    var resp MiniMaxVideoQueryResponse
    _, errWithCode := c.Requester.SendRequest(httpReq, &resp, false)
    if errWithCode != nil {
        return nil, &errWithCode.OpenAIError
    }

    resp.Normalize()

    return &resp, nil
}

// RetrieveFile 获取文件下载信息
func (c *MiniMaxVideoClient) RetrieveFile(fileID string) (*MiniMaxFileRetrieveResponse, *types.OpenAIError) {
	if v, ok := c.buildAuthorizationValue(); !ok || strings.TrimSpace(v) == "" {
		return nil, &types.OpenAIError{
			Message: "MiniMax video channel key is empty. Please set channel key or custom_parameter.video.api_key",
			Type:    "channel_config_missing",
			Code:    "invalid_config",
		}
	}
	filePath := c.buildFileRetrievePath()
	values := url.Values{}
	values.Set("file_id", fileID)
	fullURL := c.GetFullRequestURL(fmt.Sprintf("%s?%s", filePath, values.Encode()), "")

	headers := c.buildHeaders()

	httpReq, err := c.Requester.NewRequest(http.MethodGet, fullURL, c.Requester.WithHeader(headers))
	if err != nil {
		return nil, &types.OpenAIError{Message: fmt.Sprintf("create request failed: %s", err.Error()), Type: "minimax_file_request_error"}
	}

	var resp MiniMaxFileRetrieveResponse
	_, errWithCode := c.Requester.SendRequest(httpReq, &resp, false)
	if errWithCode != nil {
		return nil, &errWithCode.OpenAIError
	}

	return &resp, nil
}

func (c *MiniMaxVideoClient) prepareSubmitPayload(req *MiniMaxVideoCreateRequest) interface{} {
	clone := *req

	if clone.CallbackURL == "" {
		clone.CallbackURL = c.config.DefaultCallbackURL
	}
	if clone.EnablePromptExpansion == nil && c.config.EnablePromptExpansion != nil {
		value := *c.config.EnablePromptExpansion
		clone.EnablePromptExpansion = &value
	}

    // PPInfra 上游：使用其兼容字段与命名
    if c.config.Upstream == MiniMaxVideoUpstreamPPInfra {
        payload := map[string]interface{}{}
		if clone.Prompt != "" {
			payload["prompt"] = clone.Prompt
		}
		if clone.FirstFrameImage != "" {
			payload["image"] = clone.FirstFrameImage
		}
		if clone.LastFrameImage != "" {
			payload["end_image"] = clone.LastFrameImage
		}
		if clone.ReferenceImage != "" {
			payload["reference_image"] = clone.ReferenceImage
		}
		if len(clone.SubjectReference) > 0 {
			payload["subject_reference"] = clone.SubjectReference
		}
		if clone.Duration > 0 {
			payload["duration"] = clone.Duration
		}
		if clone.Resolution != "" {
			payload["resolution"] = clone.Resolution
		}
		if clone.EnablePromptExpansion != nil {
			payload["enable_prompt_expansion"] = *clone.EnablePromptExpansion
		}
		if clone.CallbackURL != "" {
			payload["callback_url"] = clone.CallbackURL
		}
		if clone.ExternalTaskID != "" {
			payload["external_task_id"] = clone.ExternalTaskID
		}
		return payload
	}

    // Polloi 上游：外层包裹 input + webhookUrl，input 内对齐我们内部统一字段
    if c.config.Upstream == MiniMaxVideoUpstreamPolloi {
        input := map[string]interface{}{}
        if clone.Prompt != "" {
            input["prompt"] = clone.Prompt
        }
        if clone.FirstFrameImage != "" {
            input["image"] = clone.FirstFrameImage
        }
        if clone.LastFrameImage != "" {
            input["end_image"] = clone.LastFrameImage
        }
        if clone.ReferenceImage != "" {
            input["reference_image"] = clone.ReferenceImage
        }
        if len(clone.SubjectReference) > 0 {
            input["subject_reference"] = clone.SubjectReference
        }
        if clone.Duration > 0 {
            input["duration"] = clone.Duration
        }
        if clone.Resolution != "" {
            input["resolution"] = clone.Resolution
        }
        if clone.EnablePromptExpansion != nil {
            input["enable_prompt_expansion"] = *clone.EnablePromptExpansion
        } else if clone.PromptOptimizer != nil {
            input["enable_prompt_expansion"] = *clone.PromptOptimizer
        }
        if clone.ExternalTaskID != "" {
            input["external_task_id"] = clone.ExternalTaskID
        }

        payload := map[string]interface{}{"input": input}
        if clone.CallbackURL != "" {
            payload["webhookUrl"] = clone.CallbackURL
        }
        return payload
    }

    // 官方上游：严格对齐 minimaxi 官方字段
    payload := map[string]interface{}{}
	if clone.Model != "" {
		payload["model"] = clone.Model
	}
	if clone.Prompt != "" {
		payload["prompt"] = clone.Prompt
	}
	if clone.FirstFrameImage != "" {
		payload["first_frame_image"] = clone.FirstFrameImage
	}
	if clone.LastFrameImage != "" {
		payload["last_frame_image"] = clone.LastFrameImage
	}
	if clone.ReferenceImage != "" {
		payload["reference_image"] = clone.ReferenceImage
	}
	if len(clone.SubjectReference) > 0 {
		payload["subject_reference"] = clone.SubjectReference
	}
	if clone.Duration > 0 {
		payload["duration"] = clone.Duration
	}
	if clone.Resolution != "" {
		payload["resolution"] = clone.Resolution
	}
	if clone.CallbackURL != "" {
		payload["callback_url"] = clone.CallbackURL
	}
	if clone.ExternalTaskID != "" {
		payload["external_task_id"] = clone.ExternalTaskID
	}

	// prompt_optimizer 优先：若未显式传入，则复用 enable_prompt_expansion
	if clone.PromptOptimizer != nil {
		payload["prompt_optimizer"] = *clone.PromptOptimizer
	} else if clone.EnablePromptExpansion != nil {
		payload["prompt_optimizer"] = *clone.EnablePromptExpansion
	}
	if clone.FastPretreatment != nil {
		payload["fast_pretreatment"] = *clone.FastPretreatment
	}
	if clone.AIGCWatermark != nil {
		payload["aigc_watermark"] = *clone.AIGCWatermark
	}

	return payload
}

// miniMaxVideoErrorHandle 解析视频接口错误
func miniMaxVideoErrorHandle(resp *http.Response) *types.OpenAIError {
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return &types.OpenAIError{
			Message: fmt.Sprintf("minimax video upstream error: %s", resp.Status),
			Type:    "minimax_video_error",
			Code:    fmt.Sprintf("%d", resp.StatusCode),
		}
	}

	var baseWrap struct {
		BaseResp BaseResp `json:"base_resp"`
		Message  string   `json:"message"`
		Code     any      `json:"code"`
		Error    any      `json:"error"`
	}
	if err := json.Unmarshal(bodyBytes, &baseWrap); err == nil {
		if baseWrap.BaseResp.StatusCode != 0 {
			return &types.OpenAIError{
				Message: baseWrap.BaseResp.StatusMsg,
				Type:    "minimax_video_error",
				Code:    fmt.Sprintf("%d", baseWrap.BaseResp.StatusCode),
			}
		}
		if baseWrap.Message != "" {
			return &types.OpenAIError{
				Message: baseWrap.Message,
				Type:    "minimax_video_error",
			}
		}
	}

	return &types.OpenAIError{
		Message: string(bodyBytes),
		Type:    "minimax_video_error",
		Code:    fmt.Sprintf("%d", resp.StatusCode),
	}
}

// MiniMaxVideoCreateRequest 统一描述提交任务所需参数
type MiniMaxVideoCreateRequest struct {
    // 模型名称，例如：MiniMax-Hailuo-2.3
    Model            string                    `json:"model,omitempty" example:"MiniMax-Hailuo-2.3"`
    // 提示词
    Prompt           string                    `json:"prompt,omitempty" example:"A man picks up a book [Pedestal up], then reads [Static shot]."`
    // 首帧参考图 URL
    FirstFrameImage  string                    `json:"first_frame_image,omitempty"`
    // 末帧参考图 URL
    LastFrameImage   string                    `json:"last_frame_image,omitempty"`
    // 参考图 URL（单张）
    ReferenceImage   string                    `json:"reference_image,omitempty"`
    // 主体参考
    SubjectReference []MiniMaxSubjectReference `json:"subject_reference,omitempty"`
    // 时长（秒）
    Duration         int                       `json:"duration,omitempty"`
    // 分辨率
    Resolution       string                    `json:"resolution,omitempty" enums:"720p,1080p,1080P,2k,4k" example:"1080P"`
    // 回调地址
    CallbackURL      string                    `json:"callback_url,omitempty"`
	// 兼容历史/上游差异：
	// - 官方文档字段：prompt_optimizer（是否自动优化提示词）
	// - 历史/PPInfra：enable_prompt_expansion（提示词扩写开关）
    EnablePromptExpansion *bool `json:"enable_prompt_expansion,omitempty"`
    PromptOptimizer       *bool `json:"prompt_optimizer,omitempty"`
	// 仅 MiniMax-Hailuo-02 生效：缩短 prompt_optimizer 优化耗时
    FastPretreatment *bool `json:"fast_pretreatment,omitempty"`
	// 是否在生成视频中添加 AIGC 水印
    AIGCWatermark  *bool  `json:"aigc_watermark,omitempty"`
    ExternalTaskID string `json:"external_task_id,omitempty"`
}

// MiniMaxVideoCreateResponse 表示创建任务返回的数据
type MiniMaxVideoCreateResponse struct {
    // 任务 ID
    TaskID   string   `json:"task_id" example:"106916112212032"`
    // 基本返回体
    BaseResp BaseResp `json:"base_resp"`
}

// MiniMaxVideoQueryResponse 表示查询任务返回的数据
type MiniMaxVideoQueryResponse struct {
    // 任务 ID
    TaskID          string   `json:"task_id,omitempty"`
    // 任务状态
    Status          string   `json:"status,omitempty" enums:"queueing,processing,success,failed"`
    // 文件 ID
    FileID          string   `json:"file_id,omitempty"`
    // 视频直链
    VideoURL        string   `json:"video_url,omitempty"`
    // 封面图
    CoverImage      string   `json:"cover_image,omitempty"`
    // 水印视频直链
    WatermarkedURL  string   `json:"watermarked_url,omitempty"`
    // 宽度
    VideoWidth      int      `json:"video_width,omitempty"`
    // 高度
    VideoHeight     int      `json:"video_height,omitempty"`
    // 错误码
    ErrorCode       string   `json:"error_code,omitempty"`
    // 错误信息
    ErrorMessage    string   `json:"error_message,omitempty"`
    // 进度百分比
    ProgressPercent float64  `json:"progress_percent,omitempty"`
    // 预计剩余时间（秒）
    ETA             int      `json:"eta,omitempty"`
    // 基本返回体
    BaseResp        BaseResp `json:"base_resp"`

	Extra      *MiniMaxPPInfraExtra  `json:"extra,omitempty"`
	TaskDetail *MiniMaxPPInfraTask   `json:"task,omitempty"`
	Videos     []MiniMaxPPInfraVideo `json:"videos,omitempty"`
	Images     []MiniMaxPPInfraImage `json:"images,omitempty"`
	Audios     []MiniMaxPPInfraAudio `json:"audios,omitempty"`
}

type MiniMaxFileRetrieveResponse struct {
	File     MiniMaxFileObject `json:"file"`
	BaseResp BaseResp          `json:"base_resp"`
}

type MiniMaxFileObject struct {
	FileID      types.StringOrNumber `json:"file_id,omitempty"`
	Bytes       int64                `json:"bytes,omitempty"`
	CreatedAt   int64                `json:"created_at,omitempty"`
	Filename    string               `json:"filename,omitempty"`
	Purpose     string               `json:"purpose,omitempty"`
	DownloadURL string               `json:"download_url,omitempty"`
}

type MiniMaxPPInfraExtra struct {
	Seed      string                   `json:"seed,omitempty"`
	DebugInfo *MiniMaxPPInfraDebugInfo `json:"debug_info,omitempty"`
}

type MiniMaxPPInfraDebugInfo struct {
	RequestInfo    string `json:"request_info,omitempty"`
	SubmitTimeMS   string `json:"submit_time_ms,omitempty"`
	ExecuteTimeMS  string `json:"execute_time_ms,omitempty"`
	CompleteTimeMS string `json:"complete_time_ms,omitempty"`
}

type MiniMaxPPInfraTask struct {
	TaskID          string  `json:"task_id,omitempty"`
	Status          string  `json:"status,omitempty"`
	Reason          string  `json:"reason,omitempty"`
	TaskType        string  `json:"task_type,omitempty"`
	ETA             int     `json:"eta,omitempty"`
	ProgressPercent float64 `json:"progress_percent,omitempty"`
}

type MiniMaxPPInfraVideo struct {
	VideoURL    string `json:"video_url,omitempty"`
	VideoURLTTL string `json:"video_url_ttl,omitempty"`
	VideoType   string `json:"video_type,omitempty"`
}

type MiniMaxPPInfraImage struct {
	ImageURL    string `json:"image_url,omitempty"`
	ImageURLTTL int    `json:"image_url_ttl,omitempty"`
	ImageType   string `json:"image_type,omitempty"`
}

type MiniMaxPPInfraAudio struct {
	AudioURL      string                   `json:"audio_url,omitempty"`
	AudioURLTTL   string                   `json:"audio_url_ttl,omitempty"`
	AudioType     string                   `json:"audio_type,omitempty"`
	AudioMetadata *MiniMaxPPInfraAudioMeta `json:"audio_metadata,omitempty"`
}

type MiniMaxPPInfraAudioMeta struct {
	Text      string `json:"text,omitempty"`
	StartTime int    `json:"start_time,omitempty"`
	EndTime   int    `json:"end_time,omitempty"`
}

func (r *MiniMaxVideoQueryResponse) Normalize() {
	if r == nil {
		return
	}
	if r.TaskDetail != nil {
		if r.TaskID == "" {
			r.TaskID = r.TaskDetail.TaskID
		}
		if r.Status == "" {
			r.Status = r.TaskDetail.Status
		}
		if r.ProgressPercent == 0 && r.TaskDetail.ProgressPercent > 0 {
			r.ProgressPercent = r.TaskDetail.ProgressPercent
		}
		if r.ETA == 0 && r.TaskDetail.ETA > 0 {
			r.ETA = r.TaskDetail.ETA
		}
		if r.ErrorMessage == "" {
			r.ErrorMessage = r.TaskDetail.Reason
		}
	}

	if r.CoverImage == "" && len(r.Images) > 0 {
		for _, img := range r.Images {
			if img.ImageURL != "" {
				r.CoverImage = img.ImageURL
				break
			}
		}
	}

	if len(r.Videos) > 0 {
		for _, video := range r.Videos {
			typeLower := strings.ToLower(video.VideoType)
			if r.VideoURL == "" && (typeLower == "" || strings.Contains(typeLower, "origin") || strings.Contains(typeLower, "mp4")) {
				r.VideoURL = video.VideoURL
			}
			if r.WatermarkedURL == "" && (strings.Contains(typeLower, "watermark") || strings.Contains(typeLower, "wm")) {
				r.WatermarkedURL = video.VideoURL
			}
		}
		if r.VideoURL == "" {
			r.VideoURL = r.Videos[0].VideoURL
		}
	}

	// 统一状态字段：全部转为小写，并合并同义词
	statusLower := strings.ToLower(strings.TrimSpace(r.Status))
	switch statusLower {
	case "task_status_succeed", "success", "succeeded", "completed", "finish", "finished", "done":
		statusLower = "success"
	case "task_status_failed", "fail", "failed", "error", "exception":
		statusLower = "failed"
	case "task_status_processing", "processing", "pending", "waiting", "running", "in_progress":
		statusLower = "processing"
	case "task_status_queued", "queueing", "queued":
		statusLower = "queueing"
	}
	r.Status = statusLower

	// 对齐官方响应体：移除上游（PPInfra）特有的字段，避免干扰调用方
	r.Extra = nil
	r.TaskDetail = nil
	r.Videos = nil
	r.Images = nil
	r.Audios = nil
}
