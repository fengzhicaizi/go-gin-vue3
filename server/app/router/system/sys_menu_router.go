package system

import (
	v1 "github.com/fengzhicaizi/gin-vue3/app/api/v1"
	"github.com/fengzhicaizi/gin-vue3/middleware"

	"github.com/gin-gonic/gin"
)

type MenuRouter struct{}

func (MenuRouter) InitMenuRouter(Router *gin.RouterGroup) {
	menuApi := v1.ApiGroupApp.System.MenuApi
	menuRouter := Router.Group("menu")
	menuRouter.Use(middleware.JWT())
	{
		menuRouter.GET("getMenu", menuApi.GetMenuList)          // 获取所有菜单
		menuRouter.POST("createMenu", menuApi.CreateMenu)       // 新增菜单
		menuRouter.PUT("updateMenu", menuApi.UpdateMenu)        // 修改菜单
		menuRouter.DELETE("deleteMenu/:id", menuApi.DeleteMenu) // 删除菜单
	}
}
