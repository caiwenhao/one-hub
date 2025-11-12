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

// 不做别名归一化：直接返回原模型名
func normalizeVeoModel(model string) string { return model }

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
func buildVeoInitPayload(prompt string, seconds int, size string, refs []string, extendVideoURL string) map[string]any {
    ar, res := parseSizeToGemini(size)
    // 续写场景固定 720p，忽略 seconds（上游固定 +7s）
    if strings.TrimSpace(extendVideoURL) != "" {
        res = "720p"
    } else {
        // 官方 1080p 仅支持 8 秒（保护性约束）
        if res == "1080p" && seconds != 8 {
            seconds = 8
        }
        if seconds != 4 && seconds != 6 && seconds != 8 {
            // 回退到 6 秒
            seconds = 6
        }
    }
    params := map[string]any{
        "aspectRatio":     ar,
        "resolution":      res,
    }
    if strings.TrimSpace(extendVideoURL) == "" {
        params["durationSeconds"] = strconv.Itoa(seconds)
    }
    payload := map[string]any{
        "instances": []map[string]any{{"prompt": prompt}},
        "parameters": params,
    }
    inst := payload["instances"].([]map[string]any)[0]
    // 参考图/首尾帧
    if len(refs) == 1 {
        inst["image"] = map[string]any{"uri": strings.TrimSpace(refs[0])}
    } else if len(refs) >= 2 {
        // 首尾帧：首帧作为 image、尾帧作为 lastFrame
        inst["image"] = map[string]any{"uri": strings.TrimSpace(refs[0])}
        payload["parameters"].(map[string]any)["lastFrame"] = map[string]any{"uri": strings.TrimSpace(refs[1])}
        if len(refs) > 2 {
            // 额外参考图
            refImgs := make([]map[string]any, 0, len(refs)-2)
            for _, r := range refs[2:] {
                refImgs = append(refImgs, map[string]any{"image": map[string]any{"uri": strings.TrimSpace(r)}})
            }
            payload["parameters"].(map[string]any)["referenceImages"] = refImgs
        }
    }
    // 扩展视频
    if strings.TrimSpace(extendVideoURL) != "" {
        payload["video"] = map[string]any{"uri": extendVideoURL}
    }
    return payload
}

func (p *GeminiProvider) CreateVideo(request *types.VideoCreateRequest) (*types.VideoJob, *types.OpenAIErrorWithStatusCode) {
    modelName := strings.TrimSpace(normalizeVeoModel(request.Model))

    // 预处理：扩展视频需要获得历史成片的 uri
    extendURI := ""
    if strings.TrimSpace(request.ExtendFrom) != "" {
        prev, e := p.RetrieveVideo(strings.TrimSpace(request.ExtendFrom))
        if e != nil { return nil, e }
        if prev == nil || !strings.EqualFold(prev.Status, "completed") {
            return nil, common.StringErrorWrapperLocal("extend_from target must be a completed Veo video", "invalid_arguments", http.StatusUnprocessableEntity)
        }
        if strings.TrimSpace(prev.VideoURL) != "" {
            extendURI = prev.VideoURL
        } else if prev.Result != nil && strings.TrimSpace(prev.Result.VideoURL) != "" {
            extendURI = prev.Result.VideoURL
        }
        if extendURI == "" {
            return nil, common.StringErrorWrapperLocal("extend_from target missing video uri", "invalid_arguments", http.StatusUnprocessableEntity)
        }
    }

    // 分供应商处理
    vendor := p.detectVeoVendor()

    // 允许的型号按供应商划分：
    allowed := map[string]bool{
        "veo-3.1-generate-preview":      true,
        "veo-3.1-fast-generate-preview": true,
    }
    if vendor == "ezlinkai" {
        // ezlinkai 支持 3.0 preview/fast 族
        allowed = map[string]bool{
            "veo-3.0-generate-preview":      true,
            "veo-3.0-fast-generate-preview": true,
            // 兼容 3.1 如需
            "veo-3.1-generate-preview":      true,
            "veo-3.1-fast-generate-preview": true,
        }
    }
    if !allowed[strings.ToLower(modelName)] {
        return nil, common.StringErrorWrapperLocal("model not supported by Gemini Veo adapter", "invalid_model", http.StatusBadRequest)
    }
    var id string
    switch vendor {
    case "apimart":
        if extendURI != "" {
            return nil, common.StringErrorWrapperLocal("extend video is not supported for apimart vendor", "not_implemented", http.StatusNotImplemented)
        }
        taskID, e := p.relayPredictLongRunningViaApimart(modelName, request.Prompt, request.Seconds, request.Size, request.InputReference)
        if e != nil { return nil, e }
        id = taskID
    case "ezlinkai":
        if extendURI != "" {
            return nil, common.StringErrorWrapperLocal("extend video is not supported for ezlinkai vendor", "not_implemented", http.StatusNotImplemented)
        }
        taskID, e := p.relayPredictLongRunningViaEzlinkai(modelName, request.Prompt, request.Seconds, request.Size, request.InputReference)
        if e != nil { return nil, e }
        id = taskID
    default:
        base := strings.TrimSuffix(p.GetBaseURL(), "/")
        version := "v1beta"
        if p.Channel.Other != "" { version = p.Channel.Other }
        fullURL := fmt.Sprintf("%s/%s/models/%s:predictLongRunning", base, version, modelName)
        headers := p.GetRequestHeaders()
        body := buildVeoInitPayload(request.Prompt, request.Seconds, request.Size, request.InputReference, extendURI)
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
        id = strings.TrimPrefix(name, "operations/")
    }
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
    vendor := p.detectVeoVendor()
    var data map[string]any
    if vendor == "apimart" {
        // GET /v1/tasks/{id}
        base := strings.TrimSuffix(p.GetBaseURL(), "/")
        fullURL := fmt.Sprintf("%s/v1/tasks/%s", base, strings.TrimSpace(videoID))
        headers := map[string]string{"Authorization": fmt.Sprintf("Bearer %s", p.Channel.Key)}
        req, err := p.Requester.NewRequest(http.MethodGet, fullURL, p.Requester.WithHeader(headers))
        if err != nil { return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError) }
        if req.Body != nil { defer req.Body.Close() }
        if _, e := p.Requester.SendRequest(req, &data, false); e != nil { return nil, e }
    } else if vendor == "ezlinkai" {
        base := strings.TrimSuffix(p.GetBaseURL(), "/")
        fullURL := fmt.Sprintf("%s/v1/video/generations/result?taskid=%s", base, strings.TrimSpace(videoID))
        headers := map[string]string{"Authorization": fmt.Sprintf("Bearer %s", p.Channel.Key)}
        req, err := p.Requester.NewRequest(http.MethodGet, fullURL, p.Requester.WithHeader(headers))
        if err != nil { return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError) }
        if req.Body != nil { defer req.Body.Close() }
        if _, e := p.Requester.SendRequest(req, &data, false); e != nil { return nil, e }
    } else {
        base := strings.TrimSuffix(p.GetBaseURL(), "/")
        version := "v1beta"
        if p.Channel.Other != "" { version = p.Channel.Other }
        fullURL := fmt.Sprintf("%s/%s/operations/%s", base, version, strings.TrimSpace(videoID))
        headers := p.GetRequestHeaders()
        req, err := p.Requester.NewRequest(http.MethodGet, fullURL, p.Requester.WithHeader(headers))
        if err != nil { return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError) }
        if req.Body != nil { defer req.Body.Close() }
        if _, e := p.Requester.SendRequest(req, &data, false); e != nil { return nil, e }
    }

    // 基本骨架
    job := &types.VideoJob{ ID: strings.TrimSpace(videoID), Object: "video" }

    if vendor == "apimart" {
        // 适配 apimart 任务查询结构
        // 典型：{ code:200, data:{ id,status,progress,result:{videos:[{url,expires_at}]}, ... } }
        var node any
        if v, ok := data["data"]; ok { node = v } else { node = data }
        if m, ok := node.(map[string]any); ok {
            st, _ := m["status"].(string)
            job.Status = strings.ToLower(strings.TrimSpace(st))
            if prog, ok2 := m["progress"].(float64); ok2 { job.Progress = prog }
            if job.Status == "completed" {
                job.Progress = 100
                if res, ok3 := m["result"].(map[string]any); ok3 {
                    // 优先 videos
                    if vs, ok4 := res["videos"].([]any); ok4 && len(vs) > 0 {
                        if v0, ok5 := vs[0].(map[string]any); ok5 {
                            // url 可能为 string 或 [string]
                            if u, ok6 := v0["url"].(string); ok6 {
                                job.VideoURL = strings.TrimSpace(u)
                            } else if us, ok7 := v0["url"].([]any); ok7 {
                                for _, it := range us {
                                    if s, ok := it.(string); ok && strings.TrimSpace(s) != "" {
                                        job.VideoURL = strings.TrimSpace(s)
                                        break
                                    }
                                }
                            } else if us2, ok8 := v0["url"].([]string); ok8 && len(us2) > 0 {
                                if strings.TrimSpace(us2[0]) != "" { job.VideoURL = strings.TrimSpace(us2[0]) }
                            }
                        }
                    }
                }
            } else if job.Status == "submitted" || job.Status == "pending" || job.Status == "processing" {
                job.Status = "in_progress"
            }
        }
        return job, nil
    } else if vendor == "ezlinkai" {
        // 适配 ezlinkai 查询结构
        // 典型：{ task_id, task_status, video_result, video_results:[], duration }
        st, _ := data["task_status"].(string)
        status := strings.ToLower(strings.TrimSpace(st))
        switch status {
        case "succeed", "succeeded", "success":
            job.Status = "completed"
            job.Progress = 100
        case "failed":
            job.Status = "failed"
        default:
            job.Status = "in_progress"
        }
        if job.Status == "completed" {
            if u, ok := data["video_result"].(string); ok && strings.TrimSpace(u) != "" {
                job.VideoURL = strings.TrimSpace(u)
            }
            if job.VideoURL == "" {
                if arr, ok := data["video_results"].([]any); ok && len(arr) > 0 {
                    if s, ok2 := arr[0].(string); ok2 { job.VideoURL = strings.TrimSpace(s) }
                }
            }
            if d, ok := data["duration"].(string); ok && strings.TrimSpace(d) != "" {
                if n, err := strconv.Atoi(strings.Split(strings.TrimSpace(d), ".")[0]); err == nil { job.Seconds = n }
            }
        }
        return job, nil
    }

    // 解析 done（Google 官方）
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
        job.VideoURL = uri
        job.Result = &types.VideoJobResult{ VideoURL: uri }
    }
    return job, nil
}

func (p *GeminiProvider) DownloadVideo(videoID string, variant string) (*http.Response, *types.OpenAIErrorWithStatusCode) {
    // 先获取最新的 uri
    job, err := p.RetrieveVideo(videoID)
    if err != nil { return nil, err }
    if job == nil || (strings.TrimSpace(job.VideoURL) == "" && (job.Result == nil || strings.TrimSpace(job.Result.VideoURL) == "")) {
        return nil, common.StringErrorWrapperLocal("video not ready", "video_not_ready", http.StatusBadRequest)
    }
    // 直连视频 URL：google 需要 x-goog-api-key；apimart 一般无需额外头
    vendor := p.detectVeoVendor()
    var headers map[string]string
    if vendor == "apimart" || vendor == "ezlinkai" {
        headers = map[string]string{}
    } else {
        headers = p.GetRequestHeaders()
    }
    videoURL := job.VideoURL
    if videoURL == "" && job.Result != nil { videoURL = job.Result.VideoURL }
    req, e := p.Requester.NewRequest(http.MethodGet, videoURL, p.Requester.WithHeader(headers))
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
