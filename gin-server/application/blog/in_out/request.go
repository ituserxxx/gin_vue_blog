package in_out

type PageReq struct {
	Page int `json:"page"`
}

type IDReq struct {
	ID int `json:"id" binding:"required"`
}

type TagArticleListReq struct {
	ID   int `json:"id" binding:"required"`
	Page int `json:"page" binding:"required"`
}

type SearchArticleListReq struct {
	Content string `json:"content"`
}
