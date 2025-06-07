package repl

import (
	parser "compiler/parser"
	"fmt"
	"log"

	"github.com/antlr4-go/antlr/v4"
)

// Visitor personalizado para recorrer el árbol de sintaxis
type ReplVisitor struct {
	parser.BaseVlangVisitor // Embebemos el visitor generado por ANTLR
}

// Constructor del visitor
func NewReplVisitor() *ReplVisitor {
	return &ReplVisitor{}
}
func (v *ReplVisitor) Visit(tree antlr.ParseTree) interface{} {

	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(val.GetText())
		return nil
	default:
		return tree.Accept(v)
	}

}

func (v *ReplVisitor) VisitPrograma(ctx *parser.ProgramaContext) interface{} {

	for _, stmt := range ctx.AllDeclaraciones() {
		fmt.Println("🔍 Visitando declaración:", stmt.GetText())
		v.Visit(stmt)
	}

	return nil
}

// Visitamos declaraciones de variables
func (v *ReplVisitor) VisitVariableDeclaration(ctx *parser.VariableDeclarationContext) interface{} {
	fmt.Println("📝 Declaración de variable detectada")
	return v.VisitChildren(ctx)
}

// Visitamos sentencias de impresión
func (v *ReplVisitor) VisitPrintStatement(ctx *parser.PrintStatementContext) interface{} {
	fmt.Println("📤 Print statement ejecutado (simulado)")
	return v.VisitChildren(ctx)
}

func (v *ReplVisitor) VisitParentesisexpre(ctx *parser.ParentesisexpreContext) interface{} {
	fmt.Println("🌀 Paréntesis:", ctx.GetText())
	return v.Visit(ctx.Expresion())
}

func (v *ReplVisitor) VisitCorchetesexpre(ctx *parser.CorchetesexpreContext) interface{} {
	fmt.Println("🧱 Corchetes:", ctx.GetText())
	return v.Visit(ctx.Expresion())
}

func (v *ReplVisitor) VisitUnario(ctx *parser.UnarioContext) interface{} {
	fmt.Println("🔁 Operador unario:", ctx.GetOp().GetText())
	return v.Visit(ctx.Expresion())
}

func (v *ReplVisitor) VisitMultdivmod(ctx *parser.MultdivmodContext) interface{} {
	fmt.Println("✖️ Multiplicación/División/Módulo:", ctx.GetOp().GetText())
	return v.VisitChildren(ctx)
}

func (v *ReplVisitor) VisitSumres(ctx *parser.SumresContext) interface{} {
	fmt.Println("➕ Suma/Resta:", ctx.GetOp().GetText())
	return v.VisitChildren(ctx)
}

func (v *ReplVisitor) VisitRelacionales(ctx *parser.RelacionalesContext) interface{} {
	fmt.Println("🔍 Comparación relacional:", ctx.GetOp().GetText())
	return v.VisitChildren(ctx)
}

func (v *ReplVisitor) VisitIgualdad(ctx *parser.IgualdadContext) interface{} {
	fmt.Println("🟰 Comparación de igualdad:", ctx.GetOp().GetText())
	return v.VisitChildren(ctx)
}

func (v *ReplVisitor) VisitOr(ctx *parser.OrContext) interface{} {
	fmt.Println("🚪 Operador OR lógico")
	return v.VisitChildren(ctx)
}

func (v *ReplVisitor) VisitId(ctx *parser.IdContext) interface{} {
	fmt.Println("🆔 Identificador:", ctx.GetText())
	return nil
}

func (v *ReplVisitor) VisitIncredecr(ctx *parser.IncredecrContext) interface{} {
	fmt.Println("🔼 Incremento/Decremento:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *ReplVisitor) VisitExpdotexp1(ctx *parser.Expdotexp1Context) interface{} {
	fmt.Println("📌 Acceso punto ID.ID:", ctx.GetText())
	return nil
}

func (v *ReplVisitor) VisitExpdotexp(ctx *parser.ExpdotexpContext) interface{} {
	fmt.Println("📌 Acceso punto ID.exp:", ctx.GetText())
	return v.Visit(ctx.Expresion())
}

func (v *ReplVisitor) VisitAsignacionfor(ctx *parser.AsignacionforContext) interface{} {
	fmt.Println("📤 Asignación simple:", ctx.GetText())
	return v.Visit(ctx.Expresion())
}
