package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	IDENTIFIER TokenType = "IDENTIFIER" // add, foobar, x, y, ...
	INT        TokenType = "INT"        // 123456

	EQUAL     TokenType = "="
	PLUS      TokenType = "+"
	MINUS     TokenType = "-"
	ASTERISK  TokenType = "*"
	SLASH     TokenType = "/"
	BANG      TokenType = "!"
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"

	LT     TokenType = "<"
	GT     TokenType = ">"
	LPAREN TokenType = "("
	RPAREN TokenType = ")"
	LBRACE TokenType = "{"
	RBRACE TokenType = "}"

	TRUE     TokenType = "TRUE"
	FALSE    TokenType = "FALSE"
	IF       TokenType = "IF"
	ELSE     TokenType = "ELSE"
	RETURN   TokenType = "RETURN"
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
)
