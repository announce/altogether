package web

import (
	"fmt"
	"github.com/announce/altogether/al2/util"
	"log"
)

type Web struct {
	log       *log.Logger
	Launchers *Pair
	ConfigDict
}

type Option struct {
	DtyRun  bool
	Verbose bool
}

func (w *Web) Sync(option Option) error {
	w.init()
	if err := w.Launchers.Load(); err != nil {
		w.log.Printf("[Error] Failed to load file: %v", err)
		return err
	}
	w.Launchers.Merge(w.ConfigDict)
	if option.Verbose {
		w.log.Printf("Launchers: (0,1)=(%v,%v)", w.Launchers[0].Type, w.Launchers[1].Type)
		w.log.Printf("ConfigDict count: %v", len(w.ConfigDict))
		w.log.Printf("DtyRun: %v", option.DtyRun)
	}
	if option.DtyRun {
		// @TODO use given filer to print out
		fmt.Println(w.ConfigDict)
	} else {
		if err := w.Launchers.Save(w.ConfigDict); err != nil {
			w.log.Printf("[Error] Failed to write merged config file: %v", err)
			return err
		}
	}
	return nil
}

func (w *Web) init() {
	w.log = util.CreateLogger(w)
	w.ConfigDict = make(map[Id]*SiteConfig)
}
