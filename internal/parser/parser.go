package parser

import (
	"fmt"

	"github.com/prafitradimas/interpreter/internal/ast"
	"github.com/prafitradimas/interpreter/internal/lexer"
	"github.com/prafitradimas/interpreter/internal/token"
)

type Parser struct {
	lexer        *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
	errors       []string
}

func New(lexr *lexer.Lexer) *Parser {
	parser := Parser{
		lexer:  lexr,
		errors: []string{},
	}

	parser.nextToken()
	parser.nextToken()

	return &parser
}

func (ps *Parser) peekError(_tokenType token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", _tokenType, ps.peekToken.Type)
	ps.errors = append(ps.errors, msg)
}

func (ps *Parser) Errors() []string {
	return ps.errors
}

func (ps *Parser) nextToken() {
	ps.currentToken = ps.peekToken
	ps.peekToken = ps.lexer.NextToken()
}

func (ps *Parser) Parse() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

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
	case token.RETURN:
		return ps.parseReturnStatement()
	default:
		return nil
	}
}

func (ps *Parser) parseVarStatement() ast.Statement {
	statement := &ast.VarStatement{Token: ps.currentToken}

	if !ps.expectNextToken(token.IDENT) {
		return nil
	}

	statement.Name = &ast.Identifier{Token: ps.currentToken, Value: ps.currentToken.Literal}

	if !ps.expectNextToken(token.ASSIGN) {
		return nil
	}

	// TODO:
	// skipping expression until encounter semicolon
	for !ps.expectCurrentToken(token.SEMICOLON) {
		ps.nextToken()
	}

	return statement
}

func (ps *Parser) expectCurrentToken(_tokenType token.TokenType) bool {
	return ps.currentToken.Type == _tokenType
}

func (ps *Parser) expectNextToken(_tokenType token.TokenType) bool {
	if ps.peekToken.Type == _tokenType {
		ps.nextToken()
		return true
	}

	ps.peekError(_tokenType)
	return false
}

func (ps *Parser) parseReturnStatement() ast.Statement {
	stmt := &ast.ReturnStatement{Token: ps.currentToken}
	ps.nextToken()

	for !ps.expectCurrentToken(token.SEMICOLON) {
		ps.nextToken()
	}

	return stmt
}
