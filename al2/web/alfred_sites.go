package web

import (
	"github.com/google/uuid"
)

type CustomSites map[string]*SiteConfig
type AlfredSites struct {
	CustomSites `plist:"customSites"`
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
