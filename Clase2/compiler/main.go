package main

import (
	"fmt"
	"strings"

	"compiler/errors"
	parser "compiler/parser"
	"compiler/repl"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	a := app.NewWithID("com.vlang.ide")
	w := a.NewWindow("🧠 VLang IDE")

	// Entrada de texto ampliada
	codeEntry := widget.NewMultiLineEntry()
	codeEntry.Wrapping = fyne.TextTruncate
	codeEntry.SetPlaceHolder("// Escribe tu código VLang aquí...")

	// Expandir entrada
	codeScroll := container.NewVScroll(codeEntry)
	codeScroll.SetMinSize(fyne.NewSize(600, 300))

	// Salida
	output := widget.NewMultiLineEntry()
	output.Wrapping = fyne.TextWrapWord
	output.Disable()

	// Botón ejecutar
	runButton := widget.NewButtonWithIcon("Ejecutar", theme.MediaPlayIcon(), func() {
		result := runVlang(codeEntry.Text)
		output.SetText(result)
	})

	// Layout final
	editor := container.NewVBox(
		widget.NewLabel("💡 Editor VLang"),
		codeScroll,
		runButton,
		widget.NewLabel("📤 Salida"),
		container.NewVScroll(output),
	)

	w.SetContent(container.New(layout.NewPaddedLayout(), editor))
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}

func runVlang(code string) string {
	var builder strings.Builder

	lexicalErrs := errors.NewLexicalErrorListener()
	lexer := parser.NewVlangLexer(antlr.NewInputStream(code))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexicalErrs)

	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewVlangParser(tokens)
	p.BuildParseTrees = true
	syntaxErrs := errors.NewSyntaxErrorListener(lexicalErrs.ErrorTable)
	p.RemoveErrorListeners()
	p.SetErrorHandler(errors.NewCustomErrorStrategy())
	p.AddErrorListener(syntaxErrs)

	tree := p.Programa()
	visitor := repl.NewReplVisitor()
	result := visitor.Visit(tree)

	// if lexicalErrs.ErrorTable.HasErrors() {
	// 	builder.WriteString("Errores léxicos:\n")
	// 	for _, e := range lexicalErrs.ErrorTable.Errors {
	// 		builder.WriteString(fmt.Sprintf("  - %s\n", e.String()))
	// 	}
	// }
	// if syntaxErrs.ErrorTable.HasErrors() {
	// 	builder.WriteString("Errores sintácticos:\n")
	// 	for _, e := range syntaxErrs.ErrorTable.Errors {
	// 		builder.WriteString(fmt.Sprintf("  - %s\n", e.String()))
	// 	}
	// }
	if builder.Len() == 0 {
		builder.WriteString(fmt.Sprintf("✅ Resultado: %v\n", result))
	}

	return builder.String()
}
