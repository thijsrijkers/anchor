package generator

import (
	"anchor/parser"
	"fmt"
	"strings"
)

type AssemblyGenerator struct {
	offsets map[string]int
}

func NewAssemblyGenerator() *AssemblyGenerator {
	return &AssemblyGenerator{
		offsets: make(map[string]int),
	}
}

func (gen *AssemblyGenerator) writeLines(writeLine func(string), lines ...string) {
	writeLine(strings.Join(lines, "\n"))
}

func (generator *AssemblyGenerator) generateStatement(statement parser.Statement, writeLine func(string)) {
	switch s := statement.(type) {
	case *parser.AssignmentStatement:
		generator.generateAssignmentToFile(s, writeLine)
	default:
		panic(fmt.Sprintf("unknown statement type: %T", statement))
	}
}

func (generator *AssemblyGenerator) GenerateProgram(program *parser.Program) {
	file, err := createOutputFile("output.s")
	if err != nil {
		return
	}
	defer file.Close()

	writeLine := func(s string) { fmt.Fprintln(file, s) }

	generator.writeLines(writeLine,
		".section .text",
		".global _start",
		"_start:",
		"stp x29, x30, [sp, #-16]!",
		"mov x29, sp",
	)

	for _, stmt := range program.Statements {
		generator.generateStatement(stmt, writeLine)
	}

	generator.writeLines(writeLine,
		"mov x0, #0",
		"mov x8, #93",
		"svc 0",
	)
}
