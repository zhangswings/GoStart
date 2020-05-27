package day11

import (
	"fmt"
	"strings"
	"testing"
)

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func TestStringMapFunc(t *testing.T) {
	fmt.Println(users)
	fmt.Println(distribution)

	var left int = dispatchCoin()
	fmt.Println("剩下：", left)
	fmt.Println(distribution)

	//=== RUN   TestStringMapFunc
	//[Matthew Sarah Augustus Heidi Emilie Peter Giana Adriano Aaron Elizabeth]
	//map[]
	//剩下： 25
	//map[Aaron:3 Adriano:5 Augustus:4 Elizabeth:3 Emilie:3 Giana:2 Heidi:3 Matthew:1 Peter:1]
	//--- PASS: TestStringMapFunc (0.00s)
	//PASS
}

func dispatchCoin() int {

	coinsCopy := coins
	for _, v := range users {
		//fmt.Println(v)
		if strings.Contains(v, "E") || strings.Contains(v, "e") {
			distribution[v] += 1
			coinsCopy -= 1
		}
		if strings.Contains(v, "I") || strings.Contains(v, "i") {
			distribution[v] += 2
			coinsCopy -= 2
		}
		if strings.Contains(v, "O") || strings.Contains(v, "o") {
			distribution[v] += 3
			coinsCopy -= 3
		}
		if strings.Contains(v, "U") || strings.Contains(v, "u") {
			distribution[v] += 4
			coinsCopy -= 4
		}
	}

	return coinsCopy
}
