package main

import (
	"fmt"
	"main/arm"
	"main/compiler"
	"main/cst"
	"main/errors"
	"main/repl"
	"time"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Post("/compile", func(c *fiber.Ctx) error {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r)
			}
		}()

		startTime := time.Now()

		resultChannel := make(chan string)

		go func() {
			resultChannel <- cst.CstReport(c.FormValue("code"))
		}()

		code := c.FormValue("code")

		lexicalErrorListener := errors.NewLexicalErrorListener()
		lexer := compiler.NewTSwiftLexer(antlr.NewInputStream(code))

		lexer.RemoveErrorListeners()
		lexer.AddErrorListener(lexicalErrorListener)

		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		parser := compiler.NewTSwiftLanguage(stream)
		parser.BuildParseTrees = true

		syntaxErrorListener := errors.NewSyntaxErrorListener(lexicalErrorListener.ErrorTable)
		parser.RemoveErrorListeners()
		parser.SetErrorHandler(errors.NewCustomErrorStrategy())
		parser.AddErrorListener(syntaxErrorListener)

		tree := parser.Program()

		dclVisitor := repl.NewDclVisitor(syntaxErrorListener.ErrorTable)
		dclVisitor.Visit(tree)

		replVisitor := repl.NewVisitor(dclVisitor)

		replVisitor.Visit(tree)
		// Medir el tiempo de compilacion

		/*

			Inicia la traudccion a ARM64
			-> Todo lo de ejecucion anterior
			-> Se traduce a ARM64

			// le mandamos el ScopeTrace
			Recorrer el ScopeTrace y Apartar el Espacio en los registros
			ScopeTrace -> Map[Variables][Valores]

		*/
		// ARM Visitor
		armVisitor := arm.NewArmVisitor()
		fmt.Println(armVisitor)
		intepretationEndTime := time.Now()

		cstReport := <-resultChannel

		reportEndTime := time.Now()
		// replVisitor.ScopeTrace.Print()
		// replVisitor.Console.Show()
		fmt.Println("Interpretation finished")

		fmt.Println("Interpretation time:", intepretationEndTime.Sub(startTime))
		fmt.Println("Total (with report) time:", reportEndTime.Sub(intepretationEndTime))
		fmt.Println("")

		return c.JSON(struct {
			Errors     []repl.Error     `json:"errors"`
			Output     string           `json:"output"`
			CSTSvg     string           `json:"cstSvg"`
			ScopeTrace repl.ReportTable `json:"scopeTrace"`
		}{
			Errors:     replVisitor.ErrorTable.Errors,
			Output:     replVisitor.Console.GetOutput(),
			CSTSvg:     cstReport,
			ScopeTrace: replVisitor.ScopeTrace.Report(),
		})

	})

	app.Post("/arm", func(c *fiber.Ctx) error {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r)
			}
		}()

		startTime := time.Now()

		resultChannel := make(chan string)

		go func() {
			resultChannel <- cst.CstReport(c.FormValue("code"))
		}()

		code := c.FormValue("code")

		lexicalErrorListener := errors.NewLexicalErrorListener()
		lexer := compiler.NewTSwiftLexer(antlr.NewInputStream(code))

		lexer.RemoveErrorListeners()
		lexer.AddErrorListener(lexicalErrorListener)

		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		parser := compiler.NewTSwiftLanguage(stream)
		parser.BuildParseTrees = true

		syntaxErrorListener := errors.NewSyntaxErrorListener(lexicalErrorListener.ErrorTable)
		parser.RemoveErrorListeners()
		parser.SetErrorHandler(errors.NewCustomErrorStrategy())
		parser.AddErrorListener(syntaxErrorListener)

		tree := parser.Program()

		dclVisitor := repl.NewDclVisitor(syntaxErrorListener.ErrorTable)
		dclVisitor.Visit(tree)

		replVisitor := repl.NewVisitor(dclVisitor)

		replVisitor.Visit(tree)
		// Medir el tiempo de compilacion
		intepretationEndTime := time.Now()

		cstReport := <-resultChannel

		reportEndTime := time.Now()
		// replVisitor.ScopeTrace.Print()
		// replVisitor.Console.Show()
		fmt.Println("Interpretation finished")

		fmt.Println("Interpretation time:", intepretationEndTime.Sub(startTime))
		fmt.Println("Total (with report) time:", reportEndTime.Sub(intepretationEndTime))
		fmt.Println("")

		return c.JSON(struct {
			Errors     []repl.Error     `json:"errors"`
			Output     string           `json:"output"`
			CSTSvg     string           `json:"cstSvg"`
			ScopeTrace repl.ReportTable `json:"scopeTrace"`
		}{
			Errors:     replVisitor.ErrorTable.Errors,
			Output:     replVisitor.Console.GetOutput(),
			CSTSvg:     cstReport,
			ScopeTrace: replVisitor.ScopeTrace.Report(),
		})

	})

	// Endpoints Tecnologias ->
	app.Listen(":3000")

}
