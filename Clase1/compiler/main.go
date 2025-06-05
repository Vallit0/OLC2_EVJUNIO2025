package main

import (
	"fmt"
	"os"

	compiler "compiler/parser"

	//"main/cst"
	"compiler/errors"
	// "main/repl"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	// Leer código fuente desde archivo o consola
	fmt.Println("Escribe el codigo de Vlang: (Ctrl+D) para terminar")
	inputCode, err := readStdin()
	if err != nil {
		fmt.Println("Error leyendo entrada:", err)
		return
	}

	//startTime := time.Now()

	// Canal para generar CST en paralelo (opcional)
	/*
		resultChannel := make(chan string)
		go func() {
			resultChannel <- cst.CstReport(inputCode)
		}()
	*/

	// 1. Análisis Léxico
	// Para verificar errores
	lexicalErrorListener := errors.NewLexicalErrorListener()
	//
	lexer := compiler.NewV4LangLexer(antlr.NewInputStream(inputCode))

	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexicalErrorListener)

	// 2. Tokens
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// 3. Parser + errores sintácticos
	// New<Nombre de mi gramatica>(Stream)
	parser := compiler.NewV4LangGrammar(stream)
	parser.BuildParseTrees = true

	syntaxErrorListener := errors.NewSyntaxErrorListener(lexicalErrorListener.ErrorTable)
	parser.RemoveErrorListeners()
	parser.SetErrorHandler(errors.NewCustomErrorStrategy())
	parser.AddErrorListener(syntaxErrorListener)

	// 4. Árbol sintáctico
	// En tu gramatica tienes el axioma, o simbolo inicial
	// Este es el que deberas agregar como parte del parser.
	//tree := parser.Program()

}

func readStdin() (string, error) {
	input, err := os.ReadFile("/dev/stdin")
	return string(input), err
}
