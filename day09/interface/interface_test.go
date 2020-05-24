package _interface

import "testing"

//接口
//与其他主要编程语言的差异
//1、接口为非入侵性，实现不依赖于接口定义；
//2、所以接口的定义可以包含在接口使用者包内；
type Progrommer interface {
	WriteHelloWorld() string
}

type GoProgrommer struct {
	Name string
}

func (goPro *GoProgrommer) WriteHelloWorld() string {

	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Progrommer
	p = new(GoProgrommer)

	t.Log(p.WriteHelloWorld())
}
