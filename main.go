package main

import (
	"anchor/fileloader"
	"anchor/generator"
	"anchor/lexer"
	"anchor/parser"
	"anchor/runner"
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
	debug := false
	execute := false

	for _, arg := range os.Args[2:] {
		switch arg {
		case "--debug":
			debug = true
		case "--execute":
			execute = true
		}
	}
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
	program := p.ParseProgram()

	if debug {
		parser.PrettyPrintTree(program)
	}

	g := generator.NewAssemblyGenerator()
	g.GenerateProgram(program)

	if execute {
		runner.AssembleAndRun(debug)
	}
}
