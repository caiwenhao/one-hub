package router

import (
	"one-api/middleware"
	"one-api/providers/kling"
	"one-api/relay"
	"one-api/relay/midjourney"
	"one-api/relay/task"
	"one-api/relay/task/suno"
	"one-api/relay/task/vidu"
	"one-api/relay/task/volcark"

	"github.com/gin-gonic/gin"
)

func SetRelayRouter(router *gin.Engine) {
	router.Use(middleware.CORS())
	// https://platform.openai.com/docs/api-reference/introduction
	setOpenAIRouter(router)
	setMJRouter(router)
	setSunoRouter(router)
	setClaudeRouter(router)
	setGeminiRouter(router)
	setRecraftRouter(router)
	setKlingRouter(router)
	setViduRouter(router)
	setVolcArkRouter(router)
}

func setOpenAIRouter(router *gin.Engine) {
	modelsRouter := router.Group("/v1/models")
	modelsRouter.Use(middleware.OpenaiAuth(), middleware.Distribute())
	{
		modelsRouter.GET("", relay.ListModelsByToken)
		modelsRouter.GET("/:model", relay.RetrieveModel)
	}
	relayV1Router := router.Group("/v1")
	relayV1Router.Use(middleware.RelayPanicRecover(), middleware.OpenaiAuth(), middleware.Distribute(), middleware.DynamicRedisRateLimiter())
	{
		relayV1Router.POST("/completions", relay.Relay)
		relayV1Router.POST("/chat/completions", relay.Relay)
		relayV1Router.POST("/responses", relay.Relay)
		// relayV1Router.POST("/edits", controller.Relay)
		relayV1Router.POST("/images/generations", relay.Relay)
		relayV1Router.POST("/images/edits", relay.Relay)
		relayV1Router.POST("/images/variations", relay.Relay)
		relayV1Router.POST("/embeddings", relay.Relay)
		// relayV1Router.POST("/engines/:model/embeddings", controller.RelayEmbeddings)
		relayV1Router.POST("/audio/transcriptions", relay.Relay)
		relayV1Router.POST("/audio/translations", relay.Relay)
		relayV1Router.POST("/audio/speech", relay.Relay)
		relayV1Router.POST("/moderations", relay.Relay)
		relayV1Router.POST("/rerank", relay.RelayRerank)
		relayV1Router.GET("/realtime", relay.ChatRealtime)

		relayV1Router.Use(middleware.SpecifiedChannel())
		{
			relayV1Router.Any("/files", relay.RelayOnly)
			relayV1Router.Any("/files/*any", relay.RelayOnly)
			relayV1Router.Any("/fine_tuning/*any", relay.RelayOnly)
			relayV1Router.Any("/assistants", relay.RelayOnly)
			relayV1Router.Any("/assistants/*any", relay.RelayOnly)
			relayV1Router.Any("/threads", relay.RelayOnly)
			relayV1Router.Any("/threads/*any", relay.RelayOnly)
			relayV1Router.Any("/batches/*any", relay.RelayOnly)
			relayV1Router.Any("/vector_stores/*any", relay.RelayOnly)
			relayV1Router.DELETE("/models/:model", relay.RelayOnly)
		}
	}
}

func setMJRouter(router *gin.Engine) {
	relayMjRouter := router.Group("/mj")
	registerMjRouterGroup(relayMjRouter)

	relayMjModeRouter := router.Group("/:mode/mj")
	registerMjRouterGroup(relayMjModeRouter)
}

// Author: Calcium-Ion
// GitHub: https://github.com/Calcium-Ion/new-api
// Path: router/relay-router.go
func registerMjRouterGroup(relayMjRouter *gin.RouterGroup) {
	relayMjRouter.GET("/image/:id", midjourney.RelayMidjourneyImage)
	relayMjRouter.Use(middleware.RelayMJPanicRecover(), middleware.MjAuth(), middleware.Distribute(), middleware.DynamicRedisRateLimiter())
	{
		relayMjRouter.POST("/submit/action", midjourney.RelayMidjourney)
		relayMjRouter.POST("/submit/shorten", midjourney.RelayMidjourney)
		relayMjRouter.POST("/submit/modal", midjourney.RelayMidjourney)
		relayMjRouter.POST("/submit/imagine", midjourney.RelayMidjourney)
		relayMjRouter.POST("/submit/change", midjourney.RelayMidjourney)
		relayMjRouter.POST("/submit/simple-change", midjourney.RelayMidjourney)
		relayMjRouter.POST("/submit/describe", midjourney.RelayMidjourney)
		relayMjRouter.POST("/submit/blend", midjourney.RelayMidjourney)
		relayMjRouter.POST("/notify", midjourney.RelayMidjourney)
		relayMjRouter.GET("/task/:id/fetch", midjourney.RelayMidjourney)
		relayMjRouter.GET("/task/:id/image-seed", midjourney.RelayMidjourney)
		relayMjRouter.POST("/task/list-by-condition", midjourney.RelayMidjourney)
		relayMjRouter.POST("/insight-face/swap", midjourney.RelayMidjourney)
		relayMjRouter.POST("/submit/upload-discord-images", midjourney.RelayMidjourney)
	}
}

func setSunoRouter(router *gin.Engine) {
	relaySunoRouter := router.Group("/suno")
	relaySunoRouter.Use(middleware.RelaySunoPanicRecover(), middleware.OpenaiAuth(), middleware.Distribute(), middleware.DynamicRedisRateLimiter())
	{
		relaySunoRouter.POST("/submit/:action", task.RelayTaskSubmit)
		relaySunoRouter.POST("/fetch", suno.GetFetch)
		relaySunoRouter.GET("/fetch/:id", suno.GetFetchByID)
	}
}

func setClaudeRouter(router *gin.Engine) {
	relayClaudeRouter := router.Group("/claude")
	relayV1Router := relayClaudeRouter.Group("/v1")
	relayV1Router.Use(middleware.APIEnabled("claude"), middleware.RelayCluadePanicRecover(), middleware.ClaudeAuth(), middleware.Distribute(), middleware.DynamicRedisRateLimiter())
	{
		relayV1Router.POST("/messages", relay.Relay)
		relayV1Router.GET("/models", relay.ListClaudeModelsByToken)
	}
}

func setGeminiRouter(router *gin.Engine) {
	relayGeminiRouter := router.Group("/gemini")
	relayGeminiRouter.Use(middleware.APIEnabled("gemini"), middleware.RelayGeminiPanicRecover(), middleware.GeminiAuth(), middleware.Distribute(), middleware.DynamicRedisRateLimiter())
	{
		relayGeminiRouter.POST("/:version/models/:model", relay.Relay)
		relayGeminiRouter.GET("/:version/models", relay.ListGeminiModelsByToken)
	}
}

func setRecraftRouter(router *gin.Engine) {
	relayRecraftRouter := router.Group("/recraftAI/v1")
	relayRecraftRouter.Use(middleware.RelayPanicRecover(), middleware.OpenaiAuth(), middleware.Distribute(), middleware.DynamicRedisRateLimiter())
	{
		relayRecraftRouter.POST("/images/generations", relay.Relay)
		relayRecraftRouter.POST("/images/vectorize", relay.RelayRecraftAI)
		relayRecraftRouter.POST("/images/removeBackground", relay.RelayRecraftAI)
		relayRecraftRouter.POST("/images/clarityUpscale", relay.RelayRecraftAI)
		relayRecraftRouter.POST("/images/generativeUpscale", relay.RelayRecraftAI)
		relayRecraftRouter.POST("/styles", relay.RelayRecraftAI)
	}
}

func setKlingRouter(router *gin.Engine) {
	// 新增官方API兼容路由
	setOfficialKlingRouter(router)
}

// setOfficialKlingRouter 设置完全兼容官方API的路由
func setOfficialKlingRouter(router *gin.Engine) {
	// 官方API兼容路由组 - 使用 /kling/v1 前缀
	officialKlingRouter := router.Group("/kling/v1")
	officialKlingRouter.Use(middleware.RelayKlingPanicRecover(), middleware.OpenaiAuth(), middleware.Distribute(), middleware.DynamicRedisRateLimiter())
	{
		// 文生视频
		officialKlingRouter.POST("/videos/text2video", kling.CreateOfficialText2Video)
		officialKlingRouter.GET("/videos/text2video/:id", kling.GetOfficialTask)
		officialKlingRouter.GET("/videos/text2video", kling.ListOfficialTasks)

		// 图生视频
		officialKlingRouter.POST("/videos/image2video", kling.CreateOfficialImage2Video)
		officialKlingRouter.GET("/videos/image2video/:id", kling.GetOfficialTask)
		officialKlingRouter.GET("/videos/image2video", kling.ListOfficialTasks)

		// 多图参考生视频
		officialKlingRouter.POST("/videos/multi-image2video", kling.CreateOfficialMultiImage2Video)
		officialKlingRouter.GET("/videos/multi-image2video/:id", kling.GetOfficialMultiImage2VideoTask)
		officialKlingRouter.GET("/videos/multi-image2video", kling.ListOfficialMultiImage2VideoTasks)

		// 图像生成
		officialKlingRouter.POST("/images/generations", kling.CreateOfficialImage)
		officialKlingRouter.GET("/images/generations/:id", kling.GetOfficialImageTask)
		officialKlingRouter.GET("/images/generations", kling.ListOfficialImageTasks)

		// 多图参考生图
		officialKlingRouter.POST("/images/multi-image2image", kling.CreateOfficialMultiImage2Image)
		officialKlingRouter.GET("/images/multi-image2image/:id", kling.GetOfficialMultiImage2ImageTask)
		officialKlingRouter.GET("/images/multi-image2image", kling.ListOfficialMultiImage2ImageTasks)

		// 多模态视频编辑 - 选区管理
		officialKlingRouter.POST("/videos/multi-elements/init-selection", kling.InitSelection)
		officialKlingRouter.POST("/videos/multi-elements/add-selection", kling.AddSelection)
		officialKlingRouter.POST("/videos/multi-elements/delete-selection", kling.DeleteSelection)
		officialKlingRouter.POST("/videos/multi-elements/clear-selection", kling.ClearSelection)
		officialKlingRouter.POST("/videos/multi-elements/preview-selection", kling.PreviewSelection)

		// 多模态视频编辑 - 任务管理
		officialKlingRouter.POST("/videos/multi-elements", kling.CreateMultiElementsTask)
		officialKlingRouter.GET("/videos/multi-elements/:id", kling.GetMultiElementsTask)
		officialKlingRouter.GET("/videos/multi-elements", kling.ListMultiElementsTasks)
	}
}

func setViduRouter(router *gin.Engine) {
	relayViduRouter := router.Group("/vidu")
	relayViduRouter.Use(middleware.RelayPanicRecover(), middleware.OpenaiAuth(), middleware.Distribute())

	// 查询接口
	relayViduRouter.GET("/ent/v2/task/:task_id", vidu.RelayTaskFetch)
	relayViduRouter.GET("/ent/v2/tasks", vidu.RelayTaskFetchs)
	relayViduRouter.GET("/ent/v2/tasks/:task_id/creations", vidu.RelayTaskFetch) // 官方查询接口

	// 取消任务接口
	relayViduRouter.POST("/ent/v2/tasks/:task_id/cancel", vidu.RelayTaskCancel)

	relayViduRouter.Use(middleware.DynamicRedisRateLimiter())
	{
		// 任务提交接口
		relayViduRouter.POST("/ent/v2/:action", task.RelayTaskSubmit)
	}
}

func setVolcArkRouter(router *gin.Engine) {
	relayVolcRouter := router.Group("/volcark")
	relayVolcRouter.Use(middleware.RelayPanicRecover(), middleware.OpenaiAuth(), middleware.Distribute())
	{
		relayVolcRouter.GET("/api/v3/contents/generations/tasks/:task_id", volcark.RelayTaskFetch)
		relayVolcRouter.GET("/api/v3/contents/generations/tasks", volcark.RelayTaskList)
		relayVolcRouter.DELETE("/api/v3/contents/generations/tasks/:task_id", volcark.RelayTaskCancel)
		relayVolcRouter.POST("/api/v3/images/generations", relay.Relay)
	}

	relayVolcRouter.Use(middleware.DynamicRedisRateLimiter())
	{
		relayVolcRouter.POST("/api/v3/contents/generations/tasks", task.RelayTaskSubmit)
	}
}
