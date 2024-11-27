package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type SeckillCouponGetListCommonReq struct {
	g.Meta `path:"/seckill/coupon/list" tags:"秒杀优惠券后台" method:"get" summary:"秒杀优惠券列表接口"`
	CommonPaginationReq
}

type SeckillCouponGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type SeckillCouponAddUpdateCommon struct {
	Stock     int         `json:"stock" form:"stock" v:"required#请输入库存" description:"库存"`
	StartTime *gtime.Time `json:"start_time" form:"start_time" description:"开始时间"`
	EndTime   *gtime.Time `json:"end_time" form:"end_time" description:"结束时间"`
}

type SeckillCouponAddReq struct {
	g.Meta     `path:"/seckill/coupon/add" tags:"秒杀优惠券后台" method:"post" summary:"创建秒杀优惠券接口"`
	Name       string `json:"name" form:"name" v:"required#请输入秒杀优惠券名称" description:"秒杀优惠券名称"`
	Condition  int    `json:"condition" form:"condition" v:"required#请输入秒杀优惠券条件" description:"秒杀优惠券条件(单位分)"`
	Price      int    `json:"price" form:"price" v:"required#请输入秒杀优惠券金额" description:"秒杀优惠券金额(单位分)"`
	GoodsIds   string `json:"goods_ids" form:"goods_ids" description:"可用的商品id列表"`
	CategoryId int    `json:"category_id" form:"category_id" v:"required#请选择分类" description:"可用的分类id"`
	SeckillCouponAddUpdateCommon
}

type SeckillCouponAddRes struct {
	SeckillCouponId int `json:"seckill_coupon_id"`
}

type SeckillCouponDeleteReq struct {
	g.Meta `path:"/seckill/coupon/delete" method:"delete" tags:"秒杀优惠券后台" summary:"删除秒杀优惠券接口"`
	Id     int `v:"min:1#请选择需要删除的秒杀优惠券" dc:"秒杀优惠券id"`
}
type SeckillCouponDeleteRes struct{}

type SeckillCouponUpdateReq struct {
	g.Meta `path:"/seckill/coupon/update" method:"post" tags:"秒杀优惠券后台" summary:"修改秒杀优惠券接口"`
	Id     int `json:"id" form:"id" v:"required#请输入id" dc:"秒杀优惠券id"`
	SeckillCouponAddUpdateCommon
}
type SeckillCouponUpdateRes struct{}
