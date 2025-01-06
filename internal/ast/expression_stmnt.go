package ast

import "github.com/myselfBZ/interpreter/internal/token"

type ExpressionStatement struct {
	Token      *token.Token
	Expression Expression
}

func (e *ExpressionStatement) statementNode() { return }

func (e *ExpressionStatement) String() string {
	if e.Expression != nil {
		return e.Expression.String()
	}
	return ""
}

func (e *ExpressionStatement) TokenLiteral() string {
	return e.Token.Literal
}
