package model

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

type CouponGetListOutputItem struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	GoodsIds   string `json:"goods_ids"`
	CategoryId int    `json:"category_id"`
	TimeCommon
}

type CouponCreateUpdateBase struct {
	Name       string `json:"name"`
	Price      int    `json:"price"`
	GoodsIds   string `json:"goods_ids"`
	CategoryId int    `json:"category_id"`
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
