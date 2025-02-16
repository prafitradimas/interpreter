package parser_test

import (
	"testing"

	"github.com/prafitradimas/interpreter/internal/ast"
	"github.com/prafitradimas/interpreter/internal/lexer"
	"github.com/prafitradimas/interpreter/internal/parser"
	"github.com/prafitradimas/interpreter/internal/token"
)

func TestVarStatements(t *testing.T) {
	inputs := `
	var foo = 5;
	var bar = 10;
	var bazz = 42069;
	`

	lexr := lexer.New(inputs)
	parser := parser.New(lexr)

	program := parser.Parse()
	checkParserErrors(t, parser)

	if program == nil {
		t.Fatal("Parse() returns nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	testData := []struct {
		expectedIdentifier string
	}{{"foo"}, {"bar"}, {"bazz"}}

	for i, td := range testData {
		statement := program.Statements[i]
		if !expectVarStatement(t, statement, td.expectedIdentifier, td.expectedIdentifier) {
			return
		}
	}
}

func TestReturnStatements(t *testing.T) {
	inputs := `
	return 5;
	return 10;
	return 42069;
	`

	lexr := lexer.New(inputs)
	ps := parser.New(lexr)

	program := ps.Parse()
	checkParserErrors(t, ps)

	if program == nil {
		t.Fatal("Parse() returns nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		retStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("Invalid Statement, expect `*ast.ReturnStatement`. got=%T", stmt)
			continue
		}

		if retStmt.TokenLiteral() != "return" {
			t.Errorf("Invalid Statement, expect `*ast.ReturnStatement`. got=%s", retStmt.TokenLiteral())
			continue
		}
	}
}

func TestString(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.VarStatement{
				Token: token.Token{Type: token.VAR, Literal: "var"},
				Name: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "foo"},
					Value: "foo",
				},
				Value: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "bar"},
					Value: "bar",
				},
			},
		},
	}

	if program.String() != "var foo = bar;" {
		t.Errorf("program.String() wrong. got=%s", program.String())
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

	if varStatement.Name.TokenLiteral() != name {
		t.Errorf("Invalid Identifier Token Literal, expect: %s found: %s", name, varStatement.Name.Token.Literal)
		return false
	}

	if varStatement.Name.TokenLiteral() != name {
		t.Errorf("Invalid Identifier Value, expect: %s found: %s", value, varStatement.Name.Value)
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, parser *parser.Parser) {
	errs := parser.Errors()
	if len(errs) == 0 {
		return
	}

	t.Errorf("Found %d errors while parsing\n", len(errs))
	for _, msg := range errs {
		t.Errorf("parser error: %q\n", msg)
	}
	t.FailNow()
}
