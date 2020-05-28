package day12

import (
	"fmt"
	"testing"
)

//执行上面的代码会引发panic，为什么呢？ 在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，
//否则我们的值就没办法存储。而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。
//要分配内存，就引出来今天的new和make。 Go语言中new和make是内建的两个函数，主要用来分配内存。
func TestNewMake(t *testing.T) {
	var a *int
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b["A"] = 100
	fmt.Println(b)
}

//new 是一个内置的函数，它的签名如下：
// func new(Type) *Type
// Type 表示类型，new函数只接受一个参数，这个参数是一个类型。
// *Type 表示类型指针，new函数返回一个指向该类型内存地址的指针。

func TestNewFunc(t *testing.T) {
	a := new(int)
	*a = 100
	fmt.Printf("Type %T\n", a)
	fmt.Println(a)
	fmt.Println(*a)

	b := new(bool)
	fmt.Printf("Type %T\n", b)
	fmt.Println(b)
	fmt.Println(*b)

	*b = true
	fmt.Printf("Type %T\n", b)
	fmt.Println(b)
	fmt.Println(*b)
}

//make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，
//而且它返回的类型就是这三个类型本身，而不是他们的指针类型，
//因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。
//func make(t Type, size ...IntegerType) Type
func TestMakeFunc(t *testing.T) {
	var b map[string]int
	b = make(map[string]int, 10)
	b["A"] = 100
	fmt.Println(b)
}

//new与make的区别
//二者都是用来做内存分配的。
//make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
//而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
