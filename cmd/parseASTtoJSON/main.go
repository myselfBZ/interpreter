package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/myselfBZ/interpreter/internal/lexer"
	"github.com/myselfBZ/interpreter/internal/parser"
	"github.com/tidwall/pretty"
)

func main() {
	src := `
    fn (x, y, t){
        return 23;
    }
    `
	l := lexer.New(src)
	p := parser.New(l)
	program := p.ParseProgram()
	log.Println(program.String())
	programBytes, err := json.Marshal(program)
	if err != nil {
		log.Fatal("error marshaling the program", err)
	}
	prettyProgram := pretty.Pretty(programBytes)
	err = os.WriteFile("cmd/parseASTtoJSON/program.json", prettyProgram, 0666)
	if err != nil {
		log.Fatal("error writing to a file", err)
	}
}
