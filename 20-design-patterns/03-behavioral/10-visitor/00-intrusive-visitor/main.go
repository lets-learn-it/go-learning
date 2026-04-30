package main

import (
	"fmt"
	"strings"
)

// Intrusive approach adds new methods/functionality to existing heirarchy.
// This violets open closed principle & single responsibility principle.
// THIS VISITOR SHOULD NOT BE USED.

// (1+2)+3
type Expression interface {
	// intrusive (method added)
	Print(sb *strings.Builder)
}

type DoubleExpression struct {
	value float64
}

func (de *DoubleExpression) Print(sb *strings.Builder) {
	sb.WriteString(fmt.Sprintf("%g", de.value))
}

type AdditionExpression struct {
	left, right Expression
}

func (ae *AdditionExpression) Print(sb *strings.Builder) {
	sb.WriteRune('(')
	ae.left.Print(sb)
	sb.WriteRune('+')
	ae.right.Print(sb)
	sb.WriteRune(')')
}

func main() {
	// 1 + (2 + 3)
	e := &AdditionExpression{
		left: &DoubleExpression{1},
		right: &AdditionExpression{
			left:  &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}

	// This string builder is visitor
	// as it is passed to all parts of hierarchy.
	sb := strings.Builder{}

	e.Print(&sb)

	fmt.Println(sb.String())
}
