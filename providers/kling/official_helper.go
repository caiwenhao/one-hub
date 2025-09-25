package kling

import (
	"fmt"
	"net/http"
	"strings"

	"one-api/model"
	"one-api/types"

	"github.com/gin-gonic/gin"
)

// getKlingProviderByModel 根据模型名称选择可用渠道并返回供应商实例
func getKlingProviderByModel(c *gin.Context, modelName string) (*KlingProvider, *model.Channel, bool) {
	channel, err := model.ChannelGroup.Next("default", modelName)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, OfficialResponse{
			Code:    503,
			Message: fmt.Sprintf("无可用渠道: %v", err),
		})
		return nil, nil, false
	}

	provider := KlingProviderFactory{}.Create(channel)
	klingProvider, ok := provider.(*KlingProvider)
	if !ok {
		c.JSON(http.StatusServiceUnavailable, OfficialResponse{
			Code:    503,
			Message: "Provider类型错误",
		})
		return nil, nil, false
	}

	return klingProvider, channel, true
}

// writeProviderError 将供应商调用错误统一返回
func writeProviderError(c *gin.Context, errWithCode *types.OpenAIErrorWithStatusCode) bool {
	if errWithCode == nil {
		return false
	}

	message := "调用可灵服务失败"
	if errWithCode.Message != "" {
		message = errWithCode.Message
	}

	c.JSON(errWithCode.StatusCode, OfficialResponse{
		Code:    errWithCode.StatusCode,
		Message: message,
	})
	return true
}

// writeOperationResponse 按官方格式返回通用响应
func writeOperationResponse[T any](c *gin.Context, resp *KlingResponse[T]) {
	if resp == nil {
		c.JSON(http.StatusInternalServerError, OfficialResponse{
			Code:    500,
			Message: "可灵服务返回为空",
		})
		return
	}

	body := gin.H{
		"code":       resp.Code,
		"message":    resp.Message,
		"request_id": resp.RequestID,
		"data":       resp.Data,
	}

	c.JSON(http.StatusOK, body)
}

// resolveActionsByPath 根据路由路径推导需要过滤的任务动作
func resolveActionsByPath(fullPath string) []string {
	switch {
	case strings.Contains(fullPath, "/videos/text2video"):
		return []string{"text2video"}
	case strings.Contains(fullPath, "/videos/image2video"):
		return []string{"image2video"}
	case strings.Contains(fullPath, "/videos/multi-image2video"):
		return []string{"multi-image2video"}
	case strings.HasSuffix(fullPath, "/videos/multi-elements") || strings.Contains(fullPath, "/videos/multi-elements/:"):
		return []string{"multi-elements"}
	case strings.Contains(fullPath, "/images/generations"):
		return []string{"image_generation"}
	case strings.Contains(fullPath, "/images/multi-image2image"):
		return []string{"multi-image2image"}
	default:
		return nil
	}
}
