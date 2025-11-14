package model

// UserModelGroupPermission 控制某个用户对某模型某分组的可用性
// 默认分组在 MVP 阶段视为所有用户默认可用，因此主要用于非默认分组
type UserModelGroupPermission struct {
	ID int `json:"id" gorm:"primaryKey"`

	UserID    int    `json:"user_id" gorm:"index:idx_perm_user_model_group,priority:1"`
	Model     string `json:"model" gorm:"type:varchar(100);index:idx_perm_user_model_group,priority:2"`
	GroupCode string `json:"group_code" gorm:"type:varchar(64);index:idx_perm_user_model_group,priority:3"`

	Enabled   bool  `json:"enabled" gorm:"default:true;index"`
	CreatedAt int64 `json:"created_at" gorm:"bigint"`
	UpdatedAt int64 `json:"updated_at" gorm:"bigint"`
}
