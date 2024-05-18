package tokens

type Add struct {
	tokenT int
}

func NewAdd() Add {
	return Add{tokenT: TokAdd}
}

func (i Add) GetType() int {
	return i.tokenT
}

func (i Add) GetValue() any {
	return "+"
}

func (i Add) ToString() string {
	return "+"
}
