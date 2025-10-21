package relay

import (
	"strings"

	"one-api/common/config"

	"github.com/gin-gonic/gin"
)

// MiniMaxRelay 将 allow_channel_type 限定为 MiniMax，再复用通用 Relay
func MiniMaxRelay(c *gin.Context) {
	c.Set("allow_channel_type", []int{config.ChannelTypeMiniMax})

	originalPath := c.Request.URL.Path
	if strings.HasPrefix(originalPath, "/minimaxi") {
		c.Request.URL.Path = strings.TrimPrefix(originalPath, "/minimaxi")
		if c.Request.URL.Path == "" {
			c.Request.URL.Path = "/"
		}
	}
	defer func() {
		c.Request.URL.Path = originalPath
	}()

	Relay(c)
}
