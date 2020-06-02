package main

func main() {

	// 示例1：死循环
	/*
		fmt.Println("开始")
		for {
			fmt.Println("红鲤鱼与绿鲤鱼与驴")
			time.Sleep(time.Second * 1) // 等一秒再继续执行
		}
		fmt.Println("结束")
	*/

	// 示例2：
	/*
		fmt.Println("开始")
		for 2 > 1 {
			fmt.Println("红鲤鱼与绿鲤鱼与驴")
			time.Sleep(time.Second * 1) // 等一秒再继续执行
		}
		fmt.Println("结束")
	*/

	// 示例3：
	/*

		fmt.Println("开始")
		number := 1
		for number < 5 {
			fmt.Println("钓鱼要掉刀鱼，刀鱼要到岛上钓")
			number = 10
		}
		fmt.Println("结束")
	*/

	// 示例4：
	/*
		fmt.Println("开始")
		number := 1
		for number < 5 {
			fmt.Println("钓鱼要掉刀鱼，刀鱼要到岛上钓")
			number = number + 1
		}
		fmt.Println("结束")

	*/

	// 示例5：布尔类型的变量
	/*
		fmt.Println("开始")
		flag := true
		for flag {
			fmt.Println("钓鱼要掉刀鱼，刀鱼要到岛上钓")
			flag = false
		}
		fmt.Println("结束")

	*/

	//for i := 1; i < 10; i = i + 2 {
	//	fmt.Println("钓鱼要掉刀鱼，刀鱼要到岛上钓")
	//}

	//num := 10
	//fmt.Println(num)
	//num++ // 等价于 num = num + 1
	//fmt.Println(num)

	// ######################## 阶段练习题 ########################
	// 1. 猜数字，设定一个理想数字比如：66，一直提示让用户输入数字，如果比66大，则显示猜测的结果大了；如果比66小，则显示猜测的结果小了;只有输入等于66，显示猜测结果正确，然后退出循环。
	/*
		data := 66
		flag := true
		for flag {
			var userInputNumber int
			fmt.Print("请输入数字：")
			fmt.Scanln(&userInputNumber)
			if userInputNumber > data {
				fmt.Println("大了")
			} else if userInputNumber < data {
				fmt.Println("小了")
			} else {
				fmt.Println("恭喜你猜对了")
				flag = false
			}
		}
	*/

	// 2. 使用循环输出1~100所有整数。
	/*
		for i := 1; i <= 100; i++ {
			fmt.Println(i)
		}

	*/

	// 3. 使用循环输出 1 2 3 4 5 6 8 9 10，即：10以内除7以外的整数。
	/*
		for i := 1; i <= 10; i++ {
			if i != 7{
				fmt.Println(i)
			}
		}
	*/

	// 4. 输出 1~100 内的所有奇数。
	/*
		for i := 1; i <= 100; i++ {
			if i%2 == 1{
				fmt.Println(i)
			}
		}
	*/

	// 5. 输出 1~100 内的所有偶数。
	/*
		for i := 1; i <= 100; i++ {
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
	*/
	// 6. 求 1~100 的所有整数的和。 1 + 2 + 3 + 4 + 5 ...+ 100
	/*
		sum := 0
		for i := 1; i <= 100; i++ {
			sum = sum + i
		}
		fmt.Println(sum)
	*/
	// 7. 输出10~1 所有整数。
	/*
		for i := 10; i > 0; i-- {
			fmt.Println(i)
		}
	*/

}
