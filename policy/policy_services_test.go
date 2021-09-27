//go:build go1.17
// +build go1.17

package policy

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"sort"
	"testing"
)

type SvcUnionTestEntry struct {
	a Services
	b Services
	e Services
}

func (s Services) hash() []byte {
	var clist []string

	// global hash
	gh := sha256.New()

	for svc, rpcs := range s {
		for rpc := range rpcs {
			// hash svc, rpc tuples
			str := fmt.Sprintf("s:%s.r:%s", svc, rpc)
			clist = append(clist, str)
		}
	}

	sort.Strings(clist)
	for _, c := range clist {
		c256 := sha256.Sum256([]byte(c))
		gh.Write(c256[:])
	}

	ghash := gh.Sum(nil)
	return ghash[:]
}

var (
	ServiceUnionTestVector = []SvcUnionTestEntry{
		{
			Services{
				"svc": {
					"call1": true,
					"call2": true,
				},
				"svc2": {
					"toru": true,
				},
				"public": {
					"call": true,
					"ping": true,
					"pong": true,
				},
			},
			Services{
				"svc": {
					"call2": true,
				},
				"public": {
					"call": true,
				},
			},
			Services{ // EXPECTED
				"svc":    {"call2": true},
				"public": {"call": true},
			},
		},
		{
			Services{
				"svc": {
					"ping": true,
					"pong": true,
				},
				"osvc": {
					"ping": true,
					"pong": true,
				},
			},
			Services{"*": {"ping": true}},
			Services{ // EXPECTED
				"svc":  {"ping": true},
				"osvc": {"ping": true},
			},
		},
		{
			Services{
				"svc": {
					"ping": true,
					"pong": true,
				},
				"osvc": {
					"ping": true,
					"pong": true,
				},
			},
			Services{"*": {"*": true}},
			Services{ // EXPECTED
				"svc":  {"ping": true, "pong": true},
				"osvc": {"ping": true, "pong": true},
			},
		},
		{
			Services{
				"svc": {
					"ping": true,
					"pong": true,
				},
				"osvc": {
					"ping": true,
					"pong": true,
				},
			},
			Services{"svc": {"*": true}},
			Services{ // EXPECTED
				"svc": {"ping": true, "pong": true},
			},
		},
	}
)

func TestServiceUnion(t *testing.T) {
	for i, te := range ServiceUnionTestVector {
		usvc := te.a.Union("*", te.b)
		usvch := usvc.hash()
		teh := te.e.hash()

		//if bytes.Compare(usvch, teh) != 0 {
		if !bytes.Equal(usvch, teh) {
			t.Fatalf("\n-- inconsistent[%d] --\nusvcHash: %x\npolicy: %v\nVS\nexpectedHash: %x\nexpected:\n%v\n",
				i,
				usvch,
				usvc,
				teh,
				te.e)
		}
	}
}
