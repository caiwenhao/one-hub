package relay

import (
	"io"
	"net/http"
	"strings"

	"one-api/common"
	"one-api/common/config"
	"one-api/common/logger"
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

	// 直透上游 /v1/videos/generations
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

	for key, values := range resp.Header {
		for _, v := range values {
			c.Writer.Header().Add(key, v)
		}
	}
	c.Status(resp.StatusCode)
	if _, copyErr := io.Copy(c.Writer, resp.Body); copyErr != nil {
		logger.LogError(c.Request.Context(), "copy_video_generations_failed:"+copyErr.Error())
	}

	// 成功则按次消耗
	quota.Consume(c, &types.Usage{PromptTokens: 0, TotalTokens: 0}, false)
}
