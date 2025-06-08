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
	// ScopeTrace  *ScopeTrace
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

	for _, stmt := range ctx.AllStmt() {
		fmt.Println("🔍 Visitando declaración:", stmt.GetText())
		v.Visit(stmt)
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
