package main

import (
	"strconv"

	"parser" // importa el paquete generado por ANTLR
)

// Interpreter implementa el Visitor para evaluación
type Interpreter struct {
	*parser.BaseV4LangGrammarVisitor
	Env map[string]int
}

func (v *Interpreter) VisitIntLiteral(ctx *parser.IntLiteralContext) interface{} {
	n, _ := strconv.Atoi(ctx.INT().GetText())
	return n
}

func (v *Interpreter) VisitAddExpr(ctx *parser.AddExprContext) interface{} {
	left := ctx.Expr(0).Accept(v).(int)
	right := ctx.Expr(1).Accept(v).(int)
	return left + right
}

func (v *Interpreter) VisitMulExpr(ctx *parser.MulExprContext) interface{} {
	left := ctx.Expr(0).Accept(v).(int)
	right := ctx.Expr(1).Accept(v).(int)
	return left * right
}

func (v *Interpreter) VisitNegExpr(ctx *parser.NegExprContext) interface{} {
	val := ctx.Expr().Accept(v).(int)
	return -val
}

func (v *Interpreter) VisitParenExpr(ctx *parser.ParenExprContext) interface{} {
	return ctx.Expr().Accept(v)
}

func (v *Interpreter) VisitVarExpr(ctx *parser.VarExprContext) interface{} {
	name := ctx.ID().GetText()
	val, ok := v.Env[name]
	if !ok {
		panic("variable no definida: " + name)
	}
	return val
}
