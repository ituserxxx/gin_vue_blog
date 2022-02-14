package dao

import (
	"gin-server/application/entity"
	"gin-server/constant"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/errors/gerror"
	"gorm.io/gorm"
)

var TagArticle *tagArticleDao

type tagArticleDao struct {
}

func (ta *tagArticleDao) AddRecord(tx *gorm.DB, ctx *gin.Context, ent entity.TagArticle) error {
	r := DB(tx, ctx).Create(&ent)
	if err := r.Error; err != nil {
		return err
	}
	return nil
}
func (ta *tagArticleDao) AddSomeRecord(tx *gorm.DB, ctx *gin.Context, entList []entity.TagArticle) error {
	r := DB(tx, ctx).Create(&entList)
	if err := r.Error; err != nil {
		return err
	}
	return nil
}
func (ta *tagArticleDao) IsExistsRecor(tx *gorm.DB, ctx *gin.Context, ent entity.TagArticle) (bool, error) {
	var c int64
	err := DB(tx, ctx).Where("tag_id =? ", ent.TagID).
		Where("article_id =? ", ent.ArticleID).
		Count(&c).Error
	if err != nil {
		return false, err
	}
	if c == 1 {
		return true, nil
	}
	return false, gerror.New("err count tag_article")
}

func (ta *tagArticleDao) DelSomeRecordByArticleIdAndTagIdList(tx *gorm.DB, ctx *gin.Context, tagIdList []int, articleId int) error {
	err := DB(tx, ctx).Debug().Where("article_id = ?", articleId).
		Where("tag_id IN ? ", tagIdList).
		Delete(entity.TagArticle{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (ta *tagArticleDao) TagArticleList(tx *gorm.DB, ctx *gin.Context, tagId, page int) ([]entity.TagArticle, error) {
	var entList []entity.TagArticle
	result := DB(tx, ctx).Where("tag_id =?", tagId).
		Offset((page - 1) * constant.PageMinList).
		Limit(constant.PageMinList).
		Order("id desc").
		Find(&entList)
	return entList, result.Error
}

func (a *tagArticleDao) GetBlogTagArticleList(tx *gorm.DB, ctx *gin.Context, tagID, page int) ([]entity.Article, error) {
	var ent []entity.Article
	err := DB(tx, ctx).Model(entity.Article{}).
		Joins("join blog_tag_article as ta on ta.article_id = blog_article.id").
		Where("blog_article.status =?", TableFieldStatus["publish"]).
		Where("ta.tag_id =?", tagID).
		//Offset((page - 1) * constant.PageMinList).
		//Limit(constant.PageMinList).
		Order("id desc").
		Find(&ent).
		Error
	if err != nil {
		return nil, nil
	}
	return ent, nil
}
func (a *tagArticleDao) GetBlogTagArticleSum(ctx *gin.Context, tagID int) (int64, error) {
	var cu int64
	err := DB(nil, ctx).Model(entity.Article{}).
		Joins("join blog_tag_article as ta on ta.article_id = blog_article.id").
		Where("blog_article.status =?", TableFieldStatus["publish"]).
		Where("ta.tag_id =?", tagID).
		Count(&cu).
		Error
	if err != nil {
		return 0, nil
	}
	return cu, nil
}
