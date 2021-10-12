package dao

import (
	"gin-server/application/entity"
	"gin-server/constant"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Visitor *visitorDao

type visitorDao struct {
}

func (v *visitorDao) AddRecord(tx *gorm.DB, ctx *gin.Context, ent *entity.Visitor) error {
	r := DB(tx, ctx).Create(ent)
	if err := r.Error; err != nil {
		return err
	}
	return nil
}

func (v *visitorDao) VisitorList(tx *gorm.DB, ctx *gin.Context, page int) ([]entity.Visitor, error) {
	var entList []entity.Visitor
	err := DB(tx, ctx).
		Offset((page - 1) * constant.PageMinList).
		Limit(constant.PageMinList).
		Order("id desc").
		Find(&entList).Error
	if err != nil {
		return nil, err
	}
	return entList, nil
}
