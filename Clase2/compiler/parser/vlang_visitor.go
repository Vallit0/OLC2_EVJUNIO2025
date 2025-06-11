// Code generated from parser/Vlang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // Vlang
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by VlangParser.
type VlangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by VlangParser#programa.
	VisitPrograma(ctx *ProgramaContext) interface{}

	// Visit a parse tree produced by VlangParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by VlangParser#IfStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by VlangParser#IfChain.
	VisitIfChain(ctx *IfChainContext) interface{}

	// Visit a parse tree produced by VlangParser#ElseStmt.
	VisitElseStmt(ctx *ElseStmtContext) interface{}

	// Visit a parse tree produced by VlangParser#PrintlnStmt.
	VisitPrintlnStmt(ctx *PrintlnStmtContext) interface{}

	// Visit a parse tree produced by VlangParser#PrintStmt.
	VisitPrintStmt(ctx *PrintStmtContext) interface{}

	// Visit a parse tree produced by VlangParser#WhileStmt.
	VisitWhileStmt(ctx *WhileStmtContext) interface{}

	// Visit a parse tree produced by VlangParser#ForStmt.
	VisitForStmt(ctx *ForStmtContext) interface{}

	// Visit a parse tree produced by VlangParser#NumericRange.
	VisitNumericRange(ctx *NumericRangeContext) interface{}

	// Visit a parse tree produced by VlangParser#FuncCall.
	VisitFuncCall(ctx *FuncCallContext) interface{}

	// Visit a parse tree produced by VlangParser#FuncDecl.
	VisitFuncDecl(ctx *FuncDeclContext) interface{}

	// Visit a parse tree produced by VlangParser#StructDecl.
	VisitStructDecl(ctx *StructDeclContext) interface{}

	// Visit a parse tree produced by VlangParser#ReturnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by VlangParser#BreakStmt.
	VisitBreakStmt(ctx *BreakStmtContext) interface{}

	// Visit a parse tree produced by VlangParser#ContinueStmt.
	VisitContinueStmt(ctx *ContinueStmtContext) interface{}

	// Visit a parse tree produced by VlangParser#DirectAssign.
	VisitDirectAssign(ctx *DirectAssignContext) interface{}

	// Visit a parse tree produced by VlangParser#DeclAssign.
	VisitDeclAssign(ctx *DeclAssignContext) interface{}

	// Visit a parse tree produced by VlangParser#DeclAssignType.
	VisitDeclAssignType(ctx *DeclAssignTypeContext) interface{}

	// Visit a parse tree produced by VlangParser#IdPattern.
	VisitIdPattern(ctx *IdPatternContext) interface{}

	// Visit a parse tree produced by VlangParser#corchetesexpre.
	VisitCorchetesexpre(ctx *CorchetesexpreContext) interface{}

	// Visit a parse tree produced by VlangParser#funcionexpre.
	VisitFuncionexpre(ctx *FuncionexpreContext) interface{}

	// Visit a parse tree produced by VlangParser#unario.
	VisitUnario(ctx *UnarioContext) interface{}

	// Visit a parse tree produced by VlangParser#parentesisexpre.
	VisitParentesisexpre(ctx *ParentesisexpreContext) interface{}

	// Visit a parse tree produced by VlangParser#incredecr.
	VisitIncredecr(ctx *IncredecrContext) interface{}

	// Visit a parse tree produced by VlangParser#valorexpresion.
	VisitValorexpresion(ctx *ValorexpresionContext) interface{}

	// Visit a parse tree produced by VlangParser#asignacionfor.
	VisitAsignacionfor(ctx *AsignacionforContext) interface{}

	// Visit a parse tree produced by VlangParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by VlangParser#expdotexp1.
	VisitExpdotexp1(ctx *Expdotexp1Context) interface{}

	// Visit a parse tree produced by VlangParser#expdotexp.
	VisitExpdotexp(ctx *ExpdotexpContext) interface{}

	// Visit a parse tree produced by VlangParser#BinaryExp.
	VisitBinaryExp(ctx *BinaryExpContext) interface{}

	// Visit a parse tree produced by VlangParser#ParamList.
	VisitParamList(ctx *ParamListContext) interface{}

	// Visit a parse tree produced by VlangParser#funcParam.
	VisitFuncParam(ctx *FuncParamContext) interface{}

	// Visit a parse tree produced by VlangParser#ArgList.
	VisitArgList(ctx *ArgListContext) interface{}

	// Visit a parse tree produced by VlangParser#FuncArg.
	VisitFuncArg(ctx *FuncArgContext) interface{}

	// Visit a parse tree produced by VlangParser#valores.
	VisitValores(ctx *ValoresContext) interface{}

	// Visit a parse tree produced by VlangParser#valorEntero.
	VisitValorEntero(ctx *ValorEnteroContext) interface{}

	// Visit a parse tree produced by VlangParser#valorDecimal.
	VisitValorDecimal(ctx *ValorDecimalContext) interface{}

	// Visit a parse tree produced by VlangParser#valorCadena.
	VisitValorCadena(ctx *ValorCadenaContext) interface{}

	// Visit a parse tree produced by VlangParser#valorBooleano.
	VisitValorBooleano(ctx *ValorBooleanoContext) interface{}

	// Visit a parse tree produced by VlangParser#valorCaracter.
	VisitValorCaracter(ctx *ValorCaracterContext) interface{}

	// Visit a parse tree produced by VlangParser#valorFloat.
	VisitValorFloat(ctx *ValorFloatContext) interface{}

	// Visit a parse tree produced by VlangParser#StructInstancia.
	VisitStructInstancia(ctx *StructInstanciaContext) interface{}

	// Visit a parse tree produced by VlangParser#StructAttr.
	VisitStructAttr(ctx *StructAttrContext) interface{}

	// Visit a parse tree produced by VlangParser#IntType.
	VisitIntType(ctx *IntTypeContext) interface{}

	// Visit a parse tree produced by VlangParser#FloatType.
	VisitFloatType(ctx *FloatTypeContext) interface{}

	// Visit a parse tree produced by VlangParser#StringType.
	VisitStringType(ctx *StringTypeContext) interface{}

	// Visit a parse tree produced by VlangParser#BoolType.
	VisitBoolType(ctx *BoolTypeContext) interface{}

	// Visit a parse tree produced by VlangParser#CharType.
	VisitCharType(ctx *CharTypeContext) interface{}

	// Visit a parse tree produced by VlangParser#VoidType.
	VisitVoidType(ctx *VoidTypeContext) interface{}

	// Visit a parse tree produced by VlangParser#StructVector.
	VisitStructVector(ctx *StructVectorContext) interface{}

	// Visit a parse tree produced by VlangParser#incremento.
	VisitIncremento(ctx *IncrementoContext) interface{}

	// Visit a parse tree produced by VlangParser#decremento.
	VisitDecremento(ctx *DecrementoContext) interface{}
}
