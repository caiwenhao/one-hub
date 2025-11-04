package relay

import (
	"io"
	"net/http"
	"one-api/common"
	"one-api/common/config"
	"one-api/common/logger"
	"strings"

	"github.com/gin-gonic/gin"
)

// NewAPITaskRetrieve 直透上游 /v1/tasks/{id}（仅 NewAPI 渠道）。
// 不计费，仅查询任务状态。
func NewAPITaskRetrieve(c *gin.Context) {
	// 仅允许 NewAPI 渠道类型
	c.Set("allow_channel_type", []int{config.ChannelTypeNewAPI})

	taskID := strings.TrimSpace(c.Param("id"))
	if taskID == "" {
		common.AbortWithMessage(c, http.StatusBadRequest, "task id is required")
		return
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

	// 直透响应（原样 JSON）
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

	for key, values := range resp.Header {
		for _, v := range values {
			c.Writer.Header().Add(key, v)
		}
	}
	c.Status(resp.StatusCode)
	if _, copyErr := io.Copy(c.Writer, resp.Body); copyErr != nil {
		logger.LogError(c.Request.Context(), "copy_newapi_task_failed:"+copyErr.Error())
	}
}
