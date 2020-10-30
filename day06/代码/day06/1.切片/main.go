/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import "fmt"

func main() {

	v1 := [][]int{make([]int, 2, 5), make([]int, 2, 5)}

	v2 := append(v1[0], 99)
	v2[0] = 69
	fmt.Println(v1)
}
