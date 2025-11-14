package controller

import (
	"net/http"
	"one-api/common"
	"one-api/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// ModelGroupDTO 用于管理端展示模型分组与客户价配置
type ModelGroupDTO struct {
	Model       string `json:"model"`
	GroupCode   string `json:"group_code"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
	IsDefault   bool   `json:"is_default"`
	BillingType string `json:"billing_type"`

	// 客户价相关（可能为空）
	HasCustomerPrice bool    `json:"has_customer_price"`
	Type             string  `json:"type"`
	InputRate        float64 `json:"input_rate"`
	OutputRate       float64 `json:"output_rate"`

	// 非默认分组授权信息
	Permitted bool `json:"permitted"`
}

// GetModelGroups 获取某模型的所有分组定义（管理员）
func GetModelGroups(c *gin.Context) {
	modelName := c.Query("model")
	if modelName == "" {
		common.APIRespondWithError(c, http.StatusOK, gin.Error{Err: http.ErrNoLocation})
		return
	}

	var groups []model.ModelGroup
	if err := model.DB.Where("model = ?", modelName).Find(&groups).Error; err != nil {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    groups,
	})
}

// UpsertModelGroupRequest 管理员创建/更新模型分组的请求
type UpsertModelGroupRequest struct {
	Model       string `json:"model" binding:"required"`
	GroupCode   string `json:"group_code" binding:"required"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
	BillingType string `json:"billing_type"` // tokens / times / seconds
	IsDefault   bool   `json:"is_default"`
}

// UpsertModelGroup 管理员创建或更新模型分组
func UpsertModelGroup(c *gin.Context) {
	var req UpsertModelGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}

	now := time.Now().Unix()
	var group model.ModelGroup
	err := model.DB.Where("model = ? AND group_code = ?", req.Model, req.GroupCode).First(&group).Error
	if err != nil {
		// 新建
		group = model.ModelGroup{
			Model:       req.Model,
			GroupCode:   req.GroupCode,
			DisplayName: req.DisplayName,
			Description: req.Description,
			BillingType: req.BillingType,
			IsDefault:   req.IsDefault,
			CreatedAt:   now,
			UpdatedAt:   now,
		}
		if err = model.DB.Create(&group).Error; err != nil {
			common.APIRespondWithError(c, http.StatusOK, err)
			return
		}
	} else {
		// 更新
		group.DisplayName = req.DisplayName
		group.Description = req.Description
		if req.BillingType != "" {
			group.BillingType = req.BillingType
		}
		group.UpdatedAt = now
		if req.IsDefault {
			group.IsDefault = true
		}
		if err = model.DB.Save(&group).Error; err != nil {
			common.APIRespondWithError(c, http.StatusOK, err)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
	})
}

// GetUserModelPricing 管理员查看某用户在某模型下的分组价与授权情况
func GetUserModelPricing(c *gin.Context) {
	idStr := c.Param("id")
	userID, err := strconv.Atoi(idStr)
	if err != nil || userID <= 0 {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}
	modelName := c.Query("model")
	if modelName == "" {
		common.APIRespondWithError(c, http.StatusOK, gin.Error{Err: http.ErrNoLocation})
		return
	}

	var groups []model.ModelGroup
	if err := model.DB.Where("model = ?", modelName).Find(&groups).Error; err != nil {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}

	var prices []model.UserModelGroupPrice
	if err := model.DB.Where("user_id = ? AND model = ?", userID, modelName).Find(&prices).Error; err != nil {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}
	priceMap := make(map[string]model.UserModelGroupPrice)
	for _, p := range prices {
		priceMap[p.GroupCode] = p
	}

	var perms []model.UserModelGroupPermission
	if err := model.DB.Where("user_id = ? AND model = ?", userID, modelName).Find(&perms).Error; err != nil {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}
	permMap := make(map[string]model.UserModelGroupPermission)
	for _, p := range perms {
		permMap[p.GroupCode] = p
	}

	var result []ModelGroupDTO
	for _, g := range groups {
		dto := ModelGroupDTO{
			Model:       g.Model,
			GroupCode:   g.GroupCode,
			DisplayName: g.DisplayName,
			Description: g.Description,
			IsDefault:   g.IsDefault,
			BillingType: g.BillingType,
		}
		if up, ok := priceMap[g.GroupCode]; ok && up.Enabled {
			dto.HasCustomerPrice = true
			dto.Type = up.Type
			dto.InputRate = up.InputRate
			dto.OutputRate = up.OutputRate
		}
		if perm, ok := permMap[g.GroupCode]; ok && perm.Enabled {
			dto.Permitted = true
		}
		// 默认分组在当前阶段视为所有用户可用
		if g.IsDefault {
			dto.Permitted = true
		}
		result = append(result, dto)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    result,
	})
}

// UpdateUserModelPricingRequest 管理员更新某用户在某模型下的分组价与授权
type UpdateUserModelPricingRequest struct {
	Model  string          `json:"model" binding:"required"`
	Groups []ModelGroupDTO `json:"groups" binding:"required"`
}

// UpdateUserModelPricing 批量更新某用户对某模型的分组客户价与授权
func UpdateUserModelPricing(c *gin.Context) {
	idStr := c.Param("id")
	userID, err := strconv.Atoi(idStr)
	if err != nil || userID <= 0 {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}

	var req UpdateUserModelPricingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}

	now := time.Now().Unix()

	for _, item := range req.Groups {
		groupCode := item.GroupCode
		if groupCode == "" {
			continue
		}

		// 1) 处理客户价
		var up model.UserModelGroupPrice
		err := model.DB.Where("user_id = ? AND model = ? AND group_code = ?", userID, req.Model, groupCode).First(&up).Error
		if item.HasCustomerPrice && (item.Type != "" || item.InputRate != 0 || item.OutputRate != 0) {
			// 写入/更新客户价
			if err != nil {
				up = model.UserModelGroupPrice{
					UserID:    userID,
					Model:     req.Model,
					GroupCode: groupCode,
				}
			}
			up.Type = item.Type
			up.InputRate = item.InputRate
			up.OutputRate = item.OutputRate
			up.Enabled = true
			up.UpdatedAt = now
			if up.CreatedAt == 0 {
				up.CreatedAt = now
			}
			if up.ID == 0 {
				if err = model.DB.Create(&up).Error; err != nil {
					common.APIRespondWithError(c, http.StatusOK, err)
					return
				}
			} else {
				if err = model.DB.Save(&up).Error; err != nil {
					common.APIRespondWithError(c, http.StatusOK, err)
					return
				}
			}
		} else if err == nil {
			// 没有客户价配置时，如果存在记录则禁用
			up.Enabled = false
			up.UpdatedAt = now
			_ = model.DB.Save(&up).Error
		}

		// 2) 处理授权（非默认分组才需要）
		if !item.IsDefault {
			var perm model.UserModelGroupPermission
			err = model.DB.Where("user_id = ? AND model = ? AND group_code = ?", userID, req.Model, groupCode).First(&perm).Error
			if item.Permitted {
				if err != nil {
					perm = model.UserModelGroupPermission{
						UserID:    userID,
						Model:     req.Model,
						GroupCode: groupCode,
					}
				}
				perm.Enabled = true
				perm.UpdatedAt = now
				if perm.CreatedAt == 0 {
					perm.CreatedAt = now
				}
				if perm.ID == 0 {
					if err = model.DB.Create(&perm).Error; err != nil {
						common.APIRespondWithError(c, http.StatusOK, err)
						return
					}
				} else {
					if err = model.DB.Save(&perm).Error; err != nil {
						common.APIRespondWithError(c, http.StatusOK, err)
						return
					}
				}
			} else if err == nil {
				perm.Enabled = false
				perm.UpdatedAt = now
				_ = model.DB.Save(&perm).Error
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
	})
}

// GetUserAvailableModels 返回某个用户可配置客户价的模型列表（仅包含已接入渠道且对该用户分组可用的模型）
// 逻辑：基于用户当前 Group，复用现有 AvailableModels（ChannelGroup + 全局价）计算可用模型集合。
func GetUserAvailableModels(c *gin.Context) {
	idStr := c.Param("id")
	userID, err := strconv.Atoi(idStr)
	if err != nil || userID <= 0 {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}

	// 加载用户信息，获取分组
	user, err := model.GetUserById(userID, false)
	if err != nil {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}
	groupName := user.Group
	if groupName == "" {
		groupName = "default"
	}

	// 基于 ChannelsChooser 的路由规则获取「该用户分组可用的模型」集合
	modelNames, err := model.ChannelGroup.GetGroupModels(groupName)
	if err != nil {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    modelNames,
	})
}

// ConfiguredPriceDTO 用于返回用户已配置的客户价概览
type ConfiguredPriceDTO struct {
	Model  string          `json:"model"`
	Groups []ModelGroupDTO `json:"groups"`
}

// GetUserConfiguredPrices 获取某用户已配置的所有客户价（概览）
func GetUserConfiguredPrices(c *gin.Context) {
	idStr := c.Param("id")
	userID, err := strconv.Atoi(idStr)
	if err != nil || userID <= 0 {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}

	// 查询该用户所有已启用的客户价记录
	var prices []model.UserModelGroupPrice
	if err := model.DB.Where("user_id = ? AND enabled = ?", userID, true).Find(&prices).Error; err != nil {
		common.APIRespondWithError(c, http.StatusOK, err)
		return
	}

	// 按模型分组
	modelMap := make(map[string][]model.UserModelGroupPrice)
	for _, p := range prices {
		modelMap[p.Model] = append(modelMap[p.Model], p)
	}

	// 查询所有相关的模型分组信息
	var allModels []string
	for m := range modelMap {
		allModels = append(allModels, m)
	}

	var modelGroups []model.ModelGroup
	if len(allModels) > 0 {
		if err := model.DB.Where("model IN ?", allModels).Find(&modelGroups).Error; err != nil {
			common.APIRespondWithError(c, http.StatusOK, err)
			return
		}
	}

	// 构建分组信息映射
	groupInfoMap := make(map[string]model.ModelGroup)
	for _, mg := range modelGroups {
		key := mg.Model + ":" + mg.GroupCode
		groupInfoMap[key] = mg
	}

	// 查询授权信息
	var perms []model.UserModelGroupPermission
	if len(allModels) > 0 {
		if err := model.DB.Where("user_id = ? AND model IN ? AND enabled = ?", userID, allModels, true).Find(&perms).Error; err != nil {
			common.APIRespondWithError(c, http.StatusOK, err)
			return
		}
	}

	permMap := make(map[string]bool)
	for _, p := range perms {
		key := p.Model + ":" + p.GroupCode
		permMap[key] = true
	}

	// 构建返回结果
	var result []ConfiguredPriceDTO
	for model, priceList := range modelMap {
		dto := ConfiguredPriceDTO{
			Model:  model,
			Groups: []ModelGroupDTO{},
		}

		for _, p := range priceList {
			key := p.Model + ":" + p.GroupCode
			groupInfo, hasGroupInfo := groupInfoMap[key]

			groupDTO := ModelGroupDTO{
				Model:            p.Model,
				GroupCode:        p.GroupCode,
				HasCustomerPrice: true,
				Type:             p.Type,
				InputRate:        p.InputRate,
				OutputRate:       p.OutputRate,
			}

			if hasGroupInfo {
				groupDTO.DisplayName = groupInfo.DisplayName
				groupDTO.Description = groupInfo.Description
				groupDTO.IsDefault = groupInfo.IsDefault
				groupDTO.BillingType = groupInfo.BillingType
			}

			// 默认分组始终视为可用
			if groupDTO.IsDefault {
				groupDTO.Permitted = true
			} else {
				groupDTO.Permitted = permMap[key]
			}

			dto.Groups = append(dto.Groups, groupDTO)
		}

		result = append(result, dto)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    result,
	})
}
