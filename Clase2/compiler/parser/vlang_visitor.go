// Code generated from parser/Vlang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // Vlang
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by VlangParser.
type VlangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by VlangParser#programa.
	VisitPrograma(ctx *ProgramaContext) interface{}

	// Visit a parse tree produced by VlangParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by VlangParser#declarationStatement.
	VisitDeclarationStatement(ctx *DeclarationStatementContext) interface{}

	// Visit a parse tree produced by VlangParser#assignmentStatement.
	VisitAssignmentStatement(ctx *AssignmentStatementContext) interface{}

	// Visit a parse tree produced by VlangParser#printStatement.
	VisitPrintStatement(ctx *PrintStatementContext) interface{}

	// Visit a parse tree produced by VlangParser#DirectAssign.
	VisitDirectAssign(ctx *DirectAssignContext) interface{}

	// Visit a parse tree produced by VlangParser#decl_stmt.
	VisitDecl_stmt(ctx *Decl_stmtContext) interface{}

	// Visit a parse tree produced by VlangParser#mutType.
	VisitMutType(ctx *MutTypeContext) interface{}

	// Visit a parse tree produced by VlangParser#corchetesexpre.
	VisitCorchetesexpre(ctx *CorchetesexpreContext) interface{}

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

	// Visit a parse tree produced by VlangParser#parametros.
	VisitParametros(ctx *ParametrosContext) interface{}

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

	// Visit a parse tree produced by VlangParser#incremento.
	VisitIncremento(ctx *IncrementoContext) interface{}

	// Visit a parse tree produced by VlangParser#decremento.
	VisitDecremento(ctx *DecrementoContext) interface{}
}
