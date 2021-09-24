//go:build go1.16
// +build go1.16

package config

var (
	FlavorKeyword = map[string]int{
		"go-tls":  FlavorGoTLS,  // this is just VerifyPeerCertificates + CN check.
		"go-grpc": FlavorGoGRPC, // this is GRPC + TLS interceptor
		"pem":     FlavorPEM,    // CA signature is the only check you have here.
		//"java-grpc":     FlavorJavaGrpc,      // CA signature is the only check you have here.
		//"java-tls":     FlavorJavaSocket,      // CA signature is the only check you have here.
	}

	wc = "*"
)
