package system

import (
	"errors"

	"github.com/fengzhicaizi/gin-vue3/app/model/system"
	. "github.com/fengzhicaizi/gin-vue3/global"

	"github.com/jinzhu/gorm"
)

type DictTypeService struct{}

// @description: 获取所有字典类型
// @author: [zb](https://github.com/fengzhicaizi)
// @function: Login
// @param: auth *model.SysUser
// @return: token string, err error
func (*DictTypeService) GetDictTypeList(d *system.SysDictType) (dictsR []system.SysDictType, errR error) {
	if errR = GVA_DB.Find(&dictsR).Error; errR != nil {
		return
	}
	return
}

// @description: 新增字典
// @author: [zb](https://github.com/fengzhicaizi)
// @function: Login
// @param: auth *model.SysUser
// @return: token string, err error
func (*DictTypeService) CreateDictType(d *system.SysDictType) (errR error) {
	if err := GVA_DB.Where("type = ?", d.Type).First(&system.SysDictType{}).Error; err != gorm.ErrRecordNotFound {
		errR = errors.New("存在相同的type")
		return
	}
	if errR = GVA_DB.Create(&d).Error; errR != nil {
		return
	}
	return
}

// @description: 修改
// @author: [zb](https://github.com/fengzhicaizi)
// @function: Login
// @param: auth *model.SysUser
// @return: token string, err error
func (*DictTypeService) UpdateDictType(d *system.SysDictType) (errR error) {
	// 不能修改type值
	d.Type = ""

	if errR = GVA_DB.Model(&system.SysDictType{}).Updates(&d).Error; errR != nil {
		return
	}
	return
}

// @description: 删除
// @author: [zb](https://github.com/fengzhicaizi)
// @function: Login
// @param: auth *model.SysUser
// @return: token string, err error
func (*DictTypeService) DeleteDictType(d *system.SysDictType) (errR error) {
	if errR = GVA_DB.Delete(&d).Error; errR != nil {
		return
	}
	return
}
