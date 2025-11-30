package lexer

type TokenType string

const (
    ENDOFFILE TokenType = "EOF"
    ILLEGAL TokenType = "ILLEGAL"

    NUMBER TokenType = "NUMBER"
    IDENTIFER TokenType = "IDENT"

    PLUS TokenType = "PLUS"
    MINUS TokenType = "MINUS"
    ASSIGN TokenType = "ASSIGN"
)

type Token struct {
    Type TokenType
    Value string
}
