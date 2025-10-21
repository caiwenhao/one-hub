package relay

import (
	"net/http"
	"one-api/common"
	"one-api/common/config"
	"one-api/model"
	"one-api/providers"
	miniProvider "one-api/providers/minimaxi"
	"one-api/types"
	"strings"

	"github.com/gin-gonic/gin"
)

// MiniMaxAsyncQuery 查询 MiniMax 同步语音异步任务状态，仅作用于 /minimaxi 官方路由
func MiniMaxAsyncQuery(c *gin.Context) {
	taskID := strings.TrimSpace(c.Query("task_id"))
	if taskID == "" {
		err := common.StringErrorWrapperLocal("task_id is required", "invalid_request", http.StatusBadRequest)
		relayResponseWithOpenAIErr(c, err)
		return
	}

	// 该查询接口不需要 channel_id，默认按分组择优遍历可用的 MiniMax 渠道
	var channels []*model.Channel
	allChannels, dbErr := model.GetAllChannels()
	if dbErr != nil {
		errWithCode := common.ErrorWrapper(dbErr, "fetch_channels_failed", http.StatusInternalServerError)
		relayResponseWithOpenAIErr(c, errWithCode)
		return
	}
	group := strings.TrimSpace(c.GetString("token_group"))
	for _, ch := range allChannels {
		if ch == nil || ch.Status != config.ChannelStatusEnabled || ch.Type != config.ChannelTypeMiniMax {
			continue
		}
		if group != "" && !channelInGroup(ch, group) {
			continue
		}
		channels = append(channels, ch)
	}

	if len(channels) == 0 {
		err := common.StringErrorWrapperLocal("no available MiniMax channel", "channel_not_found", http.StatusServiceUnavailable)
		relayResponseWithOpenAIErr(c, err)
		return
	}

	var lastErr *types.OpenAIErrorWithStatusCode
	for _, channel := range channels {
		provider := providers.GetProvider(channel, c)
		if provider == nil {
			continue
		}
		mini, ok := provider.(*miniProvider.MiniMaxProvider)
		if !ok {
			continue
		}
		mini.SetContext(c)
		resp, errWithCode := mini.QuerySpeechAsync(taskID)
		if errWithCode != nil {
			lastErr = errWithCode
			continue
		}
		if resp == nil {
			continue
		}
		if err := responseMultipart(c, resp); err != nil {
			relayResponseWithOpenAIErr(c, err)
			return
		}
		return
	}

	if lastErr != nil {
		relayResponseWithOpenAIErr(c, lastErr)
		return
	}

	err := common.StringErrorWrapperLocal("query failed", "upstream_error", http.StatusBadGateway)
	relayResponseWithOpenAIErr(c, err)
}

func channelInGroup(channel *model.Channel, group string) bool {
	if channel == nil {
		return false
	}
	if strings.TrimSpace(channel.Group) == "" {
		return true
	}
	for _, item := range strings.Split(channel.Group, ",") {
		if strings.TrimSpace(item) == group {
			return true
		}
	}
	return false
}
