package web

import (
	"bytes"
	"flag"
	"github.com/ToQoz/gopwt"
	"github.com/ToQoz/gopwt/assert"
	"github.com/announce/altogether/al2/helper"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	gopwt.Empower()
	os.Exit(m.Run())
}

func TestWeb_Sync(t *testing.T) {
	out := bytes.Buffer{}
	errOut := out
	pair := &Pair{&Launcher{
		Type:     Alfred,
		BasePath: helper.MustProjectPath("testdata/Alfred.alfredpreferences"),
	},
		&Launcher{
			Type:     Albert,
			BasePath: helper.MustProjectPath("testdata/albert"),
		}}
	web := Web{
		Launchers: pair,
		Out:       &out,
		ErrOut:    &errOut,
	}
	t.Run("it works with no dry-run option", func(t *testing.T) {
		err := web.Sync(Option{DtyRun: false, Verbose: true})
		assert.OK(t, err == nil)
	})
}
