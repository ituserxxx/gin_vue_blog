package redis

import (
	"gin-server/application/entity"
	"github.com/gin-gonic/gin"
)

var ArticelCache *articleCache

type articleCache struct {
}

func (ac *articleCache) ArticleList(ctx *gin.Context, page int) ([]entity.Article, error) {
	return nil, nil
	//redis := utils.Redis()
	//str, _ := json.Marshal(entList)
	//err = redis.Set(ctx, "user_info", str, 1*time.Hour).Err()
	//if err != nil {
	//	fmt.Print(err.Error())
	//}
	//val, err := redis.Get(ctx, "user_info").Result()
	//if err != nil {
	//	fmt.Print(err.Error())
	//}
	//var ss []entity.User
	//st := json.Unmarshal([]byte(val), &ss)
}
