// Code generated from pksc.g4 by ANTLR 4.13.2. DO NOT EDIT.

package config // pksc
import "github.com/antlr4-go/antlr/v4"

// pkscListener is a complete listener for a parse tree produced by pkscParser.
type pkscListener interface {
	antlr.ParseTreeListener

	// EnterGen_type is called when entering the gen_type production.
	EnterGen_type(c *Gen_typeContext)

	// EnterName_elt_list is called when entering the name_elt_list production.
	EnterName_elt_list(c *Name_elt_listContext)

	// EnterName_list is called when entering the name_list production.
	EnterName_list(c *Name_listContext)

	// EnterName_block is called when entering the name_block production.
	EnterName_block(c *Name_blockContext)

	// EnterPass_stmt is called when entering the pass_stmt production.
	EnterPass_stmt(c *Pass_stmtContext)

	// EnterPass_block is called when entering the pass_block production.
	EnterPass_block(c *Pass_blockContext)

	// EnterDaemon_stmt is called when entering the daemon_stmt production.
	EnterDaemon_stmt(c *Daemon_stmtContext)

	// EnterConfig is called when entering the config production.
	EnterConfig(c *ConfigContext)

	// ExitGen_type is called when exiting the gen_type production.
	ExitGen_type(c *Gen_typeContext)

	// ExitName_elt_list is called when exiting the name_elt_list production.
	ExitName_elt_list(c *Name_elt_listContext)

	// ExitName_list is called when exiting the name_list production.
	ExitName_list(c *Name_listContext)

	// ExitName_block is called when exiting the name_block production.
	ExitName_block(c *Name_blockContext)

	// ExitPass_stmt is called when exiting the pass_stmt production.
	ExitPass_stmt(c *Pass_stmtContext)

	// ExitPass_block is called when exiting the pass_block production.
	ExitPass_block(c *Pass_blockContext)

	// ExitDaemon_stmt is called when exiting the daemon_stmt production.
	ExitDaemon_stmt(c *Daemon_stmtContext)

	// ExitConfig is called when exiting the config production.
	ExitConfig(c *ConfigContext)
}
