package main

import "fmt"

func main() {
	var name string
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)

	if name == "wupeiqi" {
		// svip
		goto SVIP
	} else if name == "yuanhao" {
		// vip
		goto VIP
	}
	fmt.Println("预约...")
VIP:
	fmt.Println("等号...")
SVIP:
	fmt.Println("进入...")
}
