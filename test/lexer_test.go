package monkey_test

import (
	"testing"

	"github.com/Olian04/monkey/lexer"
	"github.com/Olian04/monkey/token"
)

type ILexer interface {
	NextToken() token.Token
}

type Tester interface {
	Fatalf(format string, args ...any)
}

type TestContext struct {
	Name   string
	Lexer  ILexer
	Tester Tester
}

func RunTokenTest(t Tester, lexer ILexer, expected []token.Token) {
	for _, expectedToken := range expected {
		gotToken := lexer.NextToken()
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
	RunTokenTest(t, lexer.New(input), []token.Token{
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

func TestSmallProgram(t *testing.T) {
	input := `
		let five = 5;
		let ten = 10;

		let add = fn(x, y) {
			x + y;
		};

		let result = add(five, ten);
	`
	RunTokenTest(t, lexer.New(input), []token.Token{
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENT, Literal: "five"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENT, Literal: "ten"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.INT, Literal: "10"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENT, Literal: "add"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.FUNCTION, Literal: "fn"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.IDENT, Literal: "x"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.IDENT, Literal: "y"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.LBRACE, Literal: "{"},
		{Type: token.IDENT, Literal: "x"},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.IDENT, Literal: "y"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.RBRACE, Literal: "}"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENT, Literal: "result"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.IDENT, Literal: "add"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.IDENT, Literal: "five"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.IDENT, Literal: "ten"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.SEMICOLON, Literal: ";"},
	})
}

func BenchmarkLexer(b *testing.B) {
	input := `
		let five = 5;
		let ten = 10;

		let add = fn(x, y) {
			x + y;
		};

		let result = add(five, ten);
	`
	for i := 0; i < b.N; i++ {
		RunTokenTest(b, lexer.New(input), []token.Token{
			{Type: token.LET, Literal: "let"},
			{Type: token.IDENT, Literal: "five"},
			{Type: token.ASSIGN, Literal: "="},
			{Type: token.INT, Literal: "5"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.LET, Literal: "let"},
			{Type: token.IDENT, Literal: "ten"},
			{Type: token.ASSIGN, Literal: "="},
			{Type: token.INT, Literal: "10"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.LET, Literal: "let"},
			{Type: token.IDENT, Literal: "add"},
			{Type: token.ASSIGN, Literal: "="},
			{Type: token.FUNCTION, Literal: "fn"},
			{Type: token.LPAREN, Literal: "("},
			{Type: token.IDENT, Literal: "x"},
			{Type: token.COMMA, Literal: ","},
			{Type: token.IDENT, Literal: "y"},
			{Type: token.RPAREN, Literal: ")"},
			{Type: token.LBRACE, Literal: "{"},
			{Type: token.IDENT, Literal: "x"},
			{Type: token.PLUS, Literal: "+"},
			{Type: token.IDENT, Literal: "y"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.RBRACE, Literal: "}"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.LET, Literal: "let"},
			{Type: token.IDENT, Literal: "result"},
			{Type: token.ASSIGN, Literal: "="},
			{Type: token.IDENT, Literal: "add"},
			{Type: token.LPAREN, Literal: "("},
			{Type: token.IDENT, Literal: "five"},
			{Type: token.COMMA, Literal: ","},
			{Type: token.IDENT, Literal: "ten"},
			{Type: token.RPAREN, Literal: ")"},
			{Type: token.SEMICOLON, Literal: ";"},
		})
	}
}
