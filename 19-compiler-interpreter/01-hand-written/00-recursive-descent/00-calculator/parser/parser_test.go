package parser

import (
	"fmt"
	"testing"

	"github.com/pskp-95/calc/lexer/lexer"
)

func TestParser(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"1+2+3", 6},
		{"12+234+12-13-23", 222},
		{"(3 * 6 / 9) * 5", 10},
		{"(2 * 1)", 2},
		{"(6 - 7 / 4) + 8", 13},
		{"(((3 / 9) - 2) / 5) * 9", 0},
		{"((2 - 8 / 8) + 6)", 7},
		{"4 - 1 * 7", -3},
		{"(((8 - 8) * 1) * 6 - 6)", -6},
	}

	for i, test := range tests {
		l := lexer.New(test.input)
		p := New(l)

		o := p.Run()

		if test.output != o {
			t.Fatalf("tests[%d] - expected output=%d, got=%d", i, test.output, o)
		}
		fmt.Printf("tests[%d] - expected output=%d, got=%d\n", i, test.output, o)
	}
}
