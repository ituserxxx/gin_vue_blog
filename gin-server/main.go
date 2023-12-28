package main

import (
	"fmt"
	"gin-server/router"
	"gin-server/utils"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
	"log"
	"net/http"
)

func main() {
	c := g.Config()
	gin.SetMode(c.GetString("server.runMode"))
	r := gin.New()
	r = gin.Default()
	utils.ConnectDB()
	utils.Meili.Init("http://localhost:7700","RWzzZQj1UWqKQLmOaP7HleQIWaRAc8HA-LjpBNfJCAo")
	defer utils.DisconnectDB()
	router.LoadBlogRoutes(r)
	router.LoadAdminRoutes(r)
	srv := &http.Server{
		Addr:    c.GetString("server.serverAddr"),
		Handler: r,
	}
	fmt.Print("服务已启动~"+srv.Addr)
	// 服务连接
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
