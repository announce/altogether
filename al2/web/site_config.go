package web

import (
	"bytes"
	"crypto/sha1"
	"strings"
)

type SiteConfig struct {
	Uuid     string `plist:"uuid,omitempty" json:"-"`
	Enabled  bool   `plist:"enabled" json:"-"`
	Utf8     bool   `plist:"utf8" json:"-"`
	Trigger  string `plist:"keyword" json:"trigger"`
	Name     string `plist:"text" json:"name"`
	Url      string `plist:"url" json:"url"`
	IconPath string `plist:"-" json:"iconPath"`
}

func (s *SiteConfig) Id() Id {
	a := s
	a.Normalize()
	b := bytes.Buffer{}
	b.WriteString(a.Trigger)
	b.WriteString(a.Url)
	return sha1.Sum(b.Bytes())
}

func (s *SiteConfig) PreserveUuid(key string) {
	s.Uuid = key
}

func (s *SiteConfig) Normalize() {
	s.Albert()
}

const Spacer = " "

func (s *SiteConfig) Albert() {
	s.Url = strings.Replace(s.Url, "{query}", "%s", -1)
	s.Name = strings.Replace(s.Name, "{query}", "%s", -1)
	s.Trigger = strings.Trim(s.Trigger, Spacer) + Spacer
}

func (s *SiteConfig) Alfred() {
	s.Url = strings.Replace(s.Url, "%s", "{query}", -1)
	s.Name = strings.Replace(s.Name, "%s", "{query}", -1)
	s.Trigger = strings.Trim(s.Trigger, Spacer)
}
