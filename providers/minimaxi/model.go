package minimaxi

import "errors"

// GetModelList 返回 minimaxi 支持的模型列表
// 优先尝试通过 OpenAI 兼容的 /v1/models 获取；失败时回退到内置静态清单。
func (p *MiniMaxProvider) GetModelList() ([]string, error) {
    // 优先走上游 OpenAI 兼容接口
    if p.Config.ModelList != "" {
        if list, err := p.OpenAIProvider.GetModelList(); err == nil && len(list) > 0 {
            return list, nil
        }
    }

    // 本地静态列表（覆盖文本 / 语音 / 视频核心模型）
    list := []string{
        // 文本
        "MiniMax-M1",
        "MiniMax-Text-01",
        // 语音（TTS）
        // 2.6 新版语音模型：与 02 同价位，保持静态清单可见性，便于前端选择
        "speech-2.6-hd",
        "speech-2.6-turbo",
        "speech-2.5-hd-preview",
        "speech-02-hd",
        "speech-01-hd",
        "speech-2.5-turbo-preview",
        "speech-02-turbo",
        "speech-01-turbo",
        // 视频（官方/PPInfra 常用）
        "MiniMax-Hailuo-02",
        "MiniMax-Hailuo-2.3",
        "MiniMax-Hailuo-2.3-Fast",
        "T2V-01",
        "T2V-01-Director",
        "I2V-01",
        "I2V-01-Director",
        "I2V-01-live",
        "S2V-01",
        // 音乐与图像
        "music-1.5",
        "image-01",
        "image-01-live",
    }
    if len(list) == 0 {
        return nil, errors.New("no models available")
    }
    return list, nil
}
