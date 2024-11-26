package main

import (
	"goshop/internal/cmd"
	_ "goshop/internal/logic"
	_ "goshop/internal/packed"
	_ "goshop/utility/captcha"
	_ "goshop/utility/redis_lock"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
