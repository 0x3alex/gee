package gee

import (
	"fmt"

	"github.com/0x3alex/gee/internal"
	"github.com/0x3alex/gee/internal/tokens"
)

/*
@int - 0, if it is a number, 1 if it is a bool, 2 if it is a string

@any - the value

@error - is set to an error if an error occured
*/
func Eval(s string) (int, any, error) {
	l := internal.NewLexer(s)
	res, err := l.Lex()
	if err != nil {
		return 0, nil, err
	}
	n, err := internal.BuildAST(res)
	if err != nil {
		return 0, nil, err
	}
	if n == nil {
		return 0, nil, fmt.Errorf("ast returned nil")
	}
	result, err := internal.EvalAST(n)
	if err != nil {
		return 0, nil, err
	}
	var t int
	if result.GetType() == tokens.TokTrue || result.GetType() == tokens.TokFalse {
		t = 1
	} else if result.GetType() == tokens.TokStr {
		t = 2
	}
	return t, result.GetValue(), nil
}
