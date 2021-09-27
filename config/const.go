//go:build go1.17
// +build go1.17

package config

const (
	FlavorNone = 0
)

const (
	FlavorGoTLS int = 1 << iota
	FlavorGoGRPC
	FlavorPEM
	FlavorJavaGRPC
	FlavorJavaTLS
	FlavorCppGRPC
	FlavorCppSocketOSSL
	FlavorCSocketOSSL
)
