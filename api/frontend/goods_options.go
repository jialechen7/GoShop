package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GoodsOptionsGetListCommonReq struct {
	g.Meta `path:"/goods_options/list" tags:"GoodsOptions" method:"get" summary:"商品规格列表接口"`
	CommonPaginationReq
	GoodsId int `json:"goods_id" v:"min:1#请选择需要查询的商品,required#请选择商品id" form:"query" dc:"商品id"`
}

type GoodsOptionsGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type GoodsOptionsDetailReq struct {
	g.Meta `path:"/goods_options/detail" method:"get" tags:"商品规格" summary:"商品规格详情接口"`
	Id     int `v:"min:1#请选择需要查询的商品规格" form:"query" dc:"商品规格id"`
}

type GoodsOptionsDetailRes struct {
	Id      int    `json:"id"`
	GoodsId int    `json:"goods_id"`
	PicUrl  string `json:"pic_url"`
	Name    string `json:"name"`
	Price   int    `json:"price"`
	Stock   int    `json:"stock"`
}
