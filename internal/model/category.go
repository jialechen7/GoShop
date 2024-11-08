package model

import "github.com/gogf/gf/v2/os/gtime"

// CategoryGetListInput 获取分类列表
type CategoryGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

// CategoryGetListWithStatusInput 获取分类列表
type CategoryGetListWithParentIdInput struct {
	Page     int // 分页号码
	Size     int // 分页数量，最大50
	ParentId int // 父级ID
}

// CategoryGetListOutput 查询列表结果
type CategoryGetListOutput struct {
	List  []CategoryGetListOutputItem `json:"list" description:"列表"`
	Page  int                         `json:"page" description:"分页码"`
	Size  int                         `json:"size" description:"分页数量"`
	Total int                         `json:"total" description:"数据总数"`
}

type CategoryGetListOutputItem struct {
	Id        int         `json:"id"` // 自增ID
	ParentId  int         `json:"parent_id"`
	Name      string      `json:"name"`
	PicUrl    string      `json:"pic_url"`
	Level     int         `json:"level"`
	Sort      int         `json:"sort"`
	CreatedAt *gtime.Time `json:"created_at"`
}

type CategoryAddInput struct {
	Name     string `json:"name"`
	PicUrl   string `json:"pic_url"`
	Level    int    `json:"level"`
	Sort     int    `json:"sort"`
	ParentId int    `json:"parent_id"`
}

type CategoryAddOutput struct {
	CategoryId int
}

type CategoryDeleteInput struct {
	Id int `json:"id"`
}

type CategoryDeleteOutput struct{}
