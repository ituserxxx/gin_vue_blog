package dao

import (
	"gin-server/application/entity"
	"gin-server/constant"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/errors/gerror"
	"gorm.io/gorm"
)

var Article *articleDao

var TableFieldStatus = map[string]int{
	"publish": 1, //发布
	"draft":   2, //草稿
}

type articleDao struct {
}

// 文章列表
func (a *articleDao) ArticleList(tx *gorm.DB, ctx *gin.Context, page int) ([]entity.Article, error) {
	var entList []entity.Article
	//redis.ArticelCache.ArticleList(ctx, page)
	err := DB(tx, ctx).Offset((page - 1) * constant.PageMinList).
		Limit(constant.PageMinList).
		Order("id desc").
		Find(&entList).Error
	if err != nil {
		return nil, err
	}
	return entList, nil
}

// 添加文章
func (a *articleDao) AddArticle(tx *gorm.DB, ctx *gin.Context, article *entity.Article) (*entity.Article, error) {
	err := DB(tx, ctx).Create(&article).Error
	if err != nil {
		return article, err
	}
	return article, nil
}

// 更新文章信息
func (a *articleDao) UpdateArticle(tx *gorm.DB, ctx *gin.Context, article *entity.Article) (*entity.Article, error) {
	err := DB(tx, ctx).Updates(&article).Error
	if err != nil {
		return article, err
	}
	return article, nil
}

// 通过id获取文章行数据
func (a *articleDao) ArticleInfoById(tx *gorm.DB, ctx *gin.Context, id int) (*entity.Article, error) {
	var ent *entity.Article
	r := DB(tx, ctx).Where("id =? ", id).Find(&ent)
	if err := r.Error; err != nil {
		return nil, err
	}
	if r.RowsAffected == 0 {
		return nil, gerror.New("文章已被击杀~")
	}
	return ent, nil
}

// 标签列表
func (a *articleDao) GetTagList(tx *gorm.DB, ctx *gin.Context, id int) ([]entity.Tag, error) {
	var ent []entity.Tag
	err := DB(tx, ctx).Model(entity.Article{}).
		Joins("join blog_tag_article as ta on ta.article_id = blog_article.id").
		Joins("join blog_tag as t on t.id = ta.tag_id").
		Where(" blog_article.id =? ", id).
		Select("t.*").
		Find(&ent).
		Error
	if err != nil {
		return nil, err
	}
	return ent, nil
}

// 博客前端-首页文章列表
func (a *articleDao) GetBlogHomeArticleList(tx *gorm.DB, ctx *gin.Context, page int) ([]entity.Article, error) {
	var ent []entity.Article
	err := DB(tx, ctx).Model(entity.Article{}).
		Where("status =?", TableFieldStatus["publish"]).
		Offset((page - 1) * constant.PageMinList).
		Limit(constant.PageMinList).
		Order("id desc").
		Find(&ent).
		Error
	if err != nil {
		return nil, err
	}
	return ent, nil
}

// 博客前端-首页文章总数
func (a *articleDao) BlogHomeArticleSum(ctx *gin.Context) (int64, error) {
	var cu int64
	err := DB(nil, ctx).Model(entity.Article{}).
		Where("status =?", TableFieldStatus["publish"]).
		Count(&cu).
		Error
	if err != nil {
		return 0, err
	}
	return cu, nil
}

// 文章更新状态
func (a *articleDao) ArticleUpdateStatus(tx *gorm.DB, ctx *gin.Context, id, status int) error {
	err := DB(tx, ctx).Model(entity.Article{}).
		Where("id =?", id).
		Update("status", status).
		Error
	if err != nil {
		return err
	}
	return nil
}

// 文章删除
func (a *articleDao) ArticleDelete(tx *gorm.DB, ctx *gin.Context, id int) error {
	err := DB(tx, ctx).Delete(entity.Article{
		ID: id,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
