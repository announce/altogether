package web

import (
	"bytes"
	"github.com/DHowett/go-plist"
	"github.com/google/uuid"
	"os"
)

type CustomSites map[string]*SiteConfig
type AlfredSites struct {
	CustomSites `plist:"customSites"`
}

func (a *AlfredSites) Decode(file *os.File) error {
	decoder := plist.NewDecoder(file)
	if err := decoder.Decode(a); err != nil {
		return err
	}
	return nil
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
