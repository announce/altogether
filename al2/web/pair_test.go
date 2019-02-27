package web_test

import (
	"github.com/ToQoz/gopwt/assert"
	"github.com/announce/altogether/al2/domain"
	"github.com/announce/altogether/al2/helper"
	"github.com/announce/altogether/al2/web"
	"testing"
	"time"
)

func newPair() (*web.Pair, error) {
	pair := &web.Pair{
		&web.Launcher{
			Type:     domain.Alfred,
			BasePath: helper.EnsureDataPath(domain.Alfred, web.Config),
		},
		&web.Launcher{
			Type:     domain.Albert,
			BasePath: helper.EnsureDataPath(domain.Albert, web.Config),
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
		pair.Merge(make(web.ConfigDict))
		assert.OK(t, pair[0].ConfigPath == path0)
		assert.OK(t, pair[1].ConfigPath == path1)
	})
	t.Run("touch path1 first and path0 later", func(t *testing.T) {
		//log.Printf("(p1, p1)=(%v,%v)(%v, %v)",
		//	pair[0].Type, pair[1].Type, pair[0].Mtime(), pair[1].Mtime())
		helper.MustTouchFile(path1, 0)
		helper.MustTouchFile(path0, 1*time.Nanosecond)
		_ = pair.Load()
		pair.Merge(make(web.ConfigDict))
		//log.Printf("(p1, p1)=(%v,%v)(%v, %v)",
		//	pair[0].Type, pair[1].Type, pair[0].Mtime(), pair[1].Mtime())
		assert.OK(t, pair[1].ConfigPath == path0)
		assert.OK(t, pair[0].ConfigPath == path1)
	})
}

func TestPair_Diff(t *testing.T) {
	defer helper.MustRemoveTmpDir()
	pair, err := newPair()
	if err != nil {
		panic(err)
	}
	dict := make(web.ConfigDict)
	assert.OK(t, pair.Save(dict) == nil)
}

func TestPair_Save(t *testing.T) {
	defer helper.MustRemoveTmpDir()
	pair, err := newPair()
	if err != nil {
		panic(err)
	}
	dict := web.ConfigDict{}
	pair.Merge(dict)
	assert.OK(t, pair.Diff(dict) != "")
}
