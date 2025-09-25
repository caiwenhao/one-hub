package kling

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"one-api/common/logger"
	"one-api/model"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

// CreateOfficialText2Video 官方API - 创建文生视频任务
func CreateOfficialText2Video(c *gin.Context) {
	var req OfficialText2VideoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, OfficialResponse{
			Code:    400,
			Message: fmt.Sprintf("参数错误: %v", err),
		})
		return
	}

	// 参数验证
	if req.Prompt == "" {
		c.JSON(http.StatusBadRequest, OfficialResponse{
			Code:    400,
			Message: "prompt是必填参数",
		})
		return
	}

	// 设置默认值
	if req.ModelName == "" {
		req.ModelName = ModelKlingV1
	}
	if req.Mode == "" {
		req.Mode = ModeStd
	}
	if req.Duration == "" {
		req.Duration = Duration5s
	}
	if req.AspectRatio == "" {
		req.AspectRatio = AspectRatio16x9
	}

	// 参数验证
	if err := validateText2VideoRequest(&req); err != nil {
		c.JSON(http.StatusBadRequest, OfficialResponse{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	// 转换为内部请求格式
	internalReq := convertToInternalRequest(&req, nil)

	// 调用内部处理逻辑
	response := handleVideoTask(c, internalReq, "text2video", req.ExternalTaskID)
	c.JSON(response.StatusCode, response.Response)
}

// CreateOfficialImage2Video 官方API - 创建图生视频任务
func CreateOfficialImage2Video(c *gin.Context) {
	var req OfficialImage2VideoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, OfficialResponse{
			Code:    400,
			Message: fmt.Sprintf("参数错误: %v", err),
		})
		return
	}

	// 参数验证
	if req.Image == "" {
		c.JSON(http.StatusBadRequest, OfficialResponse{
			Code:    400,
			Message: "image是必填参数",
		})
		return
	}

	// 设置默认值
	if req.ModelName == "" {
		req.ModelName = ModelKlingV1
	}
	if req.Mode == "" {
		req.Mode = ModeStd
	}
	if req.Duration == "" {
		req.Duration = Duration5s
	}
	if req.AspectRatio == "" {
		req.AspectRatio = AspectRatio16x9
	}

	// 参数验证
	if err := validateImage2VideoRequest(&req); err != nil {
		c.JSON(http.StatusBadRequest, OfficialResponse{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	// 转换为内部请求格式
	internalReq := convertToInternalRequest(nil, &req)

	// 调用内部处理逻辑
	response := handleVideoTask(c, internalReq, "image2video", req.ExternalTaskID)
	c.JSON(response.StatusCode, response.Response)
}

// GetOfficialTask 官方API - 查询单个任务
func GetOfficialTask(c *gin.Context) {
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
			c.JSON(http.StatusNotFound, OfficialResponse{
				Code:    404,
				Message: "任务不存在",
			})
			return
		}
		task = tasks[0] // 取第一个匹配的任务
	}

	// 转换为官方格式
	officialData := convertToOfficialFormat(task)

	c.JSON(http.StatusOK, OfficialResponse{
		Code:      0,
		Message:   "success",
		RequestID: generateRequestID(),
		Data:      officialData,
	})
}

// ListOfficialTasks 官方API - 查询任务列表
func ListOfficialTasks(c *gin.Context) {
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
		c.JSON(http.StatusInternalServerError, OfficialResponse{
			Code:    500,
			Message: "查询任务列表失败",
		})
		return
	}

	// 转换为官方格式
	var officialTasks []*OfficialTaskData
	for _, task := range tasks {
		officialData := convertToOfficialFormat(task)
		officialTasks = append(officialTasks, officialData)
	}

	c.JSON(http.StatusOK, OfficialResponse{
		Code:      0,
		Message:   "success",
		RequestID: generateRequestID(),
		DataList:  officialTasks,
	})
}

// 参数验证函数
func validateText2VideoRequest(req *OfficialText2VideoRequest) error {
	// 验证模型名称
	validModels := []string{ModelKlingV1, ModelKlingV15, ModelKlingV16, ModelKlingV2Master, ModelKlingV21, ModelKlingV21Master}
	if !contains(validModels, req.ModelName) {
		return fmt.Errorf("不支持的模型: %s", req.ModelName)
	}

	// 验证模式
	if req.Mode != ModeStd && req.Mode != ModePro {
		return fmt.Errorf("不支持的模式: %s", req.Mode)
	}

	// 验证时长
	if req.Duration != Duration5s && req.Duration != Duration10s {
		return fmt.Errorf("不支持的时长: %s", req.Duration)
	}

	// 验证纵横比
	validRatios := []string{AspectRatio16x9, AspectRatio9x16, AspectRatio1x1}
	if !contains(validRatios, req.AspectRatio) {
		return fmt.Errorf("不支持的纵横比: %s", req.AspectRatio)
	}

	// 验证cfg_scale
	if req.CfgScale != nil && (*req.CfgScale < 0 || *req.CfgScale > 1) {
		return fmt.Errorf("cfg_scale必须在[0,1]范围内")
	}

	// 验证摄像机控制
	if req.CameraControl != nil {
		if err := validateCameraControl(req.CameraControl); err != nil {
			return err
		}
	}

	return nil
}

func validateImage2VideoRequest(req *OfficialImage2VideoRequest) error {
	// 验证模型名称
	validModels := []string{ModelKlingV1, ModelKlingV15, ModelKlingV16, ModelKlingV2Master, ModelKlingV21, ModelKlingV21Master}
	if !contains(validModels, req.ModelName) {
		return fmt.Errorf("不支持的模型: %s", req.ModelName)
	}

	// 验证模式
	if req.Mode != ModeStd && req.Mode != ModePro {
		return fmt.Errorf("不支持的模式: %s", req.Mode)
	}

	// 验证时长
	if req.Duration != Duration5s && req.Duration != Duration10s {
		return fmt.Errorf("不支持的时长: %s", req.Duration)
	}

	// 验证纵横比
	validRatios := []string{AspectRatio16x9, AspectRatio9x16, AspectRatio1x1}
	if !contains(validRatios, req.AspectRatio) {
		return fmt.Errorf("不支持的纵横比: %s", req.AspectRatio)
	}

	// 验证cfg_scale
	if req.CfgScale != nil && (*req.CfgScale < 0 || *req.CfgScale > 1) {
		return fmt.Errorf("cfg_scale必须在[0,1]范围内")
	}

	// 验证摄像机控制
	if req.CameraControl != nil {
		if err := validateCameraControl(req.CameraControl); err != nil {
			return err
		}
	}

	// 验证动态笔刷
	if len(req.DynamicMasks) > 0 {
		if err := validateDynamicMasks(req.DynamicMasks); err != nil {
			return err
		}
	}

	// 验证互斥参数组合
	if err := validateImage2VideoConstraints(req); err != nil {
		return err
	}

	return nil
}

// 验证图生视频的约束条件
func validateImage2VideoConstraints(req *OfficialImage2VideoRequest) error {
	// image 参数与 image_tail 参数至少二选一
	if req.Image == "" && req.ImageTail == "" {
		return fmt.Errorf("image 参数与 image_tail 参数至少二选一，二者不能同时为空")
	}

	// image + image_tail参数、dynamic_masks/static_mask参数、camera_control参数三选一
	hasImageTail := req.ImageTail != ""
	hasMasks := req.StaticMask != "" || len(req.DynamicMasks) > 0
	hasCameraControl := req.CameraControl != nil

	combinationCount := 0
	if hasImageTail {
		combinationCount++
	}
	if hasMasks {
		combinationCount++
	}
	if hasCameraControl {
		combinationCount++
	}

	if combinationCount > 1 {
		return fmt.Errorf("image + image_tail参数、dynamic_masks/static_mask参数、camera_control参数三选一，不能同时使用")
	}

	return nil
}

// 验证动态笔刷参数
func validateDynamicMasks(dynamicMasks []DynamicMask) error {
	if len(dynamicMasks) > 6 {
		return fmt.Errorf("动态笔刷配置最多6组")
	}

	for i, mask := range dynamicMasks {
		if len(mask.Trajectories) < 2 {
			return fmt.Errorf("第%d组动态笔刷的轨迹点至少需要2个", i+1)
		}
		if len(mask.Trajectories) > 77 {
			return fmt.Errorf("第%d组动态笔刷的轨迹点最多77个", i+1)
		}
	}

	return nil
}

func validateCameraControl(cc *CameraControl) error {
	validTypes := []string{CameraTypeSimple, CameraTypeDownBack, CameraTypeForwardUp, CameraTypeRightTurnForward, CameraTypeLeftTurnForward}
	if !contains(validTypes, cc.Type) {
		return fmt.Errorf("不支持的摄像机控制类型: %s", cc.Type)
	}

	if cc.Type == CameraTypeSimple {
		if cc.Config == nil {
			return fmt.Errorf("simple类型摄像机控制必须提供config参数")
		}

		// 验证只能有一个参数不为0
		nonZeroCount := 0
		if cc.Config.Horizontal != nil && *cc.Config.Horizontal != 0 {
			nonZeroCount++
		}
		if cc.Config.Vertical != nil && *cc.Config.Vertical != 0 {
			nonZeroCount++
		}
		if cc.Config.Pan != nil && *cc.Config.Pan != 0 {
			nonZeroCount++
		}
		if cc.Config.Tilt != nil && *cc.Config.Tilt != 0 {
			nonZeroCount++
		}
		if cc.Config.Roll != nil && *cc.Config.Roll != 0 {
			nonZeroCount++
		}
		if cc.Config.Zoom != nil && *cc.Config.Zoom != 0 {
			nonZeroCount++
		}

		if nonZeroCount != 1 {
			return fmt.Errorf("simple类型摄像机控制必须且只能有一个参数不为0")
		}

		// 验证参数范围
		if cc.Config.Horizontal != nil && (*cc.Config.Horizontal < -10 || *cc.Config.Horizontal > 10) {
			return fmt.Errorf("horizontal参数必须在[-10,10]范围内")
		}
		if cc.Config.Vertical != nil && (*cc.Config.Vertical < -10 || *cc.Config.Vertical > 10) {
			return fmt.Errorf("vertical参数必须在[-10,10]范围内")
		}
		if cc.Config.Pan != nil && (*cc.Config.Pan < -10 || *cc.Config.Pan > 10) {
			return fmt.Errorf("pan参数必须在[-10,10]范围内")
		}
		if cc.Config.Tilt != nil && (*cc.Config.Tilt < -10 || *cc.Config.Tilt > 10) {
			return fmt.Errorf("tilt参数必须在[-10,10]范围内")
		}
		if cc.Config.Roll != nil && (*cc.Config.Roll < -10 || *cc.Config.Roll > 10) {
			return fmt.Errorf("roll参数必须在[-10,10]范围内")
		}
		if cc.Config.Zoom != nil && (*cc.Config.Zoom < -10 || *cc.Config.Zoom > 10) {
			return fmt.Errorf("zoom参数必须在[-10,10]范围内")
		}
	} else {
		// 非simple类型不需要config参数
		cc.Config = nil
	}

	return nil
}

// 内部处理逻辑
func handleVideoTask(c *gin.Context, internalReq *KlingTask, action, externalTaskID string) *TaskResponse {
	// 获取用户ID
	userId := c.GetInt("id")

	// 构建模型名称
	modelName := fmt.Sprintf("kling-video_%s_%s_%s", internalReq.ModelName, internalReq.Mode, internalReq.Duration)

	// 通过模型获取渠道
	group := "default" // 默认组
	channel, err := model.ChannelGroup.Next(group, modelName)
	if err != nil {
		return &TaskResponse{
			StatusCode: http.StatusServiceUnavailable,
			Response: OfficialResponse{
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
		return &TaskResponse{
			StatusCode: http.StatusServiceUnavailable,
			Response: OfficialResponse{
				Code:    503,
				Message: "Provider类型错误",
			},
		}
	}

	// 提交任务
	resp, errWithCode := klingProvider.Submit("videos", action, internalReq)
	if errWithCode != nil {
		return &TaskResponse{
			StatusCode: errWithCode.StatusCode,
			Response: OfficialResponse{
				Code:    errWithCode.StatusCode,
				Message: errWithCode.Error(),
			},
		}
	}

	// 兜底生成任务标识与时间，确保响应结构与官方一致
	now := time.Now().UnixMilli()
	taskID := generateTaskID()
	taskStatus := TaskStatusSubmitted
	createdAt := now
	updatedAt := now
	var officialResult *OfficialTaskResult

	if resp != nil {
		if resp.Data.TaskID != "" {
			taskID = resp.Data.TaskID
		}
		if resp.Data.TaskStatus != "" {
			taskStatus = resp.Data.TaskStatus
		}
		if resp.Data.CreatedAt != 0 {
			createdAt = resp.Data.CreatedAt
		}
		if resp.Data.UpdatedAt != 0 {
			updatedAt = resp.Data.UpdatedAt
		}
		if resp.Data.TaskResult != nil && len(resp.Data.TaskResult.Videos) > 0 {
			officialResult = &OfficialTaskResult{
				Videos: make([]OfficialVideoResult, len(resp.Data.TaskResult.Videos)),
			}
			for i, video := range resp.Data.TaskResult.Videos {
				officialResult.Videos[i] = OfficialVideoResult{
					ID:       video.ID,
					URL:      video.URL,
					Duration: video.Duration,
				}
			}
		}
	}

	// 落库任务
	task := &model.Task{
		TaskID:         taskID,
		ExternalTaskID: externalTaskID,
		Platform:       model.TaskPlatformKling,
		UserId:         userId,
		ChannelId:      channel.Id,
		Action:         action,
		Status:         model.TaskStatusSubmitted,
		Progress:       0,
		SubmitTime:     createdAt,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}

	// 保存任务数据（无返回时组装一个最小响应）
	if resp == nil {
		resp = &KlingResponse[KlingTaskData]{
			Code:    0,
			Message: "success",
			Data: KlingTaskData{
				TaskID:     taskID,
				TaskStatus: taskStatus,
				CreatedAt:  createdAt,
				UpdatedAt:  updatedAt,
			},
		}
	} else {
		resp.Data.TaskID = taskID
		resp.Data.TaskStatus = taskStatus
		resp.Data.CreatedAt = createdAt
		resp.Data.UpdatedAt = updatedAt
	}

	taskData, _ := json.Marshal(resp)
	task.Data = datatypes.JSON(taskData)

	if err := model.CreateTask(task); err != nil {
		logger.SysError(fmt.Sprintf("保存任务失败: %v", err))
	}

	// 转换响应格式
	officialData := &OfficialTaskData{
		TaskID:     taskID,
		TaskStatus: taskStatus,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
		TaskResult: officialResult,
		TaskInfo: &OfficialTaskInfo{
			ExternalTaskID: externalTaskID,
		},
	}

	return &TaskResponse{
		StatusCode: http.StatusOK,
		Response: OfficialResponse{
			Code:      0,
			Message:   "success",
			RequestID: generateRequestID(),
			Data:      officialData,
		},
	}
}

// 转换函数
func convertToInternalRequest(text2VideoReq *OfficialText2VideoRequest, image2VideoReq *OfficialImage2VideoRequest) *KlingTask {
	internalReq := &KlingTask{}

	if text2VideoReq != nil {
		internalReq.Prompt = text2VideoReq.Prompt
		internalReq.ModelName = text2VideoReq.ModelName
		internalReq.Mode = text2VideoReq.Mode
		internalReq.NegativePrompt = text2VideoReq.NegativePrompt
		internalReq.CfgScale = text2VideoReq.CfgScale
		internalReq.CameraControl = text2VideoReq.CameraControl
		internalReq.AspectRatio = text2VideoReq.AspectRatio
		internalReq.Duration = text2VideoReq.Duration
		internalReq.CallbackURL = text2VideoReq.CallbackURL
	}

	if image2VideoReq != nil {
		internalReq.Image = image2VideoReq.Image
		internalReq.ImageTail = image2VideoReq.ImageTail
		internalReq.StaticMask = image2VideoReq.StaticMask
		internalReq.DynamicMasks = image2VideoReq.DynamicMasks
		internalReq.Prompt = image2VideoReq.Prompt
		internalReq.ModelName = image2VideoReq.ModelName
		internalReq.Mode = image2VideoReq.Mode
		internalReq.NegativePrompt = image2VideoReq.NegativePrompt
		internalReq.CfgScale = image2VideoReq.CfgScale
		internalReq.CameraControl = image2VideoReq.CameraControl
		internalReq.AspectRatio = image2VideoReq.AspectRatio
		internalReq.Duration = image2VideoReq.Duration
		internalReq.CallbackURL = image2VideoReq.CallbackURL
	}

	return internalReq
}

func convertToOfficialFormat(task *model.Task) *OfficialTaskData {
	// 解析任务数据
	var klingResp KlingResponse[KlingTaskData]
	if task.Data != nil {
		json.Unmarshal(task.Data, &klingResp)
	}

	officialData := &OfficialTaskData{
		TaskID:        task.TaskID,
		TaskStatus:    string(task.Status),
		CreatedAt:     task.CreatedAt,
		UpdatedAt:     task.UpdatedAt,
		TaskStatusMsg: task.FailReason,
		TaskInfo: &OfficialTaskInfo{
			ExternalTaskID: task.ExternalTaskID,
		},
	}

	// 转换任务状态
	switch task.Status {
	case model.TaskStatusSubmitted, model.TaskStatusQueued:
		officialData.TaskStatus = TaskStatusSubmitted
	case model.TaskStatusInProgress:
		officialData.TaskStatus = TaskStatusProcessing
	case model.TaskStatusSuccess:
		officialData.TaskStatus = TaskStatusSucceed
	case model.TaskStatusFailure:
		officialData.TaskStatus = TaskStatusFailed
	}

	// 转换任务结果
	if klingResp.Data.TaskResult != nil && len(klingResp.Data.TaskResult.Videos) > 0 {
		officialResult := &OfficialTaskResult{
			Videos: make([]OfficialVideoResult, len(klingResp.Data.TaskResult.Videos)),
		}

		for i, video := range klingResp.Data.TaskResult.Videos {
			officialResult.Videos[i] = OfficialVideoResult{
				ID:       video.ID,
				URL:      video.URL,
				Duration: video.Duration,
			}
		}

		officialData.TaskResult = officialResult
	}

	return officialData
}

// 工具函数
type TaskResponse struct {
	StatusCode int
	Response   OfficialResponse
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func generateRequestID() string {
	return fmt.Sprintf("req_%d_%s", time.Now().UnixNano(), generateRandomString(8))
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	buffer := make([]byte, length)
	if _, err := rand.Read(buffer); err != nil {
		now := time.Now().UnixNano()
		for i := range result {
			result[i] = charset[now%int64(len(charset))]
			now = now / int64(len(charset))
		}
		return string(result)
	}
	for i := range result {
		result[i] = charset[int(buffer[i])%len(charset)]
	}
	return string(result)
}
