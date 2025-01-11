package ast

import (
	"bytes"

	"github.com/myselfBZ/interpreter/internal/token"
)

type BlockStatement struct {
	Token *token.Token
    Statements []Statement
}

func (b *BlockStatement) String() string {
    var out bytes.Buffer
    for _, v := range b.Statements{
        out.WriteString(v.String())
    }
    return out.String()
}

func (b *BlockStatement) statementNode() {return}

func (b *BlockStatement) TokenLiteral() string{
    return b.Token.Literal
}
