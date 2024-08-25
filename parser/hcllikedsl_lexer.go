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
		"", "'import'", "'{'", "'}'", "'='", "'repeated'", "'['", "','", "']'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "STRING", "NUMBER", "BOOLEAN", "IDENTIFIER",
		"WS", "COMMENT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "STRING",
		"NUMBER", "BOOLEAN", "IDENTIFIER", "WS", "COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 14, 113, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 8, 1, 8, 5, 8, 60, 8, 8, 10, 8, 12, 8, 63, 9, 8, 1, 8, 1, 8, 1,
		9, 4, 9, 68, 8, 9, 11, 9, 12, 9, 69, 1, 9, 1, 9, 4, 9, 74, 8, 9, 11, 9,
		12, 9, 75, 3, 9, 78, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 3, 10, 89, 8, 10, 1, 11, 1, 11, 5, 11, 93, 8, 11, 10,
		11, 12, 11, 96, 9, 11, 1, 12, 4, 12, 99, 8, 12, 11, 12, 12, 12, 100, 1,
		12, 1, 12, 1, 13, 1, 13, 5, 13, 107, 8, 13, 10, 13, 12, 13, 110, 9, 13,
		1, 13, 1, 13, 0, 0, 14, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 1, 0, 6, 3, 0, 10, 10,
		13, 13, 34, 34, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48,
		57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10, 10,
		13, 13, 120, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7,
		1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0,
		15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0,
		0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 1, 29, 1, 0, 0,
		0, 3, 36, 1, 0, 0, 0, 5, 38, 1, 0, 0, 0, 7, 40, 1, 0, 0, 0, 9, 42, 1, 0,
		0, 0, 11, 51, 1, 0, 0, 0, 13, 53, 1, 0, 0, 0, 15, 55, 1, 0, 0, 0, 17, 57,
		1, 0, 0, 0, 19, 67, 1, 0, 0, 0, 21, 88, 1, 0, 0, 0, 23, 90, 1, 0, 0, 0,
		25, 98, 1, 0, 0, 0, 27, 104, 1, 0, 0, 0, 29, 30, 5, 105, 0, 0, 30, 31,
		5, 109, 0, 0, 31, 32, 5, 112, 0, 0, 32, 33, 5, 111, 0, 0, 33, 34, 5, 114,
		0, 0, 34, 35, 5, 116, 0, 0, 35, 2, 1, 0, 0, 0, 36, 37, 5, 123, 0, 0, 37,
		4, 1, 0, 0, 0, 38, 39, 5, 125, 0, 0, 39, 6, 1, 0, 0, 0, 40, 41, 5, 61,
		0, 0, 41, 8, 1, 0, 0, 0, 42, 43, 5, 114, 0, 0, 43, 44, 5, 101, 0, 0, 44,
		45, 5, 112, 0, 0, 45, 46, 5, 101, 0, 0, 46, 47, 5, 97, 0, 0, 47, 48, 5,
		116, 0, 0, 48, 49, 5, 101, 0, 0, 49, 50, 5, 100, 0, 0, 50, 10, 1, 0, 0,
		0, 51, 52, 5, 91, 0, 0, 52, 12, 1, 0, 0, 0, 53, 54, 5, 44, 0, 0, 54, 14,
		1, 0, 0, 0, 55, 56, 5, 93, 0, 0, 56, 16, 1, 0, 0, 0, 57, 61, 5, 34, 0,
		0, 58, 60, 8, 0, 0, 0, 59, 58, 1, 0, 0, 0, 60, 63, 1, 0, 0, 0, 61, 59,
		1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 64, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0,
		64, 65, 5, 34, 0, 0, 65, 18, 1, 0, 0, 0, 66, 68, 7, 1, 0, 0, 67, 66, 1,
		0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70,
		77, 1, 0, 0, 0, 71, 73, 5, 46, 0, 0, 72, 74, 7, 1, 0, 0, 73, 72, 1, 0,
		0, 0, 74, 75, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 78,
		1, 0, 0, 0, 77, 71, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 20, 1, 0, 0, 0,
		79, 80, 5, 116, 0, 0, 80, 81, 5, 114, 0, 0, 81, 82, 5, 117, 0, 0, 82, 89,
		5, 101, 0, 0, 83, 84, 5, 102, 0, 0, 84, 85, 5, 97, 0, 0, 85, 86, 5, 108,
		0, 0, 86, 87, 5, 115, 0, 0, 87, 89, 5, 101, 0, 0, 88, 79, 1, 0, 0, 0, 88,
		83, 1, 0, 0, 0, 89, 22, 1, 0, 0, 0, 90, 94, 7, 2, 0, 0, 91, 93, 7, 3, 0,
		0, 92, 91, 1, 0, 0, 0, 93, 96, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 94, 95,
		1, 0, 0, 0, 95, 24, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 97, 99, 7, 4, 0, 0,
		98, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 100, 101,
		1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 103, 6, 12, 0, 0, 103, 26, 1, 0,
		0, 0, 104, 108, 5, 35, 0, 0, 105, 107, 8, 5, 0, 0, 106, 105, 1, 0, 0, 0,
		107, 110, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109,
		111, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 111, 112, 6, 13, 0, 0, 112, 28,
		1, 0, 0, 0, 9, 0, 61, 69, 75, 77, 88, 94, 100, 108, 1, 6, 0, 0,
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
	HCLLikeDSLLexerSTRING     = 9
	HCLLikeDSLLexerNUMBER     = 10
	HCLLikeDSLLexerBOOLEAN    = 11
	HCLLikeDSLLexerIDENTIFIER = 12
	HCLLikeDSLLexerWS         = 13
	HCLLikeDSLLexerCOMMENT    = 14
)
