package model

type UserCreateUpdateBase struct {
	Name         string
	Password     string
	UserSalt     string
	Avatar       string
	Sign         string
	Sex          int
	SecretAnswer string
	Status       int
}

type UserCreateInput struct {
	UserCreateUpdateBase
}

type UserCreateOutput struct {
	Id int `json:"id"`
}

type UserUpdateInput struct {
	UserCreateUpdateBase
	Id int
}

type UserUpdateOutput struct{}

// UserGetListInput 获取用户列表
type UserGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

// UserGetListOutput 查询列表结果
type UserGetListOutput struct {
	List  []UserGetListOutputItem `json:"list" description:"列表"`
	Page  int                     `json:"page" description:"分页码"`
	Size  int                     `json:"size" description:"分页数量"`
	Total int                     `json:"total" description:"数据总数"`
}

type UserGetListOutputItem struct {
	Id     int    `json:"id"` // 自增ID
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Sex    string `json:"sex"`
	Sign   string `json:"sign"`
	Status int    `json:"status"`
	TimeCommon
}
