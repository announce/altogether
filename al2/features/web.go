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
	"strings"
)

type Type int

const (
	Alfred Type = iota
	Albert
)

func (t Type) String() string {
	return [...]string{"Alfred", "Albert"}[t]
}

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
	AlfredSites
	AlbertSites
	ConfigDict map[string]NormalizableConfig
}

type Option struct {
	DtyRun  bool
	Verbose bool
}

func (w *Web) Sync(option Option) error {
	w.init()
	if err := w.load(); err != nil {
		return err
	}
	for _, launcher := range w.Launchers {
		if err := w.parse(launcher); err != nil {
			w.log.Printf("[Error] Failed to parse %s-config: %v", launcher.Type, err)
			return err
		}
	}
	sort.Sort(w.Launchers)
	if option.Verbose {
		w.log.Printf("Launchers: (0,1)=(%v,%v)", w.Launchers[0].Type, w.Launchers[1].Type)
	}
	w.merge()
	if option.Verbose {
		w.log.Printf("ConfigDict: %+v", w.ConfigDict)
	}
	if option.DtyRun {

	} else {

	}
	return nil
}

func (w *Web) init() {
	w.log = util.CreateLogger(w)
	for _, launcher := range w.Launchers {
		launcher.ConfigPath = filepath.Join(launcher.BasePath, ConfigPath[launcher.Type])
	}
	w.ConfigDict = make(map[string]NormalizableConfig)
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

type NormalizableConfig interface {
	Normalize() NormalizableConfig
	//Name()
	//Url()
	//Trigger()
}
type AlfredSites struct {
	CustomSites map[string]AlfredSite `plist:"customSites"`
}

type AlfredSite struct {
	Enabled bool   `plist:"enabled"`
	Trigger string `plist:"keyword"`
	Name    string `plist:"text"`
	Url     string `plist:"url"`
	Utf8    bool   `plist:"utf8"`
}

func (a AlfredSite) Normalize() NormalizableConfig {
	a.Url = strings.Replace(a.Url, "{query}", "%s", -1)
	a.Name = strings.Replace(a.Name, "{query}", "%s", -1)
	return a
}

type AlbertSites []AlbertSite

type AlbertSite struct {
	IconPath string
	Name     string
	Trigger  string
	Url      string
}

func (a AlbertSite) Normalize() NormalizableConfig {
	return a
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
			return decoder.Decode(&w.AlfredSites)
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
	for _, launcher := range w.Launchers {
		switch launcher.Type {
		case Alfred:
			{
				for _, v := range w.AlfredSites.CustomSites {
					w.ConfigDict[v.Trigger] = v.Normalize()
				}
			}
		case Albert:
			{
				for _, v := range w.AlbertSites {
					w.ConfigDict[v.Trigger] = v.Normalize()
				}
			}
		default:
			w.log.Fatalln("Unexpected type.")
		}
	}
}
