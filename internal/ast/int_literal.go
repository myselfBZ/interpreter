package ast

import "github.com/myselfBZ/interpreter/internal/token"

type IntLiteral struct {
	Value int64
	Token *token.Token
}

func (i *IntLiteral) expressionNode() { return }
func (i *IntLiteral) String() string {
	return i.Token.Literal
}
func (i *IntLiteral) TokenLiteral() string {
	return i.Token.Literal
}
