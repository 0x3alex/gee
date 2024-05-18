package gee

import (
	"fmt"

	"github.com/0x3alex/gee/internal"
	"github.com/0x3alex/gee/internal/tokens"
)

func Eval(s string) (bool, any, error) {
	l := internal.NewLexer(s)
	res, err := l.Lex()
	if err != nil {
		return false, nil, err
	}
	n, err := internal.BuildAST(res)
	if err != nil {
		return false, nil, err
	}
	if n == nil {
		return false, nil, fmt.Errorf("ast returned nil")
	}
	result, err := internal.EvalAST(n)
	if err != nil {
		return false, nil, err
	}
	var isBool bool
	if result.GetType() == tokens.TokTrue || result.GetType() == tokens.TokFalse {
		isBool = true
	}
	return isBool, result.GetValue(), nil
}
