// Code generated from parser/Vlang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // Vlang
import "github.com/antlr4-go/antlr/v4"

type BaseVlangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseVlangVisitor) VisitPrograma(ctx *ProgramaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitDeclarationStatement(ctx *DeclarationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitAssignmentStatement(ctx *AssignmentStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitPrintStatement(ctx *PrintStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitDirectAssign(ctx *DirectAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitDecl_stmt(ctx *Decl_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitMutType(ctx *MutTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitCorchetesexpre(ctx *CorchetesexpreContext) interface{} {
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

func (v *BaseVlangVisitor) VisitIncremento(ctx *IncrementoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitDecremento(ctx *DecrementoContext) interface{} {
	return v.VisitChildren(ctx)
}
