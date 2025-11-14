package relay_util

import (
	"context"
	"encoding/json"
	"errors"
	"math"
	"net/http"
	"one-api/common"
	"one-api/common/config"
	"one-api/common/logger"
	"one-api/model"
	"one-api/types"
	"time"

	"github.com/gin-gonic/gin"
	"strings"
)

type Quota struct {
	modelName        string
	promptTokens     int
	price            model.Price
	groupName        string
	groupRatio       float64
	inputRatio       float64
	outputRatio      float64
	preConsumedQuota int
	cacheQuota       int
	userId           int
	channelId        int
	tokenId          int
	HandelStatus     bool

	startTime         time.Time
	firstResponseTime time.Time
	extraBillingData  map[string]ExtraBillingData
}

func NewQuota(c *gin.Context, modelName string, promptTokens int) *Quota {
	quota := &Quota{
		modelName:    modelName,
		promptTokens: promptTokens,
		userId:       c.GetInt("id"),
		channelId:    c.GetInt("channel_id"),
		tokenId:      c.GetInt("token_id"),
		HandelStatus: false,
	}

	// 通过客户价 + 模型分组解析最终价格
	groupCode := c.GetString("model_group")
	userID := c.GetInt("id")
	price, resolvedGroup, err := model.ResolveCustomerPrice(userID, quota.modelName, groupCode)
	if err != nil {
		// 解析失败时，回退到旧逻辑，避免因为配置不完整导致服务整体不可用
		// 同时在系统日志中记录，便于后续排查
		logger.SysError("ResolveCustomerPrice failed: " + err.Error())
		price = model.PricingInstance.GetPrice(quota.modelName)
		resolvedGroup = ""
	}
	if price == nil {
		// 理论上不应出现；为安全起见，构造一个零价以避免 panic
		price = &model.Price{Model: quota.modelName}
	}
	_ = resolvedGroup // 预留：后续可写入日志或 context

	quota.price = *price
	quota.groupRatio = c.GetFloat64("group_ratio")
	quota.groupName = c.GetString("token_group")
	quota.inputRatio = quota.price.GetInput() * quota.groupRatio
	quota.outputRatio = quota.price.GetOutput() * quota.groupRatio

	// 动态价档：Gemini 2.5 Pro（含 computer-use 族）大上下文（>200k tokens）按更高档计价
	lowerModel := strings.ToLower(strings.TrimSpace(quota.modelName))
	if (strings.HasPrefix(lowerModel, "gemini-2.5-pro") || strings.HasPrefix(lowerModel, "gemini-2.5-computer-use")) && promptTokens > 200000 {
		// 输入 $2.50/M → 0.0025/1k → 1.25；输出 $15/M → 0.015/1k → 7.5
		quota.inputRatio = 1.25 * quota.groupRatio
		quota.outputRatio = 7.5 * quota.groupRatio
	}

	// 动态模态：Gemini 2.5 Flash（非 lite/image/native-audio）若包含音频输入，输入单价切换至 $1.00/M（ratio=0.5）
	if strings.HasPrefix(lowerModel, "gemini-2.5-flash") &&
		!strings.Contains(lowerModel, "flash-lite") &&
		!strings.Contains(lowerModel, "image") &&
		!strings.Contains(lowerModel, "native-audio") &&
		c.GetBool("gemini_audio_input") {
		quota.inputRatio = 0.5 * quota.groupRatio
	}

	return quota
}

func (q *Quota) PreQuotaConsumption() *types.OpenAIErrorWithStatusCode {
	if q.price.Type == model.TimesPriceType {
		q.preConsumedQuota = int(1000 * q.outputRatio)
	} else if q.price.Input != 0 || q.price.Output != 0 {
		// Sora 按秒计费：预扣时同样按 seconds×1000 计算
		pt := q.promptTokens
		lowerModel := strings.ToLower(strings.TrimSpace(q.modelName))
		if strings.HasPrefix(lowerModel, "sora-2") {
			pt = pt * 1000
		}
		q.preConsumedQuota = int(float64(pt)*q.inputRatio) + config.PreConsumedQuota
	}

	if q.preConsumedQuota == 0 {
		return nil
	}

	userQuota, err := model.CacheGetUserQuota(q.userId)
	if err != nil {
		return common.ErrorWrapper(err, "get_user_quota_failed", http.StatusInternalServerError)
	}

	if userQuota < q.preConsumedQuota {
		return common.ErrorWrapper(errors.New("user quota is not enough"), "insufficient_user_quota", http.StatusPaymentRequired)
	}

	err = model.CacheDecreaseUserQuota(q.userId, q.preConsumedQuota)
	if err != nil {
		return common.ErrorWrapper(err, "decrease_user_quota_failed", http.StatusInternalServerError)
	}

	if userQuota > 100*q.preConsumedQuota {
		// in this case, we do not pre-consume quota
		// because the user has enough quota
		q.preConsumedQuota = 0
		// common.LogInfo(c.Request.Context(), fmt.Sprintf("user %d has enough quota %d, trusted and no need to pre-consume", userId, userQuota))
	}

	if q.preConsumedQuota > 0 {
		err := model.PreConsumeTokenQuota(q.tokenId, q.preConsumedQuota)
		if err != nil {
			return common.ErrorWrapper(err, "pre_consume_token_quota_failed", http.StatusForbidden)
		}
		q.HandelStatus = true
	}

	return nil
}

// 更新用户实时配额
func (q *Quota) UpdateUserRealtimeQuota(usage *types.UsageEvent, nowUsage *types.UsageEvent) error {
	usage.Merge(nowUsage)

	// 不开启Redis，则不更新实时配额
	if !config.RedisEnabled {
		return nil
	}

	promptTokens, completionTokens := q.getComputeTokensByUsageEvent(nowUsage)
	increaseQuota := q.GetTotalQuota(promptTokens, completionTokens, nil)

	cacheQuota, err := model.CacheIncreaseUserRealtimeQuota(q.userId, increaseQuota)
	if err != nil {
		return errors.New("error update user realtime quota cache: " + err.Error())
	}

	q.cacheQuota += increaseQuota
	userQuota, err := model.CacheGetUserQuota(q.userId)
	if err != nil {
		return errors.New("error get user quota cache: " + err.Error())
	}

	if cacheQuota >= int64(userQuota) {
		return errors.New("user quota is not enough")
	}

	return nil
}

func (q *Quota) completedQuotaConsumption(usage *types.Usage, tokenName string, isStream bool, sourceIp string, ctx context.Context) error {
	defer func() {
		if q.cacheQuota > 0 {
			model.CacheDecreaseUserRealtimeQuota(q.userId, q.cacheQuota)
		}
	}()

	quota := q.GetTotalQuotaByUsage(usage)

	if quota > 0 {
		quotaDelta := quota - q.preConsumedQuota
		err := model.PostConsumeTokenQuota(q.tokenId, quotaDelta)
		if err != nil {
			return errors.New("error consuming token remain quota: " + err.Error())
		}
		err = model.CacheUpdateUserQuota(q.userId)
		if err != nil {
			return errors.New("error consuming token remain quota: " + err.Error())
		}
		model.UpdateChannelUsedQuota(q.channelId, quota)
	}

	model.RecordConsumeLog(
		ctx,
		q.userId,
		q.channelId,
		usage.PromptTokens,
		usage.CompletionTokens,
		q.modelName,
		tokenName,
		quota,
		"",
		q.getRequestTime(),
		isStream,
		q.GetLogMeta(usage),
		sourceIp,
	)
	model.UpdateUserUsedQuotaAndRequestCount(q.userId, quota)

	return nil
}

func (q *Quota) Undo(c *gin.Context) {
	tokenId := c.GetInt("token_id")
	if q.HandelStatus {
		go func(ctx context.Context) {
			// return pre-consumed quota
			err := model.PostConsumeTokenQuota(tokenId, -q.preConsumedQuota)
			if err != nil {
				logger.LogError(ctx, "error return pre-consumed quota: "+err.Error())
			}
		}(c.Request.Context())
	}
}

func (q *Quota) Consume(c *gin.Context, usage *types.Usage, isStream bool) {
	tokenName := c.GetString("token_name")
	q.startTime = c.GetTime("requestStartTime")
	// 如果没有报错，则消费配额
	go func(ctx context.Context) {
		err := q.completedQuotaConsumption(usage, tokenName, isStream, c.ClientIP(), ctx)
		if err != nil {
			logger.LogError(ctx, err.Error())
		}
	}(c.Request.Context())
}

func (q *Quota) GetInputRatio() float64 {
	return q.inputRatio
}

func (q *Quota) GetGroupRatio() float64 {
	if q.groupRatio == 0 {
		return 1
	}
	return q.groupRatio
}

func (q *Quota) GetLogMeta(usage *types.Usage) map[string]any {
	meta := map[string]any{
		"group_name":   q.groupName,
		"price_type":   q.price.Type,
		"group_ratio":  q.groupRatio,
		"input_ratio":  q.price.GetInput(),
		"output_ratio": q.price.GetOutput(),
	}

	// 标注 MiniMax 上游来源（official/ppinfra），用于对账与观测
	if q.channelId > 0 {
		if ch, err := model.GetChannelById(q.channelId); err == nil && ch != nil {
			if ch.Type == config.ChannelTypeMiniMax {
				upstream := "official"
				raw := ch.GetCustomParameter()
				if strings.TrimSpace(raw) != "" {
					var payload map[string]json.RawMessage
					if json.Unmarshal([]byte(raw), &payload) == nil {
						if vRaw, ok := payload["audio"]; ok {
							var audio map[string]any
							if json.Unmarshal(vRaw, &audio) == nil {
								if up, ok2 := audio["upstream"].(string); ok2 && strings.TrimSpace(up) != "" {
									upstream = strings.ToLower(strings.TrimSpace(up))
								}
							}
						}
						if upstream == "official" {
							if upRaw, ok := payload["upstream"]; ok {
								var up string
								if json.Unmarshal(upRaw, &up) == nil && strings.TrimSpace(up) != "" {
									upstream = strings.ToLower(strings.TrimSpace(up))
								}
							}
						}
					}
				}
				base := strings.ToLower(strings.TrimSpace(ch.GetBaseURL()))
				if upstream == "official" && base != "" && strings.Contains(base, "ppinfra") {
					upstream = "ppinfra"
				}
				meta["upstream"] = upstream
			}
		}
	}

	// 对 Sora（sora-2 系列）标注秒单位，并记录视频秒数（由 usage.PromptTokens 提供）
	lowerModel := strings.ToLower(strings.TrimSpace(q.modelName))
	if strings.HasPrefix(lowerModel, "sora-2") {
		meta["unit"] = "sec"
		if usage != nil && usage.PromptTokens > 0 {
			meta["video_seconds"] = usage.PromptTokens
		}
	}
	// 对 Veo（veo-* 系列）标注秒单位
	if strings.HasPrefix(lowerModel, "veo-") {
		meta["unit"] = "sec"
		if usage != nil && usage.PromptTokens > 0 {
			meta["video_seconds"] = usage.PromptTokens
		}
	}

	firstResponseTime := q.GetFirstResponseTime()
	if firstResponseTime > 0 {
		meta["first_response"] = firstResponseTime
	}

	if usage != nil {
		extraTokens := usage.GetExtraTokens()

		for key, value := range extraTokens {
			meta[key] = value
			extraRatio := q.price.GetExtraRatio(key)
			meta[key+"_ratio"] = extraRatio
		}
	}

	if q.extraBillingData != nil {
		meta["extra_billing"] = q.extraBillingData
	}

	return meta
}

func (q *Quota) getRequestTime() int {
	return int(time.Since(q.startTime).Milliseconds())
}

// 通过 token 数获取消费配额
func (q *Quota) GetTotalQuota(promptTokens, completionTokens int, extraBilling map[string]types.ExtraBilling) (quota int) {
	if q.price.Type == model.TimesPriceType {
		quota = int(1000 * q.outputRatio)
	} else {
		// Sora（OpenAI /v1/videos）按秒计费：内部换算为 seconds × 1000（对齐 tokens 基线 1k）
		lowerModel := strings.ToLower(strings.TrimSpace(q.modelName))
		if strings.HasPrefix(lowerModel, "sora-2") {
			promptTokens = promptTokens * 1000
		}
		// Veo（Gemini /models/veo-*）按秒计费：同样按 seconds × 1000 结算
		if strings.HasPrefix(lowerModel, "veo-") {
			promptTokens = promptTokens * 1000
		}
		quota = int(math.Ceil((float64(promptTokens) * q.inputRatio) + (float64(completionTokens) * q.outputRatio)))
	}

	q.GetExtraBillingData(extraBilling)
	extraBillingQuota := 0
	if q.extraBillingData != nil {
		for _, value := range q.extraBillingData {
			extraBillingQuota += int(math.Ceil(
				float64(value.Price)*float64(config.QuotaPerUnit),
			)) * value.CallCount
		}
	}

	if extraBillingQuota > 0 {
		quota += int(math.Ceil(
			float64(extraBillingQuota) * q.groupRatio,
		))
	}

	if q.inputRatio != 0 && quota <= 0 {
		quota = 1
	}
	totalTokens := promptTokens + completionTokens
	if totalTokens == 0 {
		// in this case, must be some error happened
		// we cannot just return, because we may have to return the pre-consumed quota
		quota = 0
	}

	return quota
}

// 获取计算的 token 数
func (q *Quota) getComputeTokensByUsage(usage *types.Usage) (promptTokens, completionTokens int) {
	promptTokens = usage.PromptTokens
	completionTokens = usage.CompletionTokens

	extraTokens := usage.GetExtraTokens()

	for key, value := range extraTokens {
		extraRatio := q.price.GetExtraRatio(key)
		if model.GetExtraPriceIsPrompt(key) {
			promptTokens += model.GetIncreaseTokens(value, extraRatio)
		} else {
			completionTokens += model.GetIncreaseTokens(value, extraRatio)
		}
	}

	return
}

func (q *Quota) getComputeTokensByUsageEvent(usage *types.UsageEvent) (promptTokens, completionTokens int) {
	promptTokens = usage.InputTokens
	completionTokens = usage.OutputTokens
	extraTokens := usage.GetExtraTokens()

	for key, value := range extraTokens {
		extraRatio := q.price.GetExtraRatio(key)
		if model.GetExtraPriceIsPrompt(key) {
			promptTokens += model.GetIncreaseTokens(value, extraRatio)
		} else {
			completionTokens += model.GetIncreaseTokens(value, extraRatio)
		}
	}

	return
}

// 通过 usage 获取消费配额
func (q *Quota) GetTotalQuotaByUsage(usage *types.Usage) (quota int) {
	promptTokens, completionTokens := q.getComputeTokensByUsage(usage)
	return q.GetTotalQuota(promptTokens, completionTokens, usage.ExtraBilling)
}

func (q *Quota) GetFirstResponseTime() int64 {
	// 先判断 firstResponseTime 是否为0
	if q.firstResponseTime.IsZero() {
		return 0
	}

	return q.firstResponseTime.Sub(q.startTime).Milliseconds()
}

func (q *Quota) SetFirstResponseTime(firstResponseTime time.Time) {
	q.firstResponseTime = firstResponseTime
}

type ExtraBillingData struct {
	Type      string  `json:"type"`
	CallCount int     `json:"call_count"`
	Price     float64 `json:"price"`
}

func (q *Quota) GetExtraBillingData(extraBilling map[string]types.ExtraBilling) {
	if extraBilling == nil {
		return
	}

	extraBillingData := make(map[string]ExtraBillingData)
	for serviceType, value := range extraBilling {
		extraBillingData[serviceType] = ExtraBillingData{
			Type:      value.Type,
			CallCount: value.CallCount,
			Price:     getDefaultExtraServicePrice(serviceType, q.modelName, value.Type),
		}

	}

	if len(extraBillingData) == 0 {
		return
	}

	q.extraBillingData = extraBillingData
}
