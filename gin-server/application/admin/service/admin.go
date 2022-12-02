package service

import (
	"gin-server/application/admin/aggregation"
	"gin-server/application/admin/in_out"
	"gin-server/application/dao"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
)

var Admin *adminService

type adminService struct {
}

//后台账号添加
func (ad *adminService) AddAdmin(ctx *gin.Context, req *in_out.AddAdminReq) error {
	userEntity := aggregation.User.NewAddAdminAggregation(req)
	err := dao.User.AddAdmin(nil, ctx, userEntity)
	if err != nil {
		return err
	}
	return nil
}

//后台账号删除
func (ad *adminService) DelAdmin(ctx *gin.Context, req *in_out.AdminUserIDReq) error {
	err := dao.User.DelAdmin(nil, ctx, req.UID)
	if err != nil {
		return err
	}

	return nil
}

//后台账号列表
func (ad *adminService) AdminList(ctx *gin.Context, req *in_out.PageReq) ([]in_out.AdminUserInfoResp, error) {
	list, err := dao.User.AdminList(nil, ctx, req.Page)
	if err != nil {
		return nil, err
	}
	var resp []in_out.AdminUserInfoResp
	if err := gconv.Struct(list, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
