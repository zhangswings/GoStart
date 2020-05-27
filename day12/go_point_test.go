package day12

import (
	"fmt"
	"testing"
)

/**
取变量指针的语法如下：
ptr := &v    // v的类型为T
v:代表被取地址的变量，类型为T
ptr:用于接收地址的变量，ptr的类型就为*T，称做T的指针类型。*代表指针。
*/
func TestPointFunc(t *testing.T) {
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a)
	fmt.Printf("a:%d type:%T\n", a, a)
	fmt.Printf("b:%p type:%T\n", b, b)
}

func TestPointFunc2(t *testing.T) {
	//指针取值
	a := 10
	b := &a //取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", b)

	c := *b
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)
	//总结： 取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。
}

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

//变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：
//
//对变量进行取地址（&）操作，可以获得这个变量的指针变量。
//指针变量的值是指针地址。
//对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。
func TestModify(t *testing.T) {
	//var a *int
	//*a=100
	//fmt.Println(&a)
	a := 10
	modify1(a)
	fmt.Println(a)

	modify2(&a)
	fmt.Println(a)

}
