package model

import (
	"errors"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const (
	TaskPlatformSuno  = "suno"
	TaskPlatformKling = "kling"
	TaskPlatformVidu  = "vidu"
)

type TaskStatus string

const (
	TaskStatusNotStart   TaskStatus = "NOT_START"
	TaskStatusSubmitted             = "SUBMITTED"
	TaskStatusQueued                = "QUEUED"
	TaskStatusInProgress            = "IN_PROGRESS"
	TaskStatusFailure               = "FAILURE"
	TaskStatusSuccess               = "SUCCESS"
	TaskStatusUnknown               = "UNKNOWN"
)

type Task struct {
	ID             int64          `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt      int64          `json:"created_at" gorm:"index"`
	UpdatedAt      int64          `json:"updated_at"`
	TaskID         string         `json:"task_id" gorm:"type:varchar(50);index"`           // 第三方id，不一定有/ song id\ Task id
	ExternalTaskID string         `json:"external_task_id" gorm:"type:varchar(100);index"` // 用户自定义任务ID
	Platform       string         `json:"platform" gorm:"type:varchar(30);index"`          // 平台
	UserId         int            `json:"user_id" gorm:"index"`
	ChannelId      int            `json:"channel_id" gorm:"index"`
	Quota          int            `json:"quota"`
	Action         string         `json:"action" gorm:"type:varchar(40);index"` // 任务类型, song, lyrics, description-mode
	Status         TaskStatus     `json:"status" gorm:"type:varchar(20);index"` // 任务状态
	FailReason     string         `json:"fail_reason"`
	SubmitTime     int64          `json:"submit_time" gorm:"index"`
	StartTime      int64          `json:"start_time" gorm:"index"`
	FinishTime     int64          `json:"finish_time" gorm:"index"`
	Progress       int            `json:"progress"`
	Properties     datatypes.JSON `json:"properties" gorm:"type:json"`
	Data           datatypes.JSON `json:"data" gorm:"type:json"`
	NotifyHook     string         `json:"notify_hook"`
	TokenID        int            `json:"token_id" gorm:"default:0"`
}

func GetTaskByTaskIds(platform string, userId int, taskIds []string) (task []*Task, err error) {
	// 最多返回100个任务
	err = DB.Omit("channel_id", "quota", "user_id").Where("platform = ? and user_id = ? and task_id in (?)", platform, userId, taskIds).Limit(100).
		Find(&task).Error

	return
}

func GetTaskActionByTaskIds(platform string, taskIds []string) (task []*Task, err error) {
	err = DB.Select("id,action,task_id").Where("platform = ? and task_id in (?)", platform, taskIds).Find(&task).Error

	return
}

func GetByUserAndTaskId(userId int, taskId string) *Task {
	var task Task
	err := DB.Where("user_id = ? and task_id = ?", userId, taskId).First(&task).Error
	if err != nil {
		return nil
	}
	return &task
}

func GetTaskByTaskId(platform string, userId int, taskId string) (task *Task, err error) {
	task = &Task{}
	err = DB.Where("platform = ? and user_id = ? and task_id = ?", platform, userId, taskId).First(task).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return
}

func GetTaskByTaskIdAndActions(platform string, userId int, taskId string, actions []string) (*Task, error) {
	task := &Task{}
	tx := DB.Where("platform = ? and user_id = ? and task_id = ?", platform, userId, taskId)
	if len(actions) > 0 {
		tx = tx.Where("action IN ?", actions)
	}

	err := tx.First(task).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return task, err
}

func (Task *Task) Insert() error {
	return DB.Create(Task).Error
}

func (Task *Task) Update() error {
	return DB.Save(Task).Error
}

func TaskBulkUpdate(TaskIds []string, params map[string]any) error {
	if len(TaskIds) == 0 {
		return nil
	}
	return DB.Model(&Task{}).
		Where("task_id in (?)", TaskIds).
		Updates(params).Error
}

func TaskBulkUpdateByTaskIds(taskIDs []int64, params map[string]any) error {
	if len(taskIDs) == 0 {
		return nil
	}
	return DB.Model(&Task{}).
		Where("id in (?)", taskIDs).
		Updates(params).Error
}

func TaskBulkUpdateByID(ids []int64, params map[string]any) error {
	if len(ids) == 0 {
		return nil
	}
	return DB.Model(&Task{}).
		Where("id in (?)", ids).
		Updates(params).Error
}

func GetAllUnFinishSyncTasks(limit int) []*Task {
	var tasks []*Task
	// get all tasks progress is not 100%
	err := DB.Where("progress != ?", "100").Limit(limit).Order("id").Find(&tasks).Error
	if err != nil {
		return nil
	}
	return tasks
}

type TaskQueryParams struct {
	PaginationParams
	Platform       string `form:"platform"`
	ChannelID      string `form:"channel_id"`
	TaskID         string `form:"task_id"`
	UserID         string `form:"user_id"`
	Action         string `form:"action"`
	Status         string `form:"status"`
	StartTimestamp int64  `form:"start_timestamp"`
	EndTimestamp   int64  `form:"end_timestamp"`
	UserIDs        []int  `form:"user_ids"`
	TokenID        int    `form:"token_id"`
}

var allowedTaskOrderFields = map[string]bool{
	"id":          true,
	"created_at":  true,
	"submit_time": true,
	"finish_time": true,
	"channel_id":  true,
	"user_id":     true,
	"platform":    true,
}

func GetAllTasks(params *TaskQueryParams) (*DataResult[Task], error) {
	tx := DB
	var tasks []*Task

	if params.ChannelID != "" {
		tx = tx.Where("channel_id = ?", params.ChannelID)
	}
	if params.Platform != "" {
		tx = tx.Where("platform = ?", params.Platform)
	}
	if params.UserID != "" {
		tx = tx.Where("user_id = ?", params.UserID)
	}
	if len(params.UserIDs) != 0 {
		tx = tx.Where("user_id in (?)", params.UserIDs)
	}
	if params.TaskID != "" {
		tx = tx.Where("task_id = ?", params.TaskID)
	}
	if params.Action != "" {
		tx = tx.Where("action = ?", params.Action)
	}
	if params.Status != "" {
		tx = tx.Where("status = ?", params.Status)
	}
	if params.StartTimestamp != 0 {
		tx = tx.Where("submit_time >= ?", params.StartTimestamp)
	}
	if params.EndTimestamp != 0 {
		tx = tx.Where("submit_time <= ?", params.EndTimestamp)
	}

	return PaginateAndOrder(tx, &params.PaginationParams, &tasks, allowedTaskOrderFields)
}

func GetAllUserTasks(userId int, params *TaskQueryParams) (*DataResult[Task], error) {
	tx := DB.Omit("channel_id").Where("user_id = ?", userId)
	var tasks []*Task

	if params.TokenID > 0 {
		tx = tx.Where("token_id = ?", params.TokenID)
	}

	if params.Platform != "" {
		tx = tx.Where("platform = ?", params.Platform)
	}

	if params.TaskID != "" {
		tx = tx.Where("task_id = ?", params.TaskID)
	}
	if params.Action != "" {
		tx = tx.Where("action = ?", params.Action)
	}
	if params.Status != "" {
		tx = tx.Where("status = ?", params.Status)
	}
	if params.StartTimestamp != 0 {
		tx = tx.Where("submit_time >= ?", params.StartTimestamp)
	}
	if params.EndTimestamp != 0 {
		tx = tx.Where("submit_time <= ?", params.EndTimestamp)
	}

	return PaginateAndOrder(tx, &params.PaginationParams, &tasks, allowedTaskOrderFields)
}

// GetTasksByExternalTaskID 通过external_task_id查询任务
func GetTasksByExternalTaskID(platform string, userId int, externalTaskID string) ([]*Task, error) {
	var tasks []*Task
	err := DB.Where("platform = ? and user_id = ? and external_task_id = ?", platform, userId, externalTaskID).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTasksByExternalTaskIDAndActions(platform string, userId int, externalTaskID string, actions []string) ([]*Task, error) {
	var tasks []*Task
	tx := DB.Where("platform = ? and user_id = ? and external_task_id = ?", platform, userId, externalTaskID)
	if len(actions) > 0 {
		tx = tx.Where("action IN ?", actions)
	}
	err := tx.Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// GetTasksList 查询任务列表（带分页）
func GetTasksList(platform string, userId int, pageNum, pageSize int) ([]*Task, error) {
	var tasks []*Task
	offset := (pageNum - 1) * pageSize
	err := DB.Where("platform = ? and user_id = ?", platform, userId).
		Order("created_at desc").
		Limit(pageSize).
		Offset(offset).
		Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTasksListByActions(platform string, userId int, actions []string, pageNum, pageSize int) ([]*Task, error) {
	var tasks []*Task
	offset := (pageNum - 1) * pageSize
	tx := DB.Where("platform = ? and user_id = ?", platform, userId)
	if len(actions) > 0 {
		tx = tx.Where("action IN ?", actions)
	}
	err := tx.Order("created_at desc").Limit(pageSize).Offset(offset).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// CreateTask 创建任务
func CreateTask(task *Task) error {
	return DB.Create(task).Error
}
