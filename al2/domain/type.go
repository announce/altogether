package domain

import (
	"errors"
)

var UnexpectedType = errors.New("unexpected type")

type Type int

const (
	Alfred Type = iota
	Albert
)

func (t Type) String() string {
	return [...]string{
		"Alfred",
		"Albert",
	}[t]
}
