package web

import (
	"bytes"
	"github.com/DHowett/go-plist"
	"github.com/google/uuid"
	"io"
)

type CustomSites map[string]*SiteConfig
type AlfredSites struct {
	CustomSites `plist:"customSites"`
	list        []*SiteConfig
}

func (a *AlfredSites) List() []*SiteConfig {
	for _, v := range a.CustomSites {
		a.list = append(a.list, v)
	}
	return a.list
}

func (a *AlfredSites) Populate(dict ConfigDict) {
	for k, v := range a.CustomSites {
		v.PreserveUuid(k)
		dict[v.Id()] = v
	}
}

func (a *AlfredSites) Decode(r io.ReadSeeker) error {
	decoder := plist.NewDecoder(r)
	return decoder.Decode(a)
}

func (a *AlfredSites) Encode(dict ConfigDict) ([]byte, error) {
	data := &bytes.Buffer{}
	encoder := plist.NewEncoder(data)
	encoder.Indent("\t")
	if err := encoder.Encode(a.Convert(dict)); err != nil {
		return nil, err
	}
	return data.Bytes(), nil
}

func (a *AlfredSites) Convert(dict ConfigDict) Sites {
	sites := make(CustomSites)
	for _, site := range dict {
		config := site
		if config.Uuid == "" {
			config.Uuid = uuid.New().String()
		}
		config.Alfred()
		sites[config.Uuid] = config
	}
	return &AlfredSites{CustomSites: sites}
}
