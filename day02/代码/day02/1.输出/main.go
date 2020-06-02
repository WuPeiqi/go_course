package main

import "fmt"

func main() {
	// 基于内置函数（不推荐）

	/*
	print("好吃不过饺子 \n")   // 下下下
	print("好玩不过嫂子 \n")
	println("好吃不过饺子")
	println("好玩不过嫂子")
	*/

	//fmt包（推荐）
	//fmt.Print("好吃不过饺子 \n")
	//fmt.Print("好玩不过嫂子 \n")
	//fmt.Println("好玩不过嫂子")
	//fmt.Println("好玩不过嫂子")
	//fmt.Println("我叫", "Alex", "我媳妇", "是个", "....")
	//fmt.Println("我叫Alex我媳妇是个....")

	// fmt包 扩展：格式化输出
	// %s，占位符 "文本"
	// %d，占位符 整数
	// %f，占位符 小数（浮点数）
	// %.2f，占位符 小数（浮点数）
	// 百分比
	fmt.Printf("老汉开着%s,去接%s的媳妇，多少钱一次？%d块。嫂子给打折吧，%.2f怎么样？小哥哥包你100%%满意", "车", "老男孩", 100, 3.889)
}
