package repl

import (
	parser "compiler/parser"
)

// Visitor personalizado para recorrer el árbol de sintaxis

// Constructor del visitor
type ARMVisitor struct {
	parser.BaseVlangVisitor
	ScopeTrace  *ScopeTrace
	CallStack   *CallStack
	Console     *Console
	StructNames []string
}
