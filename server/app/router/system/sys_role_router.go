package system

import (
	v1 "github.com/fengzhicaizi/gin-vue3/app/api/v1"
	"github.com/fengzhicaizi/gin-vue3/middleware"

	"github.com/gin-gonic/gin"
)

type RoleRouter struct{}

func (RoleRouter) InitRoleRouter(Router *gin.RouterGroup) {
	roleApi := v1.ApiGroupApp.System.RoleApi
	roleRouter := Router.Group("role")
	roleRouter.Use(middleware.JWT())
	{
		roleRouter.GET("getRoleList", roleApi.GetRoleList)      // 获取所有角色
		roleRouter.GET("getRole/:id", roleApi.GetRole)          // 获取所有角色
		roleRouter.POST("createRole", roleApi.CreateRole)       // 新增角色
		roleRouter.PUT("updateRole", roleApi.UpdateRole)        // 修改角色
		roleRouter.DELETE("deleteRole/:id", roleApi.DeleteRole) // 删除角色
	}
}
