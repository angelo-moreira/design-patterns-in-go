package reflective_visitor

import (
	"fmt"
	"strings"
)

type Expression interface{}

type DoubleExpression struct {
	value float64
}

type AdditionExpression struct {
	left, right Expression
}

func Print(e Expression, sb *strings.Builder) {
	if de, ok := e.(*DoubleExpression); ok {
		sb.WriteString(fmt.Sprintf("%g", de.value))
	} else if ae, ok := e.(*AdditionExpression); ok {
		sb.WriteRune('(')
		Print(ae.left, sb)
		sb.WriteRune('+')
		Print(ae.right, sb)
		sb.WriteRune(')')
	}
	// this is not good because it violates the open/closed principle
	// because once you add a new expression you'll have to modify this code
}

// 1 + (2+3)
func Run() string {
	e := &AdditionExpression{
		left: &DoubleExpression{1},
		right: &AdditionExpression{
			left:  &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}

	sb := strings.Builder{}
	Print(e, &sb)

	return sb.String()
}
