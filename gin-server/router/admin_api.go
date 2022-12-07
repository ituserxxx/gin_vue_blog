package router

import (
	Admin "gin-server/application/admin/api"
	"gin-server/middleware"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
)

func LoadAdminRoutes(r *gin.Engine) {
	prefix := g.Config().GetString("app.BlogAdminBaseUrl")
	r.Use(middleware.CorsMiddleware(),middleware.Visitor())
	r.POST(prefix+"/user/login", Admin.LoginApi.AdminLogin)
	r.Use(middleware.Jwt())
	api := r.Group(prefix)
	{
		api.POST("/xx", Admin.XxxApi.Golist)
		api.POST("/user/info", Admin.UserApi.AdminUserInfo)
		api.POST("/user/logout", Admin.UserApi.AdminUserLogout)

		api.POST("/user/add", Admin.UserApi.AddAdmin)
		api.POST("/user/del", Admin.UserApi.DelAdmin)
		api.POST("/list", Admin.UserApi.AdminList)
		api.POST("/article/list", Admin.ArticleApi.ArticleList)
		api.POST("/article/add", Admin.ArticleApi.AddArticle)
		api.POST("/article/detail", Admin.ArticleApi.ArticleDetail)
		api.POST("/article/update", Admin.ArticleApi.ArticleUpdate)
		api.POST("/article/draft", Admin.ArticleApi.ArticleSaveDraft)
		api.POST("/article/publish", Admin.ArticleApi.ArticlePublish)
		api.POST("/article/delete", Admin.ArticleApi.ArticleDelete)
		api.POST("/tag/list", Admin.TagApi.TagList)
		api.POST("/tag/add", Admin.TagApi.AddTag)
		api.POST("/tag/detail", Admin.TagApi.TagDetail)
		api.POST("/tag/update", Admin.TagApi.TagUpdate)
		api.POST("/tag/del", Admin.TagApi.TagDel)
		api.POST("/visitor/list", Admin.VisitorApi.VisitorList)

	}
}
