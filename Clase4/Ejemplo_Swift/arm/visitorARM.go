// Aqui tiramos el visitor de ARM
package arm

import (
	"fmt"
	"main/compiler"
)

type ArmVisitor struct {
	compiler.BaseTSwiftLanguageVisitor
	Generator *ArmGenerator // Generador de instrucciones ARM
}

// visitor que nos ayuda a traducr
// constructor de ArmVisitor
func NewArmVisitor() *ArmVisitor {
	return &ArmVisitor{
		Generator: &ArmGenerator{
			Instructions: []string{},
			Variables:    make(map[string]string),
			Output:       "",
		},
	}
}

// visitProgram para recorrer los statements
/*
public override Object? VisitProgram(LanguageParser.ProgramContext context)
    {
        foreach (var dcl in context.dcl())
        {
            Visit(dcl);
        }
        return null;
    }
*/

func (v *ArmVisitor) VisitProgram(ctx *compiler.ProgramContext) interface{} {
	for _, dcl := range ctx.AllStmt() {
		v.Visit(dcl)
	}
	return nil
}

// Traduccion cuando se visite una literal de tipo int
func (v *ArmVisitor) VisitIntLiteral(ctx *compiler.IntLiteralContext) interface{} {
	// funcion para agregar comentarios a nuestro codigo de assembler
	// tomamos el valor del INT
	value := ctx.GetText()
	fmt.Println("Visiting IntLiteral:", value)
	v.Generator.Comment("Literal agregada: " + value)
	// agregamos a una variable el valor
	v.Generator.Add("MOV X0, #" + value) // MOV X0, #valor
	v.Generator.Comment("Valor movido a X0: " + value)

	// Aqui es donde ya tenemos que agregar las traducciones
	return nil
}

/*
public override Object? VisitPrintStmt(LanguageParser.PrintStmtContext context)
    {
        c.Comment("Print statement");
        c.Comment("Visiting expression");
        Visit(context.expr());
        c.Comment("Popping value to print");
        c.Pop(Register.X0); // Pop the value to print
        c.PrintInteger(Register.X0); // Call the print function

        return null;
    }

*/

// --> llego al print ("aqui")

// std:library
// std::printf()

/*

//--------------------------------------------------------------
// print_integer - Prints a signed integer to stdout
//
// Input:
//   x0 - The integer value to print
//--------------------------------------------------------------
print_integer:
    // Save registers
    stp x29, x30, [sp, #-16]!  // Save frame pointer and link register
    stp x19, x20, [sp, #-16]!  // Save callee-saved registers
    stp x21, x22, [sp, #-16]!
    stp x23, x24, [sp, #-16]!
    stp x25, x26, [sp, #-16]!
    stp x27, x28, [sp, #-16]!

    // Check if number is negative
    mov x19, x0                // Save original number
    cmp x19, #0                // Compare with zero
    bge positive_number        // Branch if greater or equal to zero

    // Handle negative number
    mov x0, #1                 // fd = 1 (stdout)
    adr x1, minus_sign         // Address of minus sign
    mov x2, #1                 // Length = 1
    mov w8, #64                // Syscall write
    svc #0

    neg x19, x19               // Make number positive

positive_number:
    // Prepare buffer for converting result to ASCII
    sub sp, sp, #32            // Reserve space on stack
    mov x22, sp                // x22 points to buffer

    // Initialize digit counter
    mov x23, #0                // Digit counter

    // Handle special case for zero
    cmp x19, #0
    bne convert_loop

    // If number is zero, just write '0'
    mov w24, #48               // ASCII '0'
    strb w24, [x22, x23]       // Store in buffer
    add x23, x23, #1           // Increment counter
    b print_result             // Skip conversion loop

convert_loop:
    // Divide the number by 10
    mov x24, #10
    udiv x25, x19, x24         // x25 = x19 / 10 (quotient)
    msub x26, x25, x24, x19    // x26 = x19 - (x25 * 10) (remainder)

    // Convert remainder to ASCII and store in buffer
    add x26, x26, #48          // Convert to ASCII ('0' = 48)
    strb w26, [x22, x23]       // Store digit in buffer
    add x23, x23, #1           // Increment digit counter

    // Prepare for next iteration
    mov x19, x25               // Quotient becomes the new number
    cbnz x19, convert_loop     // If number is not zero, continue

    // Reverse the buffer since digits are in reverse order
    mov x27, #0                // Start index
reverse_loop:
    sub x28, x23, x27          // x28 = length - current index
    sub x28, x28, #1           // x28 = length - current index - 1

    cmp x27, x28               // Compare indices
    bge print_result           // If crossed, finish reversing

    // Swap characters
    ldrb w24, [x22, x27]       // Load character from start
    ldrb w25, [x22, x28]       // Load character from end
    strb w25, [x22, x27]       // Store end character at start
    strb w24, [x22, x28]       // Store start character at end

    add x27, x27, #1           // Increment start index
    b reverse_loop             // Continue reversing

print_result:
    // Add newline
    mov w24, #10               // Newline character
    strb w24, [x22, x23]       // Add to end of buffer
    add x23, x23, #1           // Increment counter

    // Print the result
    mov x0, #1                 // fd = 1 (stdout)
    mov x1, x22                // Buffer address
    mov x2, x23                // Buffer length
    mov w8, #64                // Syscall write
    svc #0

    // Clean up and restore registers
    add sp, sp, #32            // Free buffer space
    ldp x27, x28, [sp], #16    // Restore callee-saved registers
    ldp x25, x26, [sp], #16
    ldp x23, x24, [sp], #16
    ldp x21, x22, [sp], #16
    ldp x19, x20, [sp], #16
    ldp x29, x30, [sp], #16    // Restore frame pointer and link register
    ret                        // Return to caller

minus_sign:
    .ascii ""-""               // Minus sign"
*/
