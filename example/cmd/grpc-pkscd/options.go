package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func _pkscGrpcCreds(PEMca, PEMcert, PEMkey string) (credentials.TransportCredentials, error) {

	certificate, err := tls.X509KeyPair([]byte(PEMcert), []byte(PEMkey))
	if err != nil {
		return nil, err
	}

	// TODO exit the pool to be global to the file we need to load only
	// once..
	caPool := x509.NewCertPool()
	ok := caPool.AppendCertsFromPEM([]byte(PEMca))
	if !ok {
		return nil, fmt.Errorf("bleh")
	}

	tlsConfig := &tls.Config{
		ClientAuth:         tls.RequireAndVerifyClientCert,
		Certificates:       []tls.Certificate{certificate},
		ClientCAs:          caPool,
		InsecureSkipVerify: false,
		// XXX this handler is NOT necessary for GRPC, since you have the interceptor
		// but for normal sockets with single cert it's the basic client verification
		//VerifyPeerCertificate: PscVerifyPeerCertificate,
		MinVersion: tls.VersionTLS13,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
			tls.TLS_AES_128_GCM_SHA256,
			tls.TLS_AES_256_GCM_SHA384,
			tls.TLS_CHACHA20_POLY1305_SHA256},
	}

	return credentials.NewTLS(tlsConfig), nil
}

var (
	optIntercept = map[string]func() grpc.ServerOption{
		"servA": pkscGrpcOptInterceptor_servA,
		"servB": pkscGrpcOptInterceptor_servB,
		//"servC": pkscGrpcOptInterceptor_servC,
	}
	optTlsId = map[string]func() (grpc.ServerOption, error){
		"servA": pkscGrpcOptTlsServer_servA,
		"servB": pkscGrpcOptTlsServer_servB,
	}
)

func pkscGrpcOptInterceptorByName(name string) (opt grpc.ServerOption, err error) {
	f, ok := optIntercept[name]
	if !ok {
		err = fmt.Errorf("no such name: %s", name)
		return
	}

	opt = f()
	return
}

func pkscGrpcOptTlsServerByName(name string) (opt grpc.ServerOption, err error) {
	f, ok := optTlsId[name]
	if !ok {
		err = fmt.Errorf("no such name: %s", name)
		return
	}

	opt, err = f()
	return
}

/*
func pkscGrpcOptInterceptorByName(name string) (opt grpc.ServerOption, err error) {
	f, ok := interceptors[name]
	if !ok {
		err = fmt.Errorf("prout")
		return
	}
	opt = grpc.UnaryInterceptor(f)
	return
}

func pkscGrpcOptTlsServerByName(name string) (opt grpc.ServerOption, err error) {
	tlsId, ok := tlsCerts[name]
	if !ok {
		err = fmt.Errorf("prout")
		return
	}

	tlscreds, err := pkscGrpcCredsServ(tlsId.ca, tlsId.cert, tlsId.key)
	if err != nil {
		return
	}

	opt = grpc.Creds(tlscreds)
	return

}
*/
