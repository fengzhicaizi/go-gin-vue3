package initialize

import (
	"github.com/fengzhicaizi/gin-vue3/app/router"
	"github.com/fengzhicaizi/gin-vue3/middleware"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.New()
	systemRouter := router.RouterGroupApp.System

	// exampleRouter := router.RouterGroupApp.Example

	r.Use(middleware.Cros()).Use(gin.Logger()).Use(gin.Recovery())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	// r.POST("/upload/images", common.UploadImages)

	PrivateGroup := r.Group("/v1/")

	{
		systemRouter.InitAuthRouter(PrivateGroup)
		systemRouter.InitDictRouter(PrivateGroup)
		systemRouter.InitRoleRouter(PrivateGroup)
		systemRouter.InitMenuRouter(PrivateGroup)
	}

	return r
}
