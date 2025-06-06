package repl

import (
	parser "compiler/parser"
	"fmt"
)

// Este visitor -> Interpretar el codigo

type ReplVisitor struct {
	*parser.BaseVlangVisitor
	// CallStack
	// ScopeTrace
	// Scope *BaseScope
	// Tabla de Error
	// ErrorTable *ErrorTable
	// Console
	//Console *Console
	// append

}

// fmt.Println(visitor.console)
// Agregamos un Constructor para inicializar el visitor
func NewReplVisitor() *ReplVisitor {
	return &ReplVisitor{}
}

// Base para recorrer todas nuestras declaraciones
func (v *ReplVisitor) VisitPrograma(ctx *parser.ProgramaContext) interface{} {

	for _, declaraciones := range ctx.AllExpresion() {
		fmt.Println("Visitando una declaracion")
		v.Visit(declaraciones)
	}

	return v.VisitChildren(ctx)

}
