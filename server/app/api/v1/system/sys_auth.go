package system

import (
	"net/http"

	"github.com/fengzhicaizi/gin-vue3/app/model/common/request"
	"github.com/fengzhicaizi/gin-vue3/app/model/common/response"
	"github.com/fengzhicaizi/gin-vue3/app/model/system"

	"github.com/gin-gonic/gin"
)

type AuthApi struct{}

// @Summary   获取所有用户信息
// @Tags      Authority
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority                                                true  "权限id, 权限名, 父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "创建角色,返回包括系统角色详情"
func (a *AuthApi) GetAuthList(c *gin.Context) {
	var query struct {
		request.PageRequestStruct
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	data, len, err := authService.GetAuthList(&system.SysAuth{}, &query.PageRequestStruct)

	if err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	response.Response(&response.ResponseStruct{Code: http.StatusOK, Msg: "查询成功", Data: map[string]interface{}{
		"list": data,
		"page": response.PageResponseStruct{
			Page:     query.Page,
			PageSize: query.PageSize,
			Total:    len,
		},
	}}, c)
}

// @Summary   获取所有用户信息
// @Tags      Authority
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority                                                true  "权限id, 权限名, 父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "创建角色,返回包括系统角色详情"
func (a *AuthApi) GetAuth(c *gin.Context) {
	var uri struct {
		Id int `uri:"id"`
	}

	if err := c.ShouldBindUri(&uri); err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	data, err := authService.GetAuth(&system.SysAuth{Id: uri.Id})

	if err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	response.Response(&response.ResponseStruct{Code: http.StatusOK, Msg: "查询成功", Data: data}, c)
}

// 根据token获取用户信息
func (a *AuthApi) GetAuthByToken(c *gin.Context) {
	var header struct {
		Token string `header:"Authorization"`
	}

	if err := c.ShouldBindHeader(&header); err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusUnauthorized, Msg: err.Error()}, c)
		return
	}

	auth, err := authService.GetAuthByToken(header.Token)

	if err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusUnauthorized, Msg: err.Error()}, c)
		return
	}

	response.Response(&response.ResponseStruct{Code: http.StatusOK, Msg: "查询成功", Data: auth}, c)
}

// @Summary   新增用户
// @Tags      Authority
// @Path 			/v1/auth/createAuth
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority                                                true  "权限id, 权限名, 父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "创建角色,返回包括系统角色详情"
func (a *AuthApi) CreateAuth(c *gin.Context) {
	var data struct {
		Username string `json:"username" binding:"required"`
		RoleId   int    `json:"roleId" binding:"required"`
		RealName string `json:"realName"`
		Phone    string `json:"phone"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	err := authService.CreateAuth(&system.SysAuth{
		Username: data.Username,
		RealName: data.RealName,
		Phone:    data.Phone,
		RoleId:   data.RoleId,
	})

	if err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	response.Response(&response.ResponseStruct{Code: http.StatusOK, Msg: "添加成功"}, c)
}

// @Summary   删除用户
// @Tags      Authority
// @Path 			/v1/auth/deleteAuth/:id
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority                                                true  "权限id, 权限名, 父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "创建角色,返回包括系统角色详情"
func (a *AuthApi) DeleteAuth(c *gin.Context) {
	var uri struct {
		Id int `uri:"id"`
	}

	if err := c.ShouldBindUri(&uri); err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	err := authService.DeleteAuth(&system.SysAuth{Id: uri.Id})

	if err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	response.Response(&response.ResponseStruct{Code: http.StatusOK, Msg: "删除成功"}, c)
}

// 更新用户
func (a *AuthApi) UpdateAuth(c *gin.Context) {
	var data struct {
		Id       int    `json:"id" binding:"required"`
		Username string `json:"username"`
		Password string `json:"password"`
		Phone    string `json:"phone"`
		RoleId   int    `json:"roleId"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	err := authService.UpdateAuth(&system.SysAuth{
		Id:       data.Id,
		Username: data.Username,
		Password: data.Password,
		Phone:    data.Phone,
		RoleId:   data.RoleId,
	})

	if err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	response.Response(&response.ResponseStruct{Code: http.StatusOK, Msg: "修改成功"}, c)
}

// @Summary   登录
// @Tags      Authority
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority                                                true  "权限id, 权限名, 父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "创建角色,返回包括系统角色详情"
func (a *AuthApi) Login(c *gin.Context) {
	var data struct {
		Username string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusOK, Msg: err.Error()}, c)
		return
	}

	// 检测用户名密码是否正确
	token, err := authService.Login(&system.SysAuth{
		Username: data.Username,
		Password: data.Password,
	})

	if err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	response.Response(&response.ResponseStruct{Code: http.StatusOK, Msg: "登录成功", Data: token}, c)
}

// 重置密码
func (a *AuthApi) ResetPassword(c *gin.Context) {
	var uri struct {
		Id int `uri:"id"`
	}

	if err := c.ShouldBindUri(&uri); err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	err := authService.ResetPassword(&system.SysAuth{Id: uri.Id})

	if err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusUnauthorized, Msg: err.Error()}, c)
		return
	}

	response.Response(&response.ResponseStruct{Code: http.StatusOK, Msg: "重置成功"}, c)
}
