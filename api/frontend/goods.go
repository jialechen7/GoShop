package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GoodsGetListByLevelReq struct {
	g.Meta `path:"/goods/level/list/" tags:"商品前台" method:"get" summary:"商品列表接口"`
	CommonPaginationReq
	LevelId int `json:"level_id" v:"min:1#请选择需要查询的分类" form:"query" dc:"分类ID"`
}

type GoodsGetListCommonReq struct {
	g.Meta `path:"/goods/list" tags:"商品前台" method:"get" summary:"商品列表接口"`
	CommonPaginationReq
}

type GoodsGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type GoodsDetailReq struct {
	g.Meta `path:"/goods/detail" method:"get" tags:"商品前台" summary:"商品详情接口"`
	Id     int `v:"min:1#请选择需要查询的商品" form:"query" dc:"商品id"`
}

type GoodsDetailRes struct {
	Id               int    `json:"id"`
	PicUrl           string `json:"pic_url"`
	Name             string `json:"name"`
	Price            int    `json:"price"`
	Level1CategoryId int    `json:"level1_category_id"`
	Level2CategoryId int    `json:"level2_category_id"`
	Level3CategoryId int    `json:"level3_category_id"`
	Brand            string `json:"brand"`
	Stock            int    `json:"stock"`
	Sale             int    `json:"sale"`
	Tags             string `json:"tags"`
	DetailInfo       string `json:"detail_info"`
	IsPraise         int    `json:"is_praise"`
	IsCollect        int    `json:"is_collect"`
}
