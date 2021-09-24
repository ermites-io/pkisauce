package main

import (
	"context"
	"fmt"

	"ermites.io/pkisauce/example/svc"
)

type DifferentService struct {
}

// CallMe(context.Context, *Request) (*Response, error)
func (s *DifferentService) CallDifferent(ctx context.Context, r *svc.Request) (rp *svc.Response, err error) {
	fmt.Printf("CallDifferent() request!\n")
	rp = &svc.Response{
		Ret: r.Cmd,
		Msg: r.Args,
	}
	return
}

func (s *DifferentService) Shazz(ctx context.Context, r *svc.Request) (rp *svc.Response, err error) {
	fmt.Printf("CallThere() request: %#v\n", r)

	rp = &svc.Response{
		Ret: r.Cmd,
		Msg: r.Args,
	}
	return
}
