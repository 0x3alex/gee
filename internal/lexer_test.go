package internal

import (
	"fmt"
	"testing"
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
	//str := "'hi' == 'hi'" //"((1==1)&&(1.5+1>3)) || !(1 == 1)" //"1.5+(2/(4*5))-0.6"
	str := "2^0.5" //"((('hi' == 'hey') || (1.5+2 > 1)) || !(2>2))&&False"
	l := NewLexer(str)
	res, err := l.Lex()
	if err != nil {
		t.Fatal(err.Error())
	}
	for _, v := range res {
		print(v.ToString())
	}
	println()
	n, err := BuildAST(res)
	if err != nil {
		t.Fatal(err.Error())
	}
	//printInOrder(n, 0)
	println()
	println()
	eval, err := EvalAST(n)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Printf("evalAST(n).ToString(): %s\n", eval.ToString())
}
