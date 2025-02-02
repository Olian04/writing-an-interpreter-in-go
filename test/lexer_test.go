package monkey_test

import (
	"testing"

	"github.com/Olian04/monkey/lexer"
	"github.com/Olian04/monkey/token"
)

func RunTokenTest(t *testing.T, input string, expected []token.Token) {
	l := lexer.NewLexer(input)
	for _, expectedToken := range expected {
		gotToken := l.NextToken()
		if gotToken.Type != expectedToken.Type {
			t.Fatalf("tok.Type is not %q. got=%q", expectedToken.Type, gotToken.Type)
		}
		if gotToken.Literal != expectedToken.Literal {
			t.Fatalf("tok.Literal is not \"%q\". got=\"%q\"", expectedToken.Literal, gotToken.Literal)
		}
	}
}

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	RunTokenTest(t, input, []token.Token{
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.LBRACE, Literal: "{"},
		{Type: token.RBRACE, Literal: "}"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.SEMICOLON, Literal: ";"},
	})
}

func TestLetExpression(t *testing.T) {
	input := `let x = 5;`
	RunTokenTest(t, input, []token.Token{
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENT, Literal: "x"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
	})
}

func TestMultilineLetExpressions(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 838383;
	`
	RunTokenTest(t, input, []token.Token{
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENT, Literal: "x"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENT, Literal: "y"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.INT, Literal: "10"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENT, Literal: "foobar"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.INT, Literal: "838383"},
		{Type: token.SEMICOLON, Literal: ";"},
	})
}

func TestFunctionExpression(t *testing.T) {
	input := `fn(x) { x + 2; }`
	RunTokenTest(t, input, []token.Token{
		{Type: token.FUNCTION, Literal: "fn"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.IDENT, Literal: "x"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.LBRACE, Literal: "{"},
		{Type: token.IDENT, Literal: "x"},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.INT, Literal: "2"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.RBRACE, Literal: "}"},
	})
}
