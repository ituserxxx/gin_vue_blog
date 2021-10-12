package api

import (
	"encoding/json"
	"fmt"
	"gin-server/application/dao"
	"gin-server/application/entity"
	"gin-server/constant"
	"gin-server/library/response"
	"gin-server/utils"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
	"time"
)

var XxxApi *xxxApi

type xxxApi struct {
}

func (x *xxxApi) Golist(ctx *gin.Context) {
	var entList []entity.User
	err := dao.DB(nil, ctx).Where("id=?", 1).Find(&entList).Error
	if err != nil {
		response.Err(ctx, constant.Error, err.Error())
	}
	redis := utils.Redis()
	str, _ := json.Marshal(entList)
	err = redis.Set(ctx, "user_info", str, 1*time.Hour).Err()
	if err != nil {
		fmt.Print(err.Error())
	}
	val, err := redis.Get(ctx, "user_info").Result()
	if err != nil {
		fmt.Print(err.Error())
	}
	var ss []entity.User
	st := json.Unmarshal([]byte(val), &ss)
	response.Succ(ctx, g.Slice{entList, st})

}
