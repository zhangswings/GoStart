package day10

import (
	"errors"
	"testing"
)

func addNum(x, y int) int {
	return x + y
}
func subNum(x, y int) int {
	return x - y
}

func calcNum(x, y int, op func(int, int) int) int {
	return op(x, y)
}
func TestFuncOp(t *testing.T) {
	ret := calcNum(10, 20, addNum)
	t.Log(ret)
	t.Logf("%T", ret)

	f, err := do("+")
	t.Log(err)

	ret2 := f(1, 2)
	t.Log(ret2)
}

//函数作为返回值
func do(s string) (func(int, int) int, error) {

	switch s {
	case "+":
		return addNum, nil
	case "-":
		return subNum, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}
