package day10

import (
	"fmt"
	"testing"
)

// test commit
func TestTypeUsage(t *testing.T) {
	// t.Log("Hello World")

	t.Log("intSum(1, 2)= ", intSum(1, 2))
	t.Log("intSum2(1, 2, 3, 4, 5)= ", intSum2(1, 2, 3, 4, 5))
	t.Log("intSum3(1, 2, 3, 4, 5, 6)= ", intSum3(1, 2, 3, 4, 5, 6))
}

func intSum(x, y int) int {
	return x + y
}

func intSum2(x ...int) int {
	fmt.Println(x) //x 是一个切片

	sum := 0
	for _, v := range x {
		sum += v
	}

	return sum
}

func intSum3(x int, y ...int) int {
	sum := 0
	sum += x
	for _, value := range y {
		sum += value
	}
	return sum
}
