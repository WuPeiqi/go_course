/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串 转换 布尔类型
	// true:"1", "t", "T", "true", "TRUE", "True"
	// false:"0", "f", "F", "false", "FALSE", "False"
	//false,err错误信息
	v1, err := strconv.ParseBool("t")
	fmt.Println(v1, err)

	// 布尔类型 转换 字符串
	v2 := strconv.FormatBool(false)
	fmt.Println(v2)
}
