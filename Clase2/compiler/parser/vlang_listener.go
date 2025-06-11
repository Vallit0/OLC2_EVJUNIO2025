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

	// EnterIfStmt is called when entering the IfStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterIfChain is called when entering the IfChain production.
	EnterIfChain(c *IfChainContext)

	// EnterElseStmt is called when entering the ElseStmt production.
	EnterElseStmt(c *ElseStmtContext)

	// EnterPrintlnStmt is called when entering the PrintlnStmt production.
	EnterPrintlnStmt(c *PrintlnStmtContext)

	// EnterPrintStmt is called when entering the PrintStmt production.
	EnterPrintStmt(c *PrintStmtContext)

	// EnterWhileStmt is called when entering the WhileStmt production.
	EnterWhileStmt(c *WhileStmtContext)

	// EnterForStmt is called when entering the ForStmt production.
	EnterForStmt(c *ForStmtContext)

	// EnterNumericRange is called when entering the NumericRange production.
	EnterNumericRange(c *NumericRangeContext)

	// EnterFuncCall is called when entering the FuncCall production.
	EnterFuncCall(c *FuncCallContext)

	// EnterFuncDecl is called when entering the FuncDecl production.
	EnterFuncDecl(c *FuncDeclContext)

	// EnterStructDecl is called when entering the StructDecl production.
	EnterStructDecl(c *StructDeclContext)

	// EnterReturnStmt is called when entering the ReturnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterBreakStmt is called when entering the BreakStmt production.
	EnterBreakStmt(c *BreakStmtContext)

	// EnterContinueStmt is called when entering the ContinueStmt production.
	EnterContinueStmt(c *ContinueStmtContext)

	// EnterDirectAssign is called when entering the DirectAssign production.
	EnterDirectAssign(c *DirectAssignContext)

	// EnterDeclAssign is called when entering the DeclAssign production.
	EnterDeclAssign(c *DeclAssignContext)

	// EnterDeclAssignType is called when entering the DeclAssignType production.
	EnterDeclAssignType(c *DeclAssignTypeContext)

	// EnterIdPattern is called when entering the IdPattern production.
	EnterIdPattern(c *IdPatternContext)

	// EnterCorchetesexpre is called when entering the corchetesexpre production.
	EnterCorchetesexpre(c *CorchetesexpreContext)

	// EnterFuncionexpre is called when entering the funcionexpre production.
	EnterFuncionexpre(c *FuncionexpreContext)

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

	// EnterStructInstancia is called when entering the StructInstancia production.
	EnterStructInstancia(c *StructInstanciaContext)

	// EnterStructAttr is called when entering the StructAttr production.
	EnterStructAttr(c *StructAttrContext)

	// EnterIntType is called when entering the IntType production.
	EnterIntType(c *IntTypeContext)

	// EnterFloatType is called when entering the FloatType production.
	EnterFloatType(c *FloatTypeContext)

	// EnterStringType is called when entering the StringType production.
	EnterStringType(c *StringTypeContext)

	// EnterBoolType is called when entering the BoolType production.
	EnterBoolType(c *BoolTypeContext)

	// EnterCharType is called when entering the CharType production.
	EnterCharType(c *CharTypeContext)

	// EnterVoidType is called when entering the VoidType production.
	EnterVoidType(c *VoidTypeContext)

	// EnterStructVector is called when entering the StructVector production.
	EnterStructVector(c *StructVectorContext)

	// EnterIncremento is called when entering the incremento production.
	EnterIncremento(c *IncrementoContext)

	// EnterDecremento is called when entering the decremento production.
	EnterDecremento(c *DecrementoContext)

	// ExitPrograma is called when exiting the programa production.
	ExitPrograma(c *ProgramaContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitIfStmt is called when exiting the IfStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitIfChain is called when exiting the IfChain production.
	ExitIfChain(c *IfChainContext)

	// ExitElseStmt is called when exiting the ElseStmt production.
	ExitElseStmt(c *ElseStmtContext)

	// ExitPrintlnStmt is called when exiting the PrintlnStmt production.
	ExitPrintlnStmt(c *PrintlnStmtContext)

	// ExitPrintStmt is called when exiting the PrintStmt production.
	ExitPrintStmt(c *PrintStmtContext)

	// ExitWhileStmt is called when exiting the WhileStmt production.
	ExitWhileStmt(c *WhileStmtContext)

	// ExitForStmt is called when exiting the ForStmt production.
	ExitForStmt(c *ForStmtContext)

	// ExitNumericRange is called when exiting the NumericRange production.
	ExitNumericRange(c *NumericRangeContext)

	// ExitFuncCall is called when exiting the FuncCall production.
	ExitFuncCall(c *FuncCallContext)

	// ExitFuncDecl is called when exiting the FuncDecl production.
	ExitFuncDecl(c *FuncDeclContext)

	// ExitStructDecl is called when exiting the StructDecl production.
	ExitStructDecl(c *StructDeclContext)

	// ExitReturnStmt is called when exiting the ReturnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitBreakStmt is called when exiting the BreakStmt production.
	ExitBreakStmt(c *BreakStmtContext)

	// ExitContinueStmt is called when exiting the ContinueStmt production.
	ExitContinueStmt(c *ContinueStmtContext)

	// ExitDirectAssign is called when exiting the DirectAssign production.
	ExitDirectAssign(c *DirectAssignContext)

	// ExitDeclAssign is called when exiting the DeclAssign production.
	ExitDeclAssign(c *DeclAssignContext)

	// ExitDeclAssignType is called when exiting the DeclAssignType production.
	ExitDeclAssignType(c *DeclAssignTypeContext)

	// ExitIdPattern is called when exiting the IdPattern production.
	ExitIdPattern(c *IdPatternContext)

	// ExitCorchetesexpre is called when exiting the corchetesexpre production.
	ExitCorchetesexpre(c *CorchetesexpreContext)

	// ExitFuncionexpre is called when exiting the funcionexpre production.
	ExitFuncionexpre(c *FuncionexpreContext)

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

	// ExitStructInstancia is called when exiting the StructInstancia production.
	ExitStructInstancia(c *StructInstanciaContext)

	// ExitStructAttr is called when exiting the StructAttr production.
	ExitStructAttr(c *StructAttrContext)

	// ExitIntType is called when exiting the IntType production.
	ExitIntType(c *IntTypeContext)

	// ExitFloatType is called when exiting the FloatType production.
	ExitFloatType(c *FloatTypeContext)

	// ExitStringType is called when exiting the StringType production.
	ExitStringType(c *StringTypeContext)

	// ExitBoolType is called when exiting the BoolType production.
	ExitBoolType(c *BoolTypeContext)

	// ExitCharType is called when exiting the CharType production.
	ExitCharType(c *CharTypeContext)

	// ExitVoidType is called when exiting the VoidType production.
	ExitVoidType(c *VoidTypeContext)

	// ExitStructVector is called when exiting the StructVector production.
	ExitStructVector(c *StructVectorContext)

	// ExitIncremento is called when exiting the incremento production.
	ExitIncremento(c *IncrementoContext)

	// ExitDecremento is called when exiting the decremento production.
	ExitDecremento(c *DecrementoContext)
}
