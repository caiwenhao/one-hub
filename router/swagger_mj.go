package router

import "github.com/gin-gonic/gin"

// Midjourney-compatible endpoints. Placeholders for Swagger only.

// @Summary      Get task image
// @Tags         Midjourney
// @Produce      image/png
// @Security     BearerAuth
// @Param        id path string true "Image ID"
// @Success      200 {file} binary "Image content"
// @Router       /mj/image/{id} [get]
func docMjGetImage(*gin.Context) {}

// @Summary      Submit action
// @Tags         Midjourney
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} map[string]interface{}
// @Router       /mj/submit/action [post]
func docMjSubmitAction(*gin.Context) {}

// @Summary      Submit shorten
// @Tags         Midjourney
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Router       /mj/submit/shorten [post]
func docMjSubmitShorten(*gin.Context) {}

// @Summary      Submit modal
// @Tags         Midjourney
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Router       /mj/submit/modal [post]
func docMjSubmitModal(*gin.Context) {}

// @Summary      Submit imagine
// @Tags         Midjourney
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Router       /mj/submit/imagine [post]
func docMjSubmitImagine(*gin.Context) {}

// @Summary      Submit change
// @Tags         Midjourney
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Router       /mj/submit/change [post]
func docMjSubmitChange(*gin.Context) {}

// @Summary      Submit simple change
// @Tags         Midjourney
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Router       /mj/submit/simple-change [post]
func docMjSubmitSimpleChange(*gin.Context) {}

// @Summary      Submit describe
// @Tags         Midjourney
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Router       /mj/submit/describe [post]
func docMjSubmitDescribe(*gin.Context) {}

// @Summary      Submit blend
// @Tags         Midjourney
// @Accept       multipart/form-data
// @Produce      json
// @Security     BearerAuth
// @Router       /mj/submit/blend [post]
func docMjSubmitBlend(*gin.Context) {}

// @Summary      Notify callback
// @Tags         Midjourney
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Router       /mj/notify [post]
func docMjNotify(*gin.Context) {}

// @Summary      Fetch task
// @Tags         Midjourney
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Task ID"
// @Router       /mj/task/{id}/fetch [get]
func docMjTaskFetch(*gin.Context) {}

// @Summary      Get image seed
// @Tags         Midjourney
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Task ID"
// @Router       /mj/task/{id}/image-seed [get]
func docMjTaskImageSeed(*gin.Context) {}

// @Summary      List tasks by condition
// @Tags         Midjourney
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Router       /mj/task/list-by-condition [post]
func docMjTaskListByCondition(*gin.Context) {}

// @Summary      Insight face swap
// @Tags         Midjourney
// @Accept       multipart/form-data
// @Produce      json
// @Security     BearerAuth
// @Router       /mj/insight-face/swap [post]
func docMjInsightFaceSwap(*gin.Context) {}

// @Summary      Upload Discord images
// @Tags         Midjourney
// @Accept       multipart/form-data
// @Produce      json
// @Security     BearerAuth
// @Router       /mj/submit/upload-discord-images [post]
func docMjUploadDiscordImages(*gin.Context) {}

// Duplicate routes with mode prefix

// @Summary      Submit action (mode)
// @Tags         Midjourney
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        mode path string true "Mode"
// @Router       /{mode}/mj/submit/action [post]
func docMjModeSubmitAction(*gin.Context) {}

