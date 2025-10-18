package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"one-api/common/config"
	"one-api/common/logger"
	"one-api/common/utils"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/spf13/viper"
)

// PricingInstance is the Pricing instance
var PricingInstance *Pricing

type PriceUpdateMode string

const (
	PriceUpdateModeSystem    PriceUpdateMode = "system"
	PriceUpdateModeAdd       PriceUpdateMode = "add"
	PriceUpdateModeOverwrite PriceUpdateMode = "overwrite"
	PriceUpdateModeUpdate    PriceUpdateMode = "update"
)

// Pricing is a struct that contains the pricing data
type Pricing struct {
	sync.RWMutex
	Prices map[string]*Price `json:"models"`
	Match  []string          `json:"-"`
}

type BatchPrices struct {
	Models []string `json:"models" binding:"required"`
	Price  Price    `json:"price" binding:"required"`
}

// NewPricing creates a new Pricing instance
func NewPricing() {
	logger.SysLog("Initializing Pricing")
	logger.SysLog("Update Price Mode:" + viper.GetString("auto_price_updates_mode"))
	PricingInstance = &Pricing{
		Prices: make(map[string]*Price),
		Match:  make([]string, 0),
	}

    err := PricingInstance.Init()

	if err != nil {
		logger.SysError("Failed to initialize Pricing:" + err.Error())
		return
	}

    // 初始化时，需要检测是否有更新
    if viper.GetString("auto_price_updates_mode") == "system" && (viper.GetBool("auto_price_updates") || len(PricingInstance.Prices) == 0) {
        logger.SysLog("Checking for pricing updates")
        prices := GetDefaultPrice()
        PricingInstance.SyncPricing(prices, "system")
        logger.SysLog("Pricing initialized")
    }

    // 无论是否启用自动更新，都尝试“仅新增”一次系统内置价格，确保新加的精确计费模型（例如 minimaxi 视频组合）被补齐；不覆盖已有配置
    // 这一步是幂等的（SyncPriceWithoutOverwrite），不会影响已存在或 locked 的价格
    func() {
        defer func() { recover() }()
        _ = PricingInstance.SyncPricing(GetDefaultPrice(), string(PriceUpdateModeAdd))
    }()
}

// initializes the Pricing instance
func (p *Pricing) Init() error {
	prices, err := GetAllPrices()
	if err != nil {
		return err
	}

	if len(prices) == 0 {
		return nil
	}

	newPrices := make(map[string]*Price)
	newMatch := make(map[string]bool)

	for _, price := range prices {
		price.Normalize()
		newPrices[price.Model] = price
		if strings.HasSuffix(price.Model, "*") {
			if _, ok := newMatch[price.Model]; !ok {
				newMatch[price.Model] = true
			}
		}
	}

	var newMatchList []string
	for match := range newMatch {
		newMatchList = append(newMatchList, match)
	}

	p.Lock()
	defer p.Unlock()

	p.Prices = newPrices
	p.Match = newMatchList

	return nil
}

// GetPrice returns the price of a model
func (p *Pricing) GetPrice(modelName string) *Price {
	p.RLock()
	defer p.RUnlock()

	if price, ok := p.Prices[modelName]; ok {
		price.Normalize()
		return price
	}

	// 兼容旧版 Vidu 命名（vidu-<action>-<model>-<duration>s-<resolution>）
	if normalized, changed := normalizeLegacyViduModelName(modelName); changed {
		if price, ok := p.Prices[normalized]; ok {
			return price
		}
		modelName = normalized
	}

	matchModel := utils.GetModelsWithMatch(&p.Match, modelName)
	if price, ok := p.Prices[matchModel]; ok {
		price.Normalize()
		return price
	}

	fallback := &Price{
		Model:       modelName,
		Type:        TokensPriceType,
		ChannelType: config.ChannelTypeUnknown,
		Input:       DefaultPrice,
		Output:      DefaultPrice,
	}
	fallback.Normalize()

	return fallback
}

func normalizeLegacyViduModelName(modelName string) (string, bool) {
	// 旧格式：vidu-<action>-<model>-<duration>s-<resolution>(-style)
	// 新格式：vidu-<action>-<model>-<resolution>-<duration>s(-style)
	if !strings.HasPrefix(modelName, "vidu-") {
		return "", false
	}

	segments := strings.Split(modelName, "-")
	if len(segments) < 5 {
		return "", false
	}

	durationSeg := segments[3]
	resolutionSeg := segments[4]

	if !strings.HasSuffix(durationSeg, "s") {
		return "", false
	}

	resolutionSeg = strings.ToLower(resolutionSeg)
	if !(strings.HasSuffix(resolutionSeg, "p")) {
		return "", false
	}

	segments[3], segments[4] = resolutionSeg, durationSeg

	newName := strings.Join(segments, "-")
	if newName == modelName {
		return "", false
	}

	return newName, true
}

func (p *Pricing) GetAllPrices() map[string]*Price {
	return p.Prices
}

func (p *Pricing) GetAllPricesList() []*Price {
	var prices []*Price
	for _, price := range p.Prices {
		price.Normalize()
		prices = append(prices, price)
	}

	return prices
}

func (p *Pricing) updateRawPrice(modelName string, price *Price) error {
	if _, ok := p.Prices[modelName]; !ok {
		return errors.New("model not found")
	}

	if strings.TrimSpace(price.Model) == "" {
		price.Model = modelName
	}

	if _, ok := p.Prices[price.Model]; modelName != price.Model && ok {
		return errors.New("model names cannot be duplicated")
	}

	if err := p.deleteRawPrice(modelName); err != nil {
		return err
	}

	return price.Insert()
}

// UpdatePrice updates the price of a model
func (p *Pricing) UpdatePrice(modelName string, price *Price) error {

	if err := p.updateRawPrice(modelName, price); err != nil {
		return err
	}

	err := p.Init()

	return err
}

func (p *Pricing) addRawPrice(price *Price) error {
	if _, ok := p.Prices[price.Model]; ok {
		return errors.New("model already exists")
	}

	return price.Insert()
}

// AddPrice adds a new price to the Pricing instance
func (p *Pricing) AddPrice(price *Price) error {
	if err := p.addRawPrice(price); err != nil {
		return err
	}

	err := p.Init()

	return err
}

func (p *Pricing) deleteRawPrice(modelName string) error {
	item, ok := p.Prices[modelName]
	if !ok {
		return errors.New("model not found")
	}

	return item.Delete()
}

// DeletePrice deletes a price from the Pricing instance
func (p *Pricing) DeletePrice(modelName string) error {
	if err := p.deleteRawPrice(modelName); err != nil {
		return err
	}

	err := p.Init()

	return err
}

// SyncPricing syncs the pricing data
func (p *Pricing) SyncPricing(pricing []*Price, mode string) error {
	logger.SysLog("prices update mode：" + mode)
	var err error
	switch mode {
	case string(PriceUpdateModeSystem):
		err = p.SyncPriceWithoutOverwrite(pricing)
		return err
	case string(PriceUpdateModeUpdate):
		err = p.SyncPriceOnlyUpdate(pricing)
		return err
	case string(PriceUpdateModeOverwrite):
		err = p.SyncPriceWithOverwrite(pricing)
		return err
	case string(PriceUpdateModeAdd):
		err = p.SyncPriceWithoutOverwrite(pricing)
		return err
	default:
		err = p.SyncPriceWithoutOverwrite(pricing)
		return err
	}
}

func UpdatePriceByPriceService() error {
	updatePriceMode := viper.GetString("auto_price_updates_mode")
	if updatePriceMode == string(PriceUpdateModeSystem) {
		// 使用程序内置更新
		return nil
	}
	prices, err := GetPriceByPriceService()
	if err != nil {
		return err
	}
	if updatePriceMode == string(PriceUpdateModeAdd) {
		// 仅仅新增
		p := &Pricing{
			Prices: make(map[string]*Price),
			Match:  make([]string, 0),
		}
		err := p.Init()
		if err != nil {
			logger.SysError("Failed to initialize Pricing:" + err.Error())
			return err
		}
		err = p.SyncPriceWithoutOverwrite(prices)
		if err != nil {
			return err
		}
		return nil
	}
	if updatePriceMode == string(PriceUpdateModeOverwrite) {
		// 覆盖所有
		p := &Pricing{
			Prices: make(map[string]*Price),
			Match:  make([]string, 0),
		}
		err := p.Init()
		if err != nil {
			logger.SysError("Failed to initialize Pricing:" + err.Error())
			return err
		}
		err = p.SyncPriceWithOverwrite(prices)
		if err != nil {
			return err
		}
		return nil
	}
	if updatePriceMode == string(PriceUpdateModeUpdate) {
		// 只更新现有数据
		p := &Pricing{
			Prices: make(map[string]*Price),
			Match:  make([]string, 0),
		}
		err := p.Init()
		if err != nil {
			logger.SysError("Failed to initialize Pricing:" + err.Error())
			return err
		}
		err = p.SyncPriceOnlyUpdate(prices)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("更新模式错误，更新模式仅能选择：add、overwrite、system，详见配置文件auto_price_updates_mode部分的说明")
}

// GetPriceByPriceService 只插入系统没有的数据
func GetPriceByPriceService() ([]*Price, error) {
	api := viper.GetString("update_price_service")
	if api == "" {
		return nil, errors.New("update_price_service is not configured")
	}
	logger.SysLog("Start Update Price,Prices Service URL：" + api)
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(api)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch prices from service: %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}
	var result struct {
		Data []*Price `json:"data"`
	}
	// 尝试解析为带data字段的格式
	if err := json.Unmarshal(body, &result); err == nil && len(result.Data) > 0 {
		logger.SysLog(fmt.Sprintf("成功解析带data字段的数据，共获取到 %d 个价格配置", len(result.Data)))
		return result.Data, nil
	}
	// 如果不是带data字段的格式，尝试直接解析为数组
	var prices []*Price
	if err := json.Unmarshal(body, &prices); err != nil {
		return nil, fmt.Errorf("failed to parse price data: %v", err)
	}
	logger.SysLog(fmt.Sprintf("成功解析数组格式数据，共获取到 %d 个价格配置", len(prices)))
	return prices, nil
}

// SyncPriceWithOverwrite 删除系统所有数据并插入所有查询到的新数据 不含lock的数据
func (p *Pricing) SyncPriceWithOverwrite(pricing []*Price) error {
	tx := DB.Begin()
	logger.SysLog(fmt.Sprintf("系统内已有价格配置 %d 个(包含locked价格)", len(p.Prices)))
	if err := DeleteAllPricesNotLock(tx); err != nil {
		tx.Rollback()
		return err
	}

	defaultPricing := GetDefaultPrice()
	defaultPriceMap := make(map[string]*Price, len(defaultPricing))
	for _, defaultPrice := range defaultPricing {
		defaultPriceMap[defaultPrice.Model] = defaultPrice
	}

	incomingModels := make(map[string]struct{}, len(pricing))
	for _, price := range pricing {
		incomingModels[price.Model] = struct{}{}
		if price.ChannelType == config.ChannelTypeUnknown {
			if defaultPrice, ok := defaultPriceMap[price.Model]; ok {
				price.ChannelType = defaultPrice.ChannelType
				if price.Type == "" {
					price.Type = defaultPrice.Type
				}
				if price.Type == TimesPriceType && price.Output == 0 {
					price.Output = price.Input
				}
			}
		}
		price.Normalize()
	}

	for _, defaultPrice := range defaultPricing {
		if _, ok := incomingModels[defaultPrice.Model]; ok {
			continue
		}
		clone := *defaultPrice
		clone.Normalize()
		pricing = append(pricing, &clone)
	}

	var newPrices []*Price
	for _, price := range pricing {
		if existing, ok := p.Prices[price.Model]; !ok {
			newPrices = append(newPrices, price)
		} else if !existing.Locked {
			newPrices = append(newPrices, price)
		}
	}

	if len(newPrices) == 0 {
		tx.Commit()
		return p.Init()
	}

	if err := InsertPrices(tx, newPrices); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	logger.SysLog(fmt.Sprintf("本次修改加新增 %d 个价格配置", len(newPrices)))
	return p.Init()
}

// SyncPriceOnlyUpdate 只更新系统现有的数据 不含lock的数据
func (p *Pricing) SyncPriceOnlyUpdate(pricing []*Price) error {
	tx := DB.Begin()
	logger.SysLog(fmt.Sprintf("系统内已有价格配置 %d 个(包含locked价格)", len(p.Prices)))
	var newPrices []*Price
	var newPricesName []string
	//系统内存在并且非lock的模型价格加入new price
	for _, price := range pricing {
		price.Normalize()
		if p, ok := p.Prices[price.Model]; ok && !p.Locked {
			newPrices = append(newPrices, price)
			newPricesName = append(newPricesName, price.Model)
		}
	}
	if len(newPrices) == 0 {
		return nil
	}
	logger.SysLog(fmt.Sprintf("系统内需要更新 %d 个模型价格", len(newPrices)))
	// 删除需要更新的模型价格
	err := DeletePricesByModelNameAndNotLock(tx, newPricesName)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = InsertPrices(tx, newPrices)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	logger.SysLog(fmt.Sprintf("本次更新修改 %d 个价格配置", len(newPrices)))
	return p.Init()
}

// SyncPriceWithoutOverwrite 只插入系统没有的数据，并修正历史缺失的渠道类型
func (p *Pricing) SyncPriceWithoutOverwrite(pricing []*Price) error {
	var newPrices []*Price
	type channelTypeUpdate struct {
		Model       string
		ChannelType int
	}
	type ownedByTypeUpdate struct {
		Model       string
		OwnedByType int
	}
	var channelTypeUpdates []channelTypeUpdate
	var ownedByTypeUpdates []ownedByTypeUpdate

	logger.SysLog(fmt.Sprintf("系统内已有价格配置 %d 个", len(p.Prices)))
	for _, price := range pricing {
		price.Normalize()
		if existing, ok := p.Prices[price.Model]; !ok {
			// 系统内不存在的模型直接新增
			newPrices = append(newPrices, price)
		} else {
			existing.Normalize()
			if existing.ChannelType == config.ChannelTypeUnknown && price.ChannelType != config.ChannelTypeUnknown {
				// 针对旧数据没有渠道类型的情况做一次补齐
				channelTypeUpdates = append(channelTypeUpdates, channelTypeUpdate{Model: price.Model, ChannelType: price.ChannelType})
			}
			if price.OwnedByType != config.ChannelTypeUnknown && price.OwnedByType != existing.OwnedByType {
				ownedByTypeUpdates = append(ownedByTypeUpdates, ownedByTypeUpdate{Model: price.Model, OwnedByType: price.OwnedByType})
			}
		}
	}

	if len(newPrices) == 0 && len(channelTypeUpdates) == 0 && len(ownedByTypeUpdates) == 0 {
		return nil
	}

	tx := DB.Begin()

	if len(newPrices) > 0 {
		if err := InsertPrices(tx, newPrices); err != nil {
			tx.Rollback()
			return err
		}
		logger.SysLog(fmt.Sprintf("本次新增 %d 个价格配置", len(newPrices)))
	}

	if len(channelTypeUpdates) > 0 {
		for _, item := range channelTypeUpdates {
			if err := tx.Model(&Price{}).Where("model = ?", item.Model).Update("channel_type", item.ChannelType).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
		logger.SysLog(fmt.Sprintf("本次修正渠道类型 %d 个模型", len(channelTypeUpdates)))
	}

	if len(ownedByTypeUpdates) > 0 {
		for _, item := range ownedByTypeUpdates {
			if err := tx.Model(&Price{}).Where("model = ?", item.Model).Update("owned_by_type", item.OwnedByType).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
		logger.SysLog(fmt.Sprintf("本次修正品牌归属 %d 个模型", len(ownedByTypeUpdates)))
	}

	tx.Commit()

	return p.Init()
}

// BatchDeletePrices deletes the prices of multiple models
func (p *Pricing) BatchDeletePrices(models []string) error {
	tx := DB.Begin()

	err := DeletePrices(tx, models)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	p.Lock()
	defer p.Unlock()

	for _, model := range models {
		delete(p.Prices, model)
	}

	return nil
}

func (p *Pricing) BatchSetPrices(batchPrices *BatchPrices, originalModels []string) error {
	// 查找需要删除的model
	var deletePrices []string
	var addPrices []*Price
	var updatePrices []string

	for _, model := range originalModels {
		if !utils.Contains(model, batchPrices.Models) {
			deletePrices = append(deletePrices, model)
		} else {
			updatePrices = append(updatePrices, model)
		}
	}

	for _, model := range batchPrices.Models {
		if !utils.Contains(model, originalModels) {
			addPrice := batchPrices.Price
			addPrice.Model = model
			addPrices = append(addPrices, &addPrice)
		}
	}

	tx := DB.Begin()
	if len(addPrices) > 0 {
		err := InsertPrices(tx, addPrices)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if len(updatePrices) > 0 {
		err := UpdatePrices(tx, updatePrices, &batchPrices.Price)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if len(deletePrices) > 0 {
		err := DeletePrices(tx, deletePrices)
		if err != nil {
			tx.Rollback()
			return err
		}

	}
	tx.Commit()

	return p.Init()
}

func GetPricesList(pricingType string) []*Price {
	var prices []*Price

	switch pricingType {
	case "default":
		prices = GetDefaultPrice()
	case "db":
		prices = PricingInstance.GetAllPricesList()
	case "old":
		prices = GetOldPricesList()
	default:
		return nil
	}

	sort.Slice(prices, func(i, j int) bool {
		if prices[i].ChannelType == prices[j].ChannelType {
			return prices[i].Model < prices[j].Model
		}
		return prices[i].ChannelType < prices[j].ChannelType
	})

	return prices
}

func GetOldPricesList() []*Price {
	oldDataJson, err := GetOption("ModelRatio")
	if err != nil || oldDataJson.Value == "" {
		return nil
	}

	oldData := make(map[string][]float64)
	err = json.Unmarshal([]byte(oldDataJson.Value), &oldData)

	if err != nil {
		return nil
	}

	var prices []*Price
	for modelName, oldPrice := range oldData {
		price := PricingInstance.GetPrice(modelName)
		prices = append(prices, &Price{
			Model:       modelName,
			Type:        TokensPriceType,
			ChannelType: price.ChannelType,
			Input:       oldPrice[0],
			Output:      oldPrice[1],
		})
	}

	return prices
}

func BatchUpdatePriceOwnedByType(models []string, ownedByType int) error {
	if len(models) == 0 {
		return nil
	}

	cleanSet := make(map[string]struct{}, len(models))
	cleanList := make([]string, 0, len(models))
	for _, modelName := range models {
		trimmed := strings.TrimSpace(modelName)
		if trimmed == "" {
			continue
		}
		if _, exists := cleanSet[trimmed]; exists {
			continue
		}
		cleanSet[trimmed] = struct{}{}
		cleanList = append(cleanList, trimmed)
	}

	if len(cleanList) == 0 {
		return nil
	}

	tx := DB.Begin()
	if err := tx.Model(&Price{}).Where("model IN (?)", cleanList).Update("owned_by_type", ownedByType).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return PricingInstance.Init()
}

func BatchUpdatePriceChannelType(models []string, channelType int) error {
    if len(models) == 0 {
        return nil
    }

    cleanSet := make(map[string]struct{}, len(models))
    cleanList := make([]string, 0, len(models))
    for _, modelName := range models {
        trimmed := strings.TrimSpace(modelName)
        if trimmed == "" {
            continue
        }
        if _, exists := cleanSet[trimmed]; exists {
            continue
        }
        cleanSet[trimmed] = struct{}{}
        cleanList = append(cleanList, trimmed)
    }

    if len(cleanList) == 0 {
        return nil
    }

    tx := DB.Begin()
    if err := tx.Model(&Price{}).Where("model IN (?)", cleanList).Update("channel_type", channelType).Error; err != nil {
        tx.Rollback()
        return err
    }

    if err := tx.Commit().Error; err != nil {
        return err
    }

    return PricingInstance.Init()
}

// func ConvertBatchPrices(prices []*Price) []*BatchPrices {
// 	batchPricesMap := make(map[string]*BatchPrices)
// 	for _, price := range prices {
// 		key := fmt.Sprintf("%s-%d-%g-%g", price.Type, price.ChannelType, price.Input, price.Output)
// 		batchPrice, exists := batchPricesMap[key]
// 		if exists {
// 			batchPrice.Models = append(batchPrice.Models, price.Model)
// 		} else {
// 			batchPricesMap[key] = &BatchPrices{
// 				Models: []string{price.Model},
// 				Price:  *price,
// 			}
// 		}
// 	}

// 	var batchPrices []*BatchPrices
// 	for _, batchPrice := range batchPricesMap {
// 		batchPrices = append(batchPrices, batchPrice)
// 	}

// 	return batchPrices
// }

func GetDefaultModelMapping() map[string]string {
	return map[string]string{}
}
