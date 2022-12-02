package api

import (
	"gin-server/application/admin/in_out"
	"gin-server/application/admin/service"
	"gin-server/constant"
	"gin-server/library/response"
	"github.com/gin-gonic/gin"
)

var VisitorApi *visitApi

type visitApi struct {
}

func (c *visitApi) VisitorList(ctx *gin.Context) {
	var req *in_out.PageReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.Err(ctx, constant.Error, err.Error())
	}
	list, err := service.Visitor.VisitorList(ctx, req)
	if err != nil {
		response.Err(ctx, constant.Handle_Error, err.Error())
	}
	response.Succ(ctx, list)
}
