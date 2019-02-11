package al2

import (
	"bytes"
	"github.com/ToQoz/gopwt/assert"
	"github.com/announce/altogether/al2/helper"
	"testing"
)

func TestHandler_Perform(t *testing.T) {
	out := bytes.Buffer{}
	errOut := out
	h := &Handler{
		AlfredPath: helper.MustProjectPath("testdata/Alfred.alfredpreferences"),
		AlbertPath: helper.MustProjectPath("testdata/albert"),
		Mode: &Mode{
			DryRun:  true,
			Verbose: false,
		},
		Writer:    &out,
		ErrWriter: &errOut,
	}
	err := h.Perform()
	assert.OK(t, err == nil)
}
