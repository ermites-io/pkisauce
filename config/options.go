//go:build go1.16
// +build go1.16

package config

import (
	"errors"
	"os"
	"path/filepath"

	"ermites.io/pkisauce/policy"
)

//type Daemons map[string]Options
type Options struct {
	Flavor int    // flags defining the type of sauce: gRPC-go, regular tls, pem, private
	Path   string // the directory where to write our sauce
	Policy *policy.Policy
	Debug  bool
}

func (opt Options) IsValid(name string) (err error) {
	// check that options are valid..
	// 1. is the name correct?
	// wildcard is a protected name
	if name == "*" {
		err = errors.New("invalid entry")
		return
	}

	//fmt.Printf("OPT: %v FLAVORKEYWORD: %v\n", opt.Flavor, FlavorKeyword)

	// 2. is flavor correct?
	switch opt.Flavor {
	case FlavorGoTLS, FlavorGoGRPC, FlavorPEM:
	default:
		err = errors.New("invalid flavor")
		return
	}

	// 3. is Path correct?
	path, err := filepath.Abs(opt.Path)
	if err != nil {
		return
	}

	// XXX activate when we get close to release the first version.
	_, err = os.Stat(path)
	return
}

func (opt Options) Clients() (list []string) {
	for dn := range opt.Policy.Hosts() {
		list = append(list, dn)
	}
	return
}

func NewOptions(path string, flavor int, p *policy.Policy, debug bool) Options {
	return Options{
		Flavor: flavor,
		Path:   path,
		Policy: p,
		Debug:  debug,
	}
}
