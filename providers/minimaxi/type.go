package minimaxi

type MiniMaxBaseResp struct {
	BaseResp BaseResp `json:"base_resp"`
}

type BaseResp struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type SpeechRequest struct {
	Model        string        `json:"model"`
	Text         string        `json:"text"`
	VoiceSetting VoiceSetting  `json:"voice_setting"`
	AudioSetting *AudioSetting `json:"audio_setting"`
}

type VoiceSetting struct {
	Speed     float64  `json:"speed,omitempty"`
	Vol       *float64 `json:"vol,omitempty"`
	VoiceID   string   `json:"voice_id"`
	Emotion   string   `json:"emotion,omitempty"`
	LatexRead bool     `json:"latex_read,omitempty"`
}

type AudioSetting struct {
	SampleRate int64  `json:"sample_rate,omitempty"`
	Bitrate    int64  `json:"bitrate,omitempty"`
	Format     string `json:"format"`
	Channel    int64  `json:"channel,omitempty"`
}

type SpeechResponse struct {
	BaseResp  BaseResp  `json:"base_resp"`
	Data      Data      `json:"data"`
	ExtraInfo ExtraInfo `json:"extra_info"`
}

type Data struct {
	Audio        string  `json:"audio"` // hex编码的audio
	SubtitleFile *string `json:"subtitle_file,omitempty"`
	Status       int     `json:"status"`
}

type ExtraInfo struct {
	AudioLength             *int64   `json:"audio_length,omitempty"`
	AudioSampleRate         *int64   `json:"audio_sample_rate,omitempty"`
	AudioSize               *int64   `json:"audio_size,omitempty"`
	Bitrate                 *int64   `json:"bitrate,omitempty"`
	WordCount               *int64   `json:"word_count,omitempty"`
	InvisibleCharacterRatio *float64 `json:"invisible_character_ratio,omitempty"`
	AudioFormat             *string  `json:"audio_format,omitempty"`
	AudioChannel            *int64   `json:"audio_channel,omitempty"`
	UsageCharacters         *int     `json:"usage_characters,omitempty"`
}
