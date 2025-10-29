package minimaxi

type MiniMaxBaseResp struct {
    // 基本返回体
    BaseResp BaseResp `json:"base_resp"`
}

type BaseResp struct {
    // 状态码
    StatusCode int64  `json:"status_code"`
    // 状态信息
    StatusMsg  string `json:"status_msg"`
}

type SpeechRequest struct {
    // 模型名称，例如：speech-2.5-hd-preview
    Model        string        `json:"model" example:"speech-2.5-hd-preview"`
    // 合成文本内容
    Text         string        `json:"text" example:"欢迎使用 MiniMax 同步合成"`
    // 音色与情感
    VoiceSetting VoiceSetting  `json:"voice_setting"`
    // 音频设置
    AudioSetting *AudioSetting `json:"audio_setting"`
}

type VoiceSetting struct {
    // 语速（0.5~2.0）
    Speed     float64  `json:"speed,omitempty"`
    // 音量（0~2）
    Vol       *float64 `json:"vol,omitempty"`
    // 音色 ID
    VoiceID   string   `json:"voice_id" example:"female-chengshu"`
    // 情感
    Emotion   string   `json:"emotion,omitempty" enums:"neutral,happy,sad,angry"`
    // 是否朗读 LaTeX 公式
    LatexRead bool     `json:"latex_read,omitempty"`
}

type AudioSetting struct {
    // 采样率（Hz）
    SampleRate int64  `json:"sample_rate,omitempty"`
    // 比特率（bps）
    Bitrate    int64  `json:"bitrate,omitempty"`
    // 编码格式
    Format     string `json:"format" enums:"mp3,wav,aac,flac,pcm"`
    // 声道
    Channel    int64  `json:"channel,omitempty" enums:"1,2"`
}

type SpeechResponse struct {
    // 基本返回体
    BaseResp  BaseResp  `json:"base_resp"`
    // 语音数据
    Data      Data      `json:"data"`
    // 额外信息
    ExtraInfo ExtraInfo `json:"extra_info"`
}

type Data struct {
    // hex 编码的音频
    Audio        string  `json:"audio" example:"48656c6c6f2c206d703320..."`
    // 字幕文件 URL（可选）
    SubtitleFile *string `json:"subtitle_file,omitempty"`
    // 状态码
    Status       int     `json:"status"`
}

type ExtraInfo struct {
    // 音频时长（ms）
    AudioLength             *int64   `json:"audio_length,omitempty"`
    // 采样率
    AudioSampleRate         *int64   `json:"audio_sample_rate,omitempty"`
    // 文件大小（byte）
    AudioSize               *int64   `json:"audio_size,omitempty"`
    // 比特率（bps）
    Bitrate                 *int64   `json:"bitrate,omitempty"`
    // 字数
    WordCount               *int64   `json:"word_count,omitempty"`
    // 不可见字符占比
    InvisibleCharacterRatio *float64 `json:"invisible_character_ratio,omitempty"`
    // 音频格式
    AudioFormat             *string  `json:"audio_format,omitempty"`
    // 声道
    AudioChannel            *int64   `json:"audio_channel,omitempty"`
    // 计费字符数
    UsageCharacters         *int     `json:"usage_characters,omitempty"`
}
