//go:build go1.17
// +build go1.17

package policy

import (
	"ermites.io/pkisauce/scramble"
)

type RPCs map[string]bool
type Services map[string]RPCs
type Hosts map[string]Services

type Policy struct {
	wc string
	ph *scramble.Hasher
	h  Hosts
}

var (
	wc = "*"
)
