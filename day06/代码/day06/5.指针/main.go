/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

func main() {
	name := "武沛齐"
	var p1 *string = &name
	var p2 **string = &p1
	var p3 ***string = &p2

	println(name, &name)
	println(p1, &p1)
	println(p2, &p2)
	println(p3, &p3)
}
