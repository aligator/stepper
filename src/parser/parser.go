package parser

import "regexp"

type Utf8String []rune

type Parser interface {
	Peek() rune
	StartsWith(text string) bool
	Next() rune
	While(test func(r rune) bool) Utf8String
	Eof() bool
}

type BasicParser struct {
	position int
	input    Utf8String
}

func NewBasicParser(input string) BasicParser {
	return BasicParser{
		position: 0,
		input:    Utf8String(input),
	}
}

func (b BasicParser) Peek() rune {
	next := b.input[b.position]
	return next
}

func (b *BasicParser) Next() rune {
	next := b.input[b.position]
	b.position++
	return next
}

func (b *BasicParser) While(test func(r rune) bool) Utf8String {
	tillPos := b.position

	for !b.Eof() && test(b.input[tillPos]) {
		tillPos++
	}

	result := b.input[b.position:tillPos]

	b.position = tillPos

	return result
}

var validTagNameRune = regexp.MustCompile(`^[a-zA-Z0-9]$`)

func (b *BasicParser) TagName() Utf8String {
	return b.While(func(r rune) bool {
		return validTagNameRune.MatchString(string(r))
	})
}

func (b BasicParser) StartsWith(text string) bool {
	textRunes := Utf8String(text)

	for i, val := range textRunes {
		if val != b.input[i] {
			return false
		}
	}

	return true
}

func (b BasicParser) Eof() bool {
	return b.position >= len(b.input)
}

