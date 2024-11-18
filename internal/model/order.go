package model

import "github.com/gogf/gf/v2/os/gtime"

// OrderGetListInput 获取订单列表
type OrderGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

// OrderGetListWithStatusInput 获取订单列表
type OrderGetListWithStatusInput struct {
	Page   int // 分页号码
	Size   int // 分页数量，最大50
	Status int // 订单状态
}

// OrderGetListOutput 查询列表结果
type OrderGetListOutput struct {
	List  []OrderGetListOutputItem `json:"list" description:"列表"`
	Page  int                      `json:"page" description:"分页码"`
	Size  int                      `json:"size" description:"分页数量"`
	Total int                      `json:"total" description:"数据总数"`
}

type OrderGetListOutputItem struct {
	Id               int         `json:"id"` // 自增ID
	Number           string      `json:"number"`
	UserId           int         `json:"userId"`
	PayType          int         `json:"pay_type"`
	Remark           string      `json:"remark"`
	PayAt            *gtime.Time `json:"pay_at"`
	Status           int         `json:"status"`
	ConsigneeName    string      `json:"consignee_name"`
	ConsigneePhone   string      `json:"consignee_phone" `
	ConsigneeAddress string      `json:"consignee_address"`
	Price            int         `json:"price"`
	CouponPrice      int         `json:"coupon_price"`
	ActualPrice      int         `json:"actual_price"`
	TimeCommon
}

type OrderAddInput struct {
	Number             string
	UserId             int
	PayType            int
	Remark             string
	PayAt              *gtime.Time
	Status             int
	ConsigneeName      string
	ConsigneePhone     string
	ConsigneeAddress   string
	Price              int
	CouponPrice        int
	ActualPrice        int
	OrderAddGoodsInfos []*OrderAddGoodsInfo
}

type OrderAddGoodsInfo struct {
	Id             int
	OrderId        int
	GoodsId        int
	GoodsOptionsId int
	Count          int
	Remark         string
	Price          int
	CouponPrice    int
	ActualPrice    int
}

type OrderAddOutput struct {
	OrderId int `json:"order_id"`
}
