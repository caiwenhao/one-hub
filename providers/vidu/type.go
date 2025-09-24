package vidu

// Vidu任务提交请求结构体
type ViduTaskRequest struct {
	Model             string   `json:"model"`                        // 模型名，根据官方文档支持：viduq1, vidu1.5, vidu2.0, viduq1-classic
	Images            []string `json:"images,omitempty"`             // 图片URL或Base64 (img2video)
	ReferenceVideos   []string `json:"reference_videos,omitempty"`   // 参考视频URL (reference2video)
	StartImage        string   `json:"start_image,omitempty"`        // 起始图片 (start-end2video)
	EndImage          string   `json:"end_image,omitempty"`          // 终止图片 (start-end2video)
	Prompt            string   `json:"prompt,omitempty"`             // 文本描述，最长2000字符
	Duration          *int     `json:"duration,omitempty"`           // 视频时长 4s/5s/8s
	Seed              *int     `json:"seed,omitempty"`               // 随机种子
	Resolution        string   `json:"resolution,omitempty"`         // 分辨率 1080p/720p/360p
	MovementAmplitude string   `json:"movement_amplitude,omitempty"` // 运动幅度 auto/small/medium/large
	BGM               *bool    `json:"bgm,omitempty"`                // 是否添加背景音乐
	Payload           string   `json:"payload,omitempty"`            // 透传参数
	OffPeak           *bool    `json:"off_peak,omitempty"`           // 离峰模式
	CallbackURL       string   `json:"callback_url,omitempty"`       // 任务状态回调接口
}

// Vidu查询请求结构体
type ViduQueryRequest struct {
	TaskID string `json:"task_id"`
}

// Vidu响应结构体
type ViduResponse[T any] struct {
	TaskID    string `json:"task_id"`
	Status    string `json:"status"`      // created/queueing/processing/success/failed
	Message   string `json:"message"`     
	Data      T      `json:"data,omitempty"`
	Credits   int    `json:"credits"`     // 消耗点数
	CreatedAt string `json:"created_at"`  // 创建时间，格式为 RFC3339 字符串
}

// Vidu任务数据
type ViduTaskData struct {
	TaskID            string                 `json:"task_id"`
	Status            string                 `json:"status"`
	Message           string                 `json:"message,omitempty"`
	CreatedAt         string                 `json:"created_at"`  // RFC3339 格式字符串
	UpdatedAt         string                 `json:"updated_at,omitempty"` // RFC3339 格式字符串
	Model             string                 `json:"model"`
	Prompt            string                 `json:"prompt,omitempty"`
	Duration          *int                   `json:"duration,omitempty"`
	Resolution        string                 `json:"resolution,omitempty"`
	MovementAmplitude string                 `json:"movement_amplitude,omitempty"`
	BGM               *bool                  `json:"bgm,omitempty"`
	Credits           int                    `json:"credits"`
	Videos            []ViduVideoResult      `json:"videos,omitempty"`
	Images            []string               `json:"images,omitempty"`
	ReferenceVideos   []string               `json:"reference_videos,omitempty"`
	StartImage        string                 `json:"start_image,omitempty"`
	EndImage          string                 `json:"end_image,omitempty"`
	Extra             map[string]interface{} `json:"extra,omitempty"`
}

// Vidu视频结果
type ViduVideoResult struct {
	ID       string `json:"id"`
	URL      string `json:"url"`
	Duration string `json:"duration"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
}

// Vidu查询响应中的创作结果
type ViduCreation struct {
	ID       string `json:"id"`        // 创作ID
	URL      string `json:"url"`       // 生成结果的URL，有效期1小时
	CoverURL string `json:"cover_url"` // 封面URL，有效期1小时
}

// Vidu查询任务状态的响应结构体（根据官方文档）
type ViduQueryResponse struct {
	State     string          `json:"state"`     // created/queueing/processing/success/failed
	ErrCode   string          `json:"err_code"`  // 错误代码
	Credits   int             `json:"credits"`   // 消耗的积分数
	Payload   string          `json:"payload"`   // 透传参数
	BGM       *bool           `json:"bgm"`       // 是否使用BGM
	OffPeak   *bool           `json:"off_peak"`  // 是否使用离峰模式
	Creations []ViduCreation `json:"creations"` // 生成结果数组
}

// Vidu错误响应
type ViduErrorResponse struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error"`
}

// Vidu任务状态常量
const (
	ViduStatusCreated    = "created"
	ViduStatusQueueing   = "queueing"  
	ViduStatusProcessing = "processing"
	ViduStatusSuccess    = "success"
	ViduStatusFailed     = "failed"
)

// Vidu接口路径
const (
	ViduActionImg2Video        = "img2video"
	ViduActionReference2Video  = "reference2video"
	ViduActionStartEnd2Video   = "start-end2video"
	ViduActionText2Video       = "text2video"
)