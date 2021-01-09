/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import "fmt"

func main() {
	r1 := 5 & 99
	r2 := 5 | 99
	r3 := 5 ^ 99
	r4 := 5 << 2
	r5 := 5 >> 1
	r6 := 5 &^ 99
	fmt.Println(r1, r2, r3, r4, r5, r6)
}
