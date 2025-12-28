package parser

import (
	"anchor/lexer"
)

type Parser struct {
	tokens []lexer.Token
	pos    int
}

func NewParser(tokens []lexer.Token) *Parser {
	return &Parser{tokens: tokens}
}

func (parser *Parser) currentToken() lexer.Token {
	if parser.pos >= len(parser.tokens) {
		return lexer.Token{Type: lexer.ENDOFFILE}
	}
	return parser.tokens[parser.pos]
}

func (parser *Parser) nextToken() {
	parser.pos++
}

func (p *Parser) parseStatement() Statement {
	token := p.currentToken()

	switch token.Type {
	case lexer.IDENTIFER:
		return p.parseAssignment()
	default:
		p.nextToken()
		return nil
	}
}

func (parser *Parser) ParseProgram() *Program {
	program := &Program{}

	for parser.currentToken().Type != lexer.ENDOFFILE {
		if parser.currentToken().Type == lexer.ILLEGAL {
			parser.nextToken()
			continue
		}

		stmt := parser.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
	}

	return program
}
