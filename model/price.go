package model

import (
	"one-api/common/config"
	"strings"

	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const (
	TokensPriceType    = "tokens"
	TimesPriceType     = "times"
	DefaultPrice       = 30.0
	DollarRate         = 0.002
	RMBRate            = 0.014
	DefaultCacheRatios = 0.5
	DefaultAudioRatio  = 40

	DefaultCachedWriteRatio = 1.25
	DefaultCachedReadRatio  = 0.1
)

func GetIncreaseTokens(tokens int, ratio float64) int {
	return int(float64(tokens) * (ratio - 1))
}

var ExtraKeyIsPrompt = map[string]bool{
	config.UsageExtraCache:             true,
	config.UsageExtraCachedWrite:       true,
	config.UsageExtraCachedRead:        true,
	config.UsageExtraInputAudio:        true,
	config.UsageExtraOutputAudio:       false,
	config.UsageExtraReasoning:         false,
	config.UsageExtraInputTextTokens:   true,
	config.UsageExtraOutputTextTokens:  false,
	config.UsageExtraInputImageTokens:  true,
	config.UsageExtraOutputImageTokens: false,
}

func GetExtraPriceIsPrompt(key string) bool {
	return ExtraKeyIsPrompt[key]
}

var defaultExtraPrice = map[string]float64{
	config.UsageExtraCache:            1,
	config.UsageExtraCachedWrite:      1.25,
	config.UsageExtraCachedRead:       0.1,
	config.UsageExtraInputAudio:       1,
	config.UsageExtraOutputAudio:      1,
	config.UsageExtraReasoning:        1,
	config.UsageExtraInputTextTokens:  1,
	config.UsageExtraOutputTextTokens: 1,
}

type Price struct {
	Model       string  `json:"model" gorm:"type:varchar(100)" binding:"required"`
	Type        string  `json:"type"  gorm:"default:'tokens'" binding:"required"`
	ChannelType int     `json:"channel_type" gorm:"default:0" binding:"gte=0"`
	OwnedByType int     `json:"owned_by_type" gorm:"default:0" binding:"gte=0"`
	Input       float64 `json:"input" gorm:"default:0" binding:"gte=0"`
	Output      float64 `json:"output" gorm:"default:0" binding:"gte=0"`
	Locked      bool    `json:"locked" gorm:"default:false"` // 如果模型为locked 则覆盖模式不会更新locked的模型价格

	ExtraRatios *datatypes.JSONType[map[string]float64] `json:"extra_ratios,omitempty" gorm:"type:json"`
}

func (price *Price) Normalize() {
	if price == nil {
		return
	}

	price.Model = strings.TrimSpace(price.Model)
	if price.Type == "" {
		price.Type = TokensPriceType
	}
	if price.Type == TimesPriceType && price.Output == 0 {
		price.Output = price.Input
	}
	price.OwnedByType = resolveOwnedByType(price.Model, price.OwnedByType, price.ChannelType)
}

func (price *Price) GetOwnedByType() int {
	if price == nil {
		return config.ChannelTypeUnknown
	}

	return resolveOwnedByType(price.Model, price.OwnedByType, price.ChannelType)
}

func GetAllPrices() ([]*Price, error) {
	var prices []*Price
	if err := DB.Find(&prices).Error; err != nil {
		return nil, err
	}
	for _, price := range prices {
		price.Normalize()
	}
	// if config.ExtraTokenPriceJson == "" {
	// 	return prices, nil
	// }

	// extraRatios := make(map[string]map[string]float64)
	// err := json.Unmarshal([]byte(config.ExtraTokenPriceJson), &extraRatios)
	// if err != nil {
	// 	return nil, err
	// }

	// for _, price := range prices {
	// 	if ratio, ok := extraRatios[price.Model]; ok {
	// 		price.ExtraRatios = ratio
	// 	}
	// }

	return prices, nil
}

func (price *Price) Update(modelName string) error {
	if price.Model == "" {
		price.Model = modelName
	}
	price.Normalize()
	if err := DB.Model(price).Select("*").Where("model = ?", modelName).Updates(price).Error; err != nil {
		return err
	}

	return nil
}

func (price *Price) Insert() error {
	price.Normalize()
	if err := DB.Create(price).Error; err != nil {
		return err
	}

	return nil
}

func (price *Price) GetInput() float64 {
	if price.Input <= 0 {
		return 0
	}
	return price.Input
}

func (price *Price) GetOutput() float64 {
	if price.Output <= 0 || price.Type == TimesPriceType {
		return 0
	}

	return price.Output
}

func (price *Price) GetExtraRatio(key string) float64 {
	if price.ExtraRatios != nil {
		extraRatios := price.ExtraRatios.Data()
		if ratio, ok := extraRatios[key]; ok {
			return ratio
		}
	}

	ratio, ok := defaultExtraPrice[key]
	if !ok {
		return 1
	}

	return ratio
}

func (price *Price) FetchInputCurrencyPrice(rate float64) string {
	r := decimal.NewFromFloat(price.GetInput()).Mul(decimal.NewFromFloat(rate))
	return r.String()
}

func (price *Price) FetchOutputCurrencyPrice(rate float64) string {
	r := decimal.NewFromFloat(price.GetOutput()).Mul(decimal.NewFromFloat(rate))
	return r.String()
}

func UpdatePrices(tx *gorm.DB, models []string, prices *Price) error {
	err := tx.Model(Price{}).Where("model IN (?)", models).Select("*").Omit("model").Updates(
		Price{
			Type:        prices.Type,
			ChannelType: prices.ChannelType,
			Input:       prices.Input,
			Output:      prices.Output,
			Locked:      prices.Locked,
			ExtraRatios: prices.ExtraRatios,
		}).Error

	return err
}

func DeletePrices(tx *gorm.DB, models []string) error {
	err := tx.Where("model IN (?)", models).Delete(&Price{}).Error

	return err
}

func InsertPrices(tx *gorm.DB, prices []*Price) error {
	err := tx.CreateInBatches(prices, 100).Error
	return err
}

func DeleteAllPrices(tx *gorm.DB) error {
	err := tx.Where("1=1").Delete(&Price{}).Error
	return err
}

// 只删除未lock的价格
func DeleteAllPricesNotLock(tx *gorm.DB) error {
	err := tx.Where("locked = ?", false).Delete(&Price{}).Error
	return err
}

// 只删除指定的未lock的数据
func DeletePricesByModelNameAndNotLock(tx *gorm.DB, models []string) error {
	err := tx.Where("locked = ? and model IN (?)", false, models).Delete(&Price{}).Error
	return err
}

func (price *Price) Delete() error {
	return DB.Where("model = ?", price.Model).Delete(&Price{}).Error
}

type ModelType struct {
	Ratio []float64
	Type  int
}

var ownedByKeywordRules = []struct {
	Type     int
	Keywords []string
}{
	{config.ChannelTypeDeepseek, []string{"deepseek"}},
	{config.ChannelTypeAnthropic, []string{"claude", "anthropic/"}},
	{config.ChannelTypeOpenAI, []string{"gpt-", "gpt", "davinci", "curie", "babbage", "ada", "whisper", "dall-e", "text-embedding", "o1-", "o3-"}},
	{config.ChannelTypeGemini, []string{"gemini", "palm-"}},
	{config.ChannelTypeAli, []string{"qwen", "tongyi"}},
	{config.ChannelTypeLLAMA, []string{"llama", "meta-llama"}},
	{config.ChannelTypeMistral, []string{"mistral", "mixtral"}},
	{config.ChannelTypeMoonshot, []string{"moonshot"}},
	{config.ChannelTypeLingyi, []string{"yi-"}},
	{config.ChannelTypeZhipu, []string{"glm", "cogview"}},
	{config.ChannelTypeBaichuan, []string{"baichuan"}},
	{config.ChannelTypeMiniMax, []string{"abab", "minimax"}},
	{config.ChannelTypeGroq, []string{"groq"}},
	{config.ChannelTypeKling, []string{"kling"}},
	{config.ChannelTypeVidu, []string{"vidu"}},
	{config.ChannelTypeVolcArk, []string{"doubao", "wan2.1", "seaweed", "seedance"}},
	{config.ChannelTypeHunyuan, []string{"hunyuan"}},
	{config.ChannelTypeSuno, []string{"suno", "chirp"}},
	{config.ChannelTypeStabilityAI, []string{"stable-diffusion", "sd3"}},
	{config.ChannelTypeCohere, []string{"cohere", "command-"}},
	{config.ChannelTypeCloudflareAI, []string{"@cf/"}},
}

func resolveOwnedByType(modelName string, ownedByType int, channelType int) int {
	if ownedByType != 0 && ownedByType != config.ChannelTypeUnknown {
		return ownedByType
	}

	detected := detectOwnedByType(modelName, channelType)
	if detected != config.ChannelTypeUnknown {
		return detected
	}

	if channelType != 0 && channelType != config.ChannelTypeUnknown {
		return channelType
	}

	return config.ChannelTypeUnknown
}

func detectOwnedByType(modelName string, fallback int) int {
	lower := strings.ToLower(strings.TrimSpace(modelName))
	if lower == "" {
		if fallback != 0 {
			return fallback
		}
		return config.ChannelTypeUnknown
	}

	for _, rule := range ownedByKeywordRules {
		for _, keyword := range rule.Keywords {
			if strings.Contains(lower, keyword) {
				return rule.Type
			}
		}
	}

	if fallback != 0 {
		return fallback
	}

	return config.ChannelTypeUnknown
}

// 1 === $0.002 / 1K tokens
// 1 === ￥0.014 / 1k tokens
func GetDefaultPrice() []*Price {
	ModelTypes := map[string]ModelType{
		// 	$0.03 / 1K tokens	$0.06 / 1K tokens
		"gpt-4":      {[]float64{15, 30}, config.ChannelTypeOpenAI},
		"gpt-4-0314": {[]float64{15, 30}, config.ChannelTypeOpenAI},
		"gpt-4-0613": {[]float64{15, 30}, config.ChannelTypeOpenAI},
		// 	$0.06 / 1K tokens	$0.12 / 1K tokens
		"gpt-4-32k":      {[]float64{30, 60}, config.ChannelTypeOpenAI},
		"gpt-4-32k-0314": {[]float64{30, 60}, config.ChannelTypeOpenAI},
		"gpt-4-32k-0613": {[]float64{30, 60}, config.ChannelTypeOpenAI},
		// 	$0.01 / 1K tokens	$0.03 / 1K tokens
		"gpt-4-preview":          {[]float64{5, 15}, config.ChannelTypeOpenAI},
		"gpt-4-turbo":            {[]float64{5, 15}, config.ChannelTypeOpenAI},
		"gpt-4-turbo-2024-04-09": {[]float64{5, 15}, config.ChannelTypeOpenAI},
		"gpt-4-1106-preview":     {[]float64{5, 15}, config.ChannelTypeOpenAI},
		"gpt-4-0125-preview":     {[]float64{5, 15}, config.ChannelTypeOpenAI},
		"gpt-4-turbo-preview":    {[]float64{5, 15}, config.ChannelTypeOpenAI},
		"gpt-4-vision-preview":   {[]float64{5, 15}, config.ChannelTypeOpenAI},
		// $0.005 / 1K tokens	$0.015 / 1K tokens
		"gpt-4o": {[]float64{2.5, 7.5}, config.ChannelTypeOpenAI},
		// 	$0.0005 / 1K tokens	$0.0015 / 1K tokens
		"gpt-3.5-turbo":      {[]float64{0.25, 0.75}, config.ChannelTypeOpenAI},
		"gpt-3.5-turbo-0125": {[]float64{0.25, 0.75}, config.ChannelTypeOpenAI},
		// 	$0.0015 / 1K tokens	$0.002 / 1K tokens
		"gpt-3.5-turbo-0301":     {[]float64{0.75, 1}, config.ChannelTypeOpenAI},
		"gpt-3.5-turbo-0613":     {[]float64{0.75, 1}, config.ChannelTypeOpenAI},
		"gpt-3.5-turbo-instruct": {[]float64{0.75, 1}, config.ChannelTypeOpenAI},
		// 	$0.003 / 1K tokens	$0.004 / 1K tokens
		"gpt-3.5-turbo-16k":      {[]float64{1.5, 2}, config.ChannelTypeOpenAI},
		"gpt-3.5-turbo-16k-0613": {[]float64{1.5, 2}, config.ChannelTypeOpenAI},
		// 	$0.001 / 1K tokens	$0.002 / 1K tokens
		"gpt-3.5-turbo-1106": {[]float64{0.5, 1}, config.ChannelTypeOpenAI},
		// 	$0.0020 / 1K tokens
		"davinci-002": {[]float64{1, 1}, config.ChannelTypeOpenAI},
		// 	$0.0004 / 1K tokens
		"babbage-002": {[]float64{0.2, 0.2}, config.ChannelTypeOpenAI},
		// $0.006 / minute -> $0.006 / 150 words -> $0.006 / 200 tokens -> $0.03 / 1k tokens
		"whisper-1": {[]float64{15, 15}, config.ChannelTypeOpenAI},
		// $0.015 / 1K characters
		"tts-1":      {[]float64{7.5, 7.5}, config.ChannelTypeOpenAI},
		"tts-1-1106": {[]float64{7.5, 7.5}, config.ChannelTypeOpenAI},
		// $0.030 / 1K characters
		"tts-1-hd":               {[]float64{15, 15}, config.ChannelTypeOpenAI},
		"tts-1-hd-1106":          {[]float64{15, 15}, config.ChannelTypeOpenAI},
		"text-embedding-ada-002": {[]float64{0.05, 0.05}, config.ChannelTypeOpenAI},
		// 	$0.00002 / 1K tokens
		"text-embedding-3-small": {[]float64{0.01, 0.01}, config.ChannelTypeOpenAI},
		// 	$0.00013 / 1K tokens
		"text-embedding-3-large":  {[]float64{0.065, 0.065}, config.ChannelTypeOpenAI},
		"doubao-embedding":        {[]float64{0.035714286, 0.035714286}, config.ChannelTypeVolcArk},
		"doubao-embedding-large":  {[]float64{0.05, 0.05}, config.ChannelTypeVolcArk},
		"doubao-embedding-vision": {[]float64{0.05, 0.05}, config.ChannelTypeVolcArk},
		"text-moderation-stable":  {[]float64{0.1, 0.1}, config.ChannelTypeOpenAI},
		"text-moderation-latest":  {[]float64{0.1, 0.1}, config.ChannelTypeOpenAI},
		// $0.016 - $0.020 / image
		"dall-e-2": {[]float64{8, 8}, config.ChannelTypeOpenAI},
		// $0.040 - $0.120 / image
		"dall-e-3": {[]float64{20, 20}, config.ChannelTypeOpenAI},

		// $0.80/million tokens $2.40/million tokens
		"claude-instant-1.2": {[]float64{0.4, 1.2}, config.ChannelTypeAnthropic},
		// $8.00/million tokens $24.00/million tokens
		"claude-2.0": {[]float64{4, 12}, config.ChannelTypeAnthropic},
		"claude-2.1": {[]float64{4, 12}, config.ChannelTypeAnthropic},
		// $15 / M $75 / M
		"claude-3-opus-20240229": {[]float64{7.5, 22.5}, config.ChannelTypeAnthropic},
		//  $3 / M $15 / M
		"claude-3-sonnet-20240229": {[]float64{1.3, 3.9}, config.ChannelTypeAnthropic},
		//  $0.25 / M $1.25 / M  0.00025$ / 1k tokens 0.00125$ / 1k tokens
		"claude-3-haiku-20240307": {[]float64{0.125, 0.625}, config.ChannelTypeAnthropic},

		// ￥0.004 / 1k tokens ￥0.008 / 1k tokens
		"ERNIE-Speed": {[]float64{0.2857, 0.5714}, config.ChannelTypeBaidu},
		// ￥0.012 / 1k tokens ￥0.012 / 1k tokens
		"ERNIE-Bot":    {[]float64{0.8572, 0.8572}, config.ChannelTypeBaidu},
		"ERNIE-3.5-8K": {[]float64{0.8572, 0.8572}, config.ChannelTypeBaidu},
		// 0.024元/千tokens 0.048元/千tokens
		"ERNIE-Bot-8k": {[]float64{1.7143, 3.4286}, config.ChannelTypeBaidu},
		// ￥0.008 / 1k tokens ￥0.008 / 1k tokens
		"ERNIE-Bot-turbo": {[]float64{0.5715, 0.5715}, config.ChannelTypeBaidu},
		// ￥0.12 / 1k tokens ￥0.12 / 1k tokens
		"ERNIE-Bot-4": {[]float64{8.572, 8.572}, config.ChannelTypeBaidu},
		"ERNIE-4.0":   {[]float64{8.572, 8.572}, config.ChannelTypeBaidu},
		// ￥0.002 / 1k tokens
		"Embedding-V1": {[]float64{0.1429, 0.1429}, config.ChannelTypeBaidu},
		// ￥0.004 / 1k tokens
		"BLOOMZ-7B": {[]float64{0.2857, 0.2857}, config.ChannelTypeBaidu},

		"PaLM-2": {[]float64{1, 1}, config.ChannelTypePaLM},
		// $0.50 / 1 million tokens  $1.50 / 1 million tokens
		// 0.0005$ / 1k tokens 0.0015$ / 1k tokens
		"gemini-pro":        {[]float64{0.25, 0.75}, config.ChannelTypeGemini},
		"gemini-pro-vision": {[]float64{0.25, 0.75}, config.ChannelTypeGemini},
		"gemini-1.0-pro":    {[]float64{0.25, 0.75}, config.ChannelTypeGemini},
		// $7 / 1 million tokens  $21 / 1 million tokens
		"gemini-1.5-pro":          {[]float64{1.75, 5.25}, config.ChannelTypeGemini},
		"gemini-1.5-pro-latest":   {[]float64{1.75, 5.25}, config.ChannelTypeGemini},
		"gemini-1.5-flash":        {[]float64{0.175, 0.265}, config.ChannelTypeGemini},
		"gemini-1.5-flash-latest": {[]float64{0.175, 0.265}, config.ChannelTypeGemini},
		"gemini-ultra":            {[]float64{1, 1}, config.ChannelTypeGemini},

		// ￥0.005 / 1k tokens
		"glm-3-turbo": {[]float64{0.3572, 0.3572}, config.ChannelTypeZhipu},
		// ￥0.1 / 1k tokens
		"glm-4":  {[]float64{7.143, 7.143}, config.ChannelTypeZhipu},
		"glm-4v": {[]float64{7.143, 7.143}, config.ChannelTypeZhipu},
		// ￥0.0005 / 1k tokens
		"embedding-2": {[]float64{0.0357, 0.0357}, config.ChannelTypeZhipu},
		// ￥0.25 / 1张图片
		"cogview-3": {[]float64{17.8571, 17.8571}, config.ChannelTypeZhipu},

		// ￥0.008 / 1k tokens
		"qwen-turbo": {[]float64{0.5715, 0.5715}, config.ChannelTypeAli},
		// ￥0.02 / 1k tokens
		"qwen-plus":   {[]float64{1.4286, 1.4286}, config.ChannelTypeAli},
		"qwen-vl-max": {[]float64{1.4286, 1.4286}, config.ChannelTypeAli},
		// 0.12元/1,000tokens
		"qwen-max":             {[]float64{8.5714, 8.5714}, config.ChannelTypeAli},
		"qwen-max-longcontext": {[]float64{8.5714, 8.5714}, config.ChannelTypeAli},
		// 0.008元/1,000tokens
		"qwen-vl-plus": {[]float64{0.5715, 0.5715}, config.ChannelTypeAli},
		// ￥0.0007 / 1k tokens
		"text-embedding-v1": {[]float64{0.05, 0.05}, config.ChannelTypeAli},

		// ￥0.018 / 1k tokens
		"SparkDesk":      {[]float64{1.2858, 1.2858}, config.ChannelTypeXunfei},
		"SparkDesk-v1.1": {[]float64{0, 0}, config.ChannelTypeXunfei},
		"SparkDesk-v2.1": {[]float64{2.1429, 2.1429}, config.ChannelTypeXunfei},
		"SparkDesk-v3.1": {[]float64{2.1429, 2.1429}, config.ChannelTypeXunfei},
		"SparkDesk-v3.5": {[]float64{2.1429, 2.1429}, config.ChannelTypeXunfei},
		"SparkDesk-v4.0": {[]float64{7.1429, 7.1429}, config.ChannelTypeXunfei},

		// ¥0.012 / 1k tokens
		"360GPT_S2_V9": {[]float64{0.8572, 0.8572}, config.ChannelType360},
		// ¥0.001 / 1k tokens
		"embedding-bert-512-v1":     {[]float64{0.0715, 0.0715}, config.ChannelType360},
		"embedding_s1_v1":           {[]float64{0.0715, 0.0715}, config.ChannelType360},
		"semantic_similarity_s1_v1": {[]float64{0.0715, 0.0715}, config.ChannelType360},

		// ¥0.1 / 1k tokens  // https://cloud.tencent.com/document/product/1729/97731#e0e6be58-60c8-469f-bdeb-6c264ce3b4d0
		"hunyuan": {[]float64{7.143, 7.143}, config.ChannelTypeTencent},
		// https://cloud.tencent.com/document/product/1729/97731#e0e6be58-60c8-469f-bdeb-6c264ce3b4d0
		// ¥0.01 / 1k tokens
		"ChatStd": {[]float64{0.7143, 0.7143}, config.ChannelTypeTencent},
		//¥0.1 / 1k tokens
		"ChatPro": {[]float64{7.143, 7.143}, config.ChannelTypeTencent},

		"Baichuan2-Turbo":         {[]float64{0.5715, 0.5715}, config.ChannelTypeBaichuan}, // ¥0.008 / 1k tokens
		"Baichuan2-Turbo-192k":    {[]float64{1.143, 1.143}, config.ChannelTypeBaichuan},   // ¥0.016 / 1k tokens
		"Baichuan2-53B":           {[]float64{1.4286, 1.4286}, config.ChannelTypeBaichuan}, // ¥0.02 / 1k tokens
		"Baichuan-Text-Embedding": {[]float64{0.0357, 0.0357}, config.ChannelTypeBaichuan}, // ¥0.0005 / 1k tokens

		"abab5.5s-chat": {[]float64{0.3572, 0.3572}, config.ChannelTypeMiniMax},   // ¥0.005 / 1k tokens
		"abab5.5-chat":  {[]float64{1.0714, 1.0714}, config.ChannelTypeMiniMax},   // ¥0.015 / 1k tokens
		"abab6-chat":    {[]float64{14.2857, 14.2857}, config.ChannelTypeMiniMax}, // ¥0.2 / 1k tokens
		"embo-01":       {[]float64{0.0357, 0.0357}, config.ChannelTypeMiniMax},   // ¥0.0005 / 1k tokens

		"deepseek-coder": {[]float64{0.75, 0.75}, config.ChannelTypeDeepseek}, // 暂定 $0.0015 / 1K tokens
		"deepseek-chat":  {[]float64{0.75, 0.75}, config.ChannelTypeDeepseek}, // 暂定 $0.0015 / 1K tokens

		// 火山方舟默认模型（按 Ark 官方 RMB/百万 token 折算）
		"deepseek-v3.1":                       {[]float64{0.285714286, 0.857142857}, config.ChannelTypeVolcArk},
		"deepseek-v3-1-terminus":              {[]float64{0.285714286, 0.857142857}, config.ChannelTypeVolcArk},
		"deepseek-v3-1-250821":                {[]float64{0.285714286, 0.857142857}, config.ChannelTypeVolcArk},
		"doubao-seed-1.6":                     {[]float64{0.057142857, 0.171428571}, config.ChannelTypeVolcArk},
		"doubao-seed-1-6-250615":              {[]float64{0.057142857, 0.171428571}, config.ChannelTypeVolcArk},
		"doubao-seed-1.6-vision":              {[]float64{0.057142857, 0.171428571}, config.ChannelTypeVolcArk},
		"doubao-seed-1-6-vision-250815":       {[]float64{0.057142857, 0.171428571}, config.ChannelTypeVolcArk},
		"doubao-seed-1.6-flash":               {[]float64{0.010714286, 0.042857143}, config.ChannelTypeVolcArk},
		"doubao-seed-1-6-flash-250828":        {[]float64{0.010714286, 0.042857143}, config.ChannelTypeVolcArk},
		"doubao-seed-1.6-thinking":            {[]float64{0.057142857, 0.171428571}, config.ChannelTypeVolcArk},
		"doubao-seed-1-6-thinking-250715":     {[]float64{0.057142857, 0.171428571}, config.ChannelTypeVolcArk},
		"doubao-1.5-thinking-pro":             {[]float64{0.285714286, 1.142857143}, config.ChannelTypeVolcArk},
		"doubao-1.5-thinking-vision-pro":      {[]float64{0.214285714, 0.642857143}, config.ChannelTypeVolcArk},
		"doubao-1-5-ui-tars":                  {[]float64{0.25, 0.857142857}, config.ChannelTypeVolcArk},
		"doubao-1-5-ui-tars-250428":           {[]float64{0.25, 0.857142857}, config.ChannelTypeVolcArk},
		"doubao-seed-translation":             {[]float64{0.085714286, 0.257142857}, config.ChannelTypeVolcArk},
		"doubao-seed-translation-250915":      {[]float64{0.085714286, 0.257142857}, config.ChannelTypeVolcArk},
		"doubao-1.5-pro-32k":                  {[]float64{0.057142857, 0.142857143}, config.ChannelTypeVolcArk},
		"doubao-1-5-pro-32k-character-250715": {[]float64{0.057142857, 0.142857143}, config.ChannelTypeVolcArk},
		"doubao-1.5-pro-256k":                 {[]float64{0.357142857, 0.642857143}, config.ChannelTypeVolcArk},
		"doubao-1.5-lite-32k":                 {[]float64{0.021428571, 0.042857143}, config.ChannelTypeVolcArk},
		"doubao-pro-32k":                      {[]float64{0.057142857, 0.142857143}, config.ChannelTypeVolcArk},
		"doubao-lite-32k":                     {[]float64{0.021428571, 0.042857143}, config.ChannelTypeVolcArk},
		"kimi-k2":                             {[]float64{0.285714286, 1.142857143}, config.ChannelTypeVolcArk},
		"kimi-k2-250905":                      {[]float64{0.285714286, 1.142857143}, config.ChannelTypeVolcArk},
		"doubao-1.5-vision-pro":               {[]float64{0.214285714, 0.642857143}, config.ChannelTypeVolcArk},
		"doubao-1.5-vision-lite":              {[]float64{0.107142857, 0.321428571}, config.ChannelTypeVolcArk},
		"doubao-vision-pro-32k":               {[]float64{0.214285714, 0.642857143}, config.ChannelTypeVolcArk},
		"doubao-vision-lite-32k":              {[]float64{0.107142857, 0.321428571}, config.ChannelTypeVolcArk},
		"doubao-pro-4k":                       {[]float64{0.057142857, 0.142857143}, config.ChannelTypeVolcArk},
		"doubao-pro-128k":                     {[]float64{0.357142857, 0.642857143}, config.ChannelTypeVolcArk},
		"doubao-pro-256k":                     {[]float64{0.357142857, 0.642857143}, config.ChannelTypeVolcArk},
		"doubao-seedance-1-0-pro":             {[]float64{1.071428571, 1.071428571}, config.ChannelTypeVolcArk}, // ¥15 / 1M tokens（文/图生视频）
		"doubao-seedance-1-0-lite":            {[]float64{0.714285714, 0.714285714}, config.ChannelTypeVolcArk}, // ¥10 / 1M tokens（文/图生视频）
		"doubao-seaweed":                      {[]float64{2.142857143, 2.142857143}, config.ChannelTypeVolcArk}, // ¥30 / 1M tokens（文/图生视频）
		"wan2.1-14b":                          {[]float64{3.571428571, 3.571428571}, config.ChannelTypeVolcArk}, // ¥50 / 1M tokens（文/图生视频）

		"moonshot-v1-8k":   {[]float64{0.8572, 0.8572}, config.ChannelTypeMoonshot}, // ¥0.012 / 1K tokens
		"moonshot-v1-32k":  {[]float64{1.7143, 1.7143}, config.ChannelTypeMoonshot}, // ¥0.024 / 1K tokens
		"moonshot-v1-128k": {[]float64{4.2857, 4.2857}, config.ChannelTypeMoonshot}, // ¥0.06 / 1K tokens

		"open-mistral-7b":       {[]float64{0.125, 0.125}, config.ChannelTypeMistral}, // 0.25$ / 1M tokens	0.25$ / 1M tokens  0.00025$ / 1k tokens
		"open-mixtral-8x7b":     {[]float64{0.35, 0.35}, config.ChannelTypeMistral},   // 0.7$ / 1M tokens	0.7$ / 1M tokens  0.0007$ / 1k tokens
		"mistral-small-latest":  {[]float64{1, 3}, config.ChannelTypeMistral},         // 2$ / 1M tokens	6$ / 1M tokens  0.002$ / 1k tokens
		"mistral-medium-latest": {[]float64{1.35, 4.05}, config.ChannelTypeMistral},   // 2.7$ / 1M tokens	8.1$ / 1M tokens  0.0027$ / 1k tokens
		"mistral-large-latest":  {[]float64{4, 12}, config.ChannelTypeMistral},        // 8$ / 1M tokens	24$ / 1M tokens  0.008$ / 1k tokens
		"mistral-embed":         {[]float64{0.05, 0.05}, config.ChannelTypeMistral},   // 0.1$ / 1M tokens 0.1$ / 1M tokens  0.0001$ / 1k tokens

		// $0.70/$0.80 /1M Tokens 0.0007$ / 1k tokens
		"llama2-70b-4096": {[]float64{0.35, 0.4}, config.ChannelTypeGroq},
		// $0.10/$0.10 /1M Tokens 0.0001$ / 1k tokens
		"llama2-7b-2048": {[]float64{0.05, 0.05}, config.ChannelTypeGroq},
		"gemma-7b-it":    {[]float64{0.05, 0.05}, config.ChannelTypeGroq},
		// $0.27/$0.27 /1M Tokens 0.00027$ / 1k tokens
		"mixtral-8x7b-32768": {[]float64{0.135, 0.135}, config.ChannelTypeGroq},

		// 2.5 元 / 1M tokens 0.0025 / 1k tokens
		"yi-34b-chat-0205": {[]float64{0.1786, 0.1786}, config.ChannelTypeLingyi},
		// 12 元 / 1M tokens 0.012 / 1k tokens
		"yi-34b-chat-200k": {[]float64{0.8571, 0.8571}, config.ChannelTypeLingyi},
		// 	6 元 / 1M tokens 0.006 / 1k tokens
		"yi-vl-plus": {[]float64{0.4286, 0.4286}, config.ChannelTypeLingyi},

		"@cf/stabilityai/stable-diffusion-xl-base-1.0": {[]float64{0, 0}, config.ChannelTypeCloudflareAI},
		"@cf/lykon/dreamshaper-8-lcm":                  {[]float64{0, 0}, config.ChannelTypeCloudflareAI},
		"@cf/bytedance/stable-diffusion-xl-lightning":  {[]float64{0, 0}, config.ChannelTypeCloudflareAI},
		"@cf/qwen/qwen1.5-7b-chat-awq":                 {[]float64{0, 0}, config.ChannelTypeCloudflareAI},
		"@cf/qwen/qwen1.5-14b-chat-awq":                {[]float64{0, 0}, config.ChannelTypeCloudflareAI},
		"@hf/thebloke/deepseek-coder-6.7b-base-awq":    {[]float64{0, 0}, config.ChannelTypeCloudflareAI},
		"@hf/google/gemma-7b-it":                       {[]float64{0, 0}, config.ChannelTypeCloudflareAI},
		"@hf/thebloke/llama-2-13b-chat-awq":            {[]float64{0, 0}, config.ChannelTypeCloudflareAI},
		"@cf/openai/whisper":                           {[]float64{0, 0}, config.ChannelTypeCloudflareAI},
		//$0.50 /1M TOKENS   $1.50/1M TOKENS
		"command-r": {[]float64{0.25, 0.75}, config.ChannelTypeCohere},
		//$3 /1M TOKENS   $15/1M TOKENS
		"command-r-plus": {[]float64{1.5, 7.5}, config.ChannelTypeCohere},

		// 0.065
		"sd3": {[]float64{32.5, 32.5}, config.ChannelTypeStabilityAI},
		// 0.04
		"sd3-turbo": {[]float64{20, 20}, config.ChannelTypeStabilityAI},
		// 0.03
		"stable-image-core": {[]float64{15, 15}, config.ChannelTypeStabilityAI},

		// hunyuan
		"hunyuan-lite":          {[]float64{0, 0}, config.ChannelTypeHunyuan},
		"hunyuan-standard":      {[]float64{0.3214, 0.3571}, config.ChannelTypeHunyuan},
		"hunyuan-standard-256k": {[]float64{1.0714, 4.2857}, config.ChannelTypeHunyuan},
		"hunyuan-pro":           {[]float64{2.1429, 7.1429}, config.ChannelTypeHunyuan},
	}

	arkExtraRatios := map[string]map[string]float64{
		"deepseek-v3.1": {
			config.UsageExtraCache:       0.001214286,
			config.UsageExtraCachedWrite: 0.057142857,
		},
		"deepseek-v3-1-terminus": {
			config.UsageExtraCache:       0.001214286,
			config.UsageExtraCachedWrite: 0.057142857,
		},
		"deepseek-v3-1-250821": {
			config.UsageExtraCache:       0.001214286,
			config.UsageExtraCachedWrite: 0.057142857,
		},
		"doubao-seed-1.6": {
			config.UsageExtraCache:       0.001214286,
			config.UsageExtraCachedWrite: 0.011428571,
		},
		"doubao-seed-1-6-250615": {
			config.UsageExtraCache:       0.001214286,
			config.UsageExtraCachedWrite: 0.011428571,
		},
		"doubao-seed-1.6-vision": {
			config.UsageExtraCache:       0.001214286,
			config.UsageExtraCachedWrite: 0.011428571,
		},
		"doubao-seed-1-6-vision-250815": {
			config.UsageExtraCache:       0.001214286,
			config.UsageExtraCachedWrite: 0.011428571,
		},
		"doubao-seed-1.6-thinking": {
			config.UsageExtraCache:       0.001214286,
			config.UsageExtraCachedWrite: 0.011428571,
		},
		"doubao-seed-1-6-thinking-250715": {
			config.UsageExtraCache:       0.001214286,
			config.UsageExtraCachedWrite: 0.011428571,
		},
		"doubao-seed-1.6-flash": {
			config.UsageExtraCache:       0.001214286,
			config.UsageExtraCachedWrite: 0.002142857,
		},
		"doubao-seed-1-6-flash-250828": {
			config.UsageExtraCache:       0.001214286,
			config.UsageExtraCachedWrite: 0.002142857,
		},
		"doubao-1-5-ui-tars": {
			config.UsageExtraCache:       0.2,
			config.UsageExtraCachedWrite: 0.02125,
		},
		"doubao-1-5-ui-tars-250428": {
			config.UsageExtraCache:       0.2,
			config.UsageExtraCachedWrite: 0.02125,
		},
		"doubao-1.5-pro-32k": {
			config.UsageExtraCache:       0.2,
			config.UsageExtraCachedWrite: 0.02125,
		},
		"doubao-1-5-pro-32k-character-250715": {
			config.UsageExtraCache:       0.2,
			config.UsageExtraCachedWrite: 0.02125,
		},
		"doubao-1.5-lite-32k": {
			config.UsageExtraCache:       0.2,
			config.UsageExtraCachedWrite: 0.056666667,
		},
		"doubao-pro-32k": {
			config.UsageExtraCache:       0.2,
			config.UsageExtraCachedWrite: 0.02125,
		},
		"doubao-lite-32k": {
			config.UsageExtraCache:       0.2,
			config.UsageExtraCachedWrite: 0.056666667,
		},
		"kimi-k2": {
			config.UsageExtraCache:       0.2,
			config.UsageExtraCachedWrite: 0.00425,
		},
		"kimi-k2-250905": {
			config.UsageExtraCache:       0.2,
			config.UsageExtraCachedWrite: 0.00425,
		},
		"doubao-embedding-vision": {
			config.UsageExtraInputImageTokens: 2.571428571,
		},
	}

	ownedByOverrides := map[string]int{
		"deepseek-v3.1":          config.ChannelTypeDeepseek,
		"deepseek-v3-1-terminus": config.ChannelTypeDeepseek,
		"deepseek-v3-1-250821":   config.ChannelTypeDeepseek,
	}

	var prices []*Price

	for model, modelType := range ModelTypes {
		price := &Price{
			Model:       model,
			Type:        TokensPriceType,
			ChannelType: modelType.Type,
			Input:       modelType.Ratio[0],
			Output:      modelType.Ratio[1],
		}
		if ratios, ok := arkExtraRatios[model]; ok {
			jsonData := datatypes.NewJSONType(ratios)
			price.ExtraRatios = &jsonData
		}
		if ownedType, ok := ownedByOverrides[model]; ok {
			price.OwnedByType = ownedType
		}
		price.Normalize()
		prices = append(prices, price)
	}

	var DefaultMJPrice = map[string]float64{
		"mj_imagine":        50,
		"mj_variation":      50,
		"mj_reroll":         50,
		"mj_blend":          50,
		"mj_modal":          50,
		"mj_zoom":           50,
		"mj_shorten":        50,
		"mj_high_variation": 50,
		"mj_low_variation":  50,
		"mj_pan":            50,
		"mj_inpaint":        0,
		"mj_custom_zoom":    0,
		"mj_describe":       25,
		"mj_upscale":        25,
		"mj_swap_face":      25,
		"mj_upload":         0,

		"mj_turbo_imagine":        50,
		"mj_turbo_variation":      50,
		"mj_turbo_reroll":         50,
		"mj_turbo_blend":          50,
		"mj_turbo_modal":          50,
		"mj_turbo_zoom":           50,
		"mj_turbo_shorten":        50,
		"mj_turbo_high_variation": 50,
		"mj_turbo_low_variation":  50,
		"mj_turbo_pan":            50,
		"mj_turbo_inpaint":        0,
		"mj_turbo_custom_zoom":    0,
		"mj_turbo_describe":       25,
		"mj_turbo_upscale":        25,
		"mj_turbo_swap_face":      25,
		"mj_turbo_upload":         0,

		"mj_relax_imagine":        50,
		"mj_relax_variation":      50,
		"mj_relax_reroll":         50,
		"mj_relax_blend":          50,
		"mj_relax_modal":          50,
		"mj_relax_zoom":           50,
		"mj_relax_shorten":        50,
		"mj_relax_high_variation": 50,
		"mj_relax_low_variation":  50,
		"mj_relax_pan":            50,
		"mj_relax_inpaint":        0,
		"mj_relax_custom_zoom":    0,
		"mj_relax_describe":       25,
		"mj_relax_upscale":        25,
		"mj_relax_swap_face":      25,
		"mj_relax_upload":         0,
	}

	for model, mjPrice := range DefaultMJPrice {
		price := &Price{
			Model:       model,
			Type:        TimesPriceType,
			ChannelType: config.ChannelTypeMidjourney,
			Input:       mjPrice,
			Output:      mjPrice,
		}
		price.Normalize()
		prices = append(prices, price)
	}

	var DefaultSunoPrice = map[string]float64{
		"suno_lyrics": 5,
		"chirp-v3-0":  50,
		"chirp-v3-5":  50,
	}
	for model, sunoPrice := range DefaultSunoPrice {
		price := &Price{
			Model:       model,
			Type:        TimesPriceType,
			ChannelType: config.ChannelTypeSuno,
			Input:       sunoPrice,
			Output:      sunoPrice,
		}
		price.Normalize()
		prices = append(prices, price)
	}

	// 每积分折算成人民币：0.3125 RMB；系统基准值为“单位值 * 汇率”得到货币
	// 为在现有展示逻辑下正确显示货币，这里将积分折算为“基准值”：
	// base = credits * (RMB_per_credit / RMBRate)
	// 这样前端使用 $: base*DollarRate、￥: base*RMBRate 可分别得到 USD/人民币价格
	const viduCreditToYuan = 0.3125

	// Vidu 默认价格配置（值为积分，后续换算成人民币）
	var defaultViduCredits = map[string]float64{
		// === 新模型 viduq2-pro ===
		// 图生视频 - viduq2-pro
		"vidu-img2video-viduq2-pro-720p-2s":  3,
		"vidu-img2video-viduq2-pro-1080p-2s": 8,
		"vidu-img2video-viduq2-pro-720p-3s":  4,
		"vidu-img2video-viduq2-pro-1080p-3s": 10,
		"vidu-img2video-viduq2-pro-720p-4s":  5,
		"vidu-img2video-viduq2-pro-1080p-4s": 12,
		"vidu-img2video-viduq2-pro-720p-5s":  6,
		"vidu-img2video-viduq2-pro-1080p-5s": 14,
		"vidu-img2video-viduq2-pro-720p-6s":  7,
		"vidu-img2video-viduq2-pro-1080p-6s": 16,
		"vidu-img2video-viduq2-pro-720p-7s":  8,
		"vidu-img2video-viduq2-pro-1080p-7s": 18,
		"vidu-img2video-viduq2-pro-720p-8s":  9,
		"vidu-img2video-viduq2-pro-1080p-8s": 20,
		// 首尾帧 - viduq2-pro
		"vidu-start-end2video-viduq2-pro-720p-2s":  3,
		"vidu-start-end2video-viduq2-pro-1080p-2s": 8,
		"vidu-start-end2video-viduq2-pro-720p-3s":  4,
		"vidu-start-end2video-viduq2-pro-1080p-3s": 10,
		"vidu-start-end2video-viduq2-pro-720p-4s":  5,
		"vidu-start-end2video-viduq2-pro-1080p-4s": 12,
		"vidu-start-end2video-viduq2-pro-720p-5s":  6,
		"vidu-start-end2video-viduq2-pro-1080p-5s": 14,
		"vidu-start-end2video-viduq2-pro-720p-6s":  7,
		"vidu-start-end2video-viduq2-pro-1080p-6s": 16,
		"vidu-start-end2video-viduq2-pro-720p-7s":  8,
		"vidu-start-end2video-viduq2-pro-1080p-7s": 18,
		"vidu-start-end2video-viduq2-pro-720p-8s":  9,
		"vidu-start-end2video-viduq2-pro-1080p-8s": 20,

		// === 新模型 viduq2-turbo ===
		// 图生视频 - viduq2-turbo
		"vidu-img2video-viduq2-turbo-720p-2s":  1,
		"vidu-img2video-viduq2-turbo-1080p-2s": 5,
		"vidu-img2video-viduq2-turbo-720p-3s":  2,
		"vidu-img2video-viduq2-turbo-1080p-3s": 6,
		"vidu-img2video-viduq2-turbo-720p-4s":  3,
		"vidu-img2video-viduq2-turbo-1080p-4s": 7,
		"vidu-img2video-viduq2-turbo-720p-5s":  4,
		"vidu-img2video-viduq2-turbo-1080p-5s": 8,
		"vidu-img2video-viduq2-turbo-720p-6s":  5,
		"vidu-img2video-viduq2-turbo-1080p-6s": 9,
		"vidu-img2video-viduq2-turbo-720p-7s":  6,
		"vidu-img2video-viduq2-turbo-1080p-7s": 10,
		"vidu-img2video-viduq2-turbo-720p-8s":  7,
		"vidu-img2video-viduq2-turbo-1080p-8s": 11,
		// 首尾帧 - viduq2-turbo
		"vidu-start-end2video-viduq2-turbo-720p-2s":  1,
		"vidu-start-end2video-viduq2-turbo-1080p-2s": 5,
		"vidu-start-end2video-viduq2-turbo-720p-3s":  2,
		"vidu-start-end2video-viduq2-turbo-1080p-3s": 6,
		"vidu-start-end2video-viduq2-turbo-720p-4s":  3,
		"vidu-start-end2video-viduq2-turbo-1080p-4s": 7,
		"vidu-start-end2video-viduq2-turbo-720p-5s":  4,
		"vidu-start-end2video-viduq2-turbo-1080p-5s": 8,
		"vidu-start-end2video-viduq2-turbo-720p-6s":  5,
		"vidu-start-end2video-viduq2-turbo-1080p-6s": 9,
		"vidu-start-end2video-viduq2-turbo-720p-7s":  6,
		"vidu-start-end2video-viduq2-turbo-1080p-7s": 10,
		"vidu-start-end2video-viduq2-turbo-720p-8s":  7,
		"vidu-start-end2video-viduq2-turbo-1080p-8s": 11,

		// === 模型 viduq1 ===
		// 图生视频 - viduq1
		"vidu-img2video-viduq1-1080p-5s": 8,
		// 参考生视频 - viduq1
		"vidu-reference2video-viduq1-1080p-5s": 8,
		// 首尾帧 - viduq1
		"vidu-start-end2video-viduq1-1080p-5s": 8,
		// 文生视频 - viduq1
		"vidu-text2video-viduq1-1080p-5s":       8,
		"vidu-text2video-viduq1-1080p-5s-anime": 8,
		// 参考生图 - viduq1
		"vidu-reference2image-viduq1": 2,

		// === 模型 viduq1-classic ===
		// 图生视频 - viduq1-classic
		"vidu-img2video-viduq1-classic-1080p-5s": 8,
		// 首尾帧 - viduq1-classic
		"vidu-start-end2video-viduq1-classic-1080p-5s": 8,

		// === 模型 vidu2.0 ===
		// 图生视频 - vidu2.0
		"vidu-img2video-vidu2.0-360p-4s":  2,
		"vidu-img2video-vidu2.0-720p-4s":  4,
		"vidu-img2video-vidu2.0-1080p-4s": 10,
		"vidu-img2video-vidu2.0-720p-8s":  10,
		// 参考生视频 - vidu2.0
		"vidu-reference2video-vidu2.0-360p-4s":  8,
		"vidu-reference2video-vidu2.0-720p-4s":  8,
		"vidu-reference2video-vidu2.0-1080p-4s": 20,
		"vidu-reference2video-vidu2.0-720p-8s":  20,
		// 首尾帧 - vidu2.0
		"vidu-start-end2video-vidu2.0-360p-4s":  2,
		"vidu-start-end2video-vidu2.0-720p-4s":  4,
		"vidu-start-end2video-vidu2.0-1080p-4s": 10,
		"vidu-start-end2video-vidu2.0-720p-8s":  10,
		// 文生视频（扩展支持） - vidu2.0
		"vidu-text2video-vidu2.0-720p-4s":       4,
		"vidu-text2video-vidu2.0-720p-8s":       10,
		"vidu-text2video-vidu2.0-720p-4s-anime": 4,
		"vidu-text2video-vidu2.0-720p-8s-anime": 10,

		// === 模型 vidu1.5 ===
		// 图生视频 - vidu1.5
		"vidu-img2video-vidu1.5-360p-4s":  4,
		"vidu-img2video-vidu1.5-720p-4s":  10,
		"vidu-img2video-vidu1.5-1080p-4s": 20,
		"vidu-img2video-vidu1.5-720p-8s":  20,
		// 参考生视频 - vidu1.5
		"vidu-reference2video-vidu1.5-360p-4s":  8,
		"vidu-reference2video-vidu1.5-720p-4s":  20,
		"vidu-reference2video-vidu1.5-1080p-4s": 40,
		"vidu-reference2video-vidu1.5-720p-8s":  40,
		// 首尾帧 - vidu1.5
		"vidu-start-end2video-vidu1.5-360p-4s":  4,
		"vidu-start-end2video-vidu1.5-720p-4s":  10,
		"vidu-start-end2video-vidu1.5-1080p-4s": 20,
		"vidu-start-end2video-vidu1.5-720p-8s":  20,
		// 文生视频 - vidu1.5 通用风格
		"vidu-text2video-vidu1.5-360p-4s":  4,
		"vidu-text2video-vidu1.5-720p-4s":  10,
		"vidu-text2video-vidu1.5-1080p-4s": 20,
		"vidu-text2video-vidu1.5-720p-8s":  20,
		// 文生视频 - vidu1.5 动漫风格
		"vidu-text2video-vidu1.5-360p-4s-anime":  4,
		"vidu-text2video-vidu1.5-720p-4s-anime":  10,
		"vidu-text2video-vidu1.5-1080p-4s-anime": 20,
		"vidu-text2video-vidu1.5-720p-8s-anime":  20,
	}
	for model, credits := range defaultViduCredits {
		// 将“积分”转为系统基准值，保证前端展示的 $ 与 ￥计算正确
		base := credits * (viduCreditToYuan / RMBRate)
		price := &Price{
			Model:       model,
			Type:        TimesPriceType,
			ChannelType: config.ChannelTypeVidu,
			Input:       base,
			Output:      base,
		}
		price.Normalize()
		prices = append(prices, price)
	}

	var defaultVolcArkImagePrice = map[string]float64{
		"doubao-seedream-4.0":            0.2,
		"doubao-seedream-3.0-t2i":        0.259,
		"doubao-seedream-3-0-t2i-250415": 0.259,
		"doubao-seededit-3.0-i2i":        0.3,
		"doubao-seededit-3-0-i2i-250628": 0.3,
	}

	for model, pricePerImage := range defaultVolcArkImagePrice {
		base := pricePerImage / RMBRate
		price := &Price{
			Model:       model,
			Type:        TimesPriceType,
			ChannelType: config.ChannelTypeVolcArk,
			Input:       base,
			Output:      base,
		}
		price.Normalize()
		prices = append(prices, price)
	}

	// MiniMax 视频默认价格（单位：人民币）
	var defaultMiniMaxVideoPrice = map[string]float64{
		// 文生视频（PPInfra 官方报价）
		"minimax-text2video-minimax-hailuo-02-512p-6s":  0.6,
		"minimax-text2video-minimax-hailuo-02-512p-10s": 1.0,
		"minimax-text2video-minimax-hailuo-02-768p-6s":  2.0,
		"minimax-text2video-minimax-hailuo-02-768p-10s": 4.0,
		"minimax-text2video-minimax-hailuo-02-1080p-6s": 6.0,

		// 图生视频（暂与文生视频保持一致，可按需调整）
		"minimax-image2video-minimax-hailuo-02-512p-6s":  0.6,
		"minimax-image2video-minimax-hailuo-02-512p-10s": 1.0,
		"minimax-image2video-minimax-hailuo-02-768p-6s":  2.0,
		"minimax-image2video-minimax-hailuo-02-768p-10s": 4.0,
		"minimax-image2video-minimax-hailuo-02-1080p-6s": 6.0,

		// 首尾帧视频（默认与图生一致，便于覆盖）
		"minimax-start-end2video-minimax-hailuo-02-512p-6s":  0.6,
		"minimax-start-end2video-minimax-hailuo-02-512p-10s": 1.0,
		"minimax-start-end2video-minimax-hailuo-02-768p-6s":  2.0,
		"minimax-start-end2video-minimax-hailuo-02-768p-10s": 4.0,
		"minimax-start-end2video-minimax-hailuo-02-1080p-6s": 6.0,
	}

	for model, miniMaxPrice := range defaultMiniMaxVideoPrice {
		base := miniMaxPrice / RMBRate
		price := &Price{
			Model:       model,
			Type:        TimesPriceType,
			ChannelType: config.ChannelTypeMiniMax,
			Input:       base,
			Output:      base,
		}
		price.Normalize()
		prices = append(prices, price)
	}

	var defaultSoraVideoPrice = map[string]map[string]float64{
		"sora-2": {
			"1280x720": 0.10,
			"720x1280": 0.10,
		},
		"sora-2-pro": {
			"1280x720":  0.30,
			"720x1280":  0.30,
			"1792x1024": 0.50,
			"1024x1792": 0.50,
		},
	}

	for model, resolutionPrices := range defaultSoraVideoPrice {
		for resolution, usdPrice := range resolutionPrices {
			base := usdPrice / DollarRate
			price := &Price{
				Model:       model + "-" + resolution,
				Type:        TokensPriceType,
				ChannelType: config.ChannelTypeOpenAI,
				Input:       base,
			}
			price.Normalize()
			prices = append(prices, price)
		}
	}

	// 可灵AI 默认价格配置 (按次计费)
	var DefaultKlingPrice = map[string]float64{
		// 视频生成模型
		"kling-video_kling-v1_std_5":    1,   // V1 标准模式 5秒，单价 1 元
		"kling-video_kling-v1_std_10":   2,   // V1 标准模式 10秒，单价 2 元
		"kling-video_kling-v1_pro_5":    3.5, // V1 高品质 5秒，单价 3.5 元
		"kling-video_kling-v1_pro_10":   7,   // V1 高品质 10秒，单价 7 元
		"kling-video_kling-v1.5_std_5":  2,   // V1.5 标准模式 5秒，单价 2 元
		"kling-video_kling-v1.5_std_10": 4,   // V1.5 标准模式 10秒，单价 4 元
		"kling-video_kling-v1.5_pro_5":  3.5, // V1.5 高品质 5秒，单价 3.5 元
		"kling-video_kling-v1.5_pro_10": 7,   // V1.5 高品质 10秒，单价 7 元

		// 别名：v1-10 映射到 v1 款 10s（std/pro）
		"kling-video_kling-v1-10_std_5":  1,   // 等价 v1_std_5
		"kling-video_kling-v1-10_std_10": 2,   // 等价 v1_std_10
		"kling-video_kling-v1-10_pro_5":  3.5, // 等价 v1_pro_5
		"kling-video_kling-v1-10_pro_10": 7,   // 等价 v1_pro_10

		// 别名：v1-5 映射到 v1.5 款（std/pro）
		"kling-video_kling-v1-5_std_5":  2,   // 等价 v1.5_std_5
		"kling-video_kling-v1-5_std_10": 4,   // 等价 v1.5_std_10
		"kling-video_kling-v1-5_pro_5":  3.5, // 等价 v1.5_pro_5
		"kling-video_kling-v1-5_pro_10": 7,   // 等价 v1.5_pro_10

		"kling-video_kling-v1-6_std_5":         2,   // V1.6 标准模式 5秒，单价 2 元
		"kling-video_kling-v1-6_std_10":        4,   // V1.6 标准模式 10秒，单价 4 元
		"kling-video_kling-v1-6_pro_5":         3.5, // V1.6 高品质 5秒，单价 3.5 元
		"kling-video_kling-v1-6_pro_10":        7,   // V1.6 高品质 10秒，单价 7 元
		"kling-video_kling-v2-master_std_5":    10,  // V2.0 大师 5秒，单价 10 元
		"kling-video_kling-v2-master_std_10":   20,  // V2.0 大师 10秒，单价 20 元
		"kling-video_kling-v2-master_pro_5":    10,  // V2.0 大师 5秒，单价 10 元（未区分 pro/std，取一致）
		"kling-video_kling-v2-master_pro_10":   20,  // V2.0 大师 10秒，单价 20 元（未区分 pro/std，取一致）
		"kling-video_kling-v2-1-master_std_5":  10,  // V2.1 大师 5秒，单价 10 元
		"kling-video_kling-v2-1-master_std_10": 20,  // V2.1 大师 10秒，单价 20 元
		"kling-video_kling-v2-1-master_pro_5":  10,  // V2.1 大师 5秒，单价 10 元
		"kling-video_kling-v2-1-master_pro_10": 20,  // V2.1 大师 10秒，单价 20 元

		// 图像生成模型
		"kling-image_kling-v1_std":          0.025, // 图片 V1.0 文/图 0.025 元
		"kling-image_kling-v1_pro":          0.025, // 图片 V1.0 文/图 0.025 元
		"kling-image_kling-v1.5_std":        0.1,   // 图片 V1.5 文生图 0.1 元
		"kling-image_kling-v1.5_pro":        0.2,   // 图片 V1.5 图生图 0.2 元
		"kling-image_kling-v1-6_std":        0.1,   // 暂按 V1.5 文档价 0.1 元（无官方明细时取近似）
		"kling-image_kling-v1-6_pro":        0.2,   // 暂按 V1.5 图像价 0.2 元（无官方明细时取近似）
		"kling-image_kling-v2-master_std":   0.2,   // 图片 V2.0-master 未单列，取 0.2 元近似
		"kling-image_kling-v2-master_pro":   0.2,   // 同上
		"kling-image_kling-v2-1-master_std": 0.1,   // 图片 V2.1 文生图 0.1 元
		"kling-image_kling-v2-1-master_pro": 0.2,   // 假定图生图 0.2 元

		// 新增的图像生成模型（按照官方接口命名）
		"kling-image_kling-v1":     0.025, // 图片 V1.0 文/图 0.025 元
		"kling-image_kling-v1-5":   0.1,   // 图片 V1.5 文生图 0.1 元（默认）
		"kling-image_kling-v2":     0.1,   // 图片 V2.0 文生图 0.1 元
		"kling-image_kling-v2-new": 0.2,   // 图片 V2.0-new 图生图 0.2 元
		"kling-image_kling-v2-1":   0.1,   // 图片 V2.1 文生图 0.1 元

		// 多图参考生图模型
		"kling-multi-image2image_kling-v2": 0.4, // 多图参考生图 V2，单价 0.4 元

		// 虚拟试穿模型（按“资源包次数”对齐）
		"kling-try-on_kling-v1_std":   2,   // std 2 元
		"kling-try-on_kling-v1_pro":   3.5, // pro 3.5 元
		"kling-try-on_kling-v1.5_std": 2,   // std 2 元
		"kling-try-on_kling-v1.5_pro": 3.5, // pro 3.5 元
		"kling-try-on_kling-v1-6_std": 2,   // std 2 元
		"kling-try-on_kling-v1-6_pro": 3.5, // pro 3.5 元

		// 为了与现有实现兼容，也添加一些简化的模型名称
		"kling-v1":          1,  // 兼容简化键：V1 标准 5秒 1 元
		"kling-v1.5":        2,  // 兼容简化键：V1.5 标准 5秒 2 元
		"kling-v1-6":        2,  // 兼容简化键：V1.6 标准 5秒 2 元
		"kling-v2-master":   10, // 兼容简化键：V2.0 大师 5秒 10 元
		"kling-v2-1-master": 10, // 兼容简化键：V2.1 大师 5秒 10 元

		// 多图参考生视频模型（按“资源包次数”对齐）
		"kling-multi-image2video_kling-v1-6_std_5":  2,   // std 5s 2 元
		"kling-multi-image2video_kling-v1-6_std_10": 4,   // std 10s 4 元
		"kling-multi-image2video_kling-v1-6_pro_5":  3.5, // pro 5s 3.5 元
		"kling-multi-image2video_kling-v1-6_pro_10": 7,   // pro 10s 7 元

		// 多模态视频编辑模型（按“资源包次数”对齐）
		"kling-multi-elements_kling-v1-6_std_5":  3,  // std 5s 3 元
		"kling-multi-elements_kling-v1-6_std_10": 6,  // std 10s 6 元
		"kling-multi-elements_kling-v1-6_pro_5":  5,  // pro 5s 5 元
		"kling-multi-elements_kling-v1-6_pro_10": 10, // pro 10s 10 元

		// 多模态视频编辑辅助功能（每次操作收费）
		"kling-init-selection":    5, // 初始化待编辑视频，单价 5 元
		"kling-add-selection":     2, // 增加视频选区，单价 2 元
		"kling-delete-selection":  2, // 删减视频选区，单价 2 元
		"kling-clear-selection":   1, // 清除视频选区，单价 1 元
		"kling-preview-selection": 3, // 预览已选区视频，单价 3 元
	}

	for model, klingPrice := range DefaultKlingPrice {
		// 将人民币价格折算为系统基准值，保证前端 $ / ￥ 展示正确
		base := klingPrice / RMBRate
		price := &Price{
			Model:       model,
			Type:        TimesPriceType,
			ChannelType: config.ChannelTypeKling,
			Input:       base,
			Output:      base,
		}
		price.Normalize()
		prices = append(prices, price)
	}

	return prices
}

func GetDefaultExtraRatio() string {
	return `{"gpt-4o-audio-preview":{"input_audio_tokens":40,"output_audio_tokens":20},"gpt-4o-audio-preview-2024-10-01":{"input_audio_tokens":40,"output_audio_tokens":20},"gpt-4o-audio-preview-2024-12-17":{"input_audio_tokens":16,"output_audio_tokens":8},"gpt-4o-mini-audio-preview":{"input_audio_tokens":67,"output_audio_tokens":34},"gpt-4o-mini-audio-preview-2024-12-17":{"input_audio_tokens":67,"output_audio_tokens":34},"gpt-4o-realtime-preview":{"input_audio_tokens":20,"output_audio_tokens":10},"gpt-4o-realtime-preview-2024-10-01":{"input_audio_tokens":20,"output_audio_tokens":10},"gpt-4o-realtime-preview-2024-12-17":{"input_audio_tokens":8,"output_audio_tokens":4},"gpt-4o-mini-realtime-preview":{"input_audio_tokens":17,"output_audio_tokens":8.4},"gpt-4o-mini-realtime-preview-2024-12-17":{"input_audio_tokens":17,"output_audio_tokens":8.4},"gemini-2.5-flash-preview-04-17":{"reasoning_tokens":5.833},"gpt-image-1":{"input_text_tokens": 0.5}}`

}
