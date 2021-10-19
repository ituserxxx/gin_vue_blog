package dao

import (
	"fmt"
	"gin-server/application/entity"
	"gin-server/constant"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Tag *tagDao

type tagDao struct {
}

func (t *tagDao) TagList(tx *gorm.DB, ctx *gin.Context, page int) ([]entity.Tag, error) {
	var entList []entity.Tag
	m := DB(tx, ctx)
	if page > -1 {
		m = m.Offset((page - 1) * constant.PageMaxList).Limit(constant.PageMaxList)
	}
	result := m.Order("id desc").Find(&entList)
	return entList, result.Error
}

func (t *tagDao) AddTag(tx *gorm.DB, ctx *gin.Context, tag *entity.Tag) error {
	r := DB(tx, ctx).Create(tag)
	if err := r.Error; err != nil {
		return err
	}
	return nil
}

func (t *tagDao) TagInfo(tx *gorm.DB, ctx *gin.Context, id int) (*entity.Tag, error) {
	var ent *entity.Tag
	r := DB(tx, ctx).Where("id =? ", id).Find(&ent)
	if err := r.Error; err != nil {
		return ent, err
	}
	return ent, nil
}

func (t *tagDao) TagUpdate(tx *gorm.DB, ctx *gin.Context, tag *entity.Tag) error {
	r := DB(tx, ctx).Updates(tag)
	if err := r.Error; err != nil {
		return err
	}
	return nil
}

func (t *tagDao) TagDel(tx *gorm.DB, ctx *gin.Context, id int) error {
	r := DB(tx, ctx).
		Where("id =?", id).
		Delete(&entity.Tag{})
	if err := r.Error; err != nil {
		return err
	}
	return nil
}

func (t *tagDao) UpdateArticleSum(tx *gorm.DB, ctx *gin.Context, tagIdList []int, isIncr bool) error {
	ope := "+"
	if !isIncr {
		ope = "-"
	}
	str := "article_sum %s %d"
	str = fmt.Sprintf(str, ope, 1)
	r := DB(tx, ctx).Model(entity.Tag{}).
		Where("id IN ?", tagIdList).
		UpdateColumn("article_sum", gorm.Expr(str))
	if err := r.Error; err != nil {
		return err
	}
	return nil
}
