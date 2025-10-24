package router

import "github.com/gin-gonic/gin"

// Kling official-compatible endpoints

// @Summary      Text to video
// @Tags         Kling
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Router       /kling/v1/videos/text2video [post]
func docKlingText2Video(*gin.Context) {}

// @Summary      Get text2video task
// @Tags         Kling
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Task ID"
// @Router       /kling/v1/videos/text2video/{id} [get]
func docKlingGetText2Video(*gin.Context) {}

// @Summary      List text2video tasks
// @Tags         Kling
// @Produce      json
// @Security     BearerAuth
// @Router       /kling/v1/videos/text2video [get]
func docKlingListText2Video(*gin.Context) {}

// @Summary      Image to video
// @Tags         Kling
// @Accept       multipart/form-data
// @Produce      json
// @Security     BearerAuth
// @Router       /kling/v1/videos/image2video [post]
func docKlingImage2Video(*gin.Context) {}

// @Summary      Get image2video task
// @Tags         Kling
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Task ID"
// @Router       /kling/v1/videos/image2video/{id} [get]
func docKlingGetImage2Video(*gin.Context) {}

// @Summary      List image2video tasks
// @Tags         Kling
// @Produce      json
// @Security     BearerAuth
// @Router       /kling/v1/videos/image2video [get]
func docKlingListImage2Video(*gin.Context) {}

// @Summary      Multi-image to video
// @Tags         Kling
// @Accept       multipart/form-data
// @Produce      json
// @Security     BearerAuth
// @Router       /kling/v1/videos/multi-image2video [post]
func docKlingMultiImage2Video(*gin.Context) {}

// @Summary      Get multi-image2video task
// @Tags         Kling
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Task ID"
// @Router       /kling/v1/videos/multi-image2video/{id} [get]
func docKlingGetMultiImage2Video(*gin.Context) {}

// @Summary      List multi-image2video tasks
// @Tags         Kling
// @Produce      json
// @Security     BearerAuth
// @Router       /kling/v1/videos/multi-image2video [get]
func docKlingListMultiImage2Video(*gin.Context) {}

// @Summary      Create image
// @Tags         Kling Image
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Router       /kling/v1/images/generations [post]
func docKlingCreateImage(*gin.Context) {}

// @Summary      Get image task
// @Tags         Kling Image
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Task ID"
// @Router       /kling/v1/images/generations/{id} [get]
func docKlingGetImageTask(*gin.Context) {}

// @Summary      List image tasks
// @Tags         Kling Image
// @Produce      json
// @Security     BearerAuth
// @Router       /kling/v1/images/generations [get]
func docKlingListImageTasks(*gin.Context) {}

// @Summary      Multi-image to image
// @Tags         Kling Image
// @Accept       multipart/form-data
// @Produce      json
// @Security     BearerAuth
// @Router       /kling/v1/images/multi-image2image [post]
func docKlingMultiImage2Image(*gin.Context) {}

// @Summary      Get multi-image2image task
// @Tags         Kling Image
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Task ID"
// @Router       /kling/v1/images/multi-image2image/{id} [get]
func docKlingGetMultiImage2Image(*gin.Context) {}

// @Summary      List multi-image2image tasks
// @Tags         Kling Image
// @Produce      json
// @Security     BearerAuth
// @Router       /kling/v1/images/multi-image2image [get]
func docKlingListMultiImage2Image(*gin.Context) {}

// @Summary      Init selection
// @Tags         Kling Edit
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Router       /kling/v1/videos/multi-elements/init-selection [post]
func docKlingInitSelection(*gin.Context) {}

// @Summary      Add selection
// @Tags         Kling Edit
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Router       /kling/v1/videos/multi-elements/add-selection [post]
func docKlingAddSelection(*gin.Context) {}

// @Summary      Delete selection
// @Tags         Kling Edit
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Router       /kling/v1/videos/multi-elements/delete-selection [post]
func docKlingDeleteSelection(*gin.Context) {}

// @Summary      Clear selection
// @Tags         Kling Edit
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Router       /kling/v1/videos/multi-elements/clear-selection [post]
func docKlingClearSelection(*gin.Context) {}

// @Summary      Preview selection
// @Tags         Kling Edit
// @Accept       json
// @Produce      application/octet-stream
// @Security     BearerAuth
// @Router       /kling/v1/videos/multi-elements/preview-selection [post]
func docKlingPreviewSelection(*gin.Context) {}

// @Summary      Create multi-elements task
// @Tags         Kling Edit
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Router       /kling/v1/videos/multi-elements [post]
func docKlingCreateMultiElementsTask(*gin.Context) {}

// @Summary      Get multi-elements task
// @Tags         Kling Edit
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Task ID"
// @Router       /kling/v1/videos/multi-elements/{id} [get]
func docKlingGetMultiElementsTask(*gin.Context) {}

// @Summary      List multi-elements tasks
// @Tags         Kling Edit
// @Produce      json
// @Security     BearerAuth
// @Router       /kling/v1/videos/multi-elements [get]
func docKlingListMultiElementsTasks(*gin.Context) {}

