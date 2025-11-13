package gemini

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "mime"
    "mime/multipart"
    "net/http"
    "one-api/common"
    "one-api/common/logger"
    img "one-api/common/image"
    "one-api/common/requester"
    "one-api/model"
    "one-api/providers/base"
    "one-api/providers/openai"
    "one-api/types"
    "strconv"
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
        // 补齐 Videos 路径，确保当 GeminiProvider 参与 OpenAI 风格的视频创建时，
        // OpenAIProvider.CreateVideo 能拿到 /v1/videos（例如上游为 Sutui/OpenAI 兼容）。
        Videos:            "/v1/videos",
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

    full := fmt.Sprintf("%s/%s/models/%s:%s", baseURL, version, modelName, requestURL)
    if strings.Contains(strings.ToLower(requestURL), "/v1/videos") || strings.Contains(strings.ToLower(full), "/v1/videos") {
        logger.SysDebug(fmt.Sprintf("gemini.getfullurl.warn -> requestURL=%s model=%s full=%s", requestURL, modelName, full))
    }
    return full

}

// 获取请求头
func (p *GeminiProvider) GetRequestHeaders() (headers map[string]string) {
    headers = make(map[string]string)
    p.CommonRequestHeaders(headers)
    headers["x-goog-api-key"] = p.Channel.Key

    return headers
}

// detectVeoVendor: 返回 Gemini Veo 上游供应商标识，google(默认)/sutui
func (p *GeminiProvider) detectVeoVendor() string {
    if p.Channel != nil && p.Channel.Plugin != nil {
        if plugin := p.Channel.Plugin.Data(); plugin != nil {
            if gm, ok := plugin["gemini_video"]; ok {
                // plugin 的类型为 PluginType(map[string]map[string]interface{})
                if v, ok3 := gm["vendor"].(string); ok3 {
                    if strings.EqualFold(v, "sutui") { return "sutui" }
                    if strings.EqualFold(v, "apimart") { return "apimart" }
                    if strings.EqualFold(v, "ezlinkai") { return "ezlinkai" }
                }
            }
            // 兼容形式：plugin.gemini.video.vendor
            if gm, ok := plugin["gemini"]; ok {
                if vm, ok3 := gm["video"].(map[string]any); ok3 {
                    if v, ok4 := vm["vendor"].(string); ok4 {
                        if strings.EqualFold(v, "sutui") { return "sutui" }
                        if strings.EqualFold(v, "apimart") { return "apimart" }
                        if strings.EqualFold(v, "ezlinkai") { return "ezlinkai" }
                    }
                }
            }
        }
    }
    base := strings.ToLower(strings.TrimSpace(p.GetBaseURL()))
    // 识别 Sutui 系列域名：sutui / st-ai / sora2.pub
    if base != "" && (strings.Contains(base, "sutui") || strings.Contains(base, "st-ai") || strings.Contains(base, "sora2.pub")) {
        return "sutui"
    }
    if base != "" && strings.Contains(base, "apimart.ai") {
        return "apimart"
    }
    if base != "" && strings.Contains(base, "ezlinkai.com") {
        return "ezlinkai"
    }
    return "google"
}

// deduceSize 将官方的 aspectRatio 与 resolution 推导为 sutui 上游所需的像素尺寸
// 约定：
// - 720p + 16:9 => 1280x720
// - 720p + 9:16 => 720x1280
// - 1080p + 16:9 => 1920x1080
// - 1080p + 9:16 => 1080x1920
// - 缺省回退：
//   - 仅 AR 给出：16:9 => 1600x900，9:16 => 900x1600，1:1 => 720x720
func deduceSize(aspectRatio, resolution string) string {
    ar := strings.TrimSpace(aspectRatio)
    res := strings.TrimSpace(strings.ToLower(resolution))
    switch res {
    case "720p":
        switch ar {
        case "16:9":
            return "1280x720"
        case "9:16":
            return "720x1280"
        }
    case "1080p":
        switch ar {
        case "16:9":
            return "1920x1080"
        case "9:16":
            return "1080x1920"
        }
    }
    // 回退：仅根据 AR 推断
    switch ar {
    case "16:9":
        return "1600x900"
    case "9:16":
        return "900x1600"
    case "1:1":
        return "720x720"
    }
    if strings.Contains(ar, "x") { return ar }
    return ""
}

// 将 Veo :predictLongRunning 初始化映射到 sutui 的 /v1/videos
func (p *GeminiProvider) relayPredictLongRunningViaSutui(modelName string) (any, *types.OpenAIErrorWithStatusCode) {
    base := strings.TrimSuffix(p.GetBaseURL(), "/")
    fullURL := base + "/v1/videos"

    // Sutui 使用 OpenAI 风格鉴权（Bearer），不要使用 x-goog-api-key
    headers := map[string]string{}
    p.CommonRequestHeaders(headers)
    headers["Authorization"] = fmt.Sprintf("Bearer %s", p.Channel.Key)
    contentType := ""
    if p.Context != nil && p.Context.Request != nil {
        contentType = p.Context.Request.Header.Get("Content-Type")
    }
    if contentType != "" { headers["Content-Type"] = contentType }

    logger.SysDebug(fmt.Sprintf("gemini.veo.sutui.init -> base=%s model=%s", p.GetBaseURL(), modelName))
    var httpReq *http.Request
    // multipart：仅改写 model 字段为 sutui 值，其它（含 seconds）保持
    if strings.Contains(strings.ToLower(contentType), "multipart/form-data") {
        raw, ok := p.GetRawBody()
        if !ok {
            return nil, common.StringErrorWrapperLocal("request body not found", "request_body_not_found", http.StatusInternalServerError)
        }
        bodyReader, newCT, err := p.rewriteMultipartModelForSutui(raw, contentType, modelName)
        if err != nil || bodyReader == nil || strings.TrimSpace(newCT) == "" {
            // 解析失败则回退直透，避免阻断
            r, e := p.Requester.NewRequest(http.MethodPost, fullURL, p.Requester.WithBody(raw), p.Requester.WithHeader(headers))
            if e != nil { return nil, common.ErrorWrapper(e, "new_request_failed", http.StatusInternalServerError) }
            httpReq = r
        } else {
            headers["Content-Type"] = newCT
            r, e := p.Requester.NewRequest(http.MethodPost, fullURL, p.Requester.WithBody(bodyReader), p.Requester.WithHeader(headers))
            if e != nil { return nil, common.ErrorWrapper(e, "new_request_failed", http.StatusInternalServerError) }
            httpReq = r
        }
    } else {
        // 将官方 JSON 形态解析后，改写为 Sutui 需要的 multipart/form-data：
        // 1) 官方形态：
        //    {
        //      "instances": [{ "prompt": "..." }],
        //      "parameters": { "durationSeconds": "8", "aspectRatio": "16:9", "resolution": "720p" }
        //    }
        // 2) 兼容旧形态：
        //    { "input": {"prompt": "..."}, "config": {"durationSeconds": 8, "aspectRatio": "16:9"} }
        raw, ok := p.GetRawBody()
        if !ok {
            return nil, common.StringErrorWrapperLocal("request body not found", "request_body_not_found", http.StatusInternalServerError)
        }
        var m map[string]any
        _ = json.Unmarshal(raw, &m)
        var prompt string
        var seconds int
        var size string

        // 优先解析官方形态
        if inst, ok := m["instances"].([]any); ok && len(inst) > 0 {
            if first, ok2 := inst[0].(map[string]any); ok2 {
                if s, ok3 := first["prompt"].(string); ok3 { prompt = s }
            }
        }
        if params, ok := m["parameters"].(map[string]any); ok {
            // durationSeconds 可能为字符串或数字
            if v, ok2 := params["durationSeconds"]; ok2 {
                switch t := v.(type) {
                case string:
                    if n, err := strconv.Atoi(strings.TrimSpace(t)); err == nil { seconds = n }
                case float64:
                    seconds = int(t)
                case int:
                    seconds = t
                }
            }
            var ar, res string
            if s, ok2 := params["aspectRatio"].(string); ok2 { ar = strings.TrimSpace(s) }
            if s, ok2 := params["resolution"].(string); ok2 { res = strings.TrimSpace(s) }
            if size == "" {
                size = deduceSize(ar, res)
            }
        }

        // 兼容旧形态 input/config
        if prompt == "" {
            if in, ok := m["input"].(map[string]any); ok {
                if s, ok2 := in["prompt"].(string); ok2 { prompt = s }
            }
        }
        if seconds == 0 || size == "" {
            if cfg, ok := m["config"].(map[string]any); ok {
                if seconds == 0 {
                    if f, ok2 := cfg["durationSeconds"].(float64); ok2 { seconds = int(f) }
                }
                if s, ok2 := cfg["aspectRatio"].(string); ok2 {
                    s = strings.TrimSpace(s)
                    if size == "" { size = deduceSize(s, "") }
                }
            }
        }

    // 官方 Veo -> Sutui veo3.* 模型映射
    // - veo-3.1-fast-generate-preview    -> veo3.1
    // - veo-3.1-generate-preview         -> veo3.1
    sutuiModel := modelName
    low := strings.ToLower(strings.TrimSpace(modelName))
    switch low {
    case "veo-3.1-fast-generate-preview":
        sutuiModel = "veo3.1"      // 快速系
    case "veo-3.1-generate-preview":
        sutuiModel = "veo3.1-pro"  // 高质量
    default:
        if strings.Contains(low, "fast-generate-preview") {
            sutuiModel = "veo3.1"
        } else if strings.HasPrefix(low, "veo-") {
            sutuiModel = "veo3.1-pro"
        }
    }

        // 组装 multipart 表单（文本字段）
        var form bytes.Buffer
        builder := p.Requester.CreateFormBuilder(&form)
        if sutuiModel != "" { _ = builder.WriteField("model", sutuiModel) }
        if prompt != "" { _ = builder.WriteField("prompt", prompt) }
        if seconds > 0 { _ = builder.WriteField("seconds", strconv.Itoa(seconds)) }
        if size != "" { _ = builder.WriteField("size", size) }
        _ = builder.Close()
        headers["Content-Type"] = builder.FormDataContentType()
        r, err := p.Requester.NewRequest(http.MethodPost, fullURL, p.Requester.WithBody(&form), p.Requester.WithHeader(headers))
        if err != nil { return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError) }
        httpReq = r
    }
    if httpReq.Body != nil { defer httpReq.Body.Close() }
    func() {
        keys := []string{}
        for k := range headers { keys = append(keys, k) }
        logger.SysDebug(fmt.Sprintf("gemini.veo.sutui.init -> url=%s headers=%v", fullURL, keys))
    }()

    var resp struct{ ID string `json:"id"` }
    if _, e := p.Requester.SendRequest(httpReq, &resp, false); e != nil {
        return nil, e
    }
    if strings.TrimSpace(resp.ID) == "" {
        return nil, common.StringErrorWrapperLocal("missing id from sutui response", "upstream_error", http.StatusBadGateway)
    }
    return map[string]any{ "name": fmt.Sprintf("operations/%s", resp.ID), "done": false }, nil
}

// rewriteMultipartModelForSutui: 在 sutui 场景只改写 multipart 表单的 model 值，不动 seconds 与其它字段/文件。
func (p *GeminiProvider) rewriteMultipartModelForSutui(raw []byte, contentType string, modelName string) (io.Reader, string, error) {
    mediaType, params, err := mime.ParseMediaType(contentType)
    if err != nil || !strings.HasPrefix(strings.ToLower(mediaType), "multipart/") {
        return nil, "", fmt.Errorf("invalid multipart content-type")
    }
    boundary := strings.TrimSpace(params["boundary"])
    if boundary == "" { return nil, "", fmt.Errorf("missing boundary") }

    // 模型映射函数
    mapModel := func(in string) string {
        m := strings.ToLower(strings.TrimSpace(in))
        if strings.HasPrefix(m, "veo3") { return in }
        if strings.Contains(m, "fast-generate-preview") { return "veo3.1" }
        if strings.HasPrefix(m, "veo-") { return "veo3.1-pro" }
        // 兜底：参考路径模型
        low := strings.ToLower(strings.TrimSpace(modelName))
        if strings.Contains(low, "fast-generate-preview") { return "veo3.1" }
        if strings.HasPrefix(low, "veo-") { return "veo3.1-pro" }
        return in
    }

    // 读取原表单
    mr := multipart.NewReader(bytes.NewReader(raw), boundary)
    type filePart struct{ field, filename string; data []byte }
    fields := map[string][]string{}
    files := []filePart{}
    for {
        part, e := mr.NextPart()
        if e == io.EOF { break }
        if e != nil { return nil, "", e }
        name := part.FormName()
        if name == "" { _ = part.Close(); continue }
        if fn := part.FileName(); fn != "" {
            b, e2 := io.ReadAll(part); _ = part.Close()
            if e2 != nil { return nil, "", e2 }
            files = append(files, filePart{field: name, filename: fn, data: b})
            continue
        }
        b, e3 := io.ReadAll(part); _ = part.Close()
        if e3 != nil { return nil, "", e3 }
        fields[name] = append(fields[name], string(b))
    }

    // 计算目标模型
    target := ""
    if arr, ok := fields["model"]; ok && len(arr) > 0 {
        target = mapModel(arr[0])
    } else {
        target = mapModel("")
    }

    // 重建表单
    var buf bytes.Buffer
    builder := p.Requester.CreateFormBuilder(&buf)
    if strings.TrimSpace(target) != "" {
        if err := builder.WriteField("model", target); err != nil { return nil, "", err }
    }
    for k, vals := range fields {
        if strings.EqualFold(k, "model") { continue }
        for _, v := range vals {
            if err := builder.WriteField(k, v); err != nil { return nil, "", err }
        }
    }
    for _, f := range files {
        if err := builder.CreateFormFileReader(f.field, bytes.NewReader(f.data), f.filename); err != nil { return nil, "", err }
    }
    _ = builder.Close()
    return &buf, builder.FormDataContentType(), nil
}

// 将 Veo 初始化映射到 Apimart 的 /v1/videos/generations
// 要求：
// - Authorization: Bearer <key>
// - JSON: {model,prompt,duration=8,aspect_ratio,image_urls}
func (p *GeminiProvider) relayPredictLongRunningViaApimart(modelName string, prompt string, seconds int, size string, refs []string) (string, *types.OpenAIErrorWithStatusCode) {
    base := strings.TrimSuffix(p.GetBaseURL(), "/")
    fullURL := base + "/v1/videos/generations"

    // 模型映射：官方 -> apimart
    m := strings.ToLower(strings.TrimSpace(modelName))
    apimartModel := m
    switch m {
    case "veo-3.1-generate-preview":
        apimartModel = "veo3.1" // 质量档
    case "veo-3.1-fast-generate-preview":
        apimartModel = "veo3.1-fast"
    }

    // AR 映射
    ar, _ := parseSizeToGemini(size)
    // Apimart 固定 duration=8
    duration := 8
    // image_urls 从 refs 直接透传

    body := map[string]any{
        "model":        apimartModel,
        "prompt":       prompt,
        "duration":     duration,
        "aspect_ratio": ar,
    }
    if len(refs) > 0 {
        body["image_urls"] = refs
    }
    // Apimart 暂不支持 1080p 指定，但通过 AR 控制画幅；忽略 res

    headers := map[string]string{
        "Authorization": fmt.Sprintf("Bearer %s", p.Channel.Key),
        "Content-Type":  "application/json",
    }
    req, err := p.Requester.NewRequest(http.MethodPost, fullURL, p.Requester.WithBody(body), p.Requester.WithHeader(headers))
    if err != nil { return "", common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError) }
    defer req.Body.Close()

    var resp struct{
        Code int `json:"code"`
        Data []struct{ Status string `json:"status"`; TaskID string `json:"task_id"` } `json:"data"`
        Error any `json:"error"`
    }
    if _, e := p.Requester.SendRequest(req, &resp, false); e != nil { return "", e }
    if len(resp.Data) == 0 || strings.TrimSpace(resp.Data[0].TaskID) == "" {
        return "", common.StringErrorWrapperLocal("apimart missing task id", "upstream_error", http.StatusBadGateway)
    }
    return strings.TrimSpace(resp.Data[0].TaskID), nil
}

// 将 Veo 初始化映射到 ezlinkai 的 /v1/video/generations
// 需求：
// - Authorization: Bearer <key>
// - JSON: { model, instances:[{prompt, image?{bytesBase64Encoded,mimeType}}], parameters:{aspectRatio,durationSeconds,enhancePrompt,generateAudio,seed?} }
func (p *GeminiProvider) relayPredictLongRunningViaEzlinkai(modelName string, prompt string, seconds int, size string, refs []string) (string, *types.OpenAIErrorWithStatusCode) {
    base := strings.TrimSuffix(p.GetBaseURL(), "/")
    fullURL := base + "/v1/video/generations"

    // 组装 instances[0]
    inst := map[string]any{"prompt": prompt}
    if len(refs) > 0 && strings.TrimSpace(refs[0]) != "" {
        if mt, data, err := img.GetImageFromUrl(strings.TrimSpace(refs[0])); err == nil && strings.TrimSpace(data) != "" {
            inst["image"] = map[string]any{
                "bytesBase64Encoded": data,
                "mimeType":           mt,
            }
        }
    }

    // aspectRatio / durationSeconds
    ar, _ := parseSizeToGemini(size)
    m := strings.ToLower(strings.TrimSpace(modelName))
    // veo-3.0-generate-preview 仅支持 16:9（保护性回退）
    if m == "veo-3.0-generate-preview" && ar != "16:9" {
        ar = "16:9"
    }
    if seconds <= 0 { seconds = 6 }
    if seconds < 1 { seconds = 1 }
    if seconds > 8 { seconds = 8 }

    // 默认增强与生成音频，可后续从插件中读取开关
    params := map[string]any{
        "aspectRatio":     ar,
        "durationSeconds": seconds,
        "enhancePrompt":   true,
        "generateAudio":   true,
    }

    body := map[string]any{
        "model":      modelName,
        "instances":  []map[string]any{inst},
        "parameters": params,
    }

    headers := map[string]string{
        "Authorization": fmt.Sprintf("Bearer %s", p.Channel.Key),
        "Content-Type":  "application/json",
    }

    req, err := p.Requester.NewRequest(http.MethodPost, fullURL, p.Requester.WithBody(body), p.Requester.WithHeader(headers))
    if err != nil { return "", common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError) }
    defer req.Body.Close()

    // 响应示例：{ task_id: "...", task_status: "pending" }
    var resp map[string]any
    if _, e := p.Requester.SendRequest(req, &resp, false); e != nil { return "", e }

    // 容错提取 task_id
    taskID := ""
    if s, ok := resp["task_id"].(string); ok { taskID = strings.TrimSpace(s) }
    if taskID == "" {
        // 兼容 data.task_id
        if data, ok := resp["data"].(map[string]any); ok {
            if s, ok2 := data["task_id"].(string); ok2 { taskID = strings.TrimSpace(s) }
        }
    }
    if taskID == "" {
        return "", common.StringErrorWrapperLocal("ezlinkai missing task id", "upstream_error", http.StatusBadGateway)
    }
    return taskID, nil
}

// RelayModelAction 透传任意 models/<model>:<action> 原生请求（非流式）。
// 适配 Imagen 的 :predict 与 Veo 的 :predictLongRunning 初始化调用。
func (p *GeminiProvider) RelayModelAction(modelName, action string) (any, *types.OpenAIErrorWithStatusCode) {
    // 若为 Veo 初始化且配置为 sutui 上游，则改走 sutui /v1/videos
    if strings.HasPrefix(strings.ToLower(strings.TrimSpace(modelName)), "veo-") && strings.EqualFold(action, "predictLongRunning") && p.detectVeoVendor() == "sutui" {
        return p.relayPredictLongRunningViaSutui(modelName)
    }
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

// DetectVeoVendorForOps 暴露给 relay 用于 operations 分支判断
func (p *GeminiProvider) DetectVeoVendorForOps() string { return p.detectVeoVendor() }
