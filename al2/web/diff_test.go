package web_test

import (
	"github.com/ToQoz/gopwt/assert"
	"github.com/announce/altogether/al2/domain"
	"github.com/announce/altogether/al2/web"
	"testing"
)

func TestOrder_String(t *testing.T) {
	t.Run("1 element", func(t *testing.T) {
		order := web.Order{domain.Alfred}
		assert.OK(t, order.String() == "Alfred")
	})
	t.Run("2 elements", func(t *testing.T) {
		order := web.Order{domain.Alfred, domain.Albert}
		assert.OK(t, order.String() == "Alfred < Albert")
	})
}
