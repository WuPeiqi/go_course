/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import (
	"fmt"
	"reflect"
)

func main() {
	type Person struct {
		name string "姓名"
		age  int32  "年龄"
		blog string "博客"
	}

	p1 := Person{name: "武沛齐", age: 18, blog: "https://www.pythonav.com"}

	p1Type := reflect.TypeOf(p1)
	// 方式1
	filed1 := p1Type.Field(0)
	fmt.Println(filed1.Tag)

	// 方式2
	filed2, _ := p1Type.FieldByName("name")
	fmt.Println(filed2.Tag)

	// 循环获取
	fieldNum := p1Type.NumField()
	for index := 0; index < fieldNum; index++ {
		field := p1Type.Field(index)
		fmt.Println(field.Name, field.Tag)
	}
}
