// Code generated from HCLLikeDSL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

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

type HCLLikeDSLLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var HCLLikeDSLLexerLexerStaticData struct {
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

func hcllikedsllexerLexerInit() {
	staticData := &HCLLikeDSLLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'HCLCONFIG'", "'{'", "'}'", "'='", "'import'", "'repeated'", "'['",
		"','", "']'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "STRING", "NUMBER", "BOOLEAN",
		"IDENTIFIER", "WS", "COMMENT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"STRING", "NUMBER", "BOOLEAN", "IDENTIFIER", "WS", "COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 15, 125, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 9, 1, 9, 5, 9, 72, 8, 9, 10, 9, 12, 9, 75, 9, 9, 1, 9, 1, 9, 1,
		10, 4, 10, 80, 8, 10, 11, 10, 12, 10, 81, 1, 10, 1, 10, 4, 10, 86, 8, 10,
		11, 10, 12, 10, 87, 3, 10, 90, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 101, 8, 11, 1, 12, 1, 12, 5, 12, 105,
		8, 12, 10, 12, 12, 12, 108, 9, 12, 1, 13, 4, 13, 111, 8, 13, 11, 13, 12,
		13, 112, 1, 13, 1, 13, 1, 14, 1, 14, 5, 14, 119, 8, 14, 10, 14, 12, 14,
		122, 9, 14, 1, 14, 1, 14, 0, 0, 15, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6,
		13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 1,
		0, 6, 3, 0, 10, 10, 13, 13, 34, 34, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95,
		97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32,
		32, 2, 0, 10, 10, 13, 13, 132, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5,
		1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 1, 31, 1, 0, 0, 0, 3, 41, 1, 0, 0, 0, 5, 43, 1, 0, 0,
		0, 7, 45, 1, 0, 0, 0, 9, 47, 1, 0, 0, 0, 11, 54, 1, 0, 0, 0, 13, 63, 1,
		0, 0, 0, 15, 65, 1, 0, 0, 0, 17, 67, 1, 0, 0, 0, 19, 69, 1, 0, 0, 0, 21,
		79, 1, 0, 0, 0, 23, 100, 1, 0, 0, 0, 25, 102, 1, 0, 0, 0, 27, 110, 1, 0,
		0, 0, 29, 116, 1, 0, 0, 0, 31, 32, 5, 72, 0, 0, 32, 33, 5, 67, 0, 0, 33,
		34, 5, 76, 0, 0, 34, 35, 5, 67, 0, 0, 35, 36, 5, 79, 0, 0, 36, 37, 5, 78,
		0, 0, 37, 38, 5, 70, 0, 0, 38, 39, 5, 73, 0, 0, 39, 40, 5, 71, 0, 0, 40,
		2, 1, 0, 0, 0, 41, 42, 5, 123, 0, 0, 42, 4, 1, 0, 0, 0, 43, 44, 5, 125,
		0, 0, 44, 6, 1, 0, 0, 0, 45, 46, 5, 61, 0, 0, 46, 8, 1, 0, 0, 0, 47, 48,
		5, 105, 0, 0, 48, 49, 5, 109, 0, 0, 49, 50, 5, 112, 0, 0, 50, 51, 5, 111,
		0, 0, 51, 52, 5, 114, 0, 0, 52, 53, 5, 116, 0, 0, 53, 10, 1, 0, 0, 0, 54,
		55, 5, 114, 0, 0, 55, 56, 5, 101, 0, 0, 56, 57, 5, 112, 0, 0, 57, 58, 5,
		101, 0, 0, 58, 59, 5, 97, 0, 0, 59, 60, 5, 116, 0, 0, 60, 61, 5, 101, 0,
		0, 61, 62, 5, 100, 0, 0, 62, 12, 1, 0, 0, 0, 63, 64, 5, 91, 0, 0, 64, 14,
		1, 0, 0, 0, 65, 66, 5, 44, 0, 0, 66, 16, 1, 0, 0, 0, 67, 68, 5, 93, 0,
		0, 68, 18, 1, 0, 0, 0, 69, 73, 5, 34, 0, 0, 70, 72, 8, 0, 0, 0, 71, 70,
		1, 0, 0, 0, 72, 75, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0,
		74, 76, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 76, 77, 5, 34, 0, 0, 77, 20, 1,
		0, 0, 0, 78, 80, 7, 1, 0, 0, 79, 78, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81,
		79, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 89, 1, 0, 0, 0, 83, 85, 5, 46,
		0, 0, 84, 86, 7, 1, 0, 0, 85, 84, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 85,
		1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 90, 1, 0, 0, 0, 89, 83, 1, 0, 0, 0,
		89, 90, 1, 0, 0, 0, 90, 22, 1, 0, 0, 0, 91, 92, 5, 116, 0, 0, 92, 93, 5,
		114, 0, 0, 93, 94, 5, 117, 0, 0, 94, 101, 5, 101, 0, 0, 95, 96, 5, 102,
		0, 0, 96, 97, 5, 97, 0, 0, 97, 98, 5, 108, 0, 0, 98, 99, 5, 115, 0, 0,
		99, 101, 5, 101, 0, 0, 100, 91, 1, 0, 0, 0, 100, 95, 1, 0, 0, 0, 101, 24,
		1, 0, 0, 0, 102, 106, 7, 2, 0, 0, 103, 105, 7, 3, 0, 0, 104, 103, 1, 0,
		0, 0, 105, 108, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0,
		107, 26, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 109, 111, 7, 4, 0, 0, 110, 109,
		1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112, 113, 1, 0,
		0, 0, 113, 114, 1, 0, 0, 0, 114, 115, 6, 13, 0, 0, 115, 28, 1, 0, 0, 0,
		116, 120, 5, 35, 0, 0, 117, 119, 8, 5, 0, 0, 118, 117, 1, 0, 0, 0, 119,
		122, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 123,
		1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 123, 124, 6, 14, 0, 0, 124, 30, 1, 0,
		0, 0, 9, 0, 73, 81, 87, 89, 100, 106, 112, 120, 1, 6, 0, 0,
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

// HCLLikeDSLLexerInit initializes any static state used to implement HCLLikeDSLLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewHCLLikeDSLLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func HCLLikeDSLLexerInit() {
	staticData := &HCLLikeDSLLexerLexerStaticData
	staticData.once.Do(hcllikedsllexerLexerInit)
}

// NewHCLLikeDSLLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewHCLLikeDSLLexer(input antlr.CharStream) *HCLLikeDSLLexer {
	HCLLikeDSLLexerInit()
	l := new(HCLLikeDSLLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &HCLLikeDSLLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "HCLLikeDSL.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// HCLLikeDSLLexer tokens.
const (
	HCLLikeDSLLexerT__0       = 1
	HCLLikeDSLLexerT__1       = 2
	HCLLikeDSLLexerT__2       = 3
	HCLLikeDSLLexerT__3       = 4
	HCLLikeDSLLexerT__4       = 5
	HCLLikeDSLLexerT__5       = 6
	HCLLikeDSLLexerT__6       = 7
	HCLLikeDSLLexerT__7       = 8
	HCLLikeDSLLexerT__8       = 9
	HCLLikeDSLLexerSTRING     = 10
	HCLLikeDSLLexerNUMBER     = 11
	HCLLikeDSLLexerBOOLEAN    = 12
	HCLLikeDSLLexerIDENTIFIER = 13
	HCLLikeDSLLexerWS         = 14
	HCLLikeDSLLexerCOMMENT    = 15
)
