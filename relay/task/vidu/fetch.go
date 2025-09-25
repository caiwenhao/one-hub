package vidu

import (
	"net/http"
	"one-api/model"
	"one-api/providers"
	viduProvider "one-api/providers/vidu"

	"github.com/gin-gonic/gin"
)

func RelayTaskFetch(c *gin.Context) {
	taskId := c.Param("task_id")
	if taskId == "" {
		StringError(c, http.StatusBadRequest, "invalid_request", "task_id is required")
		return
	}

	// 获取用户任务
	userId := c.GetInt("id")
	task := model.GetByUserAndTaskId(userId, taskId)
	if task == nil {
		StringError(c, http.StatusNotFound, "task_not_found", "task not found")
		return
	}

	if task.Platform != model.TaskPlatformVidu {
		StringError(c, http.StatusBadRequest, "invalid_platform", "task platform mismatch")
		return
	}

	// 获取provider
	channel, err := model.GetChannelById(task.ChannelId)
	if err != nil {
		StringError(c, http.StatusServiceUnavailable, "channel_not_found", err.Error())
		return
	}

	provider := providers.GetProvider(channel, c)

	viduProvider, ok := provider.(*viduProvider.ViduProvider)
	if !ok {
		StringError(c, http.StatusServiceUnavailable, "provider_not_found", "provider not found")
		return
	}

	// 查询最新状态 - 使用新的官方接口
	resp, openaiErr := viduProvider.QueryCreations(task.TaskID)
	if openaiErr != nil {
		StringError(c, http.StatusInternalServerError, "query_failed", openaiErr.Message)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// 取消任务接口
func RelayTaskCancel(c *gin.Context) {
	taskId := c.Param("task_id")
	if taskId == "" {
		StringError(c, http.StatusBadRequest, "invalid_request", "task_id is required")
		return
	}

	// 获取用户任务
	userId := c.GetInt("id")
	task := model.GetByUserAndTaskId(userId, taskId)
	if task == nil {
		StringError(c, http.StatusNotFound, "task_not_found", "task not found")
		return
	}

	if task.Platform != model.TaskPlatformVidu {
		StringError(c, http.StatusBadRequest, "invalid_platform", "task platform mismatch")
		return
	}

	// 获取provider
	channel, err := model.GetChannelById(task.ChannelId)
	if err != nil {
		StringError(c, http.StatusServiceUnavailable, "channel_not_found", err.Error())
		return
	}

	provider := providers.GetProvider(channel, c)

	viduProvider, ok := provider.(*viduProvider.ViduProvider)
	if !ok {
		StringError(c, http.StatusServiceUnavailable, "provider_not_found", "provider not found")
		return
	}

	// 取消任务
	if _, openaiErr := viduProvider.CancelTask(task.TaskID); openaiErr != nil {
		StringError(c, http.StatusInternalServerError, "cancel_failed", openaiErr.Message)
		return
	}

	// 更新本地任务状态
	task.Status = model.TaskStatusFailure
	task.FailReason = "cancelled_by_user"
	if err := task.Update(); err != nil {
		StringError(c, http.StatusInternalServerError, "cancel_failed", err.Error())
		return
	}

	// 取消成功返回空响应
	c.JSON(http.StatusOK, gin.H{})
}

func RelayTaskFetchs(c *gin.Context) {
	userId := c.GetInt("id")

	// 获取查询参数
	var params model.TaskQueryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		StringError(c, http.StatusBadRequest, "invalid_request", err.Error())
		return
	}

	params.Platform = model.TaskPlatformVidu

	// 查询用户任务列表
	result, err := model.GetAllUserTasks(userId, &params)
	if err != nil {
		StringError(c, http.StatusInternalServerError, "query_failed", err.Error())
		return
	}

	// 转换为响应格式
	var responses []interface{}
	for _, task := range *result.Data {
		dto := TaskModel2Dto(task)
		responses = append(responses, dto)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data":  responses,
		"total": result.Size,
		"page":  result.Page,
		"size":  result.Size,
	})
}
