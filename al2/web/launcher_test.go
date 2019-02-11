package web

import (
	"github.com/ToQoz/gopwt/assert"
	"github.com/announce/altogether/al2/domain"
	"github.com/announce/altogether/al2/helper"
	"testing"
)

func newAlfred() *Launcher {
	launcher := &Launcher{
		Type:     domain.Alfred,
		BasePath: helper.EnsureDataPath(domain.Alfred, Config),
	}
	launcher.Init()
	return launcher
}

func newAlbert() *Launcher {
	launcher := &Launcher{
		Type:     domain.Albert,
		BasePath: helper.EnsureDataPath(domain.Albert, Config),
	}
	launcher.Init()
	return launcher
}

func TestLauncher_Init(t *testing.T) {
	launcher := newAlfred()
	assert.OK(t, launcher != nil)
}

func TestLauncher_Mtime(t *testing.T) {
	launcher := newAlfred()
	err := launcher.Load()
	assert.OK(t, err == nil)
	mtime := launcher.Mtime().String()
	assert.OK(t, len(mtime) > 0)
}

func TestLauncher_Parse(t *testing.T) {
	defer helper.MustRemoveTmpDir()
	t.Run("it works with Alfred", func(t *testing.T) {
		launcher := newAlfred()
		err := launcher.Parse()
		assert.OK(t, err == nil)
		assert.OK(t, launcher.AlfredSites != nil)
	})
	t.Run("it works with Albert", func(t *testing.T) {
		launcher := newAlbert()
		err := launcher.Parse()
		assert.OK(t, err == nil)
		assert.OK(t, launcher.AlbertSites != nil)
	})
}

func TestLauncher_Populate(t *testing.T) {
	defer helper.MustRemoveTmpDir()
	t.Run("it works with Alfred", func(t *testing.T) {
		dict := make(ConfigDict)
		launcher := newAlfred()
		_ = launcher.Parse()
		launcher.Populate(dict)
		assert.OK(t, len(dict) > 0)
	})
	t.Run("it works with Albert", func(t *testing.T) {
		dict := make(ConfigDict)
		launcher := newAlbert()
		_ = launcher.Parse()
		launcher.Populate(dict)
		assert.OK(t, len(dict) > 0)
	})
}

func TestLauncher_Save(t *testing.T) {
	defer helper.MustRemoveTmpDir()
	t.Run("it works with Alfred", func(t *testing.T) {
		dict := make(ConfigDict)
		launcher := newAlfred()
		_ = launcher.Load()
		_ = launcher.Parse()
		launcher.Populate(dict)
		err := launcher.Save(dict)
		assert.OK(t, err == nil)
	})
	t.Run("it works with Albert", func(t *testing.T) {
		dict := make(ConfigDict)
		launcher := newAlbert()
		_ = launcher.Load()
		_ = launcher.Parse()
		launcher.Populate(dict)
		err := launcher.Save(dict)
		assert.OK(t, err == nil)
	})
}
