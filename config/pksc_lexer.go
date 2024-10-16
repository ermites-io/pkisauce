// Code generated from pksc.g4 by ANTLR 4.13.2. DO NOT EDIT.

package config

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type pkscLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var PkscLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func pksclexerLexerInit() {
	staticData := &PkscLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"WS", "COMMENT", "LINE_COMMENT", "TERMINATOR", "LCURLY", "RCURLY", "COMMA",
		"QUOTE", "SEMI", "KWANY", "KWDBG", "KWPASS", "KWON", "KWRPC", "KWDAEMON",
		"KWPATH", "KWUSE", "KWGOGRPC", "KWGOTLS", "KWPEM", "HNAME", "IDENTIFIER",
		"LETTER", "DECIMAL_DIGIT", "PATH", "CHAR_VALUE", "CHAR_ESCAPE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 23, 204, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 1, 0, 4, 0, 57, 8, 0, 11, 0, 12, 0, 58, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 1, 1, 1, 5, 1, 67, 8, 1, 10, 1, 12, 1, 70, 9, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 81, 8, 2, 10, 2, 12, 2, 84,
		9, 2, 1, 2, 1, 2, 1, 3, 4, 3, 89, 8, 3, 11, 3, 12, 3, 90, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 20, 1, 20, 3, 20, 164, 8, 20, 1, 20, 1, 20, 1, 21, 1,
		21, 1, 21, 5, 21, 171, 8, 21, 10, 21, 12, 21, 174, 9, 21, 1, 22, 1, 22,
		1, 23, 1, 23, 1, 24, 1, 24, 5, 24, 182, 8, 24, 10, 24, 12, 24, 185, 9,
		24, 1, 24, 1, 24, 1, 24, 5, 24, 190, 8, 24, 10, 24, 12, 24, 193, 9, 24,
		1, 24, 3, 24, 196, 8, 24, 1, 25, 1, 25, 3, 25, 200, 8, 25, 1, 26, 1, 26,
		1, 26, 3, 68, 183, 191, 0, 27, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16,
		33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 0, 47, 0, 49, 23, 51,
		0, 53, 0, 1, 0, 6, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 13, 13, 3,
		0, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 3, 0, 0, 0, 10, 10, 92, 92, 9,
		0, 34, 34, 39, 39, 92, 92, 97, 98, 102, 102, 110, 110, 114, 114, 116, 116,
		118, 118, 210, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0,
		7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0,
		0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0,
		0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0,
		0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1,
		0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 49,
		1, 0, 0, 0, 1, 56, 1, 0, 0, 0, 3, 62, 1, 0, 0, 0, 5, 76, 1, 0, 0, 0, 7,
		88, 1, 0, 0, 0, 9, 94, 1, 0, 0, 0, 11, 96, 1, 0, 0, 0, 13, 98, 1, 0, 0,
		0, 15, 100, 1, 0, 0, 0, 17, 102, 1, 0, 0, 0, 19, 104, 1, 0, 0, 0, 21, 108,
		1, 0, 0, 0, 23, 114, 1, 0, 0, 0, 25, 119, 1, 0, 0, 0, 27, 122, 1, 0, 0,
		0, 29, 126, 1, 0, 0, 0, 31, 133, 1, 0, 0, 0, 33, 138, 1, 0, 0, 0, 35, 142,
		1, 0, 0, 0, 37, 150, 1, 0, 0, 0, 39, 157, 1, 0, 0, 0, 41, 161, 1, 0, 0,
		0, 43, 167, 1, 0, 0, 0, 45, 175, 1, 0, 0, 0, 47, 177, 1, 0, 0, 0, 49, 195,
		1, 0, 0, 0, 51, 199, 1, 0, 0, 0, 53, 201, 1, 0, 0, 0, 55, 57, 7, 0, 0,
		0, 56, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 58, 59,
		1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 61, 6, 0, 0, 0, 61, 2, 1, 0, 0, 0,
		62, 63, 5, 47, 0, 0, 63, 64, 5, 42, 0, 0, 64, 68, 1, 0, 0, 0, 65, 67, 9,
		0, 0, 0, 66, 65, 1, 0, 0, 0, 67, 70, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 68,
		66, 1, 0, 0, 0, 69, 71, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 71, 72, 5, 42,
		0, 0, 72, 73, 5, 47, 0, 0, 73, 74, 1, 0, 0, 0, 74, 75, 6, 1, 1, 0, 75,
		4, 1, 0, 0, 0, 76, 77, 5, 47, 0, 0, 77, 78, 5, 47, 0, 0, 78, 82, 1, 0,
		0, 0, 79, 81, 8, 1, 0, 0, 80, 79, 1, 0, 0, 0, 81, 84, 1, 0, 0, 0, 82, 80,
		1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 85, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0,
		85, 86, 6, 2, 1, 0, 86, 6, 1, 0, 0, 0, 87, 89, 7, 1, 0, 0, 88, 87, 1, 0,
		0, 0, 89, 90, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 92,
		1, 0, 0, 0, 92, 93, 6, 3, 1, 0, 93, 8, 1, 0, 0, 0, 94, 95, 5, 123, 0, 0,
		95, 10, 1, 0, 0, 0, 96, 97, 5, 125, 0, 0, 97, 12, 1, 0, 0, 0, 98, 99, 5,
		44, 0, 0, 99, 14, 1, 0, 0, 0, 100, 101, 5, 34, 0, 0, 101, 16, 1, 0, 0,
		0, 102, 103, 5, 59, 0, 0, 103, 18, 1, 0, 0, 0, 104, 105, 5, 97, 0, 0, 105,
		106, 5, 110, 0, 0, 106, 107, 5, 121, 0, 0, 107, 20, 1, 0, 0, 0, 108, 109,
		5, 100, 0, 0, 109, 110, 5, 101, 0, 0, 110, 111, 5, 98, 0, 0, 111, 112,
		5, 117, 0, 0, 112, 113, 5, 103, 0, 0, 113, 22, 1, 0, 0, 0, 114, 115, 5,
		112, 0, 0, 115, 116, 5, 97, 0, 0, 116, 117, 5, 115, 0, 0, 117, 118, 5,
		115, 0, 0, 118, 24, 1, 0, 0, 0, 119, 120, 5, 111, 0, 0, 120, 121, 5, 110,
		0, 0, 121, 26, 1, 0, 0, 0, 122, 123, 5, 114, 0, 0, 123, 124, 5, 112, 0,
		0, 124, 125, 5, 99, 0, 0, 125, 28, 1, 0, 0, 0, 126, 127, 5, 100, 0, 0,
		127, 128, 5, 97, 0, 0, 128, 129, 5, 101, 0, 0, 129, 130, 5, 109, 0, 0,
		130, 131, 5, 111, 0, 0, 131, 132, 5, 110, 0, 0, 132, 30, 1, 0, 0, 0, 133,
		134, 5, 112, 0, 0, 134, 135, 5, 97, 0, 0, 135, 136, 5, 116, 0, 0, 136,
		137, 5, 104, 0, 0, 137, 32, 1, 0, 0, 0, 138, 139, 5, 117, 0, 0, 139, 140,
		5, 115, 0, 0, 140, 141, 5, 101, 0, 0, 141, 34, 1, 0, 0, 0, 142, 143, 5,
		103, 0, 0, 143, 144, 5, 111, 0, 0, 144, 145, 5, 45, 0, 0, 145, 146, 5,
		103, 0, 0, 146, 147, 5, 114, 0, 0, 147, 148, 5, 112, 0, 0, 148, 149, 5,
		99, 0, 0, 149, 36, 1, 0, 0, 0, 150, 151, 5, 103, 0, 0, 151, 152, 5, 111,
		0, 0, 152, 153, 5, 45, 0, 0, 153, 154, 5, 116, 0, 0, 154, 155, 5, 108,
		0, 0, 155, 156, 5, 115, 0, 0, 156, 38, 1, 0, 0, 0, 157, 158, 5, 112, 0,
		0, 158, 159, 5, 101, 0, 0, 159, 160, 5, 109, 0, 0, 160, 40, 1, 0, 0, 0,
		161, 163, 3, 15, 7, 0, 162, 164, 3, 43, 21, 0, 163, 162, 1, 0, 0, 0, 163,
		164, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 166, 3, 15, 7, 0, 166, 42,
		1, 0, 0, 0, 167, 172, 3, 45, 22, 0, 168, 171, 3, 45, 22, 0, 169, 171, 3,
		47, 23, 0, 170, 168, 1, 0, 0, 0, 170, 169, 1, 0, 0, 0, 171, 174, 1, 0,
		0, 0, 172, 170, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 44, 1, 0, 0, 0,
		174, 172, 1, 0, 0, 0, 175, 176, 7, 2, 0, 0, 176, 46, 1, 0, 0, 0, 177, 178,
		7, 3, 0, 0, 178, 48, 1, 0, 0, 0, 179, 183, 5, 39, 0, 0, 180, 182, 3, 51,
		25, 0, 181, 180, 1, 0, 0, 0, 182, 185, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0,
		183, 181, 1, 0, 0, 0, 184, 186, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 186,
		196, 5, 39, 0, 0, 187, 191, 5, 34, 0, 0, 188, 190, 3, 51, 25, 0, 189, 188,
		1, 0, 0, 0, 190, 193, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 191, 189, 1, 0,
		0, 0, 192, 194, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 194, 196, 5, 34, 0, 0,
		195, 179, 1, 0, 0, 0, 195, 187, 1, 0, 0, 0, 196, 50, 1, 0, 0, 0, 197, 200,
		3, 53, 26, 0, 198, 200, 8, 4, 0, 0, 199, 197, 1, 0, 0, 0, 199, 198, 1,
		0, 0, 0, 200, 52, 1, 0, 0, 0, 201, 202, 5, 92, 0, 0, 202, 203, 7, 5, 0,
		0, 203, 54, 1, 0, 0, 0, 12, 0, 58, 68, 82, 90, 163, 170, 172, 183, 191,
		195, 199, 2, 6, 0, 0, 0, 1, 0,
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

// pkscLexerInit initializes any static state used to implement pkscLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewpkscLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func PkscLexerInit() {
	staticData := &PkscLexerLexerStaticData
	staticData.once.Do(pksclexerLexerInit)
}

// NewpkscLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewpkscLexer(input antlr.CharStream) *pkscLexer {
	PkscLexerInit()
	l := new(pkscLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &PkscLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "pksc.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// pkscLexer tokens.
const (
	pkscLexerWS           = 1
	pkscLexerCOMMENT      = 2
	pkscLexerLINE_COMMENT = 3
	pkscLexerTERMINATOR   = 4
	pkscLexerLCURLY       = 5
	pkscLexerRCURLY       = 6
	pkscLexerCOMMA        = 7
	pkscLexerQUOTE        = 8
	pkscLexerSEMI         = 9
	pkscLexerKWANY        = 10
	pkscLexerKWDBG        = 11
	pkscLexerKWPASS       = 12
	pkscLexerKWON         = 13
	pkscLexerKWRPC        = 14
	pkscLexerKWDAEMON     = 15
	pkscLexerKWPATH       = 16
	pkscLexerKWUSE        = 17
	pkscLexerKWGOGRPC     = 18
	pkscLexerKWGOTLS      = 19
	pkscLexerKWPEM        = 20
	pkscLexerHNAME        = 21
	pkscLexerIDENTIFIER   = 22
	pkscLexerPATH         = 23
)
