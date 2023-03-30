package system

import (
	"fmt"
	"sort"

	"github.com/fengzhicaizi/gin-vue3/app/model/system"
	. "github.com/fengzhicaizi/gin-vue3/global"
)

type MenuService struct{}

type MenuTree struct {
	system.SysMenu
	Children []*MenuTree `json:"children"`
}

// @description: 查询所有菜单
// @author: [zb](https://github.com/fengzhicaizi)
// @function: Login
// @param: auth *model.SysUser
// @return: token string, err error
func (*MenuService) GetMenu(m *system.SysMenu) (menuTree *MenuTree, errR error) {
	var menulist []system.SysMenu

	if errR = GVA_DB.Find(&menulist).Error; errR != nil {
		return
	}

	for _, sm := range menulist {
		fmt.Printf("%v", sm)
	}

	menuTree = RecurrenceMenuTree(m.Id, menulist)

	return
}

// @description: 新增菜单
// @author: [zb](https://github.com/fengzhicaizi)
// @function: Login
// @param: auth *model.SysUser
// @return: token string, err error
func (*MenuService) CreateMenu(r *system.SysMenu) (errR error) {
	if errR = GVA_DB.Create(&r).Error; errR != nil {
		return
	}
	return
}

// @description: 修改菜单
// @author: [zb](https://github.com/fengzhicaizi)
// @function: Login
// @param: auth *model.SysUser
// @return: token string, err error
func (*MenuService) UpdateMenu(r *system.SysMenu) (errR error) {
	if errR = GVA_DB.Model(&system.SysMenu{}).Updates(&r).Error; errR != nil {
		return
	}
	return
}

// @description: 删除菜单
// @author: [zb](https://github.com/fengzhicaizi)
// @function: Login
// @param: auth *model.SysUser
// @return: token string, err error
func (*MenuService) DeleteMenu(r *system.SysMenu) (errR error) {
	if errR = GVA_DB.Delete(&r).Error; errR != nil {
		return
	}
	return
}

// 获取menu的树形结构
func RecurrenceMenuTree(i int, menulist []system.SysMenu) (m *MenuTree) {
	var (
		currentMenu system.SysMenu
		children    []*MenuTree
	)

	for _, sm := range menulist {
		if i == sm.Id {
			currentMenu = sm
		}
		if i == sm.Pid {
			children = append(children, RecurrenceMenuTree(sm.Id, menulist))
			sort.Slice(children, func(i, j int) bool {
				return children[i].Sort < children[j].Sort
			})
		}
	}

	m = &MenuTree{
		SysMenu:  currentMenu,
		Children: children,
	}

	return
}
