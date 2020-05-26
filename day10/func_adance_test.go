package day10

import (
	"fmt"
	"testing"
)

//定义全局变量
var num int64 = 10

func testGlobalVar() {
	//函数中可以访问全局变量num
	fmt.Printf("num=%d\n", num)
}

func TestGlobalVar(t *testing.T) {
	testGlobalVar()
}

func testLocalVar() {
	//定义一个函数局部变量x，仅在该函数内生效
	var x int64 = 100

	fmt.Printf("x=%d\n", x)
}

func TestLocalVar(t *testing.T) {
	testLocalVar()
}

//如果局部变量和全局变量重名，优先访问局部变量
func testNum() {
	num := 100
	fmt.Printf("num=%d\n", num) // 函数中优先使用局部变量
}

func TestGlobalLocalVar(t *testing.T) {
	testNum()
}

func testLocalVar2(x, y int) {
	fmt.Println(x, y)
	if x > 0 {
		z := 100 //变量z只在if语句块生效
		fmt.Println(z)
	}
	//fmt.Println(z)//此处无法使用变量z
}

func testLocalVar3() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func TestGlobalLocalVar2(t *testing.T) {
	testLocalVar2(1, 2)
	testLocalVar3()
}

/**
使用type关键字来定义一个函数类型
*/
type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func TestTypeFuncMethod(t *testing.T) {

	var c calculation
	c = add
	t.Logf("%T", c) //day10.calculation
	t.Log(c(1, 2))

	f := sub
	t.Logf("%T", f) //func(int, int) int
	t.Log(f(2, 1))
}
