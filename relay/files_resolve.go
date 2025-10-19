package relay

import (
    "net/http"
    "one-api/common"
    "one-api/common/config"
    "one-api/model"
    "strings"

    "github.com/gin-gonic/gin"
)

// ResolveFileMapping
// 通过 file_id 解析：平台/任务/渠道/上游等信息，便于“无感”路由与调试
// 请求：GET /v1/files/resolve?file_id=xxx
// 响应示例：
// {
//   "file_id": "3249...",
//   "platform": "minimax",
//   "task_id": "task_xxx",
//   "channel_id": 12,
//   "channel_type": 27,
//   "upstream": "official|ppinfra|unknown",
//   "download_url": "..." // 如有
// }
func ResolveFileMapping(c *gin.Context) {
    fileID := strings.TrimSpace(c.Query("file_id"))
    if fileID == "" {
        common.AbortWithMessage(c, http.StatusBadRequest, "file_id is required")
        return
    }

    userID := c.GetInt("id")

    // 目前优先支持 MiniMax 视频产物的映射解析（已在任务与 artifact 表中落库）
    platform := ""
    var channelID int
    taskID := ""
    downloadURL := ""

    if a, _ := model.GetArtifactByFileID(userID, model.TaskPlatformMiniMax, fileID); a != nil {
        platform = model.TaskPlatformMiniMax
        channelID = a.ChannelId
        taskID = a.TaskID
        downloadURL = strings.TrimSpace(a.DownloadURL)
    } else {
        if t, _ := model.GetMiniMaxTaskByFileID(userID, fileID); t != nil {
            platform = model.TaskPlatformMiniMax
            channelID = t.ChannelId
            taskID = t.TaskID
        }
    }

    if platform == "" || channelID <= 0 {
        common.AbortWithMessage(c, http.StatusNotFound, "未找到包含该文件的任务（可在提交或查询任务后重试）")
        return
    }

    ch, err := model.GetChannelById(channelID)
    if err != nil || ch == nil {
        common.AbortWithMessage(c, http.StatusServiceUnavailable, "channel not found")
        return
    }

    upstream := "unknown"
    switch platform {
    case model.TaskPlatformMiniMax:
        // 复用与文件下载一致的上游判定逻辑（relay.parseMiniMaxUpstream）
        upstream = parseMiniMaxUpstream(ch)
        if upstream == "" {
            upstream = "official"
        }
    default:
        // 其它平台后续按需扩展
        upstream = "unknown"
    }

    // 构造响应（字段尽量通用，便于前端直接展示/调试）
    payload := gin.H{
        "file_id":      fileID,
        "platform":     platform,
        "task_id":      taskID,
        "channel_id":   ch.Id,
        "channel_type": ch.Type,
        "channel_name": ch.Name,
        "group":        ch.Group,
        "upstream":     upstream,
        "base_url":     ch.GetBaseURL(),
    }
    if downloadURL != "" {
        payload["download_url"] = downloadURL
    }

    // 附加：若是 OpenAI/Azure 渠道，则标注更清晰的别名
    switch ch.Type {
    case config.ChannelTypeOpenAI:
        payload["provider"] = "openai"
    case config.ChannelTypeAzure:
        payload["provider"] = "azureopenai"
    case config.ChannelTypeMiniMax:
        payload["provider"] = "minimax"
    }

    c.JSON(http.StatusOK, payload)
}
