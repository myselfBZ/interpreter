package parser

import "github.com/myselfBZ/interpreter/internal/token"

func (p *Parser) nextToken() {
    p.currentToken = p.peekToken
    p.peekToken = p.lexer.NextToken()
}

func (p *Parser) currentTokenIs(t token.TokenType) bool{
    return t == p.currentToken.Type
}

func (p *Parser) peekTokenIs(t token.TokenType) bool{
    return t == p.peekToken.Type
}

func (p *Parser) registerPrefix(t token.TokenType, f parsePrefix){
    p.parsePrefixFn[t] = f
}

func (p *Parser) registerInfix(t token.TokenType, f parseInfix){
    p.parseInfixFn[t] = f
}
