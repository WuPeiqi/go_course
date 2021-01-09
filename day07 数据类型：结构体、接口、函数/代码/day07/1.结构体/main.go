/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import "fmt"

func main() {

	type Person struct {
		name    string
		age     int
		hobby   [2]string
		num     []int
		parent  map[string]string
	}

	p1 := Person{
		name:   "二狗子",
		age:    19,
		hobby:  [2]string{"裸奔", "大保健"}, // 拷贝
		num:    []int{69, 19, 99, 38}, // 未拷贝 (内部维护指针指向数据存储的地方)
		parent: map[string]string{"father": "Alex", "mother": "Monika"}, // 未拷贝 (内部维护指针指向数据存储的地方)
	}
	p2 := p1

	fmt.Println(p1)
	fmt.Println(p2)
	p1.parent["father"] = "武沛齐"
	fmt.Println(p1)
	fmt.Println(p2)
}

























