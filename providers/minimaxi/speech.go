package minimaxi

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"one-api/common"
	"one-api/common/config"
	"one-api/common/logger"
	"one-api/common/utils"
	"one-api/types"
	"strconv"
	"strings"
	"unicode"
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

// getPPInfraSpeechPath 根据模型与同步/异步选择返回 PPInfra 的语音合成路径
// 参考文档（2025-11 已验证）：
// - 同步 HD:   /v3/minimax-speech-2.6-hd
// - 异步 HD:   /v3/async/minimax-speech-2.6-hd
// - 同步 Turbo:/v3/minimax-speech-2.6-turbo
// - 异步 Turbo:/v3/async/minimax-speech-2.6-turbo
// - 兼容旧型：02/2.5 系列仍按各自路径（例如 /v3/minimax-speech-02-hd 等）
func (p *MiniMaxProvider) getPPInfraSpeechPath(model string, async bool) string {
	m := strings.ToLower(strings.TrimSpace(model))
	// 允许传入 minimax-speech-* 形式或官方标准名 speech-*
	// 默认回退到 02-hd（服务稳定，且所有入参兼容）
	name := "minimax-speech-02-hd"
	switch m {
	case "speech-2.6-hd", "minimax-speech-2.6-hd":
		name = "minimax-speech-2.6-hd"
	case "speech-2.6-turbo", "minimax-speech-2.6-turbo":
		name = "minimax-speech-2.6-turbo"
	case "speech-02-hd", "minimax-speech-02-hd":
		name = "minimax-speech-02-hd"
	case "speech-02-turbo", "minimax-speech-02-turbo":
		name = "minimax-speech-02-turbo"
	case "speech-2.5-hd-preview", "minimax-speech-2.5-hd-preview":
		name = "minimax-speech-2.5-hd-preview"
	case "speech-2.5-turbo-preview", "minimax-speech-2.5-turbo-preview":
		name = "minimax-speech-2.5-turbo-preview"
	}
	if async {
		return "/v3/async/" + name
	}
	return "/v3/" + name
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

// estimateTextCharacters 估算文本字符数（用于 PPInfra 上游缺少 usage 的补偿）
func estimateTextCharacters(text string) int {
	t := strings.TrimSpace(text)
	if t == "" {
		return 0
	}
	return len([]rune(t))
}

// getAudioUsageRatio 读取渠道自定义参数中的估算系数；若未配置，则按内容语种做简易推断
// - 优先读取 audio.usage_ratio 或 usage_ratio（number）
// - 含中文（Han）字符时，默认 1.6；否则默认 1.0
func (p *MiniMaxProvider) getAudioUsageRatio(text string) float64 {
	// 默认值：含中文提升系数（对齐官方：约 1.85）
	defaultRatio := 1.0
	for _, r := range text {
		if unicode.Is(unicode.Han, r) {
			defaultRatio = 1.85
			break
		}
	}

	customParams, err := p.CustomParameterHandler()
	if err != nil || customParams == nil {
		return defaultRatio
	}
	// 尝试读取 audio.usage_ratio
	if audioRaw, ok := customParams["audio"]; ok {
		if audioMap, ok2 := audioRaw.(map[string]interface{}); ok2 {
			if v, ok3 := audioMap["usage_ratio"]; ok3 {
				if f, ok4 := anyToFloat64(v); ok4 && f > 0 {
					return f
				}
			}
		}
	}
	// 尝试读取根级 usage_ratio
	if v, ok := customParams["usage_ratio"]; ok {
		if f, ok2 := anyToFloat64(v); ok2 && f > 0 {
			return f
		}
	}
	return defaultRatio
}

func anyToFloat64(v interface{}) (float64, bool) {
	switch val := v.(type) {
	case float64:
		return val, true
	case float32:
		return float64(val), true
	case int:
		return float64(val), true
	case int64:
		return float64(val), true
	case json.Number:
		if f, err := val.Float64(); err == nil {
			return f, true
		}
		return 0, false
	case string:
		if strings.TrimSpace(val) == "" {
			return 0, false
		}
		if f, err := strconv.ParseFloat(strings.TrimSpace(val), 64); err == nil {
			return f, true
		}
		return 0, false
	default:
		return 0, false
	}
}

// estimateUsage 结合比例系数对字符数进行放大，向上取整
func estimateUsage(text string, ratio float64) int {
	base := estimateTextCharacters(text)
	if base == 0 {
		return 0
	}
	if ratio <= 0 {
		ratio = 1.0
	}
	return int(math.Ceil(float64(base) * ratio))
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
	// 根据上游选择路径：官方走统一 t2a_v2；PPInfra 走按模型拆分的 /v3/xxx
	var fullRequestURL string
	if p.isPPInfraSpeechUpstream() {
		modelName := request.Model
		if strings.TrimSpace(modelName) == "" {
			modelName = p.GetOriginalModel()
		}
		path := p.getPPInfraSpeechPath(modelName, false)
		fullRequestURL = p.GetFullRequestURL(path, modelName)
	} else {
		url, errWithCode := p.GetSupportedAPIUri(config.RelayModeAudioSpeech)
		if errWithCode != nil {
			return nil, errWithCode
		}
		fullRequestURL = p.GetFullRequestURL(url, request.Model)
	}
	headers := p.GetRequestHeaders()

	requestBody := p.getRequestBody(request)

	req, err := p.Requester.NewRequest(http.MethodPost, fullRequestURL, p.Requester.WithBody(requestBody), p.Requester.WithHeader(headers))
	if err != nil {
		return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError)
	}
	defer req.Body.Close()

	// PPInfra 上游：透传为主，但补写 usage，并可包装为官方结构体响应
	if p.isPPInfraSpeechUpstream() {
		resp, errWith := p.Requester.SendRequestRaw(req)
		if errWith != nil {
			return nil, errWith
		}
		// 补写 usage（按文本字符数×估算系数）
		if request != nil {
			ratio := p.getAudioUsageRatio(request.Input)
			usage := estimateUsage(request.Input, ratio)
			p.Usage.PromptTokens = usage
			p.Usage.TotalTokens = usage
		}
		// 尝试包装为官方结构
		// 从请求推断音频参数
		var fmtStr string
		var ch, sr, br *int64
		if requestBody != nil && requestBody.AudioSetting != nil {
			fmtStr = requestBody.AudioSetting.Format
			if requestBody.AudioSetting.Channel != 0 {
				v := requestBody.AudioSetting.Channel
				ch = &v
			}
			if requestBody.AudioSetting.SampleRate != 0 {
				v := requestBody.AudioSetting.SampleRate
				sr = &v
			}
			if requestBody.AudioSetting.Bitrate != 0 {
				v := requestBody.AudioSetting.Bitrate
				br = &v
			}
		}
		wrapped := p.wrapPPInfraOfficial(resp, fmtStr, ch, sr, br, p.Usage.TotalTokens)
		if wrapped != nil {
			return wrapped, nil
		}
		return resp, nil
	}

	speechResponse := &SpeechResponse{}
	_, errWithCode := p.Requester.SendRequest(req, speechResponse, false)
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
	// 回退：官方结构缺少 usage 时，用请求文本长度估算
	if usage == 0 && request != nil && strings.TrimSpace(request.Input) != "" {
		ratio := p.getAudioUsageRatio(request.Input)
		usage = estimateUsage(request.Input, ratio)
	}
	p.Usage.PromptTokens = usage
	p.Usage.TotalTokens = usage

	return response, nil
}

func (p *MiniMaxProvider) CreateSpeechAsync(req *types.MiniMaxAsyncSpeechRequest, rawBody []byte) (*http.Response, *types.OpenAIErrorWithStatusCode) {
	// PPInfra 使用按模型的 /v3/async/xxx
	var fullRequestURL string
	modelName := p.GetOriginalModel()
	if req != nil && strings.TrimSpace(req.Model) != "" {
		modelName = req.Model
	}
	if p.isPPInfraSpeechUpstream() {
		path := p.getPPInfraSpeechPath(modelName, true)
		fullRequestURL = p.GetFullRequestURL(path, modelName)
	} else {
		url, errWithCode := p.GetSupportedAPIUri(config.RelayModeMiniMaxSpeechAsync)
		if errWithCode != nil {
			return nil, errWithCode
		}
		fullRequestURL = p.GetFullRequestURL(url, modelName)
	}
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
	// 官方/PPInfra 分路构建 URL
	modelName := p.GetOriginalModel()
	if req != nil && req.Model != "" {
		modelName = req.Model
	}
	var fullRequestURL string
	if p.isPPInfraSpeechUpstream() {
		// 按模型映射 PPInfra 路径；支持 stream
		path := p.getPPInfraSpeechPath(modelName, false)
		fullRequestURL = p.GetFullRequestURL(path, modelName)
	} else {
		url, errWithCode := p.GetSupportedAPIUri(config.RelayModeAudioSpeech)
		if errWithCode != nil {
			return nil, errWithCode
		}
		fullRequestURL = p.GetFullRequestURL(url, modelName)
	}
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

	// 非流式 + PPInfra 上游：透传为主，但补写 usage，并可包装为官方结构
	if p.isPPInfraSpeechUpstream() {
		resp, errWithCode := p.Requester.SendRequestRaw(httpReq)
		if errWithCode != nil {
			return nil, errWithCode
		}
		// 补写 usage（优先使用 req.Text，其次从 payload 解析 text），并应用估算系数
		usage := 0
		if req != nil && strings.TrimSpace(req.Text) != "" {
			ratio := p.getAudioUsageRatio(req.Text)
			usage = estimateUsage(req.Text, ratio)
		} else if len(payload) > 0 {
			var m map[string]any
			if json.Unmarshal(payload, &m) == nil {
				if v, ok := m["text"].(string); ok {
					ratio := p.getAudioUsageRatio(v)
					usage = estimateUsage(v, ratio)
				}
			}
		}
		p.Usage.PromptTokens = usage
		p.Usage.TotalTokens = usage
		// 尝试包装为官方结构
		var fmtStr string
		var ch, sr, br *int64
		if req != nil && req.AudioSetting != nil {
			fmtStr = strings.TrimSpace(req.AudioSetting.Format)
			if req.AudioSetting.Channel != nil {
				ch = req.AudioSetting.Channel
			}
			if req.AudioSetting.SampleRate != nil {
				sr = req.AudioSetting.SampleRate
			}
			if req.AudioSetting.Bitrate != nil {
				br = req.AudioSetting.Bitrate
			}
		} else if len(payload) > 0 {
			var m map[string]any
			if json.Unmarshal(payload, &m) == nil {
				if as, ok := m["audio_setting"].(map[string]any); ok {
					if v, ok2 := as["format"].(string); ok2 {
						fmtStr = strings.TrimSpace(v)
					}
					if v, ok2 := as["channel"].(float64); ok2 {
						vv := int64(v)
						ch = &vv
					}
					if v, ok2 := as["sample_rate"].(float64); ok2 {
						vv := int64(v)
						sr = &vv
					}
					if v, ok2 := as["bitrate"].(float64); ok2 {
						vv := int64(v)
						br = &vv
					}
				}
			}
		}
		wrapped := p.wrapPPInfraOfficial(resp, fmtStr, ch, sr, br, p.Usage.TotalTokens)
		if wrapped != nil {
			return wrapped, nil
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
	// 回退：官方结构缺少 usage 时，用请求文本长度×估算系数估算
	if usage == 0 {
		if req != nil && strings.TrimSpace(req.Text) != "" {
			ratio := p.getAudioUsageRatio(req.Text)
			usage = estimateUsage(req.Text, ratio)
		} else if len(payload) > 0 {
			var m map[string]any
			if json.Unmarshal(payload, &m) == nil {
				if v, ok := m["text"].(string); ok {
					ratio := p.getAudioUsageRatio(v)
					usage = estimateUsage(v, ratio)
				}
			}
		}
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

// wrapPPInfraOfficial 尝试将 PPInfra 的简单响应包装为 MiniMax 官方结构
// 返回新的 *http.Response（application/json），若无法解析则返回 nil
func (p *MiniMaxProvider) wrapPPInfraOfficial(resp *http.Response, format string, channel, sampleRate, bitrate *int64, usage int) *http.Response {
	if resp == nil || resp.Body == nil {
		return nil
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	_ = resp.Body.Close()
	// 若已是官方结构，直接还原原响应
	var probe map[string]any
	if json.Unmarshal(body, &probe) == nil {
		if _, ok := probe["base_resp"]; ok {
			// 重新构造 response（因为我们已读尽 resp.Body）
			newResp := &http.Response{Status: "200 OK", StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(body)), Header: make(http.Header)}
			newResp.Header.Set("Content-Type", "application/json")
			return newResp
		}
		// 简单 JSON：尝试提取 audio/hex 或 url
		var audioHex string
		if v, ok := probe["audio"].(string); ok {
			audioHex = strings.TrimSpace(v)
		}
		// 构造官方结构
		extra := map[string]any{}
		if strings.TrimSpace(format) != "" {
			extra["audio_format"] = format
		}
		if channel != nil {
			extra["audio_channel"] = *channel
		}
		if sampleRate != nil {
			extra["audio_sample_rate"] = *sampleRate
		}
		if bitrate != nil {
			extra["bitrate"] = *bitrate
		}
		if usage > 0 {
			extra["usage_characters"] = usage
		}
		// 估算 audio_size（若为 hex）
		if audioHex != "" {
			if b, err2 := hex.DecodeString(audioHex); err2 == nil {
				extra["audio_size"] = int64(len(b))
			}
		}
		official := map[string]any{
			"data": map[string]any{
				"audio":  audioHex,
				"status": 2,
			},
			"extra_info": extra,
			"trace_id":   "",
			"base_resp":  map[string]any{"status_code": 0, "status_msg": "success"},
		}
		buf, err3 := json.Marshal(official)
		if err3 != nil {
			return nil
		}
		newResp := &http.Response{Status: "200 OK", StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(buf)), Header: make(http.Header)}
		newResp.Header.Set("Content-Type", "application/json")
		return newResp
	}
	// 非 JSON（可能是音频流）不包装
	// 重新放回 body 以便后续透传
	resp.Body = io.NopCloser(bytes.NewReader(body))
	return nil
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
