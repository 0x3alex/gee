package tokens

type Str struct {
	value  string
	tokenT int
}

func NewStr(value string) Str {
	return Str{
		value:  value,
		tokenT: TokStr,
	}
}

func (i Str) GetType() int {
	return i.tokenT
}

func (i Str) GetValue() any {
	return i.value
}

func (i Str) ToString() string {
	return i.value
}
