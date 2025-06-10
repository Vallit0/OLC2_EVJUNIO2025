// Code generated from parser/Vlang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // Vlang
import "github.com/antlr4-go/antlr/v4"

type BaseVlangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseVlangVisitor) VisitPrograma(ctx *ProgramaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitIfChain(ctx *IfChainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitElseStmt(ctx *ElseStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitPrintlnStmt(ctx *PrintlnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitPrintStmt(ctx *PrintStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitWhileStmt(ctx *WhileStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitForStmt(ctx *ForStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitNumericRange(ctx *NumericRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitFuncCall(ctx *FuncCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitFuncDecl(ctx *FuncDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitStructDecl(ctx *StructDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitBreakStmt(ctx *BreakStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitContinueStmt(ctx *ContinueStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitDirectAssign(ctx *DirectAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitDeclAssign(ctx *DeclAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitDeclAssignPattern(ctx *DeclAssignPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitDeclPattern(ctx *DeclPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitIdPattern(ctx *IdPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitCorchetesexpre(ctx *CorchetesexpreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitFuncionexpre(ctx *FuncionexpreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitUnario(ctx *UnarioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitParentesisexpre(ctx *ParentesisexpreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitIncredecr(ctx *IncredecrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitValorexpresion(ctx *ValorexpresionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitAsignacionfor(ctx *AsignacionforContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitExpdotexp1(ctx *Expdotexp1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitExpdotexp(ctx *ExpdotexpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitBinaryExp(ctx *BinaryExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitParametros(ctx *ParametrosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitValores(ctx *ValoresContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitValorEntero(ctx *ValorEnteroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitValorDecimal(ctx *ValorDecimalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitValorCadena(ctx *ValorCadenaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitValorBooleano(ctx *ValorBooleanoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitValorCaracter(ctx *ValorCaracterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitValorFloat(ctx *ValorFloatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitStructInstancia(ctx *StructInstanciaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitStructAttr(ctx *StructAttrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitIntType(ctx *IntTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitFloatType(ctx *FloatTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitStringType(ctx *StringTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitBoolType(ctx *BoolTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitCharType(ctx *CharTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitVoidType(ctx *VoidTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitStructVector(ctx *StructVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitIncremento(ctx *IncrementoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitDecremento(ctx *DecrementoContext) interface{} {
	return v.VisitChildren(ctx)
}
