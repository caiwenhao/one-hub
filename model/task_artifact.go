package model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// TaskArtifact 记录生成产物（文件）与渠道、任务的映射，便于 file_id → channel_id 精确路由
type TaskArtifact struct {
	ID        int64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt int64 `json:"created_at" gorm:"index"`
	UpdatedAt int64 `json:"updated_at"`

	Platform     string `json:"platform" gorm:"type:varchar(30);index:idx_artifact_platform_user_file,priority:1;index:idx_artifact_platform_task,priority:1"`
	UserId       int    `json:"user_id" gorm:"index:idx_artifact_platform_user_file,priority:2"`
	ChannelId    int    `json:"channel_id" gorm:"index"`
	TaskID       string `json:"task_id" gorm:"type:varchar(100);index:idx_artifact_platform_task,priority:2"`
	FileID       string `json:"file_id" gorm:"type:varchar(100);index:idx_artifact_platform_user_file,priority:3"`
	ArtifactType string `json:"artifact_type" gorm:"type:varchar(20)"`
	DownloadURL  string `json:"download_url" gorm:"type:text"`
	TTLAt        int64  `json:"ttl_at" gorm:"index"` // 过期时间（秒级unix时间），0 表示未知

	ExtJson datatypes.JSON `json:"ext_json" gorm:"type:json"`
}

// UpsertTaskArtifact 基于 (platform, user_id, file_id) 做幂等写入
func UpsertTaskArtifact(db *gorm.DB, a *TaskArtifact) error {
	if a == nil || a.FileID == "" || a.Platform == "" || a.UserId <= 0 {
		return nil
	}
	now := time.Now().Unix()
	if a.CreatedAt == 0 {
		a.CreatedAt = now
	}
	a.UpdatedAt = now

	return db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "platform"}, {Name: "user_id"}, {Name: "file_id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"channel_id":    a.ChannelId,
			"task_id":       a.TaskID,
			"artifact_type": a.ArtifactType,
			"download_url":  a.DownloadURL,
			"ttl_at":        a.TTLAt,
			"updated_at":    now,
		}),
	}).Create(a).Error
}

func GetArtifactByFileID(userID int, platform, fileID string) (*TaskArtifact, error) {
	if userID <= 0 || platform == "" || fileID == "" {
		return nil, nil
	}
	var a TaskArtifact
	err := DB.Where("platform = ? AND user_id = ? AND file_id = ?", platform, userID, fileID).First(&a).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &a, nil
}
