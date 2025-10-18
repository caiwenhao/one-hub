package minimaxi

import (
	"encoding/json"
	"strings"
)

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
	Audio  string `json:"audio"` // hex编码的audio
	Status int    `json:"status"`
}

type ExtraInfo struct {
	AudioLength             int64   `json:"audio_length"`
	AudioSampleRate         int64   `json:"audio_sample_rate"`
	AudioSize               int64   `json:"audio_size"`
	AudioBitrate            int64   `json:"audio_bitrate"`
	WordCount               int64   `json:"word_count"`
	InvisibleCharacterRatio float64 `json:"invisible_character_ratio"`
	AudioFormat             string  `json:"audio_format"`
	UsageCharacters         int     `json:"usage_characters"`
}

// StringOrNumber 用于兼容上游既可能返回字符串，也可能返回数字的字段
// 例如 MiniMax 的文件接口中 file_id 可能是数值类型
type StringOrNumber string

func (s *StringOrNumber) UnmarshalJSON(b []byte) error {
	// 如果是带引号的字符串，正常反序列化
	if len(b) > 0 && (b[0] == '"') {
		var v string
		if err := json.Unmarshal(b, &v); err != nil {
			return err
		}
		*s = StringOrNumber(v)
		return nil
	}
	// 否则直接把原始数字 token 作为字符串保存，避免大整型精度问题
	*s = StringOrNumber(strings.TrimSpace(string(b)))
	return nil
}
