package ast

import (
	"bytes"

	"github.com/myselfBZ/interpreter/internal/token"
)

type ReturnStatement struct {
	Token       *token.Token `json:"token"`
	ReturnValue Expression   `json:"return_value"`
}

func (r *ReturnStatement) TokenLiteral() string {
	return r.Token.Literal
}

func (r *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(r.TokenLiteral() + " ")
	if r.ReturnValue != nil {
		out.WriteString(r.ReturnValue.String())
	}
	return out.String()
}

func (r *ReturnStatement) statementNode() {
	return
}
