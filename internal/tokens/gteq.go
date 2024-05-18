package tokens

type GtEq struct {
	tokenT int
}

func NewGtEq() GtEq {
	return GtEq{tokenT: TokGtEq}
}

func (i GtEq) GetType() int {
	return i.tokenT
}

func (i GtEq) GetValue() any {
	return ">="
}

func (i GtEq) ToString() string {
	return ">="
}
