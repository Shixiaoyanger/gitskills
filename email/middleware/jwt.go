package jwt

import (
	"email/pkg/info"
	"email/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//解析token判断是否是服务器签发的，是否过期
func JWT() gin.HandlerFunc {

	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = info.SUCCESS
		token := c.Request.Header.Get("token")
		if token == "" {
			code = info.INVALID_PARAMS
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = info.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = info.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != info.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  "待补充",
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
