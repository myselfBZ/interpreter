package ast

import (
	"bytes"
	"github.com/myselfBZ/interpreter/internal/token"
)

type PrefixExpression struct {
	Token      *token.Token
	Operator   string
	Expression Expression
}

func (p *PrefixExpression) expressionNode() { return }
func (p *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(p.Operator)
	out.WriteString(p.Expression.String())
	out.WriteString(")")
	return out.String()
}
func (p *PrefixExpression) TokenLiteral() string {
	return p.Token.Literal
}
