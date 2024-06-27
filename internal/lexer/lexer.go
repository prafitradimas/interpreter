package lexer

import (
	"github.com/prafitradimas/interpreter/internal/token"
)

type Lexer struct {
	input        string
	position     int  // points to current char
	readPosition int  // points to current reading char (after current char)
	ch           byte // current char
}

func New(input string) *Lexer {
	lexer := &Lexer{
		input: input,
	}

	lexer.readChar()
	return lexer
}

func (lexer *Lexer) readChar() {
	// is EOF or have'nt read anything yet
	if lexer.readPosition >= len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.readPosition]
	}

	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.ch == ' ' || lexer.ch == '\t' || lexer.ch == '\r' || lexer.ch == '\n' {
		lexer.readChar()
	}
}

func (lexer *Lexer) readIdent() string {
	pos := lexer.position
	for isLetter(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[pos:lexer.position]
}

func isLetter(ch byte) bool {
	return ('a' <= ch && 'z' <= ch) || ('A' <= ch && 'Z' <= ch) || ch == '_'
}

func (lexer *Lexer) readNumber() string {
	pos := lexer.position
	for isNumber(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[pos:lexer.position]
}

func isNumber(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (lexer *Lexer) peekNextChar() byte {
	if lexer.readPosition >= len(lexer.input) {
		return 0
	}
	return lexer.input[lexer.readPosition]
}

func (lexer *Lexer) NextToken() token.Token {
	var _token token.Token

	// skiping whitespace
	lexer.skipWhitespace()

	switch lexer.ch {
	case '=':
		// peek next char
		if lexer.peekNextChar() == '=' {
			_token = token.Token{Type: token.EQ, Literal: string(lexer.ch) + string(lexer.peekNextChar())}

			// inc position
			lexer.readChar()
		} else {
			_token = token.NewToken(token.ASSIGN, lexer.ch)
		}
	case '+':
		_token = token.NewToken(token.PLUS, lexer.ch)
	case '-':
		_token = token.NewToken(token.MINUS, lexer.ch)
	case '(':
		_token = token.NewToken(token.LPAREN, lexer.ch)
	case ')':
		_token = token.NewToken(token.RPAREN, lexer.ch)
	case '{':
		_token = token.NewToken(token.LBRACE, lexer.ch)
	case '}':
		_token = token.NewToken(token.RBRACE, lexer.ch)
	case ';':
		_token = token.NewToken(token.SEMICOLON, lexer.ch)
	case ',':
		_token = token.NewToken(token.COMMA, lexer.ch)
	case '*':
		_token = token.NewToken(token.ASTERISK, lexer.ch)
	case '!':
		// peek next char
		if lexer.peekNextChar() == '=' {
			_token = token.Token{Type: token.NE, Literal: string(lexer.ch) + string(lexer.peekNextChar())}

			// inc position
			lexer.readChar()
		} else {
			_token = token.NewToken(token.BANG, lexer.ch)
		}
	case '/':
		_token = token.NewToken(token.SLASH, lexer.ch)
	case '>':
		if lexer.peekNextChar() == '=' {
			_token = token.Token{Type: token.GTE, Literal: string(lexer.ch) + string(lexer.peekNextChar())}
			// inc position
			lexer.readChar()
		} else {
			_token = token.NewToken(token.GT, lexer.ch)
		}
	case '<':
		if lexer.peekNextChar() == '=' {
			_token = token.Token{Type: token.LTE, Literal: string(lexer.ch) + string(lexer.peekNextChar())}
			// inc position
			lexer.readChar()
		} else {
			_token = token.NewToken(token.LT, lexer.ch)
		}
	case 0:
		_token.Type = token.EOF
		_token.Literal = ""
	default:
		if isLetter(lexer.ch) {
			str := lexer.readIdent()
			_token.Type = token.LookupIdent(str)
			_token.Literal = str
			return _token
		} else if isNumber(lexer.ch) {
			_token.Type = token.INT
			_token.Literal = lexer.readNumber()
			return _token
		} else {
			_token = token.NewToken(token.ILLEGAL, lexer.ch)
		}
	}

	lexer.readChar()
	return _token
}
