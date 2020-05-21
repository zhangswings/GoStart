package day06

import (
	"strconv"
	"strings"
	"testing"
)

/**
Unicode Utf8
1. Unicode 是一种字符集（code point）
2. Utf8 是 Unicode 的存储实现（转换为字节序列的规则）
*/
func TestStringMethod(t *testing.T) {
	var s string
	t.Log(s) //初始化的默认值为""

	s = "hello"
	t.Log(len(s))
	//string 是不可变的byte slice
	t.Log(s[1])

	//可以存储任何二进制数据
	s = "\xE4\xB8\xA5"
	t.Log(s)

	s = "中"
	t.Log(s)
	t.Log(len(s)) //是byte数

	c := []rune(s) //rune 取出字符串的 Unicode
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 utf8 %x", s)
}

func TestString(t *testing.T) {

	s := "中华人民共和国"

	//rune
	for _, c := range s {
		t.Logf("%[1]c %[1]d", c)
	}
}

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")

	t.Log(parts)

	for _, part := range parts {
		t.Log(part)
	}

	t.Log(strings.Join(parts, "-"))
}

func TestStringToInteger(t *testing.T) {

	s := strconv.Itoa(10)
	t.Log("str" + s)

	if atoi, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + atoi)
	} else {
		t.Log("strconv.Atoi ERROR")
	}

}
