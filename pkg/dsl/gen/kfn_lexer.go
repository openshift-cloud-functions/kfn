// Code generated from Kfn.g4 by ANTLR 4.7.1. DO NOT EDIT.

package gen

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 14, 106,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9,
	3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 7, 11, 66, 10, 11, 12, 11, 14,
	11, 69, 11, 11, 3, 11, 3, 11, 3, 12, 7, 12, 74, 10, 12, 12, 12, 14, 12,
	77, 11, 12, 3, 13, 3, 13, 3, 13, 5, 13, 82, 10, 13, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3,
	17, 5, 17, 98, 10, 17, 3, 18, 6, 18, 101, 10, 18, 13, 18, 14, 18, 102,
	3, 18, 3, 18, 2, 2, 19, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17,
	10, 19, 11, 21, 12, 23, 2, 25, 2, 27, 2, 29, 2, 31, 2, 33, 13, 35, 14,
	3, 2, 7, 7, 2, 47, 47, 50, 59, 67, 92, 97, 97, 99, 124, 10, 2, 36, 36,
	49, 49, 94, 94, 100, 100, 104, 104, 112, 112, 116, 116, 118, 118, 5, 2,
	50, 59, 67, 72, 99, 104, 5, 2, 2, 33, 36, 36, 94, 94, 4, 2, 11, 11, 34,
	34, 2, 106, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9,
	3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2,
	17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2,
	2, 35, 3, 2, 2, 2, 3, 37, 3, 2, 2, 2, 5, 39, 3, 2, 2, 2, 7, 41, 3, 2, 2,
	2, 9, 43, 3, 2, 2, 2, 11, 45, 3, 2, 2, 2, 13, 49, 3, 2, 2, 2, 15, 51, 3,
	2, 2, 2, 17, 56, 3, 2, 2, 2, 19, 59, 3, 2, 2, 2, 21, 62, 3, 2, 2, 2, 23,
	75, 3, 2, 2, 2, 25, 78, 3, 2, 2, 2, 27, 83, 3, 2, 2, 2, 29, 89, 3, 2, 2,
	2, 31, 91, 3, 2, 2, 2, 33, 97, 3, 2, 2, 2, 35, 100, 3, 2, 2, 2, 37, 38,
	7, 61, 2, 2, 38, 4, 3, 2, 2, 2, 39, 40, 7, 12, 2, 2, 40, 6, 3, 2, 2, 2,
	41, 42, 7, 60, 2, 2, 42, 8, 3, 2, 2, 2, 43, 44, 7, 46, 2, 2, 44, 10, 3,
	2, 2, 2, 45, 46, 7, 99, 2, 2, 46, 47, 7, 112, 2, 2, 47, 48, 7, 102, 2,
	2, 48, 12, 3, 2, 2, 2, 49, 50, 7, 63, 2, 2, 50, 14, 3, 2, 2, 2, 51, 52,
	7, 121, 2, 2, 52, 53, 7, 107, 2, 2, 53, 54, 7, 118, 2, 2, 54, 55, 7, 106,
	2, 2, 55, 16, 3, 2, 2, 2, 56, 57, 4, 99, 124, 2, 57, 58, 5, 23, 12, 2,
	58, 18, 3, 2, 2, 2, 59, 60, 4, 67, 92, 2, 60, 61, 5, 23, 12, 2, 61, 20,
	3, 2, 2, 2, 62, 67, 7, 36, 2, 2, 63, 66, 5, 25, 13, 2, 64, 66, 5, 31, 16,
	2, 65, 63, 3, 2, 2, 2, 65, 64, 3, 2, 2, 2, 66, 69, 3, 2, 2, 2, 67, 65,
	3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 70, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2,
	70, 71, 7, 36, 2, 2, 71, 22, 3, 2, 2, 2, 72, 74, 9, 2, 2, 2, 73, 72, 3,
	2, 2, 2, 74, 77, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76,
	24, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 78, 81, 7, 94, 2, 2, 79, 82, 9, 3,
	2, 2, 80, 82, 5, 27, 14, 2, 81, 79, 3, 2, 2, 2, 81, 80, 3, 2, 2, 2, 82,
	26, 3, 2, 2, 2, 83, 84, 7, 119, 2, 2, 84, 85, 5, 29, 15, 2, 85, 86, 5,
	29, 15, 2, 86, 87, 5, 29, 15, 2, 87, 88, 5, 29, 15, 2, 88, 28, 3, 2, 2,
	2, 89, 90, 9, 4, 2, 2, 90, 30, 3, 2, 2, 2, 91, 92, 10, 5, 2, 2, 92, 32,
	3, 2, 2, 2, 93, 94, 7, 63, 2, 2, 94, 98, 7, 64, 2, 2, 95, 96, 7, 47, 2,
	2, 96, 98, 7, 64, 2, 2, 97, 93, 3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 98, 34,
	3, 2, 2, 2, 99, 101, 9, 6, 2, 2, 100, 99, 3, 2, 2, 2, 101, 102, 3, 2, 2,
	2, 102, 100, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104,
	105, 8, 18, 2, 2, 105, 36, 3, 2, 2, 2, 9, 2, 65, 67, 75, 81, 97, 102, 3,
	8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "';'", "'\n'", "':'", "','", "'and'", "'='", "'with'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "COMPONENT_OPTION_SEP", "UPPER_IDENT", "LOWER_IDENT",
	"STRING_LITERAL", "ARROW", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "COMPONENT_OPTION_SEP",
	"UPPER_IDENT", "LOWER_IDENT", "STRING_LITERAL", "IDENT_TAIL", "ESC", "UNICODE",
	"HEX", "SAFECODEPOINT", "ARROW", "WS",
}

type KfnLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewKfnLexer(input antlr.CharStream) *KfnLexer {

	l := new(KfnLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Kfn.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// KfnLexer tokens.
const (
	KfnLexerT__0                 = 1
	KfnLexerT__1                 = 2
	KfnLexerT__2                 = 3
	KfnLexerT__3                 = 4
	KfnLexerT__4                 = 5
	KfnLexerT__5                 = 6
	KfnLexerCOMPONENT_OPTION_SEP = 7
	KfnLexerUPPER_IDENT          = 8
	KfnLexerLOWER_IDENT          = 9
	KfnLexerSTRING_LITERAL       = 10
	KfnLexerARROW                = 11
	KfnLexerWS                   = 12
)
