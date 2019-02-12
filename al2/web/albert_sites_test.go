package web_test

import (
	"bytes"
	"github.com/ToQoz/gopwt/assert"
	"github.com/announce/altogether/al2/web"
	"github.com/google/uuid"
	"strings"
	"testing"
)

func newAlbertConfigDict() web.ConfigDict {
	c1 := &web.SiteConfig{
		Uuid:    strings.ToUpper(uuid.New().String()),
		Enabled: true,
		Utf8:    true,
		Trigger: "foo",
		Name:    "Sample description",
		Url:     "https://foo.example.com/{query}/",
	}
	c2 := &web.SiteConfig{
		Uuid:    strings.ToUpper(uuid.New().String()),
		Enabled: true,
		Utf8:    true,
		Trigger: "bar",
		Name:    "Sample description",
		Url:     "https://bar.example.com/{query}/",
	}
	dict := make(web.ConfigDict)
	dict[c1.Id()] = c1
	dict[c2.Id()] = c2
	return dict
}

func TestAlbertSites_Encode(t *testing.T) {
	dict := newAlbertConfigDict()
	sites := web.AlbertSites{}
	json, err := sites.Encode(dict)
	assert.OK(t, json != nil)
	assert.OK(t, len(json) > 0)
	assert.OK(t, err == nil)
}

func TestAlbertSites_Decode(t *testing.T) {
	dict := newAlbertConfigDict()
	sites := web.AlbertSites{}
	json, _ := sites.Encode(dict)
	data := bytes.NewReader(json)
	err := sites.Decode(data)
	assert.OK(t, len(sites) == 2)
	assert.OK(t, err == nil)
}
