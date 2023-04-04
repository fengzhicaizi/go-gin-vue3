package common

import (
	"github.com/fengzhicaizi/gin-vue3/app/api/common"
	"github.com/gin-gonic/gin"
)

func InitUploadRouter(Router *gin.RouterGroup) {
	uploadRouter := Router.Group("upload")

	{
		uploadRouter.POST("/file", common.Upload)
	}
}
