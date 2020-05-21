package day04

import "testing"

func TestMapBasic(t *testing.T) {
	m1 := map[string]int{"one": 1, "two": 2}
	t.Log("This is Map Usage about 'm[\"one\"]=' ", m1["one"])
	t.Logf("len m1 is %d", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2= %d", len(m2))
	m3 := make(map[int]int, 20)
	t.Logf("len m3= %d", len(m3))
	t.Log(m3)

}
