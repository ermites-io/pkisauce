// Code generated from pksc.g4 by ANTLR 4.9.2. DO NOT EDIT.

package config

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 25, 206,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 3, 2, 6, 2, 59, 10, 2, 13, 2, 14, 2, 60, 3, 2, 3, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 7, 3, 69, 10, 3, 12, 3, 14, 3, 72, 11, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 83, 10, 4, 12, 4, 14, 4, 86, 11,
	4, 3, 4, 3, 4, 3, 5, 6, 5, 91, 10, 5, 13, 5, 14, 5, 92, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 22, 3, 22, 5, 22, 166, 10, 22, 3, 22, 3, 22, 3, 23,
	3, 23, 3, 23, 7, 23, 173, 10, 23, 12, 23, 14, 23, 176, 11, 23, 3, 24, 3,
	24, 3, 25, 3, 25, 3, 26, 3, 26, 7, 26, 184, 10, 26, 12, 26, 14, 26, 187,
	11, 26, 3, 26, 3, 26, 3, 26, 7, 26, 192, 10, 26, 12, 26, 14, 26, 195, 11,
	26, 3, 26, 5, 26, 198, 10, 26, 3, 27, 3, 27, 5, 27, 202, 10, 27, 3, 28,
	3, 28, 3, 28, 5, 70, 185, 193, 2, 29, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 2, 49, 2, 51,
	25, 53, 2, 55, 2, 3, 2, 8, 5, 2, 11, 12, 15, 15, 34, 34, 4, 2, 12, 12,
	15, 15, 5, 2, 67, 92, 97, 97, 99, 124, 3, 2, 50, 59, 5, 2, 2, 2, 12, 12,
	94, 94, 11, 2, 36, 36, 41, 41, 94, 94, 99, 100, 104, 104, 112, 112, 116,
	116, 118, 118, 120, 120, 2, 212, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2,
	7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2,
	2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2,
	2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2,
	2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3,
	2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45,
	3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 3, 58, 3, 2, 2, 2, 5, 64, 3, 2, 2, 2, 7,
	78, 3, 2, 2, 2, 9, 90, 3, 2, 2, 2, 11, 96, 3, 2, 2, 2, 13, 98, 3, 2, 2,
	2, 15, 100, 3, 2, 2, 2, 17, 102, 3, 2, 2, 2, 19, 104, 3, 2, 2, 2, 21, 106,
	3, 2, 2, 2, 23, 110, 3, 2, 2, 2, 25, 116, 3, 2, 2, 2, 27, 121, 3, 2, 2,
	2, 29, 124, 3, 2, 2, 2, 31, 128, 3, 2, 2, 2, 33, 135, 3, 2, 2, 2, 35, 140,
	3, 2, 2, 2, 37, 144, 3, 2, 2, 2, 39, 152, 3, 2, 2, 2, 41, 159, 3, 2, 2,
	2, 43, 163, 3, 2, 2, 2, 45, 169, 3, 2, 2, 2, 47, 177, 3, 2, 2, 2, 49, 179,
	3, 2, 2, 2, 51, 197, 3, 2, 2, 2, 53, 201, 3, 2, 2, 2, 55, 203, 3, 2, 2,
	2, 57, 59, 9, 2, 2, 2, 58, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 58,
	3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 63, 8, 2, 2, 2,
	63, 4, 3, 2, 2, 2, 64, 65, 7, 49, 2, 2, 65, 66, 7, 44, 2, 2, 66, 70, 3,
	2, 2, 2, 67, 69, 11, 2, 2, 2, 68, 67, 3, 2, 2, 2, 69, 72, 3, 2, 2, 2, 70,
	71, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 71, 73, 3, 2, 2, 2, 72, 70, 3, 2, 2,
	2, 73, 74, 7, 44, 2, 2, 74, 75, 7, 49, 2, 2, 75, 76, 3, 2, 2, 2, 76, 77,
	8, 3, 3, 2, 77, 6, 3, 2, 2, 2, 78, 79, 7, 49, 2, 2, 79, 80, 7, 49, 2, 2,
	80, 84, 3, 2, 2, 2, 81, 83, 10, 3, 2, 2, 82, 81, 3, 2, 2, 2, 83, 86, 3,
	2, 2, 2, 84, 82, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 87, 3, 2, 2, 2, 86,
	84, 3, 2, 2, 2, 87, 88, 8, 4, 3, 2, 88, 8, 3, 2, 2, 2, 89, 91, 9, 3, 2,
	2, 90, 89, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 92, 93,
	3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 95, 8, 5, 3, 2, 95, 10, 3, 2, 2, 2,
	96, 97, 7, 125, 2, 2, 97, 12, 3, 2, 2, 2, 98, 99, 7, 127, 2, 2, 99, 14,
	3, 2, 2, 2, 100, 101, 7, 46, 2, 2, 101, 16, 3, 2, 2, 2, 102, 103, 7, 36,
	2, 2, 103, 18, 3, 2, 2, 2, 104, 105, 7, 61, 2, 2, 105, 20, 3, 2, 2, 2,
	106, 107, 7, 99, 2, 2, 107, 108, 7, 112, 2, 2, 108, 109, 7, 123, 2, 2,
	109, 22, 3, 2, 2, 2, 110, 111, 7, 102, 2, 2, 111, 112, 7, 103, 2, 2, 112,
	113, 7, 100, 2, 2, 113, 114, 7, 119, 2, 2, 114, 115, 7, 105, 2, 2, 115,
	24, 3, 2, 2, 2, 116, 117, 7, 114, 2, 2, 117, 118, 7, 99, 2, 2, 118, 119,
	7, 117, 2, 2, 119, 120, 7, 117, 2, 2, 120, 26, 3, 2, 2, 2, 121, 122, 7,
	113, 2, 2, 122, 123, 7, 112, 2, 2, 123, 28, 3, 2, 2, 2, 124, 125, 7, 116,
	2, 2, 125, 126, 7, 114, 2, 2, 126, 127, 7, 101, 2, 2, 127, 30, 3, 2, 2,
	2, 128, 129, 7, 102, 2, 2, 129, 130, 7, 99, 2, 2, 130, 131, 7, 103, 2,
	2, 131, 132, 7, 111, 2, 2, 132, 133, 7, 113, 2, 2, 133, 134, 7, 112, 2,
	2, 134, 32, 3, 2, 2, 2, 135, 136, 7, 114, 2, 2, 136, 137, 7, 99, 2, 2,
	137, 138, 7, 118, 2, 2, 138, 139, 7, 106, 2, 2, 139, 34, 3, 2, 2, 2, 140,
	141, 7, 119, 2, 2, 141, 142, 7, 117, 2, 2, 142, 143, 7, 103, 2, 2, 143,
	36, 3, 2, 2, 2, 144, 145, 7, 105, 2, 2, 145, 146, 7, 113, 2, 2, 146, 147,
	7, 47, 2, 2, 147, 148, 7, 105, 2, 2, 148, 149, 7, 116, 2, 2, 149, 150,
	7, 114, 2, 2, 150, 151, 7, 101, 2, 2, 151, 38, 3, 2, 2, 2, 152, 153, 7,
	105, 2, 2, 153, 154, 7, 113, 2, 2, 154, 155, 7, 47, 2, 2, 155, 156, 7,
	118, 2, 2, 156, 157, 7, 110, 2, 2, 157, 158, 7, 117, 2, 2, 158, 40, 3,
	2, 2, 2, 159, 160, 7, 114, 2, 2, 160, 161, 7, 103, 2, 2, 161, 162, 7, 111,
	2, 2, 162, 42, 3, 2, 2, 2, 163, 165, 5, 17, 9, 2, 164, 166, 5, 45, 23,
	2, 165, 164, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167,
	168, 5, 17, 9, 2, 168, 44, 3, 2, 2, 2, 169, 174, 5, 47, 24, 2, 170, 173,
	5, 47, 24, 2, 171, 173, 5, 49, 25, 2, 172, 170, 3, 2, 2, 2, 172, 171, 3,
	2, 2, 2, 173, 176, 3, 2, 2, 2, 174, 172, 3, 2, 2, 2, 174, 175, 3, 2, 2,
	2, 175, 46, 3, 2, 2, 2, 176, 174, 3, 2, 2, 2, 177, 178, 9, 4, 2, 2, 178,
	48, 3, 2, 2, 2, 179, 180, 9, 5, 2, 2, 180, 50, 3, 2, 2, 2, 181, 185, 7,
	41, 2, 2, 182, 184, 5, 53, 27, 2, 183, 182, 3, 2, 2, 2, 184, 187, 3, 2,
	2, 2, 185, 186, 3, 2, 2, 2, 185, 183, 3, 2, 2, 2, 186, 188, 3, 2, 2, 2,
	187, 185, 3, 2, 2, 2, 188, 198, 7, 41, 2, 2, 189, 193, 7, 36, 2, 2, 190,
	192, 5, 53, 27, 2, 191, 190, 3, 2, 2, 2, 192, 195, 3, 2, 2, 2, 193, 194,
	3, 2, 2, 2, 193, 191, 3, 2, 2, 2, 194, 196, 3, 2, 2, 2, 195, 193, 3, 2,
	2, 2, 196, 198, 7, 36, 2, 2, 197, 181, 3, 2, 2, 2, 197, 189, 3, 2, 2, 2,
	198, 52, 3, 2, 2, 2, 199, 202, 5, 55, 28, 2, 200, 202, 10, 6, 2, 2, 201,
	199, 3, 2, 2, 2, 201, 200, 3, 2, 2, 2, 202, 54, 3, 2, 2, 2, 203, 204, 7,
	94, 2, 2, 204, 205, 9, 7, 2, 2, 205, 56, 3, 2, 2, 2, 14, 2, 60, 70, 84,
	92, 165, 172, 174, 185, 193, 197, 201, 4, 8, 2, 2, 2, 3, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "", "", "", "'{'", "'}'", "','", "'\"'", "';'", "'any'", "'debug'",
	"'pass'", "'on'", "'rpc'", "'daemon'", "'path'", "'use'", "'go-grpc'",
	"'go-tls'", "'pem'",
}

var lexerSymbolicNames = []string{
	"", "WS", "COMMENT", "LINE_COMMENT", "TERMINATOR", "LCURLY", "RCURLY",
	"COMMA", "QUOTE", "SEMI", "KWANY", "KWDBG", "KWPASS", "KWON", "KWRPC",
	"KWDAEMON", "KWPATH", "KWUSE", "KWGOGRPC", "KWGOTLS", "KWPEM", "HNAME",
	"IDENTIFIER", "PATH",
}

var lexerRuleNames = []string{
	"WS", "COMMENT", "LINE_COMMENT", "TERMINATOR", "LCURLY", "RCURLY", "COMMA",
	"QUOTE", "SEMI", "KWANY", "KWDBG", "KWPASS", "KWON", "KWRPC", "KWDAEMON",
	"KWPATH", "KWUSE", "KWGOGRPC", "KWGOTLS", "KWPEM", "HNAME", "IDENTIFIER",
	"LETTER", "DECIMAL_DIGIT", "PATH", "CHAR_VALUE", "CHAR_ESCAPE",
}

type pkscLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewpkscLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *pkscLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewpkscLexer(input antlr.CharStream) *pkscLexer {
	l := new(pkscLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
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
