package day11

import (
	"fmt"
	"testing"
)

//内置函数介绍

// close 主要用来关闭channel
// len 用来求长度，比如string,array,slice,map,channel
// new 用来分配内存，主要用来分配值类型，比如 int、struct。返回的是指针
// make 用来分配内存，主要用来分配引用类型，比如chan、map、slice
// append 用来追加元素到数组、slice中
// panic 和 recover 用来做错误处理

func funA() {
	fmt.Println("func A")
}

//注意：
//recover()必须搭配defer使用。
//defer一定要在可能引发panic的语句之前定义。
func funB() {

	defer func() {
		err := recover()
		//如果程序出现了panic错误，可以通过recover恢复过来
		if err != nil {
			fmt.Println("happen error: ", err)
			fmt.Println("recover in funB")
		}
	}()
	panic("panic in B")
}

func funC() {
	fmt.Println("func C")
}

func TestFunPanic(t *testing.T) {
	funA()
	funB()
	funC()
}
