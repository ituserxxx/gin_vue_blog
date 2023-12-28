package api

import (
	"gin-server/application/blog/in_out"
	"gin-server/application/blog/service"
	"gin-server/constant"
	"gin-server/library/response"
	"github.com/gin-gonic/gin"
)

var ArticleApi *articleApi

type articleApi struct {
}

func (a *articleApi) HomeArticleList(ctx *gin.Context) {
	var req *in_out.PageReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.Err(ctx, constant.Error, err.Error())
	}
	list, err := service.Article.HomeArticleList(ctx, req.Page)
	if err != nil {
		response.Err(ctx, constant.Handle_Error, err.Error())
	}
	response.Succ(ctx, list)
}

func (a *articleApi) ArticleDetail(ctx *gin.Context) {
	var req *in_out.IDReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.Err(ctx, constant.Error, err.Error())
	}
	list, err := service.Article.ArticleDetail(ctx, req.ID)
	if err != nil {
		response.Err(ctx, constant.Handle_Error, err.Error())
	}
	response.Succ(ctx, list)
}

func (a *articleApi) SearchArticleList(ctx *gin.Context) {
	var req *in_out.SearchArticleListReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.Err(ctx, constant.Error, err.Error())
	}
	list, err := service.Article.SearchArticleList(ctx, req.Content)
	if err != nil {
		response.Err(ctx, constant.Handle_Error, err.Error())
	}
	response.Succ(ctx, list)
}
