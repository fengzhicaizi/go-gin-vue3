package system

import (
	"net/http"

	"github.com/fengzhicaizi/gin-vue3/app/model/common/response"
	"github.com/fengzhicaizi/gin-vue3/app/model/system"
	"github.com/gin-gonic/gin"
)

type RoleApi struct{}

// @Summary 	获取所有角色
// @Tags      Role
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority                                                true  "权限id, 权限名, 父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "创建角色,返回包括系统角色详情"
func (*RoleApi) GetRoleList(c *gin.Context) {

	data, err := roleService.GetRoleList()

	if err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	response.Response(&response.ResponseStruct{Code: http.StatusOK, Msg: "查询成功", Data: data}, c)
}

// @Summary 	查询角色
// @Tags      Role
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority                                                true  "权限id, 权限名, 父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "创建角色,返回包括系统角色详情"
func (*RoleApi) GetRole(c *gin.Context) {
	var uri struct {
		Id int `uri:"id"`
	}

	if err := c.ShouldBindUri(&uri); err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	data, err := roleService.GetRole(&system.SysRole{
		Id: uri.Id,
	})

	if err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	response.Response(&response.ResponseStruct{Code: http.StatusOK, Msg: "查询成功", Data: data}, c)
}

// @Summary 	创建角色
// @Tags      Role
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority                                                true  "权限id, 权限名, 父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "创建角色,返回包括系统角色详情"
func (*RoleApi) CreateRole(c *gin.Context) {
	var json struct {
		Name    string `json:"name" binding:"required"`
		MenuIds string `json:"menu_ids"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	err := roleService.CreateRole(&system.SysRole{
		Name:    json.Name,
		MenuIds: json.MenuIds,
	})

	if err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	response.Response(&response.ResponseStruct{Code: http.StatusOK, Msg: "创建成功"}, c)
}

// @Summary 	修改角色
// @Tags      Role
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority                                                true  "权限id, 权限名, 父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "创建角色,返回包括系统角色详情"
func (*RoleApi) UpdateRole(c *gin.Context) {
	var json struct {
		Id      int    `json:"id" binding:"required"`
		Name    string `json:"name"`
		MenuIds string `json:"menuIds"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	err := roleService.UpdateRole(&system.SysRole{
		Id:      json.Id,
		Name:    json.Name,
		MenuIds: json.MenuIds,
	})

	if err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	response.Response(&response.ResponseStruct{Code: http.StatusOK, Msg: "修改成功"}, c)
}

// @Summary 	删除角色
// @Tags      Role
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority                                                true  "权限id, 权限名, 父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "创建角色,返回包括系统角色详情"
func (*RoleApi) DeleteRole(c *gin.Context) {
	var uri struct {
		Id int `uri:"id"`
	}

	if err := c.ShouldBindUri(&uri); err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	err := roleService.DeleteRole(&system.SysRole{
		Id: uri.Id,
	})

	if err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	response.Response(&response.ResponseStruct{Code: http.StatusOK, Msg: "删除成功"}, c)
}
