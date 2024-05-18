package tokens

type Gt struct {
	tokenT int
}

func NewGt() Gt {
	return Gt{tokenT: TokGt}
}

func (i Gt) GetType() int {
	return i.tokenT
}

func (i Gt) GetValue() any {
	return ">"
}

func (i Gt) ToString() string {
	return ">"
}
