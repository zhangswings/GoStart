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
