package al2

import (
	"flag"
	"github.com/ToQoz/gopwt"
	"github.com/ToQoz/gopwt/assert"
	"os"
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
		BasePath: "",
	},
		&Launcher{
			Type:     Albert,
			BasePath: "",
		}}
	web := Web{Launchers: pair}
	t.Run("it works with dry-run option", func(t *testing.T) {
		err := web.Sync(Option{DtyRun: true})
		assert.OK(t, err != nil)
	})
}
