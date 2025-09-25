package vidu

import (
	"encoding/json"
	"one-api/model"
	viduProvider "one-api/providers/vidu"
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

func TaskModel2Dto(task *model.Task) *viduProvider.ViduQueryResponse {
	if task == nil {
		return nil
	}

	var queryResp viduProvider.ViduQueryResponse
	if len(task.Data) > 0 {
		if err := json.Unmarshal(task.Data, &queryResp); err == nil && (queryResp.ID != "" || len(queryResp.Creations) > 0) {
			if queryResp.ID == "" {
				queryResp.ID = task.TaskID
			}
			return &queryResp
		}

		var legacyResp viduProvider.ViduResponse
		if err := json.Unmarshal(task.Data, &legacyResp); err == nil && legacyResp.TaskID != "" {
			queryResp.ID = legacyResp.TaskID
			queryResp.State = strings.ToLower(legacyResp.State)
			queryResp.Credits = legacyResp.Credits
			queryResp.Payload = legacyResp.Payload
			queryResp.BGM = legacyResp.BGM
			queryResp.OffPeak = legacyResp.OffPeak
			return &queryResp
		}
	}

	queryResp.ID = task.TaskID
	queryResp.State = mapTaskStatusToViduState(task.Status)
	if task.FailReason != "" {
		queryResp.ErrCode = task.FailReason
	}

	return &queryResp
}

func mapTaskStatusToViduState(status model.TaskStatus) string {
	switch status {
	case model.TaskStatusSuccess:
		return viduProvider.ViduStatusSuccess
	case model.TaskStatusFailure:
		return viduProvider.ViduStatusFailed
	case model.TaskStatusInProgress:
		return viduProvider.ViduStatusProcessing
	case model.TaskStatusQueued:
		return viduProvider.ViduStatusQueueing
	case model.TaskStatusSubmitted, model.TaskStatusNotStart:
		return viduProvider.ViduStatusCreated
	default:
		return "unknown"
	}
}
