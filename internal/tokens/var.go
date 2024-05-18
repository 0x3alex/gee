package tokens

type Var struct {
	value  string
	tokenT int
}

func NewVar(value string) Var {
	return Var{
		value:  value,
		tokenT: TokVar,
	}
}

func (i Var) GetType() int {
	return i.tokenT
}

func (i Var) GetValue() any {
	return i.value
}

func (i Var) ToString() string {
	return i.value
}
