/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import "fmt"

func main() {
	var name string = "武沛齐"



	// 4. for range循环获取所有字符
	for index, item := range name {
		fmt.Println(index, item,string(item))
	}


}