package parser

import (
	"anchor/lexer"
	"fmt"
)

func (parser *Parser) parseExpression() Expression {
	left := parser.parsePrimary()

	for parser.currentToken().Type == lexer.PLUS || parser.currentToken().Type == lexer.MINUS {
		operator := parser.currentToken().Value
		parser.nextToken()
		right := parser.parsePrimary()

		left = &BinaryExpression{
			Left:     left,
			Operator: operator,
			Right:    right,
		}
	}

	return left
}

func (parser *Parser) parsePrimary() Expression {
	token := parser.currentToken()
	parser.nextToken()

	switch token.Type {
	case lexer.IDENTIFER:
		return &Identifier{Value: token.Value}
	case lexer.NUMBER:
		return &NumberLiteral{Value: token.Value}
	default:
		panic(fmt.Sprintf("unexpected token: %v", token))
	}
}
