package system

import (
	"net/http"

	"github.com/fengzhicaizi/gin-vue3/app/model/common/response"
	"github.com/fengzhicaizi/gin-vue3/app/model/system"
	"github.com/gin-gonic/gin"
)

type MenuApi struct{}

// @Summary   获取所有菜单
// @Tags      Authority
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority                                                true  "权限id, 权限名, 父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "创建角色,返回包括系统角色详情"
func (*MenuApi) GetMenuList(c *gin.Context) {
	var query struct {
		Id int `form:"id"`
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	if query.Id == 0 {
		query.Id = 1
	}

	data, err := menuService.GetMenu(&system.SysMenu{Id: query.Id})

	if err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	response.Response(&response.ResponseStruct{Code: http.StatusOK, Msg: "查询成功", Data: data}, c)
}

// @Summary   创建菜单
// @Tags      Authority
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority                                                true  "权限id, 权限名, 父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "创建角色,返回包括系统角色详情"
func (*MenuApi) CreateMenu(c *gin.Context) {
	var json struct {
		Pid  int    `json:"pid"`
		Name string `json:"name"`
		Code string `json:"code"`
		Path string `json:"path"`
		Sort int    `json:"sort"`
		Mark string `json:"mark"`
		Icon string `json:"icon"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	err := menuService.CreateMenu(&system.SysMenu{
		Pid:  json.Pid,
		Name: json.Name,
		Code: json.Code,
		Path: json.Path,
		Sort: json.Sort,
		Mark: json.Mark,
		Icon: json.Icon,
	})

	if err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	response.Response(&response.ResponseStruct{Code: http.StatusOK, Msg: "新增成功"}, c)
}

// @Summary   修改菜单
// @Tags      Authority
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority                                                true  "权限id, 权限名, 父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "创建角色,返回包括系统角色详情"
func (*MenuApi) UpdateMenu(c *gin.Context) {
	var json struct {
		Id   int    `json:"id" binding:"required"`
		Pid  int    `json:"pid"`
		Name string `json:"name"`
		Code string `json:"code"`
		Path string `json:"path"`
		Sort int    `json:"sort"`
		Mark string `json:"mark"`
		Icon string `json:"icon"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	err := menuService.UpdateMenu(&system.SysMenu{
		Id:   json.Id,
		Pid:  json.Pid,
		Name: json.Name,
		Code: json.Code,
		Path: json.Path,
		Sort: json.Sort,
		Mark: json.Mark,
		Icon: json.Icon,
	})

	if err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	response.Response(&response.ResponseStruct{Code: http.StatusOK, Msg: "修改成功"}, c)
}

// @Summary   删除菜单
// @Tags      Authority
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority                                                true  "权限id, 权限名, 父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "创建角色,返回包括系统角色详情"
func (*MenuApi) DeleteMenu(c *gin.Context) {
	var uri struct {
		Id int `uri:"id"`
	}

	if err := c.ShouldBindUri(&uri); err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	err := menuService.DeleteMenu(&system.SysMenu{
		Id: uri.Id,
	})

	if err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	response.Response(&response.ResponseStruct{Code: http.StatusOK, Msg: "删除成功"}, c)
}
