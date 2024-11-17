package model

// CommentGetListInput 获取评论列表
type CommentGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

type CommentGetListFrontendInput struct {
	Page     int // 分页码
	Size     int // 分页数量
	Type     int // 评论类型
	ObjectId int // 评论对象ID
}

// CommentGetListOutput 查询列表结果
type CommentGetListOutput struct {
	List  []CommentGetListOutputItem `json:"list" description:"列表"`
	Page  int                        `json:"page" description:"分页码"`
	Size  int                        `json:"size" description:"分页数量"`
	Total int                        `json:"total" description:"数据总数"`
}

type CommentGetListOutputItem struct {
	Id       int    `json:"id"` // 自增ID
	ParentId int    `json:"parent_id"`
	UserId   int    `json:"user_id"`
	ObjectId int    `json:"object_id"`
	Type     int    `json:"type"`
	Content  string `json:"content"`
	TimeCommon
}

type CommentCreateUpdateBase struct {
	ParentId int
	UserId   int
	ObjectId int
	Type     int
	Content  string
}

type CommentAddInput struct {
	CommentCreateUpdateBase
}

type CommentAddOutput struct {
	CommentId int
}

type CommentUpdateInput struct {
	CommentCreateUpdateBase
	Id int
}

type CommentUpdateOutput struct{}

type CommentDeleteInput struct {
	Id int `json:"id"`
}

type CommentDeleteOutput struct{}
