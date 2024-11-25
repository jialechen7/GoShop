package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type SeckillCouponKillReq struct {
	g.Meta   `path:"/seckill/coupon/kill" method:"post" tags:"秒杀优惠券后台" summary:"秒杀优惠券接口"`
	CouponId int `json:"coupon_id" form:"coupon_id" v:"required#请输入id" dc:"秒杀优惠券id"`
}

type SeckillCouponKillRes struct{}
