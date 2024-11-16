package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CartGetListCommonReq struct {
	g.Meta `path:"/cart/list" tags:"Cart" method:"get" summary:"购物车列表接口"`
	CommonPaginationReq
}

type CartGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type CartAddReq struct {
	g.Meta         `path:"/cart/add" tags:"Cart" method:"post" summary:"创建购物车接口"`
	GoodsOptionsId int `json:"goods_options_id" form:"goods_options_id" v:"required#请选择商品规格" dc:"商品规格id"`
	Count          int `json:"count" form:"count" v:"required#请输入数量" dc:"数量"`
}

type CartAddRes struct {
	CartId int `json:"cart_id"`
}

type CartDeleteReq struct {
	g.Meta `path:"/cart/delete" method:"delete" tags:"Cart" summary:"删除购物车接口"`
	Ids    []int `json:"ids" v:"required#请选择需要删除的购物车" dc:"购物车ids"`
}
type CartDeleteRes struct{}

type CartUpdateReq struct {
	g.Meta         `path:"/cart/update" method:"post" tags:"购物车" summary:"修改购物车接口"`
	Id             int `json:"id" form:"id" v:"required#请输入id" dc:"购物车id"`
	GoodsOptionsId int `json:"goods_options_id" form:"goods_options_id" v:"required#请选择商品规格" dc:"商品规格id"`
	Count          int `json:"count" form:"count" v:"required#请输入数量" dc:"数量"`
}
type CartUpdateRes struct {
	Id int
}
