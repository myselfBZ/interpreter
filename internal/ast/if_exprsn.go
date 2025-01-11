package ast

import (
	"bytes"

	"github.com/myselfBZ/interpreter/internal/token"
)

type IfExpression struct {
	Token      *token.Token
	Condition Expression
    Consequence *BlockStatement
    Alternative *BlockStatement
}

func (i *IfExpression) expressionNode() {return}
func (i *IfExpression) String() string {
    var out bytes.Buffer
    out.WriteString("if")
    out.WriteString(" " + i.Condition.String() + " ")
    out.WriteString(i.Consequence.String())
    if i.Alternative != nil{
        out.WriteString("else")
        out.WriteString(i.Alternative.String())
    }
    return out.String()
}
func (i *IfExpression) TokenLiteral() string {
    return i.Token.Literal
}
