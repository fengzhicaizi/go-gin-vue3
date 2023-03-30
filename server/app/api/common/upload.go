package common

// import (
// 	"mime/multipart"
// 	"net/http"
// 	"github.com/fengzhicaizi/gin-vue3/pkg/file"
// 	"github.com/fengzhicaizi/gin-vue3/pkg/upload"

// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// )

// // Login @tags 通用模块
// // @Summary 上传图片
// // @Produce  json
// // @Param username formData string true "username"
// // @Param password formData string true "password"
// // @Success 200 {object} app.Response
// // @Router /upload/images [post]
// func UploadImages(c *gin.Context) {
// 	type File struct {
// 		Filename string `json:"filename"` // 文件名
// 		Size     int    `json:"size"`     // 文件大小
// 		Path     string `json:"path"`     //文件路径
// 		Type     string `json:"type"`     //文件类型
// 	}
// 	var p struct {
// 		Files []*multipart.FileHeader `form:"files" binding:"required"`
// 	}
// 	var res []*File

// 	if err := c.ShouldBind(&p); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"detail": err.Error(),
// 		})
// 		return
// 	}

// 	for _, v := range p.Files {
// 		filename := uuid.New().String() + file.GetExt(v.Filename)
// 		err := c.SaveUploadedFile(v, upload.GetImageFullPath()+filename)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"detail": err.Error(),
// 			})
// 			return
// 		}
// 		res = append(res, &File{
// 			Filename: filename,
// 			Size:     int(v.Size),
// 			Path:     upload.GetImageFullUrl(filename),
// 			Type:     v.Header["Content-Type"][0],
// 		})
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"detail": res,
// 	})
// }
