package token

const (
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	// Indentifiers
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"

	// Operators
	BANG     = "!"
	PLUS     = "+"
	MINUS    = "-"
	ASSIGN   = "="
	SLASH    = "/"
	ASTERISK = "*"
	LT       = "<"
	LTE      = "<="
	GT       = ">"
	GTE      = ">="
	EQ       = "=="
	NE       = "!="

	// Delimiters
	COMMA     = ","
	COLON     = ":"
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	VAR      = "VAR"
	RETURN   = "RETURN"
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"var":    VAR,
	"fn":     FUNCTION,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
}

func LookupIdent(ident string) TokenType {
	if token, ok := keywords[ident]; ok {
		return token
	}
	return IDENT
}

func NewToken(tokenType TokenType, ch byte) Token {
	return Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

func NewKeyword(tokenType TokenType, keyword []byte) Token {
	return Token{
		Type:    tokenType,
		Literal: string(keyword),
	}
}

func NewIdentifier(tokenType TokenType, ident []byte) Token {
	return Token{
		Type:    tokenType,
		Literal: string(ident),
	}
}
