package kling

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
    "strings"

    "one-api/common/logger"
    "one-api/common/utils"
    "one-api/model"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

// ===================== 多图参考生视频接口 =====================

// CreateOfficialMultiImage2Video 官方API - 创建多图参考生视频任务
func CreateOfficialMultiImage2Video(c *gin.Context) {
	var req OfficialMultiImage2VideoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, OfficialResponse{
			Code:    400,
			Message: fmt.Sprintf("参数错误: %v", err),
		})
		return
	}

	// 参数验证
	if len(req.ImageList) == 0 {
		c.JSON(http.StatusBadRequest, OfficialResponse{
			Code:    400,
			Message: "image_list是必填参数",
		})
		return
	}

	if req.Prompt == "" {
		c.JSON(http.StatusBadRequest, OfficialResponse{
			Code:    400,
			Message: "prompt是必填参数",
		})
		return
	}

	// 设置默认值
	if req.ModelName == "" {
		req.ModelName = ModelKlingV16MultiImage
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
	if err := validateMultiImage2VideoRequest(&req); err != nil {
		c.JSON(http.StatusBadRequest, OfficialResponse{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	// 转换为内部请求格式
	internalReq := convertMultiImage2VideoToInternal(&req)

	// 调用专门的多图参考生视频处理逻辑
	response := handleMultiImage2VideoTask(c, internalReq, "multi-image2video", req.ExternalTaskID)
	c.JSON(response.StatusCode, response.Response)
}

// GetOfficialMultiImage2VideoTask 官方API - 查询多图参考生视频单个任务
func GetOfficialMultiImage2VideoTask(c *gin.Context) {
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

// ListOfficialMultiImage2VideoTasks 官方API - 查询多图参考生视频任务列表
func ListOfficialMultiImage2VideoTasks(c *gin.Context) {
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "查询任务列表失败",
		})
		return
	}

	// 转换为官方格式
	var officialTasks []*OfficialTaskData
	for _, task := range tasks {
		officialData := convertToOfficialFormat(task)
		officialTasks = append(officialTasks, officialData)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":       0,
		"message":    "success",
		"request_id": generateRequestID(),
		"data":       officialTasks,
	})
}

// ===================== 多模态视频编辑接口 =====================

// InitSelection 初始化待编辑视频
func InitSelection(c *gin.Context) {
	var req OfficialInitSelectionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, OfficialResponse{
			Code:    400,
			Message: fmt.Sprintf("参数错误: %v", err),
		})
		return
	}

	// 参数验证
	if req.VideoID == "" && req.VideoURL == "" {
		c.JSON(http.StatusBadRequest, OfficialResponse{
			Code:    400,
			Message: "video_id和video_url不能同时为空",
		})
		return
	}

	if req.VideoID != "" && req.VideoURL != "" {
		c.JSON(http.StatusBadRequest, OfficialResponse{
			Code:    400,
			Message: "video_id和video_url不能同时有值",
		})
		return
	}

	provider, _, ok := getKlingProviderByModel(c, "kling-init-selection")
	if !ok {
		return
	}

	var resp KlingResponse[*OfficialInitSelectionResponse]
	errWithCode := provider.CallCustomPath(http.MethodPost, "/v1/videos/multi-elements/init-selection", &req, &resp)
	if writeProviderError(c, errWithCode) {
		return
	}

	writeOperationResponse(c, &resp)
}

// AddSelection 增加视频选区
func AddSelection(c *gin.Context) {
	var req OfficialAddSelectionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, OfficialResponse{
			Code:    400,
			Message: fmt.Sprintf("参数错误: %v", err),
		})
		return
	}

	// 参数验证
	if req.SessionID == "" {
		c.JSON(http.StatusBadRequest, OfficialResponse{
			Code:    400,
			Message: "session_id是必填参数",
		})
		return
	}

	if len(req.Points) == 0 {
		c.JSON(http.StatusBadRequest, OfficialResponse{
			Code:    400,
			Message: "points是必填参数",
		})
		return
	}

	if len(req.Points) > 10 {
		c.JSON(http.StatusBadRequest, OfficialResponse{
			Code:    400,
			Message: "单帧最多可标记10个点",
		})
		return
	}

	// 验证坐标范围
	for i, point := range req.Points {
		if point.X < 0 || point.X > 1 || point.Y < 0 || point.Y > 1 {
			c.JSON(http.StatusBadRequest, OfficialResponse{
				Code:    400,
				Message: fmt.Sprintf("第%d个坐标点超出范围[0,1]", i+1),
			})
			return
		}
	}

	provider, _, ok := getKlingProviderByModel(c, "kling-add-selection")
	if !ok {
		return
	}

	var resp KlingResponse[*OfficialSelectionResponse]
	errWithCode := provider.CallCustomPath(http.MethodPost, "/v1/videos/multi-elements/add-selection", &req, &resp)
	if writeProviderError(c, errWithCode) {
		return
	}

	writeOperationResponse(c, &resp)
}

// DeleteSelection 删减视频选区
func DeleteSelection(c *gin.Context) {
	var req OfficialDeleteSelectionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, OfficialResponse{
			Code:    400,
			Message: fmt.Sprintf("参数错误: %v", err),
		})
		return
	}

	// 参数验证逻辑与AddSelection类似
	if req.SessionID == "" {
		c.JSON(http.StatusBadRequest, OfficialResponse{
			Code:    400,
			Message: "session_id是必填参数",
		})
		return
	}

	provider, _, ok := getKlingProviderByModel(c, "kling-delete-selection")
	if !ok {
		return
	}

	var resp KlingResponse[*OfficialSelectionResponse]
	errWithCode := provider.CallCustomPath(http.MethodPost, "/v1/videos/multi-elements/delete-selection", &req, &resp)
	if writeProviderError(c, errWithCode) {
		return
	}

	writeOperationResponse(c, &resp)
}

// ClearSelection 清除视频选区
func ClearSelection(c *gin.Context) {
	var req OfficialClearSelectionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, OfficialResponse{
			Code:    400,
			Message: fmt.Sprintf("参数错误: %v", err),
		})
		return
	}

	if req.SessionID == "" {
		c.JSON(http.StatusBadRequest, OfficialResponse{
			Code:    400,
			Message: "session_id是必填参数",
		})
		return
	}

	provider, _, ok := getKlingProviderByModel(c, "kling-clear-selection")
	if !ok {
		return
	}

	var resp KlingResponse[*OfficialClearSelectionResponse]
	errWithCode := provider.CallCustomPath(http.MethodPost, "/v1/videos/multi-elements/clear-selection", &req, &resp)
	if writeProviderError(c, errWithCode) {
		return
	}

	writeOperationResponse(c, &resp)
}

// PreviewSelection 预览已选区视频
func PreviewSelection(c *gin.Context) {
	var req OfficialPreviewSelectionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, OfficialResponse{
			Code:    400,
			Message: fmt.Sprintf("参数错误: %v", err),
		})
		return
	}

	if req.SessionID == "" {
		c.JSON(http.StatusBadRequest, OfficialResponse{
			Code:    400,
			Message: "session_id是必填参数",
		})
		return
	}

	provider, _, ok := getKlingProviderByModel(c, "kling-preview-selection")
	if !ok {
		return
	}

	var resp KlingResponse[*OfficialPreviewSelectionResponse]
	errWithCode := provider.CallCustomPath(http.MethodPost, "/v1/videos/multi-elements/preview-selection", &req, &resp)
	if writeProviderError(c, errWithCode) {
		return
	}

	writeOperationResponse(c, &resp)
}

// CreateMultiElementsTask 创建多模态视频编辑任务
func CreateMultiElementsTask(c *gin.Context) {
	var req OfficialMultiElementsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, OfficialResponse{
			Code:    400,
			Message: fmt.Sprintf("参数错误: %v", err),
		})
		return
	}

	// 参数验证
	if req.SessionID == "" {
		c.JSON(http.StatusBadRequest, OfficialResponse{
			Code:    400,
			Message: "session_id是必填参数",
		})
		return
	}

	if req.EditMode == "" {
		c.JSON(http.StatusBadRequest, OfficialResponse{
			Code:    400,
			Message: "edit_mode是必填参数",
		})
		return
	}

	if req.Prompt == "" {
		c.JSON(http.StatusBadRequest, OfficialResponse{
			Code:    400,
			Message: "prompt是必填参数",
		})
		return
	}

	// 设置默认值
	if req.ModelName == "" {
		req.ModelName = ModelKlingV16MultiImage
	}
	if req.Mode == "" {
		req.Mode = ModeStd
	}
	if req.Duration == "" {
		req.Duration = Duration5s
	}

	// 参数验证
	if err := validateMultiElementsRequest(&req); err != nil {
		c.JSON(http.StatusBadRequest, OfficialResponse{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	internalReq := convertMultiElementsToInternalRequest(&req)

	modelName := fmt.Sprintf("kling-multi-elements_%s_%s_%s", internalReq.ModelName, internalReq.Mode, internalReq.Duration)
	provider, channel, ok := getKlingProviderByModel(c, modelName)
	if !ok {
		return
	}

	resp, errWithCode := provider.Submit("videos", "multi-elements", internalReq)
	if writeProviderError(c, errWithCode) {
		return
	}

	now := time.Now().UnixMilli()
	taskID := ""
	taskStatus := TaskStatusSubmitted
	var taskStatusMsg string
	var sessionID string
	var taskResult *OfficialMultiTaskResult
	createdAt := now
	updatedAt := now
	if resp != nil {
		taskID = resp.Data.TaskID
		if resp.Data.TaskStatus != "" {
			taskStatus = resp.Data.TaskStatus
		}
		taskStatusMsg = resp.Data.TaskStatusMsg
		sessionID = resp.Data.SessionID
		if resp.Data.CreatedAt != 0 {
			createdAt = resp.Data.CreatedAt
		}
		if resp.Data.UpdatedAt != 0 {
			updatedAt = resp.Data.UpdatedAt
		}
		if resp.Data.TaskResult != nil && len(resp.Data.TaskResult.Videos) > 0 {
			videos := make([]OfficialMultiVideoResult, len(resp.Data.TaskResult.Videos))
			for i, video := range resp.Data.TaskResult.Videos {
				videos[i] = OfficialMultiVideoResult{
					ID:       video.ID,
					URL:      video.URL,
					Duration: video.Duration,
				}
			}
			taskResult = &OfficialMultiTaskResult{Videos: videos}
		}
	}

	if taskID == "" {
		taskID = generateTaskID()
	}

	userId := c.GetInt("id")
    task := &model.Task{
        PlatformTaskID: utils.NewPlatformULID(),
        TaskID:         taskID,
		ExternalTaskID: req.ExternalTaskID,
		Platform:       model.TaskPlatformKling,
		UserId:         userId,
		ChannelId:      channel.Id,
		Action:         "multi-elements",
		Status:         model.TaskStatusSubmitted,
		Progress:       0,
		SubmitTime:     createdAt,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}

	if resp != nil {
		taskData, _ := json.Marshal(resp)
		task.Data = datatypes.JSON(taskData)
	}

	if err := model.CreateTask(task); err != nil {
		logger.SysError(fmt.Sprintf("保存多模态任务失败: %v", err))
	}

    // 统一以平台任务ID返回
    platformID := utils.AddTaskPrefix(task.PlatformTaskID)
    officialData := &OfficialMultiElementsTaskData{
        TaskID:        platformID,
        TaskStatus:    taskStatus,
        TaskStatusMsg: taskStatusMsg,
        SessionID:     sessionID,
        CreatedAt:     createdAt,
        UpdatedAt:     updatedAt,
		TaskInfo: &OfficialTaskInfo{
			ExternalTaskID: req.ExternalTaskID,
		},
		TaskResult: taskResult,
	}

	c.JSON(http.StatusOK, gin.H{
		"code":       0,
		"message":    "success",
		"request_id": generateRequestID(),
		"data":       officialData,
	})
}

// GetMultiElementsTask 查询多模态视频编辑单个任务
func GetMultiElementsTask(c *gin.Context) {
    taskID := c.Param("id")
    userId := c.GetInt("id")
    actions := resolveActionsByPath(c.FullPath())

	// 支持通过external_task_id查询
	var task *model.Task
	var err error

    // 平台任务ID支持：task_<ULID> 或历史 base36
    if strings.HasPrefix(strings.ToLower(taskID), utils.PlatformTaskPrefix) {
        if t, _ := model.GetTaskByPlatformTaskID(model.TaskPlatformKling, userId, utils.StripTaskPrefix(taskID)); t != nil {
            taskID = t.TaskID
        }
    } else if id, ok := utils.DecodePlatformTaskID(taskID); ok {
        if t, _ := model.GetTaskByID(id); t != nil && t.Platform == model.TaskPlatformKling && t.UserId == userId {
            taskID = t.TaskID
        }
    } else if utils.IsULID(taskID) {
        if t, _ := model.GetTaskByPlatformTaskID(model.TaskPlatformKling, userId, taskID); t != nil {
            taskID = t.TaskID
        }
    }

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
		task = tasks[0]
	}

	// 转换为多模态编辑任务格式
	officialData := convertToMultiElementsFormat(task)

	c.JSON(http.StatusOK, gin.H{
		"code":       0,
		"message":    "success",
		"request_id": generateRequestID(),
		"data":       officialData,
	})
}

// ListMultiElementsTasks 查询多模态视频编辑任务列表
func ListMultiElementsTasks(c *gin.Context) {
	var query OfficialTaskListQuery

	// 解析查询参数
	if err := c.ShouldBindQuery(&query); err != nil {
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
	var officialTasks []*OfficialMultiElementsTaskData
	for _, task := range tasks {
		officialData := convertToMultiElementsFormat(task)
		officialTasks = append(officialTasks, officialData)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":       0,
		"message":    "success",
		"request_id": generateRequestID(),
		"data":       officialTasks,
	})
}

// ===================== 验证函数 =====================

func validateMultiImage2VideoRequest(req *OfficialMultiImage2VideoRequest) error {
	// 验证模型名称
	if req.ModelName != ModelKlingV16MultiImage {
		return fmt.Errorf("多图参考生视频仅支持模型: %s", ModelKlingV16MultiImage)
	}

	// 验证图片数量
	if len(req.ImageList) > 4 {
		return fmt.Errorf("最多支持4张图片")
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

	return nil
}

func validateMultiElementsRequest(req *OfficialMultiElementsRequest) error {
	// 验证模型名称
	if req.ModelName != ModelKlingV16MultiImage {
		return fmt.Errorf("多模态视频编辑仅支持模型: %s", ModelKlingV16MultiImage)
	}

	// 验证编辑模式
	validEditModes := []string{EditModeAddition, EditModeSwap, EditModeRemoval}
	if !contains(validEditModes, req.EditMode) {
		return fmt.Errorf("不支持的编辑模式: %s", req.EditMode)
	}

	// 根据编辑模式验证图片列表
	switch req.EditMode {
	case EditModeAddition:
		if len(req.ImageList) < 1 || len(req.ImageList) > 2 {
			return fmt.Errorf("增加元素模式需要1~2张图片")
		}
	case EditModeSwap:
		if len(req.ImageList) != 1 {
			return fmt.Errorf("替换元素模式需要1张图片")
		}
	case EditModeRemoval:
		if len(req.ImageList) > 0 {
			return fmt.Errorf("删除元素模式不需要图片")
		}
	}

	// 验证模式
	if req.Mode != ModeStd && req.Mode != ModePro {
		return fmt.Errorf("不支持的模式: %s", req.Mode)
	}

	// 验证时长
	if req.Duration != Duration5s && req.Duration != Duration10s {
		return fmt.Errorf("不支持的时长: %s", req.Duration)
	}

	return nil
}

// ===================== 转换函数 =====================

func convertMultiImage2VideoToInternal(req *OfficialMultiImage2VideoRequest) *KlingTask {
	internalReq := &KlingTask{
		Prompt:      req.Prompt,
		ModelName:   req.ModelName,
		Mode:        req.Mode,
		AspectRatio: req.AspectRatio,
		Duration:    req.Duration,
		CallbackURL: req.CallbackURL,
		ImageList:   req.ImageList,
	}

	return internalReq
}

func convertMultiElementsToInternalRequest(req *OfficialMultiElementsRequest) *KlingTask {
	return &KlingTask{
		SessionID:      req.SessionID,
		EditMode:       req.EditMode,
		ImageList:      req.ImageList,
		Prompt:         req.Prompt,
		NegativePrompt: req.NegativePrompt,
		ModelName:      req.ModelName,
		Mode:           req.Mode,
		Duration:       req.Duration,
		CallbackURL:    req.CallbackURL,
	}
}

func convertToMultiElementsFormat(task *model.Task) *OfficialMultiElementsTaskData {
	// 解析任务数据
	var klingResp KlingResponse[KlingTaskData]
	if task.Data != nil {
		json.Unmarshal(task.Data, &klingResp)
	}

    platformID := utils.AddTaskPrefix(task.PlatformTaskID)
    officialData := &OfficialMultiElementsTaskData{
        TaskID:        platformID,
        TaskStatus:    string(task.Status),
        CreatedAt:     task.CreatedAt,
        UpdatedAt:     task.UpdatedAt,
		TaskStatusMsg: task.FailReason,
		TaskInfo: &OfficialTaskInfo{
			ExternalTaskID: task.ExternalTaskID,
		},
	}

	if klingResp.Data.SessionID != "" {
		officialData.SessionID = klingResp.Data.SessionID
	}

	if klingResp.Data.TaskStatus != "" {
		officialData.TaskStatus = klingResp.Data.TaskStatus
	}

	if klingResp.Data.TaskStatusMsg != "" {
		officialData.TaskStatusMsg = klingResp.Data.TaskStatusMsg
	}

	if klingResp.Data.CreatedAt != 0 {
		officialData.CreatedAt = klingResp.Data.CreatedAt
	}

	if klingResp.Data.UpdatedAt != 0 {
		officialData.UpdatedAt = klingResp.Data.UpdatedAt
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
		officialResult := &OfficialMultiTaskResult{
			Videos: make([]OfficialMultiVideoResult, len(klingResp.Data.TaskResult.Videos)),
		}

		for i, video := range klingResp.Data.TaskResult.Videos {
			officialResult.Videos[i] = OfficialMultiVideoResult{
				ID:        video.ID,
				SessionID: klingResp.Data.SessionID,
				URL:       video.URL,
				Duration:  video.Duration,
			}
		}

		officialData.TaskResult = officialResult
	}

	return officialData
}

// ===================== 工具函数 =====================

func generateSessionID() string {
	return fmt.Sprintf("session_%d_%s", time.Now().UnixNano(), generateRandomString(8))
}

func generateTaskID() string {
	return fmt.Sprintf("task_%d_%s", time.Now().UnixNano(), generateRandomString(12))
}

// handleMultiImage2VideoTask 专门处理多图参考生视频任务
func handleMultiImage2VideoTask(c *gin.Context, internalReq *KlingTask, action, externalTaskID string) *TaskResponse {
	// 获取用户ID
	userId := c.GetInt("id")

	// 构建多图参考生视频模型名称
	modelName := fmt.Sprintf("kling-multi-image2video_%s_%s_%s", internalReq.ModelName, internalReq.Mode, internalReq.Duration)

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
	officialData := &OfficialTaskData{
		TaskID:     resp.Data.TaskID,
		TaskStatus: resp.Data.TaskStatus,
		CreatedAt:  resp.Data.CreatedAt,
		UpdatedAt:  resp.Data.UpdatedAt,
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
