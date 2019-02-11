package web

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

var ConfigPath = map[Type]string{
	Alfred: "preferences/features/websearch/prefs.plist",
	Albert: "org.albert.extension.websearch/engines.json",
}

var UnexpectedType = errors.New("unexpected type")

type Type int

const (
	Alfred Type = iota
	Albert
)

func (t Type) String() string {
	return [...]string{"Alfred", "Albert"}[t]
}

type Pair [2]*Launcher

func (p *Pair) Len() int {
	return len(p)
}

func (p *Pair) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *Pair) Less(i, j int) bool {
	return p[i].Mtime() < p[j].Mtime()
}

func (p *Pair) Load() error {
	for _, launcher := range p {
		launcher.Init()
		if err := launcher.Load(); err != nil {
			return err
		}
		if err := launcher.Parse(); err != nil {
			return err
		}
	}
	return nil
}

func (p *Pair) Merge(dict ConfigDict) {
	sort.Sort(p)
	for _, launcher := range p {
		launcher.Populate(dict)
	}
}

func (p *Pair) Save(dict ConfigDict) error {
	for _, launcher := range p {
		if err := launcher.Save(dict); err != nil {
			return err
		}
	}
	return nil
}

type Launcher struct {
	Type       Type
	BasePath   string
	ConfigPath string
	FileInfo   os.FileInfo
	*AlfredSites
	AlbertSites
}

func (l *Launcher) Init() {
	l.AlfredSites = &AlfredSites{}
	l.ConfigPath = filepath.Join(l.BasePath, ConfigPath[l.Type])
}

func (l *Launcher) Load() error {
	info, err := os.Stat(l.ConfigPath)
	if err != nil {
		return err
	}
	l.FileInfo = info
	return nil
}

func (l *Launcher) Mtime() int64 {
	return l.FileInfo.ModTime().Unix()
}

func (l *Launcher) Parse() error {
	file, err := os.Open(l.ConfigPath)
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
	switch l.Type {
	case Alfred:
		{
			return l.AlfredSites.Decode(file)
		}
	case Albert:
		{
			return l.AlbertSites.Decode(file)
		}
	default:
		panic(UnexpectedType)
	}
	return nil
}

func (l *Launcher) Populate(dict ConfigDict) {
	// @TODO interception mode to sync deleted config
	switch l.Type {
	case Alfred:
		{
			for k, v := range l.AlfredSites.CustomSites {
				v.PreserveUuid(k)
				dict[v.Id()] = v
			}
		}
	case Albert:
		{
			for _, v := range l.AlbertSites {
				dict[v.Id()] = v
			}
		}
	default:
		panic(UnexpectedType)
	}

}

func (l *Launcher) Save(dict ConfigDict) error {
	var output []byte
	var err error
	if l.Type == Alfred {
		output, err = l.AlfredSites.Encode(dict)
	} else if l.Type == Albert {
		output, err = l.AlbertSites.Encode(dict)
	}
	if err != nil {
		return err
	}
	return ioutil.WriteFile(
		l.ConfigPath,
		output,
		l.FileInfo.Mode())
}
