package tokens

type Close struct {
	tokenT int
}

func NewClose() Close {
	return Close{tokenT: TokClose}
}

func (i Close) GetType() int {
	return i.tokenT
}

func (i Close) GetValue() any {
	return "("
}

func (i Close) ToString() string {
	return ")"
}
