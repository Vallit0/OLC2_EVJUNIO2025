// Code generated from parser/Vlang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // Vlang
import "github.com/antlr4-go/antlr/v4"

// BaseVlangListener is a complete listener for a parse tree produced by VlangParser.
type BaseVlangListener struct{}

var _ VlangListener = &BaseVlangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVlangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVlangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVlangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVlangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPrograma is called when production programa is entered.
func (s *BaseVlangListener) EnterPrograma(ctx *ProgramaContext) {}

// ExitPrograma is called when production programa is exited.
func (s *BaseVlangListener) ExitPrograma(ctx *ProgramaContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseVlangListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseVlangListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterDeclarationStatement is called when production declarationStatement is entered.
func (s *BaseVlangListener) EnterDeclarationStatement(ctx *DeclarationStatementContext) {}

// ExitDeclarationStatement is called when production declarationStatement is exited.
func (s *BaseVlangListener) ExitDeclarationStatement(ctx *DeclarationStatementContext) {}

// EnterAssignmentStatement is called when production assignmentStatement is entered.
func (s *BaseVlangListener) EnterAssignmentStatement(ctx *AssignmentStatementContext) {}

// ExitAssignmentStatement is called when production assignmentStatement is exited.
func (s *BaseVlangListener) ExitAssignmentStatement(ctx *AssignmentStatementContext) {}

// EnterPrintStatement is called when production printStatement is entered.
func (s *BaseVlangListener) EnterPrintStatement(ctx *PrintStatementContext) {}

// ExitPrintStatement is called when production printStatement is exited.
func (s *BaseVlangListener) ExitPrintStatement(ctx *PrintStatementContext) {}

// EnterDirectAssign is called when production DirectAssign is entered.
func (s *BaseVlangListener) EnterDirectAssign(ctx *DirectAssignContext) {}

// ExitDirectAssign is called when production DirectAssign is exited.
func (s *BaseVlangListener) ExitDirectAssign(ctx *DirectAssignContext) {}

// EnterDecl_stmt is called when production decl_stmt is entered.
func (s *BaseVlangListener) EnterDecl_stmt(ctx *Decl_stmtContext) {}

// ExitDecl_stmt is called when production decl_stmt is exited.
func (s *BaseVlangListener) ExitDecl_stmt(ctx *Decl_stmtContext) {}

// EnterMutType is called when production mutType is entered.
func (s *BaseVlangListener) EnterMutType(ctx *MutTypeContext) {}

// ExitMutType is called when production mutType is exited.
func (s *BaseVlangListener) ExitMutType(ctx *MutTypeContext) {}

// EnterCorchetesexpre is called when production corchetesexpre is entered.
func (s *BaseVlangListener) EnterCorchetesexpre(ctx *CorchetesexpreContext) {}

// ExitCorchetesexpre is called when production corchetesexpre is exited.
func (s *BaseVlangListener) ExitCorchetesexpre(ctx *CorchetesexpreContext) {}

// EnterUnario is called when production unario is entered.
func (s *BaseVlangListener) EnterUnario(ctx *UnarioContext) {}

// ExitUnario is called when production unario is exited.
func (s *BaseVlangListener) ExitUnario(ctx *UnarioContext) {}

// EnterParentesisexpre is called when production parentesisexpre is entered.
func (s *BaseVlangListener) EnterParentesisexpre(ctx *ParentesisexpreContext) {}

// ExitParentesisexpre is called when production parentesisexpre is exited.
func (s *BaseVlangListener) ExitParentesisexpre(ctx *ParentesisexpreContext) {}

// EnterIncredecr is called when production incredecr is entered.
func (s *BaseVlangListener) EnterIncredecr(ctx *IncredecrContext) {}

// ExitIncredecr is called when production incredecr is exited.
func (s *BaseVlangListener) ExitIncredecr(ctx *IncredecrContext) {}

// EnterValorexpresion is called when production valorexpresion is entered.
func (s *BaseVlangListener) EnterValorexpresion(ctx *ValorexpresionContext) {}

// ExitValorexpresion is called when production valorexpresion is exited.
func (s *BaseVlangListener) ExitValorexpresion(ctx *ValorexpresionContext) {}

// EnterAsignacionfor is called when production asignacionfor is entered.
func (s *BaseVlangListener) EnterAsignacionfor(ctx *AsignacionforContext) {}

// ExitAsignacionfor is called when production asignacionfor is exited.
func (s *BaseVlangListener) ExitAsignacionfor(ctx *AsignacionforContext) {}

// EnterId is called when production id is entered.
func (s *BaseVlangListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseVlangListener) ExitId(ctx *IdContext) {}

// EnterExpdotexp1 is called when production expdotexp1 is entered.
func (s *BaseVlangListener) EnterExpdotexp1(ctx *Expdotexp1Context) {}

// ExitExpdotexp1 is called when production expdotexp1 is exited.
func (s *BaseVlangListener) ExitExpdotexp1(ctx *Expdotexp1Context) {}

// EnterExpdotexp is called when production expdotexp is entered.
func (s *BaseVlangListener) EnterExpdotexp(ctx *ExpdotexpContext) {}

// ExitExpdotexp is called when production expdotexp is exited.
func (s *BaseVlangListener) ExitExpdotexp(ctx *ExpdotexpContext) {}

// EnterBinaryExp is called when production BinaryExp is entered.
func (s *BaseVlangListener) EnterBinaryExp(ctx *BinaryExpContext) {}

// ExitBinaryExp is called when production BinaryExp is exited.
func (s *BaseVlangListener) ExitBinaryExp(ctx *BinaryExpContext) {}

// EnterParametros is called when production parametros is entered.
func (s *BaseVlangListener) EnterParametros(ctx *ParametrosContext) {}

// ExitParametros is called when production parametros is exited.
func (s *BaseVlangListener) ExitParametros(ctx *ParametrosContext) {}

// EnterValores is called when production valores is entered.
func (s *BaseVlangListener) EnterValores(ctx *ValoresContext) {}

// ExitValores is called when production valores is exited.
func (s *BaseVlangListener) ExitValores(ctx *ValoresContext) {}

// EnterValorEntero is called when production valorEntero is entered.
func (s *BaseVlangListener) EnterValorEntero(ctx *ValorEnteroContext) {}

// ExitValorEntero is called when production valorEntero is exited.
func (s *BaseVlangListener) ExitValorEntero(ctx *ValorEnteroContext) {}

// EnterValorDecimal is called when production valorDecimal is entered.
func (s *BaseVlangListener) EnterValorDecimal(ctx *ValorDecimalContext) {}

// ExitValorDecimal is called when production valorDecimal is exited.
func (s *BaseVlangListener) ExitValorDecimal(ctx *ValorDecimalContext) {}

// EnterValorCadena is called when production valorCadena is entered.
func (s *BaseVlangListener) EnterValorCadena(ctx *ValorCadenaContext) {}

// ExitValorCadena is called when production valorCadena is exited.
func (s *BaseVlangListener) ExitValorCadena(ctx *ValorCadenaContext) {}

// EnterValorBooleano is called when production valorBooleano is entered.
func (s *BaseVlangListener) EnterValorBooleano(ctx *ValorBooleanoContext) {}

// ExitValorBooleano is called when production valorBooleano is exited.
func (s *BaseVlangListener) ExitValorBooleano(ctx *ValorBooleanoContext) {}

// EnterValorCaracter is called when production valorCaracter is entered.
func (s *BaseVlangListener) EnterValorCaracter(ctx *ValorCaracterContext) {}

// ExitValorCaracter is called when production valorCaracter is exited.
func (s *BaseVlangListener) ExitValorCaracter(ctx *ValorCaracterContext) {}

// EnterValorFloat is called when production valorFloat is entered.
func (s *BaseVlangListener) EnterValorFloat(ctx *ValorFloatContext) {}

// ExitValorFloat is called when production valorFloat is exited.
func (s *BaseVlangListener) ExitValorFloat(ctx *ValorFloatContext) {}

// EnterIncremento is called when production incremento is entered.
func (s *BaseVlangListener) EnterIncremento(ctx *IncrementoContext) {}

// ExitIncremento is called when production incremento is exited.
func (s *BaseVlangListener) ExitIncremento(ctx *IncrementoContext) {}

// EnterDecremento is called when production decremento is entered.
func (s *BaseVlangListener) EnterDecremento(ctx *DecrementoContext) {}

// ExitDecremento is called when production decremento is exited.
func (s *BaseVlangListener) ExitDecremento(ctx *DecrementoContext) {}
