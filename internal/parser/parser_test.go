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


func TestParseIdent(t *testing.T){
    input := "foobar;"
    expected := "foobar"
    l := lexer.New(input)
    p := New(l)
    program := p.ParseProgram()
    if program == nil{
        t.Fatalf("program is nil")
    }
    if len(program.Statements) != 1{
        t.Fatalf("statement: expected 1 got %v\n", len(program.Statements))
    }
    s, ok := program.Statements[0].(*ast.ExpressionStatement) 
    if !ok{
        t.Fatalf("statement is not of kind statement, got %T", program.Statements[0])
    }
    ident, ok := s.Expression.(*ast.Identifier)
    if !ok{
        t.Fatalf("expected Identifier got %T", s.Expression)
    }
    if ident.Value != "foobar"{
        t.Fatalf("expected %s got %s", expected, ident.Value)
    }
    if ident.TokenLiteral() != expected{
        t.Fatalf("expected %s as a token literal got %s", expected, ident.TokenLiteral())
    }
}


func TestIntLiterals(t *testing.T){
    input := `
    123;
    13;
    1;
    `
    expected := []string{"123", "13", "1"}
    expectedInt := []int{123, 13, 1}
    l := lexer.New(input)
    p := New(l)
    program := p.ParseProgram()
    if len(program.Statements) != 3{
        t.Fatalf("expected 3 statements got %v", len(program.Statements))
    } 
    for i, s := range program.Statements{
        exprsn, ok := s.(*ast.ExpressionStatement)
        if !ok{
            t.Fatalf("%T is not an expression statement", s.(*ast.ExpressionStatement))
        }
        if exprsn.TokenLiteral() != expected[i]{
            t.Fatalf("expected %v got %v", expected[i], exprsn.TokenLiteral())
        }
        intLiteral, ok := exprsn.Expression.(*ast.IntLiteral) 
        if !ok{
            t.Fatalf("expected IntLitral got %T",exprsn.Expression.(*ast.IntLiteral))
        }
        if intLiteral.Value != int64(expectedInt[i]){
            t.Fatalf("expected %v got %v", expectedInt[i], intLiteral.Value)
        }
    }
}
