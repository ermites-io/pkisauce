//go:build go1.16

package policy

func (r RPCs) check(wc, call string) bool {
	// handle wildcard.
	_, ok := r[wc]
	if ok {
		return true
	}

	_, ok = r[call]
	return ok
}

func (s Services) check(wc, svc, call string) bool {
	// handle wildcard.
	rpcs, ok := s[wc]
	if ok {
		if rpcs.check(wc, call) {
			return true
		}
	}

	rpcs, ok = s[svc]
	if ok {
		return rpcs.check(wc, call)
	}

	return false
}

func (h Hosts) check(wc, host, svc, call string) bool {
	// handle wildcard.
	s, ok := h[wc]
	if ok {
		if s.check(wc, svc, call) {
			return true
		}
	}

	// no matching wildcards.
	s, ok = h[host]
	if ok {
		return s.check(wc, svc, call)
	}

	return false
}
