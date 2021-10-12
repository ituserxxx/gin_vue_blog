package dao

import (
	"errors"
	"gin-server/application/entity"
	"gin-server/constant"
	"gin-server/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DB(tx *gorm.DB, ctx *gin.Context) *gorm.DB {
	//判断是否开启事务
	if tx != nil {
		return tx
	}
	return utils.DB(ctx)
}

var User *userDao

type userDao struct {
}

func (u *userDao) AddAdmin(tx *gorm.DB, ctx *gin.Context, user *entity.User) error {
	err := DB(tx, ctx).Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
func (u *userDao) DelAdmin(tx *gorm.DB, ctx *gin.Context, id int) error {
	err := DB(tx, ctx).Delete(&entity.User{
		Id: id,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
func (u *userDao) AdminList(tx *gorm.DB, ctx *gin.Context, page int) ([]entity.User, error) {
	var entList []entity.User
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
func (u *userDao) GetLoginUserInfoByPassword(tx *gorm.DB, ctx *gin.Context, user *entity.User) (*entity.User, error) {
	var ent *entity.User
	err := DB(tx, ctx).Where("username =? ", user.Username).
		Where("password=? ", utils.EncodeMD5(user.Password)).
		Find(&ent).
		Error
	if err != nil {
		return nil, err
	}
	if ent.Id == 0 {
		return nil, errors.New("no user")
	}
	return ent, nil
}
func (u *userDao) GetAdminInfoById(tx *gorm.DB, ctx *gin.Context, user *entity.User) (*entity.User, error) {
	result := DB(tx, ctx).Where("id=? ", user.Id).Find(&user)
	if err := result.Error; err != nil {
		return user, err
	}
	return user, nil
}
