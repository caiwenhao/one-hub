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
	// 模型名称，例如：speech-2.5-hd-preview
	Model             string                    `json:"model" binding:"required" example:"speech-2.5-hd-preview"`
	// 合成文本内容
	Text              string                    `json:"text" binding:"required" example:"欢迎使用 One Hub 与 MiniMax 语音合成功能"`
	// 是否开启流式（仅透传能力，非必须）
	Stream            *bool                     `json:"stream,omitempty"`
	// 流式选项
	StreamOptions     *MiniMaxStreamOptions     `json:"stream_options,omitempty"`
	// 音色及情感设置
	VoiceSetting      *MiniMaxVoiceSetting      `json:"voice_setting,omitempty"`
	// 音频输出参数
	AudioSetting      *MiniMaxAudioSetting      `json:"audio_setting,omitempty"`
	// 发音词典（可选）
	PronunciationDict *MiniMaxPronunciationDict `json:"pronunciation_dict,omitempty"`
	TimberWeights     []MiniMaxTimbreWeight     `json:"timber_weights,omitempty"`
	LanguageBoost     *string                   `json:"language_boost,omitempty"`
	// 后期变声（可选）
	VoiceModify       *MiniMaxVoiceModify       `json:"voice_modify,omitempty"`
	// 是否生成字幕（可选）
	SubtitleEnable    *bool                     `json:"subtitle_enable,omitempty"`
	// 输出格式（当上游支持）
	OutputFormat      string                    `json:"output_format,omitempty" enums:"url,base64,hex"`
	// 是否添加 AIGC 水印
	AIGCWatermark     *bool                     `json:"aigc_watermark,omitempty"`
}

type MiniMaxStreamOptions struct {
	ExcludeAggregatedAudio *bool `json:"exclude_aggregated_audio,omitempty"`
}

type MiniMaxVoiceSetting struct {
	// 音色 ID（可支持别名：如 alloy/echo 等）
	VoiceID           string   `json:"voice_id" example:"female-chengshu"`
	// 语速（0.5~2.0），1 为正常
	Speed             *float64 `json:"speed,omitempty"`
	// 音量（0~2），1 为正常
	Vol               *float64 `json:"vol,omitempty"`
	// 音高（整数，单位半音）
	Pitch             *int     `json:"pitch,omitempty"`
	// 情感
	Emotion           *string  `json:"emotion,omitempty" enums:"neutral,happy,sad,angry"`
	// 文本规范化（英文）
	TextNormalization *bool    `json:"text_normalization,omitempty"`
	// 是否朗读 LaTeX 公式
	LatexRead         *bool    `json:"latex_read,omitempty"`
}

type MiniMaxAudioSetting struct {
	// 采样率（Hz）
	SampleRate *int64 `json:"sample_rate,omitempty"`
	// 比特率（bps），如 128000
	Bitrate    *int64 `json:"bitrate,omitempty"`
	// 封装/编码格式
	Format     string `json:"format,omitempty" enums:"mp3,wav,aac,flac,pcm"`
	// 声道数
	Channel    *int64 `json:"channel,omitempty" enums:"1,2"`
	// 是否强制 CBR
	ForceCBR   *bool  `json:"force_cbr,omitempty"`
}

type MiniMaxPronunciationDict struct {
	// 语调字典，例如优先级规则等
	Tone []string `json:"tone,omitempty"`
}

type MiniMaxTimbreWeight struct {
	// 复合音色的单个音色 ID
	VoiceID string `json:"voice_id"`
	// 权重（1~100）
	Weight  int    `json:"weight"`
}

type MiniMaxVoiceModify struct {
	// 后期音高调节
	Pitch        *int    `json:"pitch,omitempty"`
	// 音强
	Intensity    *int    `json:"intensity,omitempty"`
	// 音色
	Timbre       *int    `json:"timbre,omitempty"`
	// 特效，如 echo, robot
	SoundEffects *string `json:"sound_effects,omitempty" enums:"none,echo,robot"`
}

type MiniMaxAsyncSpeechRequest struct {
	// 模型名称
	Model          string                    `json:"model" example:"speech-2.5-hd-preview"`
	// 文本（与 text_file_id 二选一）
	Text           *string                   `json:"text,omitempty" example:"hello, this is an async task"`
	// 文本文件 ID（与 text 二选一），可传上游返回的 file_id
	TextFileID     *json.RawMessage          `json:"text_file_id,omitempty"`
	// 音色设置
	VoiceSetting   *MiniMaxAsyncVoiceSetting `json:"voice_setting,omitempty"`
	// 音频设置
	AudioSetting   *MiniMaxAsyncAudioSetting `json:"audio_setting,omitempty"`
	// 发音词典
	Pronunciation  *MiniMaxPronunciationDict `json:"pronunciation_dict,omitempty"`
	// 语言增强
	LanguageBoost  *string                   `json:"language_boost,omitempty"`
	// 变声设置
	VoiceModify    *MiniMaxVoiceModify       `json:"voice_modify,omitempty"`
	// AIGC 水印
	AIGCWatermark  *bool                     `json:"aigc_watermark,omitempty"`
	// 音色权重
	TimberWeights  []MiniMaxTimbreWeight     `json:"timbre_weights,omitempty"`
	// 字幕
	SubtitleEnable *bool                     `json:"subtitle_enable,omitempty"`
	// 输出格式
	OutputFormat   *string                   `json:"output_format,omitempty" enums:"url,base64,hex"`
	// 是否流式
	Stream         *bool                     `json:"stream,omitempty"`
	// 流式选项
	StreamOptions  *MiniMaxStreamOptions     `json:"stream_options,omitempty"`
}

type MiniMaxAsyncVoiceSetting struct {
	// 音色 ID
	VoiceID              string   `json:"voice_id" example:"female-chengshu"`
	// 语速
	Speed                *float64 `json:"speed,omitempty"`
	// 音量
	Vol                  *float64 `json:"vol,omitempty"`
	// 音高
	Pitch                *int     `json:"pitch,omitempty"`
	// 情感
	Emotion              *string  `json:"emotion,omitempty" enums:"neutral,happy,sad,angry"`
	// 英文规范化
	EnglishNormalization *bool    `json:"english_normalization,omitempty"`
}

type MiniMaxAsyncAudioSetting struct {
	// 采样率（Hz）
	AudioSampleRate *int64  `json:"audio_sample_rate,omitempty"`
	// 比特率（bps）
	Bitrate         *int64  `json:"bitrate,omitempty"`
	// 编码格式
	Format          *string `json:"format,omitempty" enums:"mp3,wav,aac,flac,pcm"`
	// 声道
	Channel         *int64  `json:"channel,omitempty" enums:"1,2"`
}

type MiniMaxAsyncSpeechResponse struct {
	// 异步任务 ID
	TaskID          StringOrNumber       `json:"task_id,omitempty" example:"1234567890"`
	// 上游返回的文件 ID（当任务完成时）
	FileID          *StringOrNumber      `json:"file_id,omitempty"`
	// 任务令牌（有些上游返回）
	TaskToken       string               `json:"task_token,omitempty"`
	// 计费字符数
	UsageCharacters *int                 `json:"usage_characters,omitempty"`
	// 基本返回体
	BaseResp        MiniMaxAsyncBaseResp `json:"base_resp"`
	// Trace ID
	TraceID         string               `json:"trace_id,omitempty"`
	// 额外信息（透传）
	ExtraInfo       json.RawMessage      `json:"extra_info,omitempty"`
}

type MiniMaxAsyncBaseResp struct {
	// 状态码，0 表示成功
	StatusCode int64  `json:"status_code"`
	// 状态信息
	StatusMsg  string `json:"status_msg"`
}

type MiniMaxAsyncSpeechQueryResponse struct {
	// 任务 ID
	TaskID    StringOrNumber       `json:"task_id,omitempty" example:"1234567890"`
	// 任务状态
	Status    string               `json:"status,omitempty" enums:"queueing,processing,success,failed"`
	// 文件 ID（当成功时）
	FileID    *StringOrNumber      `json:"file_id,omitempty"`
	// 基本返回体
	BaseResp  MiniMaxAsyncBaseResp `json:"base_resp"`
	// 额外信息
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
