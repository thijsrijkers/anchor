package lexer

import(
    "unicode"
)

type Lexer struct {
    input string
    position int
}

func NewLexer(input string) *Lexer {
    return &Lexer{input:input}
}

func (l *Lexer) NextToken() Token {
    // Skipping over whitespace
    for l.position < len(l.input) && l.input[l.position] == ' ' {
        l.position++
    }

    // Set end of file
    if l.position >= len(l.input) {
        return Token{Type: ENDOFFILE, Value: ""}
    }

    character := l.input[l.position]

    if unicode.IsDigit(rune(character)) {
        l.position++
        return Token{Type: NUMBER, Value: string(character)}
    }

    switch string(character) {
        case "+":
	    l.position++
	    return Token{Type: PLUS, Value: string(character)}
        case "-":
	    l.position++
	    return Token{Type:MINUS, Value: string(character)}
    }


    l.position++
    return Token{Type: ILLEGAL, Value: string(character)}
}
