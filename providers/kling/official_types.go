package kling

// 官方API兼容的数据类型定义

// ImageListItem 图片列表项
type ImageListItem struct {
	Image string `json:"image"` // 图片URL或Base64编码
}

// Point 坐标点
type Point struct {
	X float64 `json:"x"` // x坐标 [0,1]
	Y float64 `json:"y"` // y坐标 [0,1]
}

// RLEMask RLE编码掩码
type RLEMask struct {
	Size   [2]int `json:"size"`   // [height, width]
	Counts string `json:"counts"` // RLE编码字符串
}

// PNGMask PNG编码掩码
type PNGMask struct {
	Size   [2]int `json:"size"`   // [height, width]
	Base64 string `json:"base64"` // Base64编码的PNG图片
}

// MaskObject 掩码对象
type MaskObject struct {
	ObjectID int      `json:"object_id"` // 对象ID
	RLEMask  *RLEMask `json:"rle_mask"`  // RLE掩码
	PNGMask  *PNGMask `json:"png_mask"`  // PNG掩码
}

// TrajectoryPoint 运动轨迹坐标点
type TrajectoryPoint struct {
	X int `json:"x"` // 轨迹点横坐标（以输入图片左下为原点的像素坐标）
	Y int `json:"y"` // 轨迹点纵坐标（以输入图片左下为原点的像素坐标）
}

// DynamicMask 动态笔刷配置
type DynamicMask struct {
	Mask         string            `json:"mask,omitempty"`         // 动态笔刷涂抹区域
	Trajectories []TrajectoryPoint `json:"trajectories,omitempty"` // 运动轨迹坐标序列
}

// CameraControlConfig 摄像机控制配置
type CameraControlConfig struct {
	Horizontal *float64 `json:"horizontal,omitempty"` // 水平运镜 [-10, 10]
	Vertical   *float64 `json:"vertical,omitempty"`   // 垂直运镜 [-10, 10]
	Pan        *float64 `json:"pan,omitempty"`        // 水平摇镜 [-10, 10]
	Tilt       *float64 `json:"tilt,omitempty"`       // 垂直摇镜 [-10, 10]
	Roll       *float64 `json:"roll,omitempty"`       // 旋转运镜 [-10, 10]
	Zoom       *float64 `json:"zoom,omitempty"`       // 变焦 [-10, 10]
}

// CameraControl 摄像机运动控制
type CameraControl struct {
	Type   string               `json:"type,omitempty"`   // "simple", "down_back", "forward_up", "right_turn_forward", "left_turn_forward"
	Config *CameraControlConfig `json:"config,omitempty"` // 当type为"simple"时必填
}

// OfficialText2VideoRequest 官方文生视频请求格式
type OfficialText2VideoRequest struct {
	ModelName      string         `json:"model_name,omitempty"`       // 模型名称，默认kling-v1
	Prompt         string         `json:"prompt"`                     // 正向文本提示词，必须
	NegativePrompt string         `json:"negative_prompt,omitempty"`  // 负向文本提示词
	CfgScale       *float64       `json:"cfg_scale,omitempty"`        // 生成视频的自由度 [0, 1]
	Mode           string         `json:"mode,omitempty"`             // 生成模式：std, pro
	CameraControl  *CameraControl `json:"camera_control,omitempty"`   // 摄像机运动控制
	AspectRatio    string         `json:"aspect_ratio,omitempty"`     // 画面纵横比：16:9, 9:16, 1:1
	Duration       string         `json:"duration,omitempty"`         // 视频时长：5, 10
	CallbackURL    string         `json:"callback_url,omitempty"`     // 回调通知地址
	ExternalTaskID string         `json:"external_task_id,omitempty"` // 自定义任务ID
}

// OfficialMultiImage2VideoRequest 官方多图参考生视频请求格式
type OfficialMultiImage2VideoRequest struct {
	ModelName      string          `json:"model_name,omitempty"`       // 模型名称，默认kling-v1-6
	ImageList      []ImageListItem `json:"image_list"`                 // 图片列表，最多4张
	Prompt         string          `json:"prompt"`                     // 正向文本提示词，必须
	NegativePrompt string          `json:"negative_prompt,omitempty"`  // 负向文本提示词
	Mode           string          `json:"mode,omitempty"`             // 生成模式：std, pro
	Duration       string          `json:"duration,omitempty"`         // 视频时长：5, 10
	AspectRatio    string          `json:"aspect_ratio,omitempty"`     // 画面纵横比：16:9, 9:16, 1:1
	CallbackURL    string          `json:"callback_url,omitempty"`     // 回调通知地址
	ExternalTaskID string          `json:"external_task_id,omitempty"` // 自定义任务ID
}

// OfficialImage2VideoRequest 官方图生视频请求格式
type OfficialImage2VideoRequest struct {
	ModelName      string         `json:"model_name,omitempty"`       // 模型名称，默认kling-v1
	Image          string         `json:"image"`                      // 图片URL，必须
	ImageTail      string         `json:"image_tail,omitempty"`       // 尾帧图片
	Prompt         string         `json:"prompt,omitempty"`           // 文本提示词
	NegativePrompt string         `json:"negative_prompt,omitempty"`  // 负向文本提示词
	CfgScale       *float64       `json:"cfg_scale,omitempty"`        // 生成视频的自由度 [0, 1]
	Mode           string         `json:"mode,omitempty"`             // 生成模式：std, pro
	StaticMask     string         `json:"static_mask,omitempty"`      // 静态笔刷涂抹区域
	DynamicMasks   []DynamicMask  `json:"dynamic_masks,omitempty"`    // 动态笔刷配置列表（最多6组）
	CameraControl  *CameraControl `json:"camera_control,omitempty"`   // 摄像机运动控制
	AspectRatio    string         `json:"aspect_ratio,omitempty"`     // 画面纵横比：16:9, 9:16, 1:1
	Duration       string         `json:"duration,omitempty"`         // 视频时长：5, 10
	CallbackURL    string         `json:"callback_url,omitempty"`     // 回调通知地址
	ExternalTaskID string         `json:"external_task_id,omitempty"` // 自定义任务ID
}

// OfficialTaskInfo 任务信息
type OfficialTaskInfo struct {
	ExternalTaskID string `json:"external_task_id,omitempty"` // 客户自定义任务ID
}

// OfficialVideoResult 视频结果
type OfficialVideoResult struct {
	ID       string `json:"id"`       // 生成的视频ID
	URL      string `json:"url"`      // 生成视频的URL
	Duration string `json:"duration"` // 视频总时长，单位s
}

// OfficialTaskResult 任务结果
type OfficialTaskResult struct {
	Videos []OfficialVideoResult `json:"videos"`
}

// OfficialTaskData 任务数据
type OfficialTaskData struct {
	TaskID        string              `json:"task_id"`                   // 任务ID，系统生成
	TaskInfo      *OfficialTaskInfo   `json:"task_info,omitempty"`       // 任务创建时的参数信息
	TaskStatus    string              `json:"task_status"`               // 任务状态
	TaskStatusMsg string              `json:"task_status_msg,omitempty"` // 任务状态信息
	TaskResult    *OfficialTaskResult `json:"task_result,omitempty"`     // 任务结果
	CreatedAt     int64               `json:"created_at"`                // 任务创建时间，Unix时间戳、单位ms
	UpdatedAt     int64               `json:"updated_at"`                // 任务更新时间，Unix时间戳、单位ms
}

// OfficialResponse 官方API响应格式
type OfficialResponse struct {
	Code      int                 `json:"code"`                // 错误码
	Message   string              `json:"message"`             // 错误信息
	RequestID string              `json:"request_id"`          // 请求ID
	Data      *OfficialTaskData   `json:"data,omitempty"`      // 单个任务数据
	DataList  []*OfficialTaskData `json:"data_list,omitempty"` // 任务列表数据（用于列表查询）
}

// OfficialTaskListQuery 任务列表查询参数
type OfficialTaskListQuery struct {
	PageNum  int `form:"pageNum" json:"pageNum" binding:"min=1,max=1000"`  // 页码 [1,1000]
	PageSize int `form:"pageSize" json:"pageSize" binding:"min=1,max=500"` // 每页数据量 [1,500]
}

// 模型名称枚举
const (
	ModelKlingV1        = "kling-v1"
	ModelKlingV15       = "kling-v1-5"
	ModelKlingV16       = "kling-v1-6"
	ModelKlingV2Master  = "kling-v2-master"
	ModelKlingV21       = "kling-v2-1"
	ModelKlingV21Master = "kling-v2-1-master"
)

// 生成模式枚举
const (
	ModeStd = "std" // 标准模式
	ModePro = "pro" // 专家模式
)

// 画面纵横比枚举
const (
	AspectRatio16x9 = "16:9"
	AspectRatio9x16 = "9:16"
	AspectRatio1x1  = "1:1"
)

// 视频时长枚举
const (
	Duration5s  = "5"
	Duration10s = "10"
)

// 任务状态枚举
const (
	TaskStatusSubmitted  = "submitted"  // 已提交
	TaskStatusProcessing = "processing" // 处理中
	TaskStatusSucceed    = "succeed"    // 成功
	TaskStatusFailed     = "failed"     // 失败
)

// 摄像机控制类型枚举
const (
	CameraTypeSimple           = "simple"
	CameraTypeDownBack         = "down_back"
	CameraTypeForwardUp        = "forward_up"
	CameraTypeRightTurnForward = "right_turn_forward"
	CameraTypeLeftTurnForward  = "left_turn_forward"
)

// 多模态视频编辑相关类型

// OfficialInitSelectionRequest 初始化待编辑视频请求
type OfficialInitSelectionRequest struct {
	VideoID  string `json:"video_id,omitempty"`  // 视频ID，与video_url二选一
	VideoURL string `json:"video_url,omitempty"` // 视频URL，与video_id二选一
}

// OfficialInitSelectionResponse 初始化待编辑视频响应
type OfficialInitSelectionResponse struct {
	Status           int     `json:"status"`            // 拒识码，非0为识别失败
	SessionID        string  `json:"session_id"`        // 会话ID，24小时有效
	FPS              float64 `json:"fps"`               // 解析后视频的帧数
	OriginalDuration int     `json:"original_duration"` // 解析后视频的时长
	Width            int     `json:"width"`             // 解析后视频的宽
	Height           int     `json:"height"`            // 解析后视频的高
	TotalFrame       int     `json:"total_frame"`       // 解析后视频的总帧数
	NormalizedVideo  string  `json:"normalized_video"`  // 初始化后的视频URL
}

// OfficialAddSelectionRequest 增加视频选区请求
type OfficialAddSelectionRequest struct {
	SessionID  string  `json:"session_id"`  // 会话ID
	FrameIndex int     `json:"frame_index"` // 帧号
	Points     []Point `json:"points"`      // 点选坐标列表
}

// OfficialSelectionResponse 选区操作响应
type OfficialSelectionResponse struct {
	Status    int                  `json:"status"`     // 拒识码
	SessionID string               `json:"session_id"` // 会话ID
	Res       *SelectionResultData `json:"res"`        // 图像分割返回结果
}

// SelectionResultData 选区结果数据
type SelectionResultData struct {
	FrameIndex  int          `json:"frame_index"`   // 帧索引
	RLEMaskList []MaskObject `json:"rle_mask_list"` // 掩码对象列表
}

// OfficialDeleteSelectionRequest 删减视频选区请求
type OfficialDeleteSelectionRequest struct {
	SessionID  string  `json:"session_id"`  // 会话ID
	FrameIndex int     `json:"frame_index"` // 帧号
	Points     []Point `json:"points"`      // 点选坐标列表
}

// OfficialClearSelectionRequest 清除视频选区请求
type OfficialClearSelectionRequest struct {
	SessionID string `json:"session_id"` // 会话ID
}

// OfficialClearSelectionResponse 清除视频选区响应
type OfficialClearSelectionResponse struct {
	Status    int    `json:"status"`     // 拒识码
	SessionID string `json:"session_id"` // 会话ID
}

// OfficialPreviewSelectionRequest 预览已选区视频请求
type OfficialPreviewSelectionRequest struct {
	SessionID string `json:"session_id"` // 会话ID
}

// OfficialPreviewSelectionResponse 预览已选区视频响应
type OfficialPreviewSelectionResponse struct {
	Status    int                `json:"status"`     // 拒识码
	SessionID string             `json:"session_id"` // 会话ID
	Res       *PreviewResultData `json:"res"`        // 预览结果数据
}

// PreviewResultData 预览结果数据
type PreviewResultData struct {
	Video          string `json:"video"`           // 含mask的视频
	VideoCover     string `json:"video_cover"`     // 含mask的视频的封面
	TrackingOutput string `json:"tracking_output"` // 图像分割结果
}

// OfficialMultiElementsRequest 多模态视频编辑任务请求
type OfficialMultiElementsRequest struct {
	ModelName      string          `json:"model_name,omitempty"`       // 模型名称，默认kling-v1-6
	SessionID      string          `json:"session_id"`                 // 会话ID，必须
	EditMode       string          `json:"edit_mode"`                  // 操作类型：addition, swap, removal
	ImageList      []ImageListItem `json:"image_list,omitempty"`       // 参考图像列表
	Prompt         string          `json:"prompt"`                     // 正向文本提示词，必须
	NegativePrompt string          `json:"negative_prompt,omitempty"`  // 负向文本提示词
	Mode           string          `json:"mode,omitempty"`             // 生成模式：std, pro
	Duration       string          `json:"duration,omitempty"`         // 视频时长：5, 10
	CallbackURL    string          `json:"callback_url,omitempty"`     // 回调通知地址
	ExternalTaskID string          `json:"external_task_id,omitempty"` // 自定义任务ID
}

// OfficialMultiElementsTaskData 多模态视频编辑任务数据
type OfficialMultiElementsTaskData struct {
	TaskID        string                   `json:"task_id"`                   // 任务ID
	TaskStatus    string                   `json:"task_status"`               // 任务状态
	TaskStatusMsg string                   `json:"task_status_msg,omitempty"` // 任务状态信息
	TaskInfo      *OfficialTaskInfo        `json:"task_info,omitempty"`       // 任务信息
	TaskResult    *OfficialMultiTaskResult `json:"task_result,omitempty"`     // 任务结果
	SessionID     string                   `json:"session_id,omitempty"`      // 会话ID
	CreatedAt     int64                    `json:"created_at"`                // 创建时间
	UpdatedAt     int64                    `json:"updated_at"`                // 更新时间
}

// OfficialMultiTaskResult 多模态任务结果
type OfficialMultiTaskResult struct {
	Videos []OfficialMultiVideoResult `json:"videos"`
}

// OfficialMultiVideoResult 多模态视频结果
type OfficialMultiVideoResult struct {
	ID        string `json:"id"`         // 视频ID
	SessionID string `json:"session_id"` // 会话ID
	URL       string `json:"url"`        // 视频URL
	Duration  string `json:"duration"`   // 视频时长
}

// 编辑模式枚举
const (
	EditModeAddition = "addition" // 增加元素
	EditModeSwap     = "swap"     // 替换元素
	EditModeRemoval  = "removal"  // 删除元素
)

// 新的模型版本支持
const (
	ModelKlingV16MultiImage = "kling-v1-6"   // 多图参考生视频专用模型
	ModelKlingV2            = "kling-v2"     // 图像生成模型
	ModelKlingV2New         = "kling-v2-new" // 新版图像生成模型
)

// ===================== 图像生成相关类型定义 =====================

// OfficialImageRequest 官方图像生成请求格式
type OfficialImageRequest struct {
	ModelName      string   `json:"model_name,omitempty"`       // 模型名称，默认kling-v1
	Prompt         string   `json:"prompt"`                     // 正向文本提示词，必须
	NegativePrompt string   `json:"negative_prompt,omitempty"`  // 负向文本提示词
	Image          string   `json:"image,omitempty"`            // 参考图像，Base64或URL
	ImageReference string   `json:"image_reference,omitempty"`  // 图片参考类型：subject, face
	ImageFidelity  *float64 `json:"image_fidelity,omitempty"`   // 图片参考强度 [0,1]
	HumanFidelity  *float64 `json:"human_fidelity,omitempty"`   // 面部参考强度 [0,1]
	Resolution     string   `json:"resolution,omitempty"`       // 清晰度：1k, 2k
	N              int      `json:"n,omitempty"`                // 生成图片数量 [1,9]
	AspectRatio    string   `json:"aspect_ratio,omitempty"`     // 纵横比
	CallbackURL    string   `json:"callback_url,omitempty"`     // 回调通知地址
	ExternalTaskID string   `json:"external_task_id,omitempty"` // 自定义任务ID
}

// SubjectImageListItem 多图参考生图中的主体图片项
type SubjectImageListItem struct {
	SubjectImage string `json:"subject_image"` // 主体图片URL或Base64
}

// OfficialMultiImage2ImageRequest 官方多图参考生图请求格式
type OfficialMultiImage2ImageRequest struct {
	ModelName        string                 `json:"model_name,omitempty"`       // 模型名称，默认kling-v2
	Prompt           string                 `json:"prompt,omitempty"`           // 正向文本提示词
	SubjectImageList []SubjectImageListItem `json:"subject_image_list"`         // 主体图片列表，1-4张
	SceneImage       string                 `json:"scene_image,omitempty"`      // 场景参考图
	StyleImage       string                 `json:"style_image,omitempty"`      // 风格参考图
	N                int                    `json:"n,omitempty"`                // 生成图片数量 [1,9]
	AspectRatio      string                 `json:"aspect_ratio,omitempty"`     // 纵横比
	CallbackURL      string                 `json:"callback_url,omitempty"`     // 回调通知地址
	ExternalTaskID   string                 `json:"external_task_id,omitempty"` // 自定义任务ID
}

// OfficialImageResult 图像结果
type OfficialImageResult struct {
	Index int    `json:"index"` // 图片编号 0-9
	URL   string `json:"url"`   // 图片URL
}

// OfficialImageTaskResult 图像任务结果
type OfficialImageTaskResult struct {
	Images []OfficialImageResult `json:"images"`
}

// OfficialImageTaskData 图像任务数据
type OfficialImageTaskData struct {
	TaskID        string                   `json:"task_id"`                   // 任务ID
	TaskInfo      *OfficialTaskInfo        `json:"task_info,omitempty"`       // 任务信息
	TaskStatus    string                   `json:"task_status"`               // 任务状态
	TaskStatusMsg string                   `json:"task_status_msg,omitempty"` // 任务状态信息
	TaskResult    *OfficialImageTaskResult `json:"task_result,omitempty"`     // 任务结果
	CreatedAt     int64                    `json:"created_at"`                // 创建时间
	UpdatedAt     int64                    `json:"updated_at"`                // 更新时间
}

// OfficialImageResponse 图像生成响应格式
type OfficialImageResponse struct {
	Code      int                      `json:"code"`                // 错误码
	Message   string                   `json:"message"`             // 错误信息
	RequestID string                   `json:"request_id"`          // 请求ID
	Data      *OfficialImageTaskData   `json:"data,omitempty"`      // 单个任务数据
	DataList  []*OfficialImageTaskData `json:"data_list,omitempty"` // 任务列表数据
}

// 图像相关枚举常量
const (
	// 图片参考类型
	ImageReferenceSubject = "subject" // 角色特征参考
	ImageReferenceFace    = "face"    // 人物长相参考

	// 清晰度
	Resolution1K = "1k" // 1K标清
	Resolution2K = "2k" // 2K高清

	// 扩展纵横比支持
	AspectRatio4x3  = "4:3"
	AspectRatio3x4  = "3:4"
	AspectRatio3x2  = "3:2"
	AspectRatio2x3  = "2:3"
	AspectRatio21x9 = "21:9"
)
