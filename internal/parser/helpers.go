package parser

import "github.com/myselfBZ/interpreter/internal/token"

func (p *Parser) expectPeekToken(t token.TokenType) bool {
    if p.peekToken.Type == t{
        p.nextToken() 
        return true
    }
    return false
}

func (p *Parser) registerPrefix(t token.TokenType, f parsePrefix){
    p.prefixFns[t] = f
}

func (p *Parser) registerInfix(t token.TokenType, f parseInfix){
    p.infixFns[t] = f
}

func (p *Parser) peekTokenIs(t token.TokenType) bool{
    return p.peekToken.Type == t
}

func (p *Parser) currentTokenIs(t token.TokenType) bool{
    return p.curToken.Type == t
}

