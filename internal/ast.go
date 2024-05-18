package internal

import (
	"fmt"
	"log"

	"github.com/0x3alex/gee/internal/tokens"
)

type node struct {
	t           tokens.TokenInterface[any]
	negate      bool
	left, right *node
}

func tokenToNode(t tokens.TokenInterface[any]) *node {
	return &node{
		t: t,
	}
}

func mergeNodes(nodes []*node) *node {
	return &node{
		t:      nodes[1].t,
		negate: nodes[1].negate,
		left:   nodes[0],
		right:  nodes[2],
	}
}

func BuildAST(toks []tokens.TokenInterface[any]) (*node, error) {
	_, v, err := _buildAST(0, toks)
	return v, err
}

func validateBraces(toks []tokens.TokenInterface[any]) bool {
	var braces int
	for _, v := range toks {
		if v.GetType() == tokens.TokOpen {
			braces++
		} else if v.GetType() == tokens.TokClose {
			braces--
			if braces < 0 {
				return false
			}
		}
	}
	return braces == 0
}

func _buildAST(i int, toks []tokens.TokenInterface[any]) (int, *node, error) {
	if !validateBraces(toks) {
		return 0, nil, fmt.Errorf("(,) mismatch")
	}
	var accum []*node
	idx := i
	var neg bool
	for idx < len(toks) {
		current := toks[idx]
		if current.GetType() == tokens.TokNot {
			if neg {
				return idx, nil, fmt.Errorf("!! is not a valid expression")
			}
			neg = true
			idx++
			continue
		}
		if current.GetType() == tokens.TokOpen {
			j, n, err := _buildAST(idx+1, toks)
			if err != nil {
				return idx, nil, err
			}
			n.negate = neg
			neg = false
			idx = j
			accum = append(accum, n)
		} else if current.GetType() == tokens.TokClose {
			if len(accum) == 2 || len(accum) > 3 || len(accum) == 0 {

				return idx, nil, fmt.Errorf("expression can only contain 1 or 3 elements")
			}
			if len(accum) == 3 {
				//println("inside with idx", idx)
				n := mergeNodes(accum)
				accum = accum[:0]
				accum = append(accum, n)
			}
			return idx + 1, accum[0], nil
		} else {
			if neg {
				return idx, nil,
					fmt.Errorf("! must be followed by = or (, but was followed by %s", current.GetValue())
			}
			accum = append(accum, tokenToNode(current))
			idx++
		}

		if len(accum) == 3 {
			n := mergeNodes(accum)
			accum = accum[:0]
			accum = append(accum, n)

		}
	}
	if len(accum) == 2 || len(accum) > 3 || len(accum) == 0 {
		log.Fatal("Expression can only contain 1 or 3 elements")
	}
	if len(accum) == 3 {
		//println("inside with idx", idx)
		n := mergeNodes(accum)
		accum = accum[:0]
		accum = append(accum, n)
	}
	//println(len(accum))
	return 0, accum[0], nil
}
