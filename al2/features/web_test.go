package al2

import (
	"flag"
	"github.com/ToQoz/gopwt"
	"github.com/ToQoz/gopwt/assert"
	"go/build"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	gopwt.Empower()
	os.Exit(m.Run())
}

func TestWeb_Sync(t *testing.T) {
	pair := &Pair{&Launcher{
		Type:     Alfred,
		BasePath: filepath.Join(fetchProjectRootPath(), "testdata/Alfred.alfredpreferences"),
	},
		&Launcher{
			Type:     Albert,
			BasePath: filepath.Join(fetchProjectRootPath(), "testdata/albert"),
		}}
	web := Web{Launchers: pair}
	t.Run("it works with dry-run option", func(t *testing.T) {
		err := web.Sync(Option{DtyRun: true})
		assert.OK(t, err == nil)
	})
}

func fetchProjectRootPath() string {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		goPath = build.Default.GOPATH
	}
	s := strings.Split(goPath, ":")
	return filepath.Join(s[len(s)-1], "src/github.com/announce/altogether")
}
