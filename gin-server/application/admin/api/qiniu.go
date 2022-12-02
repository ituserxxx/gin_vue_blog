package api

import (
	"gin-server/application/admin/in_out"
	"gin-server/library/response"
	"gin-server/utils"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
)

var QiniuApi *qiniuApi

type qiniuApi struct {
}

func (qa *qiniuApi) GetToken(ctx *gin.Context) {
	var resp = &in_out.QiniuResponse{}
	resp.Token = utils.GetUploadToken()
	resp.Domain = g.Config().GetString("qiniu.domain")
	//resp.Region = g.Config().GetString("qiniu.region")
	resp.Region = "https://up-z0.qiniup.com"
	response.Succ(ctx, resp)
}
