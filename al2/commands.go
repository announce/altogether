package al2

import (
	"github.com/urfave/cli"
)

var Commands = []cli.Command{
	{
		Name:  "sync-web",
		Usage: "Sync web-search configs",
		Description: `Required params:
	--alfred-path
	--albert-path`,
		Action: func(c *cli.Context) error { return (&Handler{}).SyncWeb(c) },
		Flags:  optionFlags,
	},
}

var optionFlags = []cli.Flag{
	cli.StringFlag{
		Name:   "alfred-path",
		Usage:  "Specify path to Alfred configuration's root directory",
		EnvVar: "AL2_ALFRED_PATH",
	},
	cli.StringFlag{
		Name:   "albert-path",
		Usage:  "Specify path to Albert configuration's root directory",
		EnvVar: "AL2_ALBERT_PATH",
	},
	cli.BoolFlag{
		Name:   "dry-run",
		Usage:  "Print merged configurations without actual file change",
		EnvVar: "AL2_DRY_RUN",
	},
	cli.BoolFlag{
		Name:   "verbose",
		Usage:  "Print out verbose logs",
		EnvVar: "AL2_VERBOSE",
	},
}
