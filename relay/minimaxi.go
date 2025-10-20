package relay

import (
    "one-api/common/config"
    "github.com/gin-gonic/gin"
)

// MiniMaxRelay 将 allow_channel_type 限定为 MiniMax，再复用通用 Relay
func MiniMaxRelay(c *gin.Context) {
    c.Set("allow_channel_type", []int{config.ChannelTypeMiniMax})
    Relay(c)
}

