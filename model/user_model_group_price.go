package model

import "gorm.io/datatypes"

// UserModelGroupPrice 表示某个用户（客户）在某模型某分组上的客户价
// 对默认分组：可覆盖 Price 表中的官方价
// 对非默认分组：只有存在客户价 + 授权时才视为可用
type UserModelGroupPrice struct {
	ID int `json:"id" gorm:"primaryKey"`

	UserID    int    `json:"user_id" gorm:"index:idx_user_model_group,priority:1"`
	Model     string `json:"model" gorm:"type:varchar(100);index:idx_user_model_group,priority:2"`
	GroupCode string `json:"group_code" gorm:"type:varchar(64);index:idx_user_model_group,priority:3"`

	// Type 与 Price.Type 语义一致：tokens / times / seconds
	Type string `json:"type" gorm:"type:varchar(32)"`

	// InputRate / OutputRate 采用与 Price.Input / Price.Output 一致的 Rate(K) 语义
	InputRate  float64 `json:"input_rate" gorm:"default:0"`
	OutputRate float64 `json:"output_rate" gorm:"default:0"`

	// ExtraRatios 兼容现有额外用量倍率配置
	ExtraRatios *datatypes.JSONType[map[string]float64] `json:"extra_ratios,omitempty" gorm:"type:json"`

	Enabled   bool  `json:"enabled" gorm:"default:true;index"`
	CreatedAt int64 `json:"created_at" gorm:"bigint"`
	UpdatedAt int64 `json:"updated_at" gorm:"bigint"`
}
