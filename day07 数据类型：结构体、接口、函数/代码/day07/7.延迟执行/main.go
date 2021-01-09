/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import "fmt"

func after()  {
	fmt.Println("after函数")
}

func do() int {
	fmt.Println("风吹")
	defer fmt.Println("函数执行完毕了")  // 如果在这行之前执行return，那么defer就不再执行
	defer after()  // 如果在这行之前执行return，那么defer就不再执行
	fmt.Println("屁屁凉")
	return 666
}

func main() {
	ret := do()
	fmt.Println(ret)
}

