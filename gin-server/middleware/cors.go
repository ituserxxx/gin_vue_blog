package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		//origin := c.Request.Header.Get("Origin")
		//var filterHost = [...]string{"http://127.0.0.1:6003","http://127.0.0.1:6004"}
		//// filterHost 做过滤器，防止不合法的域名访问
		//var isAccess = true
		//for _, v := range filterHost {
		//	match, _ := regexp.MatchString(v, origin)
		//	if match {
		//		isAccess = true
		//		break
		//	}
		//}
		//if isAccess {
		// 核心处理方式
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, X-Token,j_token")
		c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
		c.Set("content-type", "application/json")
		//}
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}
