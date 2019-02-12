package web_test

import (
	"bytes"
	"flag"
	"github.com/ToQoz/gopwt"
	"github.com/ToQoz/gopwt/assert"
	"github.com/announce/altogether/al2/helper"
	"github.com/announce/altogether/al2/web"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	gopwt.Empower()
	defer helper.MustRemoveTmpDir()
	os.Exit(m.Run())
}

func TestWeb_Sync(t *testing.T) {
	defer helper.MustRemoveTmpDir()
	out := bytes.Buffer{}
	errOut := out
	pair, _ := newPair()
	w := web.Web{
		Launchers: pair,
		Out:       &out,
		ErrOut:    &errOut,
	}
	t.Run("it works with no dry-run option", func(t *testing.T) {
		err := w.Sync(web.Option{DtyRun: false, Verbose: false})
		assert.OK(t, err == nil)
		assert.OK(t, w.ConfigDict != nil)
	})
	t.Run("it works with dry-run option", func(t *testing.T) {
		err := w.Sync(web.Option{DtyRun: true, Verbose: true})
		assert.OK(t, err == nil)
		assert.OK(t, strings.Count(out.String(), "\n") == len(w.ConfigDict))
	})
}
