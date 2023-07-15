package lexer

import "golang.org/x/exp/utf8string"

type Lexer struct {
	input *utf8string.String
	currPosition int
	nextposition int
	ch rune
}

func New(str string) *Lexer {
	return &Lexer{input: utf8string.NewString(str)}
}

func (l *Lexer) readChar() {
	if(l.nextposition >= l.input.RuneCount()) {
		l.ch = '\x00' // NUL char
	} else {
		l.ch = l.input.At(l.nextposition)
	}
	l.currPosition = l.nextposition
	l.nextposition++
}