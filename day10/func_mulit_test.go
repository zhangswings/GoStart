package day10

import "testing"

func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

func TestFuncMulitValue(t *testing.T) {
	sumValue, subValue := calc(10, 2)
	// t.Logf("calc sumValue=%d subValue=%d ", sumValue, subValue)
	t.Logf("calc sumValue=%d subValue=%d ", sumValue, subValue)
}

//函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过return关键字返回。
func calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

func TestFuncMulitValue2(t *testing.T) {
	sumValue, subValue := calc2(10, 2)
	t.Logf("calc sumValue=%d subValue=%d ", sumValue, subValue)
}

//当我们的一个函数返回值类型为slice时，nil可以看做是一个有效的slice，没必要显示返回一个长度为0的切片。
func someFunc(x string) []int {
	if x == "" {
		return nil // 没必要返回[]int{}
	}
}
