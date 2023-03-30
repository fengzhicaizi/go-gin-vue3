package system

import (
	"errors"
	"strings"

	"github.com/fengzhicaizi/gin-vue3/app/model/common/request"
	"github.com/fengzhicaizi/gin-vue3/app/model/system"
	. "github.com/fengzhicaizi/gin-vue3/global"
	"github.com/fengzhicaizi/gin-vue3/utils"

	"github.com/jinzhu/gorm"
)

type AuthService struct{}
type GetAuthByTokenResponseStruct struct {
	Auth *system.SysAuth `json:"auth"`
	Role *system.SysRole `json:"role"`
	Menu *MenuTree       `json:"menu"`
}

// @description: 用户登录
// @author: [zb](https://github.com/fengzhicaizi)
// @function: Login
// @param: auth *model.SysUser
// @return: token string, err error
func (a *AuthService) Login(auth *system.SysAuth) (token string, errR error) {
	var user system.SysAuth
	var err error

	if err := GVA_DB.Where("username = ?", auth.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			errR = errors.New("不存在用户")
			return
		} else {
			errR = err
			return
		}
	}

	if strings.Compare(auth.Password, user.Password) != 0 {
		errR = errors.New("密码错误")
		return
	}

	token, err = utils.GenerateToken(&user)

	if err != nil {
		errR = err
		return
	}

	return
}

// @description: 查询所有用户
// @author: [zb](https://github.com/fengzhicaizi)
// @function: GetAuthList
// @param: auth *model.SysUser
// @return: authsR []*system.SysAuth, errR error
func (a *AuthService) GetAuthList(auth *system.SysAuth, page *request.PageRequestStruct) (authsR []system.SysAuth, len int, errR error) {
	if errR = GVA_DB.Offset((page.Page - 1) * page.PageSize).Limit(page.PageSize).Find(&authsR).Error; errR != nil {
		return
	}

	if errR = GVA_DB.Model(&system.SysAuth{}).Count(&len).Error; errR != nil {
		return
	}
	return
}

// @description: 根据id查询用户
// @author: [zb](https://github.com/fengzhicaizi)
// @function: GetAuthList
// @param: auth *model.SysUser
// @return: authsR []*system.SysAuth, errR error
func (a *AuthService) GetAuth(auth *system.SysAuth) (authR system.SysAuth, errR error) {

	if errR = GVA_DB.First(&authR, auth.Id).Error; errR != nil {
		return
	}

	return
}

// @description: 根据token查询用户
// @author: [zb](https://github.com/fengzhicaizi)
// @function: GetAuthByToken
// @param: auth *model.SysUser
// @return: authR *system.SysAuth, errR error
func (*AuthService) GetAuthByToken(token string) (authR GetAuthByTokenResponseStruct, errR error) {
	var (
		auth  system.SysAuth
		role  system.SysRole
		menus []system.SysMenu
	)
	a, err := utils.ParseToken(token)
	if err != nil {
		return
	}

	if errR = GVA_DB.First(&auth, a.Id).Error; errR != nil {
		return
	}

	if errR = GVA_DB.Find(&role, auth.RoleId).Error; errR != nil {
		return
	}

	if role.IsAllMenus {
		if errR = GVA_DB.Find(&menus).Error; errR != nil {
			return
		}
	} else {
		if errR = GVA_DB.Find(&menus, strings.Split(role.MenuIds, ",")).Error; errR != nil {
			return
		}
	}

	authR.Auth = &auth
	authR.Role = &role
	authR.Menu = RecurrenceMenuTree(1, menus)

	return
}

// @description: 添加用户信息
// @author: [zb](https://github.com/fengzhicaizi)
// @function: CreateAuth
// @param: auth *model.SysUser
// @return: b bool, errR error
func (a *AuthService) CreateAuth(auth *system.SysAuth) (errR error) {
	auth.Password = utils.DefaultPassword()

	if errR = GVA_DB.Create(&auth).Error; errR != nil {
		return
	}

	return
}

// @description: 修改用户信息
// @author: [zb](https://github.com/fengzhicaizi)
// @function: UpdateAuth
// @param: auth *model.SysUser
// @return: b bool, errR error
func (a *AuthService) UpdateAuth(auth *system.SysAuth) (errR error) {

	if errR = GVA_DB.Model(&system.SysAuth{}).Updates(&auth).Error; errR != nil {
		return
	}

	return
}

// @description: 删除用户信息
// @author: [zb](https://github.com/fengzhicaizi)
// @function: DeleteAuth
// @param: auth *model.SysUser
// @return: b bool, errR error
func (a *AuthService) DeleteAuth(auth *system.SysAuth) (errR error) {

	if errR = GVA_DB.Delete(&auth, auth.Id).Error; errR != nil {
		return
	}

	return
}

// @description: 重置密码
// @author: [zb](https://github.com/fengzhicaizi)
// @function: Login
// @param: auth *model.SysUser
// @return: b bool, errR error
func (a *AuthService) ResetPassword(auth *system.SysAuth) (errR error) {
	defaultPassword := utils.DefaultPassword()
	if errR = GVA_DB.Model(&auth).Update("Password", defaultPassword).Error; errR != nil {
		return
	}
	return
}
