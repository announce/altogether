package al2_test

import (
	"bytes"
	"flag"
	"github.com/ToQoz/gopwt/assert"
	"github.com/announce/altogether/al2"
	"github.com/announce/altogether/al2/helper"
	"github.com/urfave/cli"
	"testing"
)

func newCliContext(params []string) *cli.Context {
	set := flag.NewFlagSet("altogether", 0)
	_ = set.Parse(params)
	out := bytes.Buffer{}
	errOut := out
	c := cli.NewContext(&cli.App{
		Writer:    &out,
		ErrWriter: &errOut,
		Commands:  al2.Commands,
	}, set, nil)
	return c
}

func TestHandler_SyncWeb(t *testing.T) {
	defer helper.MustRemoveTmpDir()
	h := &al2.Handler{}
	t.Run("no enough params", func(t *testing.T) {
		err := h.SyncWeb(newCliContext([]string{"sync-web"}))
		assert.OK(t, err != nil)
	})
	t.Run("valid params", func(t *testing.T) {
		t.Skip()
	})
}
