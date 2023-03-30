package system

import (
	"github.com/fengzhicaizi/gin-vue3/app/model/system"
	. "github.com/fengzhicaizi/gin-vue3/global"
)

type RoleService struct{}

// @description: 获取所有角色
// @author: [zb](https://github.com/fengzhicaizi)
// @function: Login
// @param: auth *model.SysUser
// @return: token string, err error
func (*RoleService) GetRoleList() (roleData []*system.SysRole, errR error) {
	if errR = GVA_DB.Find(&roleData).Error; errR != nil {
		return
	}
	return
}

// @description: 查询角色
// @author: [zb](https://github.com/fengzhicaizi)
// @function: Login
// @param: auth *model.SysUser
// @return: token string, err error
func (*RoleService) GetRole(r *system.SysRole) (role system.SysRole, errR error) {
	if errR = GVA_DB.First(&role, r.Id).Error; errR != nil {
		return
	}
	return
}

// @description: 新增角色
// @author: [zb](https://github.com/fengzhicaizi)
// @function: Login
// @param: auth *model.SysUser
// @return: token string, err error
func (*RoleService) CreateRole(r *system.SysRole) (errR error) {
	if errR = GVA_DB.Create(&r).Error; errR != nil {
		return
	}
	return
}

// @description: 修改角色
// @author: [zb](https://github.com/fengzhicaizi)
// @function: Login
// @param: auth *model.SysUser
// @return: token string, err error
func (*RoleService) UpdateRole(r *system.SysRole) (errR error) {
	if errR = GVA_DB.Model(&system.SysRole{}).Updates(&r).Error; errR != nil {
		return
	}
	return
}

// @description: 删除角色
// @author: [zb](https://github.com/fengzhicaizi)
// @function: Login
// @param: auth *model.SysUser
// @return: token string, err error
func (*RoleService) DeleteRole(r *system.SysRole) (errR error) {
	if errR = GVA_DB.Delete(&r).Error; errR != nil {
		return
	}
	return
}
