/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import "fmt"

func main() {
	// 存储着5个函数
	var functionList []func()

	for i := 0; i < 5; i++ {
		function := func(arg int) func() {
			return func() {
				fmt.Println(arg)
			}
		}(i)
		functionList = append(functionList, function)
	}

	// 运行函数
	functionList[0]()
	functionList[1]()
	functionList[2]()
}