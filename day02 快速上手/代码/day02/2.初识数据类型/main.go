package main

import "fmt"

func main() {
	// 整型
	fmt.Println(666)
	fmt.Println(6 + 9)
	fmt.Println(6 - 9)
	fmt.Println(6 * 9)
	fmt.Println(16 / 9) // 商
	fmt.Println(16 % 9) // 余数

	// 字符串类型，特点：通过双引号
	fmt.Println("武沛齐")
	fmt.Println("钓鱼要掉刀鱼，刀鱼到岛上钓")
	fmt.Println("alex" + "SB")
	//fmt.Println("alex" + 666)
	fmt.Println("alex" + "666")
	// 对比
	fmt.Println("1" + "2") // 结果："12"
	fmt.Println(1 + 2)     // 结果：3

	// 布尔类型，真假
	fmt.Println(1 > 2) // false  假
	fmt.Println(1 < 2) // true   真
	fmt.Println(1 == 2)
	fmt.Println(1 >= 2)
	fmt.Println("Alex" == "sb")

	// 超前
	if 2 > 1 {
		fmt.Println("叫爸爸")
	} else {
		fmt.Println("孙子")
	}
}
