package relay

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"one-api/common"
	"one-api/common/config"
	providersBase "one-api/providers/base"
	miniProvider "one-api/providers/minimaxi"
	"one-api/types"

	"github.com/gin-gonic/gin"
)

type relaySpeech struct {
	relayBase
	request           types.SpeechAudioRequest
	minimaxRawBody    []byte
	minimaxRequest    *types.MiniMaxSpeechRequest
	isMiniMaxOfficial bool
}

func NewRelaySpeech(c *gin.Context) *relaySpeech {
	relay := &relaySpeech{}
	relay.c = c
	return relay
}

func (r *relaySpeech) setRequest() error {
	if err := common.UnmarshalBodyReusable(r.c, &r.request); err != nil {
		return err
	}

	if raw, exists := r.c.Get(config.GinRequestBodyKey); exists {
		if bodyBytes, ok := raw.([]byte); ok {
			var mmReq types.MiniMaxSpeechRequest
			if err := json.Unmarshal(bodyBytes, &mmReq); err == nil && mmReq.Model != "" && mmReq.Text != "" && r.request.Input == "" {
				r.minimaxRawBody = bodyBytes
				r.minimaxRequest = &mmReq
				r.isMiniMaxOfficial = true
				r.request.Model = mmReq.Model
				r.request.Input = mmReq.Text
				if mmReq.VoiceSetting != nil && mmReq.VoiceSetting.VoiceID != "" {
					r.request.Voice = mmReq.VoiceSetting.VoiceID
				}
				if mmReq.VoiceSetting != nil && mmReq.VoiceSetting.Speed != nil {
					r.request.Speed = *mmReq.VoiceSetting.Speed
				}
			}
		}
	}

	if strings.TrimSpace(r.request.Model) == "" {
		return fmt.Errorf("model is required")
	}

	if !r.isMiniMaxOfficial {
		if strings.TrimSpace(r.request.Input) == "" {
			return fmt.Errorf("input is required")
		}
		if strings.TrimSpace(r.request.Voice) == "" {
			return fmt.Errorf("voice is required")
		}
	} else if strings.TrimSpace(r.request.Input) == "" {
		return fmt.Errorf("text is required")
	}

	r.setOriginalModel(r.request.Model)

	return nil
}

func (r *relaySpeech) getPromptTokens() (int, error) {
	return len(r.request.Input), nil
}

func (r *relaySpeech) send() (err *types.OpenAIErrorWithStatusCode, done bool) {
	provider, ok := r.provider.(providersBase.SpeechInterface)
	if !ok {
		err = common.StringErrorWrapperLocal("channel not implemented", "channel_error", http.StatusServiceUnavailable)
		done = true
		return
	}

	r.request.Model = r.modelName

	if r.isMiniMaxOfficial {
		mini, ok := provider.(*miniProvider.MiniMaxProvider)
		if !ok {
			err = common.StringErrorWrapperLocal("invalid minimaxi provider", "channel_error", http.StatusServiceUnavailable)
			done = true
			return
		}
		var reqCopy *types.MiniMaxSpeechRequest
		payload := r.minimaxRawBody
		if r.minimaxRequest != nil {
			copyReq := *r.minimaxRequest
			copyReq.Model = r.modelName
			reqCopy = &copyReq
			payload = nil
		}
		response, errWithCode := mini.CreateSpeechOfficial(reqCopy, payload)
		if errWithCode != nil {
			return errWithCode, true
		}
		err = responseMultipart(r.c, response)
		if err != nil {
			done = true
		}
		return nil, done
	}

	response, err := provider.CreateSpeech(&r.request)
	if err != nil {
		return
	}
	err = responseMultipart(r.c, response)

	if err != nil {
		done = true
	}

	return
}
