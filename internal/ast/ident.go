package ast

import (
	"github.com/myselfBZ/interpreter/internal/token"
)

type Identifier struct {
	Token *token.Token `json:"token"`
	Value string       `json:"value"`
}

func (i *Identifier) expressionNode() { return }

func (i *Identifier) String() string {
	return i.Value
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
