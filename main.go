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
	if len(os.Args) < 2 {
		log.Fatal("Missing argument for file path")
	}

	inputPath := os.Args[1]
	content, err := fileloader.Load(inputPath)

	if err != nil {
		log.Fatal(err)
	}

	var tokens []lexer.Token
	l := lexer.NewLexer(string(content))
	for token := l.NextToken(); token.Type != lexer.ENDOFFILE; token = l.NextToken() {
		tokens = append(tokens, token)
		fmt.Printf("%+v\n", token)
	}

	p := parser.NewParser(tokens)
	parsedContent := p.ParseProgram()
	parser.PrettyPrintTree(parsedContent)
}
