package minimaxi

import (
    "bytes"
    "encoding/hex"
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "net/http"
    "net/url"
    "one-api/common"
    "one-api/common/config"
    "one-api/common/logger"
    "one-api/common/utils"
    "one-api/types"
    "strconv"
    "strings"
)

const (
	audioModeJSON = "json"
	audioModeHex  = "hex"
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

func (p *MiniMaxProvider) getAudioMode() string {
	customParams, err := p.CustomParameterHandler()
	if err != nil || customParams == nil {
		return audioModeJSON
	}

	candidates := []interface{}{}
	if v, ok := customParams["audio_mode"]; ok {
		candidates = append(candidates, v)
	}
	if audioRaw, ok := customParams["audio"]; ok {
		if audioMap, ok := audioRaw.(map[string]interface{}); ok {
			if v, ok := audioMap["mode"]; ok {
				candidates = append(candidates, v)
			}
		}
	}

	for _, candidate := range candidates {
		if mode := normalizeAudioMode(candidate); mode != "" {
			return mode
		}
	}

	return audioModeJSON
}

func normalizeAudioMode(value interface{}) string {
	var mode string
	switch v := value.(type) {
	case string:
		mode = v
	case *string:
		if v == nil {
			return ""
		}
		mode = *v
	default:
		mode = fmt.Sprintf("%v", v)
	}
	mode = strings.ToLower(strings.TrimSpace(mode))
	switch mode {
	case audioModeHex:
		return audioModeHex
	case audioModeJSON:
		return audioModeJSON
	default:
		return ""
	}
}

func (p *MiniMaxProvider) logAudioError(message string) {
	if p.Context != nil && p.Context.Request != nil {
		logger.LogError(p.Context.Request.Context(), message)
		return
	}
	logger.SysError(message)
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

// isPPInfraSpeechUpstream 检测语音是否使用 PPInfra 上游（来自渠道 custom_parameter）
func (p *MiniMaxProvider) isPPInfraSpeechUpstream() bool {
    if p.Channel == nil || p.Channel.CustomParameter == nil || *p.Channel.CustomParameter == "" {
        return false
    }
    var payload map[string]any
    if err := json.Unmarshal([]byte(*p.Channel.CustomParameter), &payload); err != nil {
        return false
    }
    // audio 块优先
    if v, ok := payload["audio"]; ok {
        if audioMap, ok2 := v.(map[string]any); ok2 {
            if up, ok3 := audioMap["upstream"].(string); ok3 {
                return strings.ToLower(strings.TrimSpace(up)) == "ppinfra"
            }
        }
    }
    if up, ok := payload["upstream"].(string); ok {
        return strings.ToLower(strings.TrimSpace(up)) == "ppinfra"
    }
    return false
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

    // PPInfra 上游：直接透传（JSON 或 SSE），避免按官方结构解码失败
    if p.isPPInfraSpeechUpstream() {
        resp, errWith := p.Requester.SendRequestRaw(req)
        if errWith != nil {
            return nil, errWith
        }
        return resp, nil
    }

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

	contentType := "application/octet-stream"
	if speechResponse.ExtraInfo.AudioFormat != nil {
		format := strings.TrimSpace(*speechResponse.ExtraInfo.AudioFormat)
		if format != "" {
			contentType = "audio/" + format
		}
	}
	response.Header.Set("Content-Type", contentType)

	if speechResponse.ExtraInfo.AudioSize != nil && *speechResponse.ExtraInfo.AudioSize > 0 {
		response.Header.Set("Content-Length", strconv.FormatInt(*speechResponse.ExtraInfo.AudioSize, 10))
	}

	if speechResponse.Data.SubtitleFile != nil {
		subtitle := strings.TrimSpace(*speechResponse.Data.SubtitleFile)
		if subtitle != "" {
			response.Header.Set("X-Minimax-Subtitle-URL", subtitle)
		}
	}

	if speechResponse.ExtraInfo.AudioChannel != nil {
		response.Header.Set("X-Minimax-Audio-Channel", strconv.FormatInt(*speechResponse.ExtraInfo.AudioChannel, 10))
	}

	usage := 0
	if speechResponse.ExtraInfo.UsageCharacters != nil {
		usage = *speechResponse.ExtraInfo.UsageCharacters
	}
	p.Usage.PromptTokens = usage
	p.Usage.TotalTokens = usage

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

	usage := 0
	if asyncResp.UsageCharacters != nil {
		usage = *asyncResp.UsageCharacters
	}
	p.Usage.PromptTokens = usage
	p.Usage.TotalTokens = usage

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

    // 非流式 + PPInfra 上游：直接透传 JSON（可能为 url 或 hex），不按官方结构解码
    if p.isPPInfraSpeechUpstream() {
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

	usage := 0
	if speechResponse.ExtraInfo.UsageCharacters != nil {
		usage = *speechResponse.ExtraInfo.UsageCharacters
	}
	p.Usage.PromptTokens = usage
	p.Usage.TotalTokens = usage

    if p.getAudioMode() != audioModeHex {
        return resp, nil
    }

	audioHex := strings.TrimSpace(speechResponse.Data.Audio)
	if audioHex == "" {
		err := errors.New("empty audio data in MiniMax response")
		p.logAudioError(err.Error())
		return nil, common.ErrorWrapper(err, "speech_decode_error", http.StatusBadGateway)
	}

	audioBytes, err := hex.DecodeString(audioHex)
	if err != nil {
		errMsg := fmt.Sprintf("decode MiniMax audio hex failed: %v", err)
		p.logAudioError(errMsg)
		return nil, common.ErrorWrapper(err, "speech_decode_error", http.StatusBadGateway)
	}

	if resp.Body != nil {
		_ = resp.Body.Close()
	}

	body := io.NopCloser(bytes.NewReader(audioBytes))
	result := &http.Response{
		Status:     "200 OK",
		StatusCode: http.StatusOK,
		Body:       body,
		Header:     make(http.Header),
	}

	contentType := "application/octet-stream"
	if speechResponse.ExtraInfo.AudioFormat != nil {
		format := strings.TrimSpace(*speechResponse.ExtraInfo.AudioFormat)
		if format != "" {
			contentType = "audio/" + format
		}
	}
	result.Header.Set("Content-Type", contentType)

	if speechResponse.ExtraInfo.AudioSize != nil && *speechResponse.ExtraInfo.AudioSize > 0 {
		result.Header.Set("Content-Length", strconv.FormatInt(*speechResponse.ExtraInfo.AudioSize, 10))
	}

	if speechResponse.ExtraInfo.AudioChannel != nil {
		result.Header.Set("X-Minimax-Audio-Channel", strconv.FormatInt(*speechResponse.ExtraInfo.AudioChannel, 10))
	}

	if speechResponse.Data.SubtitleFile != nil {
		subtitle := strings.TrimSpace(*speechResponse.Data.SubtitleFile)
		if subtitle != "" {
			result.Header.Set("X-Minimax-Subtitle-URL", subtitle)
		}
	}

	return result, nil
}

func (p *MiniMaxProvider) QuerySpeechAsync(taskID string) (*http.Response, *types.OpenAIErrorWithStatusCode) {
	if strings.TrimSpace(taskID) == "" {
		return nil, common.StringErrorWrapperLocal("task_id is required", "invalid_request", http.StatusBadRequest)
	}

	path, errWithCode := p.GetSupportedAPIUri(config.RelayModeMiniMaxSpeechAsyncQuery)
	if errWithCode != nil {
		return nil, errWithCode
	}

	fullURL := p.GetFullRequestURL(path, "")
	parsed, err := url.Parse(fullURL)
	if err != nil {
		return nil, common.ErrorWrapper(err, "invalid_request_url", http.StatusInternalServerError)
	}
	query := parsed.Query()
	query.Set("task_id", taskID)
	parsed.RawQuery = query.Encode()

	if p.Requester == nil {
		return nil, common.StringErrorWrapperLocal("minimaxi requester not initialized", "channel_error", http.StatusServiceUnavailable)
	}

	httpReq, err := p.Requester.NewRequest(http.MethodGet, parsed.String(), p.Requester.WithHeader(p.GetRequestHeaders()))
	if err != nil {
		return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError)
	}
	if httpReq.Body != nil {
		defer httpReq.Body.Close()
	}

	queryResp := &types.MiniMaxAsyncSpeechQueryResponse{}
	resp, errWithCode := p.Requester.SendRequest(httpReq, queryResp, true)
	if errWithCode != nil {
		return nil, errWithCode
	}

	if queryResp.BaseResp.StatusCode != 0 {
		return nil, common.ErrorWrapper(errors.New(queryResp.BaseResp.StatusMsg), "speech_async_query_error", http.StatusBadGateway)
	}

	return resp, nil
}
