package kling

type KlingQueryRequest struct {
	TaskID string `json:"task_id"`
	// ExternalTaskID string `json:"external_task_id,omitempty"`
}

type KlingTask struct {
	Prompt    string `json:"prompt,omitempty"`
	ModelName string `json:"model_name,omitempty"`
	Mode      string `json:"mode,omitempty"`

	Image            string `json:"image,omitempty"`
	ImageTail        string `json:"image_tail,omitempty"`
	StaticMask       string `json:"static_mask,omitempty"`
	DynamicMasks     any    `json:"dynamic_masks,omitempty"`
	ImageList        any    `json:"image_list,omitempty"`
	SubjectImageList any    `json:"subject_image_list,omitempty"`
	SceneImage       string `json:"scene_image,omitempty"`
	StyleImage       string `json:"style_image,omitempty"`
	SessionID        string `json:"session_id,omitempty"`
	EditMode         string `json:"edit_mode,omitempty"`

	NegativePrompt string   `json:"negative_prompt,omitempty"`
	CfgScale       *float64 `json:"cfg_scale,omitempty"`
	CameraControl  any      `json:"camera_control,omitempty"`

	AspectRatio    any      `json:"aspect_ratio,omitempty"`
	Duration       string   `json:"duration,omitempty"`
	CallbackURL    any      `json:"callback_url,omitempty"`
	ImageReference string   `json:"image_reference,omitempty"`
	ImageFidelity  *float64 `json:"image_fidelity,omitempty"`
	HumanFidelity  *float64 `json:"human_fidelity,omitempty"`
	Resolution     string   `json:"resolution,omitempty"`
	N              int      `json:"n,omitempty"`
}

// type KlingVideoExtend struct {
// 	VideoID     string `json:"video_id"`
// 	Prompt      string `json:"prompt,omitempty"`
// 	CallbackURL any    `json:"callback_url,omitempty"`
// }

type KlingResponse[T any] struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
	Data      T      `json:"data,omitempty"`
}

type KlingTaskData struct {
	TaskID        string           `json:"task_id"`
	TaskStatus    string           `json:"task_status"`
	CreatedAt     int64            `json:"created_at"`
	UpdatedAt     int64            `json:"updated_at"`
	TaskStatusMsg string           `json:"task_status_msg,omitempty"`
	TaskResult    *KlingTaskResult `json:"task_result,omitempty"`
	SessionID     string           `json:"session_id,omitempty"`
}

type KlingTaskResult struct {
	Videos []KlingVideoResult `json:"videos"`
}

type KlingVideoResult struct {
	ID       string `json:"id"`
	URL      string `json:"url"`
	Duration string `json:"duration"`
}
