package relay

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"mime"
	"net/http"
	"one-api/common"
	"one-api/common/config"
	"one-api/common/logger"
	"one-api/common/storage"
	"one-api/common/utils"
	"one-api/common/video"
	"one-api/model"
	providersBase "one-api/providers/base"
	"one-api/relay/relay_util"
	"one-api/types"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

// filepath.Ext 的轻量封装，避免额外导入整个路径包
func filepathExtSafe(name string) string {
	// 找最后一个点
	for i := len(name) - 1; i >= 0; i-- {
		if name[i] == '.' {
			return name[i:]
		}
		// 遇到路径分隔符则停止
		if name[i] == '/' || name[i] == '\\' {
			break
		}
	}
	return ""
}

type soraTaskProperties struct {
	Model        string `json:"model"`
	Resolution   string `json:"resolution"`
	Seconds      int    `json:"seconds"`
	Orientation  string `json:"orientation"`
	SizeLabel    string `json:"size_label"`
	BillingModel string `json:"billing_model"`
	Prompt       string `json:"prompt,omitempty"`
}

type soraVideoResponse struct {
	ID                 string          `json:"id"`
	Object             string          `json:"object"`
	CreatedAt          int64           `json:"created_at,omitempty"`
	CompletedAt        int64           `json:"completed_at,omitempty"`
	ExpiresAt          int64           `json:"expires_at,omitempty"`
	Status             string          `json:"status"`
	Model              string          `json:"model,omitempty"`
	Prompt             string          `json:"prompt,omitempty"`
	Progress           int             `json:"progress"`
	Seconds            string          `json:"seconds,omitempty"`
	Size               string          `json:"size,omitempty"`
	RemixedFromVideoID string          `json:"remixed_from_video_id,omitempty"`
	Error              *soraVideoError `json:"error,omitempty"`
	VideoURL           string          `json:"video_url,omitempty"`
}

type soraVideoError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type soraVideoList struct {
	Object string               `json:"object"`
	Data   []*soraVideoResponse `json:"data"`
}

func VideoCreate(c *gin.Context) {
	var req types.VideoCreateRequest
	contentType := strings.ToLower(strings.TrimSpace(c.Request.Header.Get("Content-Type")))
	// 当为 multipart/form-data 且包含文件字段时，避免使用 ShouldBind 解析（会因文件字段类型报错）。
	if strings.Contains(contentType, "multipart/form-data") {
		// 读取原始请求体并缓存到上下文，供上游直透使用
		raw, err := io.ReadAll(c.Request.Body)
		if err != nil {
			common.AbortWithMessage(c, http.StatusBadRequest, err.Error())
			return
		}
		_ = c.Request.Body.Close()
		c.Set(config.GinRequestBodyKey, raw)

		// 复位 Body 以便后续 ParseMultipartForm 读取
		c.Request.Body = io.NopCloser(bytes.NewReader(raw))
		// 解析表单文本字段（不触碰文件）
		// 使用较小内存阈值，文件会落到临时目录（我们不访问文件内容）
		if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
			common.AbortWithMessage(c, http.StatusBadRequest, err.Error())
			return
		}
		form := c.Request.MultipartForm
		collectedRefs := []string{}
		if form != nil && form.Value != nil {
			if v := strings.TrimSpace(firstOrEmpty(form.Value["model"])); v != "" {
				req.Model = v
			}
			if v := strings.TrimSpace(firstOrEmpty(form.Value["prompt"])); v != "" {
				req.Prompt = v
			}
			if v := strings.TrimSpace(firstOrEmpty(form.Value["seconds"])); v != "" {
				if iv, e := strconv.Atoi(v); e == nil {
					req.Seconds = iv
				}
			}
			if v := strings.TrimSpace(firstOrEmpty(form.Value["size"])); v != "" {
				req.Size = v
			}
			// 支持以文本字段传递 URL/DataURI 参考图
			if vals, ok := form.Value["input_reference"]; ok {
				for _, s := range vals {
					s = strings.TrimSpace(s)
					if s != "" {
						collectedRefs = append(collectedRefs, s)
					}
				}
			}
			if vals, ok := form.Value["input_images"]; ok {
				for _, s := range vals {
					s = strings.TrimSpace(s)
					if s != "" {
						collectedRefs = append(collectedRefs, s)
					}
				}
			}
			if v := strings.TrimSpace(firstOrEmpty(form.Value["input_image"])); v != "" {
				collectedRefs = append(collectedRefs, v)
			}
		}
		// 解析文件字段：input_reference（可多次）
		if form != nil && form.File != nil {
			if fhs, ok := form.File["input_reference"]; ok {
				const maxImageSize = 10 * 1024 * 1024
				for _, fh := range fhs {
					if fh == nil {
						continue
					}
					// 打开并限制大小
					f, err := fh.Open()
					if err != nil {
						continue
					}
					b, _ := io.ReadAll(f)
					_ = f.Close()
					if len(b) == 0 || len(b) > maxImageSize {
						continue
					}
					// 识别扩展名
					ext := ".jpg"
					if fh.Filename != "" {
						if mt := mime.TypeByExtension(strings.ToLower(filepathExtSafe(fh.Filename))); mt != "" {
							// 保持 jpg/png/webp 三类
							if strings.Contains(mt, "jpeg") {
								ext = ".jpg"
							}
							if strings.Contains(mt, "png") {
								ext = ".png"
							}
							if strings.Contains(mt, "webp") {
								ext = ".webp"
							}
						}
					}
					key := utils.GetUUID() + ext
					if url := storage.Upload(b, key); strings.TrimSpace(url) != "" {
						collectedRefs = append(collectedRefs, url)
					}
				}
			}
		}
		if len(collectedRefs) > 0 {
			if req.InputReference == nil {
				req.InputReference = []string{}
			}
			req.InputReference = append(req.InputReference, collectedRefs...)
		}
		// 调试输出：记录解析到的关键字段（不输出文件内容）
		logger.LogDebug(c.Request.Context(), fmt.Sprintf("video.create.multipart parsed -> model=%s seconds=%d size=%s refs=%d ct=%s", req.Model, req.Seconds, req.Size, len(collectedRefs), c.Request.Header.Get("Content-Type")))
		// 再次复位 Body，供后续上游直透读取原始数据
		c.Request.Body = io.NopCloser(bytes.NewReader(raw))
	} else {
		if err := common.UnmarshalBodyReusable(c, &req); err != nil {
			common.AbortWithMessage(c, http.StatusBadRequest, err.Error())
			return
		}
	}

    originalModel := strings.TrimSpace(req.Model)
    if originalModel == "" {
        // 兼容旧行为：缺省走 OpenAI Sora
        originalModel = "sora-2"
    }

    // 渠道内部适配：对外统一接受 Sutui 风格的 veo3*，在内部路由到官方 Veo 两个可用模型：
    // - veo3                -> veo-3.1-fast-generate-preview
    // - veo3.1              -> veo-3.1-generate-preview
    // - veo3-pro            -> veo-3.1-generate-preview
    // - veo3.1-pro          -> veo-3.1-generate-preview
    // - veo3.1-components   -> veo-3.1-generate-preview（多图参考，强制 16:9）
    selectModel := originalModel
    lowerOrig := strings.ToLower(originalModel)
    switch lowerOrig {
    case "veo3":
        selectModel = "veo-3.1-fast-generate-preview"  // 快速首帧 → fast
    case "veo3.1":
        selectModel = "veo-3.1-fast-generate-preview"  // 快速首尾帧 → fast
    case "veo3-pro", "veo3.1-pro":
        selectModel = "veo-3.1-generate-preview"       // 高质量 → standard
    case "veo3.1-components":
        selectModel = "veo-3.1-generate-preview"       // 多图参考 → standard（并强制 16:9）
    }

    // veo3.1-components 要求 16:9，若未提供或提供为竖版，强制覆盖为 1600x900
    if lowerOrig == "veo3.1-components" {
        s := strings.ReplaceAll(strings.ToLower(strings.TrimSpace(req.Size)), " ", "")
        mustFix := (s == "")
        if !mustFix && strings.Contains(s, "x") {
            parts := strings.Split(s, "x")
            if len(parts) == 2 {
                w, _ := strconv.Atoi(parts[0])
                h, _ := strconv.Atoi(parts[1])
                if h > w { // 竖版，强制改横版 16:9
                    mustFix = true
                }
            }
        }
        if mustFix {
            req.Size = "1600x900"
        }
    }

    // 模型白名单扩展：允许 Sora 与 Veo 族（veo-* 与 sutui 的 veo3*）；内部路由后 selectModel 可能已变更为 veo-3.1-*
    // 其余模型仍拒绝，以避免误选其它视频供应商。
    if m := strings.ToLower(originalModel); !(m == "sora-2" || m == "sora-2-pro" || strings.HasPrefix(strings.ToLower(selectModel), "veo-") || strings.HasPrefix(m, "veo3")) {
        common.AbortWithMessage(c, http.StatusBadRequest, "invalid model: only 'sora-2', 'sora-2-pro', 'veo-*' or 'veo3*' are allowed")
        return
    }

	normalize := normalizeSoraSizeInfo(req.Size)
	if req.Seconds <= 0 {
		req.Seconds = defaultSoraDuration()
	}

    // 用内部路由后的 selectModel 选择渠道，以便仅在渠道侧维护官方 Veo 名称
	provider, mappedModel, err := GetProvider(c, selectModel)
	if err != nil {
		common.AbortWithMessage(c, http.StatusServiceUnavailable, err.Error())
		return
	}
	if mappedModel == "" {
		mappedModel = selectModel
	}
	req.Model = mappedModel

	// 调试日志：记录渠道与模型映射、基础地址
	func() {
		ch := provider.GetChannel()
		base := ch.GetBaseURL()
		logger.LogDebug(c.Request.Context(), fmt.Sprintf(
			"video.debug.select -> channel_id=%d type=%d base=%s original=%s select=%s mapped=%s", 
			ch.Id, ch.Type, base, originalModel, selectModel, mappedModel,
		))
	}()

	videoProvider, ok := provider.(providersBase.VideoInterface)
	if !ok {
		common.AbortWithMessage(c, http.StatusNotImplemented, "video interface not implemented for channel")
		return
	}

	// 针对 NewAPI 渠道：按次计费，不依赖 seconds
	isNewAPI := c.GetInt("channel_type") == config.ChannelTypeNewAPI
	var quota *relay_util.Quota
	var billingModel string
    if isNewAPI {
        billingModel = originalModel
        quota = relay_util.NewQuota(c, billingModel, 0)
    } else {
        // 计费模型：
        // - Sora：继续使用 buildSoraBillingModel（按分辨率带价档）
        // - Veo（veo-*）：直接使用 mappedModel（如 veo-3.1-generate-preview），在 quota 中按“秒×1000”换算
        // - Sutui Veo3（veo3*，按次计费）：直接使用 mappedModel，价格类型为 times，由配额器按次结算
        if strings.HasPrefix(strings.ToLower(mappedModel), "veo-") || strings.HasPrefix(strings.ToLower(mappedModel), "veo3") {
            billingModel = mappedModel
            quota = relay_util.NewQuota(c, billingModel, req.Seconds)
        } else {
            billingModel = buildSoraBillingModel(mappedModel, normalize.Resolution)
            quota = relay_util.NewQuota(c, billingModel, req.Seconds)
        }
    }
	if errWithCode := quota.PreQuotaConsumption(); errWithCode != nil {
		newErr := FilterOpenAIErr(c, errWithCode)
		relayResponseWithOpenAIErr(c, &newErr)
		return
	}

	var job *types.VideoJob
	job, errWithCode := videoProvider.CreateVideo(&req)
	if errWithCode != nil {
		quota.Undo(c)
		newErr := FilterOpenAIErr(c, errWithCode)
		relayResponseWithOpenAIErr(c, &newErr)
		return
	}

	if job.Object == "" {
		job.Object = "video"
	}
	job.Model = originalModel
	if job.Prompt == "" {
		job.Prompt = req.Prompt
	}
	if job.Seconds == 0 {
		job.Seconds = req.Seconds
	}
	if job.Size == "" {
		job.Size = normalize.Resolution
	}

	// 在消费前先落库，得到平台任务ID，用于统一对外返回与日志对照
	upstreamID := job.ID
	props := soraTaskProperties{
		Model:        originalModel,
		Resolution:   normalize.Resolution,
		Seconds:      req.Seconds,
		Orientation:  normalize.Orientation,
		SizeLabel:    normalize.SizeLabel,
		BillingModel: billingModel,
		Prompt:       req.Prompt,
	}
    task := saveSoraTask(c, provider.GetChannel().Id, job, props)
    if task != nil {
        platformID := utils.AddTaskPrefix(task.PlatformTaskID)
        // 覆盖返回体 ID 为平台任务ID
        job.ID = platformID
        // 在日志中记录双ID
        quota.SetTaskIDs(platformID, upstreamID)
    }

	usage := &types.Usage{}
	if !isNewAPI {
		usage.PromptTokens = req.Seconds
		usage.TotalTokens = req.Seconds
	}
	quota.Consume(c, usage, false)
    // 二次保存已移除，避免重复入库

	// 代理视频URL（隐藏真实供应商域名）
	proxyVideoURLs(job)

	c.JSON(http.StatusOK, newSoraVideoResponse(job))
}

func VideoRetrieve(c *gin.Context) {
	videoID := strings.TrimSpace(c.Param("id"))
	if videoID == "" {
		common.AbortWithMessage(c, http.StatusBadRequest, "video id is required")
		return
	}

	// 支持平台任务ID：task_<base36>
    var task *model.Task
    var err error
    if strings.HasPrefix(strings.ToLower(videoID), utils.PlatformTaskPrefix) {
        pid := utils.StripTaskPrefix(videoID)
        if t, _ := model.GetTaskByPlatformTaskID(model.TaskPlatformSora, c.GetInt("id"), pid); t != nil {
            task = t
            videoID = t.TaskID
        }
    } else if id, ok := utils.DecodePlatformTaskID(videoID); ok { // 兼容历史 base36 样式
        if t, _ := model.GetTaskByID(id); t != nil && t.Platform == model.TaskPlatformSora && t.UserId == c.GetInt("id") {
            task = t
            videoID = t.TaskID // 改为上游ID
        }
    } else if utils.IsULID(videoID) { // 兼容未带前缀的 ULID
        if t, _ := model.GetTaskByPlatformTaskID(model.TaskPlatformSora, c.GetInt("id"), videoID); t != nil {
            task = t
            videoID = t.TaskID
        }
    }
	if task == nil {
		task, err = model.GetTaskByTaskId(model.TaskPlatformSora, c.GetInt("id"), videoID)
		if err != nil {
			common.AbortWithMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	var (
		props       soraTaskProperties
		provider    providersBase.ProviderInterface
		mappedModel string
	)

	modelName := ""
	if task != nil {
		props = parseSoraTaskProperties(task.Properties)
		c.Set("specific_channel_id", task.ChannelId)
		modelName = strings.TrimSpace(props.Model)
	}
	if modelName == "" {
		modelName = "sora-2"
	}

	provider, mappedModel, err = GetProvider(c, modelName)
	if err != nil {
		common.AbortWithMessage(c, http.StatusServiceUnavailable, err.Error())
		return
	}
	if mappedModel == "" {
		mappedModel = modelName
	}

	videoProvider, ok := provider.(providersBase.VideoInterface)
	if !ok {
		common.AbortWithMessage(c, http.StatusNotImplemented, "video interface not implemented for channel")
		return
	}

	job, errWithCode := videoProvider.RetrieveVideo(videoID)
	if errWithCode != nil {
		newErr := FilterOpenAIErr(c, errWithCode)
		relayResponseWithOpenAIErr(c, &newErr)
		return
	}

	if job.Object == "" {
		job.Object = "video"
	}
	// 始终优先使用本地任务保存的官方模型名，避免上游内部 SKU 外泄
	if strings.TrimSpace(props.Model) != "" {
		job.Model = props.Model
	} else if strings.TrimSpace(job.Model) == "" {
		if mappedModel != "" {
			job.Model = mappedModel
		} else {
			job.Model = "sora-2"
		}
	}
	if job.Size == "" {
		if props.Resolution != "" {
			job.Size = props.Resolution
		} else {
			job.Size = normalizeSoraSizeInfo("").Resolution
		}
	}
	if job.Seconds == 0 {
		if props.Seconds > 0 {
			job.Seconds = props.Seconds
		} else {
			job.Seconds = defaultSoraDuration()
		}
	}
	if strings.TrimSpace(job.Prompt) == "" && strings.TrimSpace(props.Prompt) != "" {
		job.Prompt = props.Prompt
	}

	if task != nil {
		// 统一返回平台任务ID
		job.ID = utils.EncodePlatformTaskID(task.ID)
		updateSoraTask(task, job)
	}

	// 代理视频URL（隐藏真实供应商域名）
	proxyVideoURLs(job)

	c.JSON(http.StatusOK, newSoraVideoResponse(job))
}

func VideoDownload(c *gin.Context) {
	videoID := strings.TrimSpace(c.Param("id"))
	if videoID == "" {
		common.AbortWithMessage(c, http.StatusBadRequest, "video id is required")
		return
	}

	// 支持平台任务ID：task_<base36>
    var task *model.Task
    var err error
    if strings.HasPrefix(strings.ToLower(videoID), utils.PlatformTaskPrefix) {
        pid := utils.StripTaskPrefix(videoID)
        if t, _ := model.GetTaskByPlatformTaskID(model.TaskPlatformSora, c.GetInt("id"), pid); t != nil {
            task = t
            videoID = t.TaskID
        }
    } else if id, ok := utils.DecodePlatformTaskID(videoID); ok {
        if t, _ := model.GetTaskByID(id); t != nil && t.Platform == model.TaskPlatformSora && t.UserId == c.GetInt("id") {
            task = t
            videoID = t.TaskID
        }
    } else if utils.IsULID(videoID) {
        if t, _ := model.GetTaskByPlatformTaskID(model.TaskPlatformSora, c.GetInt("id"), videoID); t != nil {
            task = t
            videoID = t.TaskID
        }
    }
	if task == nil {
		task, err = model.GetTaskByTaskId(model.TaskPlatformSora, c.GetInt("id"), videoID)
		if err != nil {
			common.AbortWithMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	modelName := ""
	if task != nil {
		props := parseSoraTaskProperties(task.Properties)
		modelName = strings.TrimSpace(props.Model)
		c.Set("specific_channel_id", task.ChannelId)
	}
	if modelName == "" {
		modelName = "sora-2"
	}

	provider, _, err := GetProvider(c, modelName)
	if err != nil {
		common.AbortWithMessage(c, http.StatusServiceUnavailable, err.Error())
		return
	}
	videoProvider, ok := provider.(providersBase.VideoInterface)
	if !ok {
		common.AbortWithMessage(c, http.StatusNotImplemented, "video interface not implemented for channel")
		return
	}

	variant := strings.TrimSpace(c.Query("variant"))
	resp, errWithCode := videoProvider.DownloadVideo(videoID, variant)
	if errWithCode != nil {
		newErr := FilterOpenAIErr(c, errWithCode)
		relayResponseWithOpenAIErr(c, &newErr)
		return
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			c.Writer.Header().Add(key, value)
		}
	}

	c.Status(resp.StatusCode)
	if _, err := io.Copy(c.Writer, resp.Body); err != nil {
		logger.LogError(c.Request.Context(), "copy_video_content_failed:"+err.Error())
	}
}

// VideoRemix 实现 /v1/videos/{id}/remix
func VideoRemix(c *gin.Context) {
	videoID := strings.TrimSpace(c.Param("id"))
	if videoID == "" {
		common.AbortWithMessage(c, http.StatusBadRequest, "video id is required")
		return
	}

	var req types.VideoRemixRequest
	if err := common.UnmarshalBodyReusable(c, &req); err != nil {
		common.AbortWithMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	if strings.TrimSpace(req.Prompt) == "" {
		common.AbortWithMessage(c, http.StatusBadRequest, "prompt is required")
		return
	}

	// 优先根据历史任务锁定渠道与属性
	var props soraTaskProperties
	// 支持平台任务ID
    var task *model.Task
    if strings.HasPrefix(strings.ToLower(videoID), utils.PlatformTaskPrefix) {
        if t, _ := model.GetTaskByPlatformTaskID(model.TaskPlatformSora, c.GetInt("id"), utils.StripTaskPrefix(videoID)); t != nil {
            task = t
            videoID = t.TaskID
        }
    } else if id, ok := utils.DecodePlatformTaskID(videoID); ok {
        if t, _ := model.GetTaskByID(id); t != nil && t.Platform == model.TaskPlatformSora && t.UserId == c.GetInt("id") {
            task = t
            videoID = t.TaskID
        }
    } else if utils.IsULID(videoID) {
        if t, _ := model.GetTaskByPlatformTaskID(model.TaskPlatformSora, c.GetInt("id"), videoID); t != nil {
            task = t
            videoID = t.TaskID
        }
    }
    if task == nil {
        task, _ = model.GetTaskByTaskId(model.TaskPlatformSora, c.GetInt("id"), videoID)
    }
	if task != nil {
		props = parseSoraTaskProperties(task.Properties)
		c.Set("specific_channel_id", task.ChannelId)
	}

	// 选择供应商（若无历史任务，则默认以 sora-2 选择）
	modelName := props.Model
	if strings.TrimSpace(modelName) == "" {
		modelName = "sora-2"
	}

	provider, mappedModel, err := GetProvider(c, modelName)
	if err != nil {
		common.AbortWithMessage(c, http.StatusServiceUnavailable, err.Error())
		return
	}
	if mappedModel == "" {
		mappedModel = modelName
	}

	videoProvider, ok := provider.(providersBase.VideoInterface)
	if !ok {
		common.AbortWithMessage(c, http.StatusNotImplemented, "video interface not implemented for channel")
		return
	}

	// 计费：使用原视频的时长/分辨率作为预估
	seconds := props.Seconds
	if seconds <= 0 {
		seconds = defaultSoraDuration()
	}
	billingModel := buildSoraBillingModel(mappedModel, props.Resolution)
	quota := relay_util.NewQuota(c, billingModel, seconds)
	if errWithCode := quota.PreQuotaConsumption(); errWithCode != nil {
		newErr := FilterOpenAIErr(c, errWithCode)
		relayResponseWithOpenAIErr(c, &newErr)
		return
	}

	job, errWithCode := videoProvider.RemixVideo(videoID, req.Prompt)
	if errWithCode != nil {
		quota.Undo(c)
		newErr := FilterOpenAIErr(c, errWithCode)
		relayResponseWithOpenAIErr(c, &newErr)
		return
	}

	if job.Object == "" {
		job.Object = "video"
	}
	if job.Model == "" {
		job.Model = modelName
	}
	if job.RemixedFromVideoID == "" {
		job.RemixedFromVideoID = videoID
	}
	if job.Quality == "" {
		job.Quality = "standard"
	}
	if job.CreatedAt == 0 {
		job.CreatedAt = time.Now().Unix()
	}

	// 先保存 remix 任务，得到平台ID，再记录日志中的双ID
	upstreamID := job.ID
	// 若历史 props 缺失，构建默认属性保存
	if props.Model == "" {
		normalize := normalizeSoraSizeInfo("")
		props = soraTaskProperties{
			Model: mappedModel, Resolution: normalize.Resolution, Seconds: seconds,
			Orientation: normalize.Orientation, SizeLabel: normalize.SizeLabel,
			BillingModel: billingModel, Prompt: req.Prompt,
		}
	} else if strings.TrimSpace(props.Prompt) == "" {
		props.Prompt = req.Prompt
	}

	if props.Prompt == "" {
		props.Prompt = req.Prompt
	}

	taskSaved := saveSoraTask(c, provider.GetChannel().Id, job, props)
	if taskSaved != nil {
		platformID := utils.EncodePlatformTaskID(taskSaved.ID)
		job.ID = platformID
		quota.SetTaskIDs(platformID, upstreamID)
	}

	usage := &types.Usage{PromptTokens: seconds, TotalTokens: seconds}
	quota.Consume(c, usage, false)

	// 代理视频URL（隐藏真实供应商域名）
	proxyVideoURLs(job)

	c.JSON(http.StatusOK, newSoraVideoResponse(job))
}

// VideoList 实现 /v1/videos（官方上游列表）。
// 注意：官方为组织级别列表。此处采用上游直透策略。
func VideoList(c *gin.Context) {
	after := strings.TrimSpace(c.Query("after"))
	order := strings.TrimSpace(c.Query("order"))
	limitVal := strings.TrimSpace(c.Query("limit"))
	limit := 0
	if limitVal != "" {
		if v, err := strconv.Atoi(limitVal); err == nil {
			limit = v
		}
	}

	// 选择一个 OpenAI 视频渠道（按默认模型 sora-2 选路）
	provider, _, err := GetProvider(c, "sora-2")
	if err != nil {
		common.AbortWithMessage(c, http.StatusServiceUnavailable, err.Error())
		return
	}
	videoProvider, ok := provider.(providersBase.VideoInterface)
	if !ok {
		common.AbortWithMessage(c, http.StatusNotImplemented, "video interface not implemented for channel")
		return
	}

	list, errWithCode := videoProvider.ListVideos(after, limit, order)
	if errWithCode != nil {
		newErr := FilterOpenAIErr(c, errWithCode)
		relayResponseWithOpenAIErr(c, &newErr)
		return
	}
	resp := &soraVideoList{
		Object: "list",
		Data:   []*soraVideoResponse{},
	}
	if list != nil {
		for i := range list.Data {
			job := list.Data[i]
			resp.Data = append(resp.Data, newSoraVideoResponse(&job))
		}
	}
	c.JSON(http.StatusOK, resp)
}

// VideoDelete 实现 DELETE /v1/videos/{id}
func VideoDelete(c *gin.Context) {
	videoID := strings.TrimSpace(c.Param("id"))
	if videoID == "" {
		common.AbortWithMessage(c, http.StatusBadRequest, "video id is required")
		return
	}

	// 优先根据历史任务锁定渠道
	task, _ := model.GetTaskByTaskId(model.TaskPlatformSora, c.GetInt("id"), videoID)
	if task != nil {
		c.Set("specific_channel_id", task.ChannelId)
	}

	provider, _, err := GetProvider(c, "sora-2")
	if err != nil {
		common.AbortWithMessage(c, http.StatusServiceUnavailable, err.Error())
		return
	}
	videoProvider, ok := provider.(providersBase.VideoInterface)
	if !ok {
		common.AbortWithMessage(c, http.StatusNotImplemented, "video interface not implemented for channel")
		return
	}

	job, errWithCode := videoProvider.DeleteVideo(videoID)
	if errWithCode != nil {
		newErr := FilterOpenAIErr(c, errWithCode)
		relayResponseWithOpenAIErr(c, &newErr)
		return
	}
	if job == nil {
		job = &types.VideoJob{ID: videoID, Object: "video"}
	}
	c.JSON(http.StatusOK, newSoraVideoResponse(job))
}

func saveSoraTask(c *gin.Context, channelID int, job *types.VideoJob, props soraTaskProperties) *model.Task {
    if job == nil || job.ID == "" {
        return nil
    }
	if props.Prompt == "" && job.Prompt != "" {
		props.Prompt = job.Prompt
	}

    task := &model.Task{
        PlatformTaskID: utils.NewPlatformULID(),
        TaskID:        job.ID,
        Platform:      model.TaskPlatformSora,
        UserId:        c.GetInt("id"),
        TokenID:       c.GetInt("token_id"),
        ChannelId:     channelID,
		SubmitTime:    submitTimeFromJob(job),
		Status:        mapVideoStatus(job.Status),
		Progress:      int(math.Round(job.Progress)),
		Action:        "video.generate",
		BillingModel:  props.BillingModel,
		VideoDuration: float64(props.Seconds),
	}

	if propsJSON, err := json.Marshal(props); err == nil {
		task.Properties = datatypes.JSON(propsJSON)
	}
	if job.Result != nil {
		if dataJSON, err := json.Marshal(job.Result); err == nil {
			task.Data = datatypes.JSON(dataJSON)
		}
	}

	if job.Status == "completed" || job.Status == "failed" {
		task.FinishTime = time.Now().Unix()
	}

    if err := task.Insert(); err != nil {
        logger.LogError(c.Request.Context(), "save_sora_task_failed:"+err.Error())
        return task
    }
    return task
}

func updateSoraTask(task *model.Task, job *types.VideoJob) {
	if task == nil || job == nil {
		return
	}

	task.Status = mapVideoStatus(job.Status)
	task.Progress = int(math.Round(job.Progress))
	if job.Seconds > 0 {
		task.VideoDuration = float64(job.Seconds)
	}
	if job.Status == "completed" || job.Status == "failed" {
		task.FinishTime = time.Now().Unix()
	}
	if job.Result != nil {
		if dataJSON, err := json.Marshal(job.Result); err == nil {
			task.Data = datatypes.JSON(dataJSON)
		}
	}

	if err := task.Update(); err != nil {
		logger.LogError(context.Background(), "update_sora_task_failed:"+err.Error())
	}
}

func parseSoraTaskProperties(data datatypes.JSON) soraTaskProperties {
	var props soraTaskProperties
	if len(data) == 0 {
		return props
	}
	if err := json.Unmarshal(data, &props); err != nil {
		return props
	}
	return props
}

func mapVideoStatus(status string) model.TaskStatus {
	switch strings.ToLower(status) {
	case "queued":
		return model.TaskStatusQueued
	case "in_progress", "processing":
		return model.TaskStatusInProgress
	case "completed", "success":
		return model.TaskStatusSuccess
	case "failed", "error":
		return model.TaskStatusFailure
	default:
		return model.TaskStatusUnknown
	}
}

func normalizeSoraSizeInfo(size string) soraSizeInfoHelper {
	value := strings.ToLower(strings.TrimSpace(size))
	value = strings.ReplaceAll(value, " ", "")
	if value == "" {
		return soraSizeInfoHelper{
			Resolution:  "720x1280",
			Orientation: "portrait",
			SizeLabel:   "small",
		}
	}

	if strings.Contains(value, "x") {
		parts := strings.Split(value, "x")
		if len(parts) == 2 {
			w, _ := strconv.Atoi(parts[0])
			h, _ := strconv.Atoi(parts[1])
			orientation := "landscape"
			if h > w {
				orientation = "portrait"
			}
			resolution := fmt.Sprintf("%dx%d", w, h)
			return soraSizeInfoHelper{
				Resolution:  resolution,
				Orientation: orientation,
				SizeLabel:   mapResolutionSizeLabel(resolution),
			}
		}
	}

	switch value {
	case "landscape":
		return soraSizeInfoHelper{
			Resolution:  "1280x720",
			Orientation: "landscape",
			SizeLabel:   "small",
		}
	case "portrait":
		return soraSizeInfoHelper{
			Resolution:  "720x1280",
			Orientation: "portrait",
			SizeLabel:   "small",
		}
	default:
		return soraSizeInfoHelper{
			Resolution:  "720x1280",
			Orientation: "portrait",
			SizeLabel:   "small",
		}
	}
}

type soraSizeInfoHelper struct {
	Resolution  string
	Orientation string
	SizeLabel   string
}

func mapResolutionSizeLabel(resolution string) string {
	switch resolution {
	case "1280x720", "720x1280":
		return "small"
	case "1792x1024", "1024x1792":
		return "large"
	default:
		return "small"
	}
}

func buildSoraBillingModel(modelName, resolution string) string {
	modelName = strings.TrimSpace(modelName)
	if resolution == "" {
		return modelName
	}
	// Sutui 上游模型（sora_video2-*）已经在模型名中体现了方向/清晰度/时长等，无需追加分辨率，保持与定价键一致
	lower := strings.ToLower(modelName)
	if strings.HasPrefix(lower, "sora_video2") {
		return modelName
	}
	return fmt.Sprintf("%s-%s", modelName, resolution)
}

func submitTimeFromJob(job *types.VideoJob) int64 {
	if job == nil {
		return time.Now().Unix()
	}
	if job.CreatedAt > 0 {
		return job.CreatedAt
	}
	return time.Now().Unix()
}

func defaultSoraDuration() int {
	// 官方默认 4 秒
	return 4
}

func newSoraVideoResponse(job *types.VideoJob) *soraVideoResponse {
	if job == nil {
		return nil
	}
	resp := &soraVideoResponse{
		ID:                 strings.TrimSpace(job.ID),
		Object:             strings.TrimSpace(job.Object),
		CreatedAt:          job.CreatedAt,
		CompletedAt:        job.CompletedAt,
		ExpiresAt:          job.ExpiresAt,
		Status:             strings.TrimSpace(job.Status),
		Model:              strings.TrimSpace(job.Model),
		Prompt:             job.Prompt,
		RemixedFromVideoID: strings.TrimSpace(job.RemixedFromVideoID),
	}
	if resp.ID == "" {
		resp.ID = job.ID
	}
	if resp.Object == "" {
		resp.Object = "video"
	}
	if resp.Status == "" {
		resp.Status = "queued"
	}
	resp.Progress = clampSoraProgress(int(math.Round(job.Progress)))

	seconds := job.Seconds
	if seconds <= 0 {
		seconds = defaultSoraDuration()
	}
	resp.Seconds = strconv.Itoa(seconds)

	size := strings.TrimSpace(job.Size)
	if size == "" {
		size = normalizeSoraSizeInfo("").Resolution
	}
	resp.Size = size

	resp.Error = convertSoraVideoError(job.Error)
	// 附加 video_url（若上游已返回结果）
	if job.Result != nil && strings.TrimSpace(job.Result.VideoURL) != "" {
		resp.VideoURL = strings.TrimSpace(job.Result.VideoURL)
	}
	return resp
}

func convertSoraVideoError(err *types.VideoJobError) *soraVideoError {
	if err == nil {
		return nil
	}
	code := stringifySoraErrorCode(err.Code)
	if code == "" && strings.TrimSpace(err.Message) == "" {
		return nil
	}
	// 对外隐藏上游供应商标识
	msg := err.Message
	if msg != "" {
		msg = sanitizeProviderMarks(msg)
	}
	return &soraVideoError{
		Code:    code,
		Message: msg,
	}
}

func stringifySoraErrorCode(code any) string {
	switch v := code.(type) {
	case nil:
		return ""
	case string:
		return v
	case fmt.Stringer:
		return v.String()
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	default:
		return fmt.Sprintf("%v", v)
	}
}

func clampSoraProgress(val int) int {
	switch {
	case val < 0:
		return 0
	case val > 100:
		return 100
	default:
		return val
	}
}

func firstOrEmpty(arr []string) string {
	if len(arr) > 0 {
		return arr[0]
	}
	return ""
}

// proxyVideoURLs 将视频结果中的真实URL替换为CF Workers代理URL
func proxyVideoURLs(job *types.VideoJob) {
	if job == nil || job.Result == nil {
		return
	}

	// 代理主视频URL
	if job.Result.VideoURL != "" {
		job.Result.VideoURL = video.ProxyVideoURL(job.Result.VideoURL, job.ID)
	}

	// 代理缩略图URL
	if job.Result.ThumbnailURL != "" {
		job.Result.ThumbnailURL = video.ProxyVideoURL(job.Result.ThumbnailURL, job.ID)
	}

	// 代理下载URL
	if job.Result.DownloadURL != "" {
		job.Result.DownloadURL = video.ProxyVideoURL(job.Result.DownloadURL, job.ID)
	}

	// 代理精灵图URL
	if job.Result.SpriteSheet != "" {
		job.Result.SpriteSheet = video.ProxyVideoURL(job.Result.SpriteSheet, job.ID)
	}
}
