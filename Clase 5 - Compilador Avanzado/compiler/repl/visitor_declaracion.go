package repl

import (
	"log"

	parser "compiler/parser"
	"compiler/value"

	"github.com/antlr4-go/antlr/v4"
)

type DclVisitor struct {
	parser.BaseVlangVisitor
	ScopeTrace  *ScopeTrace
	ErrorTable  *ErrorTable
	StructNames []string
}

func NewDclVisitor(errorTable *ErrorTable) *DclVisitor {
	return &DclVisitor{
		ScopeTrace:  NewScopeTrace(),
		ErrorTable:  errorTable,
		StructNames: []string{},
	}
}

func (v *DclVisitor) Visit(tree antlr.ParseTree) interface{} {

	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(val.GetText())
		return nil
	default:
		return tree.Accept(v)
	}

}

func (v *DclVisitor) VisitPrograma(ctx *parser.ProgramaContext) interface{} {

	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}

	return nil
}

func (v *DclVisitor) VisitStmt(ctx *parser.StmtContext) interface{} {

	if ctx.Func_dcl() != nil {
		v.Visit(ctx.Func_dcl())
	} else if ctx.Struct_dcl() != nil {
		v.Visit(ctx.Struct_dcl())
	}

	return nil
}

func (v *DclVisitor) VisitFuncDecl(ctx *parser.FuncDeclContext) interface{} {

	// Entorno -> Global
	// Entorno -> Actual
	if v.ScopeTrace.CurrentScope != v.ScopeTrace.GlobalScope {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Las funciones solo pueden ser declaradas en el scope global")
	}

	funcName := ctx.ID().GetText()

	args := make([]*Param, 0)

	if ctx.Arg_list() != nil {
		args = v.Visit(ctx.Arg_list()).([]*Param)
	}

	if len(args) > 0 {

		// ParamType() -> IVOR_NIL | IVOR_INT | IVOR_FLOAT | IVOR_STRING | IVOR_BOOL
		baseParamType := args[0].ParamType()

		for _, param := range args {
			if param.ParamType() != baseParamType {
				v.ErrorTable.NewSemanticError(param.Token, "Todos los parametros de la funcion deben ser del mismo tipo")
				return nil
			}
		}

		// void | int | float | string | bool | nil
		// fn sumar(a int, b int) int
		returnType := value.IVOR_NIL
		var returnTypeToken antlr.Token = nil

		// // TODO: Como se resuelve el tipo
		// if ctx.Type_() != nil {
		// 	returnType = ctx.Type_().GetText()
		// 	returnTypeToken = ctx.Type_().GetStart()
		// }

		body := ctx.AllStmt()

		function := &Function{ // pointer ?
			Name:            funcName,
			Param:           args,
			ReturnType:      returnType,
			Body:            body,
			DeclScope:       v.ScopeTrace.CurrentScope,
			ReturnTypeToken: returnTypeToken,
			Token:           ctx.GetStart(),
		}

		ok, msg := v.ScopeTrace.AddFunction(funcName, function)

		if !ok {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
		}

		return nil
	}
	return nil
}
func (v *DclVisitor) VisitArgList(ctx *parser.ArgListContext) interface{} {

	args := make([]*Param, 0)

	for _, arg := range ctx.AllFunc_arg() {
		// TODO: IVOR -> Param -> Arg
		args = append(args, v.Visit(arg).(*Param))
	}

	return args
}

// func (v *DclVisitor) VisitFuncParam(ctx *parser.FuncParamContext) interface{} {

// 	externName := ""
// 	innerName := ""

// 	// at least ID(0) is defined
// 	// only 1 ID defined
// 	if ctx.ID(1) == nil {
// 		// innerName : type
// 		// _ : type
// 		innerName = ctx.ID(0).GetText()
// 	} else {
// 		// externName innerName : type
// 		externName = ctx.ID(0).GetText()
// 		innerName = ctx.ID(1).GetText()
// 	}

// 	passByReference := false

// 	if ctx.INOUT_KW() != nil {
// 		passByReference = true
// 	}

// 	paramType := ctx.Type_().GetText()

// 	return &Param{
// 		ExternName:      externName,
// 		InnerName:       innerName,
// 		PassByReference: passByReference,
// 		Type:            paramType,
// 		Token:           ctx.GetStart(),
// 	}

// }

// func (v *DclVisitor) VisitStructDecl(ctx *parser.StructDeclContext) interface{} {

// 	v.StructNames = append(v.StructNames, ctx.ID().GetText())
// 	// Agregar los nombres de las variables
// 	// Recorrer los hijos del struct y agregar las variables al scope actual
// 	if v.ScopeTrace.CurrentScope != v.ScopeTrace.GlobalScope {
// 		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Las estructuras solo pueden ser declaradas en el scope global")
// 		return nil
// 	}

// 	structName := ctx.ID().GetText()

// 	fields := make([]*Variable, 0)
// 	for _, field := range ctx.AllStructField() {

// 		variable := v.Visit(field).(*Variable)

// 	return nil
// }

//return nil
//}
