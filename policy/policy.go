//go:build go1.16

package policy

import (
	"ermites.io/pkisauce/scramble"
)

func New(ph *scramble.Hasher, wc string) *Policy {
	lwc := wc
	if ph != nil {
		lwc = ph.Sum64([]byte(wc))
	}

	p := &Policy{
		wc: lwc,
		ph: ph,
		h:  NewHosts(),
	}
	return p
}

func Unmarshal(key []byte, blob64 string) (h Hosts, err error) {
	// TODO: this is in the template actually.
	return
}

func (p *Policy) Hosts() Hosts {
	return p.h
}

func (p *Policy) Check(name, service, call string) bool {
	lcall := call
	lsvc := service
	lname := name

	if p.ph != nil {
		lname = p.ph.Sum64([]byte(name))
		lsvc = p.ph.Sum64([]byte(service))
		lcall = p.ph.Sum64([]byte(call))
	}

	return p.h.check(p.wc, lname, lsvc, lcall)
}

func (p *Policy) Entry(name string, service string, call string) {
	lname := name
	lsvc := service
	lcall := call

	if p.ph != nil {
		// add anothet dimension so that patterns somehow disappears while still matcheable
		// like lname := p.ph.Sum([]byte("h:"+name))
		// like lsvc := p.ph.Sum([]byte("s:"+name))
		// like lcall := p.ph.Sum([]byte("r:"+name))
		lname = p.ph.Sum64([]byte(name))
		lsvc = p.ph.Sum64([]byte(service))
		lcall = p.ph.Sum64([]byte(call))
	}
	//fmt.Printf("wc: %.6s add name:%s (%s) svc:%s (%s) rpc:%s (%s)\n", p.wc, name, lname, service, lsvc, call, lcall)
	host := p.Hosts()

	svcs, ok := host[lname]
	if !ok {
		svcs = NewServices()
		p.h[lname] = svcs
	}

	rpcs, ok := svcs[lsvc]
	if !ok {
		rpcs = NewRPCs()
		svcs[lsvc] = rpcs
	}

	_, ok = rpcs[lcall]
	if !ok {
		rpcs[lcall] = true
	}

	rpcs.clear(p.wc)
	svcs.clear(p.wc)
	host.clear(p.wc)
	//fmt.Printf("policy now:\n%#v\n", p)
}
