package router

import (
    "github.com/gin-gonic/gin"
    vidu "one-api/providers/vidu"
)

// Vidu endpoints

// swagger:response swaggerViduListResponse
type swaggerViduListResponse struct {
    // 列表数据
    // in: body
    Data  []vidu.ViduQueryResponse `json:"data"`
    Total int                      `json:"total"`
    Page  int                      `json:"page"`
    Size  int                      `json:"size"`
}

// @Summary      Vidu 获取单个任务状态
// @Tags         Vidu
// @Produce      json
// @Security     BearerAuth
// @Param        task_id path string true "Task ID"
// @Success      200 {object} vidu.ViduQueryResponse
// @Router       /vidu/ent/v2/task/{task_id} [get]
func docViduGetTask(*gin.Context) {}

// @Summary      Vidu 分页查询任务列表
// @Tags         Vidu
// @Produce      json
// @Security     BearerAuth
// @Param        page  query  int   false "页码"
// @Param        size  query  int   false "每页数量"
// @Success      200 {object} swaggerViduListResponse
// @Router       /vidu/ent/v2/tasks [get]
func docViduListTasks(*gin.Context) {}

// @Summary      Vidu 查询任务生成结果（creations）
// @Tags         Vidu
// @Produce      json
// @Security     BearerAuth
// @Param        task_id path string true "Task ID"
// @Success      200 {object} vidu.ViduQueryResponse
// @Router       /vidu/ent/v2/tasks/{task_id}/creations [get]
func docViduListTaskCreations(*gin.Context) {}

// @Summary      Vidu 取消任务
// @Tags         Vidu
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        task_id path string true "Task ID"
// @Success      200 {object} map[string]any
// @Router       /vidu/ent/v2/tasks/{task_id}/cancel [post]
func docViduCancelTask(*gin.Context) {}

// @Summary      Vidu 提交任务
// @Tags         Vidu
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        action path string true "Action" Enums(img2video,reference2video,start-end2video,text2video,reference2image)
// @Param        request body vidu.ViduTaskRequest true "任务请求体（reference2image 请使用 ViduReference2ImageRequest）"
// @Success      200 {object} vidu.ViduResponse
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
