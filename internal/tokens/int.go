package tokens

import (
	"fmt"
	"strconv"
)

type Int struct {
	value  int
	tokenT int
}

func NewInt(value string) (bool, Float) {
	if v, err := strconv.Atoi(value); err != nil {
		return false, Float{}
	} else {
		return true, Float{
			value:  float64(v),
			tokenT: TokFloat,
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
