package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/pskp-95/antlr-expr/parser"
)

func print_tokens(fis *antlr.FileStream) {
	lexer := parser.NewExprLexer(fis)

	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}

		fmt.Printf("%s (%q)\n", lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}

func interpreter(fis *antlr.FileStream) {
	memory := make(map[string]int)
	lexer := parser.NewExprLexer(fis)
	stream := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	p := parser.NewExprParser(stream)
	p.BuildParseTrees = true
	tree := p.Prog()

	visitor := NewExprVisitor(memory)

	result := visitor.Visit(tree)

	fmt.Println(result)
}

func interpreter_repl() {
	memory := make(map[string]int)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Print(input)

		is := antlr.NewInputStream(input)
		lexer := parser.NewExprLexer(is)
		stream := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
		p := parser.NewExprParser(stream)
		p.BuildParseTrees = true
		tree := p.Prog()

		visitor := NewExprVisitor(memory)

		_ = visitor.Visit(tree)
	}
}

func main() {
	// fis, err := antlr.NewFileStream("./input.txt")
	// if err != nil {
	// 	panic(err.Error())
	// }

	// print_tokens(fis)
	// interpreter(fis)
	interpreter_repl()
}
