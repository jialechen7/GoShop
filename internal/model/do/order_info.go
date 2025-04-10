// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderInfo is the golang structure of table order_info for DAO operations like Where/Data.
type OrderInfo struct {
	g.Meta           `orm:"table:order_info, do:true"`
	Id               interface{} // 订单id，使用基于Redis自增的全局唯一id
	Number           interface{} // 订单编号
	UserId           interface{} // 用户id
	PayType          interface{} // 支付方式 1微信 2支付宝 3云闪付
	Remark           interface{} // 备注
	PayAt            *gtime.Time // 支付时间
	Status           interface{} // 订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价 5已评价
	ConsigneeName    interface{} // 收货人姓名
	ConsigneePhone   interface{} // 收货人手机号
	ConsigneeAddress interface{} // 收货人详细地址
	Price            interface{} // 订单金额 单位分
	CouponPrice      interface{} // 优惠券金额 单位分
	ActualPrice      interface{} // 实际支付金额 单位分
	CreatedAt        *gtime.Time //
	UpdatedAt        *gtime.Time //
}
