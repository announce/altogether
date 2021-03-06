package al2_test

import (
	"flag"
	"github.com/ToQoz/gopwt"
	"github.com/ToQoz/gopwt/assert"
	"github.com/announce/altogether/al2"
	"github.com/announce/altogether/al2/helper"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	gopwt.Empower()
	defer helper.MustRemoveTmpDir()
	os.Exit(m.Run())
}

func TestCommands(t *testing.T) {
	assert.OK(t, al2.Commands != nil)
}
