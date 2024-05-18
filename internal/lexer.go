package internal

import (
	"fmt"
	"unicode"

	"github.com/0x3alex/gee/internal/tokens"
)

/*
@current - the current position in the text

@text - the text we want to lex
*/
type Lexer struct {
	current int
	text    string
}

// synonym to write less
type tok tokens.TokenInterface[any]

func NewLexer(text string) *Lexer {
	return &Lexer{
		current: 0,
		text:    text,
	}
}

func (l *Lexer) getNext() {
	//make sure we are not out of bounds
	if l.current >= len(l.text) {
		return
	}
	l.current++
}

func (l *Lexer) hasNext() bool {
	return l.current < len(l.text)
}

func isString(r rune) bool {
	return r == '\''
}

func isVar(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isNum(r rune) bool {
	return (r >= 48 && r <= 57)
}

func isBool(r rune) bool {
	return r == '&' || r == '|' || r == '!' || r == '>' || r == '<' || r == '='
}

func (l *Lexer) fetchString() (string, error) {
	var agg string
	l.getNext() //skip the first '
	for l.hasNext() {
		curr := rune(l.text[l.current])
		if isString(curr) { //second ' found
			return agg, nil
		}
		agg += string(curr)
		l.getNext()
	}
	return "", fmt.Errorf("no ' found")
}

/*
	The functions named fetchXXX proceed to take from the string until it the current rune
	is not a valid rune for the type.

	For bool it would proceed to take, if the text is ..==.., the first = and then as long as it
	is a valid sign for a bool operation. So it would return ==

*/

func (l *Lexer) fetchBool() string {
	var agg string
	for l.hasNext() {
		curr := rune(l.text[l.current])
		if !isBool(curr) {
			break
		}
		agg += string(curr)
		l.getNext()
	}
	return agg
}

func (l *Lexer) fetchVar() string {
	var agg string
	for l.hasNext() {
		curr := rune(l.text[l.current])
		if !isVar(curr) {
			return agg
		}
		agg += string(curr)
		l.getNext()
	}
	return agg
}

func (l *Lexer) fetchNum() (bool, string, error) {
	dot := false
	var agg string
	for l.hasNext() {
		curr := rune(l.text[l.current])
		if curr == '.' {
			if dot {
				return false, "", fmt.Errorf("one number can not contain two dots")
			}
			dot = true
		} else {
			if !isNum(curr) {
				break
			}
		}
		agg += string(curr)
		l.getNext()
	}
	return dot, agg, nil
}

/*
The functions named matchXXX take in the read string from fetch for datatypes that can be mapped
to multiple tokens.

True and False are considered Var before the machting because its just a sequence of letters.
*/
func matchVarToTok(s string) (tokens.TokenInterface[any], error) {
	switch s {
	case "True":
		return tokens.NewTrue(), nil
	case "False":
		return tokens.NewFalse(), nil
	default:
		return nil, fmt.Errorf("unsupported variable")
	}
}

func matchBoolToTok(s string) (tokens.TokenInterface[any], error) {
	switch s {
	case "==":
		return tokens.NewEq(), nil
	case "!=":
		return tokens.NewNEq(), nil
	case "&&":
		return tokens.NewAnd(), nil
	case "||":
		return tokens.NewOr(), nil
	case ">":
		return tokens.NewGt(), nil
	case "<":
		return tokens.NewLt(), nil
	case ">=":
		return tokens.NewGtEq(), nil
	case "<=":
		return tokens.NewLtEq(), nil
	case "!":
		return tokens.NewNot(), nil
	default:
		return nil, fmt.Errorf("operator %s not supported", s)
	}
}

func (l *Lexer) Lex() ([]tokens.TokenInterface[any], error) {
	var result []tokens.TokenInterface[any]
	for l.hasNext() {
		current := rune(l.text[l.current])
		//skip whitespaces
		if unicode.IsSpace(current) {
			l.getNext()
			continue
		}
		/*
			match the current rune. its either a single rune, which can be mapped by the first cases.
			if its a sequence we go into the default case an call the according fetchXX function and
			the according matchXX function (if needed)
		*/
		switch current {
		case '(':
			result = append(result, tokens.NewOpen())
		case ')':
			result = append(result, tokens.NewClose())
		case '+':
			result = append(result, tokens.NewAdd())
		case '-':
			result = append(result, tokens.NewSub())
		case '*':
			result = append(result, tokens.NewMul())
		case '/':
			result = append(result, tokens.NewDiv())
		case '^':
			result = append(result, tokens.NewPow())
		default:
			if isNum(current) {
				t, n, err := l.fetchNum()
				if err != nil {
					return nil, err
				}
				if t { //float
					if fok, f := tokens.NewFloat(n); fok {
						result = append(result, f)
					}
				} else {
					if fok, f := tokens.NewInt(n); fok {
						result = append(result, f)
					}
				}
				continue
			} else if isBool(current) {
				n := l.fetchBool()
				v, err := matchBoolToTok(n)
				if err != nil {
					return nil, err
				}
				result = append(result, v)
				continue
			} else if isString(current) {
				n, err := l.fetchString()
				if err != nil {
					return nil, err
				}
				result = append(result, tokens.NewStr(n))
			} else if isVar(current) {
				n := l.fetchVar()
				if v, err := matchVarToTok(n); err != nil {
					return nil, err
				} else {
					result = append(result, v)
				}
				continue
			} else {
				return nil, fmt.Errorf("unsupported rune %s", string(current))
			}
		}
		l.getNext()
	}
	return result, nil
}
