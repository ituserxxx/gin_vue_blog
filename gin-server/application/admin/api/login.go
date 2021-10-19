package api

import (
	"gin-server/application/admin/in_out"
	"gin-server/application/admin/service"
	"gin-server/constant"
	"gin-server/library/response"
	"github.com/gin-gonic/gin"
)

var LoginApi *loginApi

type loginApi struct {
}

func (l *loginApi) AdminLogin(ctx *gin.Context) {
	var req *in_out.LoginReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.Err(ctx, constant.Error, err.Error())
	}
	resp, err := service.Login.AdminLogin(ctx, req)
	if err != nil {
		response.Err(ctx, constant.Handle_Error, err.Error())
	}
	response.Succ(ctx, resp)
}
