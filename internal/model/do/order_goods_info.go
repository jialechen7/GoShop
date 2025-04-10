// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderGoodsInfo is the golang structure of table order_goods_info for DAO operations like Where/Data.
type OrderGoodsInfo struct {
	g.Meta         `orm:"table:order_goods_info, do:true"`
	Id             interface{} // 商品维度的订单表
	OrderId        interface{} // 关联的主订单表
	GoodsId        interface{} // 商品id
	GoodsOptionsId interface{} // 商品规格id(sku id)
	Count          interface{} // 商品数量
	Remark         interface{} // 备注
	Price          interface{} // 订单金额 单位分
	CouponPrice    interface{} // 优惠券金额 单位分
	ActualPrice    interface{} // 实际支付金额 单位分
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
}
