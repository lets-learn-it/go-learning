package main

import (
	"fmt"
	"strings"
)

type ExpressionVisitor interface {
	VisitDoubleExpression(e *DoubleExpression)
	VisitAdditionExpression(e *AdditionExpression)
}

// (1+2)+3
type Expression interface {
	Accept(ev ExpressionVisitor)
}

type DoubleExpression struct {
	value float64
}

func (de *DoubleExpression) Accept(ev ExpressionVisitor) {
	ev.VisitDoubleExpression(de)
}

type AdditionExpression struct {
	left, right Expression
}

func (ae *AdditionExpression) Accept(ev ExpressionVisitor) {
	ev.VisitAdditionExpression(ae)
}

// Visitor implementation
type ExpressionPrinter struct {
	sb *strings.Builder
}

func (ep *ExpressionPrinter) VisitAdditionExpression(ae *AdditionExpression) {
	ep.sb.WriteRune('(')
	ae.left.Accept(ep)
	ep.sb.WriteRune('+')
	ae.right.Accept(ep)
	ep.sb.WriteRune(')')
}

func (ep *ExpressionPrinter) VisitDoubleExpression(de *DoubleExpression) {
	ep.sb.WriteString(fmt.Sprintf("%g", de.value))
}

// Visitor Implementation
type ExpressionEvaluator struct {
	value float64
}

func (ee *ExpressionEvaluator) VisitAdditionExpression(ae *AdditionExpression) {
	ae.left.Accept(ee)
	l := ee.value

	ae.right.Accept(ee)
	r := ee.value

	ee.value = l + r
}

func (ee *ExpressionEvaluator) VisitDoubleExpression(de *DoubleExpression) {
	ee.value = de.value
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

	ep := ExpressionPrinter{sb: &sb}
	e.Accept(&ep)

	fmt.Println(ep.sb.String())

	ee := ExpressionEvaluator{value: 0}
	e.Accept(&ee)

	fmt.Printf("%s = %f\n", ep.sb.String(), ee.value)

}
