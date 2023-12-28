package router

import (
	blogApi "gin-server/application/blog/api"
	"gin-server/middleware"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
)

func LoadBlogRoutes(r *gin.Engine) {
	r.Use(middleware.CorsMiddleware(),middleware.Visitor())
	api := r.Group(g.Config().GetString("app.BlogBaseUrl"))
	{
		api.POST("/article/list", blogApi.ArticleApi.HomeArticleList)
		api.POST("/article/detail", blogApi.ArticleApi.ArticleDetail)
		api.POST("/tag/list", blogApi.TagApi.TagList)
		api.POST("/tag/article/list", blogApi.TagApi.TagArticleList)
		api.POST("/search/article", blogApi.ArticleApi.HomeArticleList)
	}

}
