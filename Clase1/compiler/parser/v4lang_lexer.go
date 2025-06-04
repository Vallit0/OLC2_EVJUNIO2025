// Code generated from parser/V4LangLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler

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

type V4LangLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var V4LangLexerLexerStaticData struct {
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

func v4langlexerLexerInit() {
	staticData := &V4LangLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "'+'", "'-'", "'*'", "'/'", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "PLUS", "MINUS", "MUL", "DIV", "LPAREN", "RPAREN", "WS",
	}
	staticData.RuleNames = []string{
		"INT", "PLUS", "MINUS", "MUL", "DIV", "LPAREN", "RPAREN", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 8, 41, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 4, 0, 19, 8, 0, 11,
		0, 12, 0, 20, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 7, 4, 7, 36, 8, 7, 11, 7, 12, 7, 37, 1, 7, 1, 7, 0, 0,
		8, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 1, 0, 2, 1, 0, 48,
		57, 3, 0, 9, 10, 13, 13, 32, 32, 42, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0,
		0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0,
		0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 1, 18, 1, 0, 0, 0, 3, 22, 1, 0, 0,
		0, 5, 24, 1, 0, 0, 0, 7, 26, 1, 0, 0, 0, 9, 28, 1, 0, 0, 0, 11, 30, 1,
		0, 0, 0, 13, 32, 1, 0, 0, 0, 15, 35, 1, 0, 0, 0, 17, 19, 7, 0, 0, 0, 18,
		17, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 18, 1, 0, 0, 0, 20, 21, 1, 0, 0,
		0, 21, 2, 1, 0, 0, 0, 22, 23, 5, 43, 0, 0, 23, 4, 1, 0, 0, 0, 24, 25, 5,
		45, 0, 0, 25, 6, 1, 0, 0, 0, 26, 27, 5, 42, 0, 0, 27, 8, 1, 0, 0, 0, 28,
		29, 5, 47, 0, 0, 29, 10, 1, 0, 0, 0, 30, 31, 5, 40, 0, 0, 31, 12, 1, 0,
		0, 0, 32, 33, 5, 41, 0, 0, 33, 14, 1, 0, 0, 0, 34, 36, 7, 1, 0, 0, 35,
		34, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 37, 38, 1, 0, 0,
		0, 38, 39, 1, 0, 0, 0, 39, 40, 6, 7, 0, 0, 40, 16, 1, 0, 0, 0, 3, 0, 20,
		37, 1, 6, 0, 0,
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

// V4LangLexerInit initializes any static state used to implement V4LangLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewV4LangLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func V4LangLexerInit() {
	staticData := &V4LangLexerLexerStaticData
	staticData.once.Do(v4langlexerLexerInit)
}

// NewV4LangLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewV4LangLexer(input antlr.CharStream) *V4LangLexer {
	V4LangLexerInit()
	l := new(V4LangLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &V4LangLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "V4LangLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// V4LangLexer tokens.
const (
	V4LangLexerINT    = 1
	V4LangLexerPLUS   = 2
	V4LangLexerMINUS  = 3
	V4LangLexerMUL    = 4
	V4LangLexerDIV    = 5
	V4LangLexerLPAREN = 6
	V4LangLexerRPAREN = 7
	V4LangLexerWS     = 8
)
