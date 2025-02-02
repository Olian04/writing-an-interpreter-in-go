package string_reader

import (
	"unicode"
)

type StringReader struct {
	input    string
	inputLen int
	readHead int
}

func NewStringReader(input string) *StringReader {
	return &StringReader{
		input:    input,
		inputLen: len(input),
		readHead: 0,
	}
}

func (sr *StringReader) ResetReadHead() {
	sr.readHead = 0
}

func (sr *StringReader) AdvanceReadHead(step int) {
	sr.readHead += step
	if sr.readHead > sr.inputLen {
		sr.readHead = sr.inputLen
	} else if sr.readHead < 0 {
		sr.readHead = 0
	}
}

func (sr *StringReader) IsEOF() bool {
	return sr.readHead >= sr.inputLen
}

func (sr *StringReader) ReadChar() rune {
	if sr.IsEOF() {
		return 0
	}
	return rune(sr.input[sr.readHead])
}

func (sr *StringReader) ReadWhile(predicate func(r rune) bool) string {
	if sr.IsEOF() {
		return ""
	}
	readHead := sr.readHead
	for readHead < sr.inputLen && predicate(rune(sr.input[readHead])) {
		readHead++
	}
	return sr.input[sr.readHead:readHead]
}

func (sr *StringReader) ReadWhileWithIndex(predicate func(r rune, i int) bool) string {
	i := 0
	return sr.ReadWhile(func(r rune) bool {
		val := predicate(r, i)
		i++
		return val
	})
}

func (sr *StringReader) ReadWord() string {
	return sr.ReadWhileWithIndex(func(r rune, i int) bool {
		if i == 0 {
			return unicode.IsLetter(r) || r == '_'
		}
		return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_'
	})
}

func (sr *StringReader) ReadWhitespace() string {
	return sr.ReadWhile(unicode.IsSpace)
}
