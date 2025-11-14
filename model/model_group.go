package model

import "gorm.io/gorm"

// ModelGroup 表示某个模型的分组（质量档 / 套餐档）
// 默认分组的官方价仍由 Price 表提供；ModelGroup 主要用于声明分组与计费类型等元信息。
type ModelGroup struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Model       string `json:"model" gorm:"type:varchar(100);index:idx_model_group_code,priority:1"`
	GroupCode   string `json:"group_code" gorm:"type:varchar(64);index:idx_model_group_code,priority:2"`
	DisplayName string `json:"display_name" gorm:"type:varchar(128)"`
	Description string `json:"description" gorm:"type:varchar(255)"`
	IsDefault   bool   `json:"is_default" gorm:"index"`
	// BillingType 计费类型：tokens / times / seconds
	BillingType string         `json:"billing_type" gorm:"type:varchar(32)"`
	CreatedAt   int64          `json:"created_at" gorm:"bigint"`
	UpdatedAt   int64          `json:"updated_at" gorm:"bigint"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
