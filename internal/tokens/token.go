package tokens

type TokenInterface[T any] interface {
	GetType() int
	GetValue() T
	ToString() string
}

const (
	TokAnd = iota
	TokTrue
	TokFalse
	TokOr
	TokEq
	TokNeq
	TokNot
	TokGt
	TokGtEq
	TokLt
	TokLtEq
	TokOpen
	TokClose
	TokInt
	TokFloat
	TokAdd
	TokSub
	TokMul
	TokDiv
	TokStr
	TokVar
	TokPow
)
