package main

import (
	"context"
	"fmt"

	"ermites.io/pkisauce/example/svc"
)

type RestrictedService struct {
}

// CallMe(context.Context, *Request) (*Response, error)
func (s *RestrictedService) CallHere(ctx context.Context, r *svc.Request) (rp *svc.Response, err error) {
	fmt.Printf("CallHere() request!\n")
	rp = &svc.Response{
		Ret: r.Cmd,
		Msg: r.Args,
	}
	return
}

func (s *RestrictedService) CallThere(ctx context.Context, r *svc.Request) (rp *svc.Response, err error) {
	fmt.Printf("CallThere() request: %#v\n", r)

	rp = &svc.Response{
		Ret: r.Cmd,
		Msg: r.Args,
	}
	return
}
