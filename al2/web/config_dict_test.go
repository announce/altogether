package web_test

import (
	"crypto/sha1"
	"fmt"
	"github.com/ToQoz/gopwt/assert"
	"github.com/announce/altogether/al2/web"
	"strings"
	"testing"
)

func TestConfigDict_String(t *testing.T) {
	dict := make(web.ConfigDict)
	assert.OK(t, len(fmt.Sprintf("%s", dict)) == 0)
	dict[sha1.Sum([]byte("a"))] = &web.SiteConfig{}
	assert.OK(t, strings.Count(dict.String(), "\n") == 0)
	dict[sha1.Sum([]byte("b"))] = &web.SiteConfig{}
	assert.OK(t, strings.Count(dict.String(), "\n") == 1)
}
