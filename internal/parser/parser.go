package parser

import (
	"github.com/prafitradimas/interpreter/internal/ast"
	"github.com/prafitradimas/interpreter/internal/lexer"
	"github.com/prafitradimas/interpreter/internal/token"
)

type Parser struct {
	lexer *lexer.Lexer

	currentToken token.Token
	peekToken    token.Token
}

func New(lexr *lexer.Lexer) *Parser {
	parser := Parser{
		lexer: lexr,
	}

	return &parser
}

func (ps *Parser) nextToken() {
	ps.currentToken = ps.peekToken
	ps.peekToken = ps.lexer.NextToken()
}

func (ps *Parser) Parse() *ast.Program {
	program := &ast.Program{
		Statements: []ast.Statement{},
	}

	for ps.currentToken.Type != token.EOF {
		statement := ps.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}

		ps.nextToken()
	}

	return program
}

func (ps *Parser) parseStatement() ast.Statement {
	switch ps.currentToken.Type {
	case token.VAR:
		return ps.parseVarStatement()
	default:
		return nil
	}
}

func (ps *Parser) parseVarStatement() ast.Statement {
	statement := &ast.VarStatement{
		Token: ps.currentToken,
		Name:  &ast.Identifier{},
	}

	// next token should be identifier
	if ps.peekToken.Type != token.IDENT {
		return nil
	}

	// get the variable's name
	ps.nextToken()
	statement.Name.Token = ps.currentToken

	// next token should be `=`
	if ps.peekToken.Type != token.ASSIGN {
		return nil
	}

	ps.nextToken()
	statement.Name.Value = ps.peekToken.Literal

	for ps.currentToken.Type != token.SEMICOLON {
		ps.nextToken()
	}

	return statement
}

func (ps *Parser) expectCurrentToken(_token token.TokenType) bool {
	return ps.currentToken.Type == _token
}

func (ps *Parser) expectNextToken(_token token.TokenType) bool {
	return ps.peekToken.Type == _token
}
