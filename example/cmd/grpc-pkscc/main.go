//go:generate protoc -I. --go-grpc_out=require_unimplemented_servers=false:.  --go_out=. --go_opt=paths=source_relative svc.proto

package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"ermites.io/pkisauce/example/svc"
)

func main() {
	fmt.Printf("grpc-client\n")
	var dialopts []grpc.DialOption

	// create dialopts depending on server destination
	//dialopts = append(dialopts, grpc.WithInsecure())
	opt, err := pkscGrpcOptTlsClient_client3.DialOptByName("servA")
	if err != nil {
		panic(err)
	}
	dialopts = append(dialopts, opt)

	ctx := context.Background()

	nfd, err := grpc.Dial("127.0.0.1:31337", dialopts...)
	if err != nil {
		panic(err)
	}

	public := svc.NewPublicClient(nfd)
	restricted := svc.NewRestrictedClient(nfd)

	r := svc.Request{
		Cmd:  1,
		Args: "prout",
	}
	rp, err := public.CallMe(ctx, &r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("rp: %#v\n", rp)

	pr := svc.Request{
		Cmd:  1,
		Args: "plop",
	}
	rpp, err := restricted.CallThere(ctx, &pr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("rpp: %#v\n", rpp)

}
