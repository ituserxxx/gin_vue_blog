package utils

import (
	"github.com/go-redis/redis/v8"
	"github.com/gogf/gf/frame/g"
)

var instance *redis.Client

func init() {

	c := g.Config()
	//是否开启redis缓存
	if c.GetBool("redis.open") {
		rdb := redis.NewClient(&redis.Options{
			Addr:     c.GetString("redis.host") + ":" + c.GetString("redis.port"), // 地址加端口
			Password: c.GetString("redis.password"),                               // 密码
			DB:       c.GetInt("redis.db"),                                        // use default DB
		})
		instance = rdb
	}
}

func Redis() *redis.Client {
	return instance
}
