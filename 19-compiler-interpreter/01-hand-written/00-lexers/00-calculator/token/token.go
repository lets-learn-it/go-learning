package token

type TokenType string

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	INT TokenType = "INT"

	PLUS  TokenType = "PLUS"
	MINUS TokenType = "MINUS"
	MUL   TokenType = "MUL"
	DIV   TokenType = "DIV"

	LPAREN TokenType = "LPAREN"
	RPAREN TokenType = "RPAREN"
)

type Token struct {
	Type    TokenType
	Literal string
}
