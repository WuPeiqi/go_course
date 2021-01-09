/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import "fmt"

func do(num ...int) int {
	fmt.Println(num) // [1 2 3 3]
	sum := 0
	for _, value := range num {
		sum += value
	}
	return sum
}

func main() {
	r1 := do(1, 2, 3, 3)
	r2 := do(0, 1, 1)
	fmt.Println(r1, r2)
}
