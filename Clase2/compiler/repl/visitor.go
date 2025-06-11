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

// Constructor del visitor
type ReplVisitor struct {
	parser.BaseVlangVisitor
	ScopeTrace  *ScopeTrace
	CallStack   *CallStack
	Console     *Console
	ErrorTable  *ErrorTable
	StructNames []string
}

/*
NOTA: para testear estoy usando un constructor
para ReplVisitor que solo recibe la tabla de errores
y luego inicializar el ScopeTrace y el CallStack
y la Consola

PERO: el metodo NewVisitor(dclVisitor *DclVisitor)
lo que hace es que recibe un DclVisitor que hace una pasada
antes para ver que cosas estan declaradas.

ENTONCES: NewVisitor es el que inicializa el ReplVisitor
y le pasa el ScopeTrace, ErrorTable y StructNames, estructuras
que ya tienen inicializadas las listas con lo que se agrego.

*/

// Recibe DclVisitor e inicializa el ReplVisitor
func NewVisitor(dclVisitor *DclVisitor) *ReplVisitor {
	return &ReplVisitor{
		ScopeTrace:  dclVisitor.ScopeTrace,
		ErrorTable:  dclVisitor.ErrorTable,
		StructNames: dclVisitor.StructNames,
		CallStack:   NewCallStack(),
		Console:     NewConsole(),
	}
}

// tambien podemos crear un ReplVisitor sin DclVisitor
// Esto es util para hacer tests de una vez
func NewReplVisitor(errorTable *ErrorTable) *ReplVisitor {
	return &ReplVisitor{
		ScopeTrace:  NewScopeTrace(),
		ErrorTable:  errorTable,
		StructNames: []string{},
		CallStack:   NewCallStack(),
		Console:     NewConsole(),
	}
}

// GetReplContext devuelve el contexto del REPL
// que contiene la consola, el ScopeTrace, el CallStack y la ErrorTable
func (v *ReplVisitor) GetReplContext() *ReplContext {
	return &ReplContext{
		Console:    v.Console,
		ScopeTrace: v.ScopeTrace,
		CallStack:  v.CallStack,
		ErrorTable: v.ErrorTable,
	}
}

// ValidType verifica si el tipo es valido en el ScopeTrace
// Esto se usa para validar tipos de variables, funciones, etc.
func (v *ReplVisitor) ValidType(_type string) bool {
	return v.ScopeTrace.GlobalScope.ValidType(_type)
}

// Visit es el metodo que se llama para visitar un nodo del arbol de sintaxis
// Dependiendo del tipo de nodo, se llama al metodo correspondiente
// Si es un ErrorNodeImpl, se lanza un error
func (v *ReplVisitor) Visit(tree antlr.ParseTree) interface{} {
	if tree == nil {
		fmt.Println("⚠️ Árbol nulo recibido.")
		return nil
	}
	fmt.Println("Enrutado------------------")
	switch node := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(node.GetText())

	case *parser.ProgramaContext:
		return v.VisitPrograma(node)

	case *parser.StmtContext:
		return v.VisitStmt(node)

	case *parser.PrintlnStmtContext:
		return v.VisitPrintlnStmt(node)

	case *parser.PrintStmtContext:
		return v.VisitPrintStmt(node)
	case *parser.ValorexpresionContext:
		// Visitamos el hijo de valorexpresion
		return v.VisitValorexpresion(node)

	// tiramos todos los case de los nodos literales
	case *parser.ValorEnteroContext:
		fmt.Println("🔍 Visitando ValorEntero:", node.GetText())
		return v.VisitValorEntero(node)
	case *parser.ValorFloatContext:
		return v.VisitValorFloat(node)
	case *parser.ValorDecimalContext:
		return v.VisitValorDecimal(node)
	case *parser.ValorCadenaContext:
		return v.VisitValorCadena(node)
	case *parser.ValorBooleanoContext:
		return v.VisitValorBooleano(node)
	case *parser.ValorCaracterContext:
		return v.VisitValorCaracter(node)
	case *parser.IdContext:
		return v.VisitId(node)
	// Expresiones binarias
	case *parser.BinaryExpContext:
		return v.VisitBinaryExp(node)
	/// Agregamos las declaraciones
	case *parser.DeclAssignContext:
		return v.VisitValueDeclAssign(node)
	case *parser.DirectAssignContext:
		return v.VisitDirectAssign(node)

	default:
		fmt.Printf("⚠️ Tipo inesperado en Visit(): %T\n", tree)
		return tree.Accept(v) // fallback por si acaso
	}

	return nil
}

// VisitPrograma es el metodo que se llama para visitar el nodo Programa
// Este nodo es el nodo raiz del arbol de sintaxis
// En este metodo recorremos todos los statements del programa
func (v *ReplVisitor) VisitPrograma(ctx *parser.ProgramaContext) interface{} {

	fmt.Println("🔍 Visitando el programa...")
	// Vamos a recorrer todos los statements y los visitamos
	for _, stmt := range ctx.AllStmt() {
		fmt.Println("🔍 Visitando statement*:", stmt.GetText())
		v.Visit(stmt)
	}

	return nil
}

/*

Ahora recorremos todos los stmts que tenemos en el programa

*/

func (v *ReplVisitor) VisitStmt(ctx *parser.StmtContext) interface{} {
	fmt.Println("🔍 Visitando statement*:", ctx.GetText())

	if ctx == nil || ctx.GetChildCount() == 0 {
		fmt.Println("⚠️ Stmt vacío o nulo.")
		return nil
	}

	node := ctx.GetChild(0)

	// Aquí van todos los ifs como los que hicimos antes:
	if printlnStmt, ok := node.(*parser.PrintlnStmtContext); ok {
		fmt.Println("🔔 Visitando nodo println")

		v.VisitPrintlnStmt(printlnStmt)
	} else {
		fmt.Printf("⚠️ Tipo no reconocido dentro de stmt: %T\n", node)
	}

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

		// // verificamos si es un struct
		// if v.ScopeTrace.CurrentScope.isStruct {
		// 	canMutate = v.ScopeTrace.IsMutatingEnvironment()
		// }

		ok, msg := variable.Assign(varValue, true)

		if !ok {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
		}
	}

	return nil

}

/*

Con este metodo subimos el valor de la expresion

*/

func (v *ReplVisitor) VisitValorexpresion(ctx *parser.ValorexpresionContext) interface{} {
	// solo vamos a subir al hijo
	fmt.Println("🔍 Visitando Valorexpresion:", ctx.GetText())

	/*
		Todo esta bien, pero, tenemos que verificar el tipo de objeto que
		es el ctx y mandarlo al visit de nuevo para que suba pero
		con una forma de expresion con valores atomicos

		volvemos a subir el valor a la expresion

	*/
	valor := ctx.Valor()

	return v.Visit(valor)
}

// print normal
func (v *ReplVisitor) VisitPrintStmt(ctx *parser.PrintStmtContext) interface{} {
	fmt.Println("🔍 Visitando PrintStmt:", ctx.GetText())

	// Total de expresiones
	exprs := ctx.AllExpresion()

	// Lista de resultados para imprimir
	resultados := []string{}

	for _, expr := range exprs {
		result := v.Visit(expr)
		if result == nil {
			v.ErrorTable.NewSemanticError(expr.GetStart(), "Expresión vacía dentro de print")
			continue
		}

		val, ok := result.(value.IVOR)
		if !ok {
			v.ErrorTable.NewSemanticError(expr.GetStart(), "La expresión no devuelve un valor válido")
			continue
		}

		// Convertir el valor a string y agregar a resultados
		text := fmt.Sprintf("%v", val.Value())
		resultados = append(resultados, text)
	}

	// Unir con espacios y enviar a la consola
	finalOutput := strings.Join(resultados, " ")
	v.Console.Println(finalOutput)

	return nil
}

func (v *ReplVisitor) VisitPrintlnStmt(ctx *parser.PrintlnStmtContext) interface{} {
	fmt.Println("🔍 Visitando PrintlnStmt:", ctx.GetText())

	// Evaluar expresión
	result := v.Visit(ctx.Expresion())
	if result == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Expresión vacía dentro de Println")
		return nil
	}

	// Verificar que el resultado sea IVOR
	val, ok := result.(value.IVOR)
	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "La expresión no devuelve un valor válido")
		return nil
	}

	// Obtener valor interno e imprimirlo como texto
	valInterno := val.Value()
	valTipo := val.Type()

	fmt.Printf("🔔 Imprimiendo valor: %v (tipo: %s)\n", valInterno, valTipo)

	// Agregar a consola como string
	v.Console.Println(fmt.Sprintf("%v", valInterno))
	return nil
}

func (v *ReplVisitor) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {

	runChain := true

	for _, ifStmt := range ctx.AllIf_chain() {

		runChain = !v.Visit(ifStmt).(bool)
		// Si la condicion del if es verdadera, ejecutamos el bloque
		if !runChain {
			break
		}
	}

	if runChain && ctx.Else_stmt() != nil {
		v.Visit(ctx.Else_stmt())
	}

	return nil
}

func (v *ReplVisitor) VisitIfChain(ctx *parser.IfChainContext) interface{} {

	condition := v.Visit(ctx.Expresion()).(value.IVOR)

	if condition.Type() != value.IVOR_BOOL {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "La condicion del if debe ser un booleano")
		return false

	}

	if condition.(*value.BoolValue).InternalValue {

		// Push scope
		v.ScopeTrace.PushScope("if")

		for _, stmt := range ctx.AllStmt() {
			v.Visit(stmt)
		}

		// Pop scope
		v.ScopeTrace.PopScope()

		return true
	}

	return false
}

func (v *ReplVisitor) VisitElseStmt(ctx *parser.ElseStmtContext) interface{} {

	// Push scope {   }
	v.ScopeTrace.PushScope("else")

	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}

	// Pop scope
	v.ScopeTrace.PopScope()

	return nil
}

func (v *ReplVisitor) VisitValueDeclAssign(ctx *parser.DeclAssignContext) interface{} {

	/*
		Cuando declaramos una variable tenemos que

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
	/*

	   fn main() {
	   // entorno padre

	   x = 6
	   while (true){
	    x = 5
	   }
	   while (true){
	   x = 6
	   }
	   }

	*/
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

// el hijo de valorexpresioncontext -> valorEnter | valorBooleano |

func (v *ReplVisitor) VisitValorEntero(ctx *parser.ValorEnteroContext) interface{} {

	fmt.Println("🔍 Visitando ValorEntero:", ctx.GetText())

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
