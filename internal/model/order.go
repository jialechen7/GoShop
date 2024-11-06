package model

import "github.com/gogf/gf/v2/os/gtime"

// OrderGetListInput 获取订单列表
type OrderGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
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
	CreatedAt        *gtime.Time `json:"created_at"`
}
