package relay

import (
    "io"
    "encoding/json"
    "net/http"
    "strings"

    "one-api/common"
    "one-api/common/config"
    // "one-api/common/logger"
    "one-api/common/utils"
    "one-api/model"
    "one-api/relay/relay_util"
    "one-api/types"

	"github.com/gin-gonic/gin"
)

// Apimart 风格的视频创建请求体
type videoGenerationsRequest struct {
	Model       string   `json:"model"`
	Prompt      string   `json:"prompt"`
	Duration    int      `json:"duration,omitempty"`
	AspectRatio string   `json:"aspect_ratio,omitempty"`
	ImageURLs   []string `json:"image_urls,omitempty"`
	Watermark   *bool    `json:"watermark,omitempty"`
}

// VideoGenerationsCreate 兼容 /v1/videos/generations（Apimart 风格）
// 仅限 ChannelTypeNewAPI 渠道，其他类型返回不支持
func VideoGenerationsCreate(c *gin.Context) {
	// 对非 NewAPI 渠道，作为官方 /v1/videos 的别名直接复用标准处理逻辑
	if c.GetInt("channel_type") != config.ChannelTypeNewAPI {
		VideoCreate(c)
		return
	}

	// 仅允许 newapi 渠道
	c.Set("allow_channel_type", []int{config.ChannelTypeNewAPI})

	var reqBody videoGenerationsRequest
	if err := common.UnmarshalBodyReusable(c, &reqBody); err != nil {
		common.AbortWithMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	modelName := strings.TrimSpace(reqBody.Model)
	if modelName == "" {
		common.AbortWithMessage(c, http.StatusBadRequest, "model is required")
		return
	}

	// 选择渠道（newapi）
	provider, _, err := GetProvider(c, modelName)
	if err != nil {
		common.AbortWithMessage(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	// 计费（按次）：基于 modelName 的 TimesPriceType
	quota := relay_util.NewQuota(c, modelName, 0)
	if errWithCode := quota.PreQuotaConsumption(); errWithCode != nil {
		newErr := FilterOpenAIErr(c, errWithCode)
		relayResponseWithOpenAIErr(c, &newErr)
		return
	}

    // 直透上游 /v1/videos/generations，并在成功时包装 task_id 为平台ID
	base := strings.TrimSuffix(provider.GetChannel().GetBaseURL(), "/")
	if base == "" {
		base = "https://api.openai.com" // 兜底；newapi 一般会配置第三方域名
	}
	fullURL := base + "/v1/videos/generations"

	headers := provider.GetRequestHeaders()
	if headers == nil {
		headers = map[string]string{}
	}
	headers["Content-Type"] = "application/json"

	httpReq, e := provider.GetRequester().NewRequest(http.MethodPost, fullURL,
		provider.GetRequester().WithBody(reqBody),
		provider.GetRequester().WithHeader(headers),
	)
	if e != nil {
		common.AbortWithMessage(c, http.StatusInternalServerError, "new_request_failed")
		return
	}

	// 不做 OpenAI 错误包装，保持上游 JSON 结构
    resp, errWithCode := provider.GetRequester().SendRequestRaw(httpReq)
    if errWithCode != nil {
		quota.Undo(c)
		// 仍以原始结构返回（尽量）
		status := errWithCode.StatusCode
		if status == 0 {
			status = http.StatusBadGateway
		}
		c.Status(status)
		c.Writer.Header().Set("Content-Type", "application/json")
		_, _ = c.Writer.Write([]byte(errWithCode.OpenAIError.Message))
		return
	}
	defer resp.Body.Close()

    // 尝试读取 JSON 并改写 data[].task_id
    body, _ := io.ReadAll(resp.Body)
    _ = resp.Body.Close()
    var vendor map[string]any
    if json.Unmarshal(body, &vendor) == nil {
        // 预期结构：{"code":0, "data":[{"task_id":"...",...}]}
        if arr, ok := vendor["data"].([]any); ok && len(arr) > 0 {
            first, _ := arr[0].(map[string]any)
            if first != nil {
                if upID, ok2 := first["task_id"].(string); ok2 && strings.TrimSpace(upID) != "" {
                    // 落库任务（按 Sora 平台对齐视频）
                    task := &model.Task{
                        PlatformTaskID: utils.NewPlatformULID(),
                        TaskID:         upID,
                        Platform:       model.TaskPlatformSora,
                        UserId:         c.GetInt("id"),
                        TokenID:        c.GetInt("token_id"),
                        ChannelId:      provider.GetChannel().Id,
                        Action:         "video.generate.vendor",
                        Status:         model.TaskStatusSubmitted,
                        SubmitTime:     c.GetTime("requestStartTime").Unix(),
                        CreatedAt:      c.GetTime("requestStartTime").Unix(),
                        UpdatedAt:      c.GetTime("requestStartTime").Unix(),
                    }
                    _ = task.Insert()
                    platformID := utils.AddTaskPrefix(task.PlatformTaskID)
                    first["task_id"] = platformID
                    vendor["data"] = arr
                    patched, _ := json.Marshal(vendor)
                    c.Data(resp.StatusCode, "application/json", patched)
                    quota.Consume(c, &types.Usage{PromptTokens: 0, TotalTokens: 0}, false)
                    return
                }
            }
        }
    }
    // 回退：原样转发
    for key, values := range resp.Header { for _, v := range values { c.Writer.Header().Add(key, v) } }
    c.Status(resp.StatusCode)
    c.Writer.Write(body)
    quota.Consume(c, &types.Usage{PromptTokens: 0, TotalTokens: 0}, false)
}
