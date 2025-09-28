package volcark

import (
	"encoding/json"
	"one-api/model"
	volcarkProvider "one-api/providers/volcark"
	"one-api/types"
	"strings"

	"github.com/gin-gonic/gin"
)

func StringError(c *gin.Context, httpCode int, code, message string) {
	err := &types.TaskResponse[any]{
		Code:    code,
		Message: message,
	}

	c.JSON(httpCode, err)
}

func TaskModel2Dto(task *model.Task) *volcarkProvider.VolcArkVideoTask {
	if task == nil {
		return nil
	}

	var dto volcarkProvider.VolcArkVideoTask
	if len(task.Data) > 0 {
		if err := json.Unmarshal(task.Data, &dto); err == nil && dto.ID != "" {
			return &dto
		}
	}

	dto.ID = task.TaskID
	dto.Status = mapTaskStatusToVolcStatus(task.Status)
	return &dto
}

func mapTaskStatusToVolcStatus(status model.TaskStatus) string {
	switch status {
	case model.TaskStatusSubmitted, model.TaskStatusNotStart:
		return "queued"
	case model.TaskStatusQueued:
		return "queued"
	case model.TaskStatusInProgress:
		return "running"
	case model.TaskStatusSuccess:
		return "succeeded"
	case model.TaskStatusFailure:
		return "failed"
	default:
		return strings.ToLower(string(status))
	}
}
