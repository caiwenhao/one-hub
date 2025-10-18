package relay

import (
    "fmt"
    "net/http"
    "one-api/common"
    "one-api/common/config"
    "one-api/common/utils"
    "one-api/model"
    "one-api/providers/claude"
    "one-api/providers/gemini"
    "one-api/types"
    "sort"
    "strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

    "github.com/gin-gonic/gin"
    "github.com/shopspring/decimal"
)

// https://platform.openai.com/docs/api-reference/models/list
type OpenAIModels struct {
	Id      string  `json:"id"`
	Object  string  `json:"object"`
	Created int     `json:"created"`
	OwnedBy *string `json:"owned_by"`
}

func ListModelsByToken(c *gin.Context) {
	groupName := c.GetString("token_group")
	if groupName == "" {
		groupName = c.GetString("group")
	}

	if groupName == "" {
		common.AbortWithMessage(c, http.StatusServiceUnavailable, "分组不存在")
		return
	}

	models, err := model.ChannelGroup.GetGroupModels(groupName)
	if err != nil {
		c.JSON(200, gin.H{
			"object": "list",
			"data":   []string{},
		})
		return
	}
	sort.Strings(models)

	var groupOpenAIModels []*OpenAIModels
	for _, modelName := range models {
		groupOpenAIModels = append(groupOpenAIModels, getOpenAIModelWithName(modelName))
	}

	// 根据 OwnedBy 排序
	sort.Slice(groupOpenAIModels, func(i, j int) bool {
		if groupOpenAIModels[i].OwnedBy == nil {
			return true // 假设 nil 值小于任何非 nil 值
		}
		if groupOpenAIModels[j].OwnedBy == nil {
			return false // 假设任何非 nil 值大于 nil 值
		}
		return *groupOpenAIModels[i].OwnedBy < *groupOpenAIModels[j].OwnedBy
	})

	c.JSON(200, gin.H{
		"object": "list",
		"data":   groupOpenAIModels,
	})
}

// https://generativelanguage.googleapis.com/v1beta/models?key=xxxxxxx
func ListGeminiModelsByToken(c *gin.Context) {
	groupName := c.GetString("token_group")
	if groupName == "" {
		groupName = c.GetString("group")
	}

	if groupName == "" {
		common.AbortWithMessage(c, http.StatusServiceUnavailable, "分组不存在")
		return
	}

	models, err := model.ChannelGroup.GetGroupModels(groupName)
	if err != nil {
		c.JSON(200, gemini.ModelListResponse{
			Models: []gemini.ModelDetails{},
		})
		return
	}
	sort.Strings(models)

	var geminiModels []gemini.ModelDetails
	for _, modelName := range models {
		// Get the price to check if it's a Gemini model (channel_type=25)
		price := model.PricingInstance.GetPrice(modelName)
		if price.ChannelType == config.ChannelTypeGemini {
			geminiModels = append(geminiModels, gemini.ModelDetails{
				Name:        fmt.Sprintf("models/%s", modelName),
				DisplayName: cases.Title(language.Und).String(strings.ReplaceAll(modelName, "-", " ")),
				SupportedGenerationMethods: []string{
					"generateContent",
				},
			})
		}
	}

	c.JSON(200, gemini.ModelListResponse{
		Models: geminiModels,
	})
}

func ListClaudeModelsByToken(c *gin.Context) {
	groupName := c.GetString("token_group")
	if groupName == "" {
		groupName = c.GetString("group")
	}

	if groupName == "" {
		common.AbortWithMessage(c, http.StatusServiceUnavailable, "分组不存在")
		return
	}

	models, err := model.ChannelGroup.GetGroupModels(groupName)
	if err != nil {
		c.JSON(200, claude.ModelListResponse{
			Data: []claude.Model{},
		})
		return
	}
	sort.Strings(models)

	var claudeModelsData []claude.Model
	for _, modelName := range models {
		// Get the price to check if it's a Gemini model (channel_type=25)
		price := model.PricingInstance.GetPrice(modelName)
		if price.ChannelType == config.ChannelTypeAnthropic {
			claudeModelsData = append(claudeModelsData, claude.Model{
				ID:   modelName,
				Type: "model",
			})
		}
	}

	c.JSON(200, claude.ModelListResponse{
		Data: claudeModelsData,
	})
}

func inferOwnedBy(modelName string, price *model.Price) *string {
	var channelType int
	if price != nil {
		channelType = price.ChannelType
		if ownedType := price.GetOwnedByType(); ownedType != config.ChannelTypeUnknown {
			if owned := getModelOwnedBy(ownedType); owned != nil && *owned != model.UnknownOwnedBy && *owned != "" {
				return owned
			}
		}
	}

	// 否则根据模型命名进行供应商推断，避免显示“未知”
	lower := strings.ToLower(modelName)
	if strings.HasPrefix(lower, "kling") || strings.Contains(lower, "kling-") || strings.Contains(lower, "kling_") {
		name := model.ModelOwnedBysInstance.GetName(config.ChannelTypeKling)
		if name == model.UnknownOwnedBy || name == "" {
			fallback := "Kling"
			return &fallback
		}
		return &name
	}
	if strings.HasPrefix(lower, "vidu") || strings.HasPrefix(lower, "viduq") || strings.Contains(lower, "vidu-") || strings.Contains(lower, "vidu_") {
		name := model.ModelOwnedBysInstance.GetName(config.ChannelTypeVidu)
		if name == model.UnknownOwnedBy || name == "" {
			fallback := "Vidu"
			return &fallback
		}
		return &name
	}
	if strings.HasPrefix(lower, "doubao") || strings.Contains(lower, "doubao-") || strings.Contains(lower, "wan2.1") {
		name := model.ModelOwnedBysInstance.GetName(config.ChannelTypeVolcArk)
		if name == model.UnknownOwnedBy || name == "" {
			fallback := "Volc Ark"
			return &fallback
		}
		return &name
	}
	if strings.Contains(lower, "deepseek") {
		name := model.ModelOwnedBysInstance.GetName(config.ChannelTypeDeepseek)
		if name == model.UnknownOwnedBy || name == "" {
			fallback := "Deepseek"
			return &fallback
		}
		return &name
	}
	// minimaxi 视频基础模型归属推断
	if strings.Contains(lower, "minimax-hailuo-02") ||
		lower == "t2v-01" || lower == "t2v-01-director" ||
		lower == "i2v-01" || lower == "i2v-01-live" ||
		lower == "s2v-01" {
		name := model.ModelOwnedBysInstance.GetName(config.ChannelTypeMiniMax)
		if name == model.UnknownOwnedBy || name == "" {
			fallback := "minimaxi"
			return &fallback
		}
		return &name
	}
	return getModelOwnedBy(channelType)
}

func ListModelsForAdmin(c *gin.Context) {
	prices := model.PricingInstance.GetAllPrices()
	var openAIModels []OpenAIModels
	for modelId, price := range prices {
		openAIModels = append(openAIModels, OpenAIModels{
			Id:      modelId,
			Object:  "model",
			Created: 1677649963,
			OwnedBy: inferOwnedBy(modelId, price),
		})
	}
	// 根据 OwnedBy 排序
	sort.Slice(openAIModels, func(i, j int) bool {
		if openAIModels[i].OwnedBy == nil {
			return true // 假设 nil 值小于任何非 nil 值
		}
		if openAIModels[j].OwnedBy == nil {
			return false // 假设任何非 nil 值大于 nil 值
		}
		return *openAIModels[i].OwnedBy < *openAIModels[j].OwnedBy
	})

	c.JSON(200, gin.H{
		"object": "list",
		"data":   openAIModels,
	})
}

func RetrieveModel(c *gin.Context) {
	modelName := c.Param("model")
	openaiModel := getOpenAIModelWithName(modelName)
	if *openaiModel.OwnedBy != model.UnknownOwnedBy {
		c.JSON(200, openaiModel)
	} else {
		openAIError := types.OpenAIError{
			Message: fmt.Sprintf("The model '%s' does not exist", modelName),
			Type:    "invalid_request_error",
			Param:   "model",
			Code:    "model_not_found",
		}
		c.JSON(200, gin.H{
			"error": openAIError,
		})
	}
}

func getModelOwnedBy(channelType int) (ownedBy *string) {
	ownedByName := model.ModelOwnedBysInstance.GetName(channelType)
	if ownedByName != "" {
		return &ownedByName
	}

	return &model.UnknownOwnedBy
}

func getOpenAIModelWithName(modelName string) *OpenAIModels {
	price := model.PricingInstance.GetPrice(modelName)

	return &OpenAIModels{
		Id:      modelName,
		Object:  "model",
		Created: 1677649963,
		OwnedBy: inferOwnedBy(modelName, price),
	}
}

func GetModelOwnedBy(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    model.ModelOwnedBysInstance.GetAll(),
	})
}

type ModelPrice struct {
	Type   string  `json:"type"`
	Input  float64 `json:"input"`
	Output float64 `json:"output"`
}

type AvailableModelResponse struct {
    Groups      []string     `json:"groups"`
    OwnedBy     string       `json:"owned_by"`
    OwnedByType int          `json:"owned_by_type"`
    ChannelType int          `json:"channel_type"`
    Price       *model.Price `json:"price"`
    PriceDisplay *PriceDisplay `json:"price_display,omitempty"`
    Variants    []VariantDisplay `json:"variants,omitempty"`
}

// PriceDisplay 为用户端展示的非侵入式价格提示（人民币），不包含敏感渠道信息
type PriceDisplay struct {
    Type      string `json:"type"`                  // tokens | times
    Unit      string `json:"unit"`                  // USD/1k tokens 或 USD/each
    // 美元展示（面向用户展示）
    InputUSD  string `json:"input_usd,omitempty"`
    OutputUSD string `json:"output_usd,omitempty"`
    // 兼容旧字段（可能被前端历史逻辑使用）
    InputRMB  string `json:"input_rmb,omitempty"`
    OutputRMB string `json:"output_rmb,omitempty"`
}

type VariantDisplay struct {
    Model        string        `json:"model"`
    PriceDisplay *PriceDisplay `json:"price_display,omitempty"`
}

func buildPriceDisplay(p *model.Price) *PriceDisplay {
    if p == nil {
        return nil
    }
    d := &PriceDisplay{Type: p.Type}
    // 主展示改为美元
    if p.Type == model.TimesPriceType {
        d.Unit = "USD/each"
    } else {
        d.Unit = "USD/1k tokens"
    }
    // 计算美元展示（统一保留 6 位小数，避免长尾差异）
    din := decimal.NewFromFloat(p.GetInput()).Mul(decimal.NewFromFloat(model.DollarRate)).Round(6)
    dout := decimal.NewFromFloat(p.GetOutput()).Mul(decimal.NewFromFloat(model.DollarRate)).Round(6)
    d.InputUSD = din.StringFixed(6)
    d.OutputUSD = dout.StringFixed(6)
    // 兼容保留人民币字段（前端历史逻辑回退使用，不作为主展示）
    rin := decimal.NewFromFloat(p.GetInput()).Mul(decimal.NewFromFloat(model.RMBRate)).Round(6)
    rout := decimal.NewFromFloat(p.GetOutput()).Mul(decimal.NewFromFloat(model.RMBRate)).Round(6)
    d.InputRMB = rin.StringFixed(6)
    d.OutputRMB = rout.StringFixed(6)
    return d
}

func shouldAttachVariants(base string) (bool, string) {
    lower := strings.ToLower(strings.TrimSpace(base))
    m := map[string]string{
        "minimax-hailuo-02": "minimax-hailuo-02",
        "t2v-01":             "t2v-01",
        "t2v-01-director":    "t2v-01-director",
        "i2v-01":             "i2v-01",
        "i2v-01-live":        "i2v-01-live",
        "s2v-01":             "s2v-01",
    }
    for k, seg := range m {
        if strings.Contains(lower, k) || lower == k {
            return true, seg
        }
    }
    return false, ""
}

func collectMiniMaxVariants(baseSegment string, allPrices map[string]*model.Price) []VariantDisplay {
    var variants []VariantDisplay
    needle := "-" + baseSegment + "-"
    for name, pr := range allPrices {
        lower := strings.ToLower(name)
        if pr == nil || pr.ChannelType != config.ChannelTypeMiniMax || pr.Type != model.TimesPriceType {
            continue
        }
        if !strings.HasPrefix(lower, "minimax-") {
            continue
        }
        if !strings.Contains(lower, needle) {
            continue
        }
        variants = append(variants, VariantDisplay{
            Model:        name,
            PriceDisplay: buildPriceDisplay(pr),
        })
    }
    sort.Slice(variants, func(i, j int) bool { return variants[i].Model < variants[j].Model })
    return variants
}

func AvailableModel(c *gin.Context) {
	groupName := c.GetString("group")

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    getAvailableModels(groupName),
	})
}

func GetAvailableModels(groupName string) map[string]*AvailableModelResponse {
	return getAvailableModels(groupName)
}

func getAvailableModels(groupName string) map[string]*AvailableModelResponse {
	publicModels := model.ChannelGroup.GetModelsGroups()
	publicGroups := model.GlobalUserGroupRatio.GetPublicGroupList()
	if groupName != "" && !utils.Contains(groupName, publicGroups) {
		publicGroups = append(publicGroups, groupName)
	}

	availableModels := make(map[string]*AvailableModelResponse, len(publicModels))

	for modelName, group := range publicModels {
		groups := []string{}
		for _, publicGroup := range publicGroups {
			if group[publicGroup] {
				groups = append(groups, publicGroup)
			}
		}

		if len(groups) == 0 {
			continue
		}

        if _, ok := availableModels[modelName]; !ok {
            price := model.PricingInstance.GetPrice(modelName)
            owned := inferOwnedBy(modelName, price)
            ownedName := model.UnknownOwnedBy
            if owned != nil && *owned != "" {
                ownedName = *owned
            }
            // 计算人民币展示（不改变内部定价基准）；仅在已知渠道类型下展示，避免默认价误导
            var disp *PriceDisplay
            if price.ChannelType != config.ChannelTypeUnknown {
                disp = buildPriceDisplay(price)
            }

            // 精确计费组合（仅对 minimaxi 视频基础模型等进行展开）
            var variants []VariantDisplay
            if ok, seg := shouldAttachVariants(modelName); ok {
                variants = collectMiniMaxVariants(seg, model.PricingInstance.GetAllPrices())
            }

            availableModels[modelName] = &AvailableModelResponse{
                Groups:      groups,
                OwnedBy:     ownedName,
                OwnedByType: price.GetOwnedByType(),
                ChannelType: price.ChannelType,
                Price:        price,
                PriceDisplay: disp,
                Variants:     variants,
            }
        }
    }

	return availableModels
}
