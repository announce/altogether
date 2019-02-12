package domain_test

import (
	"flag"
	"github.com/ToQoz/gopwt"
	"github.com/ToQoz/gopwt/assert"
	"github.com/announce/altogether/al2/domain"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	gopwt.Empower()
	os.Exit(m.Run())
}

func TestType_String(t *testing.T) {
	assert.OK(t, domain.Alfred.String() == "Alfred")
	assert.OK(t, domain.Albert.String() == "Albert")
}
