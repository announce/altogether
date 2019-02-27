package web

import (
	"fmt"
	"github.com/announce/altogether/al2/domain"
)

type Order []domain.Type
type Diff map[Id]Order

func (o Order) String() string {
	s := ""
	for i, v := range o {
		if i < 1 {
			s += fmt.Sprintf("%s", v)
		} else {
			s += fmt.Sprintf(" < %s", v)
		}
	}
	return s
}
