package domain_test

import (
	"github.com/ToQoz/gopwt/assert"
	"github.com/announce/altogether/al2/domain"
	"testing"
)

func newConfigPath() domain.ConfigPath {
	return domain.ConfigPath{
		domain.Alfred: "preferences/features/foo/prefs.plist",
		domain.Albert: "org.albert.extension.foo/engines.json",
	}
}

func TestConfigPath_Path(t *testing.T) {
	c := newConfigPath()
	assert.OK(t, c.Path(domain.Alfred) == c[domain.Alfred])
	assert.OK(t, c.Path(domain.Albert) == c[domain.Albert])
}
