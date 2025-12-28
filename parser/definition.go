package parser

type Node interface {
    node()
}

type Statement interface {
    Node
    statement()
}

type Expression interface {
    Node
    expression()
}

/* Program */

type Program struct {
    Statements []Statement
}

func (p *Program) node() {}

/* Statements */

type AssignmentStatement struct {
    Name  *Identifier
    Value Expression
}

func (a *AssignmentStatement) node()      {}
func (a *AssignmentStatement) statement() {}

/* Expressions */

type Identifier struct {
    Value string
}

func (i *Identifier) node()       {}
func (i *Identifier) expression() {}

type NumberLiteral struct {
    Value string
}

func (n *NumberLiteral) node()       {}
func (n *NumberLiteral) expression() {}

type BinaryExpression struct {
    Left     Expression
    Operator string
    Right    Expression
}

func (b *BinaryExpression) node()       {}
func (b *BinaryExpression) expression() {}
