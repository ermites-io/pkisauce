//go:build go1.17
// +build go1.17

package main

import (
	"embed"
	"encoding/base64"
)

var (
	// COMMONs
	defaultCertCountry       = []string{"CH"}
	defaultCertOrg           = []string{"PkiSauce BuildTime Ephemeral CA", PkiSauceVersion}
	defaultCertOrgUnit       = []string{"PkiSauce Cert Org Unit"}
	defaultCertLocality      = []string{"Lausanne"}
	defaultCertProvince      = []string{"Lausanne"}
	defaultCertStreetAddress = []string{""}
	defaultCertPostalCode    = []string{"1337"}
	defaultEmailContacts     = []string{"pkisauce@ermites.io"}

	// CA specific
	defaultCaCommonName  = string("ephemeral.ca.ermites.io")
	defaultCaCertOrgUnit = []string{"ca org unit"}

	// Server specific
	defaultServerCertOrgUnit = []string{"server cert org"}
	defaultServerCommonName  = string("server.cn.ermites.io")

	// Client specific
	defaultClientCertOrgUnit = []string{"client cert org"}
	defaultClientCommonName  = string("client.cn.ermites.io")

	pkscb64 = base64.RawURLEncoding
)

var (
	BoolKeyword = map[string]bool{
		"y":     true,
		"n":     false,
		"yes":   true,
		"no":    false,
		"true":  true,
		"false": false,
	}

	//go:embed templates
	tFS embed.FS
)
