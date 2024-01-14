package main

import (
	"os"

	"github.io/lets-learn-it/go-learning/ts-like/lexer"
)

func main() {
	src, err := os.ReadFile("examples/1.txt")
	if err != nil {
		panic(err)
	}

	lex := lexer.NewLexer(string(src))
	lex.Tokenize()
	lex.Print()
}
