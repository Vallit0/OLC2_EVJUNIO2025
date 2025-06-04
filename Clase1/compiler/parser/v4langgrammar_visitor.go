// Code generated from parser/V4LangGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // V4LangGrammar
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by V4LangGrammar.
type V4LangGrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by V4LangGrammar#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by V4LangGrammar#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by V4LangGrammar#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by V4LangGrammar#Parens.
	VisitParens(ctx *ParensContext) interface{}

	// Visit a parse tree produced by V4LangGrammar#Int.
	VisitInt(ctx *IntContext) interface{}
}
