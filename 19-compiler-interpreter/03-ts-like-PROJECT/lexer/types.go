package lexer

type TokenType int

const (
	Number TokenType = iota
	Identifier
	Equals
	OpenParen
	CloseParen
	BinaryOperator
	Let
	EOF
)

type Token struct {
	Value string
	Type  TokenType
}

func NewToken(v string, t TokenType) Token {
	return Token{Value: v, Type: t}
}

type Lexer struct {
	Src string
	Pos int
	Len int
	Tok []Token
}

func NewLexer(src string) Lexer {
	return Lexer{Src: src, Len: len(src), Pos: 0, Tok: make([]Token, 0)}
}

var KEYWORDS = map[string]TokenType{"let": Let}
