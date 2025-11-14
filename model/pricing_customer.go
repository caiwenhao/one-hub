package model

import (
	"errors"
	"strings"
	"time"
)

const (
	// SecondsPriceType 表示按秒计费的类型，仅用于少数视频模型
	SecondsPriceType = "seconds"
)

// ResolveCustomerPrice 解析指定用户在某模型/分组上的最终价格
// groupCode 为空时自动使用默认分组。
// 返回 *Price 以便与现有计费逻辑无缝衔接。
func ResolveCustomerPrice(userID int, modelName, groupCode string) (*Price, string, error) {
	modelName = strings.TrimSpace(modelName)
	if modelName == "" {
		return nil, "", errors.New("empty model name")
	}

	// 1. 查找分组信息
	var group ModelGroup
	var err error
	resolvedGroup := groupCode

	if groupCode == "" {
		// 使用默认分组
		err = DB.Where("model = ? AND is_default = ?", modelName, true).First(&group).Error
		if err != nil {
			// 找不到默认分组时，退回旧逻辑，以避免影响现有行为
			price := PricingInstance.GetPrice(modelName)
			if price == nil {
				return nil, "", errors.New("no price found for model")
			}
			price.Normalize()
			return price, "", nil
		}
		resolvedGroup = group.GroupCode
	} else {
		err = DB.Where("model = ? AND group_code = ?", modelName, groupCode).First(&group).Error
		if err != nil {
			return nil, "", errors.New("model group not found")
		}
	}

	// 2. 授权校验：默认分组 MVP 阶段默认允许；非默认分组需要显式授权
	if !group.IsDefault {
		var perm UserModelGroupPermission
		err = DB.Where("user_id = ? AND model = ? AND group_code = ? AND enabled = ?", userID, modelName, resolvedGroup, true).First(&perm).Error
		if err != nil {
			return nil, resolvedGroup, errors.New("model group not permitted for user")
		}
	}

	// 3. 优先查找客户价
	var userPrice UserModelGroupPrice
	err = DB.Where("user_id = ? AND model = ? AND group_code = ? AND enabled = ?", userID, modelName, resolvedGroup, true).First(&userPrice).Error
	if err == nil {
		// 将客户价映射为 Price 结构
		price := &Price{
			Model:       modelName,
			Type:        userPrice.Type,
			Input:       userPrice.InputRate,
			Output:      userPrice.OutputRate,
			ExtraRatios: userPrice.ExtraRatios,
		}
		// 保持 OwnedByType/ChannelType 等通过 Normalize 再次解析
		price.Normalize()
		return price, resolvedGroup, nil
	}

	// 4. 没有客户价时的兜底逻辑
	if group.IsDefault {
		// 默认分组回退到官方价
		price := PricingInstance.GetPrice(modelName)
		if price == nil {
			return nil, resolvedGroup, errors.New("no price found for model")
		}
		price.Normalize()
		return price, resolvedGroup, nil
	}

	// 非默认分组且没有客户价：视为未配置价格
	return nil, resolvedGroup, errors.New("no customer price configured for this model group")
}

// EnsureDefaultModelGroupForPrice 在系统启动时，为已有 Price 记录补齐默认分组记录
// 该函数在 migrationAfter 或 SetupDB 过程中调用，以保证新特性有合理的初始数据。
func EnsureDefaultModelGroupForPrice() {
	var prices []*Price
	if err := DB.Find(&prices).Error; err != nil {
		return
	}

	now := time.Now().Unix()
	for _, p := range prices {
		if strings.TrimSpace(p.Model) == "" {
			continue
		}
		var count int64
		DB.Model(&ModelGroup{}).Where("model = ? AND is_default = ?", p.Model, true).Count(&count)
		if count > 0 {
			continue
		}
		// 创建默认分组，计费类型与 Price.Type 保持一致
		group := &ModelGroup{
			Model:       p.Model,
			GroupCode:   "default",
			DisplayName: "默认分组",
			Description: "由系统根据现有价格自动创建的默认分组",
			IsDefault:   true,
			BillingType: p.Type,
			CreatedAt:   now,
			UpdatedAt:   now,
		}
		_ = DB.Create(group).Error
	}
}
