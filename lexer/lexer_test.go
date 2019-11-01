package lexer

import (
	"testing"
	"willard/token"
)

func TestNextToken(t *testing.T) {
	input := `
	|x> := |0>
	|y> := |0>
	H(|x>)
	H(|y>)
	R(45)(|y>)
	H(|y)
	CNOT(|xy>)

	let x = |x>
	let y = |y>

	print(x)
	print(y)
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong expected = %q, got = %q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong, expected = %q, got = %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
