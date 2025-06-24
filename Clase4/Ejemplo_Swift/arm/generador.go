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

public class ArmGenerator
{

    private readonly List<string> instructions = new List<string>();
    private readonly StandardLibrary stdLib = new StandardLibrary();


    public void Mul(string rd, string rs1, string rs2)
    {
        instructions.Add($"MUL {rd}, {rs1}, {rs2}");
    }

    public void Div(string rd, string rs1, string rs2)
    {
        instructions.Add($"DIV {rd}, {rs1}, {rs2}");
    }

    public void Addi(string rd, string rs1, int imm)
    {
        instructions.Add($"ADDI {rd}, {rs1}, #{imm}");
    }





*/

/*
Funcion para terminar el programa
    public void EndProgram()
    {
        Mov(Register.X0, 0);
        Mov(Register.X8, 93); // syscall number for exit
        Svc(); // make syscall
    }
*/

/*
Funcion para sumar registros
public void Add(string rd, string rs1, string rs2)

	{
	    instructions.Add($"ADD {rd}, {rs1}, {rs2}");
	}
*/

/*
// - Memory operations

	public void Str(string rs1, string rs2, int offset = 0)
	{
	    instructions.Add($"STR {rs1}, [{rs2}, #{offset}]");
	}

	public void Ldr(string rd, string rs1, int offset = 0)
	{
	    instructions.Add($"LDR {rd}, [{rs1}, #{offset}]");
	}

	public void Mov(string rd, int imm)
	{
	    instructions.Add($"MOV {rd}, #{imm}");
	}
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
.global _start:

nanananananana



'libreria estandar'
funciones hasta abajo


*/
