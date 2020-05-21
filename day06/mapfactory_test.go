package day06

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}

	m[1] = func(op int) int {
		return op
	}
	m[2] = func(op int) int {
		return op * op
	}
	m[3] = func(op int) int {
		return op * op * op
	}

	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}

	mySet[1] = true
	n := 3
	if mySet[n] {
		t.Log("1 is existing")
	} else {
		t.Logf("%d is not existing", n)
	}

	mySet[3] = false
	t.Log(len(mySet))
	t.Log(mySet)
	delete(mySet, 1)
	n = 1
	if mySet[n] {
		t.Log("1 is existing")
	} else {
		t.Logf("%d is not existing", n)
	}

}
