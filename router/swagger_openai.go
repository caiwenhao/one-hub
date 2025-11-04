package router

import "github.com/gin-gonic/gin"

// OpenAI-compatible endpoints (proxy). These are placeholders for Swagger docs only.

// @Summary      List models
// @Tags         OpenAI
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} map[string]interface{}
// @Router       /v1/models [get]
func docOpenAIListModels(*gin.Context) {}

// @Summary      Retrieve model
// @Tags         OpenAI
// @Produce      json
// @Security     BearerAuth
// @Param        model  path string true "Model ID"
// @Success      200 {object} map[string]interface{}
// @Router       /v1/models/{model} [get]
func docOpenAIRetrieveModel(*gin.Context) {}

// @Summary      Delete model (passthrough)
// @Tags         OpenAI
// @Produce      json
// @Security     BearerAuth
// @Param        model  path string true "Model ID"
// @Success      200 {object} map[string]interface{}
// @Router       /v1/models/{model} [delete]
func docOpenAIDeleteModel(*gin.Context) {}

// @Summary      Completions (legacy)
// @Tags         OpenAI
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body map[string]interface{} true "Completion request"
// @Success      200 {object} map[string]interface{}
// @Router       /v1/completions [post]
func docOpenAICompletions(*gin.Context) {}

// @Summary      Chat completions
// @Tags         OpenAI
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body map[string]interface{} true "Chat request"
// @Success      200 {object} map[string]interface{}
// @Router       /v1/chat/completions [post]
func docOpenAIChatCompletions(*gin.Context) {}

// @Summary      Responses API
// @Tags         OpenAI
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body map[string]interface{} true "Responses request"
// @Success      200 {object} map[string]interface{}
// @Router       /v1/responses [post]
func docOpenAIResponses(*gin.Context) {}

// @Summary      Create image
// @Tags         OpenAI Images
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body map[string]interface{} true "Image generation request"
// @Success      200 {object} map[string]interface{}
// @Router       /v1/images/generations [post]
func docOpenAIImagesGenerations(*gin.Context) {}

// @Summary      Edit image
// @Tags         OpenAI Images
// @Accept       multipart/form-data
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} map[string]interface{}
// @Router       /v1/images/edits [post]
func docOpenAIImagesEdits(*gin.Context) {}

// @Summary      Image variations
// @Tags         OpenAI Images
// @Accept       multipart/form-data
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} map[string]interface{}
// @Router       /v1/images/variations [post]
func docOpenAIImagesVariations(*gin.Context) {}

// @Summary      Embeddings
// @Tags         OpenAI
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body map[string]interface{} true "Embeddings request"
// @Success      200 {object} map[string]interface{}
// @Router       /v1/embeddings [post]
func docOpenAIEmbeddings(*gin.Context) {}

// @Summary      Audio transcription
// @Tags         OpenAI Audio
// @Accept       multipart/form-data
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} map[string]interface{}
// @Router       /v1/audio/transcriptions [post]
func docOpenAIAudioTranscriptions(*gin.Context) {}

// @Summary      Audio translation
// @Tags         OpenAI Audio
// @Accept       multipart/form-data
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} map[string]interface{}
// @Router       /v1/audio/translations [post]
func docOpenAIAudioTranslations(*gin.Context) {}

// @Summary      Text-to-speech
// @Tags         OpenAI Audio
// @Accept       json
// @Produce      audio/mpeg
// @Security     BearerAuth
// @Param        request body map[string]interface{} true "TTS request"
// @Success      200 {file} binary "Audio stream"
// @Router       /v1/audio/speech [post]
func docOpenAIAudioSpeech(*gin.Context) {}

// @Summary      Moderations
// @Tags         OpenAI
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body map[string]interface{} true "Moderations request"
// @Success      200 {object} map[string]interface{}
// @Router       /v1/moderations [post]
func docOpenAIModerations(*gin.Context) {}

// @Summary      Rerank
// @Tags         OpenAI
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body map[string]interface{} true "Rerank request"
// @Success      200 {object} map[string]interface{}
// @Router       /v1/rerank [post]
func docOpenAIRerank(*gin.Context) {}

// @Summary      Realtime (WebSocket)
// @Tags         OpenAI
// @Produce      json
// @Security     BearerAuth
// @Success      101 {string} string "Switching Protocols"
// @Router       /v1/realtime [get]
func docOpenAIRealtime(*gin.Context) {}

// @Summary      Create video
// @Tags         OpenAI Video
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body map[string]interface{} true "Video create request"
// @Success      200 {object} map[string]interface{}
// @Router       /v1/videos [post]
func docOpenAIVideoCreate(*gin.Context) {}

// @Summary      Retrieve video
// @Tags         OpenAI Video
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Video ID"
// @Success      200 {object} map[string]interface{}
// @Router       /v1/videos/{id} [get]
func docOpenAIVideoRetrieve(*gin.Context) {}

// @Summary      Download video content
// @Tags         OpenAI Video
// @Produce      application/octet-stream
// @Security     BearerAuth
// @Param        id path string true "Video ID"
// @Success      200 {file} binary "Video content"
// @Router       /v1/videos/{id}/content [get]
func docOpenAIVideoDownload(*gin.Context) {}

// @Summary      List videos
// @Tags         OpenAI Video
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} map[string]interface{}
// @Router       /v1/videos [get]
func docOpenAIVideoList(*gin.Context) {}

// @Summary      Remix video
// @Tags         OpenAI Video
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Video ID"
// @Param        request body map[string]interface{} true "Remix request"
// @Success      200 {object} map[string]interface{}
// @Router       /v1/videos/{id}/remix [post]
func docOpenAIVideoRemix(*gin.Context) {}

// @Summary      Delete video
// @Tags         OpenAI Video
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Video ID"
// @Success      200 {object} map[string]interface{}
// @Router       /v1/videos/{id} [delete]
func docOpenAIVideoDelete(*gin.Context) {}

// --- NewAPI 扩展 ---
// @Summary      Create vendor video (NewAPI)
// @Description  供应商风格：POST /v1/videos/generations（仅 NewAPI 渠道可用）。请求体字段：model/prompt/duration/aspect_ratio/image_urls/watermark；返回体为上游 JSON（{code,data}）。
// @Tags         NewAPI Video
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body map[string]interface{} true "Vendor generations request"
// @Success      200 {object} map[string]interface{}
// @Router       /v1/videos/generations [post]
func docNewAPIVideoGenerations(*gin.Context) {}

// @Summary      Retrieve vendor task (NewAPI)
// @Description  供应商风格：GET /v1/tasks/{id}（仅 NewAPI 渠道可用）。返回上游 JSON（包含 code/data 或统一结构）。
// @Tags         NewAPI Task
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Task ID"
// @Success      200 {object} map[string]interface{}
// @Router       /v1/tasks/{id} [get]
func docNewAPITaskRetrieve(*gin.Context) {}

// @Summary      Files proxy (passthrough)
// @Description  Passthrough to upstream API. Additional subpaths are supported.
// @Tags         OpenAI Passthrough
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} map[string]interface{}
// @Router       /v1/files [get]
func docOpenAIFilesProxy(*gin.Context) {}

// @Summary      Uploads proxy (passthrough)
// @Description  Passthrough to upstream API. Additional subpaths are supported.
// @Tags         OpenAI Passthrough
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} map[string]interface{}
// @Router       /v1/uploads [get]
func docOpenAIUploadsProxy(*gin.Context) {}

// @Summary      Conversations proxy (passthrough)
// @Description  Passthrough to upstream API. Additional subpaths are supported.
// @Tags         OpenAI Passthrough
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} map[string]interface{}
// @Router       /v1/conversations [get]
func docOpenAIConversationsProxy(*gin.Context) {}

// @Summary      Assistants proxy (passthrough)
// @Description  Passthrough to upstream API. Additional subpaths are supported.
// @Tags         OpenAI Passthrough
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} map[string]interface{}
// @Router       /v1/assistants [get]
func docOpenAIAssistantsProxy(*gin.Context) {}

// @Summary      Threads proxy (passthrough)
// @Description  Passthrough to upstream API. Additional subpaths are supported.
// @Tags         OpenAI Passthrough
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} map[string]interface{}
// @Router       /v1/threads [get]
func docOpenAIThreadsProxy(*gin.Context) {}

// @Summary      Batches proxy (passthrough)
// @Description  Passthrough to upstream API. Additional subpaths are supported.
// @Tags         OpenAI Passthrough
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} map[string]interface{}
// @Router       /v1/batches [get]
func docOpenAIBatchesProxy(*gin.Context) {}

// @Summary      Vector stores proxy (passthrough)
// @Description  Passthrough to upstream API. Additional subpaths are supported.
// @Tags         OpenAI Passthrough
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} map[string]interface{}
// @Router       /v1/vector_stores [get]
func docOpenAIVectorStoresProxy(*gin.Context) {}
