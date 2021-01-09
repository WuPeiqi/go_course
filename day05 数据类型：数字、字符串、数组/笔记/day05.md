# day05 数据类型

> 写程序 等价于 写作文

数据类型，其实就是各种各样类型的数据。

![image-20200607215200860](assets/image-20200607215200860.png)

Go语言中常见的数据类型有挺多，例如：

- 整型，用于表示整数。
- 浮点型，用于表示小数。
- 布尔型，用于表示真/假。
- 字符串，用于表示文本信息。
- 数组，用于表示多个数据（数据集合）
- 指针，用于表示内存地址的类型。
- 切片，用于表示多个数据（数据集合）
- 字典，用于表示键值对结合。
- 结构体，用于自定义一些数据集合。
- 接口，用于约束和泛指数据类型。

注意：关于“值类型”和“引用类型” [https://github.com/go101/go101/wiki/About-the-terminology-%22reference-type%22-in-Go](https://github.com/go101/go101/wiki/About-the-terminology-"reference-type"-in-Go)

## 今日概要

- 整型，用于表示整数。
- 浮点型，用于表示小数。
- 布尔型，用于表示真/假。
- 字符串，用于表示文本信息。
- 数组，用于表示多个数据（数据集合）["长沙","东莞","惠州"]

## 1.整型

Go中的整型分为有符号和无符号两大类，有符号的包含负值，无符号不包含负值。

有符号整型：

- int8（-128 -> 127）
- int16（-32768 -> 32767）
- int32（-2,147,483,648 -> 2,147,483,647）
- int64（-9,223,372,036,854,775,808 -> 9,223,372,036,854,775,807）
- int
  - 在 32 位操作系统上使用 32 位（-2,147,483,648 -> 2,147,483,647） 2**32
  - 在 64 位操作系统上使用 64 位（-9,223,372,036,854,775,808 -> 9,223,372,036,854,775,80）2**64

无符号整数：

- uint8（0 -> 255）
- uint16（0 -> 65,535）
- uint32（0 -> 4,294,967,295）
- uint64（0 -> 18,446,744,073,709,551,615）
- uint
  - 在 32 位操作系统上使用32 位（0 -> 4,294,967,295） 2**32
  -  64 位操作系统上使用 64 位（0 -> 18,446,744,073,709,551,615） 2**64

不同整型可表示的数据范围不同，我们需要根据自己的需求来选择适合的类型。



### 1.1 整型之间的转换

```go
data := intXXX(其他整型)
```

```go
var v1 int8 = 10
var v2 int16 = 18
v3 := int16(v1) + v2
fmt.Println(v3, reflect.TypeOf(v3))
```

注意：

- 地位转向高位，没问题。

- 高位转向低位，可能有问题

  ```go
  var v1 int16 = 130
  v2 := int8(v1)
  fmt.Println(v2)
  ```

  

### 1.2 整型与字符串的转换

```go
/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

func main() {

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
}
```

### 1.3 进制转换

- Go代码中：
  - 十进制，整型的方式存在。
  - 其他进制，是以字符串的形式存在。
- 整形，10进制数。



十进制转换为其他进制：

```go
/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	v1 := 99
	// 整型(十进制）转换为:二进制、八进制、十六进制
	r1 := strconv.FormatInt(int64(v1), 16)
	fmt.Println(r1, reflect.TypeOf(r1))

}

```



其他进制转换为十进制：

```go
/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// data，要转换的文本
	// 2，把文档当做二进制去转换成 十进制（整型）
	// 16，转换的过程中对结果进行约束
	// 结果：如果转换成功，则将err设置为nil，result则永远以int64的类型返回。
	data := "1001000101"
	result, err := strconv.ParseInt(data, 2, 0)
	fmt.Print(result, err, reflect.TypeOf(result))
}
```

提醒：通过ParseInt将字符串转换为10进制时，本质上与Atoi是相同的。

![image-20200608113745070](assets/image-20200608113745070.png)



练习题：

- 将十进制 14 用 转换成16进制的字符串。
- 将 2进制 "10011" 转换成 10进制的整型。
- 将 2进制 "10011" 转换成 16 进制的字符串。



### 1.4 常见数学运算

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Abs(-19))                // 取绝对值
	fmt.Println(math.Floor(3.14))             // 向下取整
	fmt.Println(math.Ceil(3.14))              // 向上取整
	fmt.Println(math.Round(3.3478))           // 就近取整
	fmt.Println(math.Round(3.5478*100) / 100) // 保留小数点后两位
	fmt.Println(math.Mod(11, 3))              // 取余数，同11 % 3
	fmt.Println(math.Pow(2, 5))               // 计算次方，如：2的5次方
	fmt.Println(math.Pow10(2))                // 计算10次方，如：2的10次方
	fmt.Println(math.Max(1, 2))               // 两个值，取较大值
	fmt.Println(math.Min(1, 2))               // 两个值，取较小值
	// ...
}
```

### 1.5 指针/nil/声明变量/new

- 声明变量

  ```go
  var v1 int 
  v2 := 999
  ```

  ![image-20200608195751136](assets/image-20200608195751136.png)

- 指针

  ```go
  var v3 *int    
  v4 := new(int)   
  ```

  ![image-20200608200651520](assets/image-20200608200651520.png)

- new关键字

  ```
  new用于创建内存并进行内部数据的初始化，并返回一个指针类型。
  ```

- nil

  ```
  nil指go语言中的空值。
  var v100 *int
  var v101 *int8
  ```



问题：

- 为什么要有指针？

  ```
  为了节省内存，不重复开辟空间去存储数据。
  ```

- int和*int是两种不同的数据类型，不等。（之后专门讲指针细说）

  

### 1.6 超大整型

#### 第一步：创建对象

创建超大整型对象

```go
// 第一步：创建一个超大整型的一个对象
var v1 big.Int
var v2 big.Int

// 第二步：在超大整型对象中写入值
v1.SetInt64(9223372036854775807)
fmt.Println(v1)

v2.SetString("92233720368547758089223372036854775808", 10)
fmt.Println(v2)
```

```go
// 第一步：创建一个超大整型的一个对象
v3 := new(big.Int)
v4 := new(big.Int)

// 第二步：在超大整型对象中写入值
v3.SetInt64(9223372036854775807)
fmt.Println(v3)

v4.SetString("92233720368547758089223372036854775808", 10)
fmt.Println(v4)
```

推荐：使用指针的方式，即：使用new来进行创建和初始化。



#### 第二步：加减乘除

超大对象进行加减乘除

```go
n1 := new(big.Int)
n1.SetInt64(89)

n2 := new(big.Int)
n2.SetInt64(99)

result := new(big.Int)
result.Add(n1, n2)

fmt.Println(result)
```

```go
n1 := big.NewInt(89)

n2 := big.NewInt(99)

result := new(big.Int)
result.Add(n1, n2)

fmt.Println(result)
```

其他：

```go
v1 := big.NewInt(11)
v2 := big.NewInt(3)
result := new(big.Int)

// 加
result.Add(v1, v2)
fmt.Println(result)
// 减
result.Sub(v1, v2)
fmt.Println(result)

// 乘
result.Mul(v1, v2)
fmt.Println(result)

// 除（地板除，只能得到商）
result.Div(v1, v2)
fmt.Println(result)

// 除，得商和余数
minder := new(big.Int)
result.DivMod(v1, v2, minder)
fmt.Println(result, minder)
```



#### 第三步：关于结果

```go
n1 := new(big.Int)
n1.SetString("92233720368547758089223372036854775808", 10)

n2 := new(big.Int)
n2.SetString("11111192233720368547758089223372036854775808", 10)

result := new(big.Int)
result.Add(n1, n2)

fmt.Println(result.String())
```



#### 最后建议

- 尽量new方式去初始化并返回一个指针类型的方式。

- 易错的点（int类型和*int类型）

  ```go
  var v1 big.Int
  v1.SetString("92233720368547758089223372036854775808", 10)
  
  var v2 big.Int
  v2.SetString("2", 10)
  
  //result := new(big.Int)
  var result big.Int
  result.Add(&v1, &v2)
  fmt.Println(result.String())
  ```

  

### 2.浮点型

浮点数，计算机中小数的表示方式，如：`3.14`

Go语言中提供了两种浮点型：

- float32，用32位（4个字节）来存储浮点型。
- float64，用64位（8个字节）来存储浮点型。

```go
/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import "fmt"

func main() {
	var v1 float32
	v1 = 3.14
	v2 := 99.9
	v3 := float64(v1) + v2
	fmt.Println(v1, v2, v3)
}
```



#### 2.1 非精确

float类型，计算机中小数的非精确的表示方式，如：`3.14`

```go
/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import "fmt"

func main() {
	var v1 float32
	v1 = 3.14
	v2 := 99.9
	v3 := float64(v1) + v2
	fmt.Println(v1, v2, v3)

	v4 := 0.1
	v5 := 0.2
	result := v4 + v5
	fmt.Println(result)

	v6 := 0.3
	v7 := 0.2
	data := v6 + v7
	fmt.Println(data)
}
```

```
3.14 99.9 103.04000010490418
0.30000000000000004
0.5
```



#### 2.2 float底层存储原理

```go
var price float32 = 0.29
```

##### 第一步：浮点型转换为二进制

- 整数部分，直接转换为二进制（10进制转换为2进制），即：`100111`

- 小数部分，让小数部分乘以2，结果小于1则将结果继续乘以2，结果大于1则讲结果-1继续乘以2，结果等于1则结束。

  ```
  0.29 * 2 = 0.58       // 小于1，则继续乘
  0.58 * 2 = 1.16		  // 大于1，则减1继续乘
  0.16 * 2 = 0.32		  // 小于1，则继续乘
  0.32 * 2 = 0.64       // 小于1，则继续乘
  0.64 * 2 = 1.28       // 大于1，则减1继续乘
  0.28 * 2 = 0.56       // 小于1，则继续乘
  0.56 * 2 = 1.12       // 大于1，则减1继续乘
  0.12 * 2 = 0.24       // 小于1，则继续乘    
  0.24 * 2 = 0.48       // 小于1，则继续乘
  0.48 * 2 = 0.96       // 小于1，则继续乘
  0.96 * 2 = 1.92       // 大于1，则减1继续乘
  0.92 * 2 = 1.84       // 大于1，则减1继续乘
  0.84 * 2 = 1.68       // 大于1，则减1继续乘
  0.68 * 2 = 1.36       // 大于1，则减1继续乘
  0.36 * 2 = 0.72       // 小于1，则继续乘
  0.72 * 2 = 1.44       // 大于1，则减1继续乘
  0.44 * 2 = 0.88       // 小于1，则继续乘
  0.88 * 2 = 1.76       // 大于1，则减1继续乘
  0.76 * 2 = 1.52       // 大于1，则减1继续乘
  0.52 * 2 = 1.04       // 大于1，则减1继续乘
  0.04 * 2 = 0.08       // 小于1，则继续乘
  0.08 * 2 = 0.16       // 小于1，则继续乘
  0.16 * 2 = 0.32       // 小于1，则继续乘（与第三行相同，这样会一直循环执行下去）
  ...
  
  将相乘之后等结果的整数部分拼接起来，所以 0.29的 二进制表示：010010100011110101110000101000111...
  ```

所以，最终39.29的二进制表示为：`100111.010010100011110101110000101000111...`



##### 第二步：科学计数法表示

`100111.010010100011110101110000101000111...`
$$
1.00111010010100011110101110000101000111... * 2^5
$$


##### 第三步：存储

以float32为例来进行存储，用32位来存储浮点型。

![image-20200608235249835](assets/image-20200608235249835.png)

- sign，用1位来表示浮点数正负，0表示正数；1表示负数。
- exponent，用8位来表示共有256种（0~255），含正负值（-127 ~ 128）。例如：5想要存储到exponent位的话，需要让 5 + 127 = 132，再讲132转换二进制，存储到exponent。（132的二进制是：01000010）
- fraction，存储小数点后的所有数据。



float64和float32类似，只是用于表示各部分的位数不同而已，其中：`sign=1位`、`exponent=11位`、`fraction=52位`，也就意味着可以表示的范围更大了。



#### 2.3 decimal

Go语言内部没有decimal。

第三方包，则需要在本地的Go环境中先安装再使用。第三方包源码地址：https://github.com/shopspring/decimal 。

##### 第一步：安装第三发的包

```
go get github.com/shopspring/decimal 
```

命令执行完成之后，在 `$GOPATH/src`的目录下就会出现 `github/shopspring/decimal`的目录，这就是第三方模块安排的位置。



##### 第二步：使用decimal包

```go
package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {

	var v1 = decimal.NewFromFloat(0.0000000000019)
	var v2 = decimal.NewFromFloat(0.29)

	var v3 = v1.Add(v2)

	var v4 = v1.Sub(v2)

	var v5 = v1.Mul(v2)

	var v6 = v1.Div(v2)
	
    fmt.Println(v3, v4, v5, v6) // 输出：0.2900000000019（也可以调用String方法）
    
    
	var price = decimal.NewFromFloat(3.4626)
	var data1 = price.Round(1) 		// 保留小数点后1位（四舍五入）
	var data2 = price.Truncate(1)	// 保留小数点后1位
	fmt.Println(data1, data2) // 输出：3.5  3.4
}
```



## 3.布尔类型

表示真假，一般是和条件等配合使用，用于满足某个条件时，执行某个操作。

```go
/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串 转换 布尔类型
	// true:"1", "t", "T", "true", "TRUE", "True"
	// false:"0", "f", "F", "false", "FALSE", "False"
	//false,err错误信息
	v1, err := strconv.ParseBool("t")
	fmt.Println(v1, err)

	// 布尔类型 转换 字符串
	v2 := strconv.FormatBool(false)
	fmt.Println(v2)
}

```



## 4.字符串

在编写程序时，使用字符串来进行文本的处理，例如：

```go
/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import "fmt"

func main() {
	var name string = "alex"
	fmt.Printf(name)

	title := "生活要想过得去，头上总得带点绿"
	fmt.Printf(title)
}
```



### 4.1 字符串的底层存储

计算机中所有的操作和数据最终都是二进制，即：1000100001011...

Go语言中的字符串是utf-8编码的序列。

```go
/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import "fmt"

func main() {
    // unicode字符集：文字 -> 码点（ucs4, 4个字节表示）
    // utf-8编码，对unicode字符集的码点进行编码最终得到：1000100001
	var name string = "武沛齐"
    fmt.Printf(name)
}
```

课上代码示例：

```go
/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	// 1. 本质是utf-8编码的序列
	var name string = "武沛齐"

	// 武 => 11100110 10101101 10100110
	fmt.Println(name[0], strconv.FormatInt(int64(name[0]), 2))
	fmt.Println(name[1], strconv.FormatInt(int64(name[1]), 2))
	fmt.Println(name[2], strconv.FormatInt(int64(name[2]), 2))

	// 武 => 11100110 10110010 10011011
	fmt.Println(name[3], strconv.FormatInt(int64(name[3]), 2))
	fmt.Println(name[4], strconv.FormatInt(int64(name[4]), 2))
	fmt.Println(name[5], strconv.FormatInt(int64(name[5]), 2))

	// 武 => 11101001 10111101 10010000
	fmt.Println(name[6], strconv.FormatInt(int64(name[6]), 2))
	fmt.Println(name[7], strconv.FormatInt(int64(name[7]), 2))
	fmt.Println(name[8], strconv.FormatInt(int64(name[8]), 2))

	// 2. 获取字符串的长度：9（字节长度）
	fmt.Println(len(name))

	// 3. 字符串转换为一个"字节集合"
	byteSet := []byte(name)
	fmt.Println(byteSet) // [230,173,166,230,178,155,233,189,144]

	// 4. 字节的集合转换为字符串
	byteList := []byte{230, 173, 166, 230, 178, 155, 233, 189, 144}
	targetString := string(byteList)
	fmt.Println(targetString)

	// 5. 将字符串转换为 unicode字符集码点的集合    6b66 ->  武  6c9b->沛   9f50->齐
	tempSet := []rune(name)
	fmt.Println(tempSet) // [27494 27803 40784]
	fmt.Println(tempSet[0], strconv.FormatInt(int64(tempSet[0]), 16))
	fmt.Println(tempSet[1], strconv.FormatInt(int64(tempSet[1]), 16))
	fmt.Println(tempSet[2], strconv.FormatInt(int64(tempSet[2]), 16))

	// 6. "rune集合" 转换 为字符串
	runeList := []rune{27494, 27803, 40784}
	targetName := string(runeList)
	fmt.Println(targetName)

	// 7. 长度的处理（获取字符长度）
	runeLength := utf8.RuneCountInString(name)
	fmt.Println(runeLength)
}
```



### 4.2 字符串常见功能

字符串属于在程序中最常见用的数据类型，所以Go中为字符串提供了很多常见的操作。

#### 4.2.1 获取长度

```go
package main

import (
   "fmt"
   "unicode/utf8"
)

func main() {
   var name string = "武沛齐"

   fmt.Println( len(name) )                    // 获取 字节 长度，输出：8
   fmt.Println( utf8.RuneCountInString(name) ) // 获取字符长度，输出：3
}
```

#### 4.2.2 是否以xx开头

```go
package main

import (
   "fmt"
   "strings"
)

func main() {
   name := "武沛齐"

   result := strings.HasPrefix(name, "武")

   fmt.Println(result) // 输出：true
}
```

#### 4.2.3 是否以xx结尾

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "武沛齐"

	result := strings.HasSuffix(name, "齐")

	fmt.Println(result) // 输出：true
}
```

#### 4.2.4 是否包含

```go
package main

import (
   "fmt"
   "strings"
)

func main() {
   name := "抬老子的意大利炮来"
   result := strings.Contains(name, "老子")

   fmt.Println(result) // 输出：true
}
```

#### 4.2.5 变大写

```go
package main

import (
   "fmt"
   "strings"
)

func main() {
   name := "wupeiqi"

   result := strings.ToUpper(name)

   fmt.Println(result) // 输出：WUPEIQI
}
// 注意：result是大写；name依然是小写。
```

#### 4.2.6 变小写

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "WUPEIQI"

	result := strings.ToLower(name)

	fmt.Println(result) // 输出：wupeiqi
}

```

#### 4.2.7 去两边

```go
package main

import (
   "fmt"
   "strings"
)

func main() {
   name := "wupeiqi"

   result1 := strings.TrimRight(name, "qi") // 去除右边的qi
   result2 := strings.TrimLeft(name, "w")   // 去除左边的w
   result3 := strings.Trim(name, "w")       // 去除两边的w

   fmt.Println(result1, result2, result3) // 输出：wupe upeiqi upeiqi
}
```

#### 4.2.8 替换

```go
package main

import (
   "fmt"
   "strings"
)

func main() {
   name := "wupeipeiqi"

   result1 := strings.Replace(name, "pei", "PE", 1)  // 找到pei替换为PE，从左到右找第一个替换
   result2 := strings.Replace(name, "pei", "PE", 2)  // 找到pei替换为PE，从左到右找前两个替换
   result3 := strings.Replace(name, "pei", "PE", -1) // 找到pei替换为PE，替换所有

   fmt.Println(result1, result2, result3)
}
```

#### 4.2.9 分割

```go
package main

import (
   "fmt"
   "strings"
)

func main() {
   name := "抬老子的意大利的炮来"
   result := strings.Split(name, "的")

   // 根据`的`进行切割，获取一个切片（类似于一个数组）
   fmt.Println(result) // [ 抬老子, 意大利, 炮来 ]
}
```

#### 4.2.10 拼接

可以使用 `+` 让两个字符串进行拼接，但这样的拼接效率会非常的低，不建议使用，建议大家使用以下的方式：

```go
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// 不建议
	message := "我爱" + "北京天安门"
	fmt.Println(message)

	// 建议：效率高一些
	dataList := []string{"我爱", "北京天安门"}
	result := strings.Join(dataList, "")
	fmt.Println(result) // 我爱北京天安门

	// 建议：效率更高一些（go 1.10之前）
	var buffer bytes.Buffer
	buffer.WriteString("你想")
	buffer.WriteString("我干")
	buffer.WriteString("他")
	data := buffer.String()
	fmt.Print(data)

	// 建议：效率更更更更高一些（go 1.10之后）
	var builder strings.Builder
	builder.WriteString("哈哈哈")
	builder.WriteString("去你的吧")
	value := builder.String()
	fmt.Print(value)
}
```

#### 4.2.11 string转换为int

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := "666"

	// 内部调用的就是 ParseInt
	var data, _ = strconv.Atoi(num)
	fmt.Println(data)
    
    // 整型转字符串（strconv.ParseInt 和 strconv.FormatInt 可用处理进制转换）
    // 十进制：整型； 其他进制：字符串形式 
    var result, err = strconv.ParseInt(num, 10, 32)
	fmt.Println(result, err)
}
```

#### 4.2.12 int转换为string

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var result = strconv.Itoa(888)
	fmt.Println(result)
}
```

#### 4.2.13 字符串 和 “字节集合”

```go
/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {

	var name string = "武沛齐"

	// 字符串转换为一个"字节集合"
	byteSet := []byte(name)
	fmt.Println(byteSet) // [230,173,166,230,178,155,233,189,144]

	// 字节的集合转换为字符串
	byteList := []byte{230, 173, 166, 230, 178, 155, 233, 189, 144}
	targetString := string(byteList)
	fmt.Println(targetString)

}
```



#### 4.2.14 字符串 和 “rune集合”

```go
/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {

	var name string = "武沛齐"

	// 将字符串转换为 unicode字符集码点的集合    6b66 ->  武  6c9b->沛   9f50->齐
	tempSet := []rune(name)
	fmt.Println(tempSet) // [27494 27803 40784]
	fmt.Println(tempSet[0], strconv.FormatInt(int64(tempSet[0]), 16))
	fmt.Println(tempSet[1], strconv.FormatInt(int64(tempSet[1]), 16))
	fmt.Println(tempSet[2], strconv.FormatInt(int64(tempSet[2]), 16))

	// "rune集合" 转换 为字符串
	runeList := []rune{27494, 27803, 40784}
	targetName := string(runeList)
	fmt.Println(targetName)
}
```



#### 4.2.15 string 和 字符

```go
/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 数字转字符串
	v1 := string(65)
	fmt.Println(v1) // A
    
    v2 := string(27494)
	fmt.Println(v2) // 武
    
	// 字符串转数字
	v3, size := utf8.DecodeRuneInString("A")
	fmt.Println(v3, size) // 65    1
    
    v4, size := utf8.DecodeRuneInString("武")
	fmt.Println(v4, size) // 27494 3

}
```

应用场景：生成一个随机数，然后调用string得到一个随机的字符。



### 4.3 索引切片和循环

```go
/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import "fmt"

func main() {
	var name string = "武沛齐"

	// 1. 索引获取字节
	v1 := name[0]
	fmt.Println(v1)  // 230

	// 2. 切片获取字节区间
	v2 := name[0:3]
	fmt.Println(v2)  // 武

	// 3. 手动循环获取所有字节
    /*
        0 230
        1 173
        2 166
        3 230
        4 178
        5 155
        6 233
        7 189
        8 144
    */
	for i := 0; i < len(name); i++ {
		fmt.Println(i, name[i]) 
	}

	// 4. for range 循环获取所有字符
    /*
        0 27494 武
        3 27803 沛
        6 40784 齐
	*/
	for index, item := range name {
        fmt.Println(index, item, string(item))
	}

	// 5.转换成rune集合 [27494,27803,40784]
	dataList := []rune(name)
	fmt.Println(dataList[0], string(dataList[0]))

}
```



## 5.数组

数组，定长且元素类型一致的数据集合。

```go
// 方式一：先声明再赋值（声明时内存中已开辟空间，内存初始化的值是0）
var numbers [3]int
numbers[0] = 999
numbers[1] = 666
numbers[2] = 333

// 方式二：声明+赋值
var names = [2]string{"武沛齐","alex"}

// 方式三：声明+赋值 + 指定位置
var ages = [3]int{0:87,1:73,2:99}

// 方式四：省略个数
var names = [...]string{"武沛齐","alex"}
var ages = [...]int{  0:87,  2:99 }
```

```go
// 声明 指针类型的数组（指针类型），不会开辟内存初始化数组中的值，numbers = nil
var numbers *[3]int

// 声明数组并初始化，返回的是 指针类型的数组（指针类型）
numbers := new([3]int)
```

### 5.1 数组内存管理

![image-20200610112239589](assets/image-20200610112239589.png)



数组，定长且元素类型一致的数据集合。

必备知识点：

- 数组的内存是连续的。

- 数组的内存地址实际上就是数组第一个元素的内存地址。

- 每个字符串的内部存储：`len` + `str`

  ```go
  type stringStruct struct {
  	str unsafe.Pointer
  	len int
  }
  ```

示例1：

```go
/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import "fmt"

func main() {
	nums := [3]int8{11, 22, 33}

	fmt.Printf("数组的内存地址：%p \n", &nums)
	fmt.Printf("数组第1个元素的内存地址：%p \n", &nums[0])
	fmt.Printf("数组第2个元素的内存地址：%p \n", &nums[1])
	fmt.Printf("数组第3个元素的内存地址：%p \n", &nums[2])
}

>>> 输出
数组的内存地址：0xc00001604a 
数组第1个元素的内存地址：0xc00001604a 
数组第2个元素的内存地址：0xc00001604b 
数组第3个元素的内存地址：0xc00001604c 
```



示例2：

```go
/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import "fmt"

func main() {
	nums := [3]int32{11, 22, 33}

	fmt.Printf("数组的内存地址：%p \n", &nums)
	fmt.Printf("数组第1个元素的内存地址：%p \n", &nums[0])
	fmt.Printf("数组第2个元素的内存地址：%p \n", &nums[1])
	fmt.Printf("数组第3个元素的内存地址：%p \n", &nums[2])
}

>>> 输出
数组的内存地址：0xc0000b4004 
数组第1个元素的内存地址：0xc0000b4004 
数组第2个元素的内存地址：0xc0000b4008 
数组第3个元素的内存地址：0xc0000b400c 
```



示例3：

```go
/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import "fmt"

func main() {
	names := [2]string{"武沛齐", "alex"}
	fmt.Printf("数组的内存地址：%p \n", &names)
	fmt.Printf("数组第1个元素的内存地址：%p \n", &names[0])
	fmt.Printf("数组第2个元素的内存地址：%p \n", &names[1])

}

>>> 输出：
数组的内存地址：0xc000128020 
数组第1个元素的内存地址：0xc000128020 
数组第2个元素的内存地址：0xc000128030
```



### 5.2 可变和拷贝

可变，数组的元素可以被更改（长度和类型都不可以修改）。

```go
names := [2]string{"武沛齐", "alex"}
names[1] = "苑昊"
```

注意：字符串不可以被修改。  "武沛齐"   "武陪齐"



拷贝，变量赋值时重新拷贝一份。

```go
name1 := [2]string{"武沛齐", "alex"}
name2 := name1

name1[1] = "苑昊"

fmt.Println(name1,name2)   // [武沛齐 苑昊]   [武沛齐 alex]
```



### 5.3 长度索引切片和循环

```go
/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import "fmt"

func main() {
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
	nums := [3]int32{11, 22, 33}
	for key, item := range nums {
		fmt.Println(key, item)
	}

	for key := range nums {
		fmt.Println(key)
	}

	for _,item := range nums {
		fmt.Println(item)
	}

}
```

### 5.4 数组嵌套

```go
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
```



## 作业题（21道）

1. Go语言中int占多少字节？

2. 整型中有符号和无符号是什么意思？

3. 整型可以表示的最大范围是多少？超出怎么办？

4. 什么是nil？

5. 十进制是以整型方式存在，其他进制则是以字符串的形式存在？如何实现进制之间的转换？

6. 简述如下代码的意义

   ```go
   var v1 int
   var v2 *int
   var v3 = new(int)
   ```

7. 浮点型为什么有时无法精确表示小数？

8. 如何使用第三方包 decimal？

9. 简述 ascii、unicode、utf-8的关系。

10. 判断：Go语言中的字符串是utf-8编码的字节序列。

11. 什么是rune？

12. 判断：字符串是否可变？

13. 列举你记得的字符串的常见功能？

14. 字符串和 “字节集合” 、“rune集合” 如何实现相互转换？

15. 字符串如何实现高效的拼接？

16. 简述数组的存储原理。

17. 根据要求写代码

    ```go
    names := [3]string{"Alex","超级无敌小jj","傻儿子"}
    
    a. 请根据索引获取“傻儿子”
    b. 请根据索引获取“alex”
    c. 请根据索引获取“小jj”
    d. 请将name数组的最后一个元素修改为 “大烧饼”
    ```

18. 看代码写输出结果

    ```go
    var nestData [3][2]int
    fmt.Println(nestData)
    ```

19. 请声明一个有3个元素的数组，元素的类型是有两个元素的数组，并在数组中初始化值。例如：

    ```go
    [ 
        ["alex","qwe123"],
        ["eric","admin11"],
        ["tony","pp1111"] 
    ]
    ```

20. 循环如下数组并使用字符串格式化输出如下内容：

    ```go
    dataList := [ 
        ["alex","qwe123"],
        ["eric","admin11"],
        ["tony","pp1111"] 
    ]
    
    请补充代码...
    
    
    最终实现输出：
    我是alex,我的密码是qwe123。
    我是eric,我的密码是admin11。
    我是tony,我的密码是pp1111。
    ```

21. 补充代码实现用户登录

    ```go
    // userList表示有三个用户，每个用户有用户名和密码，例如：用户名是alex，密码是qwe213
    userList :=  [ 
        ["alex","qwe123"],
        ["eric","admin11"],
        ["tony","pp1111"] 
    ]
    
    // 需求：提示让用户输入用户名和密码，然后再userList中校验用户名和密码是否正确。
    ```

    









































































































