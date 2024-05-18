package tokens

type True struct {
	tokenT int
}

func NewTrue() True {
	return True{tokenT: TokTrue}
}

func (i True) GetType() int {
	return i.tokenT
}

func (i True) GetValue() any {
	return true
}

func (i True) ToString() string {
	return "True"
}
