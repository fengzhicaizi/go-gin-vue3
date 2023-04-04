package initialize

import (
	"net/http"

	"github.com/fengzhicaizi/gin-vue3/app/router"
	"github.com/fengzhicaizi/gin-vue3/app/router/common"
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
	r.StaticFS("/upload/images", http.Dir("./"))

	PrivateGroup := r.Group("/v1/")

	{
		systemRouter.InitAuthRouter(PrivateGroup)
		systemRouter.InitDictRouter(PrivateGroup)
		systemRouter.InitRoleRouter(PrivateGroup)
		systemRouter.InitMenuRouter(PrivateGroup)
	}

	CommonGroup := r.Group("/common/")

	{
		common.InitUploadRouter(CommonGroup)
	}

	return r
}
