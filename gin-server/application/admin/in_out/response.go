package in_out

import (
	"github.com/gogf/gf/os/gtime"
)

type LoginResp struct {
	Token string `json:"token"`
}

type AdminUserInfoResp struct {
	Username     string `json:"username"`
	Id           int    `json:"id"`
	Avatar       string `json:"avatar"`
	Introduction string `json:"introduction"`
	Nickname     string `json:"nickname"`
	About        string `json:"about"`
}

type ArticleListResp struct {
	ID         int         `json:"id" `
	Title      string      `json:"title" `
	Content    string      `json:"content"`
	CreateTime *gtime.Time `json:"create_time" `
	UpdateTime *gtime.Time `json:"update_time"`
	Status     int         `json:"status"`
}

type ArticleInfoResp struct {
	ID         int         `json:"id" `
	Title      string      `json:"title" `
	Content    string      `json:"content"`
	CreateTime *gtime.Time `json:"create_time" `
	UpdateTime *gtime.Time `json:"update_time"`
	Status     int         `json:"status"`
	TagList    []int       `json:"tag_list"`
}

type TagInfoResp struct {
	ID         int    `json:"id"`
	TagName    string `json:"tag_name"`
	ArticleSum string `json:"article_sum" `
}

type IDResp struct {
	ID int `json:"id"`
}
