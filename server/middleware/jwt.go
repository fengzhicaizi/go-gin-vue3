package middleware

import (
	"net/http"

	"github.com/fengzhicaizi/gin-vue3/app/model/common/response"
	"github.com/fengzhicaizi/gin-vue3/utils"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			code   = http.StatusOK
			errMsg string
		)
		token := c.Request.Header.Get("Authorization")

		if token == "" {
			code = http.StatusUnauthorized
			errMsg = `missing token`
		} else {
			if _, err := utils.ParseToken(token); err != nil {
				code = http.StatusUnauthorized
				errMsg = err.Error()
			}
		}

		if code != http.StatusOK {
			response.Response(&response.ResponseStruct{Code: code, Msg: errMsg, Data: ""}, c)
			c.Abort()
			return
		}

		c.Next()
	}
}
