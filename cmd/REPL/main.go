package main

import (
	"fmt"
	"os"
	"github.com/peterh/liner"
	"github.com/myselfBZ/interpreter/internal/evaluator"
	"github.com/myselfBZ/interpreter/internal/lexer"
	"github.com/myselfBZ/interpreter/internal/parser"
)


func  Start() {
    l := liner.NewLiner() 
    defer l.Close()
    l.SetCtrlCAborts(true)
    his := ".history"
    if f, err := os.Open(his); err == nil{
        l.ReadHistory(f)
        f.Close()
    }

    for{
        input, err := l.Prompt(">>> ")
        if err != nil{
            if err == liner.ErrPromptAborted{
                fmt.Println("byeeee")
                break
            }
            fmt.Println("error: ", err)
            break
        }
        if input == ""{
            continue
        }
        l.AppendHistory(input)
        if input == "exit" || input == "quit" {
            break
        }
        lex := lexer.New(input)
        p := parser.New(lex)
        program := p.ParseProgram()
        if len(p.Errors()) != 0{
            fmt.Println("error: ", p.Errors()[0])
            continue
        }
        e := evaluator.Eval(program)
        fmt.Println(e.Inspect())

    }
}

func main() {
    Start()
}
