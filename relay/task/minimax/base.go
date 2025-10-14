package minimax

import (
	"bytes"
	"encoding/json"
	"net/http"
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

	writeJSONNoEscape(c, httpCode, err)
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

func writeJSONNoEscape(c *gin.Context, status int, payload interface{}) {
	c.Status(status)
	c.Header("Content-Type", "application/json; charset=utf-8")
	encoder := json.NewEncoder(c.Writer)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(payload); err != nil {
		c.Status(http.StatusInternalServerError)
		_ = encoder.Encode(&types.TaskResponse[any]{
			Code:    "encode_failed",
			Message: err.Error(),
		})
	}
}

func marshalNoEscape(v interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(v); err != nil {
		return nil, err
	}
	return bytes.TrimSpace(buf.Bytes()), nil
}
