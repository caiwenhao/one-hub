package types

type VideoCreateRequest struct {
	// 视频生成模型
	Model string `json:"model" form:"model"`
	// 文本提示词
	Prompt string `json:"prompt,omitempty" form:"prompt"`
	// 视频时长（秒）
	Seconds int `json:"seconds,omitempty,string" form:"seconds"`
	// 输出分辨率（例如：1280x720、720x1280）
	Size string `json:"size,omitempty" form:"size"`
	// 是否去除水印
	RemoveWatermark bool `json:"remove_watermark,omitempty" form:"remove_watermark"`
	// 图生视频：单张输入图
	InputImage string `json:"input_image,omitempty" form:"input_image"`
	// 图生视频：多张输入图
	InputImages []string `json:"input_images,omitempty" form:"input_images"`
	// 参考图/文件ID
	InputReference string `json:"input_reference,omitempty" form:"input_reference"`
	// Remix：来源视频ID
	RemixVideoID string `json:"remix_video_id,omitempty" form:"remix_video_id"`
	// 随机种子
	Seed string `json:"seed,omitempty" form:"seed"`
}

// VideoRemixRequest 用于 /v1/videos/{id}/remix
type VideoRemixRequest struct {
	Prompt string `json:"prompt" form:"prompt"`
}

type VideoJob struct {
	ID                 string          `json:"id"`
	Object             string          `json:"object,omitempty"`
	CreatedAt          int64           `json:"created_at,omitempty"`
	CompletedAt        int64           `json:"completed_at,omitempty"`
	ExpiresAt          int64           `json:"expires_at,omitempty"`
	Status             string          `json:"status"`
	Model              string          `json:"model,omitempty"`
	Prompt             string          `json:"prompt,omitempty"`
	Progress           float64         `json:"progress,omitempty"`
	Seconds            int             `json:"seconds,omitempty,string"`
	Size               string          `json:"size,omitempty"`
	Quality            string          `json:"quality,omitempty"`
	RemixedFromVideoID string          `json:"remixed_from_video_id,omitempty"`
	Error              *VideoJobError  `json:"error,omitempty"`
	Result             *VideoJobResult `json:"result,omitempty"`
	Metadata           any             `json:"metadata,omitempty"`
}

type VideoJobError struct {
	Code    any    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type VideoJobResult struct {
	VideoURL      string `json:"video_url,omitempty"`
	ThumbnailURL  string `json:"thumbnail_url,omitempty"`
	SpriteSheet   string `json:"spritesheet_url,omitempty"`
	Variant       string `json:"variant,omitempty"`
	DownloadURL   string `json:"download_url,omitempty"`
	AdditionalRaw any    `json:"raw,omitempty"`
}

// VideoList 对齐 OpenAI 列表返回
type VideoList struct {
	Object string     `json:"object"`
	Data   []VideoJob `json:"data"`
}
