package day05

import "testing"

func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
}

/**
if 条件
与其他主要编程语言的差异
1. condition 表达式结果必须为布尔值
2. 支持变量赋值：
	if var declaration; condition{
		//Code to be executed if condition is true
	}
*/
func TestIfCondition(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("1==1")
	}

}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := -1; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}
