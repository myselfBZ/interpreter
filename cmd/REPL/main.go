package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/myselfBZ/interpreter/internal/lexer"
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
			l.Tokenize()
		}
	}
}

func main() {
	r := New()
	r.Start()
}
