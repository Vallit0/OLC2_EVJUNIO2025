// Code generated from parser/V4LangGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // V4LangGrammar
import "github.com/antlr4-go/antlr/v4"

// V4LangGrammarListener is a complete listener for a parse tree produced by V4LangGrammar.
type V4LangGrammarListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterParens is called when entering the Parens production.
	EnterParens(c *ParensContext)

	// EnterInt is called when entering the Int production.
	EnterInt(c *IntContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitParens is called when exiting the Parens production.
	ExitParens(c *ParensContext)

	// ExitInt is called when exiting the Int production.
	ExitInt(c *IntContext)
}
