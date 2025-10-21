package minimaxi

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"one-api/common"
	"one-api/common/config"
	"one-api/common/utils"
	"one-api/types"
	"strconv"
	"strings"
)

func (p *MiniMaxProvider) GetVoiceMap() map[string][]string {
	defaultVoiceMapping := map[string][]string{
		"alloy":   {"female-chengshu"},
		"echo":    {"male-qn-qingse"},
		"fable":   {"male-qn-jingying"},
		"onyx":    {"presenter_male"},
		"nova":    {"presenter_female"},
		"shimmer": {"audiobook_female_1"},
	}

	if p.Channel.Plugin == nil {
		return defaultVoiceMapping
	}

	customVoiceMapping, ok := p.Channel.Plugin.Data()["voice"]
	if !ok {
		return defaultVoiceMapping
	}

	for key, value := range customVoiceMapping {
		if _, exists := defaultVoiceMapping[key]; !exists {
			continue
		}
		customVoiceValue, isString := value.(string)
		if !isString || customVoiceValue == "" {
			continue
		}
		customizeVoice := strings.Split(customVoiceValue, "|")
		defaultVoiceMapping[key] = customizeVoice
	}

	return defaultVoiceMapping
}

func (p *MiniMaxProvider) resolveVoiceAlias(alias string) (string, *string, bool) {
	voiceMap := p.GetVoiceMap()
	if alias == "" {
		return "", nil, false
	}
	mapping, ok := voiceMap[alias]
	if !ok || len(mapping) == 0 {
		return alias, nil, false
	}
	var emotionPtr *string
	if len(mapping) > 1 {
		emotion := mapping[1]
		emotionPtr = &emotion
	}
	return mapping[0], emotionPtr, true
}

func (p *MiniMaxProvider) patchAsyncPayloadVoice(raw []byte) ([]byte, error) {
	if len(raw) == 0 {
		return raw, nil
	}
	var payload map[string]interface{}
	if err := json.Unmarshal(raw, &payload); err != nil {
		return nil, err
	}
	voiceSettingAny, ok := payload["voice_setting"]
	if !ok {
		return raw, nil
	}
	voiceSetting, ok := voiceSettingAny.(map[string]interface{})
	if !ok {
		return raw, nil
	}
	voiceIDAny, ok := voiceSetting["voice_id"]
	if !ok {
		return raw, nil
	}
	voiceID, ok := voiceIDAny.(string)
	if !ok || voiceID == "" {
		return raw, nil
	}

	resolved, emotion, changed := p.resolveVoiceAlias(voiceID)
	if !changed {
		return raw, nil
	}

	voiceSetting["voice_id"] = resolved
	if emotion != nil {
		if current, exists := voiceSetting["emotion"]; !exists || current == nil || strings.TrimSpace(anyToString(current)) == "" {
			voiceSetting["emotion"] = *emotion
		}
	}

	updated, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (p *MiniMaxProvider) getRequestBody(request *types.SpeechAudioRequest) *SpeechRequest {

	var voice, emotion string
	voiceMap := p.GetVoiceMap()
	if voiceMap[request.Voice] != nil {
		voice = voiceMap[request.Voice][0]
		if len(voiceMap[request.Voice]) > 1 {
			emotion = voiceMap[request.Voice][1]
		}
	} else {
		voice = request.Voice
	}

	speechRequest := &SpeechRequest{
		Model: request.Model,
		Text:  request.Input,
		VoiceSetting: VoiceSetting{
			VoiceID: voice,
			Emotion: emotion,
			Speed:   request.Speed,
		},
	}

	// mp3-1-32000-128000
	if request.ResponseFormat != "" {
		formats := strings.Split(request.ResponseFormat, "-")
		speechRequest.AudioSetting = &AudioSetting{
			Format: formats[0],
		}
		if len(formats) > 1 {
			speechRequest.AudioSetting.Channel = utils.String2Int64(formats[1])
		}
		if len(formats) > 2 {
			speechRequest.AudioSetting.SampleRate = utils.String2Int64(formats[2])
		}
		if len(formats) > 3 {
			speechRequest.AudioSetting.Bitrate = utils.String2Int64(formats[3])
		}
	}

	return speechRequest
}

func (p *MiniMaxProvider) CreateSpeech(request *types.SpeechAudioRequest) (*http.Response, *types.OpenAIErrorWithStatusCode) {
	url, errWithCode := p.GetSupportedAPIUri(config.RelayModeAudioSpeech)
	if errWithCode != nil {
		return nil, errWithCode
	}
	fullRequestURL := p.GetFullRequestURL(url, request.Model)
	headers := p.GetRequestHeaders()

	requestBody := p.getRequestBody(request)

	req, err := p.Requester.NewRequest(http.MethodPost, fullRequestURL, p.Requester.WithBody(requestBody), p.Requester.WithHeader(headers))
	if err != nil {
		return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError)
	}
	defer req.Body.Close()

	speechResponse := &SpeechResponse{}
	_, errWithCode = p.Requester.SendRequest(req, speechResponse, false)
	if errWithCode != nil {
		return nil, errWithCode
	}

	if speechResponse.BaseResp.StatusCode != 0 {
		return nil, common.ErrorWrapper(errors.New(speechResponse.BaseResp.StatusMsg), "speech_error", http.StatusInternalServerError)
	}

	audioBytes, err := hex.DecodeString(speechResponse.Data.Audio)
	if err != nil {
		return nil, common.ErrorWrapper(err, "decode_audio_data_failed", http.StatusInternalServerError)
	}

	body := io.NopCloser(bytes.NewReader(audioBytes))

	response := &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Body:       body,
		Header:     make(http.Header),
	}

	response.Header.Set("Content-Type", "audio/"+speechResponse.ExtraInfo.AudioFormat) // 例如 "audio/wav"
	response.Header.Set("Content-Length", strconv.FormatInt(speechResponse.ExtraInfo.AudioSize, 10))

	p.Usage.PromptTokens = speechResponse.ExtraInfo.UsageCharacters
	p.Usage.TotalTokens = speechResponse.ExtraInfo.UsageCharacters

	return response, nil
}

func (p *MiniMaxProvider) CreateSpeechAsync(req *types.MiniMaxAsyncSpeechRequest, rawBody []byte) (*http.Response, *types.OpenAIErrorWithStatusCode) {
	url, errWithCode := p.GetSupportedAPIUri(config.RelayModeMiniMaxSpeechAsync)
	if errWithCode != nil {
		return nil, errWithCode
	}

	modelName := p.GetOriginalModel()
	if req != nil && strings.TrimSpace(req.Model) != "" {
		modelName = req.Model
	}
	fullRequestURL := p.GetFullRequestURL(url, modelName)
	headers := p.GetRequestHeaders()

	payload := rawBody
	if len(payload) == 0 && req != nil {
		body, err := json.Marshal(req)
		if err != nil {
			return nil, common.ErrorWrapper(err, "marshal_request_failed", http.StatusInternalServerError)
		}
		payload = body
	}

	if patched, err := p.patchAsyncPayloadVoice(payload); err == nil && len(patched) > 0 {
		payload = patched
	} else if err != nil {
		return nil, common.ErrorWrapper(err, "marshal_request_failed", http.StatusInternalServerError)
	}

	httpReq, err := p.Requester.NewRequest(http.MethodPost, fullRequestURL, p.Requester.WithBody(payload), p.Requester.WithHeader(headers))
	if err != nil {
		return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError)
	}
	defer httpReq.Body.Close()

	asyncResp := &types.MiniMaxAsyncSpeechResponse{}
	resp, errWithCode := p.Requester.SendRequest(httpReq, asyncResp, true)
	if errWithCode != nil {
		return nil, errWithCode
	}

	if asyncResp.BaseResp.StatusCode != 0 {
		return nil, common.ErrorWrapper(errors.New(asyncResp.BaseResp.StatusMsg), "speech_async_error", http.StatusBadGateway)
	}

	p.Usage.PromptTokens = asyncResp.UsageCharacters
	p.Usage.TotalTokens = asyncResp.UsageCharacters

	return resp, nil
}

func anyToString(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case fmt.Stringer:
		return val.String()
	case json.Number:
		return val.String()
	default:
		if v == nil {
			return ""
		}
		return fmt.Sprintf("%v", v)
	}
}

func (p *MiniMaxProvider) normalizeOfficialSpeechRequest(req *types.MiniMaxSpeechRequest) ([]byte, error) {
	if req == nil {
		return nil, nil
	}
	clone := *req
	if req.VoiceSetting != nil {
		voiceSetting := *req.VoiceSetting
		if resolved, emotion, ok := p.resolveVoiceAlias(voiceSetting.VoiceID); ok {
			voiceSetting.VoiceID = resolved
			if emotion != nil && (voiceSetting.Emotion == nil || *voiceSetting.Emotion == "") {
				voiceSetting.Emotion = emotion
			}
		}
		clone.VoiceSetting = &voiceSetting
	}
	return json.Marshal(&clone)
}

func (p *MiniMaxProvider) CreateSpeechOfficial(req *types.MiniMaxSpeechRequest, rawBody []byte) (*http.Response, *types.OpenAIErrorWithStatusCode) {
	url, errWithCode := p.GetSupportedAPIUri(config.RelayModeAudioSpeech)
	if errWithCode != nil {
		return nil, errWithCode
	}

	modelName := p.GetOriginalModel()
	if req != nil && req.Model != "" {
		modelName = req.Model
	}

	fullRequestURL := p.GetFullRequestURL(url, modelName)
	headers := p.GetRequestHeaders()

	var payload []byte
	if req != nil {
		body, err := p.normalizeOfficialSpeechRequest(req)
		if err != nil {
			return nil, common.ErrorWrapper(err, "marshal_request_failed", http.StatusInternalServerError)
		}
		payload = body
	} else if len(rawBody) > 0 {
		payload = rawBody
	}

	if len(payload) == 0 {
		return nil, common.StringErrorWrapperLocal("empty minimaxi payload", "invalid_request", http.StatusBadRequest)
	}

	httpReq, err := p.Requester.NewRequest(http.MethodPost, fullRequestURL, p.Requester.WithBody(payload), p.Requester.WithHeader(headers))
	if err != nil {
		return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError)
	}
	defer httpReq.Body.Close()

	streaming := false
	if req != nil && req.Stream != nil && *req.Stream {
		streaming = true
	}

	if streaming {
		resp, errWithCode := p.Requester.SendRequestRaw(httpReq)
		if errWithCode != nil {
			return nil, errWithCode
		}
		return resp, nil
	}

	speechResponse := &SpeechResponse{}
	resp, errWithCode := p.Requester.SendRequest(httpReq, speechResponse, true)
	if errWithCode != nil {
		return nil, errWithCode
	}

	if speechResponse.BaseResp.StatusCode != 0 {
		return nil, common.ErrorWrapper(errors.New(speechResponse.BaseResp.StatusMsg), "speech_error", http.StatusBadGateway)
	}

	p.Usage.PromptTokens = speechResponse.ExtraInfo.UsageCharacters
	p.Usage.TotalTokens = speechResponse.ExtraInfo.UsageCharacters

	return resp, nil
}
