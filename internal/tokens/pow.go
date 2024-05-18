package tokens

type Pow struct {
	tokenT int
}

func NewPow() Pow {
	return Pow{tokenT: TokPow}
}

func (i Pow) GetType() int {
	return i.tokenT
}

func (i Pow) GetValue() any {
	return "^"
}

func (i Pow) ToString() string {
	return "^"
}
