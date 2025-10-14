package minimax

import (
	"net/http"

	"one-api/model"
	"one-api/providers"
	miniProvider "one-api/providers/minimax"

	"github.com/gin-gonic/gin"
)

func RelayTaskFetch(c *gin.Context) {
	taskID := c.Param("task_id")
	if taskID == "" {
		taskID = c.Query("task_id")
	}

	if taskID == "" {
		StringError(c, http.StatusBadRequest, "invalid_request", "task_id is required")
		return
	}

	userID := c.GetInt("id")
	task := model.GetByUserAndTaskId(userID, taskID)
	if task == nil {
		StringError(c, http.StatusNotFound, "task_not_found", "task not found")
		return
	}

	if task.Platform != model.TaskPlatformMiniMax {
		StringError(c, http.StatusBadRequest, "invalid_platform", "task platform mismatch")
		return
	}

	channel, err := model.GetChannelById(task.ChannelId)
	if err != nil {
		StringError(c, http.StatusServiceUnavailable, "channel_not_found", err.Error())
		return
	}

	provider := providers.GetProvider(channel, c)
	mini, ok := provider.(*miniProvider.MiniMaxProvider)
	if !ok || mini.GetVideoClient() == nil {
		StringError(c, http.StatusServiceUnavailable, "provider_not_found", "provider not found")
		return
	}

	resp, queryErr := mini.GetVideoClient().QueryVideoTask(task.TaskID, extractModelFromTask(task))
	if queryErr != nil {
		StringError(c, http.StatusInternalServerError, "query_failed", queryErr.Message)
		return
	}

	if err := applyQueryResult(task, resp); err != nil {
		StringError(c, http.StatusInternalServerError, "update_failed", err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

func RelayTaskList(c *gin.Context) {
	userID := c.GetInt("id")

	var params model.TaskQueryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		StringError(c, http.StatusBadRequest, "invalid_request", err.Error())
		return
	}
	params.Platform = model.TaskPlatformMiniMax

	result, err := model.GetAllUserTasks(userID, &params)
	if err != nil {
		StringError(c, http.StatusInternalServerError, "query_failed", err.Error())
		return
	}

	var items []*miniProvider.MiniMaxVideoQueryResponse
	for _, task := range *result.Data {
		items = append(items, taskModelToDto(task))
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  items,
		"total": result.Size,
		"page":  result.Page,
		"size":  result.Size,
	})
}
