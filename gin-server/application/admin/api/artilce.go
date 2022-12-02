package api

import (
	"gin-server/application/admin/in_out"
	"gin-server/application/admin/service"
	"gin-server/application/dao"
	"gin-server/constant"
	"gin-server/library/response"
	"github.com/gin-gonic/gin"
)

var ArticleApi *articleApi

type articleApi struct {
}

// 文章列表
func (a *articleApi) ArticleList(ctx *gin.Context) {
	var req *in_out.PageReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.Err(ctx, constant.Error, err.Error())
	}
	list, err := service.Article.ArticleList(ctx, req)
	if err != nil {
		response.Err(ctx, constant.Handle_Error, err.Error())
	}
	response.Succ(ctx, list)
}

// 新增文章
func (a *articleApi) AddArticle(ctx *gin.Context) {
	var req *in_out.AddArticleReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.Err(ctx, constant.Error, err.Error())
	}
	err := service.Article.AddArticle(ctx, req)
	if err != nil {
		response.Err(ctx, constant.Handle_Error, err.Error())
	}
	response.Succ(ctx, "")
}

// 文章详情
func (a *articleApi) ArticleDetail(ctx *gin.Context) {
	var req *in_out.ArticleIDReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.Err(ctx, constant.Error, err.Error())
	}
	resp, err := service.Article.ArticleDetail(ctx, req)
	if err != nil {
		response.Err(ctx, constant.Handle_Error, err.Error())
	}
	response.Succ(ctx, resp)
}

// 文章更新
func (a *articleApi) ArticleUpdate(ctx *gin.Context) {
	var req *in_out.UpdateArticleReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.Err(ctx, constant.Error, err.Error())
	}
	err := service.Article.ArticleUpdate(ctx, req)
	if err != nil {
		response.Err(ctx, constant.Handle_Error, err.Error())
	}
	response.Succ(ctx, "")
}

// 文章存入草稿
func (a *articleApi) ArticleSaveDraft(ctx *gin.Context) {
	var req *in_out.UpdateArticleStatusReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.Err(ctx, constant.Error, err.Error())
	}
	req.Status = dao.TableFieldStatus["draft"]
	err := service.Article.ArticleUpdateStatus(ctx, req)
	if err != nil {
		response.Err(ctx, constant.Handle_Error, err.Error())
	}
	response.Succ(ctx, "")
}

// 草稿发布
func (a *articleApi) ArticlePublish(ctx *gin.Context) {
	var req *in_out.UpdateArticleStatusReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.Err(ctx, constant.Error, err.Error())
	}
	req.Status = dao.TableFieldStatus["publish"]
	err := service.Article.ArticleUpdateStatus(ctx, req)
	if err != nil {
		response.Err(ctx, constant.Handle_Error, err.Error())
	}
	response.Succ(ctx, "")
}

// 文章删除
func (a *articleApi) ArticleDelete(ctx *gin.Context) {
	var req *in_out.ArticleIDReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.Err(ctx, constant.Error, err.Error())
	}
	err := service.Article.ArticleDelete(ctx, req)
	if err != nil {
		response.Err(ctx, constant.Handle_Error, err.Error())
	}
	response.Succ(ctx, "")
}
