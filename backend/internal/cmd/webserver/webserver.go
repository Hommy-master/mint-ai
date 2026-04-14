package webserver

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"

	"cozeos/util"
)

var (
	Main = gcmd.Command{
		Name:        "webserver",
		Brief:       "An Web Server for ai-chat-x",
		Description: "The web service program that interfaces with the WeChat Official Account.",
		Usage:       "webserver [OPTION]",
		Examples: `
			Run:
				./webserver -c config.yaml`,
		Additional: "Find more information at: README.md",
		Arguments: []gcmd.Argument{
			{
				Name:   "version",
				Short:  "v",
				Brief:  "print version info",
				IsArg:  false,
				Orphan: true,
			},
			{
				Name:   "config",
				Short:  "c",
				Brief:  "config file (default config.yaml)",
				IsArg:  false,
				Orphan: false,
			},
		},
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 判断不带数据的选项是否存在时，可以通过GetOpt(name) != nil方式
			ver := parser.GetOpt("version")
			if ver != nil {
				util.PrintVersionInfo()
				return nil
			}

			config := parser.GetOpt("config").String()
			if config == "" {
				config = "../../manifest/config/config.yaml"
			}

			return start(ctx, config)
		},
	}
)
