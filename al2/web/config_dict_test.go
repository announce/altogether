package web

import (
	"crypto/sha1"
	"fmt"
	"github.com/ToQoz/gopwt/assert"
	"strings"
	"testing"
)

func TestConfigDict_String(t *testing.T) {
	dict := make(ConfigDict)
	assert.OK(t, len(fmt.Sprintf("%s", dict)) == 0)
	dict[sha1.Sum([]byte("a"))] = &SiteConfig{}
	assert.OK(t, strings.Count(dict.String(), "\n") == 0)
	dict[sha1.Sum([]byte("b"))] = &SiteConfig{}
	assert.OK(t, strings.Count(dict.String(), "\n") == 1)
}
