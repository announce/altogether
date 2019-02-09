package al2

import (
	"bytes"
	"crypto/sha1"
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

type Id [sha1.Size]byte
type ConfigDict map[Id]NormalizableConfig
type Web struct {
	log       *log.Logger
	Launchers *Pair
	AlfredSites
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
	sort.Sort(w.Launchers)
	w.merge()
	if option.Verbose {
		w.log.Printf("Launchers: (0,1)=(%v,%v)", w.Launchers[0].Type, w.Launchers[1].Type)
		w.log.Printf("ConfigDict: %+v", w.ConfigDict)
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
	w.ConfigDict = make(map[Id]NormalizableConfig)
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
	Id() Id
}
type AlfredSites struct {
	CustomSites map[string]AlfredSite `plist:"customSites"`
}

type AlfredSite struct {
	Enabled  bool   `plist:"enabled" json:"-"`
	Trigger  string `plist:"keyword" json:"trigger"`
	Name     string `plist:"text" json:"name"`
	Url      string `plist:"url" json:"url"`
	Utf8     bool   `plist:"utf8" json:"-"`
	IconPath string `plist:"-" json:"iconPath"`
}

func (a AlfredSite) Normalize() NormalizableConfig {
	a.Url = strings.Replace(a.Url, "{query}", "%s", -1)
	a.Name = strings.Replace(a.Name, "{query}", "%s", -1)
	a.Trigger = a.Trigger + " "
	return a
}

func (a AlfredSite) Id() Id {
	n := a.Normalize().(AlfredSite)
	b := bytes.Buffer{}
	b.WriteString(n.Trigger)
	b.WriteString(n.Url)
	return sha1.Sum(b.Bytes())
}

type AlbertSites []AlbertSite

func (a AlbertSites) Convert(dict ConfigDict) []NormalizableConfig {
	var configs []NormalizableConfig
	for _, c := range dict {
		configs = append(configs, c)
	}
	return configs
}

type AlbertSite struct {
	IconPath string `json:"iconPath"`
	Name     string `json:"name"`
	Trigger  string `json:"trigger"`
	Url      string `json:"url"`
}

func (a AlbertSite) Normalize() NormalizableConfig {
	return a
}

func (a AlbertSite) Id() Id {
	n := a.Normalize().(AlbertSite)
	b := bytes.Buffer{}
	b.WriteString(n.Trigger)
	b.WriteString(n.Url)
	return sha1.Sum(b.Bytes())
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
	// @TODO interception mode
	for _, launcher := range w.Launchers {
		switch launcher.Type {
		case Alfred:
			{
				for _, v := range w.AlfredSites.CustomSites {
					w.ConfigDict[v.Id()] = v.Normalize()
				}
			}
		case Albert:
			{
				for _, v := range w.AlbertSites {
					w.ConfigDict[v.Id()] = v.Normalize()
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

func (w *Web) applyChange() error {
	for _, launcher := range w.Launchers {
		switch launcher.Type {
		case Alfred:
			{
				//// Wrong dict key and invalid query format
				//// No enough information to build plist
				//data := &bytes.Buffer{}
				//encoder := plist.NewEncoder(data)
				//err := encoder.Encode(w.ConfigDict)
				//if err != nil {
				//	fmt.Println(err)
				//}
				//if err := ioutil.WriteFile(
				//	launcher.ConfigPath,
				//	data.Bytes(),
				//	launcher.FileInfo.Mode()); err != nil {
				//	return err
				//}
			}
		case Albert:
			{
				j, err := json.MarshalIndent(
					w.AlbertSites.Convert(w.ConfigDict), "", "  ")
				if err != nil {
					return err
				}
				if err := ioutil.WriteFile(
					launcher.ConfigPath,
					j,
					launcher.FileInfo.Mode()); err != nil {
					return err
				}
			}
		default:
			w.log.Fatalln("Unexpected type.")
		}
	}
	return nil
}
