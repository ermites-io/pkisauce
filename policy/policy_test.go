//go:build go1.17
// +build go1.17

package policy

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"sort"
	"testing"

	"ermites.io/pkisauce/scramble"
)

type PolicyEntry struct {
	host string
	svc  string
	rpc  string
	erc  bool // used only for check tests.
}

type PolicyTestVectorEntry struct {
	in     []PolicyEntry
	outmap Policy
}

// let's pimp policy for our tests.
// kind of like a merkle tree, but the moronic, string way :)
func (h Hosts) hash() []byte {
	var clist []string

	// global hash
	gh := sha256.New()

	for host, svcs := range h {
		for svc, rpcs := range svcs {
			for rpc := range rpcs {
				// hash svc, rpc tuples
				str := fmt.Sprintf("h:%s.s:%s.r:%s", host, svc, rpc)
				clist = append(clist, str)
			}
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
	TestPolicyRpcOnly = []PolicyTestVectorEntry{
		{
			[]PolicyEntry{
				{"c1", "public", "*", false},
				{"c1", "public", "prout", false},
			}, // in
			Policy{
				"*",
				nil,
				Hosts{"c1": {"public": {"*": true}}},
			}, // out expected
		}, // #1
		{
			[]PolicyEntry{
				{"c1", "public", "*", false},
				{"c1", "public", "prout", false},
				{"c1", "public", "treat", false},
			}, // in
			Policy{
				"*",
				nil,
				Hosts{"c1": {"public": {"*": true}}},
			}, // out expected
		}, // #2
		{ // #3 START
			[]PolicyEntry{
				{"c1", "public", "*", false},
				{"c1", "public", "prout", false},
				{"c1", "publica", "prout", false},
			}, // in
			Policy{
				"*",
				nil,
				Hosts{"c1": {
					"public":  {"*": true},
					"publica": {"prout": true},
				}},
			}, // out expected
		}, // #3 END
		{ // #4 START
			[]PolicyEntry{
				{"c1", "public", "*", false},
				{"c1", "public", "prout", false},  // erased by #0 rule
				{"c1", "publica", "prout", false}, // erased by #4 rule
				{"c1", "public", "test", false},   // erased by #0 rule
				{"c1", "*", "prout", false},
			}, // in
			Policy{
				"*",
				nil,
				Hosts{
					"c1": {
						"public": {"*": true},
						"*":      {"prout": true},
					},
				},
			}, // out expected
		}, // #4 END
		{ // #5 START
			[]PolicyEntry{
				{"c1", "public", "*", false},
				{"c1", "public", "prout", false},  // erased by #0 rule
				{"c1", "publica", "prout", false}, // erased by #4 rule
				{"c1", "public", "test", false},   // erased by #0 rule
				{"c1", "*", "prout", false},
				{"c2", "*", "*", false},
			}, // in
			Policy{
				"*",
				nil,
				Hosts{
					"c1": {
						"public": {"*": true},
						"*":      {"prout": true},
					},
					"c2": {
						"*": {"*": true},
					},
				},
			}, // out expected
		}, // #5 END
		{ // #6 START
			[]PolicyEntry{
				{"c1", "public", "*", false},
				{"c1", "public", "prout", false},  // erased by #0 rule
				{"c1", "publica", "prout", false}, // erased by #4 rule
				{"c1", "public", "test", false},   // erased by #0 rule
				{"c1", "*", "prout", false},
				{"c2", "*", "*", false},
				{"c1", "*", "*", false},
			}, // in
			Policy{
				"*",
				nil,
				Hosts{
					"c1": {
						"*": {"*": true},
					},
					"c2": {
						"*": {"*": true},
					},
				},
			}, // out expected
		}, // #6 END
		{ // #7 START
			[]PolicyEntry{
				{"c1", "public", "*", false},
				{"c1", "public", "prout", false}, // erased
				{"c1", "other", "prout", false},  // erased
				{"c1", "global", "prout", false}, // erased
				{"c1", "public", "test", false},  // erased
				{"c1", "*", "prout", false},
				{"c2", "*", "*", false},
				{"c3", "tortue", "ping", false},
			}, // in
			Policy{
				"*",
				nil,
				Hosts{
					"c1": {
						"public": {"*": true},
						"*":      {"prout": true},
					},
					"c2": {
						"*": {"*": true},
					},
					"c3": {
						"tortue": {"ping": true},
					},
				},
			}, // out expected
		}, // #7 END
		{ // #8 START
			[]PolicyEntry{
				{"c1", "public", "prout", false},
				{"c1", "other", "prout", false},
				{"c2", "*", "prout", false},
				{"c3", "tortue", "ping", false},
			}, // in
			Policy{
				"*",
				nil,
				Hosts{
					"c1": {
						"public": {"prout": true},
						"other":  {"prout": true},
					},
					"c2": {
						"*": {"prout": true},
					},
					"c3": {
						"tortue": {"ping": true},
					},
				},
			}, // out expected
		}, // #8 END
		{ // #9 START
			[]PolicyEntry{
				{"c1", "public", "prout", false}, // erased
				{"c1", "other", "prout", false},
				{"c2", "*", "prout", false},
				{"c3", "tortue", "ping", false},
				{"*", "public", "prout", false},
			}, // in
			Policy{
				"*",
				nil,
				Hosts{
					"*": {
						"public": {"prout": true},
					},
					"c1": {
						"other": {"prout": true},
					},
					"c2": {
						"*": {"prout": true},
					},
					"c3": {
						"tortue": {"ping": true},
					},
				},
			}, // out expected
		}, // #9 END
		{ // #10 START
			[]PolicyEntry{
				{"c1", "public", "prout", false}, // erased
				{"c4", "public", "coin", false},  // erased
				{"c1", "other", "prout", false},
				{"c2", "*", "prout", false},
				{"c3", "tortue", "ping", false},
				{"*", "public", "*", false},
			}, // in
			Policy{
				"*",
				nil,
				Hosts{
					"*": {
						"public": {"*": true},
					},
					"c1": {
						"other": {"prout": true},
					},
					"c2": {
						"*": {"prout": true},
					},
					"c3": {
						"tortue": {"ping": true},
					},
				},
			}, // out expected
		}, // #10 END
	}
	//hSvcRpc = NewHosts()
	TestSvcRpcVector = []PolicyEntry{
		{"h1", "svc2", "call", true},
		{"h1", "svc2", "callfake", false},
		{"h1", "svc1", "rpc", true},
		{"h1", "svc1", "toto", true},
		{"h1", "svc4", "rpc", false},
		{"h2", "svc2", "call", false},
		{"h1", "bla", "test", true},
		{"h1", "bli", "test", true},
		{"h2", "bli", "test", false},
		{"h1", "svc3", "prout", true},
		{"h1", "svc3", "prout3", false},
		{"h2", "svc3", "prout", false},
	}

	TestAllVector = []PolicyEntry{
		{"h1", "svc2", "call", true},
		{"h1", "svc2", "callfake", false},
		{"h1", "svc1", "rpc", true},
		{"h1", "svc1", "toto", true},
		{"h1", "svc4", "rpc", false},
		{"h2", "svc2", "call", false},
		{"h1", "bla", "test", true},
		{"h1", "bli", "test", true},
		{"h2", "bli", "test", false},
		{"h1", "svc3", "prout", true},
		{"h1", "svc3", "prout3", false},
		{"h2", "svc3", "prout", false},
		{"h2", "svc5", "rpc11", false},
		{"h2", "svc5", "rpc12", true},
		{"toto", "svc5", "rpc12", true},
		{"h7", "kjdk", "gero", true},
		{"h7", "ger", "gera", false},
		{"h7", "ger", "gero", true},
		{"h7", "ger", "ping", true},
		{"h12", "key", "ping", true},
		{"dkjsj", "key", "ping", true},
		{"h9", "key", "pong", false},
		{"h9", "service", "pong", true},
		{"h28", "service", "tero", true},
		{"h32", "service", "tero", true},
		{"h32", "sevice", "tero", false},
	}

	TestAllowAllVector = []PolicyEntry{
		{"h1", "svc2", "call", true},
		{"h1", "svc2", "callfake", true},
		{"h1", "svc1", "rpc", true},
		{"h1", "svc1", "toto", true},
		{"h1", "svc4", "rpc", true},
		{"h2", "svc2", "call", true},
		{"h1", "bla", "test", true},
		{"h1", "bli", "test", true},
		{"h2", "bli", "test", true},
		{"h1", "svc3", "prout", true},
		{"h1", "svc3", "prout3", true},
		{"h2", "svc3", "prout", true},
		{"h2", "svc5", "rpc11", true},
		{"h2", "svc5", "rpc12", true},
		{"toto", "svc5", "rpc12", true},
		{"h7", "kjdk", "gero", true},
		{"h7", "ger", "gera", true},
		{"h7", "ger", "gero", true},
		{"h7", "ger", "ping", true},
		{"h12", "key", "ping", true},
		{"dkjsj", "key", "ping", true},
		{"h9", "key", "pong", true},
		{"h9", "service", "pong", true},
		{"h28", "service", "tero", true},
		{"h32", "service", "tero", true},
		{"h32", "sevice", "tero", true},
	}
)

func setupPolicy(p *Policy) {
	p.Entry("h1", "svc2", "call")
	p.Entry("h1", "svc1", "rpc")
	p.Entry("h1", "svc1", "*")
	p.Entry("h1", "*", "test")
	p.Entry("h1", "svc3", "prout")
}

func setupAddPolicy(p *Policy) {
	p.Entry("*", "svc5", "rpc12")
	p.Entry("h7", "*", "gero")
	p.Entry("*", "*", "ping")
	p.Entry("h7", "*", "ping")
	p.Entry("*", "service", "*")
}

func setupAddPolicyAll(p *Policy) {
	p.Entry("*", "*", "*")
}

func TestPolicyEntry(t *testing.T) {
	for tei, te := range TestPolicyRpcOnly {
		p := New(nil, "*")
		//t.Logf("[%d] test policy entry\n", tei)
		for _, pe := range te.in {
			p.Entry(pe.host, pe.svc, pe.rpc)
		}
		// now compare the result to what is expected
		if len(p.h) != len(te.outmap.h) {
			t.Fatalf("[%d] incorrect cleaning\nout:\n%v\nVS\nexpected:\n%v\n", tei, p.h, te.outmap.h)
		}

		policyHash := p.h.hash()
		expectedPolicyHash := te.outmap.h.hash()

		if !bytes.Equal(policyHash, expectedPolicyHash) {
			t.Fatalf("\n-- inconsistent --\npolicyHash: %x\npolicy: %v\nVS\nexpectedHash: %x\nexpected:\n%v\n",
				policyHash,
				p,
				expectedPolicyHash,
				te.outmap)
		}
	}
}

func TestPolicySVCRPC(t *testing.T) {
	pol := New(nil, wc)
	setupPolicy(pol)

	//t.Logf("testing %d entries\n", len(TestSvcRpcVector))
	//t.Logf("rules: %v\n", hSvcRpc)
	for ti, tvec := range TestSvcRpcVector {
		trc := pol.Check(tvec.host, tvec.svc, tvec.rpc)
		if tvec.erc != trc {
			t.Fatalf("test[%d] expect: %v VS rc: %v\n", ti, tvec.erc, trc)
		}
	}
}

func TestPolicyAll(t *testing.T) {
	ph := scramble.NewHasher([]byte("proutproutoroutout"))
	pol := New(ph, wc)
	setupPolicy(pol)
	setupAddPolicy(pol)

	//t.Logf("testing %d entries\n", len(TestAllVector))
	//t.Logf("rules: %#v\n", pol)
	for ti, tvec := range TestAllVector {
		trc := pol.Check(tvec.host, tvec.svc, tvec.rpc)
		if tvec.erc != trc {
			t.Fatalf("test[%d] expect: %v VS rc: %v\n", ti, tvec.erc, trc)
		}
	}
}

func TestPolicyAllowAll(t *testing.T) {
	pol := New(nil, wc)
	setupPolicy(pol)
	setupAddPolicy(pol)
	setupAddPolicyAll(pol)

	//t.Logf("testing %d entries\n", len(TestAllowAllVector))
	//t.Logf("rules: %v\n", hSvcRpc)
	for ti, tvec := range TestAllowAllVector {
		trc := pol.Check(tvec.host, tvec.svc, tvec.rpc)
		if tvec.erc != trc {
			t.Fatalf("test[%d] expect: %v VS rc: %v\n", ti, tvec.erc, trc)
		}
	}
}
