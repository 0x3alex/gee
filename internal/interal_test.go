package internal

import (
	"testing"

	"github.com/0x3alex/gee/internal/lexer"
	"github.com/0x3alex/gee/internal/parser"
	"github.com/0x3alex/gee/internal/tokens"
)

func printInOrder(n *parser.Node, lvl int) {
	if n == nil {
		return
	}
	printInOrder(n.Left, lvl+1)
	for i := 0; i < lvl; i++ {
		print(" ")
	}
	print(n.T.ToString())
	println()
	printInOrder(n.Right, lvl+1)
}

func TestLexer(t *testing.T) {
	str := "((4.5 / 2) + 3)== 10"
	l := lexer.New(str)
	res, err := l.Lex()
	if err != nil {
		t.Fatal(err.Error())
	}
	if len(res) != 11 {
		t.Fatalf("Expected 10 but got %d", len(res))
	}
	n, err := parser.BuildAST(res)
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
