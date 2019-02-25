package web

import (
	"fmt"
	"github.com/announce/altogether/al2/domain"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var Config = domain.ConfigPath{
	domain.Alfred: "preferences/features/websearch/prefs.plist",
	domain.Albert: "org.albert.extension.websearch/engines.json",
}

type Launcher struct {
	Type       domain.Type
	BasePath   string
	ConfigPath string
	FileInfo   os.FileInfo
	Sites      Sites
}

func (l *Launcher) Init() {
	switch l.Type {
	case domain.Alfred:
		{
			l.Sites = &AlfredSites{}
		}
	case domain.Albert:
		{
			l.Sites = &AlbertSites{}
		}
	default:
		panic(domain.UnexpectedType)
	}
	l.ConfigPath = filepath.Join(l.BasePath, Config[l.Type])
}

func (l *Launcher) Load() error {
	info, err := os.Stat(l.ConfigPath)
	if err != nil {
		return err
	}
	l.FileInfo = info
	return nil
}

func (l *Launcher) Mtime() time.Time {
	return l.FileInfo.ModTime()
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
	return l.Sites.Decode(file)
}

func (l *Launcher) Populate(dict ConfigDict) {
	l.Sites.Populate(dict)
}

func (l *Launcher) Diff(diff Diff) {
	for _, v := range l.Sites.List() {
		if order, ok := diff[v.Id()]; ok {
			diff[v.Id()] = append(order, l.Type)
		} else {
			diff[v.Id()] = []domain.Type{l.Type}
		}
	}
}

func (l *Launcher) Save(dict ConfigDict) error {
	output, err := l.Sites.Encode(dict)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(
		l.ConfigPath,
		output,
		l.FileInfo.Mode())
}
