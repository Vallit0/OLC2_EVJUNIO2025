package repl

import (
	"main/compiler"

	"github.com/antlr4-go/antlr/v4"
)

type Struct struct {
	Name   string
	Fields []compiler.IStruct_propContext
	Token  antlr.Token
}
