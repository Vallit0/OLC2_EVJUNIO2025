// Code generated from parser/V4LangGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // V4LangGrammar
import "github.com/antlr4-go/antlr/v4"

type BaseV4LangGrammarVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseV4LangGrammarVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitParens(ctx *ParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV4LangGrammarVisitor) VisitInt(ctx *IntContext) interface{} {
	return v.VisitChildren(ctx)
}
