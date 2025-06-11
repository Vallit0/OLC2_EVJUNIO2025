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

// EnterStmt is called when production stmt is entered.
func (s *BaseVlangListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseVlangListener) ExitStmt(ctx *StmtContext) {}

// EnterIfStmt is called when production IfStmt is entered.
func (s *BaseVlangListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production IfStmt is exited.
func (s *BaseVlangListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterIfChain is called when production IfChain is entered.
func (s *BaseVlangListener) EnterIfChain(ctx *IfChainContext) {}

// ExitIfChain is called when production IfChain is exited.
func (s *BaseVlangListener) ExitIfChain(ctx *IfChainContext) {}

// EnterElseStmt is called when production ElseStmt is entered.
func (s *BaseVlangListener) EnterElseStmt(ctx *ElseStmtContext) {}

// ExitElseStmt is called when production ElseStmt is exited.
func (s *BaseVlangListener) ExitElseStmt(ctx *ElseStmtContext) {}

// EnterPrintlnStmt is called when production PrintlnStmt is entered.
func (s *BaseVlangListener) EnterPrintlnStmt(ctx *PrintlnStmtContext) {}

// ExitPrintlnStmt is called when production PrintlnStmt is exited.
func (s *BaseVlangListener) ExitPrintlnStmt(ctx *PrintlnStmtContext) {}

// EnterPrintStmt is called when production PrintStmt is entered.
func (s *BaseVlangListener) EnterPrintStmt(ctx *PrintStmtContext) {}

// ExitPrintStmt is called when production PrintStmt is exited.
func (s *BaseVlangListener) ExitPrintStmt(ctx *PrintStmtContext) {}

// EnterWhileStmt is called when production WhileStmt is entered.
func (s *BaseVlangListener) EnterWhileStmt(ctx *WhileStmtContext) {}

// ExitWhileStmt is called when production WhileStmt is exited.
func (s *BaseVlangListener) ExitWhileStmt(ctx *WhileStmtContext) {}

// EnterForStmt is called when production ForStmt is entered.
func (s *BaseVlangListener) EnterForStmt(ctx *ForStmtContext) {}

// ExitForStmt is called when production ForStmt is exited.
func (s *BaseVlangListener) ExitForStmt(ctx *ForStmtContext) {}

// EnterNumericRange is called when production NumericRange is entered.
func (s *BaseVlangListener) EnterNumericRange(ctx *NumericRangeContext) {}

// ExitNumericRange is called when production NumericRange is exited.
func (s *BaseVlangListener) ExitNumericRange(ctx *NumericRangeContext) {}

// EnterFuncCall is called when production FuncCall is entered.
func (s *BaseVlangListener) EnterFuncCall(ctx *FuncCallContext) {}

// ExitFuncCall is called when production FuncCall is exited.
func (s *BaseVlangListener) ExitFuncCall(ctx *FuncCallContext) {}

// EnterFuncDecl is called when production FuncDecl is entered.
func (s *BaseVlangListener) EnterFuncDecl(ctx *FuncDeclContext) {}

// ExitFuncDecl is called when production FuncDecl is exited.
func (s *BaseVlangListener) ExitFuncDecl(ctx *FuncDeclContext) {}

// EnterStructDecl is called when production StructDecl is entered.
func (s *BaseVlangListener) EnterStructDecl(ctx *StructDeclContext) {}

// ExitStructDecl is called when production StructDecl is exited.
func (s *BaseVlangListener) ExitStructDecl(ctx *StructDeclContext) {}

// EnterReturnStmt is called when production ReturnStmt is entered.
func (s *BaseVlangListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production ReturnStmt is exited.
func (s *BaseVlangListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterBreakStmt is called when production BreakStmt is entered.
func (s *BaseVlangListener) EnterBreakStmt(ctx *BreakStmtContext) {}

// ExitBreakStmt is called when production BreakStmt is exited.
func (s *BaseVlangListener) ExitBreakStmt(ctx *BreakStmtContext) {}

// EnterContinueStmt is called when production ContinueStmt is entered.
func (s *BaseVlangListener) EnterContinueStmt(ctx *ContinueStmtContext) {}

// ExitContinueStmt is called when production ContinueStmt is exited.
func (s *BaseVlangListener) ExitContinueStmt(ctx *ContinueStmtContext) {}

// EnterDirectAssign is called when production DirectAssign is entered.
func (s *BaseVlangListener) EnterDirectAssign(ctx *DirectAssignContext) {}

// ExitDirectAssign is called when production DirectAssign is exited.
func (s *BaseVlangListener) ExitDirectAssign(ctx *DirectAssignContext) {}

// EnterDeclAssign is called when production DeclAssign is entered.
func (s *BaseVlangListener) EnterDeclAssign(ctx *DeclAssignContext) {}

// ExitDeclAssign is called when production DeclAssign is exited.
func (s *BaseVlangListener) ExitDeclAssign(ctx *DeclAssignContext) {}

// EnterDeclAssignType is called when production DeclAssignType is entered.
func (s *BaseVlangListener) EnterDeclAssignType(ctx *DeclAssignTypeContext) {}

// ExitDeclAssignType is called when production DeclAssignType is exited.
func (s *BaseVlangListener) ExitDeclAssignType(ctx *DeclAssignTypeContext) {}

// EnterIdPattern is called when production IdPattern is entered.
func (s *BaseVlangListener) EnterIdPattern(ctx *IdPatternContext) {}

// ExitIdPattern is called when production IdPattern is exited.
func (s *BaseVlangListener) ExitIdPattern(ctx *IdPatternContext) {}

// EnterCorchetesexpre is called when production corchetesexpre is entered.
func (s *BaseVlangListener) EnterCorchetesexpre(ctx *CorchetesexpreContext) {}

// ExitCorchetesexpre is called when production corchetesexpre is exited.
func (s *BaseVlangListener) ExitCorchetesexpre(ctx *CorchetesexpreContext) {}

// EnterFuncionexpre is called when production funcionexpre is entered.
func (s *BaseVlangListener) EnterFuncionexpre(ctx *FuncionexpreContext) {}

// ExitFuncionexpre is called when production funcionexpre is exited.
func (s *BaseVlangListener) ExitFuncionexpre(ctx *FuncionexpreContext) {}

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

// EnterStructInstancia is called when production StructInstancia is entered.
func (s *BaseVlangListener) EnterStructInstancia(ctx *StructInstanciaContext) {}

// ExitStructInstancia is called when production StructInstancia is exited.
func (s *BaseVlangListener) ExitStructInstancia(ctx *StructInstanciaContext) {}

// EnterStructAttr is called when production StructAttr is entered.
func (s *BaseVlangListener) EnterStructAttr(ctx *StructAttrContext) {}

// ExitStructAttr is called when production StructAttr is exited.
func (s *BaseVlangListener) ExitStructAttr(ctx *StructAttrContext) {}

// EnterIntType is called when production IntType is entered.
func (s *BaseVlangListener) EnterIntType(ctx *IntTypeContext) {}

// ExitIntType is called when production IntType is exited.
func (s *BaseVlangListener) ExitIntType(ctx *IntTypeContext) {}

// EnterFloatType is called when production FloatType is entered.
func (s *BaseVlangListener) EnterFloatType(ctx *FloatTypeContext) {}

// ExitFloatType is called when production FloatType is exited.
func (s *BaseVlangListener) ExitFloatType(ctx *FloatTypeContext) {}

// EnterStringType is called when production StringType is entered.
func (s *BaseVlangListener) EnterStringType(ctx *StringTypeContext) {}

// ExitStringType is called when production StringType is exited.
func (s *BaseVlangListener) ExitStringType(ctx *StringTypeContext) {}

// EnterBoolType is called when production BoolType is entered.
func (s *BaseVlangListener) EnterBoolType(ctx *BoolTypeContext) {}

// ExitBoolType is called when production BoolType is exited.
func (s *BaseVlangListener) ExitBoolType(ctx *BoolTypeContext) {}

// EnterCharType is called when production CharType is entered.
func (s *BaseVlangListener) EnterCharType(ctx *CharTypeContext) {}

// ExitCharType is called when production CharType is exited.
func (s *BaseVlangListener) ExitCharType(ctx *CharTypeContext) {}

// EnterVoidType is called when production VoidType is entered.
func (s *BaseVlangListener) EnterVoidType(ctx *VoidTypeContext) {}

// ExitVoidType is called when production VoidType is exited.
func (s *BaseVlangListener) ExitVoidType(ctx *VoidTypeContext) {}

// EnterStructVector is called when production StructVector is entered.
func (s *BaseVlangListener) EnterStructVector(ctx *StructVectorContext) {}

// ExitStructVector is called when production StructVector is exited.
func (s *BaseVlangListener) ExitStructVector(ctx *StructVectorContext) {}

// EnterIncremento is called when production incremento is entered.
func (s *BaseVlangListener) EnterIncremento(ctx *IncrementoContext) {}

// ExitIncremento is called when production incremento is exited.
func (s *BaseVlangListener) ExitIncremento(ctx *IncrementoContext) {}

// EnterDecremento is called when production decremento is entered.
func (s *BaseVlangListener) EnterDecremento(ctx *DecrementoContext) {}

// ExitDecremento is called when production decremento is exited.
func (s *BaseVlangListener) ExitDecremento(ctx *DecrementoContext) {}
