package envinit

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"

	"cozeos/util"
)

var (
	Main = gcmd.Command{
		Name:        "envinit",
		Brief:       "Initialize backend runtime environment",
		Description: "Initialize required database tables from Go models.",
		Usage:       "envinit [OPTION]",
		Examples: `
			Run:
				./envinit -c config.yaml`,
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
			if parser.GetOpt("version") != nil {
				util.PrintVersionInfo()
				return nil
			}

			cfgPath := parser.GetOpt("config").String()
			if cfgPath == "" {
				cfgPath = "../../manifest/config/config.yaml"
			}

			return start(ctx, cfgPath)
		},
	}
)
