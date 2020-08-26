package request

// 文章接口结构体
type ArticleListStruct struct{
	ArticeID string `json:"articleId"`
	Search string `json:"search"`
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}