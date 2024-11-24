package model

import "github.com/gogf/gf/v2/os/gtime"

// SeckillCouponGetListInput 获取秒杀优惠券列表
type SeckillCouponGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

// SeckillCouponGetListOutput 查询列表结果
type SeckillCouponGetListOutput struct {
	List  []SeckillCouponGetListOutputItem `json:"list" description:"列表"`
	Page  int                              `json:"page" description:"分页码"`
	Size  int                              `json:"size" description:"分页数量"`
	Total int                              `json:"total" description:"数据总数"`
}

type SeckillCouponGetListOutputItem struct {
	CouponId   int         `json:"coupon_id"`
	Stock      int         `json:"stock"`
	StartTime  *gtime.Time `json:"start_time"`
	EndTime    *gtime.Time `json:"end_time"`
	CouponInfo CouponInfo  `json:"coupon_info" orm:"with:id=coupon_id"`
	TimeCommon
}

type SeckillCouponCreateUpdateBase struct {
	CouponId  int         `json:"coupon_id"`
	Stock     int         `json:"stock"`
	StartTime *gtime.Time `json:"start_time"`
	EndTime   *gtime.Time `json:"end_time"`
}

type SeckillCouponAddInput struct {
	Name       string `json:"name"`
	Condition  int    `json:"condition"`
	Price      int    `json:"price"`
	GoodsIds   string `json:"goods_ids"`
	CategoryId int    `json:"category_id"`
	Type       int    `json:"type"`
	SeckillCouponCreateUpdateBase
}

type SeckillCouponAddOutput struct {
	SeckillCouponId int
}

type SeckillCouponUpdateInput struct {
	Id int
	SeckillCouponCreateUpdateBase
}

type SeckillCouponUpdateOutput struct {
	SeckillCouponId int
}

type SeckillCouponDeleteInput struct {
	Id int
}

type SeckillCouponDeleteOutput struct{}
