/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import "fmt"

func main() {
	v1 := []int{6, 9}
	v2 := v1
	fmt.Println(v1, v2)              // [6 9] [6 9]
	fmt.Printf("v1的内存地址：%p \n", &v1) // 0xc0000a6020
	fmt.Printf("v2的内存地址：%p \n", &v2) // 0xc0000a6040

	v1 = append(v1, 999)
	fmt.Println(v1, v2) // [11111 9] [11111 9]
}
