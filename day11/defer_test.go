package day11

import (
	"fmt"
	"testing"
)

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

//问，上面代码的输出结果是？（提示：defer注册要延迟执行的函数时该函数所有的参数都需要确定其值）
func TestCalcFunc(t *testing.T) {
	x := 1
	y := 2

	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20

	//1. A 1 2 3
	//2. B 10 2 12
	//xx  3. AA 1 3 4
	//xx  4. BB 10 4 14
	//✅ 3.BB 10 12 22
	//✅ 4.AA 1 3 4
}
