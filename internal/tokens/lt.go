package tokens

type Lt struct {
	tokenT int
}

func NewLt() Lt {
	return Lt{tokenT: TokLt}
}

func (i Lt) GetType() int {
	return i.tokenT
}

func (i Lt) GetValue() any {
	return "<"
}

func (i Lt) ToString() string {
	return "<"
}
