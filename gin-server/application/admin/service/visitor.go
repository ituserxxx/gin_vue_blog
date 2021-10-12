package service

import (
	"gin-server/application/admin/in_out"
	"gin-server/application/dao"
	"gin-server/application/entity"
	"github.com/gin-gonic/gin"
)

var Visitor *visitorService

type visitorService struct {
}

//后台账号列表
func (v *visitorService) VisitorList(ctx *gin.Context, req *in_out.PageReq) ([]entity.Visitor, error) {
	list, err := dao.Visitor.VisitorList(nil, ctx, req.Page)
	if err != nil {
		return nil, err
	}
	return list, nil
}
