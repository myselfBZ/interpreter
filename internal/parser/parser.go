package parser

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/myselfBZ/interpreter/internal/ast"
	"github.com/myselfBZ/interpreter/internal/lexer"
	"github.com/myselfBZ/interpreter/internal/token"
)

const (
	_ int = iota
	LOWEST
	LESSGREATER
	EQUALS
	SUM
	PRODUCT
	PREFIX
	CALL
)

var precedences = map[token.TokenType]int{
	token.EQ:     EQUALS,
	token.NOT_EQ: EQUALS,
	token.LT:     LESSGREATER,
	token.GT:     LESSGREATER,
	token.PLUS:   SUM,
	// reason: it musn't mix the prefix with subtraction
	token.MINUS:          SUM,
	token.DIVISION:       PRODUCT,
	token.MULTIPLICATION: PRODUCT,
}

type (
	parsePrefixFn func() ast.Expression
	parseInfixFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	lexer        *lexer.Lexer
	peekToken    *token.Token
	currentToken *token.Token
	errors       []error
	infixFns     map[token.TokenType]parseInfixFn
	prefixFns    map[token.TokenType]parsePrefixFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer: l, currentToken: l.NextToken(),
		peekToken: l.NextToken(),
		infixFns:  make(map[token.TokenType]parseInfixFn),
		prefixFns: make(map[token.TokenType]parsePrefixFn),
	}
	p.prefixFns[token.IDENT] = p.parseIdent
	p.prefixFns[token.INT] = p.parseInt
	p.prefixFns[token.MINUS] = p.parsePrefixExpression
	p.prefixFns[token.BANG] = p.parsePrefixExpression
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.DIVISION, p.parseInfixExpression)
	p.registerInfix(token.MULTIPLICATION, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)
	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

// parsing methods beg

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLet()
	case token.RETURN:
		return p.parseReturn()
	default:
		return p.parseExpressionStmnt()
	}
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	for p.currentToken.Type != token.EOF {
		stmnt := p.parseStatement()
		if stmnt != nil {
			program.Statements = append(program.Statements, stmnt)
			p.nextToken()
		}
	}
	if len(p.errors) != 0 {
		for _, err := range p.errors {
			log.Println(err)
		}
		return nil
	}
	fmt.Println("Program: ", program.String())
	return program
}

func (p *Parser) parseInt() ast.Expression {
	number, err := strconv.Atoi(p.currentToken.Literal)
	if err != nil {
		// error message is just crazy, i know. it is what it is dont get pissed off at me okay?
		p.errors = append(p.errors, errors.New(fmt.Sprintf("%s is not a number", p.currentToken.Literal)))
		return nil
	}
	intNode := &ast.IntLiteral{Value: int64(number), Token: p.currentToken}
	return intNode
}

func (p *Parser) parseExpressionStmnt() ast.Statement {
	stmnt := &ast.ExpressionStatement{Token: p.currentToken}
	stmnt.Expression = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmnt
}

func (p *Parser) parseIdent() ast.Expression {
	return &ast.Identifier{Value: p.currentToken.Literal, Token: p.currentToken}
}

func (p *Parser) parseReturn() ast.Statement {
	returnNode := &ast.ReturnStatement{}
	returnNode.Token = p.currentToken
	p.nextToken()
	returnNode.ReturnValue = p.parseExpression(LOWEST)
	return returnNode
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	pref := p.prefixFns[p.currentToken.Type]
    if pref == nil{
        log.Println("exression func not found for this: ", p.currentToken.Literal)
        return nil
    }
    left := pref()
    for p.curPrecedence() > precedence && p.currentTokenIs(token.SEMICOLON){
        infix := p.infixFns[p.currentToken.Type]
        if infix == nil{
            return left
        }
        p.nextToken()
        left = infix(left)
    }
	// left expression
	return left
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	prefNode := &ast.PrefixExpression{Token: p.currentToken, Operator: p.currentToken.Literal}
	p.nextToken()
	prefNode.Expression = p.parseExpression(PREFIX)
	return prefNode
}


func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
    infExp := &ast.InfixExperssion{
        Token: p.currentToken,
        Left: left,
        Operator: p.currentToken.Literal,
    }
    p.nextToken()
    right := p.parseExpression(p.curPrecedence())
    infExp.Right = right
    return infExp
}

func (p *Parser) parseLet() ast.Statement {
	letNode := &ast.LetStatement{}
	letNode.Token = p.currentToken
	if !p.expectPeek(token.IDENT) {
		err := p.cannotBeIdent(p.peekToken)
		p.errors = append(p.errors, err)
		p.nextToken()
	}
	letNode.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

	if !p.peekTokenIs(token.ASSIGN) {
		err := fmt.Sprintf(errorExpectedToken, token.ASSIGN, p.currentToken.Literal, p.currentToken.Type)
		p.errors = append(p.errors, errors.New(err))
		p.nextToken()
	}
	// skip the assing
	p.nextToken()
	p.nextToken()
	letNode.Value = p.parseExpression(LOWEST)
	p.nextToken()
	return letNode
}

// parsing methods end

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.currentToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) currentTokenIs(t token.TokenType) bool {
	return p.currentToken.Type == t
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return t == p.peekToken.Type
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	return false
}
func (p *Parser) cannotBeIdent(t *token.Token) error {
	var err string
	if _, ok := token.Keywords[t.Literal]; ok {
		err = fmt.Sprintf(ast.ErrorVariableIDENT, t.Literal, "reserved keyword")
		return errors.New(err)
	}
	switch t.Type {
	case token.INT:
		err = fmt.Sprintf(ast.ErrorVariableIDENT, t.Literal, "number")
	case token.ASSIGN, token.BANG, token.MINUS, token.MULTIPLICATION, token.LBRACE, token.SEMICOLON:
		err = fmt.Sprintf(ast.ErrorVariableIDENT, t.Literal, "reserved symbol")
	}
	return errors.New(err)
}

func (p *Parser) registerInfix(t token.TokenType, fn parseInfixFn) {
	p.infixFns[t] = fn
}

func (p *Parser) registerPrefix(t token.Token, fn parsePrefixFn) {
	p.prefixFns[t.Type] = fn
}
