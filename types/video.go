package types

type VideoCreateRequest struct {
	Model            string   `json:"model"`
	Prompt           string   `json:"prompt,omitempty"`
	Seconds          int      `json:"seconds,omitempty"`
	Size             string   `json:"size,omitempty"`
	RemoveWatermark  bool     `json:"remove_watermark,omitempty"`
	InputImage       string   `json:"input_image,omitempty"`
	InputImages      []string `json:"input_images,omitempty"`
	InputReference   string   `json:"input_reference,omitempty"`
	RemixVideoID     string   `json:"remix_video_id,omitempty"`
	Seed             string   `json:"seed,omitempty"`
}

type VideoJob struct {
	ID        string          `json:"id"`
	Object    string          `json:"object,omitempty"`
	CreatedAt int64           `json:"created_at,omitempty"`
	Status    string          `json:"status"`
	Model     string          `json:"model,omitempty"`
	Progress  float64         `json:"progress,omitempty"`
	Seconds   int             `json:"seconds,omitempty"`
	Size      string          `json:"size,omitempty"`
	Error     *VideoJobError  `json:"error,omitempty"`
	Result    *VideoJobResult `json:"result,omitempty"`
	Metadata  any             `json:"metadata,omitempty"`
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
