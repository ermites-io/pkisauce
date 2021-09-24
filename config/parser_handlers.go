//go:build go1.16
// +build go1.16

package config

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func (c *Parser) VisitErrorNode(node antlr.ErrorNode) {
	errSymbol := node.GetSymbol()
	parent := node.GetParent().(*antlr.BaseParserRuleContext)
	parentToken := parent.GetStart()
	fmt.Printf("parent: %s\n", parentToken.GetText())
	fmt.Printf("error: %s at line: %d\n", errSymbol.GetText(), errSymbol.GetLine())
}

func (c *Parser) EnterDaemon_stmt(d *Daemon_stmtContext) {

	name := parseHName(d.HNAME())
	pathname := parsePath(d.PATH())
	flavorname, err := c.parseGenType(d.Gen_type())
	if err != nil {
		panic(err)
	}

	p, debug, err := c.parsePassBlock(d.Pass_block())
	if err != nil {
		panic(err)
	}

	opts := NewOptions(pathname, flavorname, p, debug)
	//name, options, err := c.parseDaemonEntry(d)

	err = c.config.Add(name, opts)
	if err != nil {
		panic(err)
	}
}
