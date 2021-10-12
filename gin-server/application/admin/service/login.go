package service

import (
	"gin-server/application/admin/aggregation"
	"gin-server/application/admin/in_out"
	"gin-server/application/dao"
	"gin-server/utils"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
)

var Login *loginService

type loginService struct {
}

//后台账号密码登录
func (l *loginService) AdminLogin(ctx *gin.Context, req *in_out.LoginReq) (*in_out.LoginResp, error) {
	var resp = &in_out.LoginResp{}
	userEntity := aggregation.User.NewAdminLoginAggregation(req)
	userEntity, err := dao.User.GetLoginUserInfoByPassword(nil, ctx, userEntity)
	if err != nil {
		return nil, err
	}
	// 生成token
	token := utils.GenerateToken(ctx, gconv.String(userEntity.Id))
	resp.Token = token
	return resp, nil
}

//后台账号信息
func (l *loginService) AdminInfo(ctx *gin.Context, uid int) (*in_out.AdminUserInfoResp, error) {
	userEntity := aggregation.User.NewGetAdminInfoAggregation(uid)
	userEntity, err := dao.User.GetAdminInfoById(nil, ctx, userEntity)
	if err != nil {
		return nil, err
	}
	var resp *in_out.AdminUserInfoResp
	if err = gconv.Struct(userEntity, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
