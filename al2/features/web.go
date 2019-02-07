package al2

import (
	"encoding/json"
	"fmt"
	"github.com/DHowett/go-plist"
	"github.com/announce/altogether/al2/util"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
)

type Type int

const (
	Alfred Type = iota
	Albert
)

var ConfigPath = map[Type]string{
	Alfred: "preferences/features/websearch/prefs.plist",
	Albert: "org.albert.extension.websearch/engines.json",
}

type Launcher struct {
	Type       Type
	BasePath   string
	ConfigPath string
	FileInfo   os.FileInfo
}

func (l Launcher) Mtime() int64 {
	return l.FileInfo.ModTime().Unix()
}

type Pair [2]*Launcher

func (p Pair) Len() int {
	return len(p)
}

func (p Pair) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Pair) Less(i, j int) bool {
	return p[i].Mtime() < p[j].Mtime()
}

type Web struct {
	log       *log.Logger
	Launchers *Pair
	*AlfredPrefs
	*AlbertEngines
}

type Option struct {
	DtyRun bool
}

func (w *Web) Sync(option Option) error {
	w.init()
	direction, err := w.compare()
	if err != nil {
		w.log.Println("[Error] Failed to compare mtime", err)
		return err
	}
	if direction == nil {
		w.log.Println("Aborting... Both target has the same mtime.")
		return nil
	}
	if err := w.parse(direction[0]); err != nil {
		w.log.Println("[Error] Failed to parse 0-config:", direction[0], err)
		return err
	}
	if err := w.parse(direction[1]); err != nil {
		w.log.Println("[Error] Failed to parse 1-config:", direction[1], err)
		return err
	}
	w.log.Printf("w.AlfredPrefs: %+v,\nw.AlbertEngines: %+v", w.AlfredPrefs, w.AlbertEngines)

	// @TODO os.Chtimes to avoid loop
	return nil
}

func (w *Web) init() {
	w.log = util.CreateLogger(w)
	for _, launcher := range w.Launchers {
		launcher.ConfigPath = filepath.Join(launcher.BasePath, ConfigPath[launcher.Type])
	}
	w.AlfredPrefs = &AlfredPrefs{}
	w.AlbertEngines = &AlbertEngines{}
}

func (w *Web) compare() (*Pair, error) {
	for _, launcher := range w.Launchers {
		info, err := os.Stat(launcher.ConfigPath)
		if err != nil {
			return nil, err
		}
		launcher.FileInfo = info
	}
	if w.Launchers[0].Mtime() == w.Launchers[1].Mtime() {
		return nil, nil
	}
	sort.Sort(w.Launchers)
	return w.Launchers, nil
}

type AlfredSite struct {
	Enabled bool   `plist:"enabled"`
	Keyword string `plist:"keyword"`
	Text    string `plist:"text"`
	Url     string `plist:"url"`
	Utf8    bool   `plist:"utf8"`
}
type AlfredPrefs struct {
	CustomSites map[string]AlfredSite `plist:"customSites"`
}

type AlbertEngines []AlbertEngine
type AlbertEngine struct {
	IconPath string
	Name     string
	Trigger  string
	Url      string
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
			return decoder.Decode(w.AlfredPrefs)
		}
	case Albert:
		{
			b, err := ioutil.ReadAll(file)
			if err != nil {
				return err
			}
			return json.Unmarshal(b, w.AlbertEngines)
		}
	default:
		w.log.Fatalln("Unexpected type.")
	}
	return nil
}
