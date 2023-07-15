package lexer

import (
	"rocket/src/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
  let 🚀five = 5;
let ten = 10;
let add = func(x, y) {
  x + y;
};
let result = add(five, ten);
`
	expectedResult := []struct {
		expectedType token.TokenType
		expectedLiteral string
	} {
		{ token.LET, "let"},
		{ token.IDENT, "🚀five"},
		{ token.ASSIGN, "="},
		{ token.INT, "5"},
		{ token.SEMICOLON, ";"},

		{ token.LET, "let"},
		{ token.IDENT, "ten"},
		{ token.ASSIGN, "="},
		{ token.INT, "10"},
		{ token.SEMICOLON, ";"},

		{ token.LET, "let"},
		{ token.IDENT, "add"},
		{ token.ASSIGN, "="},
		{ token.FUNCTION, "func"},
		{ token.LPAREN, "("},
		{ token.IDENT, "x"},
		{ token.COMMA, ","},
		{ token.IDENT, "y"},
		{ token.RPAREN, ")"},
		{ token.LBRACE, "{"},

		{ token.IDENT, "x"},
		{ token.PLUS, "+"},
		{ token.IDENT, "y"},
		{ token.SEMICOLON, ";"},

		{ token.RBRACE, "}"},
		{ token.SEMICOLON, ";"},

		{ token.LET, "let"},
		{ token.IDENT, "result"},
		{ token.ASSIGN, "="},
		{ token.IDENT, "add"},
		{ token.LPAREN, "("},
		{ token.IDENT, "five"},
		{ token.COMMA, ","},
		{ token.IDENT, "ten"},
		{ token.RPAREN, ")"},
		{ token.SEMICOLON, ";"},
	}

	l := New(input)
	for i, val := range expectedResult {
		tok := l.NextToken()
		t.Error(tok.Literal, tok.Ln, tok.Col)
		// check type
		if val.expectedType != tok.Type {
			t.Errorf("Tests[%d] - wrong tokenType. Expected: %q, got: %q", i, val.expectedType, tok.Type)
		}
		// check literal
		if val.expectedLiteral != tok.Literal {
			t.Errorf("Tests[%d] - wrong tokenLiteral. Expected: %q, got: %q", i, val.expectedLiteral, tok.Literal)
		}
	}
}