package web

type AlbertSites []*SiteConfig

func (a *AlbertSites) Convert(dict ConfigDict) AlbertSites {
	var sites AlbertSites
	for _, site := range dict {
		config := site
		config.Albert()
		sites = append(sites, config)
	}
	return sites
}
