package day10

import (
	"errors"
	"fmt"
	"strings"
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

/**
func(参数)(返回值){
    函数体
}
*/
func TestCloseFunc(t *testing.T) {
	//将匿名函数保存到变量
	add2 := func(x, y int) {
		fmt.Println(x + y)
	}
	add2(12, 20)

	//子执行函数：匿名函数定义加完（）直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(12, 13)
}

//闭包
//闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境。
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func TestAdder(t *testing.T) {
	var f = adder()
	fmt.Println(f(10)) //10
	fmt.Println(f(20)) //30
	fmt.Println(f(30)) //60

	f2 := adder()
	fmt.Println(f2(40)) //40
	fmt.Println(f2(50)) //90
	//变量f是一个函数并且它引用了其外部作用域中的x变量，此时f就是一个闭包。 在f的生命周期内，变量x也一直有效。
}

//闭包
//闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境。
func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func TestAdder2(t *testing.T) {
	var f = adder2(10)
	fmt.Println(f(10)) //20
	fmt.Println(f(20)) //40
	fmt.Println(f(30)) //70

	f2 := adder2(10)
	fmt.Println(f2(40)) //50
	fmt.Println(f2(50)) //100
	//变量f是一个函数并且它引用了其外部作用域中的x变量，此时f就是一个闭包。 在f的生命周期内，变量x也一直有效。
}

func makeSuffixFunc(suffix string) func(string) string {

	return func(name string) string {
		if !strings.HasPrefix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func TestMakeSuffixFunc(t *testing.T) {
	jpgFunc := makeSuffixFunc(".jpg")

	fmt.Println(jpgFunc("test")) //test.jpg
}

func calcPlus(base int) (func(int) int, func(int) int) {

	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}

	return add, sub
}

/**
闭包=函数+引用环境
*/
func TestCalcPlus(t *testing.T) {
	f1, f2 := calcPlus(10)
	t.Log(f1(1), f2(2)) //11 9
	t.Log(f1(3), f2(4)) //12 8
	t.Log(f1(5), f2(6)) //13 7
}
