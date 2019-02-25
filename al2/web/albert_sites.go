package web

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
)

type AlbertSites []*SiteConfig

func (a *AlbertSites) List() []*SiteConfig {
	return *a
}

func (a *AlbertSites) Populate(dict ConfigDict) {
	for _, v := range *a {
		dict[v.Id()] = v
	}
}

func (a *AlbertSites) Decode(r io.ReadSeeker) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &a)
}

func (a *AlbertSites) Encode(dict ConfigDict) ([]byte, error) {
	j, err := json.MarshalIndent(
		a.Convert(dict), "", strings.Repeat(" ", 2))
	if err != nil {
		return nil, err
	}
	return j, err
}

func (a *AlbertSites) Convert(dict ConfigDict) Sites {
	var sites AlbertSites
	for _, site := range dict {
		config := site
		config.Albert()
		sites = append(sites, config)
	}
	return &sites
}
