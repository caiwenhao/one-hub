package gemini

import (
    "fmt"
    "net/http"
    "one-api/common"
    "one-api/types"
    "strconv"
    "strings"
)

// --- Gemini Veo as OpenAI Video bridge ---

// normalizeVeoModel 将常见别名映射为官方模型代号
// 仅覆盖常用 3.1/3.0/2.0 以及 fast 变体；未知输入原样返回
func normalizeVeoModel(model string) string {
    m := strings.ToLower(strings.TrimSpace(model))
    // 已是官方代号
    if strings.HasPrefix(m, "veo-") && strings.Contains(m, ":") == false {
        // 形如 veo-3.1-generate-preview / veo-3.1-fast-generate-preview / veo-3.0-generate-001
        return model
    }
    // 常见别名 → 官方
    repl := map[string]string{
        "veo-3.1":            "veo-3.1-generate-preview",
        "veo3.1":             "veo-3.1-generate-preview",
        "veo3.1-fast":        "veo-3.1-fast-generate-preview",
        "veo-3.1-fast":       "veo-3.1-fast-generate-preview",
        "veo-3.0":            "veo-3.0-generate-001",
        "veo3.0":             "veo-3.0-generate-001",
        "veo3.0-fast":        "veo-3.0-fast-generate-001",
        "veo-3.0-fast":       "veo-3.0-fast-generate-001",
        "veo-2.0":            "veo-2.0-generate-001",
        "veo2.0":             "veo-2.0-generate-001",
    }
    if v, ok := repl[m]; ok {
        return v
    }
    return model
}

// parseSizeToGemini 将 1280x720/720x1280/1920x1080/1080x1920 转为 (aspectRatio,resolution)
func parseSizeToGemini(size string) (string, string) {
    s := strings.ToLower(strings.TrimSpace(size))
    switch s {
    case "1280x720", "1600x900":
        return "16:9", "720p"
    case "720x1280", "900x1600":
        return "9:16", "720p"
    case "1920x1080":
        return "16:9", "1080p"
    case "1080x1920":
        return "9:16", "1080p"
    }
    // 默认 16:9 720p
    return "16:9", "720p"
}

// buildVeoInitPayload 生成 predictLongRunning 的官方 JSON
func buildVeoInitPayload(prompt string, seconds int, size string) map[string]any {
    ar, res := parseSizeToGemini(size)
    // 官方 1080p 仅支持 8 秒（保护性约束）
    if res == "1080p" && seconds != 8 {
        seconds = 8
    }
    if seconds != 4 && seconds != 6 && seconds != 8 {
        // 回退到 6 秒
        seconds = 6
    }
    params := map[string]any{
        "aspectRatio":     ar,
        "resolution":      res,
        "durationSeconds": strconv.Itoa(seconds),
    }
    payload := map[string]any{
        "instances": []map[string]any{{
            "prompt": prompt,
        }},
        "parameters": params,
    }
    return payload
}

func (p *GeminiProvider) CreateVideo(request *types.VideoCreateRequest) (*types.VideoJob, *types.OpenAIErrorWithStatusCode) {
    modelName := normalizeVeoModel(request.Model)
    // 仅允许 veo-* 族
    if !strings.HasPrefix(strings.ToLower(modelName), "veo-") {
        return nil, common.StringErrorWrapperLocal("model not supported by Gemini Veo adapter", "invalid_model", http.StatusBadRequest)
    }

    base := strings.TrimSuffix(p.GetBaseURL(), "/")
    version := "v1beta"
    if p.Channel.Other != "" { version = p.Channel.Other }
    fullURL := fmt.Sprintf("%s/%s/models/%s:predictLongRunning", base, version, modelName)

    headers := p.GetRequestHeaders()
    body := buildVeoInitPayload(request.Prompt, request.Seconds, request.Size)
    req, err := p.Requester.NewRequest(http.MethodPost, fullURL, p.Requester.WithBody(body), p.Requester.WithHeader(headers))
    if err != nil { return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError) }
    defer req.Body.Close()

    var resp struct{ Name string `json:"name"` }
    if _, e := p.Requester.SendRequest(req, &resp, false); e != nil {
        return nil, e
    }
    name := strings.TrimSpace(resp.Name) // e.g. operations/abcdefg
    if name == "" {
        return nil, common.StringErrorWrapperLocal("missing operation name", "upstream_error", http.StatusBadGateway)
    }
    id := strings.TrimPrefix(name, "operations/")
    job := &types.VideoJob{
        ID:       id,
        Object:   "video",
        Status:   "queued",
        Model:    modelName,
        Prompt:   request.Prompt,
        Seconds:  request.Seconds,
        Size:     request.Size,
        Progress: 0,
    }
    return job, nil
}

func (p *GeminiProvider) RetrieveVideo(videoID string) (*types.VideoJob, *types.OpenAIErrorWithStatusCode) {
    base := strings.TrimSuffix(p.GetBaseURL(), "/")
    version := "v1beta"
    if p.Channel.Other != "" { version = p.Channel.Other }
    // operations/{id}
    fullURL := fmt.Sprintf("%s/%s/operations/%s", base, version, strings.TrimSpace(videoID))
    headers := p.GetRequestHeaders()
    req, err := p.Requester.NewRequest(http.MethodGet, fullURL, p.Requester.WithHeader(headers))
    if err != nil { return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError) }
    defer req.Body.Close()

    var data map[string]any
    if _, e := p.Requester.SendRequest(req, &data, false); e != nil {
        return nil, e
    }

    // 基本骨架
    job := &types.VideoJob{ ID: strings.TrimSpace(videoID), Object: "video" }

    // 解析 done
    done := false
    if v, ok := data["done"].(bool); ok { done = v }
    if !done {
        job.Status = "in_progress"
        job.Progress = 0
        return job, nil
    }

    job.Status = "completed"
    job.Progress = 100

    // 提取 uri 与元数据
    var seconds int
    var size string
    var uri string

    if respAny, ok := data["response"].(map[string]any); ok {
        if gvr, ok2 := respAny["generateVideoResponse"].(map[string]any); ok2 {
            if samples, ok3 := gvr["generatedSamples"].([]any); ok3 && len(samples) > 0 {
                if sample, ok4 := samples[0].(map[string]any); ok4 {
                    if vObj, ok5 := sample["video"].(map[string]any); ok5 {
                        if u, ok6 := vObj["uri"].(string); ok6 { uri = strings.TrimSpace(u) }
                    }
                    if meta, ok7 := sample["metadata"].(map[string]any); ok7 {
                        if dur, ok8 := meta["durationSeconds"].(float64); ok8 && dur > 0 { seconds = int(dur) }
                        if s, ok9 := meta["size"].(string); ok9 { size = strings.TrimSpace(s) }
                    }
                }
            }
        }
    }

    if seconds > 0 { job.Seconds = seconds }
    if size != "" { job.Size = size }
    if uri != "" {
        job.Result = &types.VideoJobResult{ VideoURL: uri }
    }
    return job, nil
}

func (p *GeminiProvider) DownloadVideo(videoID string, variant string) (*http.Response, *types.OpenAIErrorWithStatusCode) {
    // 先获取最新的 uri
    job, err := p.RetrieveVideo(videoID)
    if err != nil { return nil, err }
    if job == nil || job.Result == nil || strings.TrimSpace(job.Result.VideoURL) == "" {
        return nil, common.StringErrorWrapperLocal("video not ready", "video_not_ready", http.StatusBadRequest)
    }
    headers := p.GetRequestHeaders()
    // 直连视频 URL 并携带 x-goog-api-key（官方示例行为）
    req, e := p.Requester.NewRequest(http.MethodGet, job.Result.VideoURL, p.Requester.WithHeader(headers))
    if e != nil { return nil, common.ErrorWrapper(e, "new_request_failed", http.StatusInternalServerError) }
    // 不要提前关闭，交给上层转发
    return p.Requester.SendRequestRaw(req)
}

// 尚未对接：Veo 扩展/Remix、列表、删除
func (p *GeminiProvider) RemixVideo(videoID string, prompt string) (*types.VideoJob, *types.OpenAIErrorWithStatusCode) {
    return nil, common.StringErrorWrapperLocal("not implemented for Gemini Veo", "not_implemented", http.StatusNotImplemented)
}
func (p *GeminiProvider) ListVideos(after string, limit int, order string) (*types.VideoList, *types.OpenAIErrorWithStatusCode) {
    return nil, common.StringErrorWrapperLocal("not implemented for Gemini Veo", "not_implemented", http.StatusNotImplemented)
}
func (p *GeminiProvider) DeleteVideo(videoID string) (*types.VideoJob, *types.OpenAIErrorWithStatusCode) {
    return nil, common.StringErrorWrapperLocal("not implemented for Gemini Veo", "not_implemented", http.StatusNotImplemented)
}
