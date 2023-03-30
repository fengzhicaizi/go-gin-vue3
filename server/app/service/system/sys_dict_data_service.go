package system

import (
	"github.com/fengzhicaizi/gin-vue3/app/model/system"
	. "github.com/fengzhicaizi/gin-vue3/global"
)

type DictDataService struct{}

// @description: 获取所有字典值
// @author: [zb](https://github.com/fengzhicaizi)
// @function: Login
// @param: auth *model.SysUser
// @return: token string, err error
func (*DictDataService) GetDictDataList(t string) (dictData []*system.SysDictData, errR error) {
	var dict system.SysDictType

	if errR = GVA_DB.Where("type = ?", t).First(&dict).Error; errR != nil {
		return
	}

	if errR = GVA_DB.Where(&system.SysDictData{TypeId: dict.Id}).Order("sort").Find(&dictData).Error; errR != nil {
		return
	}
	return
}

// @description: 添加字典值
// @author: [zb](https://github.com/fengzhicaizi)
// @function: Login
// @param: auth *model.SysUser
// @return: token string, err error
func (*DictDataService) AddDictData(t string, d *system.SysDictData) (errR error) {
	var dict system.SysDictType

	if errR = GVA_DB.Model(&system.SysDictType{}).Where("type = ?", t).First(&dict).Error; errR != nil {
		return
	}

	d.TypeId = dict.Id

	if errR = GVA_DB.Create(d).Error; errR != nil {
		return
	}

	return
}

// @description: 修改字典值
// @author: [zb](https://github.com/fengzhicaizi)
// @function: Login
// @param: auth *model.SysUser
// @return: token string, err error
func (*DictDataService) UpdateDictData(d *system.SysDictData) (errR error) {
	if errR = GVA_DB.Model(&system.SysDictData{}).Updates(&d).Error; errR != nil {
		return
	}

	return
}

// @description: 删除字典值
// @author: [zb](https://github.com/fengzhicaizi)
// @function: Login
// @param: auth *model.SysUser
// @return: token string, err error
func (*DictDataService) DeleteDictData(d *system.SysDictData) (errR error) {
	if errR = GVA_DB.Delete(&d).Error; errR != nil {
		return
	}

	return
}
