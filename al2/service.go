package al2

import (
	"github.com/announce/altogether/al2/domain"
	"github.com/announce/altogether/al2/util"
	"github.com/announce/altogether/al2/web"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"log"
	"strings"
)

type Handler struct {
	log *log.Logger
}

func (h *Handler) SyncWeb(c *cli.Context) error {
	h.log = util.CreateLogger(c.App.ErrWriter, h)
	h.log.Printf("Starting sync-web %+v", c.App.Version)
	if err := h.verifyRequiredParams(c); err != nil {
		return err
	}
	pair := &web.Pair{&web.Launcher{
		Type:     domain.Alfred,
		BasePath: c.String("alfred-path"),
	},
		&web.Launcher{
			Type:     domain.Albert,
			BasePath: c.String("albert-path"),
		}}
	w := &web.Web{
		Launchers: pair,
		Out:       c.App.Writer,
		ErrOut:    c.App.ErrWriter,
	}
	return w.Sync(web.Option{
		DryRun:  c.Bool("dry-run"),
		Verbose: c.Bool("verbose"),
	})
}

func (h *Handler) verifyRequiredParams(c *cli.Context) error {
	var messages []string
	if c.String("alfred-path") == "" {
		messages = append(messages, "specify required option: alfred-path")
	}
	if c.String("albert-path") == "" {
		messages = append(messages, "specify required option: albert-path")
	}
	if len(messages) > 0 {
		return errors.New(strings.Join(messages, "\n"))
	}
	return nil
}
