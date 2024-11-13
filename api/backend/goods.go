package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GoodsGetListCommonReq struct {
	g.Meta `path:"/goods/list" tags:"Goods" method:"get" summary:"商品列表接口"`
	CommonPaginationReq
}

type GoodsGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type GoodsAddReq struct {
	g.Meta           `path:"/goods/add" tags:"Goods" method:"post" summary:"创建商品接口"`
	GoodsId          int    `json:"goods_id"`
	PicUrl           string `json:"pic_url" form:"pic_url" v:"required#请上传图片" dc:"图片地址"`
	Name             string `json:"name" form:"name" v:"required#请输入名称" dc:"名称"`
	Price            int    `json:"price" form:"price" v:"required#请输入价格" dc:"价格"`
	Level1CategoryId int    `json:"level1_category_id" form:"level1_category_id" v:"required#请输入一级分类" dc:"一级分类"`
	Level2CategoryId int    `json:"level2_category_id" form:"level2_category_id" v:"required#请输入二级分类" dc:"二级分类"`
	Level3CategoryId int    `json:"level3_category_id" form:"level3_category_id" v:"required#请输入三级分类" dc:"三级分类"`
	Brand            string `json:"brand" form:"brand" v:"required#请输入品牌" dc:"品牌"`
	Stock            int    `json:"stock" form:"stock" v:"required#请输入库存" dc:"库存"`
	Sale             int    `json:"sale" form:"sale" d:"0" dc:"销量"`
	Tags             string `json:"tags" form:"tags" v:"required#请输入标签" dc:"标签"`
	DetailInfo       string `json:"detail_info" form:"detail_info" v:"required#请输入详情" dc:"详情"`
}

type GoodsAddRes struct {
	Id int `json:"id"`
}

type GoodsDeleteReq struct {
	g.Meta `path:"/goods/delete" method:"delete" tags:"Goods" summary:"删除商品接口"`
	Id     int `v:"min:1#请选择需要删除的商品" dc:"商品id"`
}
type GoodsDeleteRes struct{}

type GoodsUpdateReq struct {
	g.Meta           `path:"/goods/update" method:"post" tags:"商品" summary:"修改商品接口"`
	Id               int    `json:"id" form:"id" v:"required#请输入id" dc:"商品id"`
	PicUrl           string `json:"pic_url" form:"pic_url" v:"required#请上传图片" dc:"图片地址"`
	Name             string `json:"name" form:"name" v:"required#请输入名称" dc:"名称"`
	Price            int    `json:"price" form:"price" v:"required#请输入价格" dc:"价格"`
	Level1CategoryId int    `json:"level1_category_id" form:"level1_category_id" v:"required#请输入一级分类" dc:"一级分类"`
	Level2CategoryId int    `json:"level2_category_id" form:"level2_category_id" v:"required#请输入二级分类" dc:"二级分类"`
	Level3CategoryId int    `json:"level3_category_id" form:"level3_category_id" v:"required#请输入三级分类" dc:"三级分类"`
	Brand            string `json:"brand" form:"brand" v:"required#请输入品牌" dc:"品牌"`
	Stock            int    `json:"stock" form:"stock" v:"required#请输入库存" dc:"库存"`
	Sale             int    `json:"sale" form:"sale" v:"required#请输入销量" dc:"销量"`
	Tags             string `json:"tags" form:"tags" v:"required#请输入标签" dc:"标签"`
	DetailInfo       string `json:"detail_info" form:"detail_info" v:"required#请输入详情" dc:"详情"`
}
type GoodsUpdateRes struct{}

type GoodsDetailReq struct {
	g.Meta `path:"/goods/detail" method:"get" tags:"商品" summary:"商品详情接口"`
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
}
