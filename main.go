package main

import (
	_ "goshop/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"goshop/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
