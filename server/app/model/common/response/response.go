package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// http返回模型
type ResponseStruct struct {
	Code int
	Data interface{}
	Msg  string
}

type PageResponseStruct struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
}

func Response(r *ResponseStruct, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": r.Code,
		"data": r.Data,
		"msg":  r.Msg,
	})
}
