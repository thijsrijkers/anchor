package parser

import "fmt"

func PrettyPrintTree(p *Program) {
	for _, stmt := range p.Statements {
		printStatementTree(stmt, "")
	}
}

func printStatementTree(stmt Statement, prefix string) {
	switch s := stmt.(type) {
	case *AssignmentStatement:
		fmt.Println(prefix + "AssignmentStatement")
		fmt.Println(prefix + " ├─ Name: " + s.Name.Value)
		fmt.Println(prefix + " └─ Value:")
		printExpressionTree(s.Value, prefix+"    ")
	default:
		fmt.Println(prefix + "UnknownStatement")
	}
}

func printExpressionTree(expr Expression, prefix string) {
	switch e := expr.(type) {
	case *NumberLiteral:
		fmt.Println(prefix + "NumberLiteral: " + e.Value)
	case *Identifier:
		fmt.Println(prefix + "Identifier: " + e.Value)
	case *BinaryExpression:
		fmt.Println(prefix + "BinaryExpression: " + e.Operator)
		fmt.Println(prefix + " ├─ Left:")
		printExpressionTree(e.Left, prefix+" │   ")
		fmt.Println(prefix + " └─ Right:")
		printExpressionTree(e.Right, prefix+"     ")
	default:
		fmt.Println(prefix + "UnknownExpression")
	}
}
