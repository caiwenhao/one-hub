package kling

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"one-api/common/logger"
	"one-api/model"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

// CreateOfficialImage 官方API - 创建图像生成任务
func CreateOfficialImage(c *gin.Context) {
	var req OfficialImageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, OfficialImageResponse{
			Code:    400,
			Message: fmt.Sprintf("参数错误: %v", err),
		})
		return
	}

	// 参数验证
	if req.Prompt == "" {
		c.JSON(http.StatusBadRequest, OfficialImageResponse{
			Code:    400,
			Message: "prompt是必填参数",
		})
		return
	}

	// 设置默认值
	if req.ModelName == "" {
		req.ModelName = ModelKlingV1
	}
	if req.Resolution == "" {
		req.Resolution = Resolution1K
	}
	if req.N == 0 {
		req.N = 1
	}
	if req.AspectRatio == "" {
		req.AspectRatio = AspectRatio16x9
	}

	// 参数验证
	if err := validateImageRequest(&req); err != nil {
		c.JSON(http.StatusBadRequest, OfficialImageResponse{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	// 转换为内部请求格式
	internalReq := convertImageToInternalRequest(&req)

	// 调用内部处理逻辑
	response := handleImageTask(c, internalReq, "image_generation", req.ExternalTaskID)
	c.JSON(response.StatusCode, response.Response)
}

// CreateOfficialMultiImage2Image 官方API - 创建多图参考生图任务
func CreateOfficialMultiImage2Image(c *gin.Context) {
	var req OfficialMultiImage2ImageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, OfficialImageResponse{
			Code:    400,
			Message: fmt.Sprintf("参数错误: %v", err),
		})
		return
	}

	// 验证主体图片列表
	if len(req.SubjectImageList) == 0 {
		c.JSON(http.StatusBadRequest, OfficialImageResponse{
			Code:    400,
			Message: "subject_image_list是必填参数，至少需要1张图片",
		})
		return
	}

	// 设置默认值
	if req.ModelName == "" {
		req.ModelName = ModelKlingV2
	}
	if req.N == 0 {
		req.N = 1
	}
	if req.AspectRatio == "" {
		req.AspectRatio = AspectRatio16x9
	}

	// 参数验证
	if err := validateMultiImage2ImageRequest(&req); err != nil {
		c.JSON(http.StatusBadRequest, OfficialImageResponse{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	// 转换为内部请求格式
	internalReq := convertMultiImage2ImageToInternalRequest(&req)

	// 调用专门的多图参考生图处理逻辑
	response := handleMultiImage2ImageTask(c, internalReq, "multi-image2image", req.ExternalTaskID)
	c.JSON(response.StatusCode, response.Response)
}

// GetOfficialImageTask 官方API - 查询单个图像任务
func GetOfficialImageTask(c *gin.Context) {
	taskID := c.Param("id")
	userId := c.GetInt("id")
	actions := resolveActionsByPath(c.FullPath())

	// 支持通过external_task_id查询
	var task *model.Task
	var err error

	// 先尝试通过task_id查询
	task, err = model.GetTaskByTaskIdAndActions(model.TaskPlatformKling, userId, taskID, actions)
	if err != nil || task == nil {
		// 再尝试通过external_task_id查询
		tasks, errList := model.GetTasksByExternalTaskIDAndActions(model.TaskPlatformKling, userId, taskID, actions)
		if errList != nil || len(tasks) == 0 {
			c.JSON(http.StatusNotFound, OfficialImageResponse{
				Code:    404,
				Message: "任务不存在",
			})
			return
		}
		task = tasks[0] // 取第一个匹配的任务
	}

	// 转换为官方格式
	officialData := convertToOfficialImageFormat(task)

	c.JSON(http.StatusOK, OfficialImageResponse{
		Code:      0,
		Message:   "success",
		RequestID: generateRequestID(),
		Data:      officialData,
	})
}

// GetOfficialMultiImage2ImageTask 官方API - 查询单个多图参考生图任务
func GetOfficialMultiImage2ImageTask(c *gin.Context) {
	taskID := c.Param("id")
	userId := c.GetInt("id")
	actions := resolveActionsByPath(c.FullPath())

	// 支持通过external_task_id查询
	var task *model.Task
	var err error

	// 先尝试通过task_id查询
	task, err = model.GetTaskByTaskIdAndActions(model.TaskPlatformKling, userId, taskID, actions)
	if err != nil || task == nil {
		// 再尝试通过external_task_id查询
		tasks, errList := model.GetTasksByExternalTaskIDAndActions(model.TaskPlatformKling, userId, taskID, actions)
		if errList != nil || len(tasks) == 0 {
			c.JSON(http.StatusNotFound, OfficialImageResponse{
				Code:    404,
				Message: "任务不存在",
			})
			return
		}
		task = tasks[0] // 取第一个匹配的任务
	}

	// 转换为官方格式
	officialData := convertToOfficialImageFormat(task)

	c.JSON(http.StatusOK, OfficialImageResponse{
		Code:      0,
		Message:   "success",
		RequestID: generateRequestID(),
		Data:      officialData,
	})
}

// ListOfficialImageTasks 官方API - 查询图像任务列表
func ListOfficialImageTasks(c *gin.Context) {
	var query OfficialTaskListQuery

	// 解析查询参数
	if err := c.ShouldBindQuery(&query); err != nil {
		// 设置默认值
		query.PageNum = 1
		query.PageSize = 30
	}

	// 参数验证
	if query.PageNum < 1 {
		query.PageNum = 1
	}
	if query.PageNum > 1000 {
		query.PageNum = 1000
	}
	if query.PageSize < 1 {
		query.PageSize = 30
	}
	if query.PageSize > 500 {
		query.PageSize = 500
	}

	userId := c.GetInt("id")
	actions := resolveActionsByPath(c.FullPath())

	// 查询任务列表
	tasks, err := model.GetTasksListByActions(model.TaskPlatformKling, userId, actions, query.PageNum, query.PageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, OfficialImageResponse{
			Code:    500,
			Message: "查询任务列表失败",
		})
		return
	}

	// 转换为官方格式
	var officialTasks []*OfficialImageTaskData
	for _, task := range tasks {
		officialData := convertToOfficialImageFormat(task)
		officialTasks = append(officialTasks, officialData)
	}

	c.JSON(http.StatusOK, OfficialImageResponse{
		Code:      0,
		Message:   "success",
		RequestID: generateRequestID(),
		DataList:  officialTasks,
	})
}

// ListOfficialMultiImage2ImageTasks 官方API - 查询多图参考生图任务列表
func ListOfficialMultiImage2ImageTasks(c *gin.Context) {
	var query OfficialTaskListQuery

	// 解析查询参数
	if err := c.ShouldBindQuery(&query); err != nil {
		// 设置默认值
		query.PageNum = 1
		query.PageSize = 30
	}

	// 参数验证
	if query.PageNum < 1 {
		query.PageNum = 1
	}
	if query.PageNum > 1000 {
		query.PageNum = 1000
	}
	if query.PageSize < 1 {
		query.PageSize = 30
	}
	if query.PageSize > 500 {
		query.PageSize = 500
	}

	userId := c.GetInt("id")
	actions := resolveActionsByPath(c.FullPath())

	// 查询任务列表
	tasks, err := model.GetTasksListByActions(model.TaskPlatformKling, userId, actions, query.PageNum, query.PageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, OfficialImageResponse{
			Code:    500,
			Message: "查询任务列表失败",
		})
		return
	}

	// 转换为官方格式
	var officialTasks []*OfficialImageTaskData
	for _, task := range tasks {
		officialData := convertToOfficialImageFormat(task)
		officialTasks = append(officialTasks, officialData)
	}

	c.JSON(http.StatusOK, OfficialImageResponse{
		Code:      0,
		Message:   "success",
		RequestID: generateRequestID(),
		DataList:  officialTasks,
	})
}

// ===================== 参数验证函数 =====================

func validateImageRequest(req *OfficialImageRequest) error {
	// 验证模型名称
	validModels := []string{ModelKlingV1, ModelKlingV15, ModelKlingV2, ModelKlingV2New, ModelKlingV21}
	if !contains(validModels, req.ModelName) {
		return fmt.Errorf("不支持的模型: %s", req.ModelName)
	}

	// 验证生成数量
	if req.N < 1 || req.N > 9 {
		return fmt.Errorf("生成数量必须在[1,9]范围内")
	}

	// 验证纵横比
	validRatios := []string{AspectRatio16x9, AspectRatio9x16, AspectRatio1x1, AspectRatio4x3, AspectRatio3x4, AspectRatio3x2, AspectRatio2x3, AspectRatio21x9}
	if !contains(validRatios, req.AspectRatio) {
		return fmt.Errorf("不支持的纵横比: %s", req.AspectRatio)
	}

	// 验证清晰度
	if req.Resolution != Resolution1K && req.Resolution != Resolution2K {
		return fmt.Errorf("不支持的清晰度: %s", req.Resolution)
	}

	// 验证图片参考类型
	if req.ImageReference != "" {
		if req.ImageReference != ImageReferenceSubject && req.ImageReference != ImageReferenceFace {
			return fmt.Errorf("不支持的图片参考类型: %s", req.ImageReference)
		}
		// 使用kling-v1-5且image参数不为空时，image_reference必填
		if req.ModelName == ModelKlingV15 && req.Image != "" && req.ImageReference == "" {
			return fmt.Errorf("使用 kling-v1-5 且 image 参数不为空时，image_reference 参数必填")
		}
	}

	// 验证参考强度
	if req.ImageFidelity != nil && (*req.ImageFidelity < 0 || *req.ImageFidelity > 1) {
		return fmt.Errorf("image_fidelity必须在[0,1]范围内")
	}
	if req.HumanFidelity != nil && (*req.HumanFidelity < 0 || *req.HumanFidelity > 1) {
		return fmt.Errorf("human_fidelity必须在[0,1]范围内")
	}

	// 图生图时不支持负向提示词
	if req.Image != "" && req.NegativePrompt != "" {
		return fmt.Errorf("图生图场景下不支持负向提示词")
	}

	return nil
}

func validateMultiImage2ImageRequest(req *OfficialMultiImage2ImageRequest) error {
	// 验证模型名称（仅支持kling-v2）
	if req.ModelName != ModelKlingV2 {
		return fmt.Errorf("多图参考生图仅支持模型: %s", ModelKlingV2)
	}

	// 验证主体图片数量
	if len(req.SubjectImageList) > 4 {
		return fmt.Errorf("最多支持4张主体图片")
	}

	// 验证生成数量
	if req.N < 1 || req.N > 9 {
		return fmt.Errorf("生成数量必须在[1,9]范围内")
	}

	// 验证纵横比
	validRatios := []string{AspectRatio16x9, AspectRatio9x16, AspectRatio1x1, AspectRatio4x3, AspectRatio3x4, AspectRatio3x2, AspectRatio2x3, AspectRatio21x9}
	if !contains(validRatios, req.AspectRatio) {
		return fmt.Errorf("不支持的纵横比: %s", req.AspectRatio)
	}

	return nil
}

// ===================== 转换函数 =====================

func convertImageToInternalRequest(req *OfficialImageRequest) *KlingTask {
	internalReq := &KlingTask{
		Prompt:         req.Prompt,
		ModelName:      req.ModelName,
		NegativePrompt: req.NegativePrompt,
		Image:          req.Image,
		AspectRatio:    req.AspectRatio,
		CallbackURL:    req.CallbackURL,
		ImageReference: req.ImageReference,
		ImageFidelity:  req.ImageFidelity,
		HumanFidelity:  req.HumanFidelity,
		Resolution:     req.Resolution,
		N:              req.N,
	}

	return internalReq
}

func convertMultiImage2ImageToInternalRequest(req *OfficialMultiImage2ImageRequest) *KlingTask {
	internalReq := &KlingTask{
		Prompt:           req.Prompt,
		ModelName:        req.ModelName,
		AspectRatio:      req.AspectRatio,
		CallbackURL:      req.CallbackURL,
		SubjectImageList: req.SubjectImageList,
		SceneImage:       req.SceneImage,
		StyleImage:       req.StyleImage,
		N:                req.N,
	}

	return internalReq
}

func convertToOfficialImageFormat(task *model.Task) *OfficialImageTaskData {
	// 解析任务数据
	var klingResp KlingResponse[KlingTaskData]
	if task.Data != nil {
		json.Unmarshal(task.Data, &klingResp)
	}

	officialData := &OfficialImageTaskData{
		TaskID:        task.TaskID,
		TaskStatus:    string(task.Status),
		CreatedAt:     task.CreatedAt,
		UpdatedAt:     task.UpdatedAt,
		TaskStatusMsg: task.FailReason,
		TaskInfo: &OfficialTaskInfo{
			ExternalTaskID: task.ExternalTaskID,
		},
	}

	// 如果任务成功，解析生成的图片
	if task.Status == model.TaskStatusSuccess && klingResp.Data.TaskResult != nil {
		var images []OfficialImageResult

		// 根据KlingTaskResult的结构解析视频结果，对于图像生成需要适配
		if taskResult := klingResp.Data.TaskResult; taskResult != nil {
			// 由于KlingTaskResult定义为Videos，但图像生成可能使用不同的结构
			// 这里我们先尝试直接从原始数据中解析
			if task.Data != nil {
				var rawData map[string]interface{}
				if err := json.Unmarshal(task.Data, &rawData); err == nil {
					if data, ok := rawData["data"].(map[string]interface{}); ok {
						if taskRes, ok := data["task_result"].(map[string]interface{}); ok {
							if imageList, exists := taskRes["images"].([]interface{}); exists {
								for i, img := range imageList {
									if imgMap, ok := img.(map[string]interface{}); ok {
										if url, urlOk := imgMap["url"].(string); urlOk {
											images = append(images, OfficialImageResult{
												Index: i,
												URL:   url,
											})
										}
									}
								}
							}
						}
					}
				}
			}
		}

		if len(images) > 0 {
			officialData.TaskResult = &OfficialImageTaskResult{
				Images: images,
			}
		}
	}

	return officialData
}

// ===================== 内部处理函数 =====================

// TaskImageResponse 图像任务响应结构
type TaskImageResponse struct {
	StatusCode int
	Response   interface{}
}

func handleImageTask(c *gin.Context, internalReq *KlingTask, action, externalTaskID string) *TaskImageResponse {
	// 获取用户ID
	userId := c.GetInt("id")

	// 构建模型名称（图像生成使用不同的命名方式）
	modelName := fmt.Sprintf("kling-image_%s", internalReq.ModelName)

	// 通过模型获取渠道
	group := "default" // 默认组
	channel, err := model.ChannelGroup.Next(group, modelName)
	if err != nil {
		return &TaskImageResponse{
			StatusCode: http.StatusServiceUnavailable,
			Response: OfficialImageResponse{
				Code:    503,
				Message: fmt.Sprintf("无可用渠道: %v", err),
			},
		}
	}

	// 获取Provider
	providerFactory := KlingProviderFactory{}
	provider := providerFactory.Create(channel)
	klingProvider, ok := provider.(*KlingProvider)
	if !ok {
		return &TaskImageResponse{
			StatusCode: http.StatusServiceUnavailable,
			Response: OfficialImageResponse{
				Code:    503,
				Message: "Provider类型错误",
			},
		}
	}

	// 提交任务
	resp, errWithCode := klingProvider.Submit("images", action, internalReq)
	if errWithCode != nil {
		return &TaskImageResponse{
			StatusCode: errWithCode.StatusCode,
			Response: OfficialImageResponse{
				Code:    errWithCode.StatusCode,
				Message: errWithCode.Error(),
			},
		}
	}

	// 保存任务到数据库
	now := time.Now().UnixMilli()
	task := &model.Task{
		TaskID:         resp.Data.TaskID,
		ExternalTaskID: externalTaskID,
		Platform:       model.TaskPlatformKling,
		UserId:         userId,
		ChannelId:      channel.Id,
		Action:         action,
		Status:         model.TaskStatusSubmitted,
		Progress:       0,
		SubmitTime:     now,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	// 保存任务数据
	taskData, _ := json.Marshal(resp)
	task.Data = datatypes.JSON(taskData)

	if err := model.CreateTask(task); err != nil {
		logger.SysError(fmt.Sprintf("保存任务失败: %v", err))
	}

	// 转换响应格式
	officialData := &OfficialImageTaskData{
		TaskID:     resp.Data.TaskID,
		TaskStatus: resp.Data.TaskStatus,
		CreatedAt:  resp.Data.CreatedAt,
		UpdatedAt:  resp.Data.UpdatedAt,
		TaskInfo: &OfficialTaskInfo{
			ExternalTaskID: externalTaskID,
		},
	}

	return &TaskImageResponse{
		StatusCode: http.StatusOK,
		Response: OfficialImageResponse{
			Code:      0,
			Message:   "success",
			RequestID: generateRequestID(),
			Data:      officialData,
		},
	}
}

// handleMultiImage2ImageTask 专门处理多图参考生图任务
func handleMultiImage2ImageTask(c *gin.Context, internalReq *KlingTask, action, externalTaskID string) *TaskImageResponse {
	// 获取用户ID
	userId := c.GetInt("id")

	// 构建多图参考生图模型名称
	modelName := fmt.Sprintf("kling-multi-image2image_%s", internalReq.ModelName)

	// 通过模型获取渠道
	group := "default" // 默认组
	channel, err := model.ChannelGroup.Next(group, modelName)
	if err != nil {
		return &TaskImageResponse{
			StatusCode: http.StatusServiceUnavailable,
			Response: OfficialImageResponse{
				Code:    503,
				Message: fmt.Sprintf("无可用渠道: %v", err),
			},
		}
	}

	// 获取Provider
	providerFactory := KlingProviderFactory{}
	provider := providerFactory.Create(channel)
	klingProvider, ok := provider.(*KlingProvider)
	if !ok {
		return &TaskImageResponse{
			StatusCode: http.StatusServiceUnavailable,
			Response: OfficialImageResponse{
				Code:    503,
				Message: "Provider类型错误",
			},
		}
	}

	// 提交任务
	resp, errWithCode := klingProvider.Submit("images", action, internalReq)
	if errWithCode != nil {
		return &TaskImageResponse{
			StatusCode: errWithCode.StatusCode,
			Response: OfficialImageResponse{
				Code:    errWithCode.StatusCode,
				Message: errWithCode.Error(),
			},
		}
	}

	// 保存任务到数据库
	now := time.Now().UnixMilli()
	task := &model.Task{
		TaskID:         resp.Data.TaskID,
		ExternalTaskID: externalTaskID,
		Platform:       model.TaskPlatformKling,
		UserId:         userId,
		ChannelId:      channel.Id,
		Action:         action,
		Status:         model.TaskStatusSubmitted,
		Progress:       0,
		SubmitTime:     now,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	// 保存任务数据
	taskData, _ := json.Marshal(resp)
	task.Data = datatypes.JSON(taskData)

	if err := model.CreateTask(task); err != nil {
		logger.SysError(fmt.Sprintf("保存任务失败: %v", err))
	}

	// 转换响应格式
	officialData := &OfficialImageTaskData{
		TaskID:     resp.Data.TaskID,
		TaskStatus: resp.Data.TaskStatus,
		CreatedAt:  resp.Data.CreatedAt,
		UpdatedAt:  resp.Data.UpdatedAt,
		TaskInfo: &OfficialTaskInfo{
			ExternalTaskID: externalTaskID,
		},
	}

	return &TaskImageResponse{
		StatusCode: http.StatusOK,
		Response: OfficialImageResponse{
			Code:      0,
			Message:   "success",
			RequestID: generateRequestID(),
			Data:      officialData,
		},
	}
}
