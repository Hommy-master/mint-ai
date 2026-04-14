package main

import (
	"github.com/gogf/gf/v2/os/gctx"

	"cozeos/internal/cmd/webserver"
)

func main() {
	webserver.Main.Run(gctx.New())
}
