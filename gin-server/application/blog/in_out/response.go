package in_out

import "github.com/gogf/gf/os/gtime"

type HomeArticleListResp struct {
	ArticleList []HomeArticleTitleItem `json:"article_list"`
	Total       int64                  `json:"total"`
}

type HomeArticleTitleItem struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type ArticleDetailResp struct {
	ID         int         `json:"id" `
	Title      string      `json:"title" `
	Content    string      `json:"content" `
	CreateTime *gtime.Time `json:"create_time" `
	UpdateTime *gtime.Time `json:"update_time" `
}
