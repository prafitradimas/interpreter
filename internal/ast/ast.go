package ast

// interfaces
type (
	Node interface {
		TokenLiteral() string
	}

	Statement interface {
		Node
		statementNode()
	}

	Expression interface {
		Node
		expressionNode()
	}
)

// impl
type Program struct {
	Statements []Statement
}

func (root *Program) TokenLiteral() string {
	if len(root.Statements) > 0 {
		return root.Statements[0].TokenLiteral()
	}
	return ""
}
