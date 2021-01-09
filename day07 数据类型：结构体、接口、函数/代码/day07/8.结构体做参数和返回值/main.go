/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import "fmt"

type Person struct {
	name string
	age  int
}

var P Person = Person{"wupeiqi", 18}

func doSomething() Person {
	return P
}

func main() {
	data := doSomething()
	P.name = "alex"
	fmt.Println(data)
	fmt.Println(P)
}
