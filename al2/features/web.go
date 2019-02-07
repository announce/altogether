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

type Pair [2]*Launcher
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
		w.log.Println("[Error] Failed to parse 0-config:", direction[0])
		return err
	}
	if err := w.parse(direction[1]); err != nil {
		w.log.Println("[Error] Failed to parse 1-config:", direction[1])
		return err
	}

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
	if w.Launchers[0].FileInfo.ModTime().Unix() == w.Launchers[1].FileInfo.ModTime().Unix() {
		return nil, nil
	}

	sort.Slice(w.Launchers, func(i, j int) bool {
		return w.Launchers[i].FileInfo.ModTime().Unix() < w.Launchers[j].FileInfo.ModTime().Unix()
	})
	return w.Launchers, nil
}

type AlfredSite struct {
	Enabled bool
	Keyword string
	Text    string
	Url     string
	Utf8    bool
}
type AlfredPrefs struct {
	Key   string
	Value map[string]AlfredSite
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
	defer func() {
		c := file.Close()
		if c == nil {
			return
		}
		err = fmt.Errorf("failed to close: %v, the original error: %v", c, err)
	}()
	if err != nil {
		return err
	}
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
