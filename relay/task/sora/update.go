package sora

import (
    "context"
    "encoding/json"
    "math"
    "strings"
    "time"

    "one-api/common/logger"
    "one-api/model"
    pbase "one-api/providers/base"
    "one-api/providers"
    openaip "one-api/providers/openai"
    "one-api/relay/task/base"
    "one-api/types"

    "gorm.io/datatypes"
    "github.com/gin-gonic/gin"
)

// SoraTask 实现 TaskInterface，用于批量轮询 OpenAI 视频任务（Sora）
type SoraTask struct {
    base.TaskBase
}

// 以下方法为接口占位，调度仅使用 UpdateTaskStatus
func (t *SoraTask) Init() *base.TaskError                   { return nil }
func (t *SoraTask) Relay() *base.TaskError                  { return nil }
func (t *SoraTask) HandleError(err *base.TaskError)         {}
func (t *SoraTask) ShouldRetry(_ *gin.Context, _ *base.TaskError) bool { return false }
func (t *SoraTask) GetModelName() string                    { return t.ModelName }
func (t *SoraTask) GetTask() *model.Task                    { return t.Task }
func (t *SoraTask) SetProvider() *base.TaskError            { return nil }
func (t *SoraTask) GetProvider() pbase.ProviderInterface    { return t.BaseProvider }
func (t *SoraTask) GinResponse()                            {}

func (t *SoraTask) UpdateTaskStatus(ctx context.Context, taskChannelM map[int][]string, taskM map[string]*model.Task) error {
    for channelID, ids := range taskChannelM {
        channel := model.ChannelGroup.GetChannel(channelID)
        if channel == nil {
            for _, id := range ids {
                if task := taskM[id]; task != nil {
                    task.Status = model.TaskStatusFailure
                    task.FailReason = "channel_not_found"
                    _ = task.Update()
                }
            }
            continue
        }

        provider := providers.GetProvider(channel, nil)
        openaiProv, ok := provider.(*openaip.OpenAIProvider)
        if !ok {
            logger.LogError(ctx, "Sora task adaptor: provider not OpenAI")
            continue
        }

        for _, id := range ids {
            job, errWithCode := openaiProv.RetrieveVideo(id)
            if errWithCode != nil || job == nil {
                // 保守处理：暂不将任务标为失败，避免短暂错误导致误判
                if task := taskM[id]; task != nil {
                    // 可选：记录最近错误信息到 FailReason（不改变状态）
                    if errWithCode != nil && task.FailReason == "" {
                        task.FailReason = errWithCode.Message
                        _ = task.Update()
                    }
                }
                continue
            }
            updateSoraTask(taskM[id], job)
        }
    }
    return nil
}

// 复用 relay/videos.go 的语义：将 job 状态映射更新到本地任务
func updateSoraTask(task *model.Task, job *types.VideoJob) {
    if task == nil || job == nil {
        return
    }

    task.Status = mapVideoStatus(job.Status)
    task.Progress = int(math.Round(job.Progress))
    if job.Seconds > 0 {
        task.VideoDuration = float64(job.Seconds)
    }
    if strings.EqualFold(job.Status, "completed") || strings.EqualFold(job.Status, "failed") || strings.EqualFold(job.Status, "success") || strings.EqualFold(job.Status, "error") {
        task.FinishTime = time.Now().Unix()
    }
    if job.Result != nil {
        if dataJSON, err := json.Marshal(job.Result); err == nil {
            task.Data = datatypes.JSON(dataJSON)
        }
    }
    if err := task.Update(); err != nil {
        logger.LogError(context.Background(), "update_sora_task_failed:"+err.Error())
    }
}

func mapVideoStatus(status string) model.TaskStatus {
    switch strings.ToLower(strings.TrimSpace(status)) {
    case "queued":
        return model.TaskStatusQueued
    case "in_progress", "processing":
        return model.TaskStatusInProgress
    case "completed", "success":
        return model.TaskStatusSuccess
    case "failed", "error":
        return model.TaskStatusFailure
    default:
        return model.TaskStatusUnknown
    }
}
