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
		"", "'HCLCONFIG'", "'{'", "'}'", "'='", "'repeated'", "'optional'",
		"'optional repeated'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "STRING", "NUMBER", "BOOLEAN", "IDENTIFIER",
		"WS", "COMMENT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "STRING", "NUMBER",
		"BOOLEAN", "IDENTIFIER", "WS", "COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 13, 135, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1,
		7, 5, 7, 82, 8, 7, 10, 7, 12, 7, 85, 9, 7, 1, 7, 1, 7, 1, 8, 4, 8, 90,
		8, 8, 11, 8, 12, 8, 91, 1, 8, 1, 8, 4, 8, 96, 8, 8, 11, 8, 12, 8, 97, 3,
		8, 100, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3,
		9, 111, 8, 9, 1, 10, 1, 10, 5, 10, 115, 8, 10, 10, 10, 12, 10, 118, 9,
		10, 1, 11, 4, 11, 121, 8, 11, 11, 11, 12, 11, 122, 1, 11, 1, 11, 1, 12,
		1, 12, 5, 12, 129, 8, 12, 10, 12, 12, 12, 132, 9, 12, 1, 12, 1, 12, 0,
		0, 13, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
		21, 11, 23, 12, 25, 13, 1, 0, 6, 3, 0, 10, 10, 13, 13, 34, 34, 1, 0, 48,
		57, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122,
		3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 13, 13, 142, 0, 1, 1, 0, 0,
		0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0,
		0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0,
		0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1,
		0, 0, 0, 1, 27, 1, 0, 0, 0, 3, 37, 1, 0, 0, 0, 5, 39, 1, 0, 0, 0, 7, 41,
		1, 0, 0, 0, 9, 43, 1, 0, 0, 0, 11, 52, 1, 0, 0, 0, 13, 61, 1, 0, 0, 0,
		15, 79, 1, 0, 0, 0, 17, 89, 1, 0, 0, 0, 19, 110, 1, 0, 0, 0, 21, 112, 1,
		0, 0, 0, 23, 120, 1, 0, 0, 0, 25, 126, 1, 0, 0, 0, 27, 28, 5, 72, 0, 0,
		28, 29, 5, 67, 0, 0, 29, 30, 5, 76, 0, 0, 30, 31, 5, 67, 0, 0, 31, 32,
		5, 79, 0, 0, 32, 33, 5, 78, 0, 0, 33, 34, 5, 70, 0, 0, 34, 35, 5, 73, 0,
		0, 35, 36, 5, 71, 0, 0, 36, 2, 1, 0, 0, 0, 37, 38, 5, 123, 0, 0, 38, 4,
		1, 0, 0, 0, 39, 40, 5, 125, 0, 0, 40, 6, 1, 0, 0, 0, 41, 42, 5, 61, 0,
		0, 42, 8, 1, 0, 0, 0, 43, 44, 5, 114, 0, 0, 44, 45, 5, 101, 0, 0, 45, 46,
		5, 112, 0, 0, 46, 47, 5, 101, 0, 0, 47, 48, 5, 97, 0, 0, 48, 49, 5, 116,
		0, 0, 49, 50, 5, 101, 0, 0, 50, 51, 5, 100, 0, 0, 51, 10, 1, 0, 0, 0, 52,
		53, 5, 111, 0, 0, 53, 54, 5, 112, 0, 0, 54, 55, 5, 116, 0, 0, 55, 56, 5,
		105, 0, 0, 56, 57, 5, 111, 0, 0, 57, 58, 5, 110, 0, 0, 58, 59, 5, 97, 0,
		0, 59, 60, 5, 108, 0, 0, 60, 12, 1, 0, 0, 0, 61, 62, 5, 111, 0, 0, 62,
		63, 5, 112, 0, 0, 63, 64, 5, 116, 0, 0, 64, 65, 5, 105, 0, 0, 65, 66, 5,
		111, 0, 0, 66, 67, 5, 110, 0, 0, 67, 68, 5, 97, 0, 0, 68, 69, 5, 108, 0,
		0, 69, 70, 5, 32, 0, 0, 70, 71, 5, 114, 0, 0, 71, 72, 5, 101, 0, 0, 72,
		73, 5, 112, 0, 0, 73, 74, 5, 101, 0, 0, 74, 75, 5, 97, 0, 0, 75, 76, 5,
		116, 0, 0, 76, 77, 5, 101, 0, 0, 77, 78, 5, 100, 0, 0, 78, 14, 1, 0, 0,
		0, 79, 83, 5, 34, 0, 0, 80, 82, 8, 0, 0, 0, 81, 80, 1, 0, 0, 0, 82, 85,
		1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 86, 1, 0, 0, 0,
		85, 83, 1, 0, 0, 0, 86, 87, 5, 34, 0, 0, 87, 16, 1, 0, 0, 0, 88, 90, 7,
		1, 0, 0, 89, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 91,
		92, 1, 0, 0, 0, 92, 99, 1, 0, 0, 0, 93, 95, 5, 46, 0, 0, 94, 96, 7, 1,
		0, 0, 95, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 97, 98,
		1, 0, 0, 0, 98, 100, 1, 0, 0, 0, 99, 93, 1, 0, 0, 0, 99, 100, 1, 0, 0,
		0, 100, 18, 1, 0, 0, 0, 101, 102, 5, 116, 0, 0, 102, 103, 5, 114, 0, 0,
		103, 104, 5, 117, 0, 0, 104, 111, 5, 101, 0, 0, 105, 106, 5, 102, 0, 0,
		106, 107, 5, 97, 0, 0, 107, 108, 5, 108, 0, 0, 108, 109, 5, 115, 0, 0,
		109, 111, 5, 101, 0, 0, 110, 101, 1, 0, 0, 0, 110, 105, 1, 0, 0, 0, 111,
		20, 1, 0, 0, 0, 112, 116, 7, 2, 0, 0, 113, 115, 7, 3, 0, 0, 114, 113, 1,
		0, 0, 0, 115, 118, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 116, 117, 1, 0, 0,
		0, 117, 22, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 119, 121, 7, 4, 0, 0, 120,
		119, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 122, 123,
		1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 125, 6, 11, 0, 0, 125, 24, 1, 0,
		0, 0, 126, 130, 5, 35, 0, 0, 127, 129, 8, 5, 0, 0, 128, 127, 1, 0, 0, 0,
		129, 132, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131,
		133, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 133, 134, 6, 12, 0, 0, 134, 26,
		1, 0, 0, 0, 9, 0, 83, 91, 97, 99, 110, 116, 122, 130, 1, 6, 0, 0,
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
	HCLLikeDSLLexerSTRING     = 8
	HCLLikeDSLLexerNUMBER     = 9
	HCLLikeDSLLexerBOOLEAN    = 10
	HCLLikeDSLLexerIDENTIFIER = 11
	HCLLikeDSLLexerWS         = 12
	HCLLikeDSLLexerCOMMENT    = 13
)
