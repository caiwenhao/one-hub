package minimax

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

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

	// 可选：用户上下文校验（保持与其他接口一致）。因该接口仅转发，可放宽为无需任务表。

	// 1) 若传入 channel_id，则直接使用该渠道
	if chStr := strings.TrimSpace(c.Query("channel_id")); chStr != "" {
		if chID, err := strconv.Atoi(chStr); err == nil && chID > 0 {
			if channel, err := model.GetChannelById(chID); err == nil && channel != nil {
				provider := providers.GetProvider(channel, c)
				if mini, ok := provider.(*miniProvider.MiniMaxProvider); ok && mini.GetVideoClient() != nil {
					if resp, errWithCode := mini.GetVideoClient().RetrieveFile(fileID); errWithCode == nil {
						// 命中后补写 artifact，后续可直接定位渠道
						_ = model.UpsertTaskArtifact(model.DB, &model.TaskArtifact{
							Platform:     model.TaskPlatformMiniMax,
							UserId:       c.GetInt("id"),
							ChannelId:    channel.Id,
							TaskID:       "",
							FileID:       fileID,
							ArtifactType: "video",
							DownloadURL:  resp.File.DownloadURL,
							TTLAt:        0,
						})
						writeJSONNoEscape(c, http.StatusOK, resp)
						return
					}
				}
			}
		}
	}

	// 2) 若传 task_id，则按任务反查渠道（增强命中率）
	taskID := strings.TrimSpace(c.Query("task_id"))
	var task *model.Task
	var err error
	if taskID != "" {
		task, err = model.GetTaskByTaskId(model.TaskPlatformMiniMax, c.GetInt("id"), taskID)
		if err != nil {
			StringError(c, http.StatusInternalServerError, "task_query_failed", err.Error())
			return
		}
	}

	if task == nil {
		// 2.1 优先从 artifact 表定位渠道
		if a, err := model.GetArtifactByFileID(c.GetInt("id"), model.TaskPlatformMiniMax, fileID); err == nil && a != nil && a.ChannelId > 0 {
			if channel, err := model.GetChannelById(a.ChannelId); err == nil && channel != nil {
				provider := providers.GetProvider(channel, c)
				if mini, ok := provider.(*miniProvider.MiniMaxProvider); ok && mini.GetVideoClient() != nil {
					if resp, errWithCode := mini.GetVideoClient().RetrieveFile(fileID); errWithCode == nil {
						writeJSONNoEscape(c, http.StatusOK, resp)
						return
					}
				}
			}
		}
		// 2.2 否则再用老的 data LIKE 反查（兼容历史数据）
		task, err = model.GetMiniMaxTaskByFileID(c.GetInt("id"), fileID)
		if err != nil {
			StringError(c, http.StatusInternalServerError, "task_query_failed", err.Error())
			return
		}
	}

	if task == nil {
		// 回退方案 A：在当前分组下遍历所有可用的 MiniMax 渠道逐一直查（官方优先：api.minimaxi.com > 其他），命中即返回
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
				provider := providers.GetProvider(ch, c)
				if mini, ok := provider.(*miniProvider.MiniMaxProvider); ok && mini.GetVideoClient() != nil {
					if resp, errWithCode := mini.GetVideoClient().RetrieveFile(fileID); errWithCode == nil {
						_ = model.UpsertTaskArtifact(model.DB, &model.TaskArtifact{
							Platform:     model.TaskPlatformMiniMax,
							UserId:       c.GetInt("id"),
							ChannelId:    ch.Id,
							TaskID:       "",
							FileID:       fileID,
							ArtifactType: "video",
							DownloadURL:  resp.File.DownloadURL,
							TTLAt:        0,
						})
						writeJSONNoEscape(c, http.StatusOK, resp)
						return
					}
				}
			}
			for _, ch := range others {
				provider := providers.GetProvider(ch, c)
				if mini, ok := provider.(*miniProvider.MiniMaxProvider); ok && mini.GetVideoClient() != nil {
					if resp, errWithCode := mini.GetVideoClient().RetrieveFile(fileID); errWithCode == nil {
						_ = model.UpsertTaskArtifact(model.DB, &model.TaskArtifact{
							Platform:     model.TaskPlatformMiniMax,
							UserId:       c.GetInt("id"),
							ChannelId:    ch.Id,
							TaskID:       "",
							FileID:       fileID,
							ArtifactType: "video",
							DownloadURL:  resp.File.DownloadURL,
							TTLAt:        0,
						})
						writeJSONNoEscape(c, http.StatusOK, resp)
						return
					}
				}
			}
		}

		// 回退方案 B：限定 MiniMax 后随机选择一个渠道直查（作为兜底）
		c.Set("allow_channel_type", []int{config.ChannelTypeMiniMax})
		if provider, _, fail := relay.GetProvider(c, "MiniMax-M1"); fail == nil && provider != nil {
			if mini, ok := provider.(*miniProvider.MiniMaxProvider); ok && mini.GetVideoClient() != nil {
				if resp, errWithCode := mini.GetVideoClient().RetrieveFile(fileID); errWithCode == nil {
					_ = model.UpsertTaskArtifact(model.DB, &model.TaskArtifact{
						Platform:     model.TaskPlatformMiniMax,
						UserId:       c.GetInt("id"),
						ChannelId:    provider.GetChannel().Id,
						TaskID:       "",
						FileID:       fileID,
						ArtifactType: "video",
						DownloadURL:  resp.File.DownloadURL,
						TTLAt:        0,
					})
					writeJSONNoEscape(c, http.StatusOK, resp)
					return
				}
			}
		}

		StringError(c, http.StatusNotFound, "task_not_found", "未找到包含该文件的任务（可加 channel_id 或 task_id 重试）")
		return
	}

	channel, channelErr := model.GetChannelById(task.ChannelId)
	if channelErr != nil {
		StringError(c, http.StatusServiceUnavailable, "channel_not_found", channelErr.Error())
		return
	}

	provider := providers.GetProvider(channel, c)
	mini, ok := provider.(*miniProvider.MiniMaxProvider)
	if !ok || mini.GetVideoClient() == nil {
		StringError(c, http.StatusServiceUnavailable, "provider_not_found", "provider not found")
		return
	}

	resp, errWithCode := mini.GetVideoClient().RetrieveFile(fileID)
	if errWithCode != nil {
		StringError(c, http.StatusInternalServerError, "file_retrieve_failed", errWithCode.Message)
		return
	}

	writeJSONNoEscape(c, http.StatusOK, resp)
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
