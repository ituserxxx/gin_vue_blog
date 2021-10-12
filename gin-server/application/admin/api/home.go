package api

import (
	"gin-server/application/admin/service"
	"gin-server/library/response"
	"github.com/gin-gonic/gin"
)

var HomeApi *homeApi

type homeApi struct {
}

func (h *homeApi) Home(ctx *gin.Context) {

	response.Succ(ctx, service.User.Home(ctx))
}
