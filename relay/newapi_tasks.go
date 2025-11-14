package relay

import (
    "encoding/json"
    "io"
    "io/ioutil"
    "net/http"
    "net/url"
    "path"
    "strings"
    "time"

    "one-api/common"
    "one-api/common/config"
    "one-api/common/logger"
    "one-api/common/storage"
    "one-api/common/utils"
    "one-api/model"

	"github.com/gin-gonic/gin"
)

// NewAPITaskRetrieve 直透上游 /v1/tasks/{id}（仅 NewAPI 渠道）。
// 不计费，仅查询任务状态。
func NewAPITaskRetrieve(c *gin.Context) {
	// 非 NewAPI 渠道作为 /v1/videos/{id} 的别名
	if c.GetInt("channel_type") != config.ChannelTypeNewAPI {
		VideoRetrieve(c)
		return
	}

	// 仅允许 NewAPI 渠道类型
	c.Set("allow_channel_type", []int{config.ChannelTypeNewAPI})

    taskID := strings.TrimSpace(c.Param("id"))
    if taskID == "" {
        common.AbortWithMessage(c, http.StatusBadRequest, "task id is required")
        return
    }

    // 支持平台任务ID：task_<ULID>/裸ULID/旧base36 → 转为上游ID
    if strings.HasPrefix(strings.ToLower(taskID), utils.PlatformTaskPrefix) {
        if t, _ := model.GetTaskByPlatformTaskID(model.TaskPlatformSora, c.GetInt("id"), utils.StripTaskPrefix(taskID)); t != nil {
            taskID = t.TaskID
        }
    } else if id, ok := utils.DecodePlatformTaskID(taskID); ok {
        if t, _ := model.GetTaskByID(id); t != nil && t.Platform == model.TaskPlatformSora && t.UserId == c.GetInt("id") {
            taskID = t.TaskID
        }
    } else if utils.IsULID(taskID) {
        if t, _ := model.GetTaskByPlatformTaskID(model.TaskPlatformSora, c.GetInt("id"), taskID); t != nil {
            taskID = t.TaskID
        }
    }

	// 选择一个 NewAPI 渠道。使用一个常见模型名进行选路（要求该渠道包含该模型）。
	provider, _, err := GetProvider(c, "sora-2")
	if err != nil {
		common.AbortWithMessage(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	base := strings.TrimSuffix(provider.GetChannel().GetBaseURL(), "/")
	if base == "" {
		common.AbortWithMessage(c, http.StatusServiceUnavailable, "channel base_url is empty")
		return
	}
	fullURL := base + "/v1/tasks/" + taskID

	headers := provider.GetRequestHeaders()
	req, e := provider.GetRequester().NewRequest(http.MethodGet, fullURL, provider.GetRequester().WithHeader(headers))
	if e != nil {
		common.AbortWithMessage(c, http.StatusInternalServerError, "new_request_failed")
		return
	}

	// 直透响应（原样 JSON），如配置开启则对结果中的图片 URL 进行重写/镜像
	resp, errWith := provider.GetRequester().SendRequestRaw(req)
	if errWith != nil {
		status := errWith.StatusCode
		if status == 0 {
			status = http.StatusBadGateway
		}
		c.Status(status)
		c.Writer.Header().Set("Content-Type", "application/json")
		_, _ = c.Writer.Write([]byte(errWith.OpenAIError.Message))
		return
	}
	defer resp.Body.Close()

	// 仅在返回为 JSON 时尝试处理
	contentType := resp.Header.Get("Content-Type")
	bodyBytes, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		// 退化为直接复制
		for key, values := range resp.Header {
			for _, v := range values {
				c.Writer.Header().Add(key, v)
			}
		}
		c.Status(resp.StatusCode)
		if _, copyErr := c.Writer.Write(bodyBytes); copyErr != nil {
			logger.LogError(c.Request.Context(), "write_newapi_task_failed:"+copyErr.Error())
		}
		return
	}

	// 默认：原样转发
	finalBody := bodyBytes

	if strings.Contains(strings.ToLower(contentType), "application/json") {
		if modified, ok := maybeMirrorNewAPIImages(c, bodyBytes); ok {
			finalBody = modified
		}
	}

	for key, values := range resp.Header {
		// 重设 Content-Length 由 Go 自动处理，避免与修改后的长度不一致
		if strings.ToLower(key) == "content-length" {
			continue
		}
		for _, v := range values {
			c.Writer.Header().Add(key, v)
		}
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Status(resp.StatusCode)
	if _, copyErr := io.Copy(c.Writer, strings.NewReader(string(finalBody))); copyErr != nil {
		logger.LogError(c.Request.Context(), "copy_newapi_task_failed:"+copyErr.Error())
	}
}

// maybeMirrorNewAPIImages 在启用 newapi.mirror_image_to_storage=true 且配置了存储驱动时，
// 将 vendor 返回的 images[].url（string 或字符串数组）下载后上传到本地配置的存储，并替换为自有域名 URL。
// 返回 (修改后的 JSON, 是否已修改)。失败或未开启均返回 (nil,false)。
func maybeMirrorNewAPIImages(c *gin.Context, body []byte) ([]byte, bool) {
	if !config.NewAPIMirrorImageToStorage {
		return nil, false
	}

	// 仅允许代理的上游域名，避免 SSRF。默认 upload.apimart.ai，可通过面板配置。
	allowed := config.NewAPIAllowedAssetHosts
	if len(allowed) == 0 {
		allowed = []string{"upload.apimart.ai"}
	}
	isAllowedHost := func(host string) bool {
		host = strings.ToLower(host)
		for _, h := range allowed {
			if strings.EqualFold(host, h) {
				return true
			}
		}
		return false
	}

	var payload map[string]any
	if err := json.Unmarshal(body, &payload); err != nil {
		return nil, false
	}

	data, ok := payload["data"].(map[string]any)
	if !ok {
		return nil, false
	}
	result, ok := data["result"].(map[string]any)
	if !ok {
		return nil, false
	}
	images, ok := result["images"].([]any)
	if !ok || len(images) == 0 {
		return nil, false
	}

	// 下载 -> 上传 -> 替换 URL
	changed := false
	client := &http.Client{Timeout: 15 * time.Second}
	for i := range images {
		m, _ := images[i].(map[string]any)
		if m == nil {
			continue
		}

		switch v := m["url"].(type) {
		case string:
			if u := mirrorOneURL(c, client, v, isAllowedHost); u != "" {
				m["url"] = u
				changed = true
			}
		case []any:
			newArr := make([]any, 0, len(v))
			for _, item := range v {
				if s, ok := item.(string); ok {
					if u := mirrorOneURL(c, client, s, isAllowedHost); u != "" {
						newArr = append(newArr, u)
						changed = true
						continue
					}
				}
				newArr = append(newArr, item)
			}
			m["url"] = newArr
		}
	}

	if !changed {
		return nil, false
	}
	// 写回
	result["images"] = images
	data["result"] = result
	payload["data"] = data
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, false
	}
	return b, true
}

// mirrorOneURL 将指定 URL 下载并上传到已配置的存储，返回新 URL；失败返回空串。
func mirrorOneURL(c *gin.Context, client *http.Client, raw string, hostAllow func(string) bool) string {
	if raw == "" {
		return ""
	}
	u, err := url.Parse(raw)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return ""
	}
	if !hostAllow(u.Host) {
		return ""
	}

	req, _ := http.NewRequestWithContext(c.Request.Context(), http.MethodGet, raw, nil)
	resp, err := client.Do(req)
	if err != nil {
		logger.LogError(c.Request.Context(), "mirror_fetch_failed:"+err.Error())
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		logger.LogError(c.Request.Context(), "mirror_fetch_status:"+resp.Status)
		return ""
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.LogError(c.Request.Context(), "mirror_read_failed:"+err.Error())
		return ""
	}

	// 推断后缀
	ext := strings.ToLower(path.Ext(u.Path))
	if ext == "" {
		// 尝试根据 Content-Type 推断
		ct := resp.Header.Get("Content-Type")
		switch {
		case strings.Contains(ct, "image/png"):
			ext = ".png"
		case strings.Contains(ct, "image/jpeg"):
			ext = ".jpg"
		case strings.Contains(ct, "image/webp"):
			ext = ".webp"
		default:
			ext = ".png"
		}
	}
	filename := utils.GetUUID() + ext

	if url := storage.Upload(data, filename); url != "" {
		return url
	}
	return ""
}
