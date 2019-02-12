package web

import (
	"encoding/hex"
	"github.com/ToQoz/gopwt/assert"
	"github.com/announce/altogether/al2/domain"
	"github.com/google/uuid"
	"strings"
	"testing"
)

func newSiteConfig(p domain.Type) *SiteConfig {
	config := &SiteConfig{}
	switch p {
	case domain.Alfred:
		{
			config = &SiteConfig{
				Uuid:    strings.ToUpper(uuid.New().String()),
				Enabled: true,
				Utf8:    true,
				Trigger: "jj",
				Name:    "Sample description",
				Url:     "https://example.com/{query}/",
			}
		}
	case domain.Albert:
		{
			config = &SiteConfig{
				Trigger: "jj ",
				Name:    "Sample description",
				Url:     "https://example.com/%s/",
			}
		}
	}
	return config
}

func TestSiteConfig_Id(t *testing.T) {
	id := newSiteConfig(domain.Alfred).Id()
	assert.OK(t, len(hex.EncodeToString(id[:])) == 40)
}

func TestSiteConfig_PreserveUuid(t *testing.T) {
	config := newSiteConfig(domain.Alfred)
	config.PreserveUuid("")
	assert.OK(t, config != nil)
}

func TestSiteConfig_Normalize(t *testing.T) {
	t.Run("Alfred", func(t *testing.T) {
		alfred := newSiteConfig(domain.Alfred)
		albert := newSiteConfig(domain.Albert)
		alfred.Normalize()
		assert.OK(t, alfred.Trigger == albert.Trigger)
		assert.OK(t, alfred.Url == albert.Url)
	})
	t.Run("Albert", func(t *testing.T) {
		alfred := newSiteConfig(domain.Alfred)
		albert := newSiteConfig(domain.Albert)
		albert.Normalize()
		assert.OK(t, alfred.Trigger != albert.Trigger)
		assert.OK(t, alfred.Url != albert.Url)
	})
}

func TestSiteConfig_Alfred(t *testing.T) {
	alfred := newSiteConfig(domain.Alfred)
	albert := newSiteConfig(domain.Albert)
	albert.Alfred()
	assert.OK(t, alfred.Trigger == albert.Trigger)
	assert.OK(t, alfred.Url == albert.Url)
}

func TestSiteConfig_Albert(t *testing.T) {
	alfred := newSiteConfig(domain.Alfred)
	albert := newSiteConfig(domain.Albert)
	alfred.Albert()
	assert.OK(t, alfred.Trigger == albert.Trigger)
	assert.OK(t, alfred.Url == albert.Url)
}
