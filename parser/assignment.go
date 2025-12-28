package parser

import (
	"anchor/lexer"
)

func (p *Parser) parseAssignment() *AssignmentStatement {
	name := &Identifier{Value: p.currentToken().Value}
	p.nextToken()

	if p.currentToken().Type != lexer.ASSIGN {
		panic("expected '='")
	}

	p.nextToken()

	value := p.parseExpression()

	if p.currentToken().Type == lexer.CLOSING {
		p.nextToken()
	}

	return &AssignmentStatement{
		Name:  name,
		Value: value,
	}
}
