package tokens

type Div struct {
	tokenT int
}

func NewDiv() Div {
	return Div{tokenT: TokDiv}
}

func (i Div) GetType() int {
	return i.tokenT
}

func (i Div) GetValue() any {
	return "/"
}

func (i Div) ToString() string {
	return "/"
}
