package main

import (
	"fmt"
	"strings"
)

type ExpressionVisitor interface {
	VisitDoubleExpression(e *DoubleExpression)
	VisitAdditionExpression(e *AdditionExpression)
}

type Expression interface {
	Accept(ev ExpressionVisitor)
}

type DoubleExpression struct {
	value float64
}

func (d *DoubleExpression) Accept(ev ExpressionVisitor) {
	ev.VisitDoubleExpression(d)
}

type AdditionExpression struct {
	left, right Expression
}

func (a *AdditionExpression) Accept(ev ExpressionVisitor) {
	ev.VisitAdditionExpression(a)
}

type ExpressionPrinter struct {
	sb strings.Builder
}

func NewExpressionPrinter() *ExpressionPrinter {
	return &ExpressionPrinter{strings.Builder{}}
}

func (ep *ExpressionPrinter) VisitDoubleExpression(e *DoubleExpression) {
	ep.sb.WriteString(fmt.Sprintf("%g", e.value))
}

func (ep *ExpressionPrinter) VisitAdditionExpression(e *AdditionExpression) {
	ep.sb.WriteRune('(')
	e.left.Accept(ep)
	ep.sb.WriteRune('+')
	e.right.Accept(ep)
	ep.sb.WriteRune(')')
}

func (ep *ExpressionPrinter) String() string {
	return ep.sb.String()
}

type ExpressionEvaluator struct {
	result float64
}

func (ee *ExpressionEvaluator) VisitDoubleExpression(e *DoubleExpression) {
	ee.result = e.value
}

func (ee *ExpressionEvaluator) VisitAdditionExpression(e *AdditionExpression) {
	e.left.Accept(ee)
	x := ee.result

	e.right.Accept(ee)
	x += ee.result

	ee.result = x
}

// 1 + (2+3)
func main() {
	e := &AdditionExpression{
		left: &DoubleExpression{1},
		right: &AdditionExpression{
			left:  &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}

	ep := NewExpressionPrinter()
	e.Accept(ep)

	ee := &ExpressionEvaluator{}
	e.Accept(ee)

	fmt.Println(fmt.Sprintf("%s = %g", ep.String(), ee.result))
}
