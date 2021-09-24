//go:build go1.16

package policy

func NewServices() Services {
	s := make(Services)
	return s
}

func (s Services) clear(wc string) {
	wcrpcs, ok := s[wc]
	switch ok {
	case true:
		// find the common rpc to the svc:*  and remove them from the more specific rules
		// if rpc is empty after that clear, remove the service, it's not allowing anything.
		// svc:* rpc: { * }
		// svc:p rpc: { x, y, z } <== REMOVE
		// svc:t rpc: { x, y }     <=== REMOVE
		// svc:u rpc: { x, b, c }  <=== REMOVE
		for svc, rpcs := range s {
			if svc == wc {
				continue
			}
			urpcs := rpcs.Union(wc, wcrpcs)
			rpcs.RmV(urpcs)

			if len(rpcs) == 0 {
				delete(s, svc)
			}
		}
	case false:
		fallthrough
	default:
		// DO NOTHING
	}
}

func (s Services) Union(wc string, os Services) (out Services) {
	if len(s) == 0 || len(os) == 0 {
		return
	}
	out = NewServices()

	// TODO: rewrite better..
	// wildcard defined in s
	rpcs, ok := s[wc]
	if ok { // s[*] exists
		if rpcs.hasWc(wc) {
			//return all the os
			out = os
		}

		for osvc, orpcs := range os {
			jrpcs := rpcs.Union(wc, orpcs)
			if len(jrpcs) > 0 {
				out[osvc] = jrpcs
			}
		}
		return
	}

	// wildcard defined in os
	rpcs, ok = os[wc]
	if ok { // s[*] exists
		if rpcs.hasWc(wc) {
			out = s
		}

		for osvc, orpcs := range s {
			jrpcs := rpcs.Union(wc, orpcs)
			if len(jrpcs) > 0 {
				out[osvc] = jrpcs
			}
		}
		return
	}

	// TODO: handle wildcard
	for svc, rpcs := range s {
		orpcs, ok := os[svc]
		switch ok {
		case true:
			jrpcs := rpcs.Union(wc, orpcs)
			if len(jrpcs) > 0 {
				out[svc] = jrpcs
			}
		case false:
			fallthrough
		default:
			// nothing to do.
		}
	}

	return
}

func (s Services) RmV(svcs Services) {
	for svc, rpcs := range svcs {
		r, ok := s[svc]
		if ok {
			r.RmV(rpcs)
		}

		if len(r) == 0 {
			delete(s, svc)
		}
	}
}
