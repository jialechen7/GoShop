package model

// ArticleGetListInput 获取文章列表
type ArticleGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

// ArticleGetListOutput 查询列表结果
type ArticleGetListOutput struct {
	List  []ArticleGetListOutputItem `json:"list" description:"列表"`
	Page  int                        `json:"page" description:"分页码"`
	Size  int                        `json:"size" description:"分页数量"`
	Total int                        `json:"total" description:"数据总数"`
}

type ArticleGetListOutputItem struct {
	Id      int    `json:"id"` // 自增ID
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Detail  string `json:"detail"`
	PicUrl  string `json:"pic_url"`
	IsAdmin int    `json:"is_admin"`
	TimeCommon
}

type ArticleCreateUpdateBase struct {
	UserId  int
	Title   string
	Desc    string
	PicUrl  string
	Detail  string
	IsAdmin int
}

type ArticleAddInput struct {
	ArticleCreateUpdateBase
}

type ArticleAddOutput struct {
	ArticleId int
}

type ArticleDeleteInput struct {
	Id int `json:"id"`
}

type ArticleDeleteOutput struct{}

type ArticleDetailInput struct {
	Id int
}

type ArticleDetailOutput struct {
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	Detail string `json:"detail"`
	PicUrl string `json:"pic_url"`
	TimeCommon
}
