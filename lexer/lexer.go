package lexer

import(
    "unicode"
)

type Lexer struct {
    input string
    inputLength int
    position int
}

func NewLexer(input string) *Lexer {
    return &Lexer{input:input, inputLength:len(input)}
}

func (l *Lexer) NextToken() Token {

    // Skipping over whitespace
    for l.position < l.inputLength && l.input[l.position] == ' ' {
        l.position++
    }

    // Set end of file
    if l.position >= l.inputLength {
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
        case "=":
	    l.position++
	    return Token{Type:ASSIGN, Value: string(character)}
    }

    if unicode.IsLetter(rune(character)) {
        keyword := ""
        for l.position <= l.inputLength {
	    if l.input[l.position] == ' ' {
	         l.position++
		 break
	    }
	    keyword += string(l.input[l.position])
	    l.position++
	}
	
	if keyword == "const" {
            keywordValue := ""
            for l.position <= l.inputLength {
	        if l.input[l.position] == ' ' {
	            l.position++
		    return Token{Type: IDENTIFER, Value: keywordValue}
	        }
	        keywordValue += string(l.input[l.position])
	        l.position++
	    }
	}
	
	return Token{Type: ILLEGAL, Value: keyword}
    }

    if l.position >= l.inputLength {
        return Token{Type: ENDOFFILE, Value: ""}
    }

    l.position++

    return Token{Type: ILLEGAL, Value: string(character)}
}
