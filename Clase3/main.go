package main

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"parser" // asegúrate de tener el import correcto
)

func main() {
	fmt.Println("Ingrese una expresión:")
	var input string
	_, err := fmt.Fscanln(os.Stdin, &input)
	if err != nil {
		fmt.Println("Error leyendo entrada:", err)
		return
	}

	// Lexer
	is := antlr.NewInputStream(input)
	lexer := parser.NewV4LangGrammarLexer(is)

	// Tokens
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Parser
	p := parser.NewV4LangGrammarParser(stream)
	tree := p.Expr()

	// Visitor
	visitor := &Interpreter{
		BaseV4LangGrammarVisitor: &parser.BaseV4LangGrammarVisitor{},
		Env:                      map[string]int{"x": 5},
	}

	// Evaluación
	result := tree.Accept(visitor).(int)
	fmt.Println("Resultado:", result)
}
