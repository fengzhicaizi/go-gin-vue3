package system

import (
	"net/http"

	"github.com/fengzhicaizi/gin-vue3/app/model/common/response"
	"github.com/fengzhicaizi/gin-vue3/app/model/system"

	"github.com/gin-gonic/gin"
)

type DictApi struct{}

// @Summary 	获取所有字典
// @Tags      Authority
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority                                                true  "权限id, 权限名, 父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "创建角色,返回包括系统角色详情"
func (a *DictApi) GetDictTypeList(c *gin.Context) {

	data, err := dictTypeService.GetDictTypeList(&system.SysDictType{})

	if err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	response.Response(&response.ResponseStruct{Code: http.StatusOK, Msg: "查询成功", Data: data}, c)
}

// @Summary 	新增字典
// @Tags      Authority
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority                                                true  "权限id, 权限名, 父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "创建角色,返回包括系统角色详情"
func (a *DictApi) CreateDictType(c *gin.Context) {
	var data struct {
		Type   string `json:"type" binding:"required"`
		Name   string `json:"name" binding:"required"`
		Remark string `json:"remark"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	err := dictTypeService.CreateDictType(&system.SysDictType{
		Type:   data.Type,
		Name:   data.Name,
		Remark: data.Remark,
	})

	if err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	response.Response(&response.ResponseStruct{Code: http.StatusOK, Msg: "新增成功"}, c)
}

// @Summary  	修改字典
// @Tags      Authority
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority                                                true  "权限id, 权限名, 父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "创建角色,返回包括系统角色详情"
func (a *DictApi) UpdateDictType(c *gin.Context) {
	var data struct {
		Id     int    `json:"id" binding:"required"`
		Type   string `json:"type" binding:"required"`
		Name   string `json:"name" binding:"required"`
		Remark string `json:"remark"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	err := dictTypeService.UpdateDictType(&system.SysDictType{
		Id:     data.Id,
		Type:   data.Type,
		Name:   data.Name,
		Remark: data.Remark,
	})

	if err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	response.Response(&response.ResponseStruct{Code: http.StatusOK, Msg: "修改成功"}, c)
}

// @Summary  	删除字典
// @Tags      Authority
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority                                                true  "权限id, 权限名, 父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "创建角色,返回包括系统角色详情"
func (a *DictApi) DeleteDictType(c *gin.Context) {
	var uri struct {
		Id int `uri:"id" binding:"required"`
	}

	if err := c.ShouldBindUri(&uri); err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	err := dictTypeService.DeleteDictType(&system.SysDictType{Id: uri.Id})

	if err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	response.Response(&response.ResponseStruct{Code: http.StatusOK, Msg: "删除成功"}, c)
}

// @Summary  	查询字典值
// @Tags      Authority
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority                                                true  "权限id, 权限名, 父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "创建角色,返回包括系统角色详情"
func (a *DictApi) GetDictData(c *gin.Context) {
	var uri struct {
		Type string `uri:"type"`
	}

	if err := c.ShouldBindUri(&uri); err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	data, err := dictDataService.GetDictDataList(uri.Type)

	if err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	response.Response(&response.ResponseStruct{Code: http.StatusOK, Msg: "查询成功", Data: data}, c)
}

// @Summary  	新增字典值
// @Tags      Authority
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority                                                true  "权限id, 权限名, 父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "创建角色,返回包括系统角色详情"
func (a *DictApi) AddDictData(c *gin.Context) {
	var uri struct {
		Type string `uri:"type"`
	}

	var json struct {
		Label  string `json:"label"`
		Value  string `json:"value"`
		Remark string `json:"remark"`
		Sort   int    `json:"sort"`
	}

	if err := c.ShouldBindUri(&uri); err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	err := dictDataService.AddDictData(uri.Type, &system.SysDictData{
		Label:  json.Label,
		Value:  json.Value,
		Remark: json.Remark,
		Sort:   json.Sort,
	})

	if err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	response.Response(&response.ResponseStruct{Code: http.StatusOK, Msg: "新增成功"}, c)
}

// @Summary  	修改字典值
// @Tags      Authority
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority                                                true  "权限id, 权限名, 父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "创建角色,返回包括系统角色详情"
func (a *DictApi) UpdateDictData(c *gin.Context) {

	var json struct {
		Id     int    `json:"id" binding:"required"`
		Label  string `json:"label"`
		Value  string `json:"value"`
		Remark string `json:"remark"`
		Sort   int    `json:"sort"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	err := dictDataService.UpdateDictData(&system.SysDictData{
		Id:     json.Id,
		Label:  json.Label,
		Value:  json.Value,
		Remark: json.Remark,
		Sort:   json.Sort,
	})

	if err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	response.Response(&response.ResponseStruct{Code: http.StatusOK, Msg: "修改成功"}, c)
}

// @Summary  	删除字典值
// @Tags      Authority
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority                                                true  "权限id, 权限名, 父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "创建角色,返回包括系统角色详情"
func (a *DictApi) DeleteDictData(c *gin.Context) {
	var uri struct {
		Id int `uri:"id"`
	}

	if err := c.ShouldBindUri(&uri); err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	err := dictDataService.DeleteDictData(&system.SysDictData{Id: uri.Id})

	if err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	response.Response(&response.ResponseStruct{Code: http.StatusOK, Msg: "删除成功"}, c)
}
