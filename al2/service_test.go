package al2

import (
	"bytes"
	"github.com/ToQoz/gopwt/assert"
	"github.com/announce/altogether/al2/domain"
	"github.com/announce/altogether/al2/helper"
	"github.com/announce/altogether/al2/web"
	"testing"
)

func TestHandler_Perform(t *testing.T) {
	defer helper.MustRemoveTmpDir()
	out := bytes.Buffer{}
	errOut := out
	h := &Handler{
		AlfredPath: helper.EnsureDataPath(domain.Alfred, web.Config),
		AlbertPath: helper.EnsureDataPath(domain.Albert, web.Config),
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
