package vidu

// Vidu任务提交请求结构体 - 完全对齐官方文档
type ViduTaskRequest struct {
	// 通用参数
	Model             string   `json:"model"`                        // 模型名：viduq2-pro、viduq2-turbo、viduq1、viduq1-classic、vidu2.0、vidu1.5
	Prompt            string   `json:"prompt,omitempty"`             // 文本提示词，最长2000字符
	Duration          *int     `json:"duration,omitempty"`           // 视频时长
	Seed              *int     `json:"seed,omitempty"`               // 随机种子
	Resolution        string   `json:"resolution,omitempty"`         // 分辨率：720p、1080p、360p
	MovementAmplitude string   `json:"movement_amplitude,omitempty"` // 运动幅度：auto、small、medium、large
	BGM               *bool    `json:"bgm,omitempty"`                // 是否添加背景音乐
    Payload           string   `json:"payload,omitempty"`            // 透传参数，最多1048576个字符
    OffPeak           *bool    `json:"off_peak,omitempty"`           // 错峰模式
    Watermark         *bool    `json:"watermark,omitempty"`          // 是否添加水印
    // 额外：与文档对齐的可选参数
    IsRec             *bool    `json:"is_rec,omitempty"`             // 是否使用推荐提示词
    WMPosition        *int     `json:"wm_position,omitempty"`        // 水印位置 1/2/3/4
    WMURL             string   `json:"wm_url,omitempty"`             // 自定义水印 URL
    MetaData          string   `json:"meta_data,omitempty"`          // 元数据标识（json 字符串）
    CallbackURL       string   `json:"callback_url,omitempty"`       // 回调URL

	// 图生视频专用 (img2video)
	Images []string `json:"images,omitempty"` // 首帧图像，只支持1张图

	// 参考生视频专用 (reference2video)
	AspectRatio string `json:"aspect_ratio,omitempty"` // 比例：16:9、9:16、1:1（仅reference2video、text2video支持）

	// 首尾帧专用 (start-end2video) - 使用images数组，第一张为首帧，第二张为尾帧

	// 文生视频专用 (text2video)
	Style string `json:"style,omitempty"` // 风格：general、anime（仅text2video支持）
}

// 参考生图请求结构体 (reference2image)
type ViduReference2ImageRequest struct {
	Model       string   `json:"model"`                    // 模型名：viduq1
	Images      []string `json:"images"`                  // 图像参考，支持1-7张图片
	Prompt      string   `json:"prompt"`                  // 文本提示词，最长2000字符
	Seed        *int     `json:"seed,omitempty"`          // 随机种子
	AspectRatio string   `json:"aspect_ratio,omitempty"` // 比例：16:9、9:16、1:1、auto
	Payload     string   `json:"payload,omitempty"`       // 透传参数
	CallbackURL string   `json:"callback_url,omitempty"`  // 回调URL
}

// 取消任务请求结构体
type ViduCancelRequest struct {
	ID string `json:"id"` // 任务ID
}

// 官方响应结构体 - 完全对齐官方文档
type ViduResponse struct {
	TaskID            string   `json:"task_id"`                      // 任务ID
	State             string   `json:"state"`                        // 处理状态：created、queueing、processing、success、failed
	Model             string   `json:"model,omitempty"`              // 模型名称
	Prompt            string   `json:"prompt,omitempty"`             // 提示词
	Images            []string `json:"images,omitempty"`             // 图像参数
	Duration          *int     `json:"duration,omitempty"`           // 视频时长
	Seed              *int     `json:"seed,omitempty"`               // 随机种子
	Resolution        string   `json:"resolution,omitempty"`         // 分辨率
	BGM               *bool    `json:"bgm,omitempty"`                // 背景音乐
	MovementAmplitude string   `json:"movement_amplitude,omitempty"` // 运动幅度
	Payload           string   `json:"payload,omitempty"`            // 透传参数
	OffPeak           *bool    `json:"off_peak,omitempty"`           // 错峰模式
	Credits           *int     `json:"credits,omitempty"`            // 消耗积分
	Watermark         *bool    `json:"watermark,omitempty"`          // 水印
	CreatedAt         string   `json:"created_at,omitempty"`         // 创建时间
	// text2video专用
	Style       string `json:"style,omitempty"`        // 风格
	AspectRatio string `json:"aspect_ratio,omitempty"` // 比例
}

// 查询任务响应结构体
type ViduQueryResponse struct {
	ID        string          `json:"id"`        // 任务ID
	State     string          `json:"state"`     // 处理状态
	ErrCode   string          `json:"err_code"`  // 错误代码
	Credits   *int            `json:"credits"`   // 消耗积分
	Payload   string          `json:"payload"`   // 透传参数
	BGM       *bool           `json:"bgm"`       // 背景音乐
	OffPeak   *bool           `json:"off_peak"`  // 错峰模式
	Creations []ViduCreation `json:"creations"` // 生成结果
}

// 生成结果
type ViduCreation struct {
	ID              string `json:"id"`               // 生成物ID
	URL             string `json:"url"`              // 生成物URL，1小时有效期
	CoverURL        string `json:"cover_url"`        // 封面URL，1小时有效期
	WatermarkedURL  string `json:"watermarked_url"`  // 带水印的URL，1小时有效期
}

// 错误响应结构体
type ViduErrorResponse struct {
	Code     int    `json:"code"`
	Reason   string `json:"reason"`
	Message  string `json:"message"`
	Metadata struct {
		TraceID string `json:"trace_id"`
	} `json:"metadata"`
}

// 任务状态常量
const (
	ViduStatusCreated    = "created"
	ViduStatusQueueing   = "queueing"
	ViduStatusProcessing = "processing"
	ViduStatusSuccess    = "success"
	ViduStatusFailed     = "failed"
)

// 接口路径常量
const (
	ViduActionImg2Video        = "img2video"
	ViduActionReference2Video  = "reference2video"
	ViduActionStartEnd2Video   = "start-end2video"
	ViduActionText2Video       = "text2video"
	ViduActionReference2Image  = "reference2image"
	ViduActionCancelTask       = "cancel"
)

// 模型常量
const (
	ViduModelQ2Pro     = "viduq2-pro"
	ViduModelQ2Turbo   = "viduq2-turbo"
	ViduModelQ1        = "viduq1"
	ViduModelQ1Classic = "viduq1-classic"
	ViduModel20        = "vidu2.0"
	ViduModel15        = "vidu1.5"
)
