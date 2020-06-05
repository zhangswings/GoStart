package day14

import "testing"

func TestStudent(t *testing.T) {

	m := make(map[string]student)
	stus := []student{
		{"A", "Beijing", 12},
		{"B", "Beijing", 13},
		{"C", "Beijing", 14},
	}
	for _, stu := range stus {
		m[stu.name] = stu
		//t.Log(&stu)
	}
	for k, v := range m {
		t.Log(k, "=>", v)
		//t.Log( m)
	}
}

type student struct {
	name, city string
	age        int8
}

func newPerson(name, city string, age int8) *student {

	return &student{
		name: name,
		city: city,
		age:  age,
	}
}

func TestStructFuncMethod(t *testing.T) {
	p9 := newPerson("zhangswings", "shanghai", 12)
	t.Logf("%#v\n", p9)
}
