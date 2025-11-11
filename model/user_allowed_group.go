package model

import (
    "errors"
    "time"
)

// UserAllowedGroup 表示某用户被授权可选择的分组白名单
type UserAllowedGroup struct {
    Id        int    `json:"id"`
    UserId    int    `json:"user_id" gorm:"index:idx_user_group,unique"`
    Group     string `json:"group" gorm:"type:varchar(50);index:idx_user_group,unique"`
    CreatedAt int64  `json:"created_at" gorm:"bigint"`
}

// GetAllowedGroupsByUser 返回用户被授权的分组符号列表
func GetAllowedGroupsByUser(userId int) ([]string, error) {
    if userId <= 0 {
        return nil, errors.New("invalid user id")
    }
    var items []UserAllowedGroup
    err := DB.Where("user_id = ?", userId).Find(&items).Error
    if err != nil {
        return nil, err
    }
    groups := make([]string, 0, len(items))
    for _, it := range items {
        groups = append(groups, it.Group)
    }
    return groups, nil
}

// ReplaceAllowedGroups 覆盖设置授权分组（先删后插）
func ReplaceAllowedGroups(userId int, groups []string) error {
    if userId <= 0 {
        return errors.New("invalid user id")
    }
    tx := DB.Begin()
    if err := tx.Where("user_id = ?", userId).Delete(&UserAllowedGroup{}).Error; err != nil {
        tx.Rollback()
        return err
    }
    now := time.Now().Unix()
    for _, g := range groups {
        if g == "" {
            continue
        }
        rec := &UserAllowedGroup{UserId: userId, Group: g, CreatedAt: now}
        if err := tx.Create(rec).Error; err != nil {
            tx.Rollback()
            return err
        }
    }
    return tx.Commit().Error
}

