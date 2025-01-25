package monkey_test

import (
	"testing"

	"github.com/Olian04/monkey/lexer"
)

func TestLexer(t *testing.T) {
	lexer := lexer.NewLexer("let five = 5;")
	t.Fatalf("test failed")
}
