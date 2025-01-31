package lexer

import (
	"unicode"

	"github.com/Olian04/monkey/token"
)

type Lexer struct {
	input        string
	readPosition int
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input:        input,
		readPosition: 0,
	}
}

func (l *Lexer) readChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return rune(l.input[l.readPosition])
}

func (l *Lexer) readWord() string {
	if l.readPosition >= len(l.input) {
		return ""
	}
	i := 0
	for l.readPosition+i < len(l.input) && !unicode.IsSpace(rune(l.input[l.readPosition+i])) {
		i++
	}
	return l.input[l.readPosition : l.readPosition+i]
}

func (l *Lexer) readWhitespace() string {
	i := 0
	for l.readPosition+i < len(l.input) && unicode.IsSpace(rune(l.input[l.readPosition+i])) {
		i++
	}
	return l.input[l.readPosition : l.readPosition+i]
}

func (l *Lexer) readSyntax() (*token.Token, bool) {
	switch l.readChar() {
	case '=':
		return &token.Token{
			Type:    token.ASSIGN,
			Literal: "=",
		}, true
	case '+':
		return &token.Token{
			Type:    token.PLUS,
			Literal: "+",
		}, true
	case '(':
		return &token.Token{
			Type:    token.LPAREN,
			Literal: "(",
		}, true
	case ')':
		return &token.Token{
			Type:    token.RPAREN,
			Literal: ")",
		}, true
	case '{':
		return &token.Token{
			Type:    token.LBRACE,
			Literal: "{",
		}, true
	case '}':
		return &token.Token{
			Type:    token.RBRACE,
			Literal: "}",
		}, true
	case ',':
		return &token.Token{
			Type:    token.COMMA,
			Literal: ",",
		}, true
	case ';':
		return &token.Token{
			Type:    token.SEMICOLON,
			Literal: ";",
		}, true
	}
	return nil, false
}

func (l *Lexer) readKeyword() (*token.Token, bool) {
	literal := l.readWord()
	switch literal {
	case "let":
		return &token.Token{
			Type:    token.LET,
			Literal: "let",
		}, true
	case "fn":
		return &token.Token{
			Type:    token.FUNCTION,
			Literal: "fn",
		}, true
	}
	return nil, false
}

func (l *Lexer) readValue() (*token.Token, bool) {
	firstChar := l.readChar()
	literal := l.readWord()

	switch {
	case unicode.IsNumber(firstChar):
		return &token.Token{
			Type:    token.INT,
			Literal: literal,
		}, true
	case unicode.IsLetter(firstChar):
		return &token.Token{
			Type:    token.IDENT,
			Literal: literal,
		}, true
	}
	return nil, false
}

func (l *Lexer) advance(step int) {
	l.readPosition += step
}

func (l *Lexer) NextToken() token.Token {
	whitespace := l.readWhitespace()
	l.advance(len(whitespace))

	if l.readPosition >= len(l.input) {
		return token.Token{
			Type:    token.EOF,
			Literal: "",
		}
	}

	if tok, ok := l.readSyntax(); ok {
		l.advance(len(tok.Literal))
		return *tok
	}
	if tok, ok := l.readKeyword(); ok {
		l.advance(len(tok.Literal))
		return *tok
	}
	if tok, ok := l.readValue(); ok {
		l.advance(len(tok.Literal))
		return *tok
	}

	literal := l.readWord()
	l.advance(len(literal))
	return token.Token{
		Type:    token.ILLEGAL,
		Literal: literal,
	}
}
