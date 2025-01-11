package ast

import (
	"bytes"

	"github.com/myselfBZ/interpreter/internal/token"
)

type InfixExperssion struct {
	Left     Expression
    Operator string
	Right    Expression
	Token    *token.Token
}

func (i *InfixExperssion) expressionNode() { return }
func (i *InfixExperssion) String() string{
    var out bytes.Buffer
    out.WriteString(i.Left.String())
    out.WriteString(i.Operator)
    out.WriteString(i.Right.String())
    return out.String()
}
func (i *InfixExperssion) TokenLiteral() string { return i.Token.Literal }
