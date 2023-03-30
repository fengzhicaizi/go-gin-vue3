package system

import (
	v1 "github.com/fengzhicaizi/gin-vue3/app/api/v1"
	"github.com/fengzhicaizi/gin-vue3/middleware"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct{}

func (a AuthRouter) InitAuthRouter(Router *gin.RouterGroup) {
	authApi := v1.ApiGroupApp.System.AuthApi
	noJwtAuthRouter := Router.Group("auth")
	authRouter := Router.Group("auth")
	authRouter.Use(middleware.JWT())
	{
		noJwtAuthRouter.POST("login", authApi.Login) // 登录
	}
	{
		authRouter.GET("getAuthList", authApi.GetAuthList)           // 获取所有用户列表
		authRouter.GET("getAuth/:id", authApi.GetAuth)               // 获取某个用户信息
		authRouter.GET("getAuthByToken", authApi.GetAuthByToken)     // 通过token获取用户信息
		authRouter.POST("createAuth", authApi.CreateAuth)            // 创建用户
		authRouter.DELETE("deleteAuth/:id", authApi.DeleteAuth)      // 删除用户
		authRouter.PUT("updateAuth", authApi.UpdateAuth)             // 更新用户
		authRouter.PATCH("resetPassword/:id", authApi.ResetPassword) // 重置密码
	}
}
