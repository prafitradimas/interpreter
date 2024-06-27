package token_test

import (
	"fmt"
	"testing"

	"github.com/prafitradimas/interpreter/internal/assert"
	"github.com/prafitradimas/interpreter/internal/lexer"
	"github.com/prafitradimas/interpreter/internal/token"
)

func TestNextToken(t *testing.T) {
	inputs := "=+(){}"

	expectedTokens := []token.Token{
		{
			Type:    token.ASSIGN,
			Literal: "=",
		},
		{
			Type:    token.PLUS,
			Literal: "+",
		},
		{
			Type:    token.LPAREN,
			Literal: "(",
		},
		{
			Type:    token.RPAREN,
			Literal: ")",
		},
		{
			Type:    token.LBRACE,
			Literal: "{",
		},
		{
			Type:    token.RBRACE,
			Literal: "}",
		},
	}

	lexr := lexer.New(inputs)
	for idx, _token := range expectedTokens {
		nextToken := lexr.NextToken()

		if _token.Type != nextToken.Type {
			t.Fatalf("Tests[%d] - Unexpected type.\nExpected: %q, Got: %q", idx, _token.Type, nextToken.Type)
		}

		if _token.Literal != nextToken.Literal {
			t.Fatalf("Tests[%d] - Unexpected type.\nExpected: %q, Got: %q", idx, _token.Literal, nextToken.Literal)
		}
	}
}

func TestNextTokenExt(t *testing.T) {
	inputs := `
	var five = 5;
	var ten = 10;

	fn add(x, y) {
		return x + y;
	};

	var result = add(five, ten);

	if (result == 0) {};
	if (result >= 0) {};
	if (result <= 0) {};
	if (result > 0) {};
	if (result < 0) {};
	`

	expectedTokens := []token.Token{
		{
			Type:    token.VAR,
			Literal: "var",
		},
		{
			Type:    token.IDENT,
			Literal: "five",
		},
		{
			Type:    token.ASSIGN,
			Literal: "=",
		},
		{
			Type:    token.INT,
			Literal: "5",
		},
		{
			Type:    token.SEMICOLON,
			Literal: ";",
		},
		{
			Type:    token.VAR,
			Literal: "var",
		},
		{
			Type:    token.IDENT,
			Literal: "ten",
		},
		{
			Type:    token.ASSIGN,
			Literal: "=",
		},
		{
			Type:    token.INT,
			Literal: "10",
		},
		{
			Type:    token.SEMICOLON,
			Literal: ";",
		},
		{
			Type:    token.FUNCTION,
			Literal: "fn",
		},
		{
			Type:    token.IDENT,
			Literal: "add",
		},
		{
			Type:    token.LPAREN,
			Literal: "(",
		},
		{
			Type:    token.IDENT,
			Literal: "x",
		},
		{
			Type:    token.COMMA,
			Literal: ",",
		},
		{
			Type:    token.IDENT,
			Literal: "y",
		},
		{
			Type:    token.RPAREN,
			Literal: ")",
		},
		{
			Type:    token.LBRACE,
			Literal: "{",
		},
		{
			Type:    token.RETURN,
			Literal: "return",
		},
		{
			Type:    token.IDENT,
			Literal: "x",
		},
		{
			Type:    token.PLUS,
			Literal: "+",
		},
		{
			Type:    token.IDENT,
			Literal: "y",
		},
		{
			Type:    token.SEMICOLON,
			Literal: ";",
		},
		{
			Type:    token.RBRACE,
			Literal: "}",
		},
		{
			Type:    token.SEMICOLON,
			Literal: ";",
		},
		{
			Type:    token.VAR,
			Literal: "var",
		},
		{
			Type:    token.IDENT,
			Literal: "result",
		},
		{
			Type:    token.ASSIGN,
			Literal: "=",
		},
		{
			Type:    token.IDENT,
			Literal: "add",
		},
		{
			Type:    token.LPAREN,
			Literal: "(",
		},
		{
			Type:    token.IDENT,
			Literal: "five",
		},
		{
			Type:    token.COMMA,
			Literal: ",",
		},
		{
			Type:    token.IDENT,
			Literal: "ten",
		},
		{
			Type:    token.RPAREN,
			Literal: ")",
		},
		{
			Type:    token.SEMICOLON,
			Literal: ";",
		},
		{
			Type:    token.IF,
			Literal: "if",
		},
		{
			Type:    token.LPAREN,
			Literal: "(",
		},
		{
			Type:    token.IDENT,
			Literal: "result",
		},
		{
			Type:    token.EQ,
			Literal: "==",
		},
		{
			Type:    token.INT,
			Literal: "0",
		},
		{
			Type:    token.RPAREN,
			Literal: ")",
		},
		{
			Type:    token.LBRACE,
			Literal: "{",
		},
		{
			Type:    token.RBRACE,
			Literal: "}",
		},
		{
			Type:    token.SEMICOLON,
			Literal: ";",
		},
		{
			Type:    token.IF,
			Literal: "if",
		},
		{
			Type:    token.LPAREN,
			Literal: "(",
		},
		{
			Type:    token.IDENT,
			Literal: "result",
		},
		{
			Type:    token.GTE,
			Literal: ">=",
		},
		{
			Type:    token.INT,
			Literal: "0",
		},
		{
			Type:    token.RPAREN,
			Literal: ")",
		},
		{
			Type:    token.LBRACE,
			Literal: "{",
		},
		{
			Type:    token.RBRACE,
			Literal: "}",
		},
		{
			Type:    token.SEMICOLON,
			Literal: ";",
		},
		{
			Type:    token.IF,
			Literal: "if",
		},
		{
			Type:    token.LPAREN,
			Literal: "(",
		},
		{
			Type:    token.IDENT,
			Literal: "result",
		},
		{
			Type:    token.LTE,
			Literal: "<=",
		},
		{
			Type:    token.INT,
			Literal: "0",
		},
		{
			Type:    token.RPAREN,
			Literal: ")",
		},
		{
			Type:    token.LBRACE,
			Literal: "{",
		},
		{
			Type:    token.RBRACE,
			Literal: "}",
		},
		{
			Type:    token.SEMICOLON,
			Literal: ";",
		},
		{
			Type:    token.IF,
			Literal: "if",
		},
		{
			Type:    token.LPAREN,
			Literal: "(",
		},
		{
			Type:    token.IDENT,
			Literal: "result",
		},
		{
			Type:    token.GT,
			Literal: ">",
		},
		{
			Type:    token.INT,
			Literal: "0",
		},
		{
			Type:    token.RPAREN,
			Literal: ")",
		},
		{
			Type:    token.LBRACE,
			Literal: "{",
		},
		{
			Type:    token.RBRACE,
			Literal: "}",
		},
		{
			Type:    token.SEMICOLON,
			Literal: ";",
		},
		{
			Type:    token.IF,
			Literal: "if",
		},
		{
			Type:    token.LPAREN,
			Literal: "(",
		},
		{
			Type:    token.IDENT,
			Literal: "result",
		},
		{
			Type:    token.LT,
			Literal: "<",
		},
		{
			Type:    token.INT,
			Literal: "0",
		},
		{
			Type:    token.RPAREN,
			Literal: ")",
		},
		{
			Type:    token.LBRACE,
			Literal: "{",
		},
		{
			Type:    token.RBRACE,
			Literal: "}",
		},
		{
			Type:    token.SEMICOLON,
			Literal: ";",
		},
		{
			Type:    token.EOF,
			Literal: "",
		},
	}

	lexr := lexer.New(inputs)
	for idx, _token := range expectedTokens {
		nextToken := lexr.NextToken()

		assert.AddTokenData(nextToken)

		assert.InvalidToken(
			_token.Type != nextToken.Type,
			nextToken,
			fmt.Sprintf("Tests[%d] - Unexpected type. Expected: %q, Got: %q", idx, _token.Type, nextToken.Type),
		)

		assert.InvalidToken(
			_token.Literal != nextToken.Literal,
			nextToken,
			fmt.Sprintf("Tests[%d] - Unexpected type. Expected: %q, Got: %q", idx, _token.Type, nextToken.Type),
		)
	}
}
