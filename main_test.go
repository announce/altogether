package main

import (
	"github.com/ToQoz/gopwt/assert"
	"testing"
)

func TestCreateApp(t *testing.T) {
	app := CreateApp()
	assert.OK(t, app != nil)
}
