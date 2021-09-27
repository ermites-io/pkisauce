//go:build go1.17
// +build go1.17

package config

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func (pr *Parser) VisitErrorNode(node antlr.ErrorNode) {
	errSymbol := node.GetSymbol()
	parent := node.GetParent().(*antlr.BaseParserRuleContext)
	parentToken := parent.GetStart()
	fmt.Printf("parent: %s\n", parentToken.GetText())
	fmt.Printf("error: %s at line: %d\n", errSymbol.GetText(), errSymbol.GetLine())
}

func (pr *Parser) EnterDaemon_stmt(d *Daemon_stmtContext) {

	name := parseHName(d.HNAME())
	pathname := parsePath(d.PATH())
	flavorname, err := pr.parseGenType(d.Gen_type())
	if err != nil {
		panic(err)
	}

	p, debug, err := pr.parsePassBlock(d.Pass_block())
	if err != nil {
		panic(err)
	}

	opts := NewOptions(pathname, flavorname, p, debug)
	//name, options, err := c.parseDaemonEntry(d)

	err = pr.config.Add(name, opts)
	if err != nil {
		panic(err)
	}

	pr.entries++
}
