// Code generated from pksc.g4 by ANTLR 4.9.2. DO NOT EDIT.

package config // pksc
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 24, 100,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 3, 5, 3, 22, 10, 3, 3, 3, 3, 3,
	7, 3, 26, 10, 3, 12, 3, 14, 3, 29, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5,
	3, 5, 3, 5, 5, 5, 38, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 54, 10, 6, 3, 7, 3, 7,
	7, 7, 58, 10, 7, 12, 7, 14, 7, 61, 11, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8,
	79, 10, 8, 3, 9, 7, 9, 82, 10, 9, 12, 9, 14, 9, 85, 11, 9, 3, 9, 7, 9,
	88, 10, 9, 12, 9, 14, 9, 91, 11, 9, 3, 9, 6, 9, 94, 10, 9, 13, 9, 14, 9,
	95, 5, 9, 98, 10, 9, 3, 9, 2, 2, 10, 2, 4, 6, 8, 10, 12, 14, 16, 2, 3,
	3, 2, 20, 21, 2, 104, 2, 18, 3, 2, 2, 2, 4, 21, 3, 2, 2, 2, 6, 30, 3, 2,
	2, 2, 8, 37, 3, 2, 2, 2, 10, 53, 3, 2, 2, 2, 12, 55, 3, 2, 2, 2, 14, 78,
	3, 2, 2, 2, 16, 97, 3, 2, 2, 2, 18, 19, 9, 2, 2, 2, 19, 3, 3, 2, 2, 2,
	20, 22, 7, 22, 2, 2, 21, 20, 3, 2, 2, 2, 21, 22, 3, 2, 2, 2, 22, 27, 3,
	2, 2, 2, 23, 24, 7, 9, 2, 2, 24, 26, 7, 22, 2, 2, 25, 23, 3, 2, 2, 2, 26,
	29, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 5, 3, 2, 2,
	2, 29, 27, 3, 2, 2, 2, 30, 31, 7, 7, 2, 2, 31, 32, 5, 4, 3, 2, 32, 33,
	7, 8, 2, 2, 33, 7, 3, 2, 2, 2, 34, 38, 7, 22, 2, 2, 35, 38, 5, 6, 4, 2,
	36, 38, 7, 12, 2, 2, 37, 34, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 37, 36, 3,
	2, 2, 2, 38, 9, 3, 2, 2, 2, 39, 40, 7, 14, 2, 2, 40, 41, 5, 8, 5, 2, 41,
	42, 7, 15, 2, 2, 42, 43, 5, 8, 5, 2, 43, 44, 7, 16, 2, 2, 44, 45, 5, 8,
	5, 2, 45, 46, 7, 11, 2, 2, 46, 54, 3, 2, 2, 2, 47, 48, 7, 14, 2, 2, 48,
	49, 5, 8, 5, 2, 49, 50, 7, 11, 2, 2, 50, 54, 3, 2, 2, 2, 51, 52, 7, 13,
	2, 2, 52, 54, 7, 11, 2, 2, 53, 39, 3, 2, 2, 2, 53, 47, 3, 2, 2, 2, 53,
	51, 3, 2, 2, 2, 54, 11, 3, 2, 2, 2, 55, 59, 7, 7, 2, 2, 56, 58, 5, 10,
	6, 2, 57, 56, 3, 2, 2, 2, 58, 61, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 59, 60,
	3, 2, 2, 2, 60, 62, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 62, 63, 7, 8, 2, 2,
	63, 13, 3, 2, 2, 2, 64, 65, 7, 17, 2, 2, 65, 66, 7, 22, 2, 2, 66, 67, 7,
	18, 2, 2, 67, 68, 7, 24, 2, 2, 68, 69, 7, 19, 2, 2, 69, 70, 5, 2, 2, 2,
	70, 71, 5, 12, 7, 2, 71, 79, 3, 2, 2, 2, 72, 73, 7, 17, 2, 2, 73, 74, 7,
	22, 2, 2, 74, 75, 7, 18, 2, 2, 75, 76, 7, 24, 2, 2, 76, 77, 7, 19, 2, 2,
	77, 79, 5, 2, 2, 2, 78, 64, 3, 2, 2, 2, 78, 72, 3, 2, 2, 2, 79, 15, 3,
	2, 2, 2, 80, 82, 7, 4, 2, 2, 81, 80, 3, 2, 2, 2, 82, 85, 3, 2, 2, 2, 83,
	81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 98, 3, 2, 2, 2, 85, 83, 3, 2, 2,
	2, 86, 88, 7, 5, 2, 2, 87, 86, 3, 2, 2, 2, 88, 91, 3, 2, 2, 2, 89, 87,
	3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 98, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2,
	92, 94, 5, 14, 8, 2, 93, 92, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 93, 3,
	2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 98, 3, 2, 2, 2, 97, 83, 3, 2, 2, 2, 97,
	89, 3, 2, 2, 2, 97, 93, 3, 2, 2, 2, 98, 17, 3, 2, 2, 2, 12, 21, 27, 37,
	53, 59, 78, 83, 89, 95, 97,
}
var literalNames = []string{
	"", "", "", "", "", "'{'", "'}'", "','", "'\"'", "';'", "'any'", "'debug'",
	"'pass'", "'on'", "'rpc'", "'daemon'", "'path'", "'use'", "'go-grpc'",
	"'go-tls'",
}
var symbolicNames = []string{
	"", "WS", "COMMENT", "LINE_COMMENT", "TERMINATOR", "LCURLY", "RCURLY",
	"COMMA", "QUOTE", "SEMI", "KWANY", "KWDBG", "KWPASS", "KWON", "KWRPC",
	"KWDAEMON", "KWPATH", "KWUSE", "KWGOGRPC", "KWGOTLS", "HNAME", "IDENTIFIER",
	"PATH",
}

var ruleNames = []string{
	"gen_type", "name_elt_list", "name_list", "name_block", "pass_stmt", "pass_block",
	"daemon_stmt", "config",
}

type pkscParser struct {
	*antlr.BaseParser
}

// NewpkscParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *pkscParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewpkscParser(input antlr.TokenStream) *pkscParser {
	this := new(pkscParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "pksc.g4"

	return this
}

// pkscParser tokens.
const (
	pkscParserEOF          = antlr.TokenEOF
	pkscParserWS           = 1
	pkscParserCOMMENT      = 2
	pkscParserLINE_COMMENT = 3
	pkscParserTERMINATOR   = 4
	pkscParserLCURLY       = 5
	pkscParserRCURLY       = 6
	pkscParserCOMMA        = 7
	pkscParserQUOTE        = 8
	pkscParserSEMI         = 9
	pkscParserKWANY        = 10
	pkscParserKWDBG        = 11
	pkscParserKWPASS       = 12
	pkscParserKWON         = 13
	pkscParserKWRPC        = 14
	pkscParserKWDAEMON     = 15
	pkscParserKWPATH       = 16
	pkscParserKWUSE        = 17
	pkscParserKWGOGRPC     = 18
	pkscParserKWGOTLS      = 19
	pkscParserHNAME        = 20
	pkscParserIDENTIFIER   = 21
	pkscParserPATH         = 22
)

// pkscParser rules.
const (
	pkscParserRULE_gen_type      = 0
	pkscParserRULE_name_elt_list = 1
	pkscParserRULE_name_list     = 2
	pkscParserRULE_name_block    = 3
	pkscParserRULE_pass_stmt     = 4
	pkscParserRULE_pass_block    = 5
	pkscParserRULE_daemon_stmt   = 6
	pkscParserRULE_config        = 7
)

// IGen_typeContext is an interface to support dynamic dispatch.
type IGen_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGen_typeContext differentiates from other interfaces.
	IsGen_typeContext()
}

type Gen_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGen_typeContext() *Gen_typeContext {
	var p = new(Gen_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pkscParserRULE_gen_type
	return p
}

func (*Gen_typeContext) IsGen_typeContext() {}

func NewGen_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Gen_typeContext {
	var p = new(Gen_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pkscParserRULE_gen_type

	return p
}

func (s *Gen_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Gen_typeContext) KWGOGRPC() antlr.TerminalNode {
	return s.GetToken(pkscParserKWGOGRPC, 0)
}

func (s *Gen_typeContext) KWGOTLS() antlr.TerminalNode {
	return s.GetToken(pkscParserKWGOTLS, 0)
}

func (s *Gen_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Gen_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Gen_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pkscListener); ok {
		listenerT.EnterGen_type(s)
	}
}

func (s *Gen_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pkscListener); ok {
		listenerT.ExitGen_type(s)
	}
}

func (p *pkscParser) Gen_type() (localctx IGen_typeContext) {
	localctx = NewGen_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, pkscParserRULE_gen_type)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		_la = p.GetTokenStream().LA(1)

		if !(_la == pkscParserKWGOGRPC || _la == pkscParserKWGOTLS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IName_elt_listContext is an interface to support dynamic dispatch.
type IName_elt_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsName_elt_listContext differentiates from other interfaces.
	IsName_elt_listContext()
}

type Name_elt_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyName_elt_listContext() *Name_elt_listContext {
	var p = new(Name_elt_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pkscParserRULE_name_elt_list
	return p
}

func (*Name_elt_listContext) IsName_elt_listContext() {}

func NewName_elt_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Name_elt_listContext {
	var p = new(Name_elt_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pkscParserRULE_name_elt_list

	return p
}

func (s *Name_elt_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Name_elt_listContext) AllHNAME() []antlr.TerminalNode {
	return s.GetTokens(pkscParserHNAME)
}

func (s *Name_elt_listContext) HNAME(i int) antlr.TerminalNode {
	return s.GetToken(pkscParserHNAME, i)
}

func (s *Name_elt_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(pkscParserCOMMA)
}

func (s *Name_elt_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(pkscParserCOMMA, i)
}

func (s *Name_elt_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Name_elt_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Name_elt_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pkscListener); ok {
		listenerT.EnterName_elt_list(s)
	}
}

func (s *Name_elt_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pkscListener); ok {
		listenerT.ExitName_elt_list(s)
	}
}

func (p *pkscParser) Name_elt_list() (localctx IName_elt_listContext) {
	localctx = NewName_elt_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, pkscParserRULE_name_elt_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == pkscParserHNAME {
		{
			p.SetState(18)
			p.Match(pkscParserHNAME)
		}

	}
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == pkscParserCOMMA {
		{
			p.SetState(21)
			p.Match(pkscParserCOMMA)
		}
		{
			p.SetState(22)
			p.Match(pkscParserHNAME)
		}

		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IName_listContext is an interface to support dynamic dispatch.
type IName_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsName_listContext differentiates from other interfaces.
	IsName_listContext()
}

type Name_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyName_listContext() *Name_listContext {
	var p = new(Name_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pkscParserRULE_name_list
	return p
}

func (*Name_listContext) IsName_listContext() {}

func NewName_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Name_listContext {
	var p = new(Name_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pkscParserRULE_name_list

	return p
}

func (s *Name_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Name_listContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(pkscParserLCURLY, 0)
}

func (s *Name_listContext) Name_elt_list() IName_elt_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IName_elt_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IName_elt_listContext)
}

func (s *Name_listContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(pkscParserRCURLY, 0)
}

func (s *Name_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Name_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Name_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pkscListener); ok {
		listenerT.EnterName_list(s)
	}
}

func (s *Name_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pkscListener); ok {
		listenerT.ExitName_list(s)
	}
}

func (p *pkscParser) Name_list() (localctx IName_listContext) {
	localctx = NewName_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, pkscParserRULE_name_list)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)
		p.Match(pkscParserLCURLY)
	}
	{
		p.SetState(29)
		p.Name_elt_list()
	}
	{
		p.SetState(30)
		p.Match(pkscParserRCURLY)
	}

	return localctx
}

// IName_blockContext is an interface to support dynamic dispatch.
type IName_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsName_blockContext differentiates from other interfaces.
	IsName_blockContext()
}

type Name_blockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyName_blockContext() *Name_blockContext {
	var p = new(Name_blockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pkscParserRULE_name_block
	return p
}

func (*Name_blockContext) IsName_blockContext() {}

func NewName_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Name_blockContext {
	var p = new(Name_blockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pkscParserRULE_name_block

	return p
}

func (s *Name_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Name_blockContext) HNAME() antlr.TerminalNode {
	return s.GetToken(pkscParserHNAME, 0)
}

func (s *Name_blockContext) Name_list() IName_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IName_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IName_listContext)
}

func (s *Name_blockContext) KWANY() antlr.TerminalNode {
	return s.GetToken(pkscParserKWANY, 0)
}

func (s *Name_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Name_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Name_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pkscListener); ok {
		listenerT.EnterName_block(s)
	}
}

func (s *Name_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pkscListener); ok {
		listenerT.ExitName_block(s)
	}
}

func (p *pkscParser) Name_block() (localctx IName_blockContext) {
	localctx = NewName_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, pkscParserRULE_name_block)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(35)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case pkscParserHNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(32)
			p.Match(pkscParserHNAME)
		}

	case pkscParserLCURLY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(33)
			p.Name_list()
		}

	case pkscParserKWANY:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(34)
			p.Match(pkscParserKWANY)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPass_stmtContext is an interface to support dynamic dispatch.
type IPass_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPass_stmtContext differentiates from other interfaces.
	IsPass_stmtContext()
}

type Pass_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPass_stmtContext() *Pass_stmtContext {
	var p = new(Pass_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pkscParserRULE_pass_stmt
	return p
}

func (*Pass_stmtContext) IsPass_stmtContext() {}

func NewPass_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pass_stmtContext {
	var p = new(Pass_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pkscParserRULE_pass_stmt

	return p
}

func (s *Pass_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Pass_stmtContext) KWPASS() antlr.TerminalNode {
	return s.GetToken(pkscParserKWPASS, 0)
}

func (s *Pass_stmtContext) AllName_block() []IName_blockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IName_blockContext)(nil)).Elem())
	var tst = make([]IName_blockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IName_blockContext)
		}
	}

	return tst
}

func (s *Pass_stmtContext) Name_block(i int) IName_blockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IName_blockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IName_blockContext)
}

func (s *Pass_stmtContext) KWON() antlr.TerminalNode {
	return s.GetToken(pkscParserKWON, 0)
}

func (s *Pass_stmtContext) KWRPC() antlr.TerminalNode {
	return s.GetToken(pkscParserKWRPC, 0)
}

func (s *Pass_stmtContext) SEMI() antlr.TerminalNode {
	return s.GetToken(pkscParserSEMI, 0)
}

func (s *Pass_stmtContext) KWDBG() antlr.TerminalNode {
	return s.GetToken(pkscParserKWDBG, 0)
}

func (s *Pass_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pass_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pass_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pkscListener); ok {
		listenerT.EnterPass_stmt(s)
	}
}

func (s *Pass_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pkscListener); ok {
		listenerT.ExitPass_stmt(s)
	}
}

func (p *pkscParser) Pass_stmt() (localctx IPass_stmtContext) {
	localctx = NewPass_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, pkscParserRULE_pass_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(37)
			p.Match(pkscParserKWPASS)
		}
		{
			p.SetState(38)
			p.Name_block()
		}
		{
			p.SetState(39)
			p.Match(pkscParserKWON)
		}
		{
			p.SetState(40)
			p.Name_block()
		}
		{
			p.SetState(41)
			p.Match(pkscParserKWRPC)
		}
		{
			p.SetState(42)
			p.Name_block()
		}
		{
			p.SetState(43)
			p.Match(pkscParserSEMI)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(45)
			p.Match(pkscParserKWPASS)
		}
		{
			p.SetState(46)
			p.Name_block()
		}
		{
			p.SetState(47)
			p.Match(pkscParserSEMI)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(49)
			p.Match(pkscParserKWDBG)
		}
		{
			p.SetState(50)
			p.Match(pkscParserSEMI)
		}

	}

	return localctx
}

// IPass_blockContext is an interface to support dynamic dispatch.
type IPass_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPass_blockContext differentiates from other interfaces.
	IsPass_blockContext()
}

type Pass_blockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPass_blockContext() *Pass_blockContext {
	var p = new(Pass_blockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pkscParserRULE_pass_block
	return p
}

func (*Pass_blockContext) IsPass_blockContext() {}

func NewPass_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pass_blockContext {
	var p = new(Pass_blockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pkscParserRULE_pass_block

	return p
}

func (s *Pass_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Pass_blockContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(pkscParserLCURLY, 0)
}

func (s *Pass_blockContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(pkscParserRCURLY, 0)
}

func (s *Pass_blockContext) AllPass_stmt() []IPass_stmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPass_stmtContext)(nil)).Elem())
	var tst = make([]IPass_stmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPass_stmtContext)
		}
	}

	return tst
}

func (s *Pass_blockContext) Pass_stmt(i int) IPass_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPass_stmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPass_stmtContext)
}

func (s *Pass_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pass_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pass_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pkscListener); ok {
		listenerT.EnterPass_block(s)
	}
}

func (s *Pass_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pkscListener); ok {
		listenerT.ExitPass_block(s)
	}
}

func (p *pkscParser) Pass_block() (localctx IPass_blockContext) {
	localctx = NewPass_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, pkscParserRULE_pass_block)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Match(pkscParserLCURLY)
	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == pkscParserKWDBG || _la == pkscParserKWPASS {
		{
			p.SetState(54)
			p.Pass_stmt()
		}

		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(60)
		p.Match(pkscParserRCURLY)
	}

	return localctx
}

// IDaemon_stmtContext is an interface to support dynamic dispatch.
type IDaemon_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDaemon_stmtContext differentiates from other interfaces.
	IsDaemon_stmtContext()
}

type Daemon_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDaemon_stmtContext() *Daemon_stmtContext {
	var p = new(Daemon_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pkscParserRULE_daemon_stmt
	return p
}

func (*Daemon_stmtContext) IsDaemon_stmtContext() {}

func NewDaemon_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Daemon_stmtContext {
	var p = new(Daemon_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pkscParserRULE_daemon_stmt

	return p
}

func (s *Daemon_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Daemon_stmtContext) KWDAEMON() antlr.TerminalNode {
	return s.GetToken(pkscParserKWDAEMON, 0)
}

func (s *Daemon_stmtContext) HNAME() antlr.TerminalNode {
	return s.GetToken(pkscParserHNAME, 0)
}

func (s *Daemon_stmtContext) KWPATH() antlr.TerminalNode {
	return s.GetToken(pkscParserKWPATH, 0)
}

func (s *Daemon_stmtContext) PATH() antlr.TerminalNode {
	return s.GetToken(pkscParserPATH, 0)
}

func (s *Daemon_stmtContext) KWUSE() antlr.TerminalNode {
	return s.GetToken(pkscParserKWUSE, 0)
}

func (s *Daemon_stmtContext) Gen_type() IGen_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGen_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGen_typeContext)
}

func (s *Daemon_stmtContext) Pass_block() IPass_blockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPass_blockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPass_blockContext)
}

func (s *Daemon_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Daemon_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Daemon_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pkscListener); ok {
		listenerT.EnterDaemon_stmt(s)
	}
}

func (s *Daemon_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pkscListener); ok {
		listenerT.ExitDaemon_stmt(s)
	}
}

func (p *pkscParser) Daemon_stmt() (localctx IDaemon_stmtContext) {
	localctx = NewDaemon_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, pkscParserRULE_daemon_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)
			p.Match(pkscParserKWDAEMON)
		}
		{
			p.SetState(63)
			p.Match(pkscParserHNAME)
		}
		{
			p.SetState(64)
			p.Match(pkscParserKWPATH)
		}
		{
			p.SetState(65)
			p.Match(pkscParserPATH)
		}
		{
			p.SetState(66)
			p.Match(pkscParserKWUSE)
		}
		{
			p.SetState(67)
			p.Gen_type()
		}
		{
			p.SetState(68)
			p.Pass_block()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(70)
			p.Match(pkscParserKWDAEMON)
		}
		{
			p.SetState(71)
			p.Match(pkscParserHNAME)
		}
		{
			p.SetState(72)
			p.Match(pkscParserKWPATH)
		}
		{
			p.SetState(73)
			p.Match(pkscParserPATH)
		}
		{
			p.SetState(74)
			p.Match(pkscParserKWUSE)
		}
		{
			p.SetState(75)
			p.Gen_type()
		}

	}

	return localctx
}

// IConfigContext is an interface to support dynamic dispatch.
type IConfigContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConfigContext differentiates from other interfaces.
	IsConfigContext()
}

type ConfigContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConfigContext() *ConfigContext {
	var p = new(ConfigContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pkscParserRULE_config
	return p
}

func (*ConfigContext) IsConfigContext() {}

func NewConfigContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConfigContext {
	var p = new(ConfigContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pkscParserRULE_config

	return p
}

func (s *ConfigContext) GetParser() antlr.Parser { return s.parser }

func (s *ConfigContext) AllCOMMENT() []antlr.TerminalNode {
	return s.GetTokens(pkscParserCOMMENT)
}

func (s *ConfigContext) COMMENT(i int) antlr.TerminalNode {
	return s.GetToken(pkscParserCOMMENT, i)
}

func (s *ConfigContext) AllLINE_COMMENT() []antlr.TerminalNode {
	return s.GetTokens(pkscParserLINE_COMMENT)
}

func (s *ConfigContext) LINE_COMMENT(i int) antlr.TerminalNode {
	return s.GetToken(pkscParserLINE_COMMENT, i)
}

func (s *ConfigContext) AllDaemon_stmt() []IDaemon_stmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDaemon_stmtContext)(nil)).Elem())
	var tst = make([]IDaemon_stmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDaemon_stmtContext)
		}
	}

	return tst
}

func (s *ConfigContext) Daemon_stmt(i int) IDaemon_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDaemon_stmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDaemon_stmtContext)
}

func (s *ConfigContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConfigContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConfigContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pkscListener); ok {
		listenerT.EnterConfig(s)
	}
}

func (s *ConfigContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pkscListener); ok {
		listenerT.ExitConfig(s)
	}
}

func (p *pkscParser) Config() (localctx IConfigContext) {
	localctx = NewConfigContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, pkscParserRULE_config)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == pkscParserCOMMENT {
			{
				p.SetState(78)
				p.Match(pkscParserCOMMENT)
			}

			p.SetState(83)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == pkscParserLINE_COMMENT {
			{
				p.SetState(84)
				p.Match(pkscParserLINE_COMMENT)
			}

			p.SetState(89)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == pkscParserKWDAEMON {
			{
				p.SetState(90)
				p.Daemon_stmt()
			}

			p.SetState(93)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}
