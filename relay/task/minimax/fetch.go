package minimax

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"one-api/common/config"
	"one-api/model"
	"one-api/providers"
	miniProvider "one-api/providers/minimaxi"
	"one-api/relay"

	"github.com/gin-gonic/gin"
)

// RelayTaskFetch：DB优先精准路由；不要求 channel_id，与官方参数保持一致
// 流程：
// 1) 仅用 task_id（与官方一致）
// 2) 先查 DB: (platform=minimax,user_id,task_id) → channel_id 命中直转发
// 3) 回填：若响应含 file_id/直链，入库 artifact，后续文件下载可 O(1) 命中
// 4) 兜底（可选）：若 DB 未命中，可按“组内所有 MiniMax 渠道”穷举上游查询，命中则补写 task 映射/或至少写 artifact
func RelayTaskFetch(c *gin.Context) {
	taskID := c.Param("task_id")
	if taskID == "" {
		taskID = c.Query("task_id")
	}
	if taskID == "" {
		StringError(c, http.StatusBadRequest, "invalid_request", "task_id is required")
		return
	}

	userID := c.GetInt("id")

	// DB-First: 先按 (platform,user,task_id) 查任务映射
	task := model.GetByUserAndTaskId(userID, taskID)
	if task != nil && task.Platform == model.TaskPlatformMiniMax && task.ChannelId > 0 {
		if respOK := fetchViaChannel(c, task.ChannelId, taskID, extractModelFromTask(task)); respOK {
			return
		}
	}

	// 兜底：组内 MiniMax 渠道逐一尝试（默认官方优先：api.minimaxi.com > 其他上游）
	// 若你希望严格模式，可将该段受配置控制。
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
			if isMiniMaxOfficialChannel(ch) {
				official = append(official, ch)
			} else {
				others = append(others, ch)
			}
		}
		for _, ch := range official {
			if respOK := fetchViaChannel(c, ch.Id, taskID, ""); respOK {
				return
			}
		}
		for _, ch := range others {
			if respOK := fetchViaChannel(c, ch.Id, taskID, ""); respOK {
				return
			}
		}
	}

	// 最后兜底一次：限定 MiniMax 后随机选择一个渠道
	c.Set("allow_channel_type", []int{config.ChannelTypeMiniMax})
	if p, _, fail := relay.GetProvider(c, "MiniMax-M1"); fail == nil {
		if mini, ok := p.(*miniProvider.MiniMaxProvider); ok && mini.GetVideoClient() != nil {
			if resp, errWithCode := mini.GetVideoClient().QueryVideoTask(taskID, ""); errWithCode == nil {
				// 命中后尝试回填 artifact（若有）
				_ = upsertMiniMaxArtifacts(&model.Task{UserId: userID, ChannelId: p.GetChannel().Id, TaskID: taskID}, resp)
				writeJSONNoEscape(c, http.StatusOK, resp)
				return
			}
		}
	}

	StringError(c, http.StatusNotFound, "upstream_not_found_or_no_permission", "该任务可能不是由本平台创建或不属于当前分组/账号")
}

// fetchViaChannel 使用指定渠道查询上游，并在成功时回填 artifact
func fetchViaChannel(c *gin.Context, channelID int, taskID string, modelName string) bool {
	channel, err := model.GetChannelById(channelID)
	if err != nil || channel == nil {
		return false
	}
	provider := providers.GetProvider(channel, c)
	mini, ok := provider.(*miniProvider.MiniMaxProvider)
	if !ok || mini.GetVideoClient() == nil {
		return false
	}
	resp, errWithCode := mini.GetVideoClient().QueryVideoTask(taskID, modelName)
	if errWithCode != nil {
		return false
	}
	// 回填 artifact（若包含 file_id/直链）
	_ = upsertMiniMaxArtifacts(&model.Task{UserId: c.GetInt("id"), ChannelId: channelID, TaskID: taskID, FinishTime: time.Now().Unix()}, resp)
	writeJSONNoEscape(c, http.StatusOK, resp)
	return true
}

// isMiniMaxOfficialChannel 判断渠道的视频上游是否为官方 api.minimaxi.com
// 依据：channel.CustomParameter.video.upstream == "official" 或 video.base_url 包含 api.minimaxi.com
// 若无配置或解析失败，按官方处理（官方优先）
func isMiniMaxOfficialChannel(ch *model.Channel) bool {
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
