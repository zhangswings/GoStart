package day03

import "testing"

func TestSliceGrowthing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
	t.Log(s)
}

/**
切片共享内存存储空间
*/
func TestSliceShareMemory(t *testing.T) {
	year := []string{"A", "B", "C", "D", "E", "F"}

	Q2 := year[2:6]
	t.Log(Q2, len(Q2), cap(Q2))

	Q1 := year[0:4]
	Q1[2] = "UNKNOWN"
	t.Log(Q1, len(Q1), cap(Q1))

	t.Log(Q2, len(Q2), cap(Q2))
}
