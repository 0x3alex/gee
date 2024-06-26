package internal

import (
	"fmt"
	"math"

	"github.com/0x3alex/gee/internal/parser"
	"github.com/0x3alex/gee/internal/tokens"
)

// synonym to write less
type tok tokens.TokenInterface[any]

func isTokNum(a tok) bool {
	return a.GetType() == tokens.TokInt || a.GetType() == tokens.TokFloat
}

func isTokBool(a tok) bool {
	return a.GetType() == tokens.TokTrue || a.GetType() == tokens.TokFalse
}

func isTokStr(a tok) bool {
	return a.GetType() == tokens.TokStr
}

func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func mul(a, b float64) float64 {
	return a * b
}

func div(a, b float64) float64 {
	return a / b
}

func gt(a, b float64) bool {
	return a > b
}

func gteq(a, b float64) bool {
	return a >= b
}

func lt(a, b float64) bool {
	return a < b
}

func lteq(a, b float64) bool {
	return a <= b
}

func eq[k comparable](a, b k) bool {
	return a == b
}

func neq[k comparable](a, b k) bool {
	return a != b
}

func and(a, b bool) bool {
	return a && b
}

func or(a, b bool) bool {
	return a || b
}

func pow(a, b float64) float64 {
	return math.Pow(a, b)
}

func boolToBoolOp(a, b tok, f func(a, b bool) bool, negate bool) (tok, error) {
	if !isTokBool(a) || !isTokBool(b) {
		return nil, fmt.Errorf("bool comparisons can only be done between bools. left was %s, right was %s", a.ToString(), b.ToString())
	}
	aV, bV := a.GetValue().(bool), b.GetValue().(bool)
	res := f(aV, bV)
	if res {
		if negate {
			return tokens.NewFalse(), nil
		}
		return tokens.NewTrue(), nil
	}
	if negate {
		return tokens.NewTrue(), nil
	}
	return tokens.NewFalse(), nil
}

/*
matchXX are helper functions to reduce redunancy in the code
it converts to the needed type for the operation and executes the function.
if the result needs to be negated, its handled here aswell
*/
func mathToBoolOp(a, b tok, f func(a, b float64) bool, negate bool) (tok, error) {
	if !isTokNum(a) || !isTokNum(b) {
		return nil, fmt.Errorf("comparisons can only be done between numbers. left was %s right was %s", a.ToString(), b.ToString())
	}
	var aV, bV float64
	if a.GetType() == tokens.TokInt {
		aV = float64(a.GetValue().(int))
	} else {
		aV = a.GetValue().(float64)
	}
	if b.GetType() == tokens.TokInt {
		bV = float64(b.GetValue().(int))
	} else {
		bV = b.GetValue().(float64)
	}
	res := f(aV, bV)

	if res {
		if negate {
			return tokens.NewFalse(), nil
		}
		return tokens.NewTrue(), nil
	}
	if negate {
		return tokens.NewTrue(), nil
	}
	return tokens.NewFalse(), nil
}

func strToBoolOp(a, b tok, f func(a, b string) bool, negate bool) (tok, error) {
	if !isTokStr(a) || !isTokStr(b) {
		return nil, fmt.Errorf("strings can only be compared with strings")
	}
	aV, bV := a.GetValue().(string), b.GetValue().(string)
	res := f(aV, bV)

	if res {
		if negate {
			return tokens.NewFalse(), nil
		}
		return tokens.NewTrue(), nil
	}
	if negate {
		return tokens.NewTrue(), nil
	}
	return tokens.NewFalse(), nil

}

func mathToMathOp(a, b tok, f func(a, b float64) float64) (tok, error) {
	if !isTokNum(a) || !isTokNum(b) {
		return nil, fmt.Errorf("math operation can only be done between two numbers")
	}
	var aV, bV float64
	if a.GetType() == tokens.TokInt {
		aV = float64(a.GetValue().(int))
	} else {
		aV = a.GetValue().(float64)
	}
	if b.GetType() == tokens.TokInt {
		bV = float64(b.GetValue().(int))
	} else {
		bV = b.GetValue().(float64)
	}
	res := f(aV, bV)
	return tokens.ExistingFloat(res), nil

}

/*
from bottom up, traverse to the leaves and then eval
left -> right -> root
*/
func EvalAST(n *parser.Node) (tok, error) {
	if n.Left == nil && n.Right == nil {
		return n.T, nil
	}
	//var fn func(a, b float64) float64
	left, lerr := EvalAST(n.Left)
	if lerr != nil {
		return nil, lerr
	}
	operator := n.T.GetType()
	right, rerr := EvalAST(n.Right)
	if rerr != nil {
		return nil, rerr
	}
	//match the operator
	switch operator {
	case tokens.TokAdd:
		t, err := mathToMathOp(left, right, add)
		return t, err
	case tokens.TokSub:
		t, err := mathToMathOp(left, right, sub)
		return t, err
	case tokens.TokMul:
		t, err := mathToMathOp(left, right, mul)
		return t, err
	case tokens.TokDiv:
		t, err := mathToMathOp(left, right, div)
		return t, err
	case tokens.TokPow:
		t, err := mathToMathOp(left, right, pow)
		return t, err
	case tokens.TokEq:
		var t tok
		var err error
		//special case for string
		if isTokStr(left) || isTokStr(right) {
			t, err = strToBoolOp(left, right, eq, n.Negate)
		} else {
			t, err = mathToBoolOp(left, right, eq, n.Negate)
		}
		return t, err
	case tokens.TokNeq:
		var t tok
		var err error
		//special case for string
		if isTokStr(left) || isTokStr(right) {
			t, err = strToBoolOp(left, right, neq, n.Negate)
		} else {
			t, err = mathToBoolOp(left, right, neq, n.Negate)
		}
		return t, err
	case tokens.TokGt:
		t, err := mathToBoolOp(left, right, gt, n.Negate)
		return t, err
	case tokens.TokGtEq:
		t, err := mathToBoolOp(left, right, gteq, n.Negate)
		return t, err
	case tokens.TokLt:
		t, err := mathToBoolOp(left, right, lt, n.Negate)
		return t, err
	case tokens.TokLtEq:
		t, err := mathToBoolOp(left, right, lteq, n.Negate)
		return t, err
	case tokens.TokAnd:
		t, err := boolToBoolOp(left, right, and, n.Negate)
		return t, err
	case tokens.TokOr:
		t, err := boolToBoolOp(left, right, or, n.Negate)
		return t, err
	case tokens.TokNot:
		return nil, fmt.Errorf("%s is not an operator", n.T.ToString())
	default:
		return nil, fmt.Errorf("%s is not an operator", n.T.ToString())
	}
}
