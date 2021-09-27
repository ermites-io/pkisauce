//go:build go1.17
// +build go1.17

package config

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"ermites.io/pkisauce/policy"
)

type Parser struct {
	*BasepkscListener
	entries int // let's count the number of entries found.
	key     []byte
	config  Config
}

func (pr *Parser) Config() Config {
	//fmt.Printf("DANS LA CONFIG: %v\n", psr.config)
	return pr.config
}

//func NewParser(key []byte) *Parser {
func NewParser() *Parser {
	// random key for the hmac'ing of rules.
	return &Parser{
		key:    nil, // might disappear in the end..
		config: NewConfig(),
	}
}

var (
	// aliases of the same func
	parseHName = parseTermNodeText
	parsePath  = parseTermNodeText
	wc         = "*"
)

func parseTermNodeText(hn antlr.TerminalNode) string {
	quoted := hn.GetSymbol().GetText()
	str := quoted[1 : len(quoted)-1]
	return str
}

// IGen_typeContext
func (pr *Parser) parseGenType(igt IGen_typeContext) (dtype int, err error) {
	gt, ok := igt.(*Gen_typeContext)
	if !ok {
		err = fmt.Errorf("error type")
		return
	}

	switch {
	case gt.KWGOGRPC() != nil:
		dtype = FlavorGoGRPC
	case gt.KWGOTLS() != nil:
		dtype = FlavorGoTLS
	case gt.KWPEM() != nil:
		dtype = FlavorPEM
	default:
		err = fmt.Errorf("unknown type")
	}

	return
}

func parseNameBlock(nbr IName_blockContext) (names []string, err error) {
	nb, ok := nbr.(*Name_blockContext)
	if !ok {
		err = fmt.Errorf("name block type error")
		return
	}

	nbhn := nb.HNAME()
	nbany := nb.KWANY()
	nbnlr := nb.Name_list() // IName_listContext

	switch {
	case nbhn != nil:
		names = append(names, parseHName(nbhn))
	case nbany != nil:
		names = append(names, "*")
	case nbnlr != nil:
		nbnl, ok := nbnlr.(*Name_listContext)
		if !ok {
			err = fmt.Errorf("name list type error")
			return
		}

		// IName_elt_listContext
		nbnlelr := nbnl.Name_elt_list()
		nbnlel, ok := nbnlelr.(*Name_elt_listContext)
		if !ok {
			err = fmt.Errorf("name elt list type error")
			return
		}

		nl := nbnlel.AllHNAME()
		for _, n := range nl {
			names = append(names, parseHName(n))
		}
	}
	//fmt.Printf("nbhn: %v nbany: %v nbnl: %v\n", nbhn, nbany, nbnl)
	return
}

//func parsePassBlock(pbr IPass_blockContext) (rules []*Rule, err error) {
func (pr *Parser) parsePassBlock(pbr IPass_blockContext) (p *policy.Policy, debug bool, err error) {
	//var rules []*Rule

	pb, ok := pbr.(*Pass_blockContext)
	if !ok {
		err = fmt.Errorf("error type")
		return
	}

	p = policy.New(nil, "*")

	// []IPass_stmtContext
	pbrules := pb.AllPass_stmt()
	for _, r := range pbrules {
		// one rul each.
		//fmt.Printf("[%d] \n", i)
		pbrule := r.(*Pass_stmtContext)
		// it returns []IName_blockContext
		pbrunb := pbrule.AllName_block()
		// nb0: srcs   nb1: svcs  nb2: rpcs
		switch len(pbrunb) {
		case 1:
			// we only have src names...
			srcs, perr := parseNameBlock(pbrunb[0])
			if perr != nil {
				err = perr
				return
			}
			//fmt.Printf("rule: %p srcs: %p\n", rule, srcs)
			for _, src := range srcs {
				p.Entry(src, wc, wc)
			}

		case 3: // srcs, svcs, rpcs
			// we have src names...
			srcs, perr := parseNameBlock(pbrunb[0])
			if perr != nil {
				err = perr
				return
			}

			// we have svc names...
			svcs, perr := parseNameBlock(pbrunb[1])
			if perr != nil {
				err = perr
				return
			}

			// we have svc names...
			rpcs, perr := parseNameBlock(pbrunb[2])
			if perr != nil {
				err = perr
				return
			}

			for _, src := range srcs {
				for _, svc := range svcs {
					for _, rpc := range rpcs {
						p.Entry(src, svc, rpc)
					}
				}
			}
		default:
			err = fmt.Errorf("unknown parse error")
			return
		}
	}

	// create the associated policy
	return
}
