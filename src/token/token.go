package token

import "fmt"

type TokenType string

type Token struct {
	Type TokenType
	Literal string
	Ln int
	Col int
}

func New(t TokenType, l string, ln int, col int) Token {
	return Token{Type: t, Literal: l, Ln: ln, Col: col }
}

func (tok *Token) Inspect() {
	fmt.Printf("%d:%d\t%s (%s)\n", tok.Ln, tok.Col, tok.Type, tok.Literal)
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...

	// data types
	INT = "INT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
	RETURN = "RETURN"
	TRUE = "TRUE"
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"
)

var Keywords = map[string]TokenType{
	"func": FUNCTION,
	"let": LET,
	"return": RETURN,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE, 
}