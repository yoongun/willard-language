package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifiers + literals
	IDENT = "IDENT"
	EQUAL = "="
	PLUS  = "+"

	// Delimiter
	COMMA     = ","
	COLON     = ":"
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"

	LBRACE = "{"
	RBRACE = "}"

	// keywords
)
