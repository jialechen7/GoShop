package main

import (
	"context"
	"goshop/internal/cmd"
	"goshop/internal/consts"
	_ "goshop/internal/logic"
	coupon "goshop/internal/logic/seckill_coupon"
	_ "goshop/internal/packed"
	_ "goshop/utility/captcha"
	_ "goshop/utility/redis_lock"

	"github.com/gogf/gf/v2/frame/g"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

func initStream() {
	// XGROUP CREATE stream.user_coupon g1 0 MKSTREAM
	_, err := g.Redis().Do(context.Background(), "XGROUP", "CREATE",
		consts.StreamUserCoupon, consts.StreamUserCouponGroup, "0", "MKSTREAM")
	if err != nil {
		return
	}
}

func main() {
	if err := consts.LoadLuaSeckillScript(); err != nil {
		g.Dump("Failed to load Lua script: %v\n", err)
		return
	}

	initStream()

	go coupon.UserCouponStreamConsumer(context.Background())

	cmd.Main.Run(gctx.GetInitCtx())
}
