/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import "fmt"

type School struct {
	band string
	city string
}

type Class struct {
	title  string
	count  int
	school *School
}

func main() {
	// 1.创建学校
	sch := &School{"路飞学城", "北京"}

	var classList []Class
	// 2.循环创建班级
	for{
		// 用户输入品牌和城市
		//var title string
		//var count int
		var cls Class
		fmt.Printf("请输入班级：")
		fmt.Scanf("%s",&cls.title)
		if cls.title == "Q"{
			break
		}
		fmt.Printf("请输入人数：")
		fmt.Scanf("%d",&cls.count)

		// 创建班级对象
		cls.school = sch

		// 加入到班级的列表（切片）
		classList = append(classList,cls)
	}
	fmt.Println(classList) // [{python1 10 0xc00000c060} {python2 90 0xc00000c060}]

	for _,item := range  classList{
		message := fmt.Sprintf("%s%s 校区，%s 班级有%d多学生。",item.school.band,item.school.city,item.title,item.count)
		fmt.Println(message)
	}
}
