package al2

import (
	"github.com/urfave/cli"
)

var Commands = []cli.Command{
	{
		Name:        "sync",
		Usage:       "Sync configuration files",
		Description: ``,
		//Action: jira.PrintAttachmentStats,
		Flags: adminFlags,
	},
}

var adminFlags = []cli.Flag{
	cli.StringFlag{
		Name:   "alfred-path",
		Usage:  "Specify path to Alfred configuration root directory",
		EnvVar: "AL2_ALFRED_PATH",
	},
	cli.StringFlag{
		Name:   "albert-path",
		Usage:  "Specify path to Albert configuration root directory",
		EnvVar: "AL2_ALBERT_PATH",
	},
	cli.BoolFlag{
		Name:   "dry-run",
		Usage:  "Print configuration diff without actual file change",
		EnvVar: "AL2_PATH_TO_ALBERT",
	},
	cli.BoolFlag{
		Name:   "daemon",
		Usage:  "Sync configurations based on file change detection",
		EnvVar: "AL2_DAEMON",
	},
}
