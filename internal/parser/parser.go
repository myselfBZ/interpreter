package parser
// line 97
import (
	"fmt"
	"strconv"

	"github.com/myselfBZ/interpreter/internal/ast"
	"github.com/myselfBZ/interpreter/internal/lexer"
	"github.com/myselfBZ/interpreter/internal/token"
)

const (
    _ int = iota
    LOWEST
    EQUALS
    // ==
    LESSGREATER // > or <
    SUM// +
    PRODUCT// *
    PREFIX// -X or !X
    CALL// myFunction(X)
)

type(
    parsePrefix func() ast.Expression
    parseInfix func(ast.Expression) ast.Expression
)

func New(lexer *lexer.Lexer) *Parser{
    p := &Parser{lexer: lexer}
    p.currentToken = p.lexer.NextToken()
    p.peekToken = p.lexer.NextToken()
    p.registerPrefix(token.IDENT, p.parseIdent)
    return p
}

type Parser struct{
    lexer   *lexer.Lexer
    currentToken *token.Token
    peekToken   *token.Token
    errors []string
    parseInfixFn map[token.TokenType]parseInfix
    parsePrefixFn map[token.TokenType]parsePrefix
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
    for _, n := range program.Statements{
        fmt.Println(n.String())
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
        return nil
    }
}

func (p *Parser) parseInt() ast.Expression {
    number, err := strconv.Atoi(p.currentToken.Literal)
    if err != nil{
        p.errors = append(p.errors, fmt.Sprintf("%s is not a number\n", p.currentToken.Literal))
        return nil
    }
    intNode := &ast.IntLiteral{Token: p.currentToken, Value: int64(number)}
    return intNode
}

func (p *Parser) parseIdent() ast.Expression {
    return &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
}

func (p *Parser) parseExpressionStatements() *ast.ExpressionStatement{
    node := &ast.ExpressionStatement{Token: p.currentToken}
    node.Expression = p.parseExpression(LOWEST)
    return node
}


func (p *Parser) parseExpression(precedence int) ast.Expression {
    prefix := p.parsePrefixFn[p.currentToken.Type]
    if prefix == nil{
        p.errors = append(p.errors, fmt.Sprintf("couldn't find function for %s symbol\n", p.currentToken.Literal))
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

