package day14

import (
	"testing"
	"unsafe"
)

type test struct {
	a, b, c, d int8
}

func TestMemoryFunc(t *testing.T) {
	n := test{
		1, 2, 3, 4,
	}

	t.Logf("n.a %p\n", &n.a)
	t.Logf("n.a %p\n", &n.b)
	t.Logf("n.a %p\n", &n.c)
	t.Logf("n.a %p\n", &n.d)
	t.Log(unsafe.Sizeof(n))

	//n.a 0xc0000a0060
	//n.b 0xc0000a0061
	//n.c 0xc0000a0062
	//n.d 0xc0000a0063
	//【进阶知识点】关于Go语言中的内存对齐推荐阅读:在 Go 中恰到好处的内存对齐
	//https://segmentfault.com/a/1190000017527311?utm_campaign=studygolang.com&utm_medium=studygolang.com&utm_source=studygolang.com

	//	空结构体
	//	空结构体是不占用空间的。
	var v struct{}
	t.Log(unsafe.Sizeof(v))
}
