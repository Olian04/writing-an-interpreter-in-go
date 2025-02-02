package monkey_test

import (
	"testing"

	"github.com/Olian04/monkey/lexer/string_reader"
)

func RunStringReaderTest(t *testing.T, sr *string_reader.StringReader, expected []struct {
	Reader   func() string
	Expected string
}) {
	for _, expected := range expected {
		got := expected.Reader()
		if got != expected.Expected {
			t.Fatalf("expected=\"%s\", got=\"%s\"", expected.Expected, got)
		}
		sr.AdvanceReadHead(len(got))
	}
}

func TestReadWord(t *testing.T) {
	sr := string_reader.NewStringReader("hello world")
	RunStringReaderTest(t, sr, []struct {
		Reader   func() string
		Expected string
	}{
		{sr.ReadWord, "hello"},
		{sr.ReadWhitespace, " "},
		{sr.ReadWord, "world"},
	})
}
