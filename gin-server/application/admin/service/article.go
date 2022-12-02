package service

import (
	"gin-server/application/admin/aggregation"
	"gin-server/application/admin/in_out"
	"gin-server/application/dao"
	"gin-server/application/entity"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/util/gconv"
	"gorm.io/gorm"
	"sync"
)

var Article *articleService

type articleService struct {
}

// 文章列表
func (a *articleService) ArticleList(ctx *gin.Context, req *in_out.PageReq) ([]in_out.ArticleListResp, error) {
	list, err := dao.Article.ArticleList(nil, ctx, req.Page)
	if err != nil {
		return nil, err
	}
	var resp []in_out.ArticleListResp
	if err := gconv.Structs(list, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// 添加文章
func (a *articleService) AddArticle(ctx *gin.Context, req *in_out.AddArticleReq) error {
	// 加锁执行当前方法
	var lock = sync.Mutex{}
	lock.Lock()
	defer lock.Unlock()
	//初始化插入文章的实体
	articleAddEnt := aggregation.Article.NewAddArticleAggregation(req)

	//开启事务
	err := dao.DB(nil, ctx).Transaction(func(tx *gorm.DB) error {
		tx = tx.WithContext(ctx)
		// 添加文章
		articleEnt, err := dao.Article.AddArticle(tx, ctx, articleAddEnt)
		if err != nil {
			return err
		}
		// 绑定标签
		for _, i2 := range req.TagList {
			var tagArticleEnt = entity.TagArticle{
				TagID:     i2,
				ArticleID: articleEnt.ID,
			}
			err := dao.TagArticle.AddRecord(tx, ctx, tagArticleEnt)
			if err != nil {
				return err
			}
		}
		// 立即发布则增加标签的文章数量
		if len(req.TagList) > 0 && req.Status == dao.TableFieldStatus["publish"] {
			// 新增标签文章数量
			err = dao.Tag.UpdateArticleSum(tx, ctx, req.TagList, true)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// 文章详情
func (a *articleService) ArticleDetail(ctx *gin.Context, req *in_out.ArticleIDReq) (*in_out.ArticleInfoResp, error) {
	articleEnt, err := dao.Article.ArticleInfoById(nil, ctx, req.ArticleID)
	if err != nil {
		return nil, err
	}
	var resp *in_out.ArticleInfoResp
	if err := gconv.Struct(articleEnt, &resp); err != nil {
		return nil, err
	}
	TagListEnt, err := dao.Article.GetTagList(nil, ctx, req.ArticleID)
	if err != nil {
		return nil, err
	}
	var tagListResp []int
	for _, i2 := range TagListEnt {
		tagListResp = append(tagListResp, i2.ID)
	}
	resp.TagList = tagListResp
	return resp, nil
}

// 文章更新
func (a *articleService) ArticleUpdate(ctx *gin.Context, req *in_out.UpdateArticleReq) error {
	// 加锁执行当前方法
	var lock = sync.Mutex{}
	lock.Lock()
	defer lock.Unlock()
	//初始化插入文章的实体
	articleUpdateEnt := aggregation.Article.NewUpdateArticleAggregation(req)
	//开启事务
	err := dao.DB(nil, ctx).Transaction(func(tx *gorm.DB) error {
		tx = tx.WithContext(ctx)
		_, err := dao.Article.UpdateArticle(tx, ctx, articleUpdateEnt)
		if err != nil {
			return err
			//return gerror.New("1111")
		}
		// 获取文章的旧标签列表
		oldTagEntList, err := dao.Article.GetTagList(tx, ctx, req.ArticleID)
		if err != nil {
			return gerror.New("2222")
		}
		// 需要增加的
		var needIncrTagArticleSumList []int
		var needCreateTagArticleList []entity.TagArticle
		for _, i2 := range req.TagList {
			mark := false
			// 存在则不需要操作
			for _, i3 := range oldTagEntList {
				if i2 == i3.ID {
					mark = true
					break
				}
			}
			if mark {
				continue
			}
			var ent = entity.TagArticle{
				TagID:     i2,
				ArticleID: req.ArticleID,
			}
			// 不存在则需要添加关系记录
			needCreateTagArticleList = append(needCreateTagArticleList, ent)
			// 增加标签文章数量
			needIncrTagArticleSumList = append(needIncrTagArticleSumList, i2)
		}
		if len(needCreateTagArticleList) > 0 {
			err = dao.TagArticle.AddSomeRecord(tx, ctx, needCreateTagArticleList)
			if err != nil {
				return err
			}
		}
		if len(needIncrTagArticleSumList) > 0 {
			err = dao.Tag.UpdateArticleSum(tx, ctx, needIncrTagArticleSumList, true)
			if err != nil {
				return gerror.New("4444")
				//return err
			}
		}

		// 删除标签文章关系
		var needDelTagArticleRecordIDList []int
		var needDescTagArrticleSunList []int
		for _, i2 := range oldTagEntList {
			mark2 := false
			for _, i4 := range req.TagList {
				if i4 == i2.ID {
					mark2 = true
					break
				}
			}
			if !mark2 {
				needDescTagArrticleSunList = append(needDescTagArrticleSunList, i2.ID)
				needDelTagArticleRecordIDList = append(needDelTagArticleRecordIDList, i2.ID)
			}
		}
		if len(needDelTagArticleRecordIDList) > 0 {
			err = dao.TagArticle.DelSomeRecordByArticleIdAndTagIdList(tx, ctx, needDelTagArticleRecordIDList, req.ArticleID)
			if err != nil {
				return gerror.New(err.Error() + "6666")
				//return err
			}
		}
		if len(needDescTagArrticleSunList) > 0 {
			err = dao.Tag.UpdateArticleSum(tx, ctx, needDescTagArrticleSunList, false)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// 文章更新状态
func (a *articleService) ArticleUpdateStatus(ctx *gin.Context, req *in_out.UpdateArticleStatusReq) error {
	err := dao.Article.ArticleUpdateStatus(nil, ctx, req.ID, req.Status)
	if err != nil {
		return err
	}
	return nil
}

// 文章删除
func (a *articleService) ArticleDelete(ctx *gin.Context, req *in_out.ArticleIDReq) error {
	err := dao.Article.ArticleDelete(nil, ctx, req.ArticleID)
	if err != nil {
		return err
	}
	return nil
}
