package service

import (
	"gin-server/application/admin/aggregation"
	"gin-server/application/admin/in_out"
	"gin-server/application/dao"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/util/gconv"
)

var Tag *tagService

type tagService struct {
}

// 标签列表
func (a *tagService) TagList(ctx *gin.Context, req *in_out.PageReq) ([]in_out.TagInfoResp, error) {
	list, err := dao.Tag.TagList(nil, ctx, req.Page)
	if err != nil {
		return nil, err
	}
	var resp []in_out.TagInfoResp
	if err := gconv.Structs(list, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// 添加标签
func (a *tagService) TagAdd(ctx *gin.Context, req *in_out.AddTagReq) error {
	tagAddEnt := aggregation.Tag.NewAddTagAggregation(req)
	err := dao.Tag.AddTag(nil, ctx, tagAddEnt)
	if err != nil {
		return err
	}
	return nil
}

// 标签详情
func (a *tagService) TagDetail(ctx *gin.Context, req *in_out.IDReq) (*in_out.TagInfoResp, error) {
	articleEnt, err := dao.Tag.TagInfo(nil, ctx, req.ID)
	if err != nil {
		return nil, err
	}
	var resp *in_out.TagInfoResp
	if err := gconv.Struct(articleEnt, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// 更新标签
func (a *tagService) TagUpdate(ctx *gin.Context, req *in_out.UpdateTagReq) error {
	tagAddEnt := aggregation.Tag.NewUpdateTagAggregation(req)
	err := dao.Tag.TagUpdate(nil, ctx, tagAddEnt)
	if err != nil {
		return err
	}
	return nil
}

// 删除标签
func (a *tagService) TagDel(ctx *gin.Context, req *in_out.IDReq) error {
	// 判断标签下面的文章数量
	tagEnt, err := dao.Tag.TagInfo(nil, ctx, req.ID)
	if err != nil {
		return err
	}
	if tagEnt.ArticleSum > 0 {
		return gerror.New("当前标签存在已发布文章")
	}
	err = dao.Tag.TagDel(nil, ctx, req.ID)
	if err != nil {
		return err
	}
	return nil
}
