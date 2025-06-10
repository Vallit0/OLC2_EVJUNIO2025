package main

import (
	"fmt"
	"os"

	compiler "compiler/parser"
	repl "compiler/repl"

	"compiler/errors"
	// "main/repl"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	// inicializamos la aplicacion con fiber

	// Leer código fuente desde archivo o consola
	fmt.Println("Escribe el codigo de Vlang: (Ctrl+D) para terminar")
	inputCode, err := readStdin()
	if err != nil {
		fmt.Println("Error leyendo entrada:", err)
		return
	}

	//startTime := time.Now()

	// Canal para generar CST en paralelo (opcional)

	// 1. Análisis Léxico
	// Para verificar errores
	lexicalErrorListener := errors.NewLexicalErrorListener()
	//
	lexer := compiler.NewVlangLexer(antlr.NewInputStream(inputCode))

	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexicalErrorListener)

	// 2. Tokens
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// 3. Parser + errores sintácticos
	// New<Nombre de mi gramatica>(Stream)
	parser := compiler.NewVlangParser(stream)
	parser.BuildParseTrees = true

	syntaxErrorListener := errors.NewSyntaxErrorListener(lexicalErrorListener.ErrorTable)
	parser.RemoveErrorListeners()
	parser.SetErrorHandler(errors.NewCustomErrorStrategy())
	parser.AddErrorListener(syntaxErrorListener)
	fmt.Println("Se termino el analisis sintactico")

	// 4. Árbol sintáctico
	// En tu gramatica tienes el axioma, o simbolo inicial
	// Este es el que deberas agregar como parte del parser.
	arbolito := parser.Programa()

	// verificamos que el arbol si funcione
	fmt.Printf("🌳 Árbol raíz: %T\n", arbolito)

	// imprimimos los errores de sintaxis y léxicos
	/*
		NOTA: Para mejorar el rendimiento y facilidad de las estructuras
		lo primero que debemos hacer es un visitor que declare funciones,
		structs, y algunos vectores.

		// -> Crear los entornos de la funciones
		// -> Anadir los vectores a esos entornos
		// Se define el main()

	*/
	// ----> Este visitor -> Crear entornos
	// dclVisitor := repl.NewDclVisitor(syntaxErrorListener.ErrorTable)
	// dclVisitor.Visit(tree)
	/*
		En visitor.go tenemos un metodo en el que INICIALIZAMOS
		ReplVisitor PERO le enviamos dclVisitor como parametro
		para que este pueda visitar los nodos de declaración

		En visitor.go tambien tenemos un metodo para inicializar
		ReplVisitor y enviarle el errorTable, este nos sirve
		para testear
	*/
	fmt.Println("Visitamos los nodos del árbol sintáctico")

	visitor := repl.NewReplVisitor(syntaxErrorListener.ErrorTable)

	visitor.Visit(arbolito)
	visitor.Console.Show()
	/*

		En console.GetOutput() tenemos la salida de la consola
	*/
	nuevaVariable := visitor.Console.GetOutput()
	fmt.Println("Salida de la consola:", nuevaVariable)
	//
}

func readStdin() (string, error) {
	input, err := os.ReadFile("/dev/stdin")
	return string(input), err
}

// Funciones para visualizar nuestro arbol
func PrintVerticalTree(node antlr.Tree, ruleNames []string) {
	printVerticalNode(node, ruleNames, "", true)
}

func printVerticalNode(node antlr.Tree, ruleNames []string, prefix string, isLast bool) {
	connector := "+-- "
	if !isLast {
		connector = "|-- "
	}

	var label string
	switch n := node.(type) {
	case antlr.RuleNode:
		label = ruleNames[n.GetRuleContext().GetRuleIndex()]
	case antlr.TerminalNode:
		label = fmt.Sprintf("\"%s\"", n.GetText())
	default:
		label = fmt.Sprintf("%T", n)
	}

	fmt.Printf("%s%s%s\n", prefix, connector, label)

	// Actualizar el prefijo para los hijos
	childCount := node.GetChildCount()
	for i := 0; i < childCount; i++ {
		child := node.GetChild(i)
		newPrefix := prefix
		if isLast {
			newPrefix += "    "
		} else {
			newPrefix += "|   "
		}
		printVerticalNode(child, ruleNames, newPrefix, i == childCount-1)
	}
}
