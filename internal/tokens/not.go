package tokens

type Not struct {
	tokenT int
}

func NewNot() Not {
	return Not{tokenT: TokNot}
}

func (i Not) GetType() int {
	return i.tokenT
}

func (i Not) GetValue() any {
	return "!"
}

func (i Not) ToString() string {
	return "!"
}
