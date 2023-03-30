package system

import "github.com/fengzhicaizi/gin-vue3/app/service"

type ApiGroup struct {
	AuthApi
	DictApi
	MenuApi
	RoleApi
}

var (
	authService     = service.ServiceGroupApp.System.AuthService
	dictTypeService = service.ServiceGroupApp.System.DictTypeService
	dictDataService = service.ServiceGroupApp.System.DictDataService
	menuService     = service.ServiceGroupApp.System.MenuService
	roleService     = service.ServiceGroupApp.System.RoleService
)
