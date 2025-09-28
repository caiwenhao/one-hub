package volcark

import (
	"net/http"
	"one-api/model"
	"one-api/providers"
	volcarkProvider "one-api/providers/volcark"

	"github.com/gin-gonic/gin"
)

func RelayTaskFetch(c *gin.Context) {
	taskID := c.Param("task_id")
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

	if task.Platform != model.TaskPlatformVolcArk {
		StringError(c, http.StatusBadRequest, "invalid_platform", "task platform mismatch")
		return
	}

	channel, err := model.GetChannelById(task.ChannelId)
	if err != nil {
		StringError(c, http.StatusServiceUnavailable, "channel_not_found", err.Error())
		return
	}

	provider := providers.GetProvider(channel, c)
	volcProvider, ok := provider.(*volcarkProvider.VolcArkProvider)
	if !ok {
		StringError(c, http.StatusServiceUnavailable, "provider_not_found", "provider not found")
		return
	}

	resp, errInfo := volcProvider.GetVideoTask(task.TaskID)
	if errInfo != nil {
		StringError(c, http.StatusInternalServerError, "query_failed", errInfo.Message)
		return
	}

	if resp != nil {
		applyResponseToTask(task, resp)
		if err := task.Update(); err != nil {
			StringError(c, http.StatusInternalServerError, "update_failed", err.Error())
			return
		}
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

	params.Platform = model.TaskPlatformVolcArk

	result, err := model.GetAllUserTasks(userID, &params)
	if err != nil {
		StringError(c, http.StatusInternalServerError, "query_failed", err.Error())
		return
	}

	var items []interface{}
	for _, task := range *result.Data {
		items = append(items, TaskModel2Dto(task))
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  items,
		"total": result.TotalCount,
		"page":  result.Page,
		"size":  result.Size,
	})
}

func RelayTaskCancel(c *gin.Context) {
	taskID := c.Param("task_id")
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

	if task.Platform != model.TaskPlatformVolcArk {
		StringError(c, http.StatusBadRequest, "invalid_platform", "task platform mismatch")
		return
	}

	channel, err := model.GetChannelById(task.ChannelId)
	if err != nil {
		StringError(c, http.StatusServiceUnavailable, "channel_not_found", err.Error())
		return
	}

	provider := providers.GetProvider(channel, c)
	volcProvider, ok := provider.(*volcarkProvider.VolcArkProvider)
	if !ok {
		StringError(c, http.StatusServiceUnavailable, "provider_not_found", "provider not found")
		return
	}

	if err := volcProvider.CancelVideoTask(task.TaskID); err != nil {
		StringError(c, http.StatusInternalServerError, "cancel_failed", err.Message)
		return
	}

	task.Status = model.TaskStatusFailure
	task.Progress = 100
	task.FailReason = "cancelled_by_user"
	if err := task.Update(); err != nil {
		StringError(c, http.StatusInternalServerError, "cancel_failed", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
