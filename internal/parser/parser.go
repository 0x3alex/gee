package parser

import (
	"fmt"

	"github.com/0x3alex/gee/internal/tokens"
)

type Node struct {
	T           tokens.TokenInterface[any]
	Negate      bool
	Left, Right *Node
}

func tokenToNode(t tokens.TokenInterface[any]) *Node {
	return &Node{
		T: t,
	}
}

func mergeNodes(nodes []*Node) *Node {
	return &Node{
		T:      nodes[1].T,
		Negate: nodes[1].Negate,
		Left:   nodes[0],
		Right:  nodes[2],
	}
}

func BuildAST(toks []tokens.TokenInterface[any]) (*Node, error) {
	_, v, err := _buildAST(0, toks)
	return v, err
}

// check if braces are valid
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

func _buildAST(i int, toks []tokens.TokenInterface[any]) (int, *Node, error) {
	//exit if braces do not match
	if !validateBraces(toks) {
		return 0, nil, fmt.Errorf("(,) mismatch")
	}
	var accum []*Node
	idx := i
	//this bool keeps track of single !
	var neg bool
	for idx < len(toks) {
		current := toks[idx]
		//handle the !, we dont need to append it. if we already had a ! and
		//did not use it, error out because !..! or !!,.. is not valid
		if current.GetType() == tokens.TokNot {
			if neg {
				return idx, nil, fmt.Errorf("!! is not a valid expression")
			}
			neg = true
			idx++
			continue
		}
		//recursivly handle braces
		if current.GetType() == tokens.TokOpen {
			j, n, err := _buildAST(idx+1, toks)
			if err != nil {
				return idx, nil, err
			}
			n.Negate = neg
			neg = false
			idx = j
			accum = append(accum, n)
		} else if current.GetType() == tokens.TokClose {
			//if we close and do not have 1 ore 3 elements, then the expression is not valid
			if len(accum) == 2 || len(accum) > 3 || len(accum) == 0 {
				return idx, nil, fmt.Errorf("expression can only contain 1 or 3 elements")
			}
			//if we have an expression pending, merge it
			if len(accum) == 3 {
				n := mergeNodes(accum)
				accum = accum[:0]
				accum = append(accum, n)
			}
			return idx + 1, accum[0], nil
		} else {
			//if we have a ! (TokNot) and its not followed by a (, then its wrong
			if neg {
				return idx, nil,
					fmt.Errorf("! must be followed by = or (, but was followed by %s", current.GetValue())
			}
			accum = append(accum, tokenToNode(current))
			idx++
		}
		//if we have an expression pending, merge it
		if len(accum) == 3 {
			n := mergeNodes(accum)
			accum = accum[:0]
			accum = append(accum, n)

		}
	}
	//if we close and do not have 1 ore 3 elements, then the expression is not valid
	if len(accum) == 2 || len(accum) > 3 || len(accum) == 0 {
		return 0, nil, fmt.Errorf("expression can only contain 1 or 3 elements")
	}
	//if we have an expression pending, merge it
	if len(accum) == 3 {
		n := mergeNodes(accum)
		accum = accum[:0]
		accum = append(accum, n)
	}
	return 0, accum[0], nil
}
