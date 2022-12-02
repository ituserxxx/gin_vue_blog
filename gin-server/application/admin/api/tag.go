package api

import (
	"gin-server/application/admin/in_out"
	"gin-server/application/admin/service"
	"gin-server/constant"
	"gin-server/library/response"
	"github.com/gin-gonic/gin"
)

var TagApi *tagApi

type tagApi struct {
}

// 标签列表
func (t *tagApi) TagList(ctx *gin.Context) {
	var req *in_out.PageReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.Err(ctx, constant.Error, err.Error())
	}
	list, err := service.Tag.TagList(ctx, req)
	if err != nil {
		response.Err(ctx, constant.Handle_Error, err.Error())
	}

	response.Succ(ctx, list)
}

// 新增标签
func (t *tagApi) AddTag(ctx *gin.Context) {
	var req *in_out.AddTagReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.Err(ctx, constant.Error, err.Error())
	}
	err := service.Tag.TagAdd(ctx, req)
	if err != nil {
		response.Err(ctx, constant.Handle_Error, err.Error())
	}
	response.Succ(ctx, "")
}

// 标签详情
func (t *tagApi) TagDetail(ctx *gin.Context) {
	var req *in_out.IDReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.Err(ctx, constant.Error, err.Error())
	}
	resp, err := service.Tag.TagDetail(ctx, req)
	if err != nil {
		response.Err(ctx, constant.Handle_Error, err.Error())
	}
	response.Succ(ctx, resp)
}

// 标签更新
func (t *tagApi) TagUpdate(ctx *gin.Context) {
	var req *in_out.UpdateTagReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.Err(ctx, constant.Error, err.Error())
	}
	err := service.Tag.TagUpdate(ctx, req)
	if err != nil {
		response.Err(ctx, constant.Handle_Error, err.Error())
	}
	response.Succ(ctx, "")
}

// 标签删除
func (t *tagApi) TagDel(ctx *gin.Context) {
	var req *in_out.IDReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.Err(ctx, constant.Error, err.Error())
	}
	err := service.Tag.TagDel(ctx, req)
	if err != nil {
		response.Err(ctx, constant.Handle_Error, err.Error())
	}
	response.Succ(ctx, "")
}
