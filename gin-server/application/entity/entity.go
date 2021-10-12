package entity

import "github.com/gogf/gf/os/gtime"

type Article struct {
	ID         int         `gorm:"primary_key" json:"id" `
	Title      string      `json:"title" `
	Content    string      `json:"content" `
	CreateTime *gtime.Time `json:"create_time" `
	UpdateTime *gtime.Time `json:"update_time" `
	Status     int         `json:"status" db:"status"` // 1-发布，2-草稿，3-下架
}

type Tag struct {
	ID         int    `gorm:"primary_key" json:"id" `
	TagName    string `json:"tag_name" `
	ArticleSum int    `json:"article_sum" `
}

type TagArticle struct {
	Id        int `gorm:"primaryKey"`
	TagID     int `json:"tag_id"`
	ArticleID int `json:"article_id"`
}

type User struct {
	Id           int    `gorm:"primaryKey" json:"id" `
	Username     string `json:"username"`
	Password     string `json:"password"`
	Avatar       string `json:"avatar"`
	Introduction string `json:"introduction"`
	Nickname     string `json:"nickname"`
	About        string `json:"about"`
}

type Visitor struct {
	ID        uint64      `gorm:"primaryKey" json:"id" db:"id"`
	Uri       string      `json:"uri" db:"uri" `
	Referer   string      `json:"referer" `
	Ua        string      `json:"ua" db:"ua"`
	IpAddress string      `json:"ip_address"`
	IP        string      `json:"ip" `
	VisitTime *gtime.Time `json:"visit_time" `
}
