package service

import (
	"gin-server/application/blog/in_out"
	"gin-server/application/dao"
	"gin-server/utils"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
)

var Article *articleService

type articleService struct {
}

func (a *articleService) HomeArticleList(ctx *gin.Context, page int) (in_out.HomeArticleListResp, error) {
	var resp in_out.HomeArticleListResp
	list, err := dao.Article.GetBlogHomeArticleList(nil, ctx, page)
	if err != nil {
		return resp, err
	}

	go func() {
		for _, a := range list {
			utils.Meili.Insert(a)
		}
	}()
	if err := gconv.Struct(list, &resp.ArticleList); err != nil {
		return resp, err
	}
	total, err := dao.Article.BlogHomeArticleSum(ctx)
	if err != nil {
		return resp, err
	}
	resp.Total = total
	return resp, nil
}

func (a *articleService) ArticleDetail(ctx *gin.Context, id int) (*in_out.ArticleDetailResp, error) {
	var resp *in_out.ArticleDetailResp
	info, err := dao.Article.ArticleInfoById(nil, ctx, id)
	if err != nil {
		return nil, err
	}
	if err = gconv.Struct(info, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
func (a *articleService) SearchArticleList(ctx *gin.Context, content string) (*in_out.SearchArticleListResp, error) {
	var resp = &in_out.SearchArticleListResp{
		ArticleList: make([]interface{}, 0),
	}
	ls := utils.Meili.SearchByContent(content)
	if ls !=nil{
		resp.ArticleList =ls
	}
	resp.Total = len(resp.ArticleList)
	return resp, nil
}
