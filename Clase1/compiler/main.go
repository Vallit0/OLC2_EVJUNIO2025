package main

import (
	// Ajusta si tu módulo se llama diferente
	"fmt"

	compiler "compiler/parser"

	"github.com/antlr4-go/antlr/v4"
)

/*
Go no provee una estructura directa para crear clases
asi que cuando lo manejamos de manera interna,
solamente podemos utilizar unas funciones y hacerles un estilo
de attachment por medio de las funciones func (v)
*/
type EvalVisitor struct {
	*compiler.BaseV4LangGrammarVisitor
}

func (v *EvalVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *EvalVisitor) VisitProg(ctx *compiler.ProgContext) interface{} {
	return v.Visit(ctx.Expr())
}

func (v *EvalVisitor) VisitAddSub(ctx *compiler.AddSubContext) interface{} {
	left := v.Visit(ctx.Expr(0)).(int)
	right := v.Visit(ctx.Expr(1)).(int)

	switch ctx.GetOp().GetTokenType() {
	case compiler.V4LangLexerPLUS:
		return left + right
	case compiler.V4LangLexerMINUS:
		return left - right
	}
	return 0
}

func (v *EvalVisitor) VisitMulDiv(ctx *compiler.MulDivContext) interface{} {
	left := v.Visit(ctx.Expr(0)).(int)
	right := v.Visit(ctx.Expr(1)).(int)

	switch ctx.GetOp().GetTokenType() {
	case compiler.V4LangLexerMUL:
		return left * right
	case compiler.V4LangLexerDIV:
		return left / right
	}
	return 0
}

func (v *EvalVisitor) VisitInt(ctx *compiler.IntContext) interface{} {
	var value int
	fmt.Sscanf(ctx.GetText(), "%d", &value)
	return value
}

func (v *EvalVisitor) VisitParens(ctx *compiler.ParensContext) interface{} {
	return v.Visit(ctx.Expr())
}

func main() {
	input := antlr.NewInputStream("3 + 4 * (2 - 1)")

	lexer := compiler.NewV4LangLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := compiler.NewV4LangGrammar(tokens)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	tree := parser.Prog()

	visitor := &EvalVisitor{}
	result := visitor.Visit(tree)

	fmt.Println("Resultado:", result)
}
