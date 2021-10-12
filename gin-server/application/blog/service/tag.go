package service

import (
	"gin-server/application/blog/in_out"
	"gin-server/application/dao"
	"gin-server/application/entity"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
)

var TagApi *tagService

type tagService struct {
}

func (t *tagService) TagList(ctx *gin.Context, page int) ([]entity.Tag, error) {
	var resp []entity.Tag
	list, err := dao.Tag.TagList(nil, ctx, page)
	if err != nil {
		return resp, err
	}
	if err = gconv.Struct(list, &resp); err != nil {
		return resp, err
	}
	return list, nil
}

func (a *tagService) TagArticleList(ctx *gin.Context, req *in_out.TagArticleListReq) (in_out.HomeArticleListResp, error) {
	var resp in_out.HomeArticleListResp
	list, err := dao.TagArticle.GetBlogTagArticleList(nil, ctx, req.ID, req.Page)
	if err != nil {
		return resp, err
	}
	if err := gconv.Struct(list, &resp.ArticleList); err != nil {
		return resp, err
	}
	total, err := dao.TagArticle.GetBlogTagArticleSum(ctx, req.ID)
	if err != nil {
		return resp, err
	}
	resp.Total = total
	return resp, nil
}
