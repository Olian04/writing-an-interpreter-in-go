package lexer

import (
	"github.com/Olian04/monkey/token"
	"github.com/Olian04/monkey/util"
)

type Lexer struct {
	input      string
	readHead   int
	charOffset int
	char       byte
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.advanceReadHead()
	return l
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()

	// Syntax token
	if tokType := l.lookupSingleCharToken(l.char); tokType != token.ILLEGAL {
		literal := string(tokType)
		l.advanceReadHead()
		return token.Token{
			Type:    tokType,
			Literal: literal,
		}
	}

	// Keyword or identifier
	if util.IsLetter(l.char) {
		ident := l.readIdentifier()
		if tokType := l.lookupKeyword(ident); tokType != token.ILLEGAL {
			return token.Token{
				Type:    tokType,
				Literal: ident,
			}
		}
		return token.Token{
			Type:    token.IDENTIFIER,
			Literal: ident,
		}
	}

	// Integer literal
	if util.IsDigit(l.char) {
		return token.Token{
			Type:    token.INT,
			Literal: l.readNumber(),
		}
	}

	// Illegal token
	literal := string(l.char)
	l.advanceReadHead()
	return token.Token{
		Type:    token.ILLEGAL,
		Literal: literal,
	}
}

func (l *Lexer) peekChar() byte {
	if l.readHead >= len(l.input) {
		return 0
	}
	return l.input[l.readHead]
}

func (l *Lexer) advanceReadHead() {
	if l.readHead >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readHead]
	}
	l.charOffset = l.readHead
	l.readHead++
}

func (l *Lexer) readIdentifier() string {
	pos := l.charOffset
	for util.IsLetter(l.char) || util.IsDigit(l.char) {
		l.advanceReadHead()
	}
	return l.input[pos:l.charOffset]
}

func (l *Lexer) readNumber() string {
	pos := l.charOffset
	for util.IsDigit(l.char) {
		l.advanceReadHead()
	}
	return l.input[pos:l.charOffset]
}

func (l *Lexer) skipWhitespace() {
	for util.IsWhitespace(l.char) {
		l.advanceReadHead()
	}
}

func (l *Lexer) lookupSingleCharToken(char byte) token.TokenType {
	switch char {
	case '=':
		return token.EQUAL
	case '+':
		return token.PLUS
	case '-':
		return token.MINUS
	case '*':
		return token.ASTERISK
	case '/':
		return token.SLASH
	case '!':
		return token.BANG
	case ',':
		return token.COMMA
	case ';':
		return token.SEMICOLON
	case '(':
		return token.LPAREN
	case ')':
		return token.RPAREN
	case '{':
		return token.LBRACE
	case '}':
		return token.RBRACE
	case '<':
		return token.LT
	case '>':
		return token.GT
	case 0:
		return token.EOF
	}
	return token.ILLEGAL
}

func (l *Lexer) lookupKeyword(ident string) token.TokenType {
	switch ident {
	case "fn":
		return token.FUNCTION
	case "let":
		return token.LET
	case "if":
		return token.IF
	case "else":
		return token.ELSE
	case "return":
		return token.RETURN
	case "true":
		return token.TRUE
	case "false":
		return token.FALSE
	}
	return token.ILLEGAL
}
