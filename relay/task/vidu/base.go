package vidu

import (
	"encoding/json"
	"one-api/model"
	"one-api/types"
	viduProvider "one-api/providers/vidu"

	"github.com/gin-gonic/gin"
)

func StringError(c *gin.Context, httpCode int, code, message string) {
	err := &types.TaskResponse[any]{
		Code:    code,
		Message: message,
	}

	c.JSON(httpCode, err)
}

func TaskModel2Dto(task *model.Task) *viduProvider.ViduResponse[*viduProvider.ViduTaskData] {
	data := &viduProvider.ViduResponse[*viduProvider.ViduTaskData]{}
	json.Unmarshal(task.Data, data)

	return data
}