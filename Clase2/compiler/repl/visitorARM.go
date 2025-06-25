package repl

import (
	parser "compiler/parser"
)

// Visitor personalizado para recorrer el árbol de sintaxis

// Constructor del visitor
type ARMVisitor struct {
	parser.BaseVlangVisitor
	ScopeTrace  *ScopeTrace
	CallStack   *CallStack
	Console     *Console
	StructNames []string
}


// NewARMVisitor 
func NewARMVisitor() *ARMVisitor {
	return &ARMVisitor{
		ScopeTrace:  NewScopeTrace(),
		CallStack:   NewCallStack(),
		Console:     NewConsole(),
		StructNames: []string{},
	}
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