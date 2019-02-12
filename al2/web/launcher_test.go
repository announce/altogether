package web_test

import (
	"github.com/ToQoz/gopwt/assert"
	"github.com/announce/altogether/al2/domain"
	"github.com/announce/altogether/al2/helper"
	"github.com/announce/altogether/al2/web"
	"testing"
)

func newLauncher(p domain.Type) *web.Launcher {
	launcher := &web.Launcher{
		Type:     p,
		BasePath: helper.EnsureDataPath(p, web.Config),
	}
	launcher.Init()
	return launcher
}

func TestLauncher_Init(t *testing.T) {
	launcher := newLauncher(domain.Alfred)
	assert.OK(t, launcher != nil)
}

func TestLauncher_Load(t *testing.T) {
	err := newLauncher(domain.Alfred).Load()
	assert.OK(t, err == nil)
}

func TestLauncher_Mtime(t *testing.T) {
	launcher := newLauncher(domain.Alfred)
	err := launcher.Load()
	assert.OK(t, err == nil)
	mtime := launcher.Mtime().String()
	assert.OK(t, len(mtime) > 0)
}

func TestLauncher_Parse(t *testing.T) {
	defer helper.MustRemoveTmpDir()
	t.Run("it works with Alfred", func(t *testing.T) {
		launcher := newLauncher(domain.Alfred)
		err := launcher.Parse()
		assert.OK(t, err == nil)
		assert.OK(t, launcher.AlfredSites != nil)
	})
	t.Run("it works with Albert", func(t *testing.T) {
		launcher := newLauncher(domain.Albert)
		err := launcher.Parse()
		assert.OK(t, err == nil)
		assert.OK(t, launcher.AlbertSites != nil)
	})
}

func TestLauncher_Populate(t *testing.T) {
	defer helper.MustRemoveTmpDir()
	t.Run("it works with Alfred", func(t *testing.T) {
		dict := make(web.ConfigDict)
		launcher := newLauncher(domain.Alfred)
		_ = launcher.Parse()
		launcher.Populate(dict)
		assert.OK(t, len(dict) > 0)
	})
	t.Run("it works with Albert", func(t *testing.T) {
		dict := make(web.ConfigDict)
		launcher := newLauncher(domain.Albert)
		_ = launcher.Parse()
		launcher.Populate(dict)
		assert.OK(t, len(dict) > 0)
	})
}

func TestLauncher_Save(t *testing.T) {
	defer helper.MustRemoveTmpDir()
	t.Run("it works with Alfred", func(t *testing.T) {
		dict := make(web.ConfigDict)
		launcher := newLauncher(domain.Alfred)
		_ = launcher.Load()
		_ = launcher.Parse()
		launcher.Populate(dict)
		err := launcher.Save(dict)
		assert.OK(t, err == nil)
	})
	t.Run("it works with Albert", func(t *testing.T) {
		dict := make(web.ConfigDict)
		launcher := newLauncher(domain.Albert)
		_ = launcher.Load()
		_ = launcher.Parse()
		launcher.Populate(dict)
		err := launcher.Save(dict)
		assert.OK(t, err == nil)
	})
}
