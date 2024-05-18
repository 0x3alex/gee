package tokens

import (
	"fmt"
	"strconv"
)

type Int struct {
	value  int
	tokenT int
}

func NewInt(value string) (bool, Int) {
	if v, err := strconv.Atoi(value); err != nil {
		return false, Int{}
	} else {
		return true, Int{
			value:  v,
			tokenT: TokInt,
		}
	}
}

func ExistingInt(i int) Int {
	return Int{
		value:  i,
		tokenT: TokInt,
	}
}

func (i Int) GetType() int {
	return i.tokenT
}

func (i Int) GetValue() any {
	return i.value
}

func (i Int) ToString() string {
	return fmt.Sprintf("%d", i.value)
}
