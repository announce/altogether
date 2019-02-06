package al2

import (
	"github.com/announce/altogether/al2/features"
	"github.com/announce/altogether/al2/util"
	"log"
	"os"
)

type Mode struct {
	Daemon bool
	DryRun bool
}

type Handler struct {
	log        *log.Logger
	AlfredPath string
	AlbertPath string
	Mode       *Mode
}

func (h *Handler) Perform() error {
	h.init()
	if err := h.validate(); err != nil {
		return err
	}
	web := al2.Web{&al2.Pair{
		al2.Launcher{
			Type:     al2.Alfred,
			BasePath: h.AlfredPath,
		},
		al2.Launcher{
			Type:     al2.Albert,
			BasePath: h.AlbertPath,
		},
	}}
	return web.Sync()
}

func (h *Handler) init() {
	h.log = util.CreateLogger(h)
}

func (h *Handler) validate() error {
	if _, err := os.Stat(h.AlfredPath); err != nil {
		return err
	}
	if _, err := os.Stat(h.AlbertPath); err != nil {
		return err
	}
	return nil
}
