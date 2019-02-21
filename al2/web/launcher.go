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
	*AlfredSites
	AlbertSites
}

func (l *Launcher) Init() {
	l.AlfredSites = &AlfredSites{}
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
	switch l.Type {
	case domain.Alfred:
		{
			return l.AlfredSites.Decode(file)
		}
	case domain.Albert:
		{
			return l.AlbertSites.Decode(file)
		}
	default:
		panic(domain.UnexpectedType)
	}
	return nil
}

func (l *Launcher) Populate(dict ConfigDict) {
	// @TODO interception mode to sync deleted config
	switch l.Type {
	case domain.Alfred:
		{
			for k, v := range l.AlfredSites.CustomSites {
				v.PreserveUuid(k)
				dict[v.Id()] = v
			}
		}
	case domain.Albert:
		{
			for _, v := range l.AlbertSites {
				dict[v.Id()] = v
			}
		}
	default:
		panic(domain.UnexpectedType)
	}

}

func (l *Launcher) Save(dict ConfigDict) error {
	var output []byte
	var err error
	if l.Type == domain.Alfred {
		output, err = l.AlfredSites.Encode(dict)
	} else if l.Type == domain.Albert {
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
