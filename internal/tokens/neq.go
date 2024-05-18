package tokens

type NEq struct {
	tokenT int
}

func NewNEq() NEq {
	return NEq{tokenT: TokNeq}
}

func (i NEq) GetType() int {
	return i.tokenT
}

func (i NEq) GetValue() any {
	return "!="
}

func (i NEq) ToString() string {
	return "!="
}
