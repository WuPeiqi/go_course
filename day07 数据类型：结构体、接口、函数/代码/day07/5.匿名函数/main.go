/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import "fmt"

func F1(n1 int, n2 int) func(int) string {

	return func(n1 int) string {
		fmt.Println("匿名函数")
		return "匿名"
	}
}

func main() {
	// 匿名函数
	v1 := func(n1 int, n2 int) int {
		return 123
	}
	data := v1(11, 22)
	fmt.Println(data)

	value := func(n1 int, n2 int) int {
		return 123
	}(11, 22)
	fmt.Println(value)

}
