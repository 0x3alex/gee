package tokens

type Eq struct {
	tokenT int
}

func NewEq() Eq {
	return Eq{tokenT: TokEq}
}

func (i Eq) GetType() int {
	return i.tokenT
}

func (i Eq) GetValue() any {
	return "=="
}

func (i Eq) ToString() string {
	return "=="
}
