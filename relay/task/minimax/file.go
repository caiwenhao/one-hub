package minimax

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"one-api/common/config"
	"one-api/model"
	"one-api/providers"
	miniProvider "one-api/providers/minimaxi"
	"one-api/relay"

	"github.com/gin-gonic/gin"
)

// RelayFileRetrieve 代理 MiniMax 文件信息查询
func RelayFileRetrieve(c *gin.Context) {
	fileID := strings.TrimSpace(c.Query("file_id"))
	if fileID == "" {
		StringError(c, http.StatusBadRequest, "invalid_request", "file_id is required")
		return
	}

	resp, ok := retrieveMiniMaxFileMetadata(c, fileID)
	if !ok {
		return
	}

	writeJSONNoEscape(c, http.StatusOK, resp)
}

func RelayFileRetrieveContent(c *gin.Context) {
	fileID := strings.TrimSpace(c.Query("file_id"))
	if fileID == "" {
		StringError(c, http.StatusBadRequest, "invalid_request", "file_id is required")
		return
	}

	resp, ok := retrieveMiniMaxFileMetadata(c, fileID)
	if !ok {
		return
	}

	downloadURL := strings.TrimSpace(resp.File.DownloadURL)
	if downloadURL == "" {
		StringError(c, http.StatusNotFound, "download_url_missing", "上游未返回下载地址，可重试或稍后再试")
		return
	}

	req, err := http.NewRequestWithContext(c.Request.Context(), http.MethodGet, downloadURL, nil)
	if err != nil {
		StringError(c, http.StatusInternalServerError, "download_request_failed", err.Error())
		return
	}

	client := &http.Client{Timeout: 120 * time.Second}
	upstreamResp, err := client.Do(req)
	if err != nil {
		StringError(c, http.StatusBadGateway, "download_failed", err.Error())
		return
	}
	defer upstreamResp.Body.Close()

	if upstreamResp.StatusCode < http.StatusOK || upstreamResp.StatusCode >= http.StatusMultipleChoices {
		body, _ := io.ReadAll(io.LimitReader(upstreamResp.Body, 4<<20))
		StringError(c, http.StatusBadGateway, "download_failed", fmt.Sprintf("上游返回状态 %d: %s", upstreamResp.StatusCode, strings.TrimSpace(string(body))))
		return
	}

	if ct := upstreamResp.Header.Get("Content-Type"); ct != "" {
		c.Header("Content-Type", ct)
	} else {
		c.Header("Content-Type", "application/octet-stream")
	}
	if cl := upstreamResp.Header.Get("Content-Length"); cl != "" {
		c.Header("Content-Length", cl)
	}

	if filename := strings.TrimSpace(resp.File.Filename); filename != "" {
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	}

	c.Status(http.StatusOK)
	io.Copy(c.Writer, upstreamResp.Body)
}

func retrieveMiniMaxFileMetadata(c *gin.Context, fileID string) (*miniProvider.MiniMaxFileRetrieveResponse, bool) {
	userID := c.GetInt("id")

	if chStr := strings.TrimSpace(c.Query("channel_id")); chStr != "" {
		if chID, err := strconv.Atoi(chStr); err == nil && chID > 0 {
			if channel, err := model.GetChannelById(chID); err == nil && channel != nil {
				if resp, ok := retrieveFileViaChannel(c, channel, userID, fileID, true); ok {
					return resp, true
				}
			}
		}
	}

	taskID := strings.TrimSpace(c.Query("task_id"))
	var task *model.Task
	var err error
	if taskID != "" {
		task, err = model.GetTaskByTaskId(model.TaskPlatformMiniMax, userID, taskID)
		if err != nil {
			StringError(c, http.StatusInternalServerError, "task_query_failed", err.Error())
			return nil, false
		}
	}

	if task == nil {
		if a, err := model.GetArtifactByFileID(userID, model.TaskPlatformMiniMax, fileID); err == nil && a != nil && a.ChannelId > 0 {
			if channel, err := model.GetChannelById(a.ChannelId); err == nil && channel != nil {
				if resp, ok := retrieveFileViaChannel(c, channel, userID, fileID, false); ok {
					return resp, true
				}
			}
		}

		task, err = model.GetMiniMaxTaskByFileID(userID, fileID)
		if err != nil {
			StringError(c, http.StatusInternalServerError, "task_query_failed", err.Error())
			return nil, false
		}
	}

	if task == nil {
		group := strings.TrimSpace(c.GetString("token_group"))
		if channels, err := model.GetAllChannels(); err == nil {
			var official []*model.Channel
			var others []*model.Channel
			for _, ch := range channels {
				if ch == nil || ch.Status != config.ChannelStatusEnabled || ch.Type != config.ChannelTypeMiniMax {
					continue
				}
				if group != "" {
					okInGroup := false
					for _, g := range strings.Split(ch.Group, ",") {
						if strings.TrimSpace(g) == group {
							okInGroup = true
							break
						}
					}
					if !okInGroup {
						continue
					}
				}
				if isMiniMaxOfficialChannelForFile(ch) {
					official = append(official, ch)
				} else {
					others = append(others, ch)
				}
			}
			for _, ch := range official {
				if resp, ok := retrieveFileViaChannel(c, ch, userID, fileID, true); ok {
					return resp, true
				}
			}
			for _, ch := range others {
				if resp, ok := retrieveFileViaChannel(c, ch, userID, fileID, true); ok {
					return resp, true
				}
			}
		}

		c.Set("allow_channel_type", []int{config.ChannelTypeMiniMax})
		if provider, _, fail := relay.GetProvider(c, "MiniMax-M1"); fail == nil && provider != nil {
			if mini, ok := provider.(*miniProvider.MiniMaxProvider); ok && mini.GetVideoClient() != nil {
				if resp, errWithCode := mini.GetVideoClient().RetrieveFile(fileID); errWithCode == nil {
					upsertMiniMaxArtifact(userID, provider.GetChannel().Id, fileID, resp.File.DownloadURL)
					return resp, true
				}
			}
		}

		StringError(c, http.StatusNotFound, "task_not_found", "未找到包含该文件的任务（可加 channel_id 或 task_id 重试）")
		return nil, false
	}

	channel, channelErr := model.GetChannelById(task.ChannelId)
	if channelErr != nil {
		StringError(c, http.StatusServiceUnavailable, "channel_not_found", channelErr.Error())
		return nil, false
	}

	provider := providers.GetProvider(channel, c)
	mini, ok := provider.(*miniProvider.MiniMaxProvider)
	if !ok || mini.GetVideoClient() == nil {
		StringError(c, http.StatusServiceUnavailable, "provider_not_found", "provider not found")
		return nil, false
	}

	resp, errWithCode := mini.GetVideoClient().RetrieveFile(fileID)
	if errWithCode != nil {
		StringError(c, http.StatusInternalServerError, "file_retrieve_failed", errWithCode.Message)
		return nil, false
	}

	return resp, true
}

func retrieveFileViaChannel(c *gin.Context, channel *model.Channel, userID int, fileID string, upsert bool) (*miniProvider.MiniMaxFileRetrieveResponse, bool) {
    // 若为 PPInfra 上游，优先使用 artifact 中的直链构造响应，避免请求不存在的 /v1/files/retrieve
    if !isMiniMaxOfficialChannelForFile(channel) {
        if a, err := model.GetArtifactByFileID(userID, model.TaskPlatformMiniMax, fileID); err == nil && a != nil && strings.TrimSpace(a.DownloadURL) != "" {
            resp := &miniProvider.MiniMaxFileRetrieveResponse{
                File: miniProvider.MiniMaxFileObject{
                    FileID:      miniProvider.StringOrNumber(fileID),
                    Bytes:       0,
                    CreatedAt:   time.Now().Unix(),
                    Filename:    "video.mp4",
                    Purpose:     "video",
                    DownloadURL: a.DownloadURL,
                },
                BaseResp: miniProvider.BaseResp{StatusCode: 0, StatusMsg: "success"},
            }
            return resp, true
        }
    }

    provider := providers.GetProvider(channel, c)
    mini, ok := provider.(*miniProvider.MiniMaxProvider)
    if !ok || mini.GetVideoClient() == nil {
        return nil, false
    }
    resp, errWithCode := mini.GetVideoClient().RetrieveFile(fileID)
    if errWithCode != nil {
        return nil, false
    }
    if upsert {
        upsertMiniMaxArtifact(userID, channel.Id, fileID, resp.File.DownloadURL)
    }
    return resp, true
}

func upsertMiniMaxArtifact(userID, channelID int, fileID, downloadURL string) {
	if userID <= 0 || channelID <= 0 {
		return
	}
	_ = model.UpsertTaskArtifact(model.DB, &model.TaskArtifact{
		Platform:     model.TaskPlatformMiniMax,
		UserId:       userID,
		ChannelId:    channelID,
		TaskID:       "",
		FileID:       fileID,
		ArtifactType: "video",
		DownloadURL:  downloadURL,
		TTLAt:        0,
	})
}

// isMiniMaxOfficialChannel 判断渠道的视频上游是否为官方 api.minimaxi.com
// 依据：channel.CustomParameter.video.upstream == "official" 或 video.base_url 包含 api.minimaxi.com
// 若无配置或解析失败，按官方处理（官方优先）
func isMiniMaxOfficialChannelForFile(ch *model.Channel) bool {
	if ch == nil {
		return false
	}
	if ch.CustomParameter == nil || strings.TrimSpace(*ch.CustomParameter) == "" {
		return true
	}
	var payload map[string]json.RawMessage
	if err := json.Unmarshal([]byte(*ch.CustomParameter), &payload); err != nil {
		return true
	}
	videoRaw, ok := payload["video"]
	if !ok {
		return true
	}
	var videoCfg struct {
		Upstream string `json:"upstream"`
		BaseURL  string `json:"base_url"`
	}
	if err := json.Unmarshal(videoRaw, &videoCfg); err != nil {
		return true
	}
	up := strings.ToLower(strings.TrimSpace(videoCfg.Upstream))
	base := strings.ToLower(strings.TrimSpace(videoCfg.BaseURL))
	if up == "ppinfra" || strings.Contains(base, "api.ppinfra.com") {
		return false
	}
	if base != "" && !strings.Contains(base, "api.minimaxi.com") {
		return false
	}
	return true
}
