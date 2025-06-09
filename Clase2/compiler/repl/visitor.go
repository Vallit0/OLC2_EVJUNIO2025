package repl

import (
	parser "compiler/parser"
	"compiler/value"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// Visitor personalizado para recorrer el árbol de sintaxis
type ReplVisitor struct {
	parser.BaseVlangVisitor
	ScopeTrace *ScopeTrace
	// CallStack   *CallStack
	Console     *Console
	ErrorTable  *ErrorTable
	StructNames []string
}

// Constructor del visitor
func NewReplVisitor() *ReplVisitor {
	return &ReplVisitor{}
}
func (v *ReplVisitor) Visit(tree antlr.ParseTree) interface{} {

	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(val.GetText())
		return nil
	default:
		return tree.Accept(v)
	}

}

func (v *ReplVisitor) VisitPrograma(ctx *parser.ProgramaContext) interface{} {

	// Vamos a recorrer todos los statements y los visitamos
	for _, stmt := range ctx.AllStmt() {
		fmt.Println("🔍 Visitando declaración:", stmt.GetText())
		v.Visit(stmt)
	}

	return nil
}

/*

Ahora recorremos todos los stmts que tenemos en el programa

*/

func (v *ReplVisitor) VisitStmt(ctx *parser.StmtContext) interface{} {
	// vamos a recorrer todos los statements que tenemos
	if ctx != nil && ctx.GetChildCount() != 0 {
		// es distinto de nil y no tiene hijos
		if ctx.Decl_stmt() != nil {
			// Mandamos a visitar la declaracion
			v.Visit(ctx.Decl_stmt())
		} else if ctx.Assign_stmt() != nil {
			// mandamos a asignar la variable que tenemos
			v.Visit(ctx.Assign_stmt()) // -> me voy a ir a hacer el visitor
		}
		// despues vemos si es otro statement
		// otro statement que sea de otro tipo
	}

	// Si es un error, lo lanzamos
	//ctx.

	// Vemos el caso de la Declaracion de Cualquier cosa

	// v.Visit(ctx.declarationStatement())
	// } else if ctx.Assign_stmt() != nil {
	// 	v.Visit(ctx.Assign_stmt())
	// }
	//
	// else if ctx.If_stmt() != nil {
	// 	v.Visit(ctx.If_stmt())
	// } else if ctx.Switch_stmt() != nil {
	// 	v.Visit(ctx.Switch_stmt())
	// } else if ctx.While_stmt() != nil {
	// 	v.Visit(ctx.While_stmt())
	// } else if ctx.For_stmt() != nil {
	// 	v.Visit(ctx.For_stmt())
	// } else if ctx.Guard_stmt() != nil {
	// 	v.Visit(ctx.Guard_stmt())
	// } else if ctx.Transfer_stmt() != nil {
	// 	v.Visit(ctx.Transfer_stmt())
	// } else if ctx.Func_call() != nil {
	// 	v.Visit(ctx.Func_call())
	// } else if ctx.Func_dcl() != nil {
	// 	v.Visit(ctx.Func_dcl())
	// } else if ctx.Strct_dcl() != nil {
	// 	v.Visit(ctx.Strct_dcl())
	// } else if ctx.Vector_func() != nil {
	// 	v.Visit(ctx.Vector_func())
	// } else {
	// 	log.Fatal("Statement not found " + ctx.GetText())
	// }

	return nil
}

func (v *ReplVisitor) VisitDirectAssign(ctx *parser.DirectAssignContext) interface{} {
	// Vamos a visitar la asignacion
	// Primero vamos a ver si es una asignacion de variable
	// lo que queremos hacer para una asignacion
	// es ver si tenemos el ID
	varName := v.Visit(ctx.Id_pattern()).(string)
	varValue := v.Visit(ctx.Expresion()).(value.IVOR)
	// le pedimos la variable al ScopeTrace
	variable := v.ScopeTrace.GetVariable(varName)

	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Variable "+varName+" no encontrada")
	} else {

		// Se copia el objeto

		// Aqui se deberia agregar la validacion del vector

		canMutate := true

		// // verificamos si es un struct
		// if v.ScopeTrace.CurrentScope.isStruct {
		// 	canMutate = v.ScopeTrace.IsMutatingEnvironment()
		// }

		ok, msg := variable.Assign(varValue, canMutate)

		if !ok {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
		}
	}

	return nil

}

func (v *ReplVisitor) VisitValueDeclAssign(ctx *parser.DeclAssignContext) interface{} {

	/*
		Cuando declaramos una variable tenemos que
		1. Saber si es constante o variable (let o var)
		2. Saber el nombre de la variable
		3. Saber el valor de la variable (Depende, si solo la declaramos)
		mut

	*/
	// verificamos si es una constante

	// isConst := isDeclConst(ctx.Var_type().GetText())
	varName := ctx.ID().GetText()
	varValue := v.Visit(ctx.Expresion()).(value.IVOR)
	varType := varValue.Type()

	if varType == "[]" {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "No se puede inferir el tipo de un vector vacio '"+varName+"'")
		return nil
	}

	// copy object
	// if obj, ok := varValue.(*ObjectValue); ok {
	// 	varValue = obj.Copy()
	// }

	// if IsVectorType(varValue.Type()) {
	// 	varValue = varValue.Copy()
	// }

	variable, msg := v.ScopeTrace.AddVariable(varName, varType, varValue, false, false, ctx.GetStart())

	// Variable already exists
	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
	}
	return nil
}

/*


Devolvemos el valor de la expresion en este caso es el tope,
es un entero asi que solamente tomaremos lo que nos da el usuario
y lo convertiremos a un entero

*/

func (v *ReplVisitor) VisitValorEntero(ctx *parser.ValorEnteroContext) interface{} {

	intVal, _ := strconv.Atoi(ctx.GetText())

	return &value.IntValue{
		InternalValue: intVal,
	}

}

// Manejar Floats
func (v *ReplVisitor) VisitValorFloat(ctx *parser.ValorFloatContext) interface{} {

	floatVal, _ := strconv.ParseFloat(ctx.GetText(), 64)

	return &value.FloatValue{
		InternalValue: floatVal,
	}

}
func (v *ReplVisitor) VisitIdPattern(ctx *parser.IdContext) interface{} {
	// hay que ir a buscar el Id a el entorno
	return ctx.GetText()
}

func (v *ReplVisitor) VisitCorchetesexpre(ctx *parser.CorchetesexpreContext) interface{} {

	return v.Visit(ctx.Expresion())
}

// Suma Re

func (v *ReplVisitor) VisitId(ctx *parser.IdContext) interface{} {
	id := ctx.ID().GetText()
	// TODO: buscar variable en entorno
	fmt.Println("Accediendo a variable:", id)
	return nil
}

func (v *ReplVisitor) VisitIncredecr(ctx *parser.IncredecrContext) interface{} {
	return v.Visit(ctx.Incredecre())
}

func (v *ReplVisitor) VisitExpdotexp1(ctx *parser.Expdotexp1Context) interface{} {
	left := ctx.ID(0).GetText()
	right := ctx.ID(1).GetText()
	// TODO: acceder a campo estructurado: ID.ID
	fmt.Println("Accediendo a campo estructurado:", left, ".", right)
	return nil
}

func (v *ReplVisitor) VisitExpdotexp(ctx *parser.ExpdotexpContext) interface{} {
	id := ctx.ID().GetText()
	expr := v.Visit(ctx.Expresion())
	// TODO: ID.EXPR, posiblemente array o campo dinámico
	fmt.Println("Accediendo a campo dinámico:", id, "con expresión:", expr)
	return nil
}

func (v *ReplVisitor) VisitAsignacionfor(ctx *parser.AsignacionforContext) interface{} {
	id := ctx.ID().GetText()
	expr := v.Visit(ctx.Expresion())
	fmt.Println("Asignando valor a variable:", id, "con expresión:", expr)
	// TODO: asignar expr a id
	return nil
}

func (v *ReplVisitor) VisitValorDecimal(ctx *parser.ValorDecimalContext) interface{} {
	text := ctx.DECIMAL().GetText()
	val, err := strconv.ParseFloat(text, 64)
	if err != nil {
		// Manejo de error
		// mandamos el error a ErrorTable
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Expresión vacía dentro de paréntesis")
		log.Printf("Error al convertir '%s' a decimal: %v", text, err)
		return nil
	}
	return &value.FloatValue{
		InternalValue: val,
	}
}

func (v *ReplVisitor) VisitValorCadena(ctx *parser.ValorCadenaContext) interface{} {
	stringVal := ctx.GetText()[1 : len(ctx.GetText())-1]

	stringVal = strings.ReplaceAll(stringVal, "\\\"", "\"")
	stringVal = strings.ReplaceAll(stringVal, "\\\\", "\\")
	stringVal = strings.ReplaceAll(stringVal, "\\n", "\n")
	stringVal = strings.ReplaceAll(stringVal, "\\r", "\r")

	// si es un caracter -> solamente un char
	if len(stringVal) == 1 {
		return &value.CharacterValue{
			InternalValue: stringVal,
		}
	}

	return &value.StringValue{
		InternalValue: stringVal,
	}
}

func (v *ReplVisitor) VisitValorBooleano(ctx *parser.ValorBooleanoContext) interface{} {
	text := ctx.BOOLEANO().GetText()
	// vemos si es true o false y devolvemos eso
	// si es true entonces devolvemos un valor booleano
	if text == "true" {
		return &value.BoolValue{
			InternalValue: true,
		}
	}
	if text == "false" {
		return &value.BoolValue{
			InternalValue: false,
		}
	}
	return nil
}

func (v *ReplVisitor) VisitValorCaracter(ctx *parser.ValorCaracterContext) interface{} {
	text := ctx.CARACTER().GetText()
	// Suponiendo que vienen entre comillas simples: 'a'
	if len(text) >= 3 {
		char := text[1 : len(text)-1]
		return &value.CharacterValue{
			InternalValue: char,
		}
	}
	return nil
}

/*
Ahora haremos la seccion de las expresiones binarias completa
en esta seccion haremos uso de un strategy, el cual
explicare en la clase del dia LUNES
*/
func (v *ReplVisitor) VisitBinaryExp(ctx *parser.BinaryExpContext) interface{} {

	op := ctx.GetOp().GetText()
	left := v.Visit(ctx.GetLeft()).(value.IVOR)
	// Verificamos si hay una estrategia de retorno anticipado para este operador
	earlyCheck, ok := EarlyReturnStrats[op]

	if ok {
		ok, _, result := earlyCheck.Validate(left)

		if ok {
			return result
		}
	}

	right := v.Visit(ctx.GetRight()).(value.IVOR)

	/*
			Aqui es donde sucede la magia, buscamos el operador
			y haciendo uso de un strategy pattern, validamos
			y ejecutamos la operacion de manera dinamica
		Si el operador no existe, lanzamos un error
		Si el operador existe, validamos los tipos de datos
	*/
	strat, ok := BinaryStrats[op]

	if !ok {
		log.Fatal("Binary operator not found")
	}

	ok, msg, result := strat.Validate(left, right)

	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetOp(), msg)
		return value.DefaultNilValue
	}

	return result
}
