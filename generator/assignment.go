package generator

import (
	"anchor/parser"
	"fmt"
)

func (gen *AssemblyGenerator) generateAssignmentToFile(a *parser.AssignmentStatement, writeLine func(string)) {
	writeLine("sub sp, sp, #8")
	writeLine("")
	gen.generateExpressionToFile(a.Value, writeLine)

	offset, ok := gen.offsets[a.Name.Value]
	if !ok {
		gen.offsets[a.Name.Value] = offset
	}

	writeLine(fmt.Sprintf("str x0, [fp,#-%d]", offset))

	writeLine("")
}

func (gen *AssemblyGenerator) generateExpressionToFile(expr parser.Expression, writeLine func(string)) {
	switch e := expr.(type) {
	case *parser.NumberLiteral:
		writeLine(fmt.Sprintf("mov x0, #%s", e.Value))
		writeLine("")
	case *parser.Identifier:
		offset, ok := gen.offsets[e.Value]
		if !ok {
			panic(fmt.Sprintf("variable %s not defined", e.Value))
		}
		writeLine(fmt.Sprintf("ldr x0, [fp,#-%d]", offset))
		writeLine("")
	case *parser.BinaryExpression:
		// Evaluate left
		gen.generateExpressionToFile(e.Left, writeLine)
		writeLine("str x0, [sp, #-16]!")

		// Evaluate right
		gen.generateExpressionToFile(e.Right, writeLine)
		writeLine("ldr x1, [sp], #16")
		writeLine("")

		// Perform operation
		switch e.Operator {
		case "+":
			writeLine("add x0, x1, x0")
		case "-":
			writeLine("sub x0, x1, x0")
		default:
			panic(fmt.Sprintf("unsupported operator: %s", e.Operator))
		}
		writeLine("")
	default:
		panic(fmt.Sprintf("unknown expression type: %T", expr))
	}
}
