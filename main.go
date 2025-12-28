package main

import (
	"anchor/fileloader"
	"anchor/lexer"
	"anchor/parser"
	"fmt"
	"log"
	"os"
)

func main() {
	amountOfArguments := len(os.Args)
	if amountOfArguments < 2 {
		log.Fatal("Missing argument for file path")
	}

	inputPath := os.Args[1]
	debug := amountOfArguments >= 3 && os.Args[2] == "--debug"

	content, err := fileloader.Load(inputPath)

	if err != nil {
		log.Fatal(err)
	}

	var tokens []lexer.Token
	l := lexer.NewLexer(string(content))
	for token := l.NextToken(); token.Type != lexer.ENDOFFILE; token = l.NextToken() {
		tokens = append(tokens, token)
		if debug {
			fmt.Printf("%+v\n", token)
		}
	}

	p := parser.NewParser(tokens)
	parsedContent := p.ParseProgram()

	if debug {
		parser.PrettyPrintTree(parsedContent)
	}
}
