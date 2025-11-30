package main

import(
    "log"
    "os"
    "fmt"
    "anchor/fileloader"
    "anchor/lexer"
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

    l := lexer.NewLexer(string(content))
    for token := l.NextToken(); token.Type != lexer.ENDOFFILE; token = l.NextToken() {
        fmt.Printf("%+v\n", token)
    }
}
