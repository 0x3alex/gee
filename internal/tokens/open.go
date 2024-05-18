package tokens

type Open struct {
	tokenT int
}

func NewOpen() Open {
	return Open{tokenT: TokOpen}
}

func (i Open) GetType() int {
	return i.tokenT
}

func (i Open) GetValue() any {
	return "("
}

func (i Open) ToString() string {
	return "("
}
