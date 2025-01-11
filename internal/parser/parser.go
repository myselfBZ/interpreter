package parser

import (
	"github.com/myselfBZ/interpreter/internal/ast"
	"github.com/myselfBZ/interpreter/internal/lexer"
	"github.com/myselfBZ/interpreter/internal/token"
)

// let's rewrite the parser
// where we left of: parseExoressionStatement method 
type(
    parsePrefix func() ast.Expression
    parseInfix func(ast.Expression) ast.Expression
)


type Parser struct{
    lexer *lexer.Lexer
    curToken *token.Token
    errors []string
    peekToken *token.Token
    prefixFns map[token.TokenType]parsePrefix
    infixFns map[token.TokenType]parseInfix
}

func New(l *lexer.Lexer) *Parser{
    p := &Parser{
        lexer: l,
        curToken: l.NextToken(),
        peekToken: l.NextToken(),
        prefixFns: make(map[token.TokenType]parsePrefix),
        infixFns: make(map[token.TokenType]parseInfix),
    }
    p.registerPrefix(token.IDENT, p.parseIdent)
    return p
}

func (p *Parser) ParseProgram() *ast.Program{
    program := &ast.Program{}
    for !p.currentTokenIs(token.EOF){
        stmnt := p.parseStatement(p.curToken)
        if stmnt != nil{
            program.Statements = append(program.Statements, stmnt)
        }
        p.nextToken()
    }
    return program
}

func (p *Parser) nextToken() {
    p.curToken = p.peekToken
    p.peekToken = p.lexer.NextToken()
}

func (p *Parser) parseStatement(t *token.Token) ast.Statement{
    switch t.Type{
    case token.LET:
        return p.parseLet()
    case token.RETURN:
        return p.parseReturnStatement()
    default:
        return nil
    }
}


// func (p *Parser) parseExpressionStatement() ast.Statement{
//     node := &ast.ExpressionStatement{Token: p.curToken}
//
// }

func (p *Parser) parseLet() ast.Statement{
    node := &ast.LetStatement{Token: p.curToken}
    if !p.expectPeekToken(token.IDENT){
        return nil
    }
    node.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
    if !p.expectPeekToken(token.ASSIGN){
        return nil
    }
    p.nextToken()
    //TODO: parse the expression
    for !p.currentTokenIs(token.SEMICOLON){
        p.nextToken()
    }
    return node
}

func (p *Parser) parseReturnStatement() ast.Statement{
    node := &ast.ReturnStatement{Token: p.curToken};    
    // TODO: parse the expression
    for !p.currentTokenIs(token.SEMICOLON){
        p.nextToken()
    }
    return node
}

func (p *Parser) parseIdent() ast.Expression{ 
    node := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
    return node
}
