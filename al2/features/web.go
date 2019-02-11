package al2

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/DHowett/go-plist"
	"github.com/announce/altogether/al2/util"
	"github.com/google/uuid"
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

type Id [sha1.Size]byte
type ConfigDict map[Id]*NormalizableConfig
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
		if err := w.printDiff(); err != nil {
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
	w.ConfigDict = make(map[Id]*NormalizableConfig)
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

type NormalizableConfig struct {
	Uuid     string
	Enabled  bool   `plist:"enabled" json:"-"`
	Utf8     bool   `plist:"utf8" json:"-"`
	Trigger  string `plist:"keyword" json:"trigger"`
	Name     string `plist:"text" json:"name"`
	Url      string `plist:"url" json:"url"`
	IconPath string `plist:"-" json:"iconPath"`
}

func (a *NormalizableConfig) Id() Id {
	a.Normalize()
	b := bytes.Buffer{}
	b.WriteString(a.Trigger)
	b.WriteString(a.Url)
	return sha1.Sum(b.Bytes())
}

func (a *NormalizableConfig) PreserveUuid(key string) {
	a.Uuid = key
}

func (a *NormalizableConfig) Normalize() {
	a.Albert()
}

const Spacer = " "

func (a *NormalizableConfig) Albert() {
	a.Url = strings.Replace(a.Url, "{query}", "%s", -1)
	a.Name = strings.Replace(a.Name, "{query}", "%s", -1)
	a.Trigger = strings.Trim(a.Trigger, Spacer) + Spacer
}

func (a *NormalizableConfig) Alfred() {
	a.Url = strings.Replace(a.Url, "%s", "{query}", -1)
	a.Name = strings.Replace(a.Name, "%s", "{query}", -1)
	a.Trigger = strings.Trim(a.Trigger, Spacer)
}

type CustomSites map[string]*NormalizableConfig
type AlfredSites struct {
	CustomSites `plist:"customSites"`
}

func (a *AlfredSites) Convert(dict ConfigDict) AlfredSites {
	sites := make(CustomSites)
	for _, site := range dict {
		config := site
		if config.Uuid == "" {
			config.Uuid = uuid.New().String()
		}
		config.Alfred()
		sites[config.Uuid] = config
	}
	return AlfredSites{sites}
}

type AlbertSites []*NormalizableConfig

func (a *AlbertSites) Convert(dict ConfigDict) AlbertSites {
	var sites AlbertSites
	for _, site := range dict {
		config := site
		config.Albert()
		sites = append(sites, config)
	}
	return sites
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
	// @TODO interception mode
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

func (w *Web) printDiff() error {
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
				encoder.Indent(Indent)
				if err := encoder.Encode(
					w.AlfredSites.Convert(w.ConfigDict)); err != nil {
					return err
				}
				//if err := ioutil.WriteFile(
				//	launcher.ConfigPath,
				//	data.Bytes(),
				//	launcher.FileInfo.Mode()); err != nil {
				//	return err
				//}
				return nil
			}
		case Albert:
			{
				//j, err := json.MarshalIndent(
				//	w.AlbertSites.Convert(w.ConfigDict), "", Indent)
				//if err != nil {
				//	return err
				//}
				//if err := ioutil.WriteFile(
				//	launcher.ConfigPath,
				//	j,
				//	launcher.FileInfo.Mode()); err != nil {
				//	return err
				//}
				return nil
			}
		default:
			w.log.Fatalln("Unexpected type.")
		}
	}
	return nil
}
