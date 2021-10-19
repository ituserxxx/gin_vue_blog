package response

import (
	"gin-server/constant"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ps struct {
	Code constant.ResponseCode `json:"code"`
	Msg  string                `json:"msg"`
	Data interface{}           `json:"data"`
}

func Succ(ctx *gin.Context, data interface{}) {
	resp := ps{
		Code: constant.SUCCESS,
		Msg:  "",
		Data: data,
	}
	ctx.Set("succ_response", resp)
	ctx.JSON(http.StatusOK, resp)
}

func Err(ctx *gin.Context, ErrCode constant.ResponseCode, messge string) {
	resp := ps{
		Code: ErrCode,
		Msg:  messge,
		Data: nil,
	}
	ctx.Set("err_response", resp)
	ctx.JSON(http.StatusOK, resp)
	panic(constant.Error)
}
