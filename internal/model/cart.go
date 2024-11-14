package model

// CartGetListInput 获取购物车列表
type CartGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

// CartGetListOutput 查询列表结果
type CartGetListOutput struct {
	List  []CartGetListOutputItem `json:"list" description:"列表"`
	Page  int                     `json:"page" description:"分页码"`
	Size  int                     `json:"size" description:"分页数量"`
	Total int                     `json:"total" description:"数据总数"`
}

type CartGetListOutputItem struct {
	Id               int               `json:"id"`                                                // 自增ID
	UserId           int               `json:"user_id"`                                           // 用户ID
	GoodsOptionsId   int               `json:"goods_options_id"`                                  // 商品规格ID
	Count            int               `json:"count"`                                             // 数量
	GoodsOptionsInfo *GoodsOptionsInfo `json:"goods_options_info" orm:"with:id=goods_options_id"` // 商品规格信息
	TimeCommon
}

type CartCreateUpdateBase struct {
	UserId         int
	GoodsOptionsId int
	Count          int
}

type CartAddInput struct {
	CartCreateUpdateBase
}

type CartAddOutput struct {
	CartId int
}

type CartUpdateInput struct {
	CartCreateUpdateBase
	Id int
}

type CartUpdateOutput struct {
	Id int `json:"id"`
}

type CartDeleteInput struct {
	Id int `json:"id"`
}

type CartDeleteOutput struct{}
