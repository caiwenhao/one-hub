package controller

import (
	"net/http"
	"one-api/common"
	"one-api/model"
	"one-api/common/video"
	"regexp"
	"encoding/json"
    "strings"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

func GetAllTask(c *gin.Context) {
	var params model.TaskQueryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}

	tasks, err := model.GetAllTasks(&params)
	if err != nil {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}
    // 改写任务数据中的视频直链为代理 URL
    rewriteTasksVideoURLs(tasks)

    c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    tasks,
	})
}

func GetUserAllTask(c *gin.Context) {
	userId := c.GetInt("id")

	var params model.TaskQueryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}

    tasks, err := model.GetAllUserTasks(userId, &params)
	if err != nil {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}

	re := regexp.MustCompile(`"user_id":\s*"[^"]*",?\s*`)

    for _, task := range *tasks.Data {
		if task.Platform == model.TaskPlatformSuno && task.Action == "MUSIC" {
			data := task.Data.String()
			if data != "" {
				data = re.ReplaceAllString(data, "")
				task.Data = datatypes.JSON(data)
			}
		}
	}

    // 改写任务数据中的视频直链为代理 URL
    rewriteTasksVideoURLs(tasks)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    tasks,
	})
}

// rewriteTasksVideoURLs 遍历任务列表，将 Data JSON 中的可识别视频直链替换为 CF Worker 代理 URL
func rewriteTasksVideoURLs(tasks *model.DataResult[model.Task]) {
    if tasks == nil || tasks.Data == nil { return }
    for _, t := range *tasks.Data {
        if len(t.Data) == 0 { continue }
        var anyData any
        if err := json.Unmarshal(t.Data, &anyData); err != nil { continue }
        changed := rewriteVideoURLsRecursive(&anyData, t.TaskID)
        if changed {
            if b, err := json.Marshal(anyData); err == nil {
                t.Data = datatypes.JSON(b)
            }
        }
    }
}

// rewriteVideoURLsRecursive 递归改写 map/array 中的视频 URL，返回是否发生变更
func rewriteVideoURLsRecursive(node *any, idHint string) bool {
    changed := false
    switch v := (*node).(type) {
    case map[string]any:
        for k, val := range v {
            // 递归处理子结构
            if m, ok := val.(map[string]any); ok {
                sub := any(m)
                if rewriteVideoURLsRecursive(&sub, idHint) { changed = true; v[k] = sub }
                continue
            }
            if arr, ok := val.([]any); ok {
                sub := any(arr)
                if rewriteVideoURLsRecursive(&sub, idHint) { changed = true; v[k] = sub }
                continue
            }
            // 针对常见字段做改写
            if s, ok := val.(string); ok {
                if shouldRewriteField(k, s) {
                    proxy := video.ProxyVideoURL(s, idHint)
                    if proxy != s {
                        v[k] = proxy
                        changed = true
                    }
                }
            }
        }
    case []any:
        for i, elem := range v {
            if m, ok := elem.(map[string]any); ok {
                sub := any(m)
                if rewriteVideoURLsRecursive(&sub, idHint) { changed = true; v[i] = sub }
                continue
            }
            if arr, ok := elem.([]any); ok {
                sub := any(arr)
                if rewriteVideoURLsRecursive(&sub, idHint) { changed = true; v[i] = sub }
                continue
            }
            if s, ok := elem.(string); ok {
                if shouldRewriteField("", s) {
                    proxy := video.ProxyVideoURL(s, idHint)
                    if proxy != s { v[i] = proxy; changed = true }
                }
            }
        }
    }
    return changed
}

// shouldRewriteField 判断该字段是否应视为视频直链并改写
func shouldRewriteField(key, val string) bool {
    if val == "" { return false }
    // 仅处理 http/https
    if !(len(val) > 7 && (val[:7] == "http://" || (len(val) > 8 && val[:8] == "https://"))) {
        return false
    }
    // 优先匹配常见视频字段名
    lower := strings.ToLower(key)
    if lower == "video_url" || lower == "thumbnail_url" || lower == "download_url" || lower == "spritesheet_url" || lower == "watermarked_url" {
        return true
    }
    // 兼容 raw.videos[].url 等，或明显视频扩展
    if lower == "url" {
        uv := strings.ToLower(val)
        if strings.HasSuffix(uv, ".mp4") || strings.HasSuffix(uv, ".webm") || strings.HasSuffix(uv, ".mov") || strings.Contains(uv, ".m3u8") {
            return true
        }
    }
    return false
}
