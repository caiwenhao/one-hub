package relay

import (
	"fmt"
	"strings"

	"one-api/common"
	"one-api/common/config"
	miniProvider "one-api/providers/minimaxi"
	"one-api/types"

	"github.com/gin-gonic/gin"
)

type relayMiniMaxAsync struct {
	relayBase
	request types.MiniMaxAsyncSpeechRequest
	rawBody []byte
}

func NewRelayMiniMaxAsync(c *gin.Context) *relayMiniMaxAsync {
	relay := &relayMiniMaxAsync{}
	relay.c = c
	return relay
}

func (r *relayMiniMaxAsync) setRequest() error {
	if err := common.UnmarshalBodyReusable(r.c, &r.request); err != nil {
		return err
	}

	if raw, exists := r.c.Get(config.GinRequestBodyKey); exists {
		if bytes, ok := raw.([]byte); ok {
			r.rawBody = bytes
		}
	}

	r.c.Set("allow_channel_type", []int{config.ChannelTypeMiniMax})

	if strings.TrimSpace(r.request.Model) == "" {
		return fmt.Errorf("model is required")
	}

	textProvided := r.request.Text != nil && strings.TrimSpace(*r.request.Text) != ""
	fileProvided := r.request.TextFileID != nil && len(*r.request.TextFileID) > 0
	if !textProvided && !fileProvided {
		return fmt.Errorf("either text or text_file_id is required")
	}

	if r.request.VoiceSetting == nil || strings.TrimSpace(r.request.VoiceSetting.VoiceID) == "" {
		return fmt.Errorf("voice_setting.voice_id is required")
	}

	r.setOriginalModel(r.request.Model)

	return nil
}

func (r *relayMiniMaxAsync) getPromptTokens() (int, error) {
	if r.request.Text != nil {
		return len(*r.request.Text), nil
	}
	return 0, nil
}

func (r *relayMiniMaxAsync) send() (err *types.OpenAIErrorWithStatusCode, done bool) {
	mini, ok := r.provider.(*miniProvider.MiniMaxProvider)
	if !ok {
		err = common.StringErrorWrapperLocal("invalid minimaxi provider", "channel_error", 503)
		done = true
		return
	}

	r.request.Model = r.modelName

	resp, errWithCode := mini.CreateSpeechAsync(&r.request, r.rawBody)
	if errWithCode != nil {
		return errWithCode, true
	}

	if err = responseMultipart(r.c, resp); err != nil {
		done = true
	}

	return nil, done
}

func (r *relayMiniMaxAsync) getRequest() any {
	return &r.request
}
