package al2

import (
	"github.com/urfave/cli"
)

var Commands = []cli.Command{
	{
		Name:        "sync",
		Usage:       "Sync configuration files",
		Description: ``,
		Action:      ExecSync,
		Flags:       optionFlags,
	},
}

var optionFlags = []cli.Flag{
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
		EnvVar: "AL2_DRY_RUN",
	},
	cli.BoolFlag{
		Name:   "verbose",
		Usage:  "Print out verbose logs",
		EnvVar: "AL2_VERBOSE",
	},
}

func ExecSync(c *cli.Context) error {
	h := &Handler{
		AlfredPath: c.String("alfred-path"),
		AlbertPath: c.String("albert-path"),
		Mode: &Mode{
			DryRun:  c.Bool("dry-run"),
			Verbose: c.Bool("verbose"),
		},
		Writer:    c.App.Writer,
		ErrWriter: c.App.ErrWriter,
	}
	return h.Perform()
}
