package main

import (
	"context"
	"fmt"

	"ermites.io/pkisauce/example/svc"
)

type PublicService struct {
}

// CallMe(context.Context, *Request) (*Response, error)
func (s *PublicService) CallMe(ctx context.Context, r *svc.Request) (rp *svc.Response, err error) {
	fmt.Printf("CallMe() request!\n")
	rp = &svc.Response{
		Ret: r.Cmd,
		Msg: r.Args,
	}
	return
}

func (s *PublicService) Ping(ctx context.Context, r *svc.Request) (rp *svc.Response, err error) {
	fmt.Printf("Ping() request: %#v\n", r)

	rp = &svc.Response{
		Ret: r.Cmd,
		Msg: r.Args,
	}
	return
}
