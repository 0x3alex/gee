package tokens

type Mul struct {
	tokenT int
}

func NewMul() Mul {
	return Mul{tokenT: TokMul}
}

func (i Mul) GetType() int {
	return i.tokenT
}

func (i Mul) GetValue() any {
	return "*"
}

func (i Mul) ToString() string {
	return "*"
}
