package ast

import "github.com/prafitradimas/interpreter/internal/token"

type Identifier struct {
	Token token.Token
	Value string
}

func NewIdent(token token.Token, val string) *Identifier {
	ident := Identifier{
		Token: token,
		Value: val,
	}
	return &ident
}

func (ident *Identifier) expressionNode()      {}
func (ident *Identifier) TokenLiteral() string { return ident.Token.Literal }
func (ident *Identifier) String() string {
	return ident.Value
}
