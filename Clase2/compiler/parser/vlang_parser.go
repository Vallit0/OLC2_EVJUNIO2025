// Code generated from parser/Vlang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // Vlang
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

type VlangParser struct {
	*antlr.BaseParser
}

var VlangParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func vlangParserInit() {
	staticData := &VlangParserStaticData
	staticData.LiteralNames = []string{
		"", "'print('", "'len'", "'cap'", "'append'", "'mut'", "", "", "", "",
		"", "", "", "'+'", "'-'", "'*'", "'/'", "'%'", "'!'", "'||'", "'&&'",
		"'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'='", "'++'", "'--'",
		"'('", "')'", "'['", "']'", "'.'", "','",
	}
	staticData.SymbolicNames = []string{
		"", "", "LEN", "CAP", "APPEND", "MUT", "BOOLEANO", "ENTERO", "FLOAT",
		"DECIMAL", "CADENA", "CARACTER", "ID", "PLUS", "MINUS", "MUL", "DIV",
		"MOD", "NOT", "OR", "AND", "EQ", "NEQ", "LT", "LE", "GT", "GE", "ASSIGN",
		"INC", "DEC", "LPAREN", "RPAREN", "LBRACK", "RBRACK", "DOT", "COMMA",
		"WS", "LINE_COMMENT", "BLOCK_COMMENT",
	}
	staticData.RuleNames = []string{
		"programa", "stmt", "assign_stmt", "decl_stmt", "var_type", "id_pattern",
		"expresion", "parametros", "valores", "valor", "incredecre",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 38, 129, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 4, 0, 24, 8, 0, 11, 0, 12, 0, 25, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 3, 1, 36, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 5, 5, 52, 8, 5, 10, 5,
		12, 5, 55, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 3, 6, 80, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 100,
		8, 6, 10, 6, 12, 6, 103, 9, 6, 1, 7, 1, 7, 1, 7, 5, 7, 108, 8, 7, 10, 7,
		12, 7, 111, 9, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9,
		121, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 127, 8, 10, 1, 10, 0, 1,
		12, 11, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 0, 5, 2, 0, 14, 14, 18,
		18, 1, 0, 15, 17, 1, 0, 13, 14, 1, 0, 23, 26, 1, 0, 21, 22, 142, 0, 23,
		1, 0, 0, 0, 2, 35, 1, 0, 0, 0, 4, 37, 1, 0, 0, 0, 6, 41, 1, 0, 0, 0, 8,
		46, 1, 0, 0, 0, 10, 48, 1, 0, 0, 0, 12, 79, 1, 0, 0, 0, 14, 104, 1, 0,
		0, 0, 16, 112, 1, 0, 0, 0, 18, 120, 1, 0, 0, 0, 20, 126, 1, 0, 0, 0, 22,
		24, 3, 2, 1, 0, 23, 22, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 23, 1, 0, 0,
		0, 25, 26, 1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 28, 5, 0, 0, 1, 28, 1, 1,
		0, 0, 0, 29, 36, 3, 6, 3, 0, 30, 36, 3, 4, 2, 0, 31, 32, 5, 1, 0, 0, 32,
		33, 3, 12, 6, 0, 33, 34, 5, 31, 0, 0, 34, 36, 1, 0, 0, 0, 35, 29, 1, 0,
		0, 0, 35, 30, 1, 0, 0, 0, 35, 31, 1, 0, 0, 0, 36, 3, 1, 0, 0, 0, 37, 38,
		3, 10, 5, 0, 38, 39, 5, 21, 0, 0, 39, 40, 3, 12, 6, 0, 40, 5, 1, 0, 0,
		0, 41, 42, 3, 8, 4, 0, 42, 43, 5, 12, 0, 0, 43, 44, 5, 21, 0, 0, 44, 45,
		3, 12, 6, 0, 45, 7, 1, 0, 0, 0, 46, 47, 5, 5, 0, 0, 47, 9, 1, 0, 0, 0,
		48, 53, 5, 12, 0, 0, 49, 50, 5, 34, 0, 0, 50, 52, 5, 12, 0, 0, 51, 49,
		1, 0, 0, 0, 52, 55, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0,
		54, 11, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 56, 57, 6, 6, -1, 0, 57, 80, 3,
		18, 9, 0, 58, 59, 5, 30, 0, 0, 59, 60, 3, 12, 6, 0, 60, 61, 5, 31, 0, 0,
		61, 80, 1, 0, 0, 0, 62, 63, 5, 32, 0, 0, 63, 64, 3, 12, 6, 0, 64, 65, 5,
		33, 0, 0, 65, 80, 1, 0, 0, 0, 66, 67, 7, 0, 0, 0, 67, 80, 3, 12, 6, 12,
		68, 80, 5, 12, 0, 0, 69, 80, 3, 20, 10, 0, 70, 71, 5, 12, 0, 0, 71, 72,
		5, 34, 0, 0, 72, 80, 5, 12, 0, 0, 73, 74, 5, 12, 0, 0, 74, 75, 5, 34, 0,
		0, 75, 80, 3, 12, 6, 2, 76, 77, 5, 12, 0, 0, 77, 78, 5, 27, 0, 0, 78, 80,
		3, 12, 6, 1, 79, 56, 1, 0, 0, 0, 79, 58, 1, 0, 0, 0, 79, 62, 1, 0, 0, 0,
		79, 66, 1, 0, 0, 0, 79, 68, 1, 0, 0, 0, 79, 69, 1, 0, 0, 0, 79, 70, 1,
		0, 0, 0, 79, 73, 1, 0, 0, 0, 79, 76, 1, 0, 0, 0, 80, 101, 1, 0, 0, 0, 81,
		82, 10, 11, 0, 0, 82, 83, 7, 1, 0, 0, 83, 100, 3, 12, 6, 12, 84, 85, 10,
		10, 0, 0, 85, 86, 7, 2, 0, 0, 86, 100, 3, 12, 6, 11, 87, 88, 10, 9, 0,
		0, 88, 89, 7, 3, 0, 0, 89, 100, 3, 12, 6, 10, 90, 91, 10, 8, 0, 0, 91,
		92, 7, 4, 0, 0, 92, 100, 3, 12, 6, 9, 93, 94, 10, 7, 0, 0, 94, 95, 5, 20,
		0, 0, 95, 100, 3, 12, 6, 8, 96, 97, 10, 6, 0, 0, 97, 98, 5, 19, 0, 0, 98,
		100, 3, 12, 6, 7, 99, 81, 1, 0, 0, 0, 99, 84, 1, 0, 0, 0, 99, 87, 1, 0,
		0, 0, 99, 90, 1, 0, 0, 0, 99, 93, 1, 0, 0, 0, 99, 96, 1, 0, 0, 0, 100,
		103, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 13, 1,
		0, 0, 0, 103, 101, 1, 0, 0, 0, 104, 109, 3, 12, 6, 0, 105, 106, 5, 35,
		0, 0, 106, 108, 3, 12, 6, 0, 107, 105, 1, 0, 0, 0, 108, 111, 1, 0, 0, 0,
		109, 107, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 15, 1, 0, 0, 0, 111, 109,
		1, 0, 0, 0, 112, 113, 3, 18, 9, 0, 113, 17, 1, 0, 0, 0, 114, 121, 5, 7,
		0, 0, 115, 121, 5, 9, 0, 0, 116, 121, 5, 10, 0, 0, 117, 121, 5, 6, 0, 0,
		118, 121, 5, 11, 0, 0, 119, 121, 5, 8, 0, 0, 120, 114, 1, 0, 0, 0, 120,
		115, 1, 0, 0, 0, 120, 116, 1, 0, 0, 0, 120, 117, 1, 0, 0, 0, 120, 118,
		1, 0, 0, 0, 120, 119, 1, 0, 0, 0, 121, 19, 1, 0, 0, 0, 122, 123, 5, 12,
		0, 0, 123, 127, 5, 28, 0, 0, 124, 125, 5, 12, 0, 0, 125, 127, 5, 29, 0,
		0, 126, 122, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 127, 21, 1, 0, 0, 0, 9,
		25, 35, 53, 79, 99, 101, 109, 120, 126,
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

// VlangParserInit initializes any static state used to implement VlangParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewVlangParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func VlangParserInit() {
	staticData := &VlangParserStaticData
	staticData.once.Do(vlangParserInit)
}

// NewVlangParser produces a new parser instance for the optional input antlr.TokenStream.
func NewVlangParser(input antlr.TokenStream) *VlangParser {
	VlangParserInit()
	this := new(VlangParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &VlangParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Vlang.g4"

	return this
}

// VlangParser tokens.
const (
	VlangParserEOF           = antlr.TokenEOF
	VlangParserT__0          = 1
	VlangParserLEN           = 2
	VlangParserCAP           = 3
	VlangParserAPPEND        = 4
	VlangParserMUT           = 5
	VlangParserBOOLEANO      = 6
	VlangParserENTERO        = 7
	VlangParserFLOAT         = 8
	VlangParserDECIMAL       = 9
	VlangParserCADENA        = 10
	VlangParserCARACTER      = 11
	VlangParserID            = 12
	VlangParserPLUS          = 13
	VlangParserMINUS         = 14
	VlangParserMUL           = 15
	VlangParserDIV           = 16
	VlangParserMOD           = 17
	VlangParserNOT           = 18
	VlangParserOR            = 19
	VlangParserAND           = 20
	VlangParserEQ            = 21
	VlangParserNEQ           = 22
	VlangParserLT            = 23
	VlangParserLE            = 24
	VlangParserGT            = 25
	VlangParserGE            = 26
	VlangParserASSIGN        = 27
	VlangParserINC           = 28
	VlangParserDEC           = 29
	VlangParserLPAREN        = 30
	VlangParserRPAREN        = 31
	VlangParserLBRACK        = 32
	VlangParserRBRACK        = 33
	VlangParserDOT           = 34
	VlangParserCOMMA         = 35
	VlangParserWS            = 36
	VlangParserLINE_COMMENT  = 37
	VlangParserBLOCK_COMMENT = 38
)

// VlangParser rules.
const (
	VlangParserRULE_programa    = 0
	VlangParserRULE_stmt        = 1
	VlangParserRULE_assign_stmt = 2
	VlangParserRULE_decl_stmt   = 3
	VlangParserRULE_var_type    = 4
	VlangParserRULE_id_pattern  = 5
	VlangParserRULE_expresion   = 6
	VlangParserRULE_parametros  = 7
	VlangParserRULE_valores     = 8
	VlangParserRULE_valor       = 9
	VlangParserRULE_incredecre  = 10
)

// IProgramaContext is an interface to support dynamic dispatch.
type IProgramaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsProgramaContext differentiates from other interfaces.
	IsProgramaContext()
}

type ProgramaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramaContext() *ProgramaContext {
	var p = new(ProgramaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_programa
	return p
}

func InitEmptyProgramaContext(p *ProgramaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_programa
}

func (*ProgramaContext) IsProgramaContext() {}

func NewProgramaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramaContext {
	var p = new(ProgramaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_programa

	return p
}

func (s *ProgramaContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramaContext) EOF() antlr.TerminalNode {
	return s.GetToken(VlangParserEOF, 0)
}

func (s *ProgramaContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ProgramaContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
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

	return t.(IStmtContext)
}

func (s *ProgramaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterPrograma(s)
	}
}

func (s *ProgramaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitPrograma(s)
	}
}

func (s *ProgramaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitPrograma(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Programa() (localctx IProgramaContext) {
	localctx = NewProgramaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VlangParserRULE_programa)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4130) != 0) {
		{
			p.SetState(22)
			p.Stmt()
		}

		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(27)
		p.Match(VlangParserEOF)
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

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Decl_stmt() IDecl_stmtContext
	Assign_stmt() IAssign_stmtContext
	Expresion() IExpresionContext
	RPAREN() antlr.TerminalNode

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Decl_stmt() IDecl_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecl_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecl_stmtContext)
}

func (s *StmtContext) Assign_stmt() IAssign_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_stmtContext)
}

func (s *StmtContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *StmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, VlangParserRULE_stmt)
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VlangParserMUT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(29)
			p.Decl_stmt()
		}

	case VlangParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(30)
			p.Assign_stmt()
		}

	case VlangParserT__0:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(31)
			p.Match(VlangParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(32)
			p.expresion(0)
		}
		{
			p.SetState(33)
			p.Match(VlangParserRPAREN)
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

// IAssign_stmtContext is an interface to support dynamic dispatch.
type IAssign_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssign_stmtContext differentiates from other interfaces.
	IsAssign_stmtContext()
}

type Assign_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssign_stmtContext() *Assign_stmtContext {
	var p = new(Assign_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_assign_stmt
	return p
}

func InitEmptyAssign_stmtContext(p *Assign_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_assign_stmt
}

func (*Assign_stmtContext) IsAssign_stmtContext() {}

func NewAssign_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_stmtContext {
	var p = new(Assign_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_assign_stmt

	return p
}

func (s *Assign_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_stmtContext) CopyAll(ctx *Assign_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Assign_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DirectAssignContext struct {
	Assign_stmtContext
}

func NewDirectAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DirectAssignContext {
	var p = new(DirectAssignContext)

	InitEmptyAssign_stmtContext(&p.Assign_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_stmtContext))

	return p
}

func (s *DirectAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectAssignContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *DirectAssignContext) EQ() antlr.TerminalNode {
	return s.GetToken(VlangParserEQ, 0)
}

func (s *DirectAssignContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *DirectAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterDirectAssign(s)
	}
}

func (s *DirectAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitDirectAssign(s)
	}
}

func (s *DirectAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitDirectAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Assign_stmt() (localctx IAssign_stmtContext) {
	localctx = NewAssign_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, VlangParserRULE_assign_stmt)
	localctx = NewDirectAssignContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(37)
		p.Id_pattern()
	}
	{
		p.SetState(38)
		p.Match(VlangParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)
		p.expresion(0)
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

// IDecl_stmtContext is an interface to support dynamic dispatch.
type IDecl_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_type() IVar_typeContext
	ID() antlr.TerminalNode
	EQ() antlr.TerminalNode
	Expresion() IExpresionContext

	// IsDecl_stmtContext differentiates from other interfaces.
	IsDecl_stmtContext()
}

type Decl_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecl_stmtContext() *Decl_stmtContext {
	var p = new(Decl_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_decl_stmt
	return p
}

func InitEmptyDecl_stmtContext(p *Decl_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_decl_stmt
}

func (*Decl_stmtContext) IsDecl_stmtContext() {}

func NewDecl_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Decl_stmtContext {
	var p = new(Decl_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_decl_stmt

	return p
}

func (s *Decl_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Decl_stmtContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *Decl_stmtContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *Decl_stmtContext) EQ() antlr.TerminalNode {
	return s.GetToken(VlangParserEQ, 0)
}

func (s *Decl_stmtContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Decl_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Decl_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Decl_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterDecl_stmt(s)
	}
}

func (s *Decl_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitDecl_stmt(s)
	}
}

func (s *Decl_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitDecl_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Decl_stmt() (localctx IDecl_stmtContext) {
	localctx = NewDecl_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, VlangParserRULE_decl_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(41)
		p.Var_type()
	}
	{
		p.SetState(42)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(43)
		p.Match(VlangParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(44)
		p.expresion(0)
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

// IVar_typeContext is an interface to support dynamic dispatch.
type IVar_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVar_typeContext differentiates from other interfaces.
	IsVar_typeContext()
}

type Var_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_typeContext() *Var_typeContext {
	var p = new(Var_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_var_type
	return p
}

func InitEmptyVar_typeContext(p *Var_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_var_type
}

func (*Var_typeContext) IsVar_typeContext() {}

func NewVar_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_typeContext {
	var p = new(Var_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_var_type

	return p
}

func (s *Var_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_typeContext) CopyAll(ctx *Var_typeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Var_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MutTypeContext struct {
	Var_typeContext
}

func NewMutTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MutTypeContext {
	var p = new(MutTypeContext)

	InitEmptyVar_typeContext(&p.Var_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Var_typeContext))

	return p
}

func (s *MutTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MutTypeContext) MUT() antlr.TerminalNode {
	return s.GetToken(VlangParserMUT, 0)
}

func (s *MutTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterMutType(s)
	}
}

func (s *MutTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitMutType(s)
	}
}

func (s *MutTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitMutType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Var_type() (localctx IVar_typeContext) {
	localctx = NewVar_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, VlangParserRULE_var_type)
	localctx = NewMutTypeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Match(VlangParserMUT)
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

// IId_patternContext is an interface to support dynamic dispatch.
type IId_patternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsId_patternContext differentiates from other interfaces.
	IsId_patternContext()
}

type Id_patternContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyId_patternContext() *Id_patternContext {
	var p = new(Id_patternContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_id_pattern
	return p
}

func InitEmptyId_patternContext(p *Id_patternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_id_pattern
}

func (*Id_patternContext) IsId_patternContext() {}

func NewId_patternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Id_patternContext {
	var p = new(Id_patternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_id_pattern

	return p
}

func (s *Id_patternContext) GetParser() antlr.Parser { return s.parser }

func (s *Id_patternContext) CopyAll(ctx *Id_patternContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Id_patternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Id_patternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IdPatternContext struct {
	Id_patternContext
}

func NewIdPatternContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdPatternContext {
	var p = new(IdPatternContext)

	InitEmptyId_patternContext(&p.Id_patternContext)
	p.parser = parser
	p.CopyAll(ctx.(*Id_patternContext))

	return p
}

func (s *IdPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdPatternContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(VlangParserID)
}

func (s *IdPatternContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserID, i)
}

func (s *IdPatternContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(VlangParserDOT)
}

func (s *IdPatternContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserDOT, i)
}

func (s *IdPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterIdPattern(s)
	}
}

func (s *IdPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitIdPattern(s)
	}
}

func (s *IdPatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitIdPattern(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Id_pattern() (localctx IId_patternContext) {
	localctx = NewId_patternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, VlangParserRULE_id_pattern)
	var _la int

	localctx = NewIdPatternContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(VlangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VlangParserDOT {
		{
			p.SetState(49)
			p.Match(VlangParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(55)
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

// IExpresionContext is an interface to support dynamic dispatch.
type IExpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpresionContext() *ExpresionContext {
	var p = new(ExpresionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_expresion
	return p
}

func InitEmptyExpresionContext(p *ExpresionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_expresion
}

func (*ExpresionContext) IsExpresionContext() {}

func NewExpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionContext {
	var p = new(ExpresionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_expresion

	return p
}

func (s *ExpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionContext) CopyAll(ctx *ExpresionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CorchetesexpreContext struct {
	ExpresionContext
}

func NewCorchetesexpreContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CorchetesexpreContext {
	var p = new(CorchetesexpreContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *CorchetesexpreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CorchetesexpreContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(VlangParserLBRACK, 0)
}

func (s *CorchetesexpreContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *CorchetesexpreContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(VlangParserRBRACK, 0)
}

func (s *CorchetesexpreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterCorchetesexpre(s)
	}
}

func (s *CorchetesexpreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitCorchetesexpre(s)
	}
}

func (s *CorchetesexpreContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitCorchetesexpre(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnarioContext struct {
	ExpresionContext
	op antlr.Token
}

func NewUnarioContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnarioContext {
	var p = new(UnarioContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *UnarioContext) GetOp() antlr.Token { return s.op }

func (s *UnarioContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnarioContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *UnarioContext) NOT() antlr.TerminalNode {
	return s.GetToken(VlangParserNOT, 0)
}

func (s *UnarioContext) MINUS() antlr.TerminalNode {
	return s.GetToken(VlangParserMINUS, 0)
}

func (s *UnarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterUnario(s)
	}
}

func (s *UnarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitUnario(s)
	}
}

func (s *UnarioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitUnario(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParentesisexpreContext struct {
	ExpresionContext
}

func NewParentesisexpreContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentesisexpreContext {
	var p = new(ParentesisexpreContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ParentesisexpreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentesisexpreContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserLPAREN, 0)
}

func (s *ParentesisexpreContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ParentesisexpreContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VlangParserRPAREN, 0)
}

func (s *ParentesisexpreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterParentesisexpre(s)
	}
}

func (s *ParentesisexpreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitParentesisexpre(s)
	}
}

func (s *ParentesisexpreContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitParentesisexpre(s)

	default:
		return t.VisitChildren(s)
	}
}

type IncredecrContext struct {
	ExpresionContext
}

func NewIncredecrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IncredecrContext {
	var p = new(IncredecrContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *IncredecrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncredecrContext) Incredecre() IIncredecreContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncredecreContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIncredecreContext)
}

func (s *IncredecrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterIncredecr(s)
	}
}

func (s *IncredecrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitIncredecr(s)
	}
}

func (s *IncredecrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitIncredecr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValorexpresionContext struct {
	ExpresionContext
}

func NewValorexpresionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorexpresionContext {
	var p = new(ValorexpresionContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ValorexpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorexpresionContext) Valor() IValorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValorContext)
}

func (s *ValorexpresionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorexpresion(s)
	}
}

func (s *ValorexpresionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorexpresion(s)
	}
}

func (s *ValorexpresionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorexpresion(s)

	default:
		return t.VisitChildren(s)
	}
}

type AsignacionforContext struct {
	ExpresionContext
}

func NewAsignacionforContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionforContext {
	var p = new(AsignacionforContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *AsignacionforContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionforContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *AsignacionforContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(VlangParserASSIGN, 0)
}

func (s *AsignacionforContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AsignacionforContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterAsignacionfor(s)
	}
}

func (s *AsignacionforContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitAsignacionfor(s)
	}
}

func (s *AsignacionforContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitAsignacionfor(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdContext struct {
	ExpresionContext
}

func NewIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdContext {
	var p = new(IdContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitId(s)
	}
}

func (s *IdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitId(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expdotexp1Context struct {
	ExpresionContext
}

func NewExpdotexp1Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expdotexp1Context {
	var p = new(Expdotexp1Context)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Expdotexp1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expdotexp1Context) AllID() []antlr.TerminalNode {
	return s.GetTokens(VlangParserID)
}

func (s *Expdotexp1Context) ID(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserID, i)
}

func (s *Expdotexp1Context) DOT() antlr.TerminalNode {
	return s.GetToken(VlangParserDOT, 0)
}

func (s *Expdotexp1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterExpdotexp1(s)
	}
}

func (s *Expdotexp1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitExpdotexp1(s)
	}
}

func (s *Expdotexp1Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitExpdotexp1(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpdotexpContext struct {
	ExpresionContext
}

func NewExpdotexpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpdotexpContext {
	var p = new(ExpdotexpContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ExpdotexpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpdotexpContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *ExpdotexpContext) DOT() antlr.TerminalNode {
	return s.GetToken(VlangParserDOT, 0)
}

func (s *ExpdotexpContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExpdotexpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterExpdotexp(s)
	}
}

func (s *ExpdotexpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitExpdotexp(s)
	}
}

func (s *ExpdotexpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitExpdotexp(s)

	default:
		return t.VisitChildren(s)
	}
}

type BinaryExpContext struct {
	ExpresionContext
	left  IExpresionContext
	op    antlr.Token
	right IExpresionContext
}

func NewBinaryExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryExpContext {
	var p = new(BinaryExpContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *BinaryExpContext) GetOp() antlr.Token { return s.op }

func (s *BinaryExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *BinaryExpContext) GetLeft() IExpresionContext { return s.left }

func (s *BinaryExpContext) GetRight() IExpresionContext { return s.right }

func (s *BinaryExpContext) SetLeft(v IExpresionContext) { s.left = v }

func (s *BinaryExpContext) SetRight(v IExpresionContext) { s.right = v }

func (s *BinaryExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryExpContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *BinaryExpContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
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

	return t.(IExpresionContext)
}

func (s *BinaryExpContext) MUL() antlr.TerminalNode {
	return s.GetToken(VlangParserMUL, 0)
}

func (s *BinaryExpContext) DIV() antlr.TerminalNode {
	return s.GetToken(VlangParserDIV, 0)
}

func (s *BinaryExpContext) MOD() antlr.TerminalNode {
	return s.GetToken(VlangParserMOD, 0)
}

func (s *BinaryExpContext) PLUS() antlr.TerminalNode {
	return s.GetToken(VlangParserPLUS, 0)
}

func (s *BinaryExpContext) MINUS() antlr.TerminalNode {
	return s.GetToken(VlangParserMINUS, 0)
}

func (s *BinaryExpContext) LT() antlr.TerminalNode {
	return s.GetToken(VlangParserLT, 0)
}

func (s *BinaryExpContext) LE() antlr.TerminalNode {
	return s.GetToken(VlangParserLE, 0)
}

func (s *BinaryExpContext) GT() antlr.TerminalNode {
	return s.GetToken(VlangParserGT, 0)
}

func (s *BinaryExpContext) GE() antlr.TerminalNode {
	return s.GetToken(VlangParserGE, 0)
}

func (s *BinaryExpContext) EQ() antlr.TerminalNode {
	return s.GetToken(VlangParserEQ, 0)
}

func (s *BinaryExpContext) NEQ() antlr.TerminalNode {
	return s.GetToken(VlangParserNEQ, 0)
}

func (s *BinaryExpContext) AND() antlr.TerminalNode {
	return s.GetToken(VlangParserAND, 0)
}

func (s *BinaryExpContext) OR() antlr.TerminalNode {
	return s.GetToken(VlangParserOR, 0)
}

func (s *BinaryExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterBinaryExp(s)
	}
}

func (s *BinaryExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitBinaryExp(s)
	}
}

func (s *BinaryExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitBinaryExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Expresion() (localctx IExpresionContext) {
	return p.expresion(0)
}

func (p *VlangParser) expresion(_p int) (localctx IExpresionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpresionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpresionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, VlangParserRULE_expresion, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewValorexpresionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(57)
			p.Valor()
		}

	case 2:
		localctx = NewParentesisexpreContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(58)
			p.Match(VlangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(59)
			p.expresion(0)
		}
		{
			p.SetState(60)
			p.Match(VlangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewCorchetesexpreContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(62)
			p.Match(VlangParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)
			p.expresion(0)
		}
		{
			p.SetState(64)
			p.Match(VlangParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewUnarioContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(66)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnarioContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == VlangParserMINUS || _la == VlangParserNOT) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnarioContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(67)
			p.expresion(12)
		}

	case 5:
		localctx = NewIdContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(68)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewIncredecrContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(69)
			p.Incredecre()
		}

	case 7:
		localctx = NewExpdotexp1Context(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(70)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.Match(VlangParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(72)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewExpdotexpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(73)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
			p.Match(VlangParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.expresion(2)
		}

	case 9:
		localctx = NewAsignacionforContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(76)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.Match(VlangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)
			p.expresion(1)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(99)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryExpContext(p, NewExpresionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(81)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(82)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&229376) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(83)

					var _x = p.expresion(12)

					localctx.(*BinaryExpContext).right = _x
				}

			case 2:
				localctx = NewBinaryExpContext(p, NewExpresionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(84)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(85)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == VlangParserPLUS || _la == VlangParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(86)

					var _x = p.expresion(11)

					localctx.(*BinaryExpContext).right = _x
				}

			case 3:
				localctx = NewBinaryExpContext(p, NewExpresionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(87)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(88)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&125829120) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(89)

					var _x = p.expresion(10)

					localctx.(*BinaryExpContext).right = _x
				}

			case 4:
				localctx = NewBinaryExpContext(p, NewExpresionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(90)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(91)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == VlangParserEQ || _la == VlangParserNEQ) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(92)

					var _x = p.expresion(9)

					localctx.(*BinaryExpContext).right = _x
				}

			case 5:
				localctx = NewBinaryExpContext(p, NewExpresionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(93)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(94)

					var _m = p.Match(VlangParserAND)

					localctx.(*BinaryExpContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(95)

					var _x = p.expresion(8)

					localctx.(*BinaryExpContext).right = _x
				}

			case 6:
				localctx = NewBinaryExpContext(p, NewExpresionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VlangParserRULE_expresion)
				p.SetState(96)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(97)

					var _m = p.Match(VlangParserOR)

					localctx.(*BinaryExpContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(98)

					var _x = p.expresion(7)

					localctx.(*BinaryExpContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParametrosContext is an interface to support dynamic dispatch.
type IParametrosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpresion() []IExpresionContext
	Expresion(i int) IExpresionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParametrosContext differentiates from other interfaces.
	IsParametrosContext()
}

type ParametrosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametrosContext() *ParametrosContext {
	var p = new(ParametrosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_parametros
	return p
}

func InitEmptyParametrosContext(p *ParametrosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_parametros
}

func (*ParametrosContext) IsParametrosContext() {}

func NewParametrosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametrosContext {
	var p = new(ParametrosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_parametros

	return p
}

func (s *ParametrosContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametrosContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *ParametrosContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
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

	return t.(IExpresionContext)
}

func (s *ParametrosContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VlangParserCOMMA)
}

func (s *ParametrosContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VlangParserCOMMA, i)
}

func (s *ParametrosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametrosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterParametros(s)
	}
}

func (s *ParametrosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitParametros(s)
	}
}

func (s *ParametrosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitParametros(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Parametros() (localctx IParametrosContext) {
	localctx = NewParametrosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, VlangParserRULE_parametros)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.expresion(0)
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VlangParserCOMMA {
		{
			p.SetState(105)
			p.Match(VlangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.expresion(0)
		}

		p.SetState(111)
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

// IValoresContext is an interface to support dynamic dispatch.
type IValoresContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Valor() IValorContext

	// IsValoresContext differentiates from other interfaces.
	IsValoresContext()
}

type ValoresContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValoresContext() *ValoresContext {
	var p = new(ValoresContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_valores
	return p
}

func InitEmptyValoresContext(p *ValoresContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_valores
}

func (*ValoresContext) IsValoresContext() {}

func NewValoresContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValoresContext {
	var p = new(ValoresContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_valores

	return p
}

func (s *ValoresContext) GetParser() antlr.Parser { return s.parser }

func (s *ValoresContext) Valor() IValorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValorContext)
}

func (s *ValoresContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValoresContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValoresContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValores(s)
	}
}

func (s *ValoresContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValores(s)
	}
}

func (s *ValoresContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValores(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Valores() (localctx IValoresContext) {
	localctx = NewValoresContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, VlangParserRULE_valores)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.Valor()
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

// IValorContext is an interface to support dynamic dispatch.
type IValorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsValorContext differentiates from other interfaces.
	IsValorContext()
}

type ValorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValorContext() *ValorContext {
	var p = new(ValorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_valor
	return p
}

func InitEmptyValorContext(p *ValorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_valor
}

func (*ValorContext) IsValorContext() {}

func NewValorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValorContext {
	var p = new(ValorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_valor

	return p
}

func (s *ValorContext) GetParser() antlr.Parser { return s.parser }

func (s *ValorContext) CopyAll(ctx *ValorContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ValorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ValorDecimalContext struct {
	ValorContext
}

func NewValorDecimalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorDecimalContext {
	var p = new(ValorDecimalContext)

	InitEmptyValorContext(&p.ValorContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValorContext))

	return p
}

func (s *ValorDecimalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorDecimalContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(VlangParserDECIMAL, 0)
}

func (s *ValorDecimalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorDecimal(s)
	}
}

func (s *ValorDecimalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorDecimal(s)
	}
}

func (s *ValorDecimalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorDecimal(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValorEnteroContext struct {
	ValorContext
}

func NewValorEnteroContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorEnteroContext {
	var p = new(ValorEnteroContext)

	InitEmptyValorContext(&p.ValorContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValorContext))

	return p
}

func (s *ValorEnteroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorEnteroContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(VlangParserENTERO, 0)
}

func (s *ValorEnteroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorEntero(s)
	}
}

func (s *ValorEnteroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorEntero(s)
	}
}

func (s *ValorEnteroContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorEntero(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValorBooleanoContext struct {
	ValorContext
}

func NewValorBooleanoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorBooleanoContext {
	var p = new(ValorBooleanoContext)

	InitEmptyValorContext(&p.ValorContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValorContext))

	return p
}

func (s *ValorBooleanoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorBooleanoContext) BOOLEANO() antlr.TerminalNode {
	return s.GetToken(VlangParserBOOLEANO, 0)
}

func (s *ValorBooleanoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorBooleano(s)
	}
}

func (s *ValorBooleanoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorBooleano(s)
	}
}

func (s *ValorBooleanoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorBooleano(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValorCaracterContext struct {
	ValorContext
}

func NewValorCaracterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorCaracterContext {
	var p = new(ValorCaracterContext)

	InitEmptyValorContext(&p.ValorContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValorContext))

	return p
}

func (s *ValorCaracterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorCaracterContext) CARACTER() antlr.TerminalNode {
	return s.GetToken(VlangParserCARACTER, 0)
}

func (s *ValorCaracterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorCaracter(s)
	}
}

func (s *ValorCaracterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorCaracter(s)
	}
}

func (s *ValorCaracterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorCaracter(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValorCadenaContext struct {
	ValorContext
}

func NewValorCadenaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorCadenaContext {
	var p = new(ValorCadenaContext)

	InitEmptyValorContext(&p.ValorContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValorContext))

	return p
}

func (s *ValorCadenaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorCadenaContext) CADENA() antlr.TerminalNode {
	return s.GetToken(VlangParserCADENA, 0)
}

func (s *ValorCadenaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorCadena(s)
	}
}

func (s *ValorCadenaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorCadena(s)
	}
}

func (s *ValorCadenaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorCadena(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValorFloatContext struct {
	ValorContext
}

func NewValorFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorFloatContext {
	var p = new(ValorFloatContext)

	InitEmptyValorContext(&p.ValorContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValorContext))

	return p
}

func (s *ValorFloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorFloatContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(VlangParserFLOAT, 0)
}

func (s *ValorFloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterValorFloat(s)
	}
}

func (s *ValorFloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitValorFloat(s)
	}
}

func (s *ValorFloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitValorFloat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Valor() (localctx IValorContext) {
	localctx = NewValorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, VlangParserRULE_valor)
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VlangParserENTERO:
		localctx = NewValorEnteroContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(114)
			p.Match(VlangParserENTERO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserDECIMAL:
		localctx = NewValorDecimalContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(115)
			p.Match(VlangParserDECIMAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserCADENA:
		localctx = NewValorCadenaContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(116)
			p.Match(VlangParserCADENA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserBOOLEANO:
		localctx = NewValorBooleanoContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(117)
			p.Match(VlangParserBOOLEANO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserCARACTER:
		localctx = NewValorCaracterContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(118)
			p.Match(VlangParserCARACTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VlangParserFLOAT:
		localctx = NewValorFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(119)
			p.Match(VlangParserFLOAT)
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

// IIncredecreContext is an interface to support dynamic dispatch.
type IIncredecreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIncredecreContext differentiates from other interfaces.
	IsIncredecreContext()
}

type IncredecreContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncredecreContext() *IncredecreContext {
	var p = new(IncredecreContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_incredecre
	return p
}

func InitEmptyIncredecreContext(p *IncredecreContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VlangParserRULE_incredecre
}

func (*IncredecreContext) IsIncredecreContext() {}

func NewIncredecreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncredecreContext {
	var p = new(IncredecreContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VlangParserRULE_incredecre

	return p
}

func (s *IncredecreContext) GetParser() antlr.Parser { return s.parser }

func (s *IncredecreContext) CopyAll(ctx *IncredecreContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IncredecreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncredecreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IncrementoContext struct {
	IncredecreContext
}

func NewIncrementoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IncrementoContext {
	var p = new(IncrementoContext)

	InitEmptyIncredecreContext(&p.IncredecreContext)
	p.parser = parser
	p.CopyAll(ctx.(*IncredecreContext))

	return p
}

func (s *IncrementoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncrementoContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *IncrementoContext) INC() antlr.TerminalNode {
	return s.GetToken(VlangParserINC, 0)
}

func (s *IncrementoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterIncremento(s)
	}
}

func (s *IncrementoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitIncremento(s)
	}
}

func (s *IncrementoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitIncremento(s)

	default:
		return t.VisitChildren(s)
	}
}

type DecrementoContext struct {
	IncredecreContext
}

func NewDecrementoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecrementoContext {
	var p = new(DecrementoContext)

	InitEmptyIncredecreContext(&p.IncredecreContext)
	p.parser = parser
	p.CopyAll(ctx.(*IncredecreContext))

	return p
}

func (s *DecrementoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecrementoContext) ID() antlr.TerminalNode {
	return s.GetToken(VlangParserID, 0)
}

func (s *DecrementoContext) DEC() antlr.TerminalNode {
	return s.GetToken(VlangParserDEC, 0)
}

func (s *DecrementoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.EnterDecremento(s)
	}
}

func (s *DecrementoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VlangListener); ok {
		listenerT.ExitDecremento(s)
	}
}

func (s *DecrementoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VlangVisitor:
		return t.VisitDecremento(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VlangParser) Incredecre() (localctx IIncredecreContext) {
	localctx = NewIncredecreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, VlangParserRULE_incredecre)
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIncrementoContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)
			p.Match(VlangParserINC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewDecrementoContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(124)
			p.Match(VlangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.Match(VlangParserDEC)
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

func (p *VlangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 6:
		var t *ExpresionContext = nil
		if localctx != nil {
			t = localctx.(*ExpresionContext)
		}
		return p.Expresion_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *VlangParser) Expresion_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
