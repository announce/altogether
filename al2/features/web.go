package al2

import (
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
	Type      Type
	BasePath  string
	PlistPath string
	FileInfo  os.FileInfo
}

type Pair [2]Launcher
type Web struct {
	Launchers *Pair
}

// @TODO os.Chtimes to avoid loop
func (w *Web) Sync() error {
	w.init()
	_, err := w.direction()
	if err != nil {
		return err
	}
	return err
}

func (w *Web) init() {
	for _, launcher := range w.Launchers {
		launcher.PlistPath = filepath.Join(launcher.BasePath, ConfigPath[launcher.Type])
	}
}

func (w *Web) direction() (*Pair, error) {
	for _, launcher := range w.Launchers {
		info, err := os.Stat(launcher.PlistPath)
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
