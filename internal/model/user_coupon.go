package model

// UserCouponGetListInput 获取优惠券列表
type UserCouponGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

// UserCouponGetListOutput 查询列表结果
type UserCouponGetListOutput struct {
	List  []UserCouponGetListOutputItem `json:"list" description:"列表"`
	Page  int                           `json:"page" description:"分页码"`
	Size  int                           `json:"size" description:"分页数量"`
	Total int                           `json:"total" description:"数据总数"`
}

type UserCouponGetListOutputItem struct {
	Id       int `json:"id"`
	UserId   int `json:"user_id"`
	CouponId int `json:"coupon_id"`
	Status   int `json:"status"`
	TimeCommon
}

type UserCouponCreateUpdateBase struct {
	UserId   int `json:"user_id"`
	CouponId int `json:"coupon_id"`
	Status   int `json:"status"`
}

type UserCouponAddInput struct {
	UserCouponCreateUpdateBase
}

type UserCouponAddOutput struct {
	UserCouponId int
}

type UserCouponUpdateInput struct {
	Id int
	UserCouponCreateUpdateBase
}

type UserCouponUpdateOutput struct {
	UserCouponId int
}

type UserCouponDeleteInput struct {
	Id int `json:"id"`
}

type UserCouponDeleteOutput struct{}
