/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

func main() {
	//n1 := 666
	//n2 := 9223372036854775807
	//n3 := -9223372036854775809
	//fmt.Println(n1, n2, n3)

	//var v1 int8 = 10
	//var v2 int16 = 18
	//v3 := int16(v1) + v2
	//fmt.Println(v3, reflect.TypeOf(v3))

	//var v1 int16 = 130
	//v2 := int8(v1)
	//fmt.Println(v2)

	// 整型转换为字符串类型
	/*
		v1 := 19
		result := strconv.Itoa(v1)
		fmt.Println(result, reflect.TypeOf(result))
		var v2 int8 = 17
		data := strconv.Itoa(int(v2))
		fmt.Println(data,reflect.TypeOf(data))
	*/

	// 字符串转换为整型：转换后是int类型;可能存在错误
	/*
		v1 := "666"
		result, err := strconv.Atoi(v1)
		if err == nil {
			fmt.Println("转换成功", result,reflect.TypeOf(result))
		} else {
			fmt.Println("转换失败")
		}

	*/

	//v1 := 99
	//// 整型(十进制）转换为:二进制、八进制、十六进制
	//r1 := strconv.FormatInt(int64(v1), 16)
	//fmt.Println(r1, reflect.TypeOf(r1))

	// data，要转换的文本
	// 2，把文档当做二进制去转换成 十进制（整型）
	// 16，转换的过程中对结果进行约束
	// 结果：如果转换成功，则将err设置为nil，result则永远以int64的类型返回。
	//data := "1001000101"
	//result, err := strconv.ParseInt(data, 2, 16)
	//fmt.Print(result, err, reflect.TypeOf(result))
	//
	////- 将十进制 14 用 转换成16进制的字符串。
	//v1 := strconv.FormatInt(14, 16)
	//fmt.Println(v1)
	//
	////- 将 2进制 "10011" 转换成 10进制的整型。
	//v2, _ := strconv.ParseInt("10011", 2, 0)
	//fmt.Println(v2)
	////- 将 2进制 "10010" 转换成 16 进制的字符串。
	//v3, _ := strconv.ParseInt("10011", 2, 0)
	//v4 := strconv.FormatInt(v3, 16)
	//fmt.Println(v4)

	//fmt.Println(math.Abs(-19))                // 取绝对值
	//fmt.Println(math.Floor(3.14))             // 向下取整
	//fmt.Println(math.Ceil(3.14))              // 向上取整
	//fmt.Println(math.Round(3.3478))           // 就近取整
	//fmt.Println(math.Round(3.5478*100) / 100) // 保留小数点后两位
	//fmt.Println(math.Mod(11, 3))              // 取余数，同11 % 3
	//fmt.Println(math.Pow(2, 5))               // 计算次方，如：2的5次方
	//fmt.Println(math.Pow10(2))                // 计算10次方，如：2的10次方
	//fmt.Println(math.Max(1, 2))               // 两个值，取较大值
	//fmt.Println(math.Min(1, 2))               // 两个值，取较小值
	//
	//// 第一步：创建一个超大整型的一个对象
	////var v1 big.Int
	////var v2 *big.Int // 一般用不到（直接复制的时候使用）
	//v3 := new(big.Int)
	//
	//// 第二步：在超大整型对象中写入值
	////v1.SetInt64(9223372036854775808)
	////fmt.Println(v1)
	////v1.SetString("92233720368547758089223372036854775808", 10)
	////fmt.Println(v1)
	////v1.SetString("100011100001111010011", 2)
	////fmt.Println(v1)
	//
	////v3.SetInt64(9223372036854775807)
	////fmt.Println(v3)
	////v3.SetString("92233720368547758089223372036854775808", 10)
	////fmt.Println(v3)

	//n1 := big.NewInt(89)
	//
	//n2 := big.NewInt(99)
	//
	//result := new(big.Int)
	//result.Add(n1, n2)
	//
	//fmt.Println(result.Int64(),result.String())

	//n1 := new(big.Int)
	//n1.SetString("92233720368547758089223372036854775808", 10)
	//
	//n2 := new(big.Int)
	//n2.SetString("11111192233720368547758089223372036854775808", 10)
	//
	//result := new(big.Int)
	//result.Add(n1, n2)
	//
	//fmt.Println(result.String())

	//var v1 big.Int
	//v1.SetString("92233720368547758089223372036854775808", 10)
	//
	//var v2 big.Int
	//v2.SetString("2", 10)
	//
	////result := new(big.Int)
	//var result big.Int
	//result.Add(&v1, &v2)
	//fmt.Println(result.String())

}
