package tokens

import (
	"fmt"
	"strconv"
)

type Float struct {
	value  float64
	tokenT int
}

func NewFloat(value string) (bool, Float) {
	if v, err := strconv.ParseFloat(value, 32); err != nil {
		return false, Float{}
	} else {
		return true, Float{
			value:  v,
			tokenT: TokFloat,
		}
	}
}

func ExistingFloat(f float64) Float {
	return Float{
		value:  f,
		tokenT: TokFloat,
	}
}

func (i Float) GetType() int {
	return i.tokenT
}

func (i Float) GetValue() any {
	return i.value
}
func (i Float) ToString() string {
	return fmt.Sprintf("%.2f", i.value)
}
