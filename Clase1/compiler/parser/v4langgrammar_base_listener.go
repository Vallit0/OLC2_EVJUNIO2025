// Code generated from parser/V4LangGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // V4LangGrammar
import "github.com/antlr4-go/antlr/v4"

// BaseV4LangGrammarListener is a complete listener for a parse tree produced by V4LangGrammar.
type BaseV4LangGrammarListener struct{}

var _ V4LangGrammarListener = &BaseV4LangGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseV4LangGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseV4LangGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseV4LangGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseV4LangGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseV4LangGrammarListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseV4LangGrammarListener) ExitProg(ctx *ProgContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseV4LangGrammarListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseV4LangGrammarListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseV4LangGrammarListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseV4LangGrammarListener) ExitAddSub(ctx *AddSubContext) {}

// EnterParens is called when production Parens is entered.
func (s *BaseV4LangGrammarListener) EnterParens(ctx *ParensContext) {}

// ExitParens is called when production Parens is exited.
func (s *BaseV4LangGrammarListener) ExitParens(ctx *ParensContext) {}

// EnterInt is called when production Int is entered.
func (s *BaseV4LangGrammarListener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production Int is exited.
func (s *BaseV4LangGrammarListener) ExitInt(ctx *IntContext) {}
