package lexer

import (
	"golang.org/x/exp/utf8string"
	"rocket/src/token"
)

const NULL_CHAR = '\x00'

type Lexer struct {
	input        *utf8string.String
	currPosition int
	nextposition int
	ch           rune
	str          string
	// meta info
	currentLine int
	currentCol  int
}

func New(str string) *Lexer {
	l := &Lexer{input: utf8string.NewString(str), currentLine: 1, currentCol: 1}
	l.readChar()
	return l
}

// LEXER Type functions

// this moves the needle
func (l *Lexer) readChar() {
	if l.nextposition >= l.input.RuneCount() {
		l.ch = NULL_CHAR // NUL char
	} else {
		l.ch = l.input.At(l.nextposition)
		l.str = string(l.ch)
	}
	l.currPosition = l.nextposition
	l.nextposition++
	l.currentCol++
}

func (l *Lexer) createToken(t token.TokenType, s string) token.Token {
	c := l.currentCol - utf8string.NewString(s).RuneCount()
	return token.Token{
		Literal: s,
		Type: t,
		Ln: l.currentLine,
		Col: c,
	}
}
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.eatWhitespace()

	switch l.ch {
	case NULL_CHAR:
		tok = l.createToken(token.EOF, "")
	case '=': // = or ==
		s := string(l.ch)
		var tt token.TokenType = token.ASSIGN
		if l.peekChar() == '=' {
			l.readChar()
			s += string(l.ch)
			tt = token.EQ
		}
		tok = l.createToken(tt, s)
	case '!': // ! or !=
		s := string(l.ch)
		var tt token.TokenType = token.BANG
		if l.peekChar() == '=' {
			l.readChar()
			s += string(l.ch)
			tt = token.NOT_EQ
		}
		tok = l.createToken(tt, s)
	case '*': // * or **
		s := string(l.ch)
		var tt token.TokenType = token.ASTERISK
		if l.peekChar() == '*' {
			l.readChar()
			s += string(l.ch)
			tt = token.EXPONENT
		}
		tok = l.createToken(tt, s)
	case '+':
		tok = l.createToken(token.PLUS, string(l.ch))
	case ':':
		tok = l.createToken(token.COLON, string(l.ch))
	case '<':
		tok = l.createToken(token.LT, string(l.ch))
	case '>':
		tok = l.createToken(token.GT, string(l.ch))
	case '/':
		tok = l.createToken(token.SLASH, string(l.ch))
	case '-':
		tok = l.createToken(token.MINUS, string(l.ch))
	case ',':
		tok = l.createToken(token.COMMA, string(l.ch))
	case ';':
		tok = l.createToken(token.SEMICOLON, string(l.ch))
	case '(':
		tok = l.createToken(token.LPAREN, string(l.ch))
	case ')':
		tok = l.createToken(token.RPAREN, string(l.ch))
	case '{':
		tok = l.createToken(token.LBRACE, string(l.ch))
	case '}':
		tok = l.createToken(token.RBRACE, string(l.ch))
	default: // keywords, identifiers, numbers, strings
		if isDigit(l.ch) {
			literal := l.parseNumber()
			tok = l.createToken(token.INT, literal)
		} else if isLetter(l.ch) {
			literal := l.parseString()
			var tt token.TokenType
			if t, ok := token.Keywords[literal]; ok {
				tt = t
			} else {
				tt = token.IDENT
			}
			tok = l.createToken(tt, literal)
		} else {
			tok = l.createToken(token.ILLEGAL, "")
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) peekChar() rune {
	if l.nextposition >= l.input.RuneCount() {
		return NULL_CHAR // NUL char
	}
	return l.input.At(l.nextposition)
}

func (l *Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		if l.ch == '\n' { // hopefully this will handle all kinds of line endings
			l.currentLine++
			l.currentCol = 1
		}
		l.readChar()
	}
}

func (l *Lexer) parseNumber() string {
	s := ""
	for ;isDigit(l.ch); {
		s += string(l.ch)
		if !isDigit(l.peekChar()) {
			return s
		}
		l.readChar()
	}
	return s
}

func (l *Lexer) parseString() string {
	s := ""
	for ;isLetter(l.ch); {
		s += string(l.ch)
		if !isLetter(l.peekChar()) {
			return s
		}
		l.readChar()
	}
	return s
}

// Other utils

func isDigit(char rune) bool {
	if char >= '0' && char <= '9' {
		return true
	} 
	return false
}

func isLetter(char rune) bool {
	if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z' ) || char == '_' || char == '🚀' {
		return true
	}
	return false
}