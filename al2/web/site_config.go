package web

import (
	"bytes"
	"crypto/sha1"
	"strings"
)

type Id [sha1.Size]byte
type ConfigDict map[Id]*SiteConfig

type SiteConfig struct {
	Uuid     string `plist:"uuid,omitempty" json:"-"`
	Enabled  bool   `plist:"enabled" json:"-"`
	Utf8     bool   `plist:"utf8" json:"-"`
	Trigger  string `plist:"keyword" json:"trigger"`
	Name     string `plist:"text" json:"name"`
	Url      string `plist:"url" json:"url"`
	IconPath string `plist:"-" json:"iconPath"`
}

func (a *SiteConfig) Id() Id {
	a.Normalize()
	b := bytes.Buffer{}
	b.WriteString(a.Trigger)
	b.WriteString(a.Url)
	return sha1.Sum(b.Bytes())
}

func (a *SiteConfig) PreserveUuid(key string) {
	a.Uuid = key
}

func (a *SiteConfig) Normalize() {
	a.Albert()
}

const Spacer = " "

func (a *SiteConfig) Albert() {
	a.Url = strings.Replace(a.Url, "{query}", "%s", -1)
	a.Name = strings.Replace(a.Name, "{query}", "%s", -1)
	a.Trigger = strings.Trim(a.Trigger, Spacer) + Spacer
}

func (a *SiteConfig) Alfred() {
	a.Url = strings.Replace(a.Url, "%s", "{query}", -1)
	a.Name = strings.Replace(a.Name, "%s", "{query}", -1)
	a.Trigger = strings.Trim(a.Trigger, Spacer)
}
