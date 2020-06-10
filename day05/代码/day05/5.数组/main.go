/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

func main() {
	//nums := [3]int32{11, 22, 33}
	//
	//fmt.Printf("数组的内存地址：%p \n", &nums)
	//fmt.Printf("数组第1个元素的内存地址：%p \n", &nums[0])
	//fmt.Printf("数组第2个元素的内存地址：%p \n", &nums[1])
	//fmt.Printf("数组第3个元素的内存地址：%p \n", &nums[2])

	//names := [2]string{"武沛齐", "alex"}
	//fmt.Printf("数组的内存地址：%p \n", &names)
	//fmt.Printf("数组第1个元素的内存地址：%p \n", &names[0])
	//fmt.Printf("数组第2个元素的内存地址：%p \n", &names[1])

	//name1 := [2]string{"武沛齐", "alex"}
	//name2 := name1
	//
	//name1[1] = "苑昊"
	//
	//fmt.Println(name1,name2)   // [武沛齐 苑昊]   [武沛齐 alex]

	// 1. 长度
	//name := [2]string{"武沛齐", "alex"}
	//fmt.Println(len(name))

	// 2. 索引
	//name := [2]string{"武沛齐", "alex"}
	//data := name[0]
	//fmt.Println(data)
	//name[0] = "eric"
	//fmt.Println(name)

	// 3. 切片
	//nums := [3]int32{11, 22, 33}
	//data := nums[0:2] // 获取   0 <= 下标 < 2
	//fmt.Println(data)

	// 4. 循环
	//nums := [3]int32{11, 22, 33}
	//for i:=0;i<len(nums);i++{
	//	fmt.Println(i, nums[i] )
	//}

	// 5.for range 循环
	//nums := [3]int32{11, 22, 33}
	//for key, item := range nums {
	//	fmt.Println(key, item)
	//}
	//
	//for key := range nums {
	//	fmt.Println(key)
	//}
	//
	//for _,item := range nums {
	//	fmt.Println(item)
	//}

	// [0,0,0]
	//var nestData [3]int

	// [ [  [0,0,0],[0,0,0]  ],[  [0,0,0],[0,0,0]  ],[  [0,0,0],[0,0,0]  ], ]
	//var nestData [3][2][3]int

	// [  [0,0,0],[0,0,0]  ]
	//var nestData [2][3]int
	//nestData[0] = [3]int{11, 22, 33}
	//nestData[1][1] = 666
	//fmt.Println(nestData)

	//nestData := [2][3]int{[3]int{11, 22, 33}, [3]int{44, 55, 66}}
	//fmt.Println(nestData)



}
