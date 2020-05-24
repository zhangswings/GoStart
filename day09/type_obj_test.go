package day09

import "testing"

type Employee struct {
	ID   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "zhangswings", 20}
	t.Log(e)

	e1 := Employee{Name: "Mike", Age: 20}
	e1.ID = "123456"
	t.Log(e1)

	e2 := new(Employee)
	e2.Age = 20
	e2.Name = "zhangswings"
	e2.ID = "100100"
	t.Log(e2)

	t.Log("e1.ID= ", e1.ID)
	t.Logf("e is %T", e)

	t.Log("e2= ", &e2)
	t.Logf("e2 is %T", e2)

}
