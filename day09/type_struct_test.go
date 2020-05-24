package day09

import (
	"fmt"
	"testing"
)

type EmployeeCopy struct {
	ID   string
	Name string
	Age  int
}

//行为（方法）的定义
//第一种：定义方式在实例对应方法被调用时，实力的实例的成员会进行值复制
func (e EmployeeCopy) String() string {
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.ID, e.Name, e.Age)
}
func TestStructOperations(t *testing.T) {
	e := EmployeeCopy{"1", "Bob", 23}

	t.Log(e.String())
}

//第二种：通常情况下为了避免内存拷贝我们使用第二种定义方式
//func (e *EmployeeCopy) String() string {
//	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.ID, e.Name, e.Age)
//}
//
//func TestStructOperations(t *testing.T) {
//	e := &EmployeeCopy{"1", "Bob", 23}
//
//	t.Log(e.String())
//}
