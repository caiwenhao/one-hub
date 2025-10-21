package router

import (
	"one-api/common/config"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerDocs "one-api/docs/swagger"
)

func setSwaggerRouter(engine *gin.Engine) {
	if engine == nil {
		return
	}

	if !(config.Debug || viper.GetBool("swagger.enable")) {
		return
	}

	swaggerDocs.SwaggerInfo.BasePath = "/"

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
