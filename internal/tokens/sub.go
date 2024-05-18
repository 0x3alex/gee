package tokens

type Sub struct {
	tokenT int
}

func NewSub() Sub {
	return Sub{tokenT: TokSub}
}

func (i Sub) GetType() int {
	return i.tokenT
}

func (i Sub) GetValue() any {
	return "-"
}

func (i Sub) ToString() string {
	return "-"
}
