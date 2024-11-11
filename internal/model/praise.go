package model

// PraiseGetListInput 获取点赞列表
type PraiseGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Type int // 类型
}

// PraiseGetListOutput 查询列表结果
type PraiseGetListOutput struct {
	List  []PraiseGetListOutputItem `json:"list" description:"列表"`
	Page  int                       `json:"page" description:"分页码"`
	Size  int                       `json:"size" description:"分页数量"`
	Total int                       `json:"total" description:"数据总数"`
}

type PraiseGetListOutputItem struct {
	Id       int `json:"id"` // 自增ID
	UserId   int `json:"user_id"`
	Type     int `json:"type"`
	ObjectId int `json:"object_id"`
	TimeCommon
}

type PraiseCreateUpdateBase struct {
	UserId   int
	Type     int
	ObjectId int
}

type PraiseAddInput struct {
	PraiseCreateUpdateBase
}

type PraiseAddOutput struct {
	PraiseId int
}

type PraiseDeleteInput struct {
	Id int `json:"id"`
}

type PraiseDeleteOutput struct{}

type PraiseDeleteByTypeInput struct {
	Type     int `json:"type"`
	UserId   int `json:"user_id"`
	ObjectId int `json:"object_id"`
}

type PraiseDeleteByTypeOutput struct{}
