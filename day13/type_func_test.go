package day13

import (
	"fmt"
	"testing"
)

//结构体
//string、整型、浮点数、布尔等数据类型
//将MyInt 定义为int 类型
//type MyInt int

//类型别名，是Go1.9版本添加的新功能
//type TypeAlias = Type
type byte = uint8
type rune = int32

//类型定义和类型别名的区别

//类型定义
type NewInt int

//类型别名
type MyInt = int

func TestTypeAlias(t *testing.T) {
	var a NewInt
	var b MyInt

	t.Logf("type of a:%T\n", a)
	t.Logf("type of b:%T\n", b)
	//	结果显示a的类型是main.NewInt，表示main包下定义的NewInt类型。
	//	b的类型是int。MyInt类型只会在代码中存在，编译完成时并不会有MyInt类型。
}

//结构体
//使用 type 和 struct 关键字来定义结构体，具体代码格式如下：
// type 类型名 struct {
//   字段名  字段类型
//   字段名  字段类型
//   ```
// }
//其中：
//
//类型名：标识自定义结构体的名称，在同一个包内不能重复。
//字段名：表示结构体字段名。结构体中的字段名必须唯一。
//字段类型：表示结构体字段的具体类型。

type person struct {
	name string
	age  int8
	city string
}

type person1 struct {
	name, city string
	age        int8
}

//结构体实例化
//只有当结构体实例化的时候，才会真正地分配内存，也就是必须实例化之后才能使用结构体的字段。
//结构体本身也是一种类型，我们可以像声明内置类型一样使用 var 关键字声明结构类型。

// var 结构体实例 结构体类型

func TestStructFunc(t *testing.T) {

	var p1 person
	p1.name = "zhangswings"
	p1.age = 22
	p1.city = "Beijing"

	t.Logf("Person is =%v\n", p1)
	t.Logf("Person is =%#v\n", p1)

	//	  type_func_test.go:70: Person is ={zhangswings 22 Beijing}
	//    type_func_test.go:71: Person is =day13.person{name:"zhangswings", age:22, city:"Beijing"}
}

//匿名结构体

func TestAnonymousFunc(t *testing.T) {
	var user struct {
		name string
		age  int8
	}

	user.age = 1
	user.name = "Nomi"
	fmt.Printf("%#v\n", user)

	//	创建指针类型结构体
	//	new 关键字
	var p2 = new(person)
	t.Logf("p2 is %#v\n", p2)
	t.Logf("p2 is %v\n", p2)

	p2.name = "Nomi"
	p2.age = 1
	p2.city = "Beijing"
	t.Logf("p2 is %#v\n", p2)
	t.Logf("p2 is %v\n", p2)

	p3 := &person{}
	t.Logf("p3 is %#v\n", p3)
	t.Logf("p3 is %v\n", p3)
	p3.name = "Nomi"
	p3.age = 1
	p3.city = "Beijing"
	t.Logf("p3 is %#v\n", p3)
	t.Logf("p3 is %v\n", p3)

	//是要键值对对结构体进行初始化时，键对应的结构体的字段，值对应该字段类型的初始值。
	p5 := person{
		name: "Nomi",
		age:  1,
		city: "Beijing",
	}
	t.Logf("p5 is %#v\n", p5)
	t.Logf("p5 is %v\n", p5)

	p6 := &person{
		name: "Nomi",
		age:  1,
		city: "Beijing",
	}
	t.Logf("p6 is %#v\n", p6)
	t.Logf("p6 is %v\n", p6)
	// 使用键值对初始化
	p7 := &person{
		name: "Nomi",
	}
	t.Logf("p7 is %#v\n", p7)
	t.Logf("p7 is %v\n", p7)
}
