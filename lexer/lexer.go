package lexer

import (
	"github.com/Olian04/monkey/token"
	"github.com/Olian04/monkey/util"
)

type Lexer struct {
	input    string
	charHead int
	readHead int
	char     byte
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.advanceOneChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	switch l.char {
	case ';':
		tok = token.Token{Type: token.SEMICOLON, Literal: ";"}
	case '=':
		tok = token.Token{Type: token.ASSIGN, Literal: "="}
	case '(':
		tok = token.Token{Type: token.LPAREN, Literal: "("}
	case ')':
		tok = token.Token{Type: token.RPAREN, Literal: ")"}
	case '{':
		tok = token.Token{Type: token.LBRACE, Literal: "{"}
	case '}':
		tok = token.Token{Type: token.RBRACE, Literal: "}"}
	case '+':
		tok = token.Token{Type: token.PLUS, Literal: "+"}
	case ',':
		tok = token.Token{Type: token.COMMA, Literal: ","}
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
	default:
		if util.IsLetter(l.char) {
			ident := l.readIdentifier()
			tok = token.Token{Type: token.LookupIdent(ident), Literal: ident}
			return tok
		} else if util.IsDigit(l.char) {
			tok = token.Token{Type: token.INT, Literal: l.readNumber()}
			return tok
		} else {
			tok = token.Token{Type: token.ILLEGAL, Literal: string(l.char)}
		}
	}
	l.advanceOneChar()
	return tok
}

func (l *Lexer) advanceOneChar() {
	if l.readHead >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readHead]
	}
	l.charHead = l.readHead
	l.readHead++
}

func (l *Lexer) readIdentifier() string {
	pos := l.charHead
	for util.IsLetter(l.char) || util.IsDigit(l.char) {
		l.advanceOneChar()
	}
	return l.input[pos:l.charHead]
}

func (l *Lexer) readNumber() string {
	pos := l.charHead
	for util.IsDigit(l.char) {
		l.advanceOneChar()
	}
	return l.input[pos:l.charHead]
}

func (l *Lexer) skipWhitespace() {
	for util.IsWhitespace(l.char) {
		l.advanceOneChar()
	}
}
