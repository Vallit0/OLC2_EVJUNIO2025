// Code generated from parser/Vlang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // Vlang
import "github.com/antlr4-go/antlr/v4"

// VlangListener is a complete listener for a parse tree produced by VlangParser.
type VlangListener interface {
	antlr.ParseTreeListener

	// EnterPrograma is called when entering the programa production.
	EnterPrograma(c *ProgramaContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterDirectAssign is called when entering the DirectAssign production.
	EnterDirectAssign(c *DirectAssignContext)

	// EnterDeclAssign is called when entering the DeclAssign production.
	EnterDeclAssign(c *DeclAssignContext)

	// EnterMutType is called when entering the mutType production.
	EnterMutType(c *MutTypeContext)

	// EnterIdPattern is called when entering the IdPattern production.
	EnterIdPattern(c *IdPatternContext)

	// EnterCorchetesexpre is called when entering the corchetesexpre production.
	EnterCorchetesexpre(c *CorchetesexpreContext)

	// EnterUnario is called when entering the unario production.
	EnterUnario(c *UnarioContext)

	// EnterParentesisexpre is called when entering the parentesisexpre production.
	EnterParentesisexpre(c *ParentesisexpreContext)

	// EnterIncredecr is called when entering the incredecr production.
	EnterIncredecr(c *IncredecrContext)

	// EnterValorexpresion is called when entering the valorexpresion production.
	EnterValorexpresion(c *ValorexpresionContext)

	// EnterAsignacionfor is called when entering the asignacionfor production.
	EnterAsignacionfor(c *AsignacionforContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// EnterExpdotexp1 is called when entering the expdotexp1 production.
	EnterExpdotexp1(c *Expdotexp1Context)

	// EnterExpdotexp is called when entering the expdotexp production.
	EnterExpdotexp(c *ExpdotexpContext)

	// EnterBinaryExp is called when entering the BinaryExp production.
	EnterBinaryExp(c *BinaryExpContext)

	// EnterParametros is called when entering the parametros production.
	EnterParametros(c *ParametrosContext)

	// EnterValores is called when entering the valores production.
	EnterValores(c *ValoresContext)

	// EnterValorEntero is called when entering the valorEntero production.
	EnterValorEntero(c *ValorEnteroContext)

	// EnterValorDecimal is called when entering the valorDecimal production.
	EnterValorDecimal(c *ValorDecimalContext)

	// EnterValorCadena is called when entering the valorCadena production.
	EnterValorCadena(c *ValorCadenaContext)

	// EnterValorBooleano is called when entering the valorBooleano production.
	EnterValorBooleano(c *ValorBooleanoContext)

	// EnterValorCaracter is called when entering the valorCaracter production.
	EnterValorCaracter(c *ValorCaracterContext)

	// EnterValorFloat is called when entering the valorFloat production.
	EnterValorFloat(c *ValorFloatContext)

	// EnterIncremento is called when entering the incremento production.
	EnterIncremento(c *IncrementoContext)

	// EnterDecremento is called when entering the decremento production.
	EnterDecremento(c *DecrementoContext)

	// ExitPrograma is called when exiting the programa production.
	ExitPrograma(c *ProgramaContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitDirectAssign is called when exiting the DirectAssign production.
	ExitDirectAssign(c *DirectAssignContext)

	// ExitDeclAssign is called when exiting the DeclAssign production.
	ExitDeclAssign(c *DeclAssignContext)

	// ExitMutType is called when exiting the mutType production.
	ExitMutType(c *MutTypeContext)

	// ExitIdPattern is called when exiting the IdPattern production.
	ExitIdPattern(c *IdPatternContext)

	// ExitCorchetesexpre is called when exiting the corchetesexpre production.
	ExitCorchetesexpre(c *CorchetesexpreContext)

	// ExitUnario is called when exiting the unario production.
	ExitUnario(c *UnarioContext)

	// ExitParentesisexpre is called when exiting the parentesisexpre production.
	ExitParentesisexpre(c *ParentesisexpreContext)

	// ExitIncredecr is called when exiting the incredecr production.
	ExitIncredecr(c *IncredecrContext)

	// ExitValorexpresion is called when exiting the valorexpresion production.
	ExitValorexpresion(c *ValorexpresionContext)

	// ExitAsignacionfor is called when exiting the asignacionfor production.
	ExitAsignacionfor(c *AsignacionforContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)

	// ExitExpdotexp1 is called when exiting the expdotexp1 production.
	ExitExpdotexp1(c *Expdotexp1Context)

	// ExitExpdotexp is called when exiting the expdotexp production.
	ExitExpdotexp(c *ExpdotexpContext)

	// ExitBinaryExp is called when exiting the BinaryExp production.
	ExitBinaryExp(c *BinaryExpContext)

	// ExitParametros is called when exiting the parametros production.
	ExitParametros(c *ParametrosContext)

	// ExitValores is called when exiting the valores production.
	ExitValores(c *ValoresContext)

	// ExitValorEntero is called when exiting the valorEntero production.
	ExitValorEntero(c *ValorEnteroContext)

	// ExitValorDecimal is called when exiting the valorDecimal production.
	ExitValorDecimal(c *ValorDecimalContext)

	// ExitValorCadena is called when exiting the valorCadena production.
	ExitValorCadena(c *ValorCadenaContext)

	// ExitValorBooleano is called when exiting the valorBooleano production.
	ExitValorBooleano(c *ValorBooleanoContext)

	// ExitValorCaracter is called when exiting the valorCaracter production.
	ExitValorCaracter(c *ValorCaracterContext)

	// ExitValorFloat is called when exiting the valorFloat production.
	ExitValorFloat(c *ValorFloatContext)

	// ExitIncremento is called when exiting the incremento production.
	ExitIncremento(c *IncrementoContext)

	// ExitDecremento is called when exiting the decremento production.
	ExitDecremento(c *DecrementoContext)
}
