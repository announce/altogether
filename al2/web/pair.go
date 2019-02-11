package web

import (
	"sort"
)

type Pair [2]*Launcher

func (p *Pair) Len() int {
	return len(p)
}

func (p *Pair) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *Pair) Less(i, j int) bool {
	return p[i].Mtime().Unix() < p[j].Mtime().Unix()
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
