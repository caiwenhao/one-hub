package types

import (
	"encoding/json"
	"mime/multipart"
)

type SpeechAudioRequest struct {
	Model          string  `json:"model"`
	Input          string  `json:"input"`
	Voice          string  `json:"voice"`
	ResponseFormat string  `json:"response_format,omitempty"`
	Speed          float64 `json:"speed,omitempty"`
}

type MiniMaxSpeechRequest struct {
	Model             string                    `json:"model" binding:"required"`
	Text              string                    `json:"text" binding:"required"`
	Stream            *bool                     `json:"stream,omitempty"`
	StreamOptions     *MiniMaxStreamOptions     `json:"stream_options,omitempty"`
	VoiceSetting      *MiniMaxVoiceSetting      `json:"voice_setting,omitempty"`
	AudioSetting      *MiniMaxAudioSetting      `json:"audio_setting,omitempty"`
	PronunciationDict *MiniMaxPronunciationDict `json:"pronunciation_dict,omitempty"`
	TimberWeights     []MiniMaxTimbreWeight     `json:"timber_weights,omitempty"`
	LanguageBoost     *string                   `json:"language_boost,omitempty"`
	VoiceModify       *MiniMaxVoiceModify       `json:"voice_modify,omitempty"`
	SubtitleEnable    *bool                     `json:"subtitle_enable,omitempty"`
	OutputFormat      string                    `json:"output_format,omitempty"`
	AIGCWatermark     *bool                     `json:"aigc_watermark,omitempty"`
}

type MiniMaxStreamOptions struct {
	ExcludeAggregatedAudio *bool `json:"exclude_aggregated_audio,omitempty"`
}

type MiniMaxVoiceSetting struct {
	VoiceID           string   `json:"voice_id"`
	Speed             *float64 `json:"speed,omitempty"`
	Vol               *float64 `json:"vol,omitempty"`
	Pitch             *int     `json:"pitch,omitempty"`
	Emotion           *string  `json:"emotion,omitempty"`
	TextNormalization *bool    `json:"text_normalization,omitempty"`
	LatexRead         *bool    `json:"latex_read,omitempty"`
}

type MiniMaxAudioSetting struct {
	SampleRate *int64 `json:"sample_rate,omitempty"`
	Bitrate    *int64 `json:"bitrate,omitempty"`
	Format     string `json:"format,omitempty"`
	Channel    *int64 `json:"channel,omitempty"`
	ForceCBR   *bool  `json:"force_cbr,omitempty"`
}

type MiniMaxPronunciationDict struct {
	Tone []string `json:"tone,omitempty"`
}

type MiniMaxTimbreWeight struct {
	VoiceID string `json:"voice_id"`
	Weight  int    `json:"weight"`
}

type MiniMaxVoiceModify struct {
	Pitch        *int    `json:"pitch,omitempty"`
	Intensity    *int    `json:"intensity,omitempty"`
	Timbre       *int    `json:"timbre,omitempty"`
	SoundEffects *string `json:"sound_effects,omitempty"`
}

type MiniMaxAsyncSpeechRequest struct {
	Model          string                    `json:"model"`
	Text           *string                   `json:"text,omitempty"`
	TextFileID     *json.RawMessage          `json:"text_file_id,omitempty"`
	VoiceSetting   *MiniMaxAsyncVoiceSetting `json:"voice_setting,omitempty"`
	AudioSetting   *MiniMaxAsyncAudioSetting `json:"audio_setting,omitempty"`
	Pronunciation  *MiniMaxPronunciationDict `json:"pronunciation_dict,omitempty"`
	LanguageBoost  *string                   `json:"language_boost,omitempty"`
	VoiceModify    *MiniMaxVoiceModify       `json:"voice_modify,omitempty"`
	AIGCWatermark  *bool                     `json:"aigc_watermark,omitempty"`
	TimberWeights  []MiniMaxTimbreWeight     `json:"timbre_weights,omitempty"`
	SubtitleEnable *bool                     `json:"subtitle_enable,omitempty"`
	OutputFormat   *string                   `json:"output_format,omitempty"`
	Stream         *bool                     `json:"stream,omitempty"`
	StreamOptions  *MiniMaxStreamOptions     `json:"stream_options,omitempty"`
}

type MiniMaxAsyncVoiceSetting struct {
	VoiceID              string   `json:"voice_id"`
	Speed                *float64 `json:"speed,omitempty"`
	Vol                  *float64 `json:"vol,omitempty"`
	Pitch                *int     `json:"pitch,omitempty"`
	Emotion              *string  `json:"emotion,omitempty"`
	EnglishNormalization *bool    `json:"english_normalization,omitempty"`
}

type MiniMaxAsyncAudioSetting struct {
	AudioSampleRate *int64  `json:"audio_sample_rate,omitempty"`
	Bitrate         *int64  `json:"bitrate,omitempty"`
	Format          *string `json:"format,omitempty"`
	Channel         *int64  `json:"channel,omitempty"`
}

type MiniMaxAsyncSpeechResponse struct {
	TaskID          StringOrNumber       `json:"task_id,omitempty"`
	FileID          *StringOrNumber      `json:"file_id,omitempty"`
	TaskToken       string               `json:"task_token,omitempty"`
	UsageCharacters *int                 `json:"usage_characters,omitempty"`
	BaseResp        MiniMaxAsyncBaseResp `json:"base_resp"`
	TraceID         string               `json:"trace_id,omitempty"`
	ExtraInfo       json.RawMessage      `json:"extra_info,omitempty"`
}

type MiniMaxAsyncBaseResp struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type MiniMaxAsyncSpeechQueryResponse struct {
	TaskID    StringOrNumber       `json:"task_id,omitempty"`
	Status    string               `json:"status,omitempty"`
	FileID    *StringOrNumber      `json:"file_id,omitempty"`
	BaseResp  MiniMaxAsyncBaseResp `json:"base_resp"`
	ExtraInfo json.RawMessage      `json:"extra_info,omitempty"`
}

type AudioRequest struct {
	File           *multipart.FileHeader `form:"file" binding:"required"`
	Model          string                `form:"model" binding:"required"`
	Language       string                `form:"language"`
	Prompt         string                `form:"prompt"`
	ResponseFormat string                `form:"response_format"`
	Temperature    float32               `form:"temperature"`
}

type AudioResponse struct {
	Task     string           `json:"task,omitempty"`
	Language string           `json:"language,omitempty"`
	Duration float64          `json:"duration,omitempty"`
	Segments any              `json:"segments,omitempty"`
	Text     string           `json:"text"`
	Words    []AudioWordsList `json:"words,omitempty"`
}

type AudioWordsList struct {
	Word  string  `json:"word"`
	Start float64 `json:"start"`
	End   float64 `json:"end"`
}

type AudioResponseWrapper struct {
	Headers map[string]string
	Body    []byte
}
