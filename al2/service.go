package al2

import (
	"github.com/announce/altogether/al2/features"
	"github.com/announce/altogether/al2/util"
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
}

func (h *Handler) Perform() error {
	h.init()
	h.log.Println("Starting with option: ", h.Mode)
	pair := &al2.Pair{&al2.Launcher{
		Type:     al2.Alfred,
		BasePath: h.AlfredPath,
	},
		&al2.Launcher{
			Type:     al2.Albert,
			BasePath: h.AlbertPath,
		}}
	web := &al2.Web{Launchers: pair}
	return web.Sync(al2.Option{
		DtyRun:  h.Mode.DryRun,
		Verbose: h.Mode.Verbose,
	})
}

func (h *Handler) init() {
	h.log = util.CreateLogger(h)
}
