package main

import (
	"github.com/gogf/gf/v2/os/gctx"

	"cozeos/internal/cmd/envinit"
)

func main() {
	envinit.Main.Run(gctx.New())
}
