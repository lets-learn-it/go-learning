package main

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/pskp-95/antlr-calc/parser"
)

func print_tokens(input string) {
	is := antlr.NewInputStream(input)

	lexer := parser.NewCalcLexer(is)

	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}

		fmt.Printf("%s (%q)\n", lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}

func calc(input string) int {
	is := antlr.NewInputStream(input)

	lexer := parser.NewCalcLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewCalcParser(stream)

	var listener calcListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start_())

	return listener.pop()
}

func main() {
	input := "1+2+54* 4"
	print_tokens(input)

	fmt.Println(calc(input))
}
