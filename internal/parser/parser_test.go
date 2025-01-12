package parser

import (
	"fmt"
	"testing"

	"github.com/myselfBZ/interpreter/internal/ast"
	"github.com/myselfBZ/interpreter/internal/lexer"
)

//	expect &ast.ReturnStatement{
//	    Token: Token{TYPE:"RETURN", Literal:"return"},
//	    ReturnValue:Expression{},
//	}
func TestReturnStatemetn(t *testing.T) {
	input := `
    return 12;
    `

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 1 {
		t.Fatalf("expected 1 got %v", len(program.Statements))
	}
	s, ok := program.Statements[0].(*ast.ReturnStatement)
	if !ok {
		t.Fatalf("statement isn't of kind ast.ReturnStatement got: %T", s)
	}
	if s.TokenLiteral() != "return" {
		t.Fatalf("s.Token.Literal is not 'return' got %s", s.TokenLiteral())
	}
}

func TestLetStatemetns(t *testing.T) {
	input := `
    let s = 12;
    let b = 21;
    `
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	expected := []string{"s", "b"}
	for i, s := range program.Statements {
		if testLetStatement(t, s, expected[i]) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("expected  let got %s\n", s.TokenLiteral())
		return false
	}
	v, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s is not *ast.LetStatement got %T", s)
		return false
	}
	if v.Name.Value != name {
		t.Errorf("expected '%s' got %s\n", name, v.Name.Value)
		return false
	}
	if v.Name.TokenLiteral() != name {
		t.Errorf("let.name.TokenLiteral() is not %s", name)
		return false
	}
	return true
}

func TestParseIdent(t *testing.T) {
	input := "foobar;"
	expected := "foobar"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("program is nil")
	}
	if len(program.Statements) != 1 {
		t.Fatalf("statement: expected 1 got %v\n", len(program.Statements))
	}
	s, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("statement is not of kind statement, got %T", program.Statements[0])
	}
	ident, ok := s.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("expected Identifier got %T", s.Expression)
	}
	if ident.Value != "foobar" {
		t.Fatalf("expected %s got %s", expected, ident.Value)
	}
	if ident.TokenLiteral() != expected {
		t.Fatalf("expected %s as a token literal got %s", expected, ident.TokenLiteral())
	}
}

func TestIntLiterals(t *testing.T) {
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
	if len(program.Statements) != 3 {
		t.Fatalf("expected 3 statements got %v", len(program.Statements))
	}
	for i, s := range program.Statements {
		exprsn, ok := s.(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("%T is not an expression statement", s.(*ast.ExpressionStatement))
		}
		if exprsn.TokenLiteral() != expected[i] {
			t.Fatalf("expected %v got %v", expected[i], exprsn.TokenLiteral())
		}
		intLiteral, ok := exprsn.Expression.(*ast.IntLiteral)
		if !ok {
			t.Fatalf("expected IntLitral got %T", exprsn.Expression.(*ast.IntLiteral))
		}
		if intLiteral.Value != int64(expectedInt[i]) {
			t.Fatalf("expected %v got %v", expectedInt[i], intLiteral.Value)
		}
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}
	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
				1, len(program.Statements))
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}
		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Expression)
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s",
				tt.operator, exp.Operator)
		}
		if !testIntegerLiteral(t, exp.Right, tt.integerValue) {
			return
		}
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}
	if integ.Value != value {
		t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
		return false
	}
	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d. got=%s", value,
			integ.TokenLiteral())
		return false
	}
	return true
}

func TestInfix(t *testing.T) {
	input := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"5 + 5;", 5, "+", 5},
		{"4 - 5;", 4, "-", 5},
		{"4 * 5;", 4, "*", 5},
		{"4 / 2;", 4, "/", 2},
		{"4==2;", 4, "==", 2},
		{"4 != 2;", 4, "!=", 2},
	}
	for _, tt := range input {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		if len(program.Statements) != 1 {
			t.Fatalf("wrong number of statements got %v", len(program.Statements))
		}
		exp, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("statement is not an expression statement got %T", program.Statements[0].(*ast.ExpressionStatement))
		}
		infixExps, ok := exp.Expression.(*ast.InfixExperssion)
		if !ok {
			t.Fatalf("expression of expression statement isnt infix expression got=%T", exp.Expression.(*ast.InfixExperssion))
		}
		if !testIntegerLiteral(t, infixExps.Left, tt.leftValue) {
			return
		}
		if !testIntegerLiteral(t, infixExps.Right, tt.rightValue) {
			return
		}
		if infixExps.Operator != tt.operator {
			t.Fatalf("expected %s got=%s", tt.operator, infixExps.Operator)
		}
	}
}

// i know it is stupid
func TestGroupedExpressions(t *testing.T) {
	input := "(123+34) + 12;"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	if len(program.Statements) != 1 {
		t.Fatalf("got more than 1 statements %v", len(program.Statements))
	}
	expressionStatement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("expected expression statement got %T", program.Statements[0].(*ast.ExpressionStatement))
	}
	infixExpression, ok := expressionStatement.Expression.(*ast.InfixExperssion)
	if !ok {
		t.Fatalf("expected infix expression got %T", expressionStatement.Expression.(*ast.InfixExperssion))
	}
	// check the left grouped expression
	leftGroupedExpression, ok := infixExpression.Left.(*ast.InfixExperssion)
	if !ok {
		t.Fatalf("left side is not grouped got: %T ", infixExpression.Left)
	}
	if !testIntegerLiteral(t, leftGroupedExpression.Left, 123) {
		return
	}
	if leftGroupedExpression.Operator != "+" {
		t.Fatalf("expected + got %s", leftGroupedExpression.Operator)
	}
	if !testIntegerLiteral(t, leftGroupedExpression.Right, 34) {
		return
	}
	if infixExpression.Operator != "+" {
		t.Fatalf("expected +  got %s", infixExpression.Operator)
	}
	if !testIntegerLiteral(t, infixExpression.Right, 12) {
		return
	}
}

func TestBoolean(t *testing.T) {
	input := []struct {
		input  string
		value  bool
		litral string
	}{
		{"true;", true, "true"},
		{"false;", false, "false"},
	}
	for _, tt := range input {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		if len(program.Statements) != 1 {
			t.Fatalf("there isn't one statement got %d", len(program.Statements))
		}
		exs, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("not an expression statement got: %T", program.Statements[0].(*ast.ExpressionStatement))
		}
		boolean, ok := exs.Expression.(*ast.Boolean)
		if !ok {
			t.Fatalf("not a boolean value: %T", exs.Expression.(*ast.Boolean))
		}
		if boolean.Value != tt.value {
			t.Fatalf("expected %v got %v", tt.value, boolean.Value)
		}
		if boolean.TokenLiteral() != tt.litral {
			t.Fatalf("falied on litral test expected %s got %s", tt.litral, boolean.TokenLiteral())
		}
	}
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("exp not *ast.Identifier. got=%T", exp)
		return false
	}
	if ident.Value != value {
		t.Errorf("ident.Value not %s. got=%s", value, ident.Value)
		return false
	}
	if ident.TokenLiteral() != value {
		t.Errorf("ident.TokenLiteral not %s. got=%s", value,
			ident.TokenLiteral())
		return false
	}
	return true
}
