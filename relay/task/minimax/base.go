package minimax

import (
	"encoding/json"
	"one-api/model"
	miniProvider "one-api/providers/minimax"
	"one-api/types"

	"github.com/gin-gonic/gin"
)

func StringError(c *gin.Context, httpCode int, code, message string) {
	err := &types.TaskResponse[any]{
		Code:    code,
		Message: message,
	}

	c.JSON(httpCode, err)
}

func taskModelToDto(task *model.Task) *miniProvider.MiniMaxVideoQueryResponse {
	if task == nil {
		return nil
	}

	var resp miniProvider.MiniMaxVideoQueryResponse
	if len(task.Data) > 0 {
		if err := json.Unmarshal(task.Data, &resp); err == nil {
			if resp.TaskID == "" {
				resp.TaskID = task.TaskID
			}
			return &resp
		}
	}

	return &miniProvider.MiniMaxVideoQueryResponse{
		TaskID: task.TaskID,
		Status: string(task.Status),
	}
}
