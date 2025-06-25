package repl

import (
	// parser "compiler/parser"

	compiler "compiler/parser"

	"github.com/antlr4-go/antlr/v4"
)

type Struct struct {
	Name string
	// Struct Name -> interfaz
	Fields []compiler.IStruct_propContext
	Token  antlr.Token
}
