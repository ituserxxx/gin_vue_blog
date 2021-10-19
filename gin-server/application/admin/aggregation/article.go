package aggregation

import (
	"gin-server/application/admin/in_out"
	"gin-server/application/entity"
	"github.com/gogf/gf/os/gtime"
)

var Article *articleAggregation

type articleAggregation struct {
	Ent     *entity.Article
	EntList []entity.Article
	TagList []entity.Tag
}

func (a *articleAggregation) NewAddArticleAggregation(req *in_out.AddArticleReq) *entity.Article {
	r := &articleAggregation{
		Ent: &entity.Article{
			Title:      req.Title,
			Content:    req.Content,
			CreateTime: gtime.Now(),
			Status:     req.Status,
		},
	}
	return r.Ent
}

func (a *articleAggregation) NewUpdateArticleAggregation(req *in_out.UpdateArticleReq) *entity.Article {
	if req.UpdateTime == nil {
		req.UpdateTime = gtime.Now()
	}
	r := &articleAggregation{
		Ent: &entity.Article{
			ID:         req.ArticleID,
			Title:      req.Title,
			Content:    req.Content,
			CreateTime: req.CreateTime,
			UpdateTime: req.UpdateTime,
		},
	}
	return r.Ent
}
