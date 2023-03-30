package system

import (
	v1 "github.com/fengzhicaizi/gin-vue3/app/api/v1"
	"github.com/fengzhicaizi/gin-vue3/middleware"

	"github.com/gin-gonic/gin"
)

type DictRouter struct{}

func (DictRouter) InitDictRouter(Router *gin.RouterGroup) {
	dictApi := v1.ApiGroupApp.System.DictApi
	dictRouter := Router.Group("dict")
	dictRouter.Use(middleware.JWT())
	{
		dictRouter.GET("getDictList", dictApi.GetDictTypeList)          // 获取所有字典
		dictRouter.POST("createDict", dictApi.CreateDictType)           // 新增字典
		dictRouter.PUT("updateDict", dictApi.UpdateDictType)            // 修改字典
		dictRouter.DELETE("deleteDict/:id", dictApi.DeleteDictType)     // 删除字典
		dictRouter.GET("getDictData/:type", dictApi.GetDictData)        // 根据type查询字典
		dictRouter.POST("addDictData/:type", dictApi.AddDictData)       // 给type类型字典添加值
		dictRouter.PUT("updateDictData", dictApi.UpdateDictData)        // 给type类型字典修改值
		dictRouter.DELETE("deleteDictData/:id", dictApi.DeleteDictData) // 给type类型字典删除值
	}
}
