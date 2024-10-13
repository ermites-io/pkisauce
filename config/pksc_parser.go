// Code generated from pksc.g4 by ANTLR 4.13.2. DO NOT EDIT.

package config // pksc
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type pkscParser struct {
	*antlr.BaseParser
}

var PkscParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func pkscParserInit() {
	staticData := &PkscParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "", "'{'", "'}'", "','", "'\"'", "';'", "'any'", "'debug'",
		"'pass'", "'on'", "'rpc'", "'daemon'", "'path'", "'use'", "'go-grpc'",
		"'go-tls'", "'pem'",
	}
	staticData.SymbolicNames = []string{
		"", "WS", "COMMENT", "LINE_COMMENT", "TERMINATOR", "LCURLY", "RCURLY",
		"COMMA", "QUOTE", "SEMI", "KWANY", "KWDBG", "KWPASS", "KWON", "KWRPC",
		"KWDAEMON", "KWPATH", "KWUSE", "KWGOGRPC", "KWGOTLS", "KWPEM", "HNAME",
		"IDENTIFIER", "PATH",
	}
	staticData.RuleNames = []string{
		"gen_type", "name_elt_list", "name_list", "name_block", "pass_stmt",
		"pass_block", "daemon_stmt", "config",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 23, 98, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 1, 3, 1, 20, 8, 1,
		1, 1, 1, 1, 5, 1, 24, 8, 1, 10, 1, 12, 1, 27, 9, 1, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 3, 3, 3, 36, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 52, 8, 4, 1, 5, 1,
		5, 5, 5, 56, 8, 5, 10, 5, 12, 5, 59, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6,
		77, 8, 6, 1, 7, 5, 7, 80, 8, 7, 10, 7, 12, 7, 83, 9, 7, 1, 7, 5, 7, 86,
		8, 7, 10, 7, 12, 7, 89, 9, 7, 1, 7, 4, 7, 92, 8, 7, 11, 7, 12, 7, 93, 3,
		7, 96, 8, 7, 1, 7, 0, 0, 8, 0, 2, 4, 6, 8, 10, 12, 14, 0, 1, 1, 0, 18,
		20, 102, 0, 16, 1, 0, 0, 0, 2, 19, 1, 0, 0, 0, 4, 28, 1, 0, 0, 0, 6, 35,
		1, 0, 0, 0, 8, 51, 1, 0, 0, 0, 10, 53, 1, 0, 0, 0, 12, 76, 1, 0, 0, 0,
		14, 95, 1, 0, 0, 0, 16, 17, 7, 0, 0, 0, 17, 1, 1, 0, 0, 0, 18, 20, 5, 21,
		0, 0, 19, 18, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 25, 1, 0, 0, 0, 21, 22,
		5, 7, 0, 0, 22, 24, 5, 21, 0, 0, 23, 21, 1, 0, 0, 0, 24, 27, 1, 0, 0, 0,
		25, 23, 1, 0, 0, 0, 25, 26, 1, 0, 0, 0, 26, 3, 1, 0, 0, 0, 27, 25, 1, 0,
		0, 0, 28, 29, 5, 5, 0, 0, 29, 30, 3, 2, 1, 0, 30, 31, 5, 6, 0, 0, 31, 5,
		1, 0, 0, 0, 32, 36, 5, 21, 0, 0, 33, 36, 3, 4, 2, 0, 34, 36, 5, 10, 0,
		0, 35, 32, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 35, 34, 1, 0, 0, 0, 36, 7, 1,
		0, 0, 0, 37, 38, 5, 12, 0, 0, 38, 39, 3, 6, 3, 0, 39, 40, 5, 13, 0, 0,
		40, 41, 3, 6, 3, 0, 41, 42, 5, 14, 0, 0, 42, 43, 3, 6, 3, 0, 43, 44, 5,
		9, 0, 0, 44, 52, 1, 0, 0, 0, 45, 46, 5, 12, 0, 0, 46, 47, 3, 6, 3, 0, 47,
		48, 5, 9, 0, 0, 48, 52, 1, 0, 0, 0, 49, 50, 5, 11, 0, 0, 50, 52, 5, 9,
		0, 0, 51, 37, 1, 0, 0, 0, 51, 45, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 52, 9,
		1, 0, 0, 0, 53, 57, 5, 5, 0, 0, 54, 56, 3, 8, 4, 0, 55, 54, 1, 0, 0, 0,
		56, 59, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 60, 1,
		0, 0, 0, 59, 57, 1, 0, 0, 0, 60, 61, 5, 6, 0, 0, 61, 11, 1, 0, 0, 0, 62,
		63, 5, 15, 0, 0, 63, 64, 5, 21, 0, 0, 64, 65, 5, 16, 0, 0, 65, 66, 5, 23,
		0, 0, 66, 67, 5, 17, 0, 0, 67, 68, 3, 0, 0, 0, 68, 69, 3, 10, 5, 0, 69,
		77, 1, 0, 0, 0, 70, 71, 5, 15, 0, 0, 71, 72, 5, 21, 0, 0, 72, 73, 5, 16,
		0, 0, 73, 74, 5, 23, 0, 0, 74, 75, 5, 17, 0, 0, 75, 77, 3, 0, 0, 0, 76,
		62, 1, 0, 0, 0, 76, 70, 1, 0, 0, 0, 77, 13, 1, 0, 0, 0, 78, 80, 5, 2, 0,
		0, 79, 78, 1, 0, 0, 0, 80, 83, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 81, 82,
		1, 0, 0, 0, 82, 96, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 84, 86, 5, 3, 0, 0,
		85, 84, 1, 0, 0, 0, 86, 89, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 88, 1,
		0, 0, 0, 88, 96, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 90, 92, 3, 12, 6, 0, 91,
		90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 93, 94, 1, 0, 0,
		0, 94, 96, 1, 0, 0, 0, 95, 81, 1, 0, 0, 0, 95, 87, 1, 0, 0, 0, 95, 91,
		1, 0, 0, 0, 96, 15, 1, 0, 0, 0, 10, 19, 25, 35, 51, 57, 76, 81, 87, 93,
		95,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// pkscParserInit initializes any static state used to implement pkscParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewpkscParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PkscParserInit() {
	staticData := &PkscParserStaticData
	staticData.once.Do(pkscParserInit)
}

// NewpkscParser produces a new parser instance for the optional input antlr.TokenStream.
func NewpkscParser(input antlr.TokenStream) *pkscParser {
	PkscParserInit()
	this := new(pkscParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &PkscParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
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
	pkscParserKWPEM        = 20
	pkscParserHNAME        = 21
	pkscParserIDENTIFIER   = 22
	pkscParserPATH         = 23
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

	// Getter signatures
	KWGOGRPC() antlr.TerminalNode
	KWGOTLS() antlr.TerminalNode
	KWPEM() antlr.TerminalNode

	// IsGen_typeContext differentiates from other interfaces.
	IsGen_typeContext()
}

type Gen_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGen_typeContext() *Gen_typeContext {
	var p = new(Gen_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pkscParserRULE_gen_type
	return p
}

func InitEmptyGen_typeContext(p *Gen_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pkscParserRULE_gen_type
}

func (*Gen_typeContext) IsGen_typeContext() {}

func NewGen_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Gen_typeContext {
	var p = new(Gen_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

func (s *Gen_typeContext) KWPEM() antlr.TerminalNode {
	return s.GetToken(pkscParserKWPEM, 0)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1835008) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IName_elt_listContext is an interface to support dynamic dispatch.
type IName_elt_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllHNAME() []antlr.TerminalNode
	HNAME(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsName_elt_listContext differentiates from other interfaces.
	IsName_elt_listContext()
}

type Name_elt_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyName_elt_listContext() *Name_elt_listContext {
	var p = new(Name_elt_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pkscParserRULE_name_elt_list
	return p
}

func InitEmptyName_elt_listContext(p *Name_elt_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pkscParserRULE_name_elt_list
}

func (*Name_elt_listContext) IsName_elt_listContext() {}

func NewName_elt_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Name_elt_listContext {
	var p = new(Name_elt_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == pkscParserHNAME {
		{
			p.SetState(18)
			p.Match(pkscParserHNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == pkscParserCOMMA {
		{
			p.SetState(21)
			p.Match(pkscParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(22)
			p.Match(pkscParserHNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IName_listContext is an interface to support dynamic dispatch.
type IName_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LCURLY() antlr.TerminalNode
	Name_elt_list() IName_elt_listContext
	RCURLY() antlr.TerminalNode

	// IsName_listContext differentiates from other interfaces.
	IsName_listContext()
}

type Name_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyName_listContext() *Name_listContext {
	var p = new(Name_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pkscParserRULE_name_list
	return p
}

func InitEmptyName_listContext(p *Name_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pkscParserRULE_name_list
}

func (*Name_listContext) IsName_listContext() {}

func NewName_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Name_listContext {
	var p = new(Name_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = pkscParserRULE_name_list

	return p
}

func (s *Name_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Name_listContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(pkscParserLCURLY, 0)
}

func (s *Name_listContext) Name_elt_list() IName_elt_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IName_elt_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)
		p.Match(pkscParserLCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(29)
		p.Name_elt_list()
	}
	{
		p.SetState(30)
		p.Match(pkscParserRCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IName_blockContext is an interface to support dynamic dispatch.
type IName_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HNAME() antlr.TerminalNode
	Name_list() IName_listContext
	KWANY() antlr.TerminalNode

	// IsName_blockContext differentiates from other interfaces.
	IsName_blockContext()
}

type Name_blockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyName_blockContext() *Name_blockContext {
	var p = new(Name_blockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pkscParserRULE_name_block
	return p
}

func InitEmptyName_blockContext(p *Name_blockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pkscParserRULE_name_block
}

func (*Name_blockContext) IsName_blockContext() {}

func NewName_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Name_blockContext {
	var p = new(Name_blockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = pkscParserRULE_name_block

	return p
}

func (s *Name_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Name_blockContext) HNAME() antlr.TerminalNode {
	return s.GetToken(pkscParserHNAME, 0)
}

func (s *Name_blockContext) Name_list() IName_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IName_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case pkscParserHNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(32)
			p.Match(pkscParserHNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPass_stmtContext is an interface to support dynamic dispatch.
type IPass_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KWPASS() antlr.TerminalNode
	AllName_block() []IName_blockContext
	Name_block(i int) IName_blockContext
	KWON() antlr.TerminalNode
	KWRPC() antlr.TerminalNode
	SEMI() antlr.TerminalNode
	KWDBG() antlr.TerminalNode

	// IsPass_stmtContext differentiates from other interfaces.
	IsPass_stmtContext()
}

type Pass_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPass_stmtContext() *Pass_stmtContext {
	var p = new(Pass_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pkscParserRULE_pass_stmt
	return p
}

func InitEmptyPass_stmtContext(p *Pass_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pkscParserRULE_pass_stmt
}

func (*Pass_stmtContext) IsPass_stmtContext() {}

func NewPass_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pass_stmtContext {
	var p = new(Pass_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = pkscParserRULE_pass_stmt

	return p
}

func (s *Pass_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Pass_stmtContext) KWPASS() antlr.TerminalNode {
	return s.GetToken(pkscParserKWPASS, 0)
}

func (s *Pass_stmtContext) AllName_block() []IName_blockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IName_blockContext); ok {
			len++
		}
	}

	tst := make([]IName_blockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IName_blockContext); ok {
			tst[i] = t.(IName_blockContext)
			i++
		}
	}

	return tst
}

func (s *Pass_stmtContext) Name_block(i int) IName_blockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IName_blockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(37)
			p.Match(pkscParserKWPASS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(38)
			p.Name_block()
		}
		{
			p.SetState(39)
			p.Match(pkscParserKWON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(40)
			p.Name_block()
		}
		{
			p.SetState(41)
			p.Match(pkscParserKWRPC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(42)
			p.Name_block()
		}
		{
			p.SetState(43)
			p.Match(pkscParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(45)
			p.Match(pkscParserKWPASS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(46)
			p.Name_block()
		}
		{
			p.SetState(47)
			p.Match(pkscParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(49)
			p.Match(pkscParserKWDBG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.Match(pkscParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPass_blockContext is an interface to support dynamic dispatch.
type IPass_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LCURLY() antlr.TerminalNode
	RCURLY() antlr.TerminalNode
	AllPass_stmt() []IPass_stmtContext
	Pass_stmt(i int) IPass_stmtContext

	// IsPass_blockContext differentiates from other interfaces.
	IsPass_blockContext()
}

type Pass_blockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPass_blockContext() *Pass_blockContext {
	var p = new(Pass_blockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pkscParserRULE_pass_block
	return p
}

func InitEmptyPass_blockContext(p *Pass_blockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pkscParserRULE_pass_block
}

func (*Pass_blockContext) IsPass_blockContext() {}

func NewPass_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pass_blockContext {
	var p = new(Pass_blockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPass_stmtContext); ok {
			len++
		}
	}

	tst := make([]IPass_stmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPass_stmtContext); ok {
			tst[i] = t.(IPass_stmtContext)
			i++
		}
	}

	return tst
}

func (s *Pass_blockContext) Pass_stmt(i int) IPass_stmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPass_stmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Match(pkscParserLCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == pkscParserKWDBG || _la == pkscParserKWPASS {
		{
			p.SetState(54)
			p.Pass_stmt()
		}

		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(60)
		p.Match(pkscParserRCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDaemon_stmtContext is an interface to support dynamic dispatch.
type IDaemon_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KWDAEMON() antlr.TerminalNode
	HNAME() antlr.TerminalNode
	KWPATH() antlr.TerminalNode
	PATH() antlr.TerminalNode
	KWUSE() antlr.TerminalNode
	Gen_type() IGen_typeContext
	Pass_block() IPass_blockContext

	// IsDaemon_stmtContext differentiates from other interfaces.
	IsDaemon_stmtContext()
}

type Daemon_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDaemon_stmtContext() *Daemon_stmtContext {
	var p = new(Daemon_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pkscParserRULE_daemon_stmt
	return p
}

func InitEmptyDaemon_stmtContext(p *Daemon_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pkscParserRULE_daemon_stmt
}

func (*Daemon_stmtContext) IsDaemon_stmtContext() {}

func NewDaemon_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Daemon_stmtContext {
	var p = new(Daemon_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGen_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGen_typeContext)
}

func (s *Daemon_stmtContext) Pass_block() IPass_blockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPass_blockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)
			p.Match(pkscParserKWDAEMON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)
			p.Match(pkscParserHNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)
			p.Match(pkscParserKWPATH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)
			p.Match(pkscParserPATH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(66)
			p.Match(pkscParserKWUSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.Match(pkscParserHNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(72)
			p.Match(pkscParserKWPATH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(73)
			p.Match(pkscParserPATH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
			p.Match(pkscParserKWUSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.Gen_type()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConfigContext is an interface to support dynamic dispatch.
type IConfigContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCOMMENT() []antlr.TerminalNode
	COMMENT(i int) antlr.TerminalNode
	AllLINE_COMMENT() []antlr.TerminalNode
	LINE_COMMENT(i int) antlr.TerminalNode
	AllDaemon_stmt() []IDaemon_stmtContext
	Daemon_stmt(i int) IDaemon_stmtContext

	// IsConfigContext differentiates from other interfaces.
	IsConfigContext()
}

type ConfigContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConfigContext() *ConfigContext {
	var p = new(ConfigContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pkscParserRULE_config
	return p
}

func InitEmptyConfigContext(p *ConfigContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pkscParserRULE_config
}

func (*ConfigContext) IsConfigContext() {}

func NewConfigContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConfigContext {
	var p = new(ConfigContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDaemon_stmtContext); ok {
			len++
		}
	}

	tst := make([]IDaemon_stmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDaemon_stmtContext); ok {
			tst[i] = t.(IDaemon_stmtContext)
			i++
		}
	}

	return tst
}

func (s *ConfigContext) Daemon_stmt(i int) IDaemon_stmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDaemon_stmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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

	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == pkscParserCOMMENT {
			{
				p.SetState(78)
				p.Match(pkscParserCOMMENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(83)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == pkscParserLINE_COMMENT {
			{
				p.SetState(84)
				p.Match(pkscParserLINE_COMMENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(89)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == pkscParserKWDAEMON {
			{
				p.SetState(90)
				p.Daemon_stmt()
			}

			p.SetState(93)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
