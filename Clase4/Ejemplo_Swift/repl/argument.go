package repl

import (
	"main/value"

	"github.com/antlr4-go/antlr/v4"
)

type Argument struct {
	Name            string
	Value           value.IVOR
	PassByReference bool
	Token           antlr.Token
	VariableRef     *Variable
}
