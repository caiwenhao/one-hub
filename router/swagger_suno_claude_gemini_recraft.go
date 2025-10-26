package router

import "github.com/gin-gonic/gin"

// Suno endpoints

// @Summary      Suno submit task
// @Tags         Suno
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        action path string true "Action"
// @Param        request body map[string]interface{} true "Task request"
// @Success      200 {object} map[string]interface{}
// @Router       /suno/submit/{action} [post]
func docSunoSubmit(*gin.Context) {}

// @Summary      Suno fetch
// @Tags         Suno
// @Produce      json
// @Security     BearerAuth
// @Router       /suno/fetch [post]
func docSunoFetch(*gin.Context) {}

// @Summary      Suno fetch by id
// @Tags         Suno
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Task ID"
// @Router       /suno/fetch/{id} [get]
func docSunoFetchByID(*gin.Context) {}

// Claude endpoints

// @Summary      Claude messages
// @Tags         Claude
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body map[string]interface{} true "Claude message request"
// @Success      200 {object} map[string]interface{}
// @Router       /claude/v1/messages [post]
func docClaudeMessages(*gin.Context) {}

// @Summary      Claude models
// @Tags         Claude
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} map[string]interface{}
// @Router       /claude/v1/models [get]
func docClaudeModels(*gin.Context) {}

// Gemini endpoints

// @Summary      Gemini model actions (chat/images/video)
// @Tags         Gemini
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        version path string true "API Version (e.g. v1beta)"
// @Param        model   path string true "Full model action, e.g.: gemini-2.5-flash:generateContent | gemini-2.5-flash:streamGenerateContent?alt=sse | imagen-4.0-generate-001:predict | veo-3.1-generate-preview:predictLongRunning"
// @Param        request body map[string]interface{} true "Request body (generateContent/predict)"
// @Success      200 {object} map[string]interface{}
// @Router       /gemini/{version}/models/{model} [post]
func docGeminiGenerate(*gin.Context) {}

// @Summary      Gemini list models
// @Tags         Gemini
// @Produce      json
// @Security     BearerAuth
// @Param        version path string true "API Version"
// @Success      200 {object} map[string]interface{}
// @Router       /gemini/{version}/models [get]
func docGeminiListModels(*gin.Context) {}

// @Summary      Gemini operations polling (Veo long-running jobs)
// @Tags         Gemini
// @Produce      json
// @Security     BearerAuth
// @Param        version path string true "API Version (e.g. v1beta)"
// @Param        name    path string true "Operation resource name (operations/...) from predictLongRunning response"
// @Param        model   query string false "Model used for selecting channel (e.g. veo-3.1-generate-preview). Default: gemini-2.5-pro"
// @Param        duration query int false "Video duration seconds (optional override for billing if not present in response)"
// @Success      200 {object} map[string]interface{}
// @Router       /gemini/{version}/operations/{name} [get]
func docGeminiOperations(*gin.Context) {}

// Recraft endpoints

// @Summary      Recraft generate image
// @Tags         Recraft
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body map[string]interface{} true "Image generation request"
// @Success      200 {object} map[string]interface{}
// @Router       /recraftAI/v1/images/generations [post]
func docRecraftGen(*gin.Context) {}

// @Summary      Recraft vectorize
// @Tags         Recraft
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body map[string]interface{} true "Vectorize request"
// @Success      200 {object} map[string]interface{}
// @Router       /recraftAI/v1/images/vectorize [post]
func docRecraftVectorize(*gin.Context) {}

// @Summary      Recraft remove background
// @Tags         Recraft
// @Accept       multipart/form-data
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} map[string]interface{}
// @Router       /recraftAI/v1/images/removeBackground [post]
func docRecraftRemoveBg(*gin.Context) {}

// @Summary      Recraft clarity upscale
// @Tags         Recraft
// @Accept       multipart/form-data
// @Produce      json
// @Security     BearerAuth
// @Router       /recraftAI/v1/images/clarityUpscale [post]
func docRecraftClarityUpscale(*gin.Context) {}

// @Summary      Recraft generative upscale
// @Tags         Recraft
// @Accept       multipart/form-data
// @Produce      json
// @Security     BearerAuth
// @Router       /recraftAI/v1/images/generativeUpscale [post]
func docRecraftGenerativeUpscale(*gin.Context) {}

// @Summary      Recraft styles
// @Tags         Recraft
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Router       /recraftAI/v1/styles [post]
func docRecraftStyles(*gin.Context) {}
