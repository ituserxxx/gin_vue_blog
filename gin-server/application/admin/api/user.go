package api

import (
	"fmt"
	"gin-server/application/admin/in_out"
	"gin-server/application/admin/service"
	"gin-server/constant"
	"gin-server/library/response"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
)

var UserApi *userApi

type userApi struct {
}

// 用户信息
func (l *userApi) AdminUserInfo(ctx *gin.Context) {
	uid, _ := ctx.Get("uid")
	resp, err := service.Login.AdminInfo(ctx, gconv.Int(uid))
	if err != nil {
		response.Err(ctx, constant.Handle_Error, err.Error())
	}
	response.Succ(ctx, resp)
}

// 用户退出
func (l *userApi) AdminUserLogout(ctx *gin.Context) {
	uid := ctx.GetHeader("X-Token")
	fmt.Printf("用户退出了-_>", uid)
	response.Succ(ctx, "")
}

// 添加管理员
func (l *userApi) AddAdmin(ctx *gin.Context) {
	var req *in_out.AddAdminReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.Err(ctx, constant.Error, err.Error())
	}
	err := service.Admin.AddAdmin(ctx, req)
	if err != nil {
		response.Err(ctx, constant.Handle_Error, err.Error())
	}
	response.Succ(ctx, "ok")
}

// 删除管理员
func (l *userApi) DelAdmin(ctx *gin.Context) {
	var req *in_out.AdminUserIDReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.Err(ctx, constant.Error, err.Error())
	}
	err := service.Admin.DelAdmin(ctx, req)
	if err != nil {
		response.Err(ctx, constant.Handle_Error, err.Error())
	}
	response.Succ(ctx, "ok")
}

// 管理员列表
func (l *userApi) AdminList(ctx *gin.Context) {
	var req *in_out.PageReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.Err(ctx, constant.Error, err.Error())
	}
	list, err := service.Admin.AdminList(ctx, req)
	if err != nil {
		response.Err(ctx, constant.Handle_Error, err.Error())
	}
	response.Succ(ctx, list)
}
