package middleware

import (
	"fmt"
	"gin-server/constant"
	"gin-server/library/response"
	"gin-server/utils"
	"github.com/gin-gonic/gin"
)

// JWTAuth 中间件，检查token
func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Jwt")
		fmt.Printf("request token is %#v", c.Request.Header.Get("Jwt"))
		if token == "" {
			response.Err(c, constant.TokenFail, "请求未携带token，无权限访问")
			c.Abort()
			return
		}
		j := utils.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				response.Err(c, constant.TokenFail, "授权已过期")
				c.Abort()
				return
			}
			response.Err(c, constant.TokenFail, err.Error())
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("uid", claims.UID)
		c.Next()
	}
}
