/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import (
	"fmt"
	"unsafe"
)

func main() {

	dataList := [3]int8{11, 22, 33}

	// 1.获取数组第一个元素的地址（指针）
	var firstDataPtr *int8 = &dataList[0]

	// 2.转换成Pointer类型
	ptr := unsafe.Pointer(firstDataPtr)

	// 3.转换成uintptr类型，然后进行内存地址的计算（即：地址加1个字节，意味着取第2个索引位置的值）。
	targetAddress := uintptr(ptr) + 2

	// 4.根据新地址，重新转换成Pointer类型
	newPtr := unsafe.Pointer(targetAddress)

	// 5.Pointer对象转换为 int8 指针类型
	value := (*int8)(newPtr)

	// 6.根据指针获取值
	fmt.Println("最终结果为：", *value)
}