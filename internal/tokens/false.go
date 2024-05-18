package tokens

type False struct {
	tokenT int
}

func NewFalse() False {
	return False{tokenT: TokFalse}
}

func (i False) GetType() int {
	return i.tokenT
}

func (i False) GetValue() any {
	return false
}

func (i False) ToString() string {
	return "False"
}
