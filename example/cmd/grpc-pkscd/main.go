package main

// the grpc test daemon
import (
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"

	"ermites.io/pkisauce/example/svc"
)

type myServiceContext struct {
	public    *PublicService
	restrict  *RestrictedService
	different *DifferentService
}

func registerByName(name string) {
}

func main() {
	var opts []grpc.ServerOption

	fmt.Printf("%s\n", os.Args[0])

	listenAddr := "127.0.0.1:31337"

	nfd, err := net.Listen("tcp", listenAddr)
	if err != nil {
		panic(err)
	}

	server := os.Args[1]
	fmt.Printf("listening server: %s\n", server)

	// to compile, run pkisauce with -test-server <path> -test-client <path> <pkisaucefiles>
	// test-server will implement these calls +
	// the one defined in the pkisauce config file.
	//
	// all server profile will be available
	// ./pksc-serv [-p <port>] <server name>
	//
	// like:
	// ./pksc-client <server name>
	//

	//map["fsd"]pkiSauceInterceptorFsd
	opt, err := pkscGrpcOptInterceptorByName(server)
	if err != nil {
		panic(err)
	}
	opts = append(opts, opt)

	opt, err = pkscGrpcOptTlsServerByName(server)
	if err != nil {
		panic(err)
	}
	opts = append(opts, opt)

	fmt.Printf("policy:\n\n%#v\n\n", pkscPolicy_servA)

	grpcd := grpc.NewServer(opts...)

	sctx := myServiceContext{
		public:    &PublicService{},
		restrict:  &RestrictedService{},
		different: &DifferentService{},
	}

	fmt.Printf("registering >> public << service\n")
	svc.RegisterPublicServer(grpcd, sctx.public)
	fmt.Printf("registering >> restricted << service\n")
	svc.RegisterRestrictedServer(grpcd, sctx.restrict)
	fmt.Printf("registering >> different << service\n")
	svc.RegisterDifferentServer(grpcd, sctx.different)

	fmt.Printf("listening on %s\n", listenAddr)
	err = grpcd.Serve(nfd)
	if err != nil {
		panic(err)
	}
}
