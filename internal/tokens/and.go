package tokens

type And struct {
	tokenT int
}

func NewAnd() And {
	return And{tokenT: TokAnd}
}

func (i And) GetType() int {
	return i.tokenT
}

func (i And) GetValue() any {
	return "&&"
}

func (i And) ToString() string {
	return "&&"
}
