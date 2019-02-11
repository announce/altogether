package al2

import (
	"github.com/announce/altogether/al2/util"
	"github.com/announce/altogether/al2/web"
	"io"
	"log"
)

type Mode struct {
	DryRun  bool
	Verbose bool
}

type Handler struct {
	log        *log.Logger
	AlfredPath string
	AlbertPath string
	Mode       *Mode
	Writer     io.Writer
	ErrWriter  io.Writer
}

func (h *Handler) Perform() error {
	h.init()
	h.log.Println("Starting with option: ", h.Mode)
	pair := &web.Pair{&web.Launcher{
		Type:     web.Alfred,
		BasePath: h.AlfredPath,
	},
		&web.Launcher{
			Type:     web.Albert,
			BasePath: h.AlbertPath,
		}}
	w := &web.Web{
		Launchers: pair,
		Out:       h.Writer,
		ErrOut:    h.ErrWriter,
	}
	return w.Sync(web.Option{
		DtyRun:  h.Mode.DryRun,
		Verbose: h.Mode.Verbose,
	})
}

func (h *Handler) init() {
	h.log = util.CreateLogger(h.ErrWriter, h)
	h.log.Println("ErrOut:", h.ErrWriter)
}
