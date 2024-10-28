package model

import "github.com/gogf/gf/v2/os/gtime"

// RotationCreateUpdateBase 创建/修改轮播图基类
type RotationCreateUpdateBase struct {
	PicUrl string
	Link   string
	Sort   int
}

// RotationCreateInput 创建轮播图
type RotationCreateInput struct {
	RotationCreateUpdateBase
}

// RotationCreateOutput 创建轮播图返回结果
type RotationCreateOutput struct {
	RotationId int `json:"rotation_id"`
}

// RotationUpdateInput 修改轮播图
type RotationUpdateInput struct {
	RotationCreateUpdateBase
	Id int
}

// RotationGetListInput 获取轮播图列表
type RotationGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型
}

// RotationGetListOutput 查询列表结果
type RotationGetListOutput struct {
	List  []RotationGetListOutputItem `json:"list" description:"列表"`
	Page  int                         `json:"page" description:"分页码"`
	Size  int                         `json:"size" description:"分页数量"`
	Total int                         `json:"total" description:"数据总数"`
}

type RotationGetListOutputItem struct {
	//Rotation  *RotationListItem `json:"Rotation"`
	Id        int         `json:"id"` // 自增ID
	PicUrl    string      `json:"picUrl"`
	Link      string      `json:"link"`
	Sort      int         `json:"sort"`       // 排序，数值越高越靠前
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}

// RotationListItem 主要用于轮播图列表展示
type RotationListItem struct {
	Id        int         `json:"id"` // 自增ID
	PicUrl    string      `json:"picUrl"`
	Link      string      `json:"link"`
	Sort      int         `json:"sort"`       // 排序，数值越高越靠前
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}
