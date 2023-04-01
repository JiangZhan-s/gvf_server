package middleware

import (
	"github.com/gin-gonic/gin"
	"gvf_server/models/res"
	"gvf_server/utils/jwts"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMessage("未携带token", c)
			c.Abort()
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token错误", c)
			c.Abort()
			return
		}
		// 登陆的用户
		c.Set("claims", claims)
	}
}
