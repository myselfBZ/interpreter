package ast

import (
	"bytes"

	"github.com/myselfBZ/interpreter/internal/token"
)

var (
	ErrorVariableIDENT = `"%s" cannot be a variable name it is a %s`
)

// LetStatement godoc
// @Param:Token for storing LET token
type LetStatement struct {
	Token *token.Token `json:"token"`
	Value Expression   `json:"value"`
	Name  *Identifier  `json:"name"`
}

func (l *LetStatement) TokenLiteral() string {
	return l.Token.Literal
}

func (l *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(l.TokenLiteral() + " ")
	out.WriteString(l.Name.String() + "=")
	if l.Value != nil {
		out.WriteString(l.Value.TokenLiteral())
	}
	return out.String()
}

func (l *LetStatement) statementNode() {
	return
}
