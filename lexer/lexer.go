package lexer

import (
	"unicode"

	"github.com/Olian04/monkey/lexer/string_reader"
	"github.com/Olian04/monkey/token"
)

type Lexer struct {
	sr *string_reader.StringReader
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		sr: string_reader.NewStringReader(input),
	}
}

func (l *Lexer) readSyntax() (*token.Token, bool) {
	switch l.sr.ReadChar() {
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
	switch l.sr.ReadWord() {
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
	firstChar := l.sr.ReadChar()

	switch {
	case unicode.IsNumber(firstChar):
		return &token.Token{
			Type:    token.INT,
			Literal: l.sr.ReadWhile(unicode.IsNumber),
		}, true
	case unicode.IsLetter(firstChar):
		return &token.Token{
			Type:    token.IDENT,
			Literal: l.sr.ReadWord(),
		}, true
	}
	return nil, false
}

func (l *Lexer) NextToken() token.Token {
	whitespace := l.sr.ReadWhitespace()
	l.sr.AdvanceReadHead(len(whitespace))

	if l.sr.IsEOF() {
		return token.Token{
			Type:    token.EOF,
			Literal: "",
		}
	}

	if tok, ok := l.readSyntax(); ok {
		l.sr.AdvanceReadHead(len(tok.Literal))
		return *tok
	}
	if tok, ok := l.readKeyword(); ok {
		l.sr.AdvanceReadHead(len(tok.Literal))
		return *tok
	}
	if tok, ok := l.readValue(); ok {
		l.sr.AdvanceReadHead(len(tok.Literal))
		return *tok
	}

	literal := l.sr.ReadWord()
	l.sr.AdvanceReadHead(len(literal))
	return token.Token{
		Type:    token.ILLEGAL,
		Literal: literal,
	}
}
