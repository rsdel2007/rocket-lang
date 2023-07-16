package lexer

import (
	"rocket/src/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let 🚀five = 5;
let ten = 10;

let add = func(x, y) {
  x + y;
};

let result = add(five, ten);
!-/*5;
5 < 10 > 5;

if (5 < 10) {
	return true;
} else {
	return false;
}

10 == 10;
10 != 9;
5 ** 5;
`
	expectedResult := []struct {
		expectedType token.TokenType
		expectedLiteral string
		ln int
		col int
	} {
		{token.LET, "let", 1, 1},
		{token.IDENT, "🚀five", 1, 5},
		{token.ASSIGN, "=", 1, 11},
		{token.INT, "5", 1, 13},
		{token.SEMICOLON, ";", 1, 14},
		{token.LET, "let", 2, 1},
		{token.IDENT, "ten", 2, 5},
		{token.ASSIGN, "=", 2, 9},
		{token.INT, "10", 2, 11},
		{token.SEMICOLON, ";", 2, 13},
		{token.LET, "let", 4, 1},
		{token.IDENT, "add", 4, 5},
		{token.ASSIGN, "=", 4, 9},
		{token.FUNCTION, "func", 4, 11},
		{token.LPAREN, "(", 4, 15},
		{token.IDENT, "x", 4, 16},
		{token.COMMA, ",", 4, 17},
		{token.IDENT, "y", 4, 19},
		{token.RPAREN, ")", 4, 20},
		{token.LBRACE, "{", 4, 22},
		{token.IDENT, "x", 5, 3},
		{token.PLUS, "+", 5, 5},
		{token.IDENT, "y", 5, 7},
		{token.SEMICOLON, ";", 5, 8},
		{token.RBRACE, "}", 6, 1},
		{token.SEMICOLON, ";", 6, 2},
		{token.LET, "let", 8, 1},
		{token.IDENT, "result", 8, 5},
		{token.ASSIGN, "=", 8, 12},
		{token.IDENT, "add", 8, 14},
		{token.LPAREN, "(", 8, 17},
		{token.IDENT, "five", 8, 18},
		{token.COMMA, ",", 8, 22},
		{token.IDENT, "ten", 8, 24},
		{token.RPAREN, ")", 8, 27},
		{token.SEMICOLON, ";", 8, 28},
		{token.BANG, "!", 9, 1},
		{token.MINUS, "-", 9, 2},
		{token.SLASH, "/", 9, 3},
		{token.ASTERISK, "*", 9, 4},
		{token.INT, "5", 9, 5},
		{token.SEMICOLON, ";", 9, 6},
		{token.INT, "5", 10, 1},
		{token.LT, "<", 10, 3},
		{token.INT, "10", 10, 5},
		{token.GT, ">", 10, 8},
		{token.INT, "5", 10, 10},
		{token.SEMICOLON, ";", 10, 11},
		{token.IF, "if", 12, 1},
		{token.LPAREN, "(", 12, 4},
		{token.INT, "5", 12, 5},
		{token.LT, "<", 12, 7},
		{token.INT, "10", 12, 9},
		{token.RPAREN, ")", 12, 11},
		{token.LBRACE, "{", 12, 13},
		{token.RETURN, "return", 13, 2},
		{token.TRUE, "true", 13, 9},
		{token.SEMICOLON, ";", 13, 13},
		{token.RBRACE, "}", 14, 1},
		{token.ELSE, "else", 14, 3},
		{token.LBRACE, "{", 14, 8},
		{token.RETURN, "return", 15, 2},
		{token.FALSE, "false", 15, 9},
		{token.SEMICOLON, ";", 15, 14},
		{token.RBRACE, "}", 16, 1},
		{token.INT, "10", 18, 1},
		{token.EQ, "==", 18, 4},
		{token.INT, "10", 18, 7},
		{token.SEMICOLON, ";", 18, 9},
		{token.INT, "10", 19, 1},
		{token.NOT_EQ, "!=", 19, 4},
		{token.INT, "9", 19, 7},
		{token.SEMICOLON, ";", 19, 8},
		{token.INT, "5", 20, 1},
		{token.EXPONENT, "**", 20, 3},
		{token.INT, "5", 20, 6},
		{token.SEMICOLON, ";", 20, 7},
		{token.EOF, "", 21, 2},
	}

	l := New(input)
	for i, val := range expectedResult {
		tok := l.NextToken()
		// check type
		if val.expectedType != tok.Type {
			t.Errorf("Tests[%d] - wrong tokenType. Expected: %q, got: %q | Ln: %d : col : %d", i, val.expectedType, tok.Type, tok.Ln, tok.Col)
		}
		// check literal
		if val.expectedLiteral != tok.Literal {
			t.Errorf("Tests[%d] - wrong tokenLiteral. Expected: %q, got: %q | Ln: %d : col : %d", i, val.expectedLiteral, tok.Literal, tok.Ln, tok.Col)
		}
		// check line number
		if val.ln != tok.Ln {
			t.Errorf("Tests[%d] - wrong line number. Expected: %d, got: %d", i, val.ln, tok.Ln)
		}
		// check column number
		if val.col != tok.Col {
			t.Errorf("Tests[%d] - wrong column number. Expected: %d, got: %d", i, val.col, tok.Col)
		}
	}
}