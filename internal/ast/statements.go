package ast

import "github.com/prafitradimas/interpreter/internal/token"

type VarStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (vs *VarStatement) statementNode() {}
func (vs *VarStatement) TokenLiteral() string {
	return vs.Token.Literal
}
