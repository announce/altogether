package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/DHowett/go-plist"
	"github.com/announce/altogether/al2/util"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Web struct {
	log       *log.Logger
	Launchers *Pair
	*AlfredSites
	AlbertSites
	ConfigDict
}

type Option struct {
	DtyRun  bool
	Verbose bool
}

func (w *Web) Sync(option Option) error {
	w.init()
	if err := w.load(); err != nil {
		w.log.Printf("[Error] Failed to load file stats: %v", err)
		return err
	}
	for _, launcher := range w.Launchers {
		if err := w.parse(launcher); err != nil {
			w.log.Printf("[Error] Failed to parse %s-config: %v", launcher.Type, err)
			return err
		}
	}
	w.merge()
	if option.Verbose {
		w.log.Printf("Launchers: (0,1)=(%v,%v)", w.Launchers[0].Type, w.Launchers[1].Type)
		w.log.Printf("ConfigDict count: %v", len(w.ConfigDict))
		w.log.Printf("DtyRun: %v", option.DtyRun)
	}
	if option.DtyRun {
		if err := w.printMerged(); err != nil {
			w.log.Printf("[Error] Failed to print diff: %v", err)
			return err
		}
	} else {
		if err := w.applyChange(); err != nil {
			w.log.Printf("[Error] Failed to write merged config file: %v", err)
			return err
		}
	}
	return nil
}

func (w *Web) init() {
	w.log = util.CreateLogger(w)
	for _, launcher := range w.Launchers {
		launcher.ConfigPath = filepath.Join(launcher.BasePath, ConfigPath[launcher.Type])
	}
	w.AlfredSites = &AlfredSites{}
	w.ConfigDict = make(map[Id]*SiteConfig)
}

func (w *Web) load() error {
	for _, launcher := range w.Launchers {
		info, err := os.Stat(launcher.ConfigPath)
		if err != nil {
			return err
		}
		launcher.FileInfo = info
	}
	return nil
}

func (w *Web) parse(launcher *Launcher) error {
	file, err := os.Open(launcher.ConfigPath)
	if err != nil {
		return err
	}
	defer func() {
		c := file.Close()
		if c == nil {
			return
		}
		err = fmt.Errorf("failed to close: %v, the original error: %v", c, err)
	}()
	switch launcher.Type {
	case Alfred:
		{
			decoder := plist.NewDecoder(file)
			if err := decoder.Decode(w.AlfredSites); err != nil {
				return err
			}
		}
	case Albert:
		{
			b, err := ioutil.ReadAll(file)
			if err != nil {
				return err
			}
			return json.Unmarshal(b, &w.AlbertSites)
		}
	default:
		w.log.Fatalln("Unexpected type.")
	}
	return nil
}

func (w *Web) merge() {
	// @TODO interception mode to sync deleted config
	sort.Sort(w.Launchers)
	for _, launcher := range w.Launchers {
		switch launcher.Type {
		case Alfred:
			{
				for k, v := range w.AlfredSites.CustomSites {
					v.PreserveUuid(k)
					w.ConfigDict[v.Id()] = v
				}
			}
		case Albert:
			{
				for _, v := range w.AlbertSites {
					w.ConfigDict[v.Id()] = v
				}
			}
		default:
			w.log.Fatalln("Unexpected type.")
		}
	}
}

func (w *Web) printMerged() error {
	// @TODO use given filer to print out
	for _, config := range w.ConfigDict {
		if _, err := fmt.Printf("%#v\n", config); err != nil {
			return err
		}
	}
	return nil
}

const Indent = "  "

func (w *Web) applyChange() error {
	for _, launcher := range w.Launchers {
		switch launcher.Type {
		case Alfred:
			{
				data := &bytes.Buffer{}
				encoder := plist.NewEncoder(data)
				encoder.Indent("\t")
				if err := encoder.Encode(
					w.AlfredSites.Convert(w.ConfigDict)); err != nil {
					return err
				}
				if err := ioutil.WriteFile(
					launcher.ConfigPath,
					data.Bytes(),
					launcher.FileInfo.Mode()); err != nil {
					return err
				}
				return nil
			}
		case Albert:
			{
				j, err := json.MarshalIndent(
					w.AlbertSites.Convert(w.ConfigDict), "", strings.Repeat(Indent, 2))
				if err != nil {
					return err
				}
				if err := ioutil.WriteFile(
					launcher.ConfigPath,
					j,
					launcher.FileInfo.Mode()); err != nil {
					return err
				}
				return nil
			}
		default:
			w.log.Fatalln("Unexpected type.")
		}
	}
	return nil
}
