package parser

import (
	"fmt"

	"github.com/myselfBZ/interpreter/internal/ast"
	"github.com/myselfBZ/interpreter/internal/lexer"
	"github.com/myselfBZ/interpreter/internal/token"
)

func New(lexer *lexer.Lexer) *Parser{
    p := &Parser{lexer: lexer}
    p.currentToken = p.lexer.NextToken()
    p.peekToken = p.lexer.NextToken()
    return p
}

type Parser struct{
    lexer   *lexer.Lexer
    currentToken *token.Token
    peekToken   *token.Token
    errors []string
}

func (p *Parser) ParseProgram() *ast.Program {
    program := &ast.Program{}
    for p.currentToken.Type != token.EOF{
        stmnt := p.parseStatement(p.currentToken)
        if stmnt != nil{
            program.Statements = append(program.Statements, stmnt)
        }
        p.nextToken()
    }
    if len(p.errors) > 0{
        for _, v := range p.errors {
            fmt.Println("error: ", v)
        }
    }
    return program
}

func (p *Parser) parseStatement(t *token.Token) ast.Statement{
    switch t.Type{
    case token.LET:
        return p.parseLet()
    case token.RETURN:
        return p.parseReturn()
    default:
        fmt.Printf("%s is currently not supported\n", p.currentToken.Type)
        return nil
    }
}

func (p *Parser) parseLet() *ast.LetStatement{
    letNode := &ast.LetStatement{Token: p.currentToken}
    if p.peekToken.Type != token.IDENT{
        p.errors = append(p.errors, fmt.Sprintf("%s cannot be a name for a variable\n", p.peekToken.Literal))
        return nil
    } 

    
    letNode.Name = &ast.Identifier{Token:p.peekToken, Value: p.peekToken.Literal}
    p.nextToken()
    
    if p.peekToken.Type != token.ASSIGN{
        p.errors = append(p.errors, fmt.Sprintf("expected '=' got '%s'", p.peekToken.Literal))
        return nil
    }

    p.nextToken()
    
    //TODO parse the expression, we are skipping it for now
    for p.currentToken.Type != token.SEMICOLON{
        p.nextToken()
    }

    return letNode
}

func (p *Parser) parseReturn() *ast.ReturnStatement {
    returnNode := &ast.ReturnStatement{Token: p.currentToken}
    // skip the expression
    for p.currentToken.Type != token.SEMICOLON{
        p.nextToken()
    } 
    return returnNode
}

func (p *Parser) nextToken() {
    p.currentToken = p.peekToken
    p.peekToken = p.lexer.NextToken()
}
