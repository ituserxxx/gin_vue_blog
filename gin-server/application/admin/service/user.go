package service

import (
	"gin-server/application/dao"
	"gin-server/application/entity"
	"gin-server/constant"
	"github.com/gin-gonic/gin"
)

var User *userService

type userService struct {
}

func (us *userService) Home(ctx *gin.Context) []entity.User {
	res, _ := dao.User.AdminList(nil, ctx, constant.PageMinList)
	return res
}
