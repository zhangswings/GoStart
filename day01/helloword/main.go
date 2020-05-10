package main

import "fmt"

// func main()  {
// 	fmt.Println("Hello world.")
// 	fmt.Println("Hello Hello")
// }

// 批量声明
var (
	name string //" "
	age  int    //0
	isOK bool   //false
)

func main() {
	name = "lixiang"
	age = 16
	isOK = true

	fmt.Println(isOK)
	fmt.Printf("name: %s\n", name)
	fmt.Println(age)

}
