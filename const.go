//go:build go1.17
// +build go1.17

package main

const (
	Defaultx509version   = 0x3
	defaultPEMValidity   = 3650 // days
	defaultPEMHeaderCert = "CERTIFICATE"
	defaultPEMHeaderKey  = "PRIVATE KEY"

	PkiSauceVersion = "v0.1.0"
)

const (
	PurposeClientCert int = 1 << iota
	PurposeServerCert
)

const (
	TemplateNameDebug     = "templates/debug.template"
	TemplateHelpersCommon = "templates/helpers_common.go.template"
	TemplateHelpers       = "templates/helpers.go.template"
	TemplateCAPEM         = "templates/pem.ca.template"
	TemplateClientPEM     = "templates/pem.client.template"
	TemplateServerPEM     = "templates/pem.server.template"
	/*
		TemplateScrambleHash    = "templates/helpers_scramble_hash.go.template"
		TemplateScrambleMarshal = "templates/helpers_scramble_marshal.go.template"
		TemplateScrambleError   = "templates/helpers_scramble_error.go.template"
	*/
)
