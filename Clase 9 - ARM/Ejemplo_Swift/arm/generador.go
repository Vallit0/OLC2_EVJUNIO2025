package arm

import "fmt"

// Objeto que nos permita
// llevar la cuenta de las instrucciones
// este objeto se usa desde otro visitor
type ArmGenerator struct {
	Instructions []string          // Lista de instrucciones generadas
	Variables    map[string]string // Mapa de variables y sus registros
	// Output Console
	Output string // Salida de la consola
	// libreria estandar
	StdLib *StandardLibrary // Biblioteca estándar de funciones
}

// Agregamos las cosas al string y el string es la traduccion

// Agregamos operaciones que o metan a un stack
// o agreguen al string lo que vamos haciendo
//

// funcion para agregar comnetarios la codigo
func (g *ArmGenerator) Comment(comment string) {
	g.Instructions = append(g.Instructions, "## "+comment)
}

/*
Agregar instrucciones
*/
func (g *ArmGenerator) Add(instruction string) {
	g.Instructions = append(g.Instructions, instruction)
}

/*
MUL ->
DIV ->
SUB ->
ADD ->
MOV ->
SVC -> Syscalls
*/

/*


 */

/*

 */

func (g *ArmGenerator) EndProgram() {
	g.StdLib.Use("end_program")
	g.Add("MOV X0, 0")  // Código de salida 0
	g.Add("MOV X8, 93") // Número de syscall para exit
	g.Add("SVC #0")     // Realizar syscall
}

/*
Funcion para sumar registros
*/

/*
// - Memory operations
*/
func Mov(g *ArmGenerator, rd, imm string) {
	g.Add(fmt.Sprintf("MOV %s, #%s", rd, imm))
}

func Ldr(g *ArmGenerator, rd, rs string) {
	g.Add(fmt.Sprintf("LDR %s, [%s]", rd, rs))
}
func Str(g *ArmGenerator, rs, rd string) {
	g.Add(fmt.Sprintf("STR %s, [%s]", rs, rd))
}

/*
push - pop
*/

func Push(g *ArmGenerator, rs string) {
	g.Add(fmt.Sprintf("STR %s, [SP, #-8]!", rs))
}

func Pop(g *ArmGenerator, rd string) {
	g.Add(fmt.Sprintf("LDR %s, [SP], #8", rd))
}

func Sub(g *ArmGenerator, rd, rs1, rs2 string) {
	g.Add(fmt.Sprintf("SUB %s, %s, %s", rd, rs1, rs2))
}

func Add(g *ArmGenerator, rd, rs1, rs2 string) {
	g.Add(fmt.Sprintf("ADD %s, %s, %s", rd, rs1, rs2))
}

func EndProgram(g *ArmGenerator) {
	g.StdLib.Use("end_program")
	g.Add("MOV X0, 0")  // Exit code 0
	g.Add("MOV X8, 93") // Syscall number for exit
	g.Add("SVC #0")     // Make syscall
}

/*
Impresion de las cosas
*/

func PrintInteger(g *ArmGenerator, rs string) {
	g.StdLib.Use("print_integer")
	g.Add(fmt.Sprintf("MOV X0, %s", rs))
	g.Add("BL print_integer")
}

func (g *ArmGenerator) String() {
	// Iniciar el codigo
	result := ".text\n.global _start:\n"

	// Agregar instrucciones
	for _, instruction := range g.Instructions {
		result += instruction + "\n"
	}

	// Agregar la biblioteca estándar
	result += "\n\n\n// Standard Library\n"
	result += NewStandardLibrary().GetFunctionDefinitions()

	// lo agregamos a consola
	g.Output = result
	fmt.Println(result) // Imprimir el resultado en consola
}

/*
Funcion para terminar el programa
syscall EndProgram

*/

func (g *ArmGenerator) GetOutput() string {
	// Retorna el output generado
	return g.Output
}
