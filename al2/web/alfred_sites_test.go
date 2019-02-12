package web

import (
	"bytes"
	"github.com/ToQoz/gopwt/assert"
	"github.com/google/uuid"
	"strings"
	"testing"
)

func newAlfredConfigDict() ConfigDict {
	c1 := &SiteConfig{
		Uuid:    strings.ToUpper(uuid.New().String()),
		Enabled: true,
		Utf8:    true,
		Trigger: "foo ",
		Name:    "Sample description",
		Url:     "https://foo.example.com/%s/",
	}
	c2 := &SiteConfig{
		Trigger: "bar ",
		Name:    "Sample description",
		Url:     "https://bar.example.com/%s/",
	}
	dict := make(ConfigDict)
	dict[c1.Id()] = c1
	dict[c2.Id()] = c2
	return dict
}

func TestAlfredSites_Encode(t *testing.T) {
	dict := newAlfredConfigDict()
	sites := &AlfredSites{}
	json, err := sites.Encode(dict)
	assert.OK(t, json != nil)
	assert.OK(t, len(json) > 0)
	assert.OK(t, err == nil)
}

func TestAlfredSites_Decode(t *testing.T) {
	dict := newAlfredConfigDict()
	sites := &AlfredSites{}
	plist, _ := sites.Encode(dict)
	data := bytes.NewReader(plist)
	err := sites.Decode(data)
	assert.OK(t, len(sites.CustomSites) == 2)
	assert.OK(t, err == nil)
}
