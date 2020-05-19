package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}

	arr3 := [...]int{1, 2, 3, 4, 56, 6}
	arr1[1]=5

	t.Log(arr1[1],arr[2])
	t.Log(arr1,arr3)
}

