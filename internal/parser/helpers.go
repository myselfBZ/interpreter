package parser

import (
	"fmt"

	"github.com/myselfBZ/interpreter/internal/token"
)

func (p *Parser) expectPeekToken(t token.TokenType) bool {
	if p.peekToken.Type == t {
		p.nextToken()
		return true
	}
	return false
}

func (p *Parser) registerPrefix(t token.TokenType, f parsePrefix) {
	p.prefixFns[t] = f
}

func (p *Parser) registerInfix(t token.TokenType, f parseInfix) {
	p.infixFns[t] = f
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) currentTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) noPrefixExpression(t token.TokenType) {
	p.errors = append(p.errors, fmt.Sprintf("no prefix func for %s\n", t))
}

func (p *Parser) curPrecedence() int{
    if p, ok := precedences[p.curToken.Type]; ok{
        return p
    }
    return LOWEST
}

func (p *Parser) peekPrecedence() int{
    if p, ok := precedences[p.peekToken.Type]; ok{
        return p
    }
    return LOWEST
}
