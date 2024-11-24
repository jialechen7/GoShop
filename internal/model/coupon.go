package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CouponGetListInput 获取优惠券列表
type CouponGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

// CouponGetListOutput 查询列表结果
type CouponGetListOutput struct {
	List  []CouponGetListOutputItem `json:"list" description:"列表"`
	Page  int                       `json:"page" description:"分页码"`
	Size  int                       `json:"size" description:"分页数量"`
	Total int                       `json:"total" description:"数据总数"`
}

type CouponGetListWithGoodsIdInput struct {
	GoodsId int
}

// CouponGetListWithGoodsIdOutput 查询所有列表结果
type CouponGetListWithGoodsIdOutput struct {
	List []CouponGetListOutputItem `json:"list" description:"列表"`
}

type CouponGetListAvailableInput struct {
	OrderGoodsInfos interface{} `json:"order_goods_infos"`
}

type CouponGetListAvailableOutput struct {
	AvailableList   []CouponGetListOutputItem `json:"available_list"`
	UnavailableList []CouponGetListOutputItem `json:"unavailable_list"`
}

type CouponGetListOutputItem struct {
	Id         int         `json:"id"`
	Name       string      `json:"name"`
	Condition  int         `json:"condition"`
	Price      int         `json:"price"`
	GoodsIds   string      `json:"goods_ids"`
	CategoryId int         `json:"category_id"`
	Type       int         `json:"type"`
	Stock      int         `json:"stock"`
	Reason     string      `json:"reason"`
	StartTime  *gtime.Time `json:"start_time"`
	EndTime    *gtime.Time `json:"end_time"`
	TimeCommon
}

type CouponCreateUpdateBase struct {
	Name       string `json:"name"`
	Condition  int    `json:"condition"`
	Price      int    `json:"price"`
	GoodsIds   string `json:"goods_ids"`
	CategoryId int    `json:"category_id"`
	Type       int    `json:"type"`
}

type CouponAddInput struct {
	CouponCreateUpdateBase
}

type CouponAddOutput struct {
	CouponId int
}

type CouponUpdateInput struct {
	Id int
	CouponCreateUpdateBase
}

type CouponUpdateOutput struct {
	CouponId int
}

type CouponDeleteInput struct {
	Id int `json:"id"`
}

type CouponDeleteOutput struct{}
