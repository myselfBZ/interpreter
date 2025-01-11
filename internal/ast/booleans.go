package ast

import "github.com/myselfBZ/interpreter/internal/token"

type Boolean struct{
    Token *token.Token
    Value bool
}

func (b *Boolean) expressionNode() {return}
func (b *Boolean) String() string{
    return b.Token.Literal
}
func (b *Boolean) TokenLiteral() string{
    return b.Token.Literal
}
