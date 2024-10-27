package main

import (
	"goshop/internal/cmd"
	_ "goshop/internal/logic"
	_ "goshop/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
