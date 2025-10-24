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

// @Summary      Gemini generate content
// @Tags         Gemini
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        version path string true "API Version"
// @Param        model   path string true "Model"
// @Param        request body map[string]interface{} true "Gemini request"
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

