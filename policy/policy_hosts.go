//go:build go1.16

package policy

func NewHosts() Hosts {
	p := make(Hosts)
	return p
}

// we have a case failing, see tests..
//func (p *Policy) clear(wc string) {
func (h Hosts) clear(wc string) {
	// host: * - svc: *
	svc, ok := h[wc]
	switch ok {
	case true:
		for host, othersvc := range h {
			if host == wc {
				continue
			}
			torm := svc.Union(wc, othersvc)
			othersvc.RmV(torm)

			if len(othersvc) == 0 {
				delete(h, host)
			}
		}
	case false:
		fallthrough
	default:
		// DO NOTHING
	} // p[wc] switch
}
