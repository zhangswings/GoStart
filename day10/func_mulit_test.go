package day10

import "testing"

func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

func TestFuncMulitValue(t *testing.T) {
	sumValue, subValue := calc(10, 2)
	t.Logf("calc sumValue=%d subValue=%d ", sumValue, subValue)
	t.Logf("calc sumValue=%d subValue=%d ", sumValue, subValue)
}
