package day08

import "testing"

//func TestXXXXX(t *testing.T) {
//	t.Log("hello")
//}

func TestFuncBasic(t *testing.T) {
	sumNum := Sum(1, 2, 3, 4, 5)
	t.Log(sumNum)
	t.Log(Sum(1, 2, 3, 4, 5))

	ret := intSum(3, 9)
	t.Log(ret)

	t.Log(intSum2(3, 4))
}

/**
求和
*/
func Sum(ops ...int) int {
	s := 0
	for _, op := range ops {
		s += op
	}
	return s
}

/**
add
*/
func intSum(x int, y int) int {

	return x + y
}

func intSum2(x, y int) int {
	return x + y
}
