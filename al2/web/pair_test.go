package web

import (
	"github.com/ToQoz/gopwt/assert"
	"github.com/announce/altogether/al2/domain"
	"github.com/announce/altogether/al2/helper"
	"log"
	"testing"
	"time"
)

func newPair() (*Pair, error) {
	pair := &Pair{
		&Launcher{
			Type:     domain.Alfred,
			BasePath: helper.EnsureDataPath(domain.Alfred, Config),
		},
		&Launcher{
			Type:     domain.Albert,
			BasePath: helper.EnsureDataPath(domain.Albert, Config),
		}}
	return pair, pair.Load()
}

func TestPair_Load(t *testing.T) {
	pair, err := newPair()
	assert.OK(t, pair != nil)
	assert.OK(t, err == nil)
}

func TestPair_Merge(t *testing.T) {
	pair, _ := newPair()
	path0 := pair[0].ConfigPath
	path1 := pair[1].ConfigPath
	t.Run("touch path0 first and path1 later", func(t *testing.T) {
		helper.MustTouchFile(path0, 0)
		helper.MustTouchFile(path1, 1*time.Nanosecond)
		_ = pair.Load()
		pair.Merge(make(ConfigDict))
		assert.OK(t, pair[0].ConfigPath == path0)
		assert.OK(t, pair[1].ConfigPath == path1)
	})
	t.Run("touch path1 first and path0 later", func(t *testing.T) {
		log.Printf("(p1, p1)=(%v,%v)(%v, %v)",
			pair[0].Type, pair[1].Type, pair[0].Mtime(), pair[1].Mtime())
		helper.MustTouchFile(path1, 0)
		helper.MustTouchFile(path0, 1*time.Nanosecond)
		_ = pair.Load()
		pair.Merge(make(ConfigDict))
		log.Printf("(p1, p1)=(%v,%v)(%v, %v)",
			pair[0].Type, pair[1].Type, pair[0].Mtime(), pair[1].Mtime())
		assert.OK(t, pair[1].ConfigPath == path0)
		assert.OK(t, pair[0].ConfigPath == path1)
	})
}

func TestPair_Save(t *testing.T) {
	defer helper.MustRemoveTmpDir()
	pair, err := newPair()
	assert.OK(t, pair != nil)
	assert.OK(t, err == nil)
}
