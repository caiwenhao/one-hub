package router

import (
	"github.com/gin-gonic/gin"

	miniapi "one-api/providers/minimaxi"
	"one-api/types"
)

// nolint:revive // 仅用于 swagger 注释类型引用
var (
	_ miniapi.SpeechResponse
	_ types.MiniMaxSpeechRequest
	_ types.MiniMaxAsyncSpeechRequest
	_ types.MiniMaxAsyncSpeechResponse
	_ types.MiniMaxAsyncSpeechQueryResponse
	_ miniapi.MiniMaxVideoCreateRequest
	_ miniapi.MiniMaxVideoCreateResponse
	_ miniapi.MiniMaxVideoQueryResponse
	_ miniapi.MiniMaxFileRetrieveResponse
)

// docMiniMaxSyncSpeech 仅用于 Swagger 文档描述。
// @Summary      minimaxi 官方接口 - 同步文本转语音
// @Description  与 minimaxi `/v1/t2a_v2` 保持一致，默认返回 JSON；当渠道配置 `audio_mode=hex` 时返回裸音频流。
// @Tags         MiniMax Speech
// @Accept       json
// @Produce      json
// @Produce      audio/mpeg
// @Security     BearerAuth
// @Param        request body types.MiniMaxSpeechRequest true "同步文本转语音请求体"
// @Success      200 {object} miniapi.SpeechResponse
// @Router       /minimaxi/v1/t2a_v2 [post]
func docMiniMaxSyncSpeech(*gin.Context) {}

// docMiniMaxAsyncSpeechCreate 仅用于 Swagger 文档描述。
// @Summary      minimaxi 官方接口 - 创建异步文本转语音任务
// @Description  透传 minimaxi `/v1/t2a_async_v2`，支持 text 或 text_file_id 方式提交长文本。
// @Tags         MiniMax Speech
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body types.MiniMaxAsyncSpeechRequest true "异步文本转语音请求体"
// @Success      200 {object} types.MiniMaxAsyncSpeechResponse
// @Router       /minimaxi/v1/t2a_async_v2 [post]
func docMiniMaxAsyncSpeechCreate(*gin.Context) {}

// docMiniMaxAsyncSpeechQuery 仅用于 Swagger 文档描述。
// @Summary      minimaxi 官方接口 - 查询异步文本转语音任务
// @Description  默认按 Token 分组遍历可用渠道，无需传递 channel_id。
// @Tags         MiniMax Speech
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        task_id    query string true  "任务 ID"
// @Success      200 {object} types.MiniMaxAsyncSpeechQueryResponse
// @Router       /minimaxi/v1/query/t2a_async_query_v2 [get]
func docMiniMaxAsyncSpeechQuery(*gin.Context) {}

// docMiniMaxAsyncSpeechCreateAlias 用于 OpenAI 兼容前缀文档描述。
// @Summary      minimaxi 官方接口 - 创建异步文本转语音任务（别名）
// @Tags         MiniMax Speech
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body types.MiniMaxAsyncSpeechRequest true "异步文本转语音请求体"
// @Success      200 {object} types.MiniMaxAsyncSpeechResponse
// @Router       /v1/t2a_async_v2 [post]
// 已移除：不再提供 /v1/t2a_async_v2 别名文档
// func docMiniMaxAsyncSpeechCreateAlias(*gin.Context) {}

// docMiniMaxAsyncSpeechQueryAlias 用于 OpenAI 兼容前缀文档描述。
// @Summary      minimaxi 官方接口 - 查询异步文本转语音任务（别名）
// @Tags         MiniMax Speech
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        task_id query string true "任务 ID"
// @Success      200 {object} types.MiniMaxAsyncSpeechQueryResponse
// @Router       /v1/query/t2a_async_query_v2 [get]
// 已移除：不再提供 /v1/query/t2a_async_query_v2 别名文档
// func docMiniMaxAsyncSpeechQueryAlias(*gin.Context) {}

// docMiniMaxVideoCreate 仅用于 Swagger 文档描述。
// @Summary      minimaxi 官方接口 - 创建视频任务
// @Description  与 minimaxi `/v1/video_generation` 一致，根据请求内容自动识别文生、图生、主体参考等模式。
// @Tags         MiniMax Video
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body miniapi.MiniMaxVideoCreateRequest true "视频任务请求体"
// @Success      200 {object} miniapi.MiniMaxVideoCreateResponse
// @Router       /minimaxi/v1/video_generation [post]
func docMiniMaxVideoCreate(*gin.Context) {}

// docMiniMaxVideoQuery 仅用于 Swagger 文档描述。
// @Summary      minimaxi 官方接口 - 查询视频任务
// @Description  查询任务状态，支持 queueing/processing/success/failed。
// @Tags         MiniMax Video
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        task_id query string true "任务 ID"
// @Success      200 {object} miniapi.MiniMaxVideoQueryResponse
// @Router       /minimaxi/v1/query/video_generation [get]
func docMiniMaxVideoQuery(*gin.Context) {}

// docMiniMaxFileRetrieve 仅用于 Swagger 文档描述。
// @Summary      minimaxi 官方接口 - 获取文件下载信息
// @Description  根据 file_id 获取下载链接，可选指定 channel_id/task_id。
// @Tags         MiniMax File
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        file_id    query string true  "文件 ID"
// @Param        channel_id query int    false "指定渠道 ID"
// @Param        task_id    query string false "关联任务 ID"
// @Success      200 {object} miniapi.MiniMaxFileRetrieveResponse
// @Router       /minimaxi/v1/files/retrieve [get]
func docMiniMaxFileRetrieve(*gin.Context) {}

// docMiniMaxFileRetrieveContent 仅用于 Swagger 文档描述。
// @Summary      minimaxi 官方接口 - 下载文件内容
// @Description  根据 file_id 下载生成的音视频资源内容。
// @Tags         MiniMax File
// @Security     BearerAuth
// @Produce      application/octet-stream
// @Param        file_id    query string true  "文件 ID"
// @Param        channel_id query int    false "指定渠道 ID"
// @Param        task_id    query string false "关联任务 ID"
// @Success      200 {file} binary "音视频文件流"
// @Router       /minimaxi/v1/files/retrieve_content [get]
func docMiniMaxFileRetrieveContent(*gin.Context) {}

// docMiniMaxVideoCreateAlias 仅用于 Swagger 文档描述。
// @Summary      minimaxi 官方接口 - 创建视频任务（别名）
// @Tags         MiniMax Video
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body miniapi.MiniMaxVideoCreateRequest true "视频任务请求体"
// @Success      200 {object} miniapi.MiniMaxVideoCreateResponse
// @Router       /v1/video_generation [post]
func docMiniMaxVideoCreateAlias(*gin.Context) {}

// docMiniMaxVideoQueryAlias 仅用于 Swagger 文档描述。
// @Summary      minimaxi 官方接口 - 查询视频任务（别名）
// @Tags         MiniMax Video
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        task_id query string true "任务 ID"
// @Success      200 {object} miniapi.MiniMaxVideoQueryResponse
// @Router       /v1/query/video_generation [get]
func docMiniMaxVideoQueryAlias(*gin.Context) {}
