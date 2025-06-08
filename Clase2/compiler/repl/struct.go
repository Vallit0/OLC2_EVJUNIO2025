package repl

import (
	// parser "compiler/parser"

	"github.com/antlr4-go/antlr/v4"
)

type Struct struct {
	Name string
	// Fields []parser.IStruct_propContext
	Token antlr.Token
}
