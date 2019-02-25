package web

import (
	"fmt"
	"sort"
	"strings"
)

type Pair [2]*Launcher

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

func (p *Pair) Len() int {
	return len(p)
}

func (p *Pair) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *Pair) Less(i, j int) bool {
	return p[i].Mtime().UnixNano() < p[j].Mtime().UnixNano()
}

func (p *Pair) Merge(dict ConfigDict) {
	sort.Sort(p)
	for _, launcher := range p {
		launcher.Populate(dict)
	}
}

func (p *Pair) Diff(dict ConfigDict) string {
	sort.Sort(p)
	diff := make(Diff)
	for _, launcher := range p {
		launcher.Diff(diff)
	}
	var s []string
	for id, config := range dict {
		s = append(s, fmt.Sprintf("%s\t%#v", diff[id], config))
	}
	return strings.Join(s, "\n")
}

func (p *Pair) Save(dict ConfigDict) error {
	for _, launcher := range p {
		if err := launcher.Save(dict); err != nil {
			return err
		}
	}
	return nil
}
