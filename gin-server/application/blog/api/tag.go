package api

import (
	"gin-server/application/blog/in_out"
	"gin-server/application/blog/service"
	"gin-server/constant"
	"gin-server/library/response"
	"github.com/gin-gonic/gin"
)

var TagApi *tagApi

type tagApi struct {
}

func (a *tagApi) TagList(ctx *gin.Context) {
	var req *in_out.PageReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.Err(ctx, constant.Error, err.Error())
	}
	list, err := service.TagApi.TagList(ctx, req.Page)
	if err != nil {
		response.Err(ctx, constant.Handle_Error, err.Error())
	}
	response.Succ(ctx, list)
}

func (a *tagApi) TagArticleList(ctx *gin.Context) {
	var req *in_out.TagArticleListReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.Err(ctx, constant.Error, err.Error())
	}
	list, err := service.TagApi.TagArticleList(ctx, req)
	if err != nil {
		response.Err(ctx, constant.Handle_Error, err.Error())
	}
	response.Succ(ctx, list)
}
