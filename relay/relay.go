package relay

import (
    "net/http"
    "one-api/common"
    "one-api/common/config"
    "one-api/model"
    "one-api/providers/azure"
    "one-api/providers/openai"
    miniProvider "one-api/providers/minimaxi"
    "encoding/json"
    "net/url"
    "io"
    "fmt"
    "strings"
    "time"
    "strconv"

    "github.com/gin-gonic/gin"
)

// 仅中继转发；在进入上游转发前，针对部分路径进行内置处理
func RelayOnly(c *gin.Context) {
    // 特判文件接口，交由定制逻辑，避免强制指定渠道
    if c.Request.URL.Path == "/v1/files/resolve" {
        ResolveFileMapping(c)
        return
    }
    if c.Request.URL.Path == "/v1/files/retrieve" || c.Request.URL.Path == "/v1/files/retrieve_content" {
        if handleFilesRetrieveCompat(c) {
            return
        }
    }
	provider, _, fail := GetProvider(c, "")
	if fail != nil {
		common.AbortWithMessage(c, http.StatusServiceUnavailable, fail.Error())
		return
	}

	channel := provider.GetChannel()
	if channel.Type != config.ChannelTypeOpenAI && channel.Type != config.ChannelTypeAzure {
		common.AbortWithMessage(c, http.StatusServiceUnavailable, "provider must be of type azureopenai or openai")
		return
	}

	// 获取请求的path
	url := ""
	path := c.Request.URL.Path
	openAIProvider, ok := provider.(*openai.OpenAIProvider)
	if !ok {
		azureProvider, ok := provider.(*azure.AzureProvider)
		if !ok {
			common.AbortWithMessage(c, http.StatusServiceUnavailable, "provider must be of type openai")
			return
		}
		url = azureProvider.GetFullRequestURL(path, "")
	} else {
		url = openAIProvider.GetFullRequestURL(path, "")
	}

	headers := c.Request.Header
	mapHeaders := provider.GetRequestHeaders()
	// 设置请求头
	for k, v := range headers {
		if _, ok := mapHeaders[k]; ok {
			continue
		}
		mapHeaders[k] = strings.Join(v, ", ")
	}

	requester := provider.GetRequester()
	req, err := requester.NewRequest(c.Request.Method, url, requester.WithBody(c.Request.Body), requester.WithHeader(mapHeaders))
	if err != nil {
		common.AbortWithMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	defer req.Body.Close()

	response, errWithCode := requester.SendRequestRaw(req)
	if errWithCode != nil {
		newErrWithCode := FilterOpenAIErr(c, errWithCode)
		relayResponseWithOpenAIErr(c, &newErrWithCode)
		return
	}

	errWithCode = responseMultipart(c, response)

	if errWithCode != nil {
		newErrWithCode := FilterOpenAIErr(c, errWithCode)
		relayResponseWithOpenAIErr(c, &newErrWithCode)
		return
	}

	requestTime := 0
	requestStartTimeValue := c.Request.Context().Value("requestStartTime")
	if requestStartTimeValue != nil {
		requestStartTime, ok := requestStartTimeValue.(time.Time)
		if ok {
			requestTime = int(time.Since(requestStartTime).Milliseconds())
		}
	}
	model.RecordConsumeLog(c.Request.Context(), c.GetInt("id"), c.GetInt("channel_id"), 0, 0, "", c.GetString("token_name"), 0, "中继:"+path, requestTime, false, nil, c.ClientIP())

}

// 兼容 /v1/files/retrieve(_content) 到 MiniMax 文件检索/下载
// 返回 true 表示已处理
func handleFilesRetrieveCompat(c *gin.Context) bool {
    fileID := strings.TrimSpace(c.Query("file_id"))
    if fileID == "" {
        common.AbortWithMessage(c, http.StatusBadRequest, "file_id is required")
        return true
    }

    userID := c.GetInt("id")

    // 1) 先通过 artifact 反查渠道
    var channel *model.Channel
    // 显式指定渠道优先（便于快速定位）
    if chStr := strings.TrimSpace(c.Query("channel_id")); chStr != "" {
        if chID, err := strconv.Atoi(chStr); err == nil && chID > 0 {
            if ch, err2 := model.GetChannelById(chID); err2 == nil {
                channel = ch
            }
        }
    }
    if a, err := model.GetArtifactByFileID(userID, model.TaskPlatformMiniMax, fileID); err == nil && a != nil && a.ChannelId > 0 && channel == nil {
        if ch, err2 := model.GetChannelById(a.ChannelId); err2 == nil {
            channel = ch
        }
    }
    // 2) 再用历史任务表反查（兼容旧数据）
    if channel == nil {
        if t, err := model.GetMiniMaxTaskByFileID(userID, fileID); err == nil && t != nil && t.ChannelId > 0 {
            if ch, err2 := model.GetChannelById(t.ChannelId); err2 == nil {
                channel = ch
            }
        }
    }
    if channel == nil {
        common.AbortWithMessage(c, http.StatusNotFound, "未找到包含该文件的任务或渠道")
        return true
    }

    // 解析上游
    upstream := parseMiniMaxUpstream(channel)
    if upstream == "ppinfra" {
        // 直接用 artifact 直链构造响应/下载
        a, _ := model.GetArtifactByFileID(userID, model.TaskPlatformMiniMax, fileID)
        if a == nil || strings.TrimSpace(a.DownloadURL) == "" {
            common.AbortWithMessage(c, http.StatusNotFound, "下载链接不存在或已过期，请先查询任务获取直链")
            return true
        }
        if c.Request.URL.Path == "/v1/files/retrieve" {
            payload := map[string]interface{}{
                "file": map[string]interface{}{
                    "file_id":      fileID,
                    "bytes":        0,
                    "created_at":   time.Now().Unix(),
                    "filename":     "video.mp4",
                    "purpose":      "video",
                    "download_url": a.DownloadURL,
                },
                "base_resp": map[string]interface{}{"status_code": 0, "status_msg": "success"},
            }
            c.JSON(http.StatusOK, payload)
            recordFilesConsumeLog(c, channel.Id)
            return true
        }
        // retrieve_content：拉取直链并回传
        resp, err := http.Get(a.DownloadURL)
        if err != nil {
            common.AbortWithMessage(c, http.StatusBadGateway, err.Error())
            return true
        }
        defer resp.Body.Close()
        if resp.StatusCode < 200 || resp.StatusCode >= 300 {
            body, _ := io.ReadAll(io.LimitReader(resp.Body, 4<<20))
            common.AbortWithMessage(c, http.StatusBadGateway, fmt.Sprintf("上游返回状态 %d: %s", resp.StatusCode, strings.TrimSpace(string(body))))
            return true
        }
        if ct := resp.Header.Get("Content-Type"); ct != "" { c.Header("Content-Type", ct) } else { c.Header("Content-Type", "application/octet-stream") }
        if cl := resp.Header.Get("Content-Length"); cl != "" { c.Header("Content-Length", cl) }
        c.Status(http.StatusOK)
        io.Copy(c.Writer, resp.Body)
        recordFilesConsumeLog(c, channel.Id)
        return true
    }

    // 官方上游：调用官方 /v1/files/retrieve；内容下载则二段式
    c.Set("specific_channel_id", channel.Id)

    // 处理 /v1/files/retrieve：在此直接转发到上游，避免进入通用转发对 provider 类型的限制
    if c.Request.URL.Path == "/v1/files/retrieve" {
        provider, _, fail := GetProvider(c, "")
        if fail != nil || provider == nil {
            common.AbortWithMessage(c, http.StatusServiceUnavailable, "provider not found")
            return true
        }

        // MiniMax 官方上游：直接通过 provider 的视频客户端取回 JSON
        if mini, ok := provider.(*miniProvider.MiniMaxProvider); ok && mini.GetVideoClient() != nil {
            respObj, errWrap := mini.GetVideoClient().RetrieveFile(fileID)
            if errWrap != nil {
                common.AbortWithMessage(c, http.StatusBadGateway, errWrap.Message)
                return true
            }
            // 建立/刷新 file_id → channel 映射
            if userID > 0 {
                _ = model.UpsertTaskArtifact(model.DB, &model.TaskArtifact{
                    Platform:     model.TaskPlatformMiniMax,
                    UserId:       userID,
                    ChannelId:    channel.Id,
                    TaskID:       "",
                    FileID:       fileID,
                    ArtifactType: "video",
                    DownloadURL:  respObj.File.DownloadURL,
                    TTLAt:        0,
                })
            }
            c.JSON(http.StatusOK, respObj)
            recordFilesConsumeLog(c, channel.Id)
            return true
        }

        // OpenAI/Azure 走直透
        var baseURL string
        if p, ok := provider.(*openai.OpenAIProvider); ok {
            baseURL = p.GetFullRequestURL("/v1/files/retrieve", "")
        } else if az, ok := provider.(*azure.AzureProvider); ok {
            baseURL = az.GetFullRequestURL("/v1/files/retrieve", "")
        } else {
            common.AbortWithMessage(c, http.StatusServiceUnavailable, "unsupported provider for files api")
            return true
        }

        fullURL := baseURL
        if strings.Contains(baseURL, "?") {
            fullURL = baseURL + "&file_id=" + url.QueryEscape(fileID)
        } else {
            fullURL = baseURL + "?file_id=" + url.QueryEscape(fileID)
        }

        headers := provider.GetRequestHeaders()
        for k, v := range c.Request.Header {
            if _, ok := headers[k]; ok {
                continue
            }
            headers[k] = strings.Join(v, ", ")
        }

        requester := provider.GetRequester()
        req, err := requester.NewRequest(http.MethodGet, fullURL, requester.WithHeader(headers))
        if err != nil {
            common.AbortWithMessage(c, http.StatusBadRequest, err.Error())
            return true
        }
        resp, errWith := requester.SendRequestRaw(req)
        if errWith != nil {
            relayResponseWithOpenAIErr(c, errWith)
            return true
        }
        defer resp.Body.Close()
        // 透传 JSON 响应
        if ct := resp.Header.Get("Content-Type"); ct != "" {
            c.Header("Content-Type", ct)
        } else {
            c.Header("Content-Type", "application/json")
        }
        c.Status(resp.StatusCode)
        io.Copy(c.Writer, resp.Body)
        recordFilesConsumeLog(c, channel.Id)
        return true
    }

    // 内容下载：先请求 /v1/files/retrieve 获取 download_url
    provider, _, fail := GetProvider(c, "")
    if fail != nil || provider == nil {
        common.AbortWithMessage(c, http.StatusServiceUnavailable, "provider not found")
        return true
    }
    var baseURL string
    if p, ok := provider.(*openai.OpenAIProvider); ok {
        baseURL = p.GetFullRequestURL("/v1/files/retrieve", "")
    } else if az, ok := provider.(*azure.AzureProvider); ok {
        baseURL = az.GetFullRequestURL("/v1/files/retrieve", "")
    } else {
        common.AbortWithMessage(c, http.StatusServiceUnavailable, "unsupported provider for files api")
        return true
    }
    fullURL := baseURL
    if strings.Contains(baseURL, "?") {
        fullURL = baseURL + "&file_id=" + url.QueryEscape(fileID)
    } else {
        fullURL = baseURL + "?file_id=" + url.QueryEscape(fileID)
    }
    headers := provider.GetRequestHeaders()
    // 透传必要头
    for k, v := range c.Request.Header {
        if _, ok := headers[k]; ok { continue }
        headers[k] = strings.Join(v, ", ")
    }
    requester := provider.GetRequester()
    req, err := requester.NewRequest(http.MethodGet, fullURL, requester.WithHeader(headers))
    if err != nil { common.AbortWithMessage(c, http.StatusBadRequest, err.Error()); return true }
    var meta struct{ File struct{ DownloadURL string `json:"download_url"` } `json:"file"` }
    _, errWith := requester.SendRequest(req, &meta, false)
    if errWith != nil { relayResponseWithOpenAIErr(c, errWith); return true }
    if strings.TrimSpace(meta.File.DownloadURL) == "" { common.AbortWithMessage(c, http.StatusBadGateway, "上游未返回下载链接"); return true }
    // 拉取内容
    resp, err := http.Get(meta.File.DownloadURL)
    if err != nil { common.AbortWithMessage(c, http.StatusBadGateway, err.Error()); return true }
    defer resp.Body.Close()
    if resp.StatusCode < 200 || resp.StatusCode >= 300 {
        body, _ := io.ReadAll(io.LimitReader(resp.Body, 4<<20))
        common.AbortWithMessage(c, http.StatusBadGateway, fmt.Sprintf("上游返回状态 %d: %s", resp.StatusCode, strings.TrimSpace(string(body))))
        return true
    }
    if ct := resp.Header.Get("Content-Type"); ct != "" { c.Header("Content-Type", ct) } else { c.Header("Content-Type", "application/octet-stream") }
    if cl := resp.Header.Get("Content-Length"); cl != "" { c.Header("Content-Length", cl) }
    c.Status(http.StatusOK)
    io.Copy(c.Writer, resp.Body)
    recordFilesConsumeLog(c, channel.Id)
    return true
}

func parseMiniMaxUpstream(channel *model.Channel) string {
    raw := channel.GetCustomParameter()
    if strings.TrimSpace(raw) == "" { return "official" }
    var payload map[string]json.RawMessage
    if err := json.Unmarshal([]byte(raw), &payload); err != nil { return "official" }
    if vRaw, ok := payload["video"]; ok {
        var video map[string]interface{}
        if err := json.Unmarshal(vRaw, &video); err == nil {
            if up, ok2 := video["upstream"].(string); ok2 && strings.TrimSpace(up) != "" {
                return strings.ToLower(strings.TrimSpace(up))
            }
        }
    }
    if upRaw, ok := payload["upstream"]; ok {
        var up string
        if err := json.Unmarshal(upRaw, &up); err == nil && strings.TrimSpace(up) != "" {
            return strings.ToLower(strings.TrimSpace(up))
        }
    }
    return "official"
}

// recordFilesConsumeLog 记录文件接口的消费日志（不计配额与 Token，仅统计请求路径与耗时）
func recordFilesConsumeLog(c *gin.Context, channelId int) {
    requestTime := 0
    requestStartTimeValue := c.Request.Context().Value("requestStartTime")
    if requestStartTimeValue != nil {
        if requestStartTime, ok := requestStartTimeValue.(time.Time); ok {
            requestTime = int(time.Since(requestStartTime).Milliseconds())
        }
    }
    model.RecordConsumeLog(c.Request.Context(), c.GetInt("id"), channelId, 0, 0, "", c.GetString("token_name"), 0, "文件:"+c.Request.URL.Path, requestTime, false, nil, c.ClientIP())
}
