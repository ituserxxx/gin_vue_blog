package in_out

import "github.com/gogf/gf/os/gtime"

type LoginReq struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type AdminUserIDReq struct {
	UID int `json:"token" binding:"required"`
}
type PageReq struct {
	Page int `json:"page"`
}
type TagPageReq struct {
	Page int `json:"page"`
	All  int `json:"all"` // == -1 所有
}

type AddAdminReq struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type ArticleIDReq struct {
	ArticleID int `json:"article_id" binding:"required"`
}

type UpdateArticleStatusReq struct {
	ID     int `json:"id" binding:"required"`
	Status int `json:"status" `
}

type AddArticleReq struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Status  int    `json:"status" binding:"required"`
	TagList []int  `json:"tag_list"`
}

type UpdateArticleReq struct {
	ArticleID  int         `json:"article_id" binding:"required"`
	Title      string      `json:"title" binding:"required"`
	Content    string      `json:"content" binding:"required"`
	CreateTime *gtime.Time `json:"create_time"`
	UpdateTime *gtime.Time `json:"update_time"`
	TagList    []int       `json:"tag_list"`
}

type IDReq struct {
	ID int `json:"id" binding:"required"`
}

type AddTagReq struct {
	TagName string `json:"tag_name" binding:"required"`
}

type UpdateTagReq struct {
	ID      int    `json:"id" binding:"required"`
	TagName string `json:"tag_name" binding:"required"`
}
