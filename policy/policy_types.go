//go:build go1.16

package policy

import (
	"encoding/base64"

	"ermites.io/pkisauce/scramble"
)

type RPCs map[string]bool
type Services map[string]RPCs
type Hosts map[string]Services

// add a dimension to the hmac.
type WC struct {
	hwc string
	swc string
	rwc string
}

type Policy struct {
	wc string
	ph *scramble.Hasher
	h  Hosts
}

var (
	b64 = base64.RawURLEncoding
	wc  = "*"
)
