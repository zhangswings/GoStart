package day08

import "testing"

//func TestXXXXX(t *testing.T) {
//	t.Log("hello")
//}

func TestFuncBasic(t *testing.T) {
	sumNum := Sum(1, 2, 3, 4, 5)
	t.Log(sumNum)
	t.Log(Sum(1, 2, 3, 4, 5))
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
