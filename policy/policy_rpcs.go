//go:build go1.17
// +build go1.17

package policy

func NewRPCs() RPCs {
	r := make(RPCs)
	return r
}

func (r RPCs) clear(wc string) {
	_, ok := r[wc]
	if ok {
		// erase the rest
		for k := range r {
			if k == wc {
				continue
			}
			delete(r, k)
		}
	}
}

func (r RPCs) hasWc(wc string) bool {
	_, ok := r[wc]
	return ok
}

//func (r RPCs) Union(o RPCs) (torm []string) {
func (r RPCs) Union(wc string, o RPCs) (torm RPCs) {
	if len(o) == 0 || len(r) == 0 {
		return
	}

	torm = NewRPCs()

	_, ok := r[wc]
	if ok {
		torm = o
		return
	}

	_, ok = o[wc]
	if ok {
		torm = r
		return
	}

	for rpc := range r {
		_, ok := o[rpc]
		if ok {
			torm[rpc] = true
		}
	}

	return
}

func (r RPCs) RmV(rpc RPCs) {
	for lrpc := range rpc {
		delete(r, lrpc)
	}
}
