package tokens

type LtEq struct {
	tokenT int
}

func NewLtEq() LtEq {
	return LtEq{tokenT: TokLtEq}
}

func (i LtEq) GetType() int {
	return i.tokenT
}

func (i LtEq) GetValue() any {
	return "<="
}

func (i LtEq) ToString() string {
	return "<="
}
