package middleware

import (
	"fmt"
	"gin-server/application/dao"
	"gin-server/application/entity"
	"gin-server/utils"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/os/gtime"
)

func Visitor() gin.HandlerFunc {
	return func(context *gin.Context) {
		go (func() {
			ip, Address, err := utils.GetIpAddress(context.ClientIP())
			if err != nil {
				fmt.Printf("%#v", err.Error())
			}
			visitor := &entity.Visitor{
				Ua:        context.Request.UserAgent(),
				IP:        ip,
				Uri:       context.Request.RequestURI,
				Referer:   context.Request.Referer(),
				VisitTime: gtime.Now(),
				IpAddress: Address,
			}
			_ = dao.Visitor.AddRecord(nil, context, visitor)
		})()
		context.Next()
	}
}
