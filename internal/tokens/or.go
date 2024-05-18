package tokens

type Or struct {
	tokenT int
}

func NewOr() Or {
	return Or{tokenT: TokOr}
}

func (i Or) GetType() int {
	return i.tokenT
}

func (i Or) GetValue() any {
	return "||"
}

func (i Or) ToString() string {
	return "||"
}
