package main

import (
	"fmt"

	"github.com/pskp-95/calc/lexer/lexer"
	"github.com/pskp-95/calc/parser/parser"
)

func main() {
	// check if lexer is available for use & also working
	// l := lexer.New("56+95")
	// for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
	// 	fmt.Printf("%+v\n", tok)
	// }

	ll := lexer.New("(3 * 6 / 9) * 5")
	p := parser.New(ll)
	fmt.Println(p.Run())
}
