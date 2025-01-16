package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/myselfBZ/interpreter/internal/lexer"
	"github.com/myselfBZ/interpreter/internal/parser"
)

func New() *REPL {
	return &REPL{
		scanner: bufio.NewScanner(os.Stdin),
	}
}

type REPL struct {
	scanner *bufio.Scanner
}

func (r *REPL) Start() {
	fmt.Print("Hello welcome to the Monkey programming language\n")
	for {
		fmt.Print(">>>")
		scanned := r.scanner.Scan()
		if !scanned {
			return
		}
		line := r.scanner.Text()
		if line != "" {
			l := lexer.New(line)
            p := parser.New(l)
            program := p.ParseProgram()
            if len(p.Errors()) != 0{
                fmt.Println("Woops, we ran into some monkey buisness here: ")
                fmt.Print("     ")
                for _, e := range p.Errors(){
                    fmt.Println(e)
                }
            } else{
                fmt.Print(program.String())
            }
		}
	}
}

func main() {
	r := New()
	r.Start()
}
