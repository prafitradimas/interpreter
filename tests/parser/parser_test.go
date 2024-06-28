package parser_test

import (
	"testing"

	"github.com/prafitradimas/interpreter/internal/ast"
	"github.com/prafitradimas/interpreter/internal/lexer"
	"github.com/prafitradimas/interpreter/internal/parser"
	"github.com/prafitradimas/interpreter/internal/token"
)

func TestVarStatement(t *testing.T) {
	inputs := `
	var foo = 5;
	var bar = 10;
	var bazz = 42069;
	`

	lexr := lexer.New(inputs)
	prog := parser.New(lexr)
	program := prog.Parse()

	if program == nil {
		t.Fatal("Parse() returns nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("Invalid `program.Statements` length, expect: %d found: %d", 3, len(program.Statements))
	}

	testData := []ast.VarStatement{
		{
			Token: token.Token{Type: token.VAR, Literal: "var"},
			Name: &ast.Identifier{
				Token: token.Token{Type: token.IDENT, Literal: "foo"},
				Value: "5",
			},
		},
		{
			Token: token.Token{Type: token.VAR, Literal: "var"},
			Name: &ast.Identifier{
				Token: token.Token{Type: token.IDENT, Literal: "bar"},
				Value: "10",
			},
		},
		{
			Token: token.Token{Type: token.VAR, Literal: "var"},
			Name: &ast.Identifier{
				Token: token.Token{Type: token.IDENT, Literal: "bazz"},
				Value: "42069",
			},
		},
	}

	for i, td := range testData {
		statement := program.Statements[i]
		if !expectVarStatement(t, statement, td.Name.Token.Literal, td.Name.Value) {
			return
		}
	}
}

func expectVarStatement(t *testing.T, statement ast.Statement, name string, value string) bool {
	if statement.TokenLiteral() != "var" {
		t.Errorf("Invalid `statement.TokenLiteral()`, expect: %s found: %s", "var", statement.TokenLiteral())
		return false
	}

	varStatement, ok := statement.(*ast.VarStatement)
	if !ok {
		t.Errorf("Invalid Type, expect: `*ast.VarStatement` found: `%T`", varStatement)
		return false
	}

	if varStatement.Name.Token.Type != token.IDENT {
		t.Errorf("Invalid Identifier Token Type, expect: %s found: %s", name, varStatement.Name.Token.Literal)
		return false
	}

	if varStatement.Name.Token.Literal != name {
		t.Errorf("Invalid Identifier Token Literal, expect: %s found: %s", name, varStatement.Name.Token.Literal)
		return false
	}

	if varStatement.Name.Value != value {
		t.Errorf("Invalid Identifier Value, expect: %s found: %s", value, varStatement.Name.Value)
		return false
	}

	return true
}
