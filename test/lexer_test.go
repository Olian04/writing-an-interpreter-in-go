package monkey_test

import (
	"testing"

	"github.com/Olian04/monkey/lexer"
	"github.com/Olian04/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	l := lexer.NewLexer(input)

	if tok := l.NextToken(); tok.Type != token.ASSIGN {
		t.Fatalf("tok.Type is not ASSIGN. got=%q, tok.Literal=%q", tok.Type, tok.Literal)
	}

	if tok := l.NextToken(); tok.Type != token.PLUS {
		t.Fatalf("tok.Type is not PLUS. got=%q, tok.Literal=%q", tok.Type, tok.Literal)
	}

	if tok := l.NextToken(); tok.Type != token.LPAREN {
		t.Fatalf("tok.Type is not LPAREN. got=%q, tok.Literal=%q", tok.Type, tok.Literal)
	}

	if tok := l.NextToken(); tok.Type != token.RPAREN {
		t.Fatalf("tok.Type is not RPAREN. got=%q, tok.Literal=%q", tok.Type, tok.Literal)
	}

	if tok := l.NextToken(); tok.Type != token.LBRACE {
		t.Fatalf("tok.Type is not LBRACE. got=%q, tok.Literal=%q", tok.Type, tok.Literal)
	}

	if tok := l.NextToken(); tok.Type != token.RBRACE {
		t.Fatalf("tok.Type is not RBRACE. got=%q, tok.Literal=%q", tok.Type, tok.Literal)
	}

	if tok := l.NextToken(); tok.Type != token.COMMA {
		t.Fatalf("tok.Type is not COMMA. got=%q, tok.Literal=%q", tok.Type, tok.Literal)
	}

	if tok := l.NextToken(); tok.Type != token.SEMICOLON {
		t.Fatalf("tok.Type is not SEMICOLON. got=%q, tok.Literal=%q", tok.Type, tok.Literal)
	}

	if tok := l.NextToken(); tok.Type != token.EOF {
		t.Fatalf("tok.Type is not EOF. got=%q, tok.Literal=%q", tok.Type, tok.Literal)
	}
}

func TestLetExpression(t *testing.T) {
	input := `let x = 5;`
	l := lexer.NewLexer(input)

	if tok := l.NextToken(); tok.Type != token.LET {
		t.Fatalf("tok.Type is not LET. got=%q, tok.Literal=%q", tok.Type, tok.Literal)
	}

	if tok := l.NextToken(); tok.Type != token.IDENT {
		t.Fatalf("tok.Type is not IDENT. got=%q, tok.Literal=%q", tok.Type, tok.Literal)
	}

	if tok := l.NextToken(); tok.Type != token.ASSIGN {
		t.Fatalf("tok.Type is not ASSIGN. got=%q, tok.Literal=%q", tok.Type, tok.Literal)
	}

	if tok := l.NextToken(); tok.Type != token.INT {
		t.Fatalf("tok.Type is not INT. got=%q, tok.Literal=%q", tok.Type, tok.Literal)
	}

	if tok := l.NextToken(); tok.Type != token.SEMICOLON {
		t.Fatalf("tok.Type is not SEMICOLON. got=%q, tok.Literal=%q", tok.Type, tok.Literal)
	}

	if tok := l.NextToken(); tok.Type != token.EOF {
		t.Fatalf("tok.Type is not EOF. got=%q, tok.Literal=%q", tok.Type, tok.Literal)
	}
}
