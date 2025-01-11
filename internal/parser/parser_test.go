package parser

import (
	"testing"

	"github.com/myselfBZ/interpreter/internal/ast"
	"github.com/myselfBZ/interpreter/internal/lexer"
)

// expect &ast.ReturnStatement{
//     Token: Token{TYPE:"RETURN", Literal:"return"},
//     ReturnValue:Expression{},
// }
func TestReturnStatemetn(t *testing.T) {
    input := `
    return 12;
    `    

    l := lexer.New(input) 
    p := New(l)
    program := p.ParseProgram()
    if program == nil{
        t.Fatalf("ParseProgram() returned nil")
    }
    if len(program.Statements) != 1{
        t.Fatalf("expected 1 got %v", len(program.Statements))
    }
    s, ok := program.Statements[0].(*ast.ReturnStatement) 
    if !ok{
        t.Fatalf("statement isn't of kind ast.ReturnStatement got: %T", s)
    }
    if s.TokenLiteral()  != "return"{
        t.Fatalf("s.Token.Literal is not 'return' got %s", s.TokenLiteral())
    }
}


func TestLetStatemetns(t *testing.T){
    input := `
    let s = 12;
    let b = 21;
    `
    l := lexer.New(input)
    p := New(l)
    program := p.ParseProgram()
    expected := []string{"s", "b"}
    for i, s := range program.Statements{
       if testLetStatement(t, s, expected[i]){
           return
       } 
    }
} 

func testLetStatement(t *testing.T, s ast.Statement, name string) bool{
    if s.TokenLiteral() != "let"{
        t.Errorf("expected  let got %s\n", s.TokenLiteral())
        return false
    }
    v, ok := s.(*ast.LetStatement)
    if !ok{
        t.Errorf("s is not *ast.LetStatement got %T", s)
        return false
    }
    if v.Name.Value != name{
        t.Errorf("expected '%s' got %s\n", name, v.Name.Value)
        return false
    } 
    if v.Name.TokenLiteral() != name{
        t.Errorf("let.name.TokenLiteral() is not %s", name)
        return false
    }
    return true
}
