package common

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/fengzhicaizi/gin-vue3/app/model/common/response"
	"github.com/gin-gonic/gin"
)

// @Summary   单文件上传
// @Tags      upload
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority                                                true  "权限id, 权限名, 父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "创建角色,返回包括系统角色详情"
func Upload(c *gin.Context) {
	var (
		files struct {
			Files []*multipart.FileHeader `form:"files" binding:"required"`
		}
		Data struct {
			Urls []string `json:"urls"`
		}
	)

	if err := c.ShouldBind(&files); err != nil {
		response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
		return
	}

	for _, fh := range files.Files {
		if err := c.SaveUploadedFile(fh, fmt.Sprintf("./runtime/upload/images/%v", fh.Filename)); err != nil {
			response.Response(&response.ResponseStruct{Code: http.StatusBadRequest, Msg: err.Error()}, c)
			return
		}
		Data.Urls = append(Data.Urls, fmt.Sprintf("./runtime/upload/images/%v", fh.Filename))
	}

	response.Response(&response.ResponseStruct{Code: http.StatusOK, Msg: "上传成功", Data: Data}, c)
}
