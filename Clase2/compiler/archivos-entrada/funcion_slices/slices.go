func (v *EvalVisitor) VisitSliceDeclInit(ctx *parser.SliceDeclInitContext) interface{} {
	fmt.Println("****************************************")
	fmt.Println("[Debug] VisitSliceDeclInit:")
	fmt.Println("****************************************")
	id := ctx.ID().GetText()
	rawTipo := ctx.Tipos().GetText()     // p.ej. "[]int" o "[][]int"
	dims := strings.Count(rawTipo, "[]") // p.ej. 1 o 2
	baseType := strings.TrimPrefix(rawTipo, strings.Repeat("[]", dims))
	exprList := ctx.ExprList()
	if exprList == nil {
		panic(&SemanticError{"Debe inicializar el slice con lista de valores", ctx.GetStart()})
	}

	fmt.Printf("[Debug:VisitSliceDeclInit] id=%q rawTipo=%q dims=%d baseType=%q\n",
		id, rawTipo, dims, baseType,
	)

	// Si viene MatrixList, lo tratamos como 2D (dims debería ser ≥2)
	if exprList.MatrixList() != nil {
		rowsCtx := exprList.MatrixList().AllExprList()
		fmt.Printf("[Debug] detectada MATRIZ %dD (%d filas)\n", dims, len(rowsCtx))
		rows := make([][]ValueWrapper, 0, len(rowsCtx))
		for i, rowCtx := range rowsCtx {
			elemsCtx := rowCtx.AllExpr()
			fmt.Printf("[Debug] fila %d: %d elementos\n", i, len(elemsCtx))
			row := make([]ValueWrapper, 0, len(elemsCtx))
			for j, eCtx := range elemsCtx {
				val := eCtx.Accept(v).(ValueWrapper)
				fmt.Printf("[Debug]   val[%d]=%v (%T)\n", j, val, val)
				if !v.isTypeAccepted(baseType, val) {
					panic(&SemanticError{
						Message: fmt.Sprintf("Matriz '%s' solo admite %s, pero encontré %T", id, baseType, val),
						Token:   ctx.GetStart(),
					})
				}
				row = append(row, val)
			}
			rows = append(rows, row)
		}
		mat := &MatrixValue{ElementType: baseType, Rows: rows}
		fmt.Printf("[Debug] declarando MATRIZ %s con %d filas\n", id, len(rows))
		v.env.Declare(id, mat, ctx.GetStart())
		v.symbolTable.AddSymbol(id, SymbolSlice, "Slice", ctx.GetStart())
		return mat
	}

	elemsCtx := exprList.AllExpr()
	fmt.Printf("[Debug] rama SLICE 1D (%d elementos)\n", len(elemsCtx))
	elems := make([]ValueWrapper, 0, len(elemsCtx))
	for i, eCtx := range elemsCtx {
		val := eCtx.Accept(v).(ValueWrapper)
		fmt.Printf("[Debug]   elem[%d]=%v (%T)\n", i, val, val)
		if !v.isTypeAccepted(baseType, val) {
			panic(&SemanticError{
				Message: fmt.Sprintf("Slice '%s' solo admite %s, pero encontré %T", id, baseType, val),
				Token:   ctx.GetStart(),
			})
		}
		elems = append(elems, val)
	}
	sl := &SliceValue{ElementType: baseType, Elements: elems}
	v.env.Declare(id, sl, ctx.GetStart())
	v.symbolTable.AddSymbol(id, SymbolSlice, "Slice", ctx.GetStart())
	return sl
}