// Code generated from MsgLexer.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 11, 177,
	8, 1, 8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6,
	4, 7, 9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12,
	9, 12, 4, 13, 9, 13, 4, 14, 9, 14, 3, 2, 6, 2, 32, 10, 2, 13, 2, 14, 2,
	33, 3, 3, 3, 3, 5, 3, 38, 10, 3, 3, 3, 6, 3, 41, 10, 3, 13, 3, 14, 3, 42,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 110, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 7, 3, 7, 7, 7, 119, 10, 7, 12, 7, 14, 7, 122, 11, 7, 3, 8, 3,
	8, 3, 8, 5, 8, 127, 10, 8, 3, 9, 3, 9, 3, 10, 6, 10, 132, 10, 10, 13, 10,
	14, 10, 133, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 7, 11, 141, 10, 11, 12,
	11, 14, 11, 144, 11, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12,
	152, 10, 12, 3, 13, 3, 13, 5, 13, 156, 10, 13, 3, 13, 3, 13, 7, 13, 160,
	10, 13, 12, 13, 14, 13, 163, 11, 13, 5, 13, 165, 10, 13, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 176, 10, 14, 3,
	142, 2, 15, 4, 3, 6, 4, 8, 5, 10, 6, 12, 7, 14, 8, 16, 9, 18, 10, 20, 11,
	22, 2, 24, 2, 26, 2, 28, 2, 4, 2, 3, 9, 3, 2, 62, 62, 4, 2, 11, 11, 34,
	34, 5, 2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124,
	5, 2, 11, 12, 15, 15, 34, 34, 3, 2, 51, 59, 3, 2, 50, 59, 2, 195, 2, 4,
	3, 2, 2, 2, 2, 6, 3, 2, 2, 2, 2, 8, 3, 2, 2, 2, 3, 10, 3, 2, 2, 2, 3, 12,
	3, 2, 2, 2, 3, 14, 3, 2, 2, 2, 3, 16, 3, 2, 2, 2, 3, 18, 3, 2, 2, 2, 3,
	20, 3, 2, 2, 2, 4, 31, 3, 2, 2, 2, 6, 40, 3, 2, 2, 2, 8, 44, 3, 2, 2, 2,
	10, 109, 3, 2, 2, 2, 12, 111, 3, 2, 2, 2, 14, 116, 3, 2, 2, 2, 16, 126,
	3, 2, 2, 2, 18, 128, 3, 2, 2, 2, 20, 131, 3, 2, 2, 2, 22, 137, 3, 2, 2,
	2, 24, 151, 3, 2, 2, 2, 26, 164, 3, 2, 2, 2, 28, 175, 3, 2, 2, 2, 30, 32,
	10, 2, 2, 2, 31, 30, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 31, 3, 2, 2, 2,
	33, 34, 3, 2, 2, 2, 34, 5, 3, 2, 2, 2, 35, 41, 9, 3, 2, 2, 36, 38, 7, 15,
	2, 2, 37, 36, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 41,
	7, 12, 2, 2, 40, 35, 3, 2, 2, 2, 40, 37, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2,
	42, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 7, 3, 2, 2, 2, 44, 45, 7, 62,
	2, 2, 45, 46, 7, 122, 2, 2, 46, 47, 7, 47, 2, 2, 47, 48, 3, 2, 2, 2, 48,
	49, 8, 4, 2, 2, 49, 9, 3, 2, 2, 2, 50, 51, 7, 111, 2, 2, 51, 52, 7, 103,
	2, 2, 52, 53, 7, 112, 2, 2, 53, 54, 7, 118, 2, 2, 54, 55, 7, 107, 2, 2,
	55, 56, 7, 113, 2, 2, 56, 110, 7, 112, 2, 2, 57, 58, 7, 107, 2, 2, 58,
	59, 7, 111, 2, 2, 59, 60, 7, 99, 2, 2, 60, 61, 7, 105, 2, 2, 61, 110, 7,
	103, 2, 2, 62, 63, 7, 120, 2, 2, 63, 64, 7, 107, 2, 2, 64, 65, 7, 102,
	2, 2, 65, 66, 7, 103, 2, 2, 66, 110, 7, 113, 2, 2, 67, 68, 7, 104, 2, 2,
	68, 69, 7, 107, 2, 2, 69, 70, 7, 110, 2, 2, 70, 110, 7, 103, 2, 2, 71,
	72, 7, 99, 2, 2, 72, 73, 7, 119, 2, 2, 73, 74, 7, 102, 2, 2, 74, 75, 7,
	107, 2, 2, 75, 110, 7, 113, 2, 2, 76, 77, 7, 110, 2, 2, 77, 78, 7, 107,
	2, 2, 78, 79, 7, 112, 2, 2, 79, 110, 7, 109, 2, 2, 80, 81, 7, 117, 2, 2,
	81, 82, 7, 118, 2, 2, 82, 83, 7, 107, 2, 2, 83, 84, 7, 101, 2, 2, 84, 85,
	7, 109, 2, 2, 85, 86, 7, 103, 2, 2, 86, 110, 7, 116, 2, 2, 87, 88, 7, 101,
	2, 2, 88, 89, 7, 113, 2, 2, 89, 90, 7, 112, 2, 2, 90, 91, 7, 118, 2, 2,
	91, 92, 7, 99, 2, 2, 92, 93, 7, 101, 2, 2, 93, 110, 7, 118, 2, 2, 94, 95,
	7, 102, 2, 2, 95, 96, 7, 113, 2, 2, 96, 97, 7, 101, 2, 2, 97, 98, 7, 119,
	2, 2, 98, 99, 7, 111, 2, 2, 99, 100, 7, 103, 2, 2, 100, 101, 7, 112, 2,
	2, 101, 110, 7, 118, 2, 2, 102, 103, 7, 111, 2, 2, 103, 104, 7, 103, 2,
	2, 104, 105, 7, 103, 2, 2, 105, 106, 7, 118, 2, 2, 106, 107, 7, 107, 2,
	2, 107, 108, 7, 112, 2, 2, 108, 110, 7, 105, 2, 2, 109, 50, 3, 2, 2, 2,
	109, 57, 3, 2, 2, 2, 109, 62, 3, 2, 2, 2, 109, 67, 3, 2, 2, 2, 109, 71,
	3, 2, 2, 2, 109, 76, 3, 2, 2, 2, 109, 80, 3, 2, 2, 2, 109, 87, 3, 2, 2,
	2, 109, 94, 3, 2, 2, 2, 109, 102, 3, 2, 2, 2, 110, 11, 3, 2, 2, 2, 111,
	112, 7, 49, 2, 2, 112, 113, 7, 64, 2, 2, 113, 114, 3, 2, 2, 2, 114, 115,
	8, 6, 3, 2, 115, 13, 3, 2, 2, 2, 116, 120, 9, 4, 2, 2, 117, 119, 9, 5,
	2, 2, 118, 117, 3, 2, 2, 2, 119, 122, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2,
	120, 121, 3, 2, 2, 2, 121, 15, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 123, 127,
	5, 22, 11, 2, 124, 127, 5, 26, 13, 2, 125, 127, 5, 28, 14, 2, 126, 123,
	3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 126, 125, 3, 2, 2, 2, 127, 17, 3, 2,
	2, 2, 128, 129, 7, 63, 2, 2, 129, 19, 3, 2, 2, 2, 130, 132, 9, 6, 2, 2,
	131, 130, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 131, 3, 2, 2, 2, 133,
	134, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 136, 8, 10, 4, 2, 136, 21,
	3, 2, 2, 2, 137, 142, 7, 36, 2, 2, 138, 141, 5, 24, 12, 2, 139, 141, 11,
	2, 2, 2, 140, 138, 3, 2, 2, 2, 140, 139, 3, 2, 2, 2, 141, 144, 3, 2, 2,
	2, 142, 143, 3, 2, 2, 2, 142, 140, 3, 2, 2, 2, 143, 145, 3, 2, 2, 2, 144,
	142, 3, 2, 2, 2, 145, 146, 7, 36, 2, 2, 146, 23, 3, 2, 2, 2, 147, 148,
	7, 94, 2, 2, 148, 152, 7, 36, 2, 2, 149, 150, 7, 94, 2, 2, 150, 152, 7,
	94, 2, 2, 151, 147, 3, 2, 2, 2, 151, 149, 3, 2, 2, 2, 152, 25, 3, 2, 2,
	2, 153, 165, 7, 50, 2, 2, 154, 156, 7, 47, 2, 2, 155, 154, 3, 2, 2, 2,
	155, 156, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 161, 9, 7, 2, 2, 158,
	160, 9, 8, 2, 2, 159, 158, 3, 2, 2, 2, 160, 163, 3, 2, 2, 2, 161, 159,
	3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162, 165, 3, 2, 2, 2, 163, 161, 3, 2,
	2, 2, 164, 153, 3, 2, 2, 2, 164, 155, 3, 2, 2, 2, 165, 27, 3, 2, 2, 2,
	166, 167, 7, 118, 2, 2, 167, 168, 7, 116, 2, 2, 168, 169, 7, 119, 2, 2,
	169, 176, 7, 103, 2, 2, 170, 171, 7, 104, 2, 2, 171, 172, 7, 99, 2, 2,
	172, 173, 7, 110, 2, 2, 173, 174, 7, 117, 2, 2, 174, 176, 7, 103, 2, 2,
	175, 166, 3, 2, 2, 2, 175, 170, 3, 2, 2, 2, 176, 29, 3, 2, 2, 2, 19, 2,
	3, 33, 37, 40, 42, 109, 120, 126, 133, 140, 142, 151, 155, 161, 164, 175,
	5, 7, 3, 2, 6, 2, 2, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE", "TAG",
}

var lexerLiteralNames = []string{
	"", "", "", "'<x-'", "", "'/>'", "", "", "'='",
}

var lexerSymbolicNames = []string{
	"", "TEXT", "SEA_WS", "TAG_OPEN", "TAG_TYPE", "TAG_CLOSE", "ATTR_NAME",
	"ATTR_VALUE", "EQ", "WS",
}

var lexerRuleNames = []string{
	"TEXT", "SEA_WS", "TAG_OPEN", "TAG_TYPE", "TAG_CLOSE", "ATTR_NAME", "ATTR_VALUE",
	"EQ", "WS", "STRING", "ESC", "INT", "BOOL",
}

type MsgLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewMsgLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *MsgLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewMsgLexer(input antlr.CharStream) *MsgLexer {
	l := new(MsgLexer)
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
	l.GrammarFileName = "MsgLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MsgLexer tokens.
const (
	MsgLexerTEXT       = 1
	MsgLexerSEA_WS     = 2
	MsgLexerTAG_OPEN   = 3
	MsgLexerTAG_TYPE   = 4
	MsgLexerTAG_CLOSE  = 5
	MsgLexerATTR_NAME  = 6
	MsgLexerATTR_VALUE = 7
	MsgLexerEQ         = 8
	MsgLexerWS         = 9
)

// MsgLexerTAG is the MsgLexer mode.
const MsgLexerTAG = 1
