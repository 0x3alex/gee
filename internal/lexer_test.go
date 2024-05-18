package internal

import (
	"testing"

	"github.com/0x3alex/gee/internal/tokens"
)

func printInOrder(n *node, lvl int) {
	if n == nil {
		return
	}
	printInOrder(n.left, lvl+1)
	for i := 0; i < lvl; i++ {
		print(" ")
	}
	print(n.t.ToString())
	println()
	printInOrder(n.right, lvl+1)
}

func TestLexer(t *testing.T) {
	str := "((4.5 / 2) + 3)== 10"
	l := NewLexer(str)
	res, err := l.Lex()
	if err != nil {
		t.Fatal(err.Error())
	}
	if len(res) != 11 {
		t.Fatalf("Expected 10 but got %d", len(res))
	}
	n, err := BuildAST(res)
	if err != nil {
		t.Fatal(err.Error())
	}
	eval, err := EvalAST(n)
	if err != nil {
		t.Fatal(err.Error())
	}
	if eval.GetType() != tokens.TokFalse {
		t.Fatalf("Expected false but got %s", eval.ToString())
	}
}
