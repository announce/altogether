package web

import (
	"crypto/sha1"
	"fmt"
	"strings"
)

type Id [sha1.Size]byte
type ConfigDict map[Id]*SiteConfig

func (d ConfigDict) String() string {
	var s []string
	for _, config := range d {
		s = append(s, fmt.Sprintf("%#v", config))
	}
	return strings.Join(s, "\n")
}
