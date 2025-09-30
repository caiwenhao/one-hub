package volcark

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"one-api/common"
	"one-api/common/logger"
	"one-api/model"
	volcarkProvider "one-api/providers/volcark"
)

type videoEstimate struct {
	Width      int
	Height     int
	FPS        float64
	Duration   float64
	Tokens     int64
	Ratio      string
	Resolution string
}

func (e videoEstimate) hasDimensions() bool {
	return e.Width > 0 && e.Height > 0
}

func (e *videoEstimate) ensureTokens() {
	if e.Tokens > 0 {
		return
	}
	if !e.hasDimensions() || e.FPS <= 0 || e.Duration <= 0 {
		return
	}
	value := float64(e.Width) * float64(e.Height) * e.FPS * e.Duration / 1024.0
	if value <= 0 {
		return
	}
	e.Tokens = int64(math.Round(value))
}

func mergeVideoEstimate(base, override videoEstimate) videoEstimate {
	result := base
	if override.Width > 0 {
		result.Width = override.Width
	}
	if override.Height > 0 {
		result.Height = override.Height
	}
	if override.FPS > 0 {
		result.FPS = override.FPS
	}
	if override.Duration > 0 {
		result.Duration = override.Duration
	}
	if override.Tokens > 0 {
		result.Tokens = override.Tokens
	}
	if override.Ratio != "" {
		result.Ratio = override.Ratio
	}
	if override.Resolution != "" {
		result.Resolution = override.Resolution
	}
	result.ensureTokens()
	return result
}

func estimateFromPayload(model string, payload map[string]any) videoEstimate {
	params := make(map[string]any)
	collectVideoParams(payload, params)

	ratio := toStringParam(params, "ratio")
	if ratio == "" {
		ratio = toStringParam(params, "aspect_ratio")
	}
	resolution := toStringParam(params, "resolution")

	width := toIntParam(params, "width", "video_width", "w")
	height := toIntParam(params, "height", "video_height", "h")
	duration := toFloatParam(params, "duration", "video_duration", "clip_duration", "duration_sec")
	fps := toFloatParam(params, "frames_per_second", "frame_rate", "fps")

	if width == 0 || height == 0 {
		width, height = lookupVolcArkDimensions(model, ratio, resolution, toStringParam(params, "size"))
	}
	if fps <= 0 {
		fps = defaultVolcArkFPS(model)
	}
	if duration <= 0 {
		duration = defaultVolcArkDuration(model)
	}

	estimate := videoEstimate{
		Width:      width,
		Height:     height,
		FPS:        fps,
		Duration:   duration,
		Ratio:      ratio,
		Resolution: resolution,
	}
	estimate.ensureTokens()

	return estimate
}

func estimateFromResponse(model string, resp *volcarkProvider.VolcArkVideoTask) videoEstimate {
	if resp == nil {
		return videoEstimate{}
	}

	estimate := videoEstimate{
		Ratio:      strings.TrimSpace(resp.Ratio),
		Resolution: strings.TrimSpace(resp.Resolution),
	}

	if resp.Duration != nil {
		estimate.Duration = float64(*resp.Duration)
	}
	if resp.FramesPerSecond != nil {
		estimate.FPS = float64(*resp.FramesPerSecond)
	}
	if estimate.FPS <= 0 {
		estimate.FPS = defaultVolcArkFPS(model)
	}
	if estimate.Duration <= 0 {
		estimate.Duration = defaultVolcArkDuration(model)
	}

	if estimate.Width == 0 || estimate.Height == 0 {
		estimate.Width, estimate.Height = lookupVolcArkDimensions(model, estimate.Ratio, estimate.Resolution, "")
	}

	if resp.Usage != nil {
		if resp.Usage.VideoTokens > 0 {
			estimate.Tokens = resp.Usage.VideoTokens
		} else if resp.Usage.TotalTokens > 0 {
			estimate.Tokens = resp.Usage.TotalTokens
		}
	}

	estimate.ensureTokens()
	return estimate
}

func collectVideoParams(input any, params map[string]any) {
	switch value := input.(type) {
	case map[string]any:
		for key, val := range value {
			lower := strings.ToLower(key)
			if _, exists := params[lower]; !exists {
				params[lower] = val
			}
			collectVideoParams(val, params)
		}
	case []any:
		for _, item := range value {
			collectVideoParams(item, params)
		}
	}
}

func toStringParam(params map[string]any, keys ...string) string {
	for _, key := range keys {
		if raw, ok := params[strings.ToLower(key)]; ok {
			switch v := raw.(type) {
			case string:
				if strings.TrimSpace(v) != "" {
					return strings.TrimSpace(v)
				}
			case fmt.Stringer:
				str := strings.TrimSpace(v.String())
				if str != "" {
					return str
				}
			case float64:
				return strconv.FormatFloat(v, 'f', -1, 64)
			case int:
				return strconv.Itoa(v)
			case int64:
				return strconv.FormatInt(v, 10)
			}
		}
	}
	return ""
}

func toIntParam(params map[string]any, keys ...string) int {
	for _, key := range keys {
		if raw, ok := params[strings.ToLower(key)]; ok {
			switch v := raw.(type) {
			case int:
				return v
			case int64:
				return int(v)
			case float64:
				return int(math.Round(v))
			case string:
				if val, err := strconv.Atoi(strings.TrimSpace(v)); err == nil {
					return val
				}
			}
		}
	}
	return 0
}

func toFloatParam(params map[string]any, keys ...string) float64 {
	for _, key := range keys {
		if raw, ok := params[strings.ToLower(key)]; ok {
			switch v := raw.(type) {
			case float32:
				if v > 0 {
					return float64(v)
				}
			case float64:
				if v > 0 {
					return v
				}
			case int:
				if v > 0 {
					return float64(v)
				}
			case int64:
				if v > 0 {
					return float64(v)
				}
			case string:
				if value, err := strconv.ParseFloat(strings.TrimSpace(v), 64); err == nil && value > 0 {
					return value
				}
			}
		}
	}
	return 0
}

func lookupVolcArkDimensions(model, ratio, resolution, rawSize string) (int, int) {
	if width, height, ok := parseSizeString(rawSize); ok {
		return width, height
	}

	modelKey := strings.ToLower(model)
	ratioKey := strings.ToLower(ratio)
	resolutionKey := strings.ToLower(resolution)

	if width, height := findDimensionsInMap(volcArkModelDimensions[modelKey], ratioKey, resolutionKey); width > 0 {
		return width, height
	}
	if width, height := findDimensionsInMap(volcArkResolutionDimensions, ratioKey, resolutionKey); width > 0 {
		return width, height
	}
	return 0, 0
}

func findDimensionsInMap(source map[string]map[string][2]int, ratioKey, resolutionKey string) (int, int) {
	if source == nil {
		return 0, 0
	}
	if resolutionKey == "" {
		return 0, 0
	}
	if ratioKey == "" {
		return 0, 0
	}
	if ratioMap, ok := source[resolutionKey]; ok {
		if dims, ok := ratioMap[ratioKey]; ok {
			return dims[0], dims[1]
		}
	}
	return 0, 0
}

func parseSizeString(raw string) (int, int, bool) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return 0, 0, false
	}
	parts := strings.Split(raw, "x")
	if len(parts) != 2 {
		parts = strings.Split(raw, "×")
	}
	if len(parts) != 2 {
		return 0, 0, false
	}
	width, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
	height, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err1 != nil || err2 != nil {
		return 0, 0, false
	}
	return width, height, true
}

func defaultVolcArkFPS(model string) float64 {
	switch strings.ToLower(model) {
	case "wan2.1-14b":
		return 16
	}
	return 24
}

func defaultVolcArkDuration(model string) float64 {
	// 默认使用 5 秒，除非后续在响应中给出更准确的时长
	return 5
}

var volcArkResolutionDimensions = map[string]map[string][2]int{
	"480p": {
		"16:9": {864, 480},
		"4:3":  {736, 544},
		"1:1":  {640, 640},
		"21:9": {960, 416},
	},
	"720p": {
		"16:9": {1248, 704},
		"4:3":  {1120, 832},
		"1:1":  {960, 960},
		"21:9": {1504, 640},
	},
	"1080p": {
		"16:9": {1920, 1088},
		"4:3":  {1664, 1248},
		"1:1":  {1440, 1440},
		"21:9": {2176, 928},
	},
}

var volcArkModelDimensions = map[string]map[string]map[string][2]int{
	"doubao-seaweed": {
		"480p": {
			"1:1":  {480, 480},
			"4:3":  {640, 480},
			"16:9": {848, 480},
		},
		"720p": {
			"1:1":  {720, 720},
			"4:3":  {960, 720},
			"16:9": {1280, 720},
		},
	},
	"wan2.1-14b": {
		"480p": {
			"16:9": {832, 480},
		},
		"720p": {
			"16:9": {1280, 720},
		},
	},
}

func finalizeVolcArkBilling(task *model.Task, resp *volcarkProvider.VolcArkVideoTask) {
	if task == nil {
		return
	}

	requestEstimate := videoEstimate{
		Width:      task.VideoWidth,
		Height:     task.VideoHeight,
		FPS:        task.VideoFps,
		Duration:   task.VideoDuration,
		Tokens:     task.VideoEstimatedTokens,
		Ratio:      "",
		Resolution: "",
	}
	requestEstimate.ensureTokens()

	responseEstimate := estimateFromResponse(task.BillingModel, resp)
	merged := mergeVideoEstimate(requestEstimate, responseEstimate)

	if merged.Duration > 0 {
		task.VideoDuration = merged.Duration
	}
	if merged.FPS > 0 {
		task.VideoFps = merged.FPS
	}
	if merged.Width > 0 {
		task.VideoWidth = merged.Width
	}
	if merged.Height > 0 {
		task.VideoHeight = merged.Height
	}

	usageTokens := merged.Tokens
	if usageTokens == 0 && resp != nil && resp.Usage != nil {
		usageTokens = resp.Usage.TotalTokens
	}
	if usageTokens == 0 {
		usageTokens = task.VideoEstimatedTokens
	}

	if usageTokens <= 0 {
		return
	}

	task.VideoUsageTokens = usageTokens
	if task.VideoEstimatedTokens == 0 {
		task.VideoEstimatedTokens = usageTokens
	}

	billingModel := task.BillingModel
	if billingModel == "" && resp != nil {
		billingModel = resp.Model
	}
	if billingModel == "" {
		return
	}
	task.BillingModel = billingModel

	price := model.PricingInstance.GetPrice(billingModel)
	if price == nil {
		return
	}
	inputRatio := price.GetInput()
	if inputRatio <= 0 {
		return
	}

	groupRatio := task.BillingGroupRatio
	if groupRatio == 0 {
		groupRatio = 1
	}

	expectedQuota := int(math.Ceil((float64(usageTokens) / 1000.0) * inputRatio * groupRatio))
	if expectedQuota <= 0 {
		return
	}

	delta := expectedQuota - task.Quota
	if delta != 0 {
		adjustVolcArkQuota(task, delta)
		message := fmt.Sprintf("火山方舟任务 %s 用量调整：%s，视频tokens=%d", task.TaskID, common.LogQuota(delta), usageTokens)
		model.RecordLog(task.UserId, model.LogTypeSystem, message)
	}

	task.Quota = expectedQuota
}

func adjustVolcArkQuota(task *model.Task, delta int) {
	if delta == 0 {
		return
	}

	if task.TokenID > 0 {
		if err := model.PostConsumeTokenQuota(task.TokenID, delta); err != nil {
			logger.SysError(fmt.Sprintf("adjust volc ark quota failed: %v", err))
			if delta > 0 {
				if err := model.DecreaseUserQuota(task.UserId, delta); err != nil {
					logger.SysError(fmt.Sprintf("decrease user quota fallback failed: %v", err))
				}
			} else {
				if err := model.IncreaseUserQuota(task.UserId, -delta); err != nil {
					logger.SysError(fmt.Sprintf("increase user quota fallback failed: %v", err))
				}
			}
		}
	} else {
		if delta > 0 {
			if err := model.DecreaseUserQuota(task.UserId, delta); err != nil {
				logger.SysError(fmt.Sprintf("decrease user quota failed: %v", err))
			}
		} else {
			if err := model.IncreaseUserQuota(task.UserId, -delta); err != nil {
				logger.SysError(fmt.Sprintf("increase user quota failed: %v", err))
			}
		}
	}

	model.UpdateChannelUsedQuota(task.ChannelId, delta)
	model.AdjustUserUsedQuota(task.UserId, delta)
}
