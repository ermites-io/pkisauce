// Code generated from pksc.g4 by ANTLR 4.13.2. DO NOT EDIT.

package config // pksc
import "github.com/antlr4-go/antlr/v4"

// BasepkscListener is a complete listener for a parse tree produced by pkscParser.
type BasepkscListener struct{}

var _ pkscListener = &BasepkscListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasepkscListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasepkscListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasepkscListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasepkscListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterGen_type is called when production gen_type is entered.
func (s *BasepkscListener) EnterGen_type(ctx *Gen_typeContext) {}

// ExitGen_type is called when production gen_type is exited.
func (s *BasepkscListener) ExitGen_type(ctx *Gen_typeContext) {}

// EnterName_elt_list is called when production name_elt_list is entered.
func (s *BasepkscListener) EnterName_elt_list(ctx *Name_elt_listContext) {}

// ExitName_elt_list is called when production name_elt_list is exited.
func (s *BasepkscListener) ExitName_elt_list(ctx *Name_elt_listContext) {}

// EnterName_list is called when production name_list is entered.
func (s *BasepkscListener) EnterName_list(ctx *Name_listContext) {}

// ExitName_list is called when production name_list is exited.
func (s *BasepkscListener) ExitName_list(ctx *Name_listContext) {}

// EnterName_block is called when production name_block is entered.
func (s *BasepkscListener) EnterName_block(ctx *Name_blockContext) {}

// ExitName_block is called when production name_block is exited.
func (s *BasepkscListener) ExitName_block(ctx *Name_blockContext) {}

// EnterPass_stmt is called when production pass_stmt is entered.
func (s *BasepkscListener) EnterPass_stmt(ctx *Pass_stmtContext) {}

// ExitPass_stmt is called when production pass_stmt is exited.
func (s *BasepkscListener) ExitPass_stmt(ctx *Pass_stmtContext) {}

// EnterPass_block is called when production pass_block is entered.
func (s *BasepkscListener) EnterPass_block(ctx *Pass_blockContext) {}

// ExitPass_block is called when production pass_block is exited.
func (s *BasepkscListener) ExitPass_block(ctx *Pass_blockContext) {}

// EnterDaemon_stmt is called when production daemon_stmt is entered.
func (s *BasepkscListener) EnterDaemon_stmt(ctx *Daemon_stmtContext) {}

// ExitDaemon_stmt is called when production daemon_stmt is exited.
func (s *BasepkscListener) ExitDaemon_stmt(ctx *Daemon_stmtContext) {}

// EnterConfig is called when production config is entered.
func (s *BasepkscListener) EnterConfig(ctx *ConfigContext) {}

// ExitConfig is called when production config is exited.
func (s *BasepkscListener) ExitConfig(ctx *ConfigContext) {}
