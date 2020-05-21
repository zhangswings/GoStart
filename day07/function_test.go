package day07

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/**
Golang 函数与其他主要编程语言的差异
1. 可以有多个返回值
2. 所有参数都是值传递：slice map channel 会有传引用的错觉
3. 函数可以作为变量的值
4. 函数可以作为参数和返回值

*/
func TestBasicFn(t *testing.T) {

	//只有下划线_ 忽略返回值
	a, _ := returnMultiValues()
	t.Log(a)

	tsSF := timeSpent(SlowFunc)
	t.Log(tsSF(10))
}

//func TestNimoFn(t *testing.T)  {
//	t.Log("hello")
//
//}

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {

	return func(op int) int {
		start := time.Now()
		ret := inner(op)

		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func SlowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}
