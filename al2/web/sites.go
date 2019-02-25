package web

import "io"

type Sites interface {
	List() []*SiteConfig
	Decode(io.ReadSeeker) error
	Encode(ConfigDict) ([]byte, error)
	Populate(ConfigDict)
	Convert(ConfigDict) Sites
}
