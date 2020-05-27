package day11

import (
	"fmt"
	"testing"
)

//Go语言中的defer语句会将其后面跟随的语句进行延迟处理。
//在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行，
//也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行。
func TestDeferFunc(t *testing.T) {
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}

//=== RUN   TestDeferFunc
//start
//end
//3
//2
//1
//--- PASS: TestDeferFunc (0.00s)
//PASS

func f1() int {
	x := 5
	defer func() {
		x++
	}()

	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func TestMultiDefer(t *testing.T) {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
