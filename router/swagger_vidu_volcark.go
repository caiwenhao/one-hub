package router

import "github.com/gin-gonic/gin"

// Vidu endpoints

// @Summary      Vidu get task
// @Tags         Vidu
// @Produce      json
// @Security     BearerAuth
// @Param        task_id path string true "Task ID"
// @Router       /vidu/ent/v2/task/{task_id} [get]
func docViduGetTask(*gin.Context) {}

// @Summary      Vidu list tasks
// @Tags         Vidu
// @Produce      json
// @Security     BearerAuth
// @Router       /vidu/ent/v2/tasks [get]
func docViduListTasks(*gin.Context) {}

// @Summary      Vidu list creations of task
// @Tags         Vidu
// @Produce      json
// @Security     BearerAuth
// @Param        task_id path string true "Task ID"
// @Router       /vidu/ent/v2/tasks/{task_id}/creations [get]
func docViduListTaskCreations(*gin.Context) {}

// @Summary      Vidu cancel task
// @Tags         Vidu
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        task_id path string true "Task ID"
// @Router       /vidu/ent/v2/tasks/{task_id}/cancel [post]
func docViduCancelTask(*gin.Context) {}

// @Summary      Vidu submit task
// @Tags         Vidu
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        action path string true "Action"
// @Param        request body map[string]interface{} true "Task request"
// @Router       /vidu/ent/v2/{action} [post]
func docViduSubmitTask(*gin.Context) {}

// VolcArk endpoints

// @Summary      VolcArk get generation task
// @Tags         VolcArk
// @Produce      json
// @Security     BearerAuth
// @Param        task_id path string true "Task ID"
// @Router       /volcark/api/v3/contents/generations/tasks/{task_id} [get]
func docVolcGetTask(*gin.Context) {}

// @Summary      VolcArk list generation tasks
// @Tags         VolcArk
// @Produce      json
// @Security     BearerAuth
// @Router       /volcark/api/v3/contents/generations/tasks [get]
func docVolcListTasks(*gin.Context) {}

// @Summary      VolcArk cancel task
// @Tags         VolcArk
// @Produce      json
// @Security     BearerAuth
// @Param        task_id path string true "Task ID"
// @Router       /volcark/api/v3/contents/generations/tasks/{task_id} [delete]
func docVolcCancelTask(*gin.Context) {}

// @Summary      VolcArk create image generation
// @Tags         VolcArk
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Router       /volcark/api/v3/images/generations [post]
func docVolcImageGen(*gin.Context) {}

// @Summary      VolcArk create content generation task
// @Tags         VolcArk
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Router       /volcark/api/v3/contents/generations/tasks [post]
func docVolcCreateTask(*gin.Context) {}

