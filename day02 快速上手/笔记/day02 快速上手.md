# day02 快速上手

## 今日概要

- 初识包管理，知道项目中文件和文件、文件和文件夹之间关系。
- 输出，写代码，在go编译器运行时会在屏幕显示内容。
- 初识数据类型
  - 整型，数字。例如：1、2、3、4
  - 字符串，表示文本信息。例如：“如家”  "锦江之星"
  - 布尔类型，真假。例如： 1>2 、 "如家" == “家”

- 变量 & 常量，当做是昵称。
- 输入，让咱们用户输入内容。
- 条件语句，开发一个猜数次程序，用户输入数字与咱们定义的数字进行比较。



## 1.初识包管理

关于包管理的总结：

- 一个文件夹可以称为一个包。
- 在文件夹（包）中可以创建多个文件。
- 在同一个包下的每个为文件中必须指定 `包名称且相同`

重点：关于包的分类

- main包，如果是main包，则必须写一个main函数，此函数就是项目的入口（main主函数）。 编译生成的就是一个可执行文件。
- 非main包，用来将代码进行分类，分别放在不同的包和文件中。

![image-20200601001233922](assets/image-20200601001233922.png)

![image-20200601001735200](assets/image-20200601001735200.png)



注意：

- 调用其他包的功能，需要先 import 导入，然后再使用；调用自己包中的功能时，直接调用即可（无需导入）

- 文件中的函数首字母是小写，表示此函数只能被当前包内部文件调用。首字母是大写，则意味着任何包都可以调用。



## 2.输出

在终端将想要展示的数据显示出来，例如：欢迎登录、请输入用户名等。。。

- 内置函数
  - print
  - println
- fmt包（推荐）
  - fmt.Print
  - fmt.Println

扩展：进程里有 stdin/stdout/stderr 。

```go
package main

import "fmt"

func main() {
	// 基于内置函数（不推荐）
	//print("好吃不过饺子 \n")
	//print("好玩不过嫂子 \n")
	//println("好吃不过饺子")
	//println("好玩不过嫂子")

	// fmt包（推荐）
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
	// %.2f，占位符 小数（保留小数点后2位，四舍五入）
	// 百分比
	fmt.Printf("老汉开着%s,去接%s的媳妇，多少钱一次？%d块。嫂子给打折吧，%.2f怎么样？小哥哥包你100%%满意", "车", "老男孩", 100, 3.889)
}
```



## 3.注释

- 单行注释， //
- 多行注释， /*    */

快捷：contrl + ?



## 4.初识数据类型

- 整型，整数
- 字符串，文本
- 布尔型，真假

```go
package main

import "fmt"

func main() {
	// 整型
	fmt.Println(666)
	fmt.Println(6 + 9)
	fmt.Println(6 - 9)
	fmt.Println(6 * 9)
	fmt.Println(16 / 9) // 商
	fmt.Println(16 % 9) // 余数

	// 字符串类型，特点：通过双引号
	fmt.Println("武沛齐")
	fmt.Println("钓鱼要掉刀鱼，刀鱼到岛上钓")
	fmt.Println("alex" + "SB")
	//fmt.Println("alex" + 666)
	fmt.Println("alex" + "666")
	// 对比
	fmt.Println("1" + "2") // 结果："12"
	fmt.Println(1 + 2)     // 结果：3

	// 布尔类型，真假
	fmt.Println(1 > 2) // false  假
	fmt.Println(1 < 2) // true   真
	fmt.Println(1 == 2)
	fmt.Println(1 >= 2)
	fmt.Println("Alex" == "sb")

	// 超前
	if 2 > 1 {
		fmt.Println("叫爸爸")
	} else {
		fmt.Println("孙子")
	}

}
```



## 5.变量

可以理解为昵称。

- 声明 + 赋值

  ```go
  var sd string = "老男孩alex"
  fmt.Println(sd)
  
  var age int = 73
  fmt.Println(age)
  
  var flag bool = true
  fmt.Println(flag)
  ```

- 先声明后赋值

  ```go
  // 声明了一个字符类型变量 sd
  var sd string
  // 给sd变量赋值
  sd = "老男孩alex"
  fmt.Println(sd)
  ```



### 5.1 声明变量的意义？

- 编写代码省事

  ```
  // 文本，请将文本输出3次："伤情最是晚凉天，憔悴斯人不堪怜。"
  var message string = "伤情最是晚凉天，憔悴斯人不堪怜。"
  fmt.Println(message)
  fmt.Println(message)
  fmt.Println(message)
  ```

- 存储结果，方便之后使用

  ```
  // 存储结果，方便之后使用
  var v1 int = 99
  var v2 int = 33
  var v3 int = v1 + v2
  var v4 int = v1 * v2
  var v5 int = v1 + v2 + v3 + v4
  fmt.Println(v1, v2, v3, v4, v5)
  ```

- 存储用户输入的值

  ```go
  var name string
  fmt.Scanf("%s", &name) // 用户输入字符串并赋值给name变量
  if name == "alex" {
      fmt.Println("用户名输入正确")
  } else {
      fmt.Println("用户名输入失败")
  }
  ```

### 5.2 变量名要求

- 【要求】变量名必须只包含：字母、数字、下划线

  ```go
  var %1 string，错误
  var $ string，错误
  ```

- 【要求】数字不能开头

  ```go
  var 1 string  错误
  var 1name string  错误
  var _ string 正确
  ```

- 【要求】不能使用go语言内置的关键字

  ```
  var var string = "南通州北通州南北通州通南北"  错误
  ```

  ```
  break、default、func、interface、select、case、defer、go、map、struct、chan、else、goto、package、switch、const、fallthrough、if、range、type、continue、for、import、return、var
  ```

- 建议
  - 变量名见名知意：name/age/num ;  v1、v2、v3
  - 驼峰式命名：myBossName / startDate

练习题：

```go
var n1 int 
var data bool
var _9 string
```



### 5.3 变量简写

- 声明+赋值

  ```go
  var name string = "武沛齐"
  
  var name = "武沛齐"
  
  name := "武沛齐"  // 推荐
  ```

- 先声明再赋值

  ```go
  var name string
  var message string
  var data string
  
  var name,message,data string
  name = "武沛齐"
  message = "中奖了"
  data = "中了5000w"
  ```



因式分解，例如：声明5个变量，分别有字符串、整型

```go
var (
    name   = "武沛齐"
    age    = 18
    hobby  = "大保健"
    salary = 1000000
    gender string  // 只声明但不赋值，有一个默认： ""
    length int     // 只声明但不赋值，有一个默认： 0
    sb bool     // 只声明但不赋值，有一个默认： false
)
fmt.Println(name, age, hobby, salary, gender,length,sb)
```



扩展：go编译器会认为声明变量不使用 就是耍流氓。



### 5.4 作用域

如果我们定义了大括号，那么在大括号中定义的变量。

- 不能被上级使用。
- 可以在同级使用。
- 可以再子级使用。

```go
package main

import "fmt"

func main() {
	name := "武沛齐"
	fmt.Println(name)
	if true {
		age := 18
		name := "alex"
		fmt.Println(age)
		fmt.Println(name)
	}
	fmt.Println(name)
}
```



全局变量和局部变量

- 全局变量，未写在函数中的变量称为全局变量；不可以使用`v1:=xx`方式进行简化；可以基于因式分解方式声明多个变量；项目中寻找变量时最后一环。
- 局部变量，编写在{}里面的变量；可以使用任意方式简化；可以基于因式分解方式声明多个变量；

```go
package main

import "fmt"

// 全局变量（不能以省略的方式）
var school string = "老男孩IT教育" // 可以
//var school = "老男孩IT教育" 	 // 可以
//school := "老男孩IT教育"  		 // 不可以

var (
	v1 = 123
	v2 = "你好"
	v3 int
)

func main() {

	name := "武沛齐" // 局部变量
	fmt.Println(name)
	if true {
		age := 18      // 局部变量
		name := "alex" // 局部变量
		fmt.Println(age)
		fmt.Println(name)
	}
	fmt.Println(name)
	fmt.Println(school)
	fmt.Println(v1, v2, v3)
}
```



### 5.5 赋值及内存相关

示例1：

```go
name := "武沛齐"
```

![image-20200601222221725](assets/image-20200601222221725.png)

示例2：

```go
name := "武沛齐"
nickname := name
```

![image-20200601222404306](assets/image-20200601222404306.png)

注意：这一点与python不同。

```go
package main

import "fmt"

func main() {
	name := "武沛齐"
	nickname := name

	fmt.Println(name, &name)
	fmt.Println(nickname, &nickname)
}

```



示例3：

```go
name := "武沛齐"
nickname := name

name = "alex"
```

![image-20200601222800118](assets/image-20200601222800118.png)

```go
package main

import "fmt"

func main() {
	name := "武沛齐"
	nickname := name
	fmt.Println(name, &name)
	fmt.Println(nickname, &nickname)

	name = "alex"
	fmt.Println(name, &name)
	fmt.Println(nickname, &nickname)
}
```

#### 注意事项

使用int、string、bool这三种数据类型时，如果遇到变量的赋值则会拷贝一份。【值类型】

```go
name := "武沛齐"
nickname := name
```

提示：后续在数据类型中会详细再讲解。



## 阶段练习题

1. `package main`的作用？

2. 单行注释和多行注释分别是什么？

3. 下面声明变量的方式错误的有哪些？

   ```go
   var name = "武沛齐"
   var age int = 18
   hobby = "嫂子"
   nickName string = "一米八大高个"
   webSite := "秃头统治地球"
   ```

   ```go
   var age int = 18
   age = 19
   ```

4. 下面声明变量的方式正确的有哪些？

   ```go
   var v1,v2,v3 int
   var (
   	v6 = 123
       v7 string
   )
   var v4,v5 = 11,22
   v4,v5 = 11,22
   v4,v5 := 11,22
   ```

5. 变量名规范。

6. 看代码写结果

   ```go
   package main
   
   import "fmt"
   
   var number = 9999
   
   func main() {
   	number := 666
   	fmt.Println(number)
   }
   ```

7. 看代码写结果

   ```go
   package main
   
   import "fmt"
   
   var number = 99
   
   func main() {
   	number := 66
   	if true {
   		number := 33
   		fmt.Println(number)
   	}
   	fmt.Println(number)
   }
   ```

8. 看代码写结果

   ```go
   package main
   
   import "fmt"
   
   var number = 99
   
   func main() {
   	number := 66
   	if true {
   		number := 33
   		fmt.Println(number)
   	}
   	number = 88
   	fmt.Println(number)
   }
   ```

9. 看代码写结果

   ```go
   package main
   
   import "fmt"
   
   var number = 99
   
   func main() {
   	if true {
   		fmt.Println(number)
   	}
   	number = 88
   	fmt.Println(number)
   }
   ```

10. 看代码写结果

    ```go
    package main
    
    import "fmt"
    
    var number = 99
    
    func main() {
    	number := 88
    	if true {
    		number = 123
    		fmt.Println(number)
    	}
    	fmt.Println(number)
    }
    
    ```

11. 看代码写结果

    ```go
    package main
    
    import "fmt"
    
    var number = 99
    
    func main() {
    	fmt.Println(number)
    	number := 88
    	if true {
    		number = 123
    		fmt.Println(number)
    	}
    	fmt.Println(number)
    }
    ```

    

## 6.常量

不可被修改的变量。

```go
package main

import "fmt"

func main() {
	// 定义变量
	//var name string = "武沛齐"
	//var name = "武沛齐"
	name := "武沛齐"
	name = "alex"
	fmt.Println(name)

	// 定义常量
	//const age int = 98
	const age = 98
	fmt.Println(age)
}
```



### 6.1 因式分解

```go
package main

import "fmt"


func main() {
	// 定义变量
	//var name string = "武沛齐"
	//var name = "武沛齐"
	name := "武沛齐"
	name = "alex"
	fmt.Println(name)

	// 定义常量
	//const age int = 98
	const age = 98
	fmt.Println(age)

	// 常量因式分解
	const (
		v1 = 123
		v2 = 456
		pi = 9.9
	)
	fmt.Println(v1, v2, pi, gender)
}

```



### 6.2 全局

```go
package main

import "fmt"

const Data = 999
const (
	pi     = 3.1415926
	gender = "男"
)

func main() {
	// 定义变量
	//var name string = "武沛齐"
	//var name = "武沛齐"
	name := "武沛齐"
	name = "alex"
	fmt.Println(name)

	// 定义常量
	//const age int = 98
	const age = 98
	fmt.Println(age)

	// 常量因式分解
	const (
		v1 = 123
		v2 = 456
		//pi = 9.9
	)
	fmt.Println(v1, v2, pi, gender)
}

```



### 6.3 iota

可有可无，当做一个在声明常量时的一个计数器。

```go
package main

const (
	monday = iota + 1
	tuesday
	wednesday
	thursday
	friday
	saturday
	sunday
)

const (
	n1 = iota
	n2
	n3
)

func main() {

	// iota
	// 示例1
	/*
		const (
			v1 = 1
			v2 = 2
			v3 = 3
			v4 = 4
			v5 = 5
		)
		fmt.Println(v1, v2, v3, v4, v5)
	*/

	// 示例2
	/*
		const (
			v1 = iota
			v2
			v3
			v4
			v5
		)
		fmt.Println(v1, v2, v3, v4, v5)
	*/

	// 示例3
	/*
		const (
			v1 = iota + 2
			v2
			v3
			v4
			v5
		)
		fmt.Println(v1, v2, v3, v4, v5)
	*/

	// 示例4：
	/*
		const (
			v1 = iota + 2
			_
			v2
			v3
			v4
			v5
		)
		fmt.Println(v1, v2, v3, v4, v5)
	*/

}
```



## 7.输入

让用户输入数据，完成项目交互。

- fmt.Scan
- fmt.Scanln
- fmt.Scanf



```go
package main

import "fmt"

func main() {
	// 示例1：fmt.Scan
	/*
		var name string
		fmt.Println("请输入用户名：")
		fmt.Scan(&name)
		fmt.Printf(name)
	*/

	// 示例2：fmt.Scan
	var name string
	var age int

	fmt.Println("请输入用户名：")
	// 当使用Scan时，会提示用户输入
	// 用户输入完成之后，会得到两个值：count,用户输入了几个值；err，用输入错误则是错误信息
	_, err := fmt.Scan(&name, &age)
	if err == nil {
		fmt.Println(name, age)
	} else {
		fmt.Println("用户输入数据错误", err)
	}
	// 特别说明：fmt.Scan 要求输入两个值，必须输入两个，否则他会一直等待。
}

```



```go
package main

import "fmt"

func main() {
	// 示例1：fmt.Scanln
	/*
		var name string
		fmt.Print("请输入用户名：")
		fmt.Scanln(&name)
		fmt.Printf(name)
	*/

	// 示例2：fmt.Scanln

	var name string
	var age int
	fmt.Print("请输入用户名：")
	// 当使用Scanln时，会提示用户输入
	// 用户输入完成之后，会得到两个值：count,用户输入了几个值；err，用输入错误则是错误信息
	count, err := fmt.Scanln(&name, &age)
	fmt.Println(count, err)
	fmt.Println(name, age)

	// 特别说明：fmt.Scanln 等待回车。
}
```



```go
package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Print("请输入用户名：")
	_, _ = fmt.Scanf("我叫%s 今年%d 岁", &name, &age)
	fmt.Println(name, age)
}
```



无法解决的难题？

<img src="assets/image-20200602112623611.png" alt="image-20200602112623611" style="zoom:25%;" />

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	// line，从stdin中读取一行的数据（字节集合 -> 转化成为字符串）
	// reader默认一次能4096个字节（4096/3)
	//    1. 一次性读完，isPrefix=false
	// 	  2. 先读一部分，isPrefix=true，再去读取isPrefix=false
	line, _, _ := reader.ReadLine()
	data := string(line)
	fmt.Println(data)
}

```



## 8.条件语句



### 8.1 最基本

```go
if 条件 {
    成立后，此代码块执行
}else{
    不成立，此代码块执行
}
```

```go
if 条件 {
    成立后，此代码块执行
}
```

示例：

```go
package main

func main() {
	/*
		if true {
			fmt.Println("666")
		}else{
			fmt.Println("999")
		}
	*/

	/*
		if 1 > 2 {
			fmt.Println("666")
		} else {
			fmt.Println("999")
		}
	*/

	/*
		flag := false
		if flag {
			fmt.Println("条件成立")
		}else{
			fmt.Println("条件不成立")
		}
	*/

	// 练习题1:用户输入姓名，判断是否正确
	/*
		var name string
		fmt.Print("请输入姓名：")
		fmt.Scanln(&name)
		if name == "alex" {
			fmt.Println("用户名输入正确")
		} else {
			fmt.Println("用户名输入错误")
		}
	*/
	// 练习题2:用户输入数字，判断奇数、偶数
	/*
		var number int
		fmt.Print("请输入数字：")
		fmt.Scanln(&number)
		if number % 2 == 0{
			fmt.Println("您输入的是偶数")
		}else{
			fmt.Println("您输入的是奇数")
		}
	*/
	// 练习题3:用户和密码，判断用户名密码是否正确。
	/*


		var username, password string
		fmt.Print("请输入用户名：")
		fmt.Scanln(&username)

		fmt.Print("请输入密码：")
		fmt.Scanln(&password)

		if username == "alex" && password == "sb" {
			fmt.Println("欢迎登录pornhub")
		} else {
			fmt.Println("用户名或密码错误")
		}
	*/
	// 练习题4:请输入用户名校验是否是VIP
	/*
		var username string
		fmt.Print("请输入用户名：")
		fmt.Scanln(&username)

		if username == "alex" || username == "eric" {
			fmt.Println("天上人间大VIP")
		} else {
			fmt.Println("屌丝")
		}
	*/
}

```



### 8.2 多条件判断

```go
if 条件A{
    ...
}else if 条件B{
    ...
}else if 条件C{
    ...
}else{
    ...
}
```

示例：

```go
package main

import "fmt"

func main() {
	var length int
	fmt.Print("请输入你的长度：")
	fmt.Scanln(&length)

	if length < 1 {
		fmt.Println("没用的东西，还特么是坑")
	} else if length < 6 {
		fmt.Println("刚刚能用")
	} else if length < 18 {
		fmt.Println("生活和谐")
	} else {
		fmt.Println("太特么大了")
	}
}
```

### 8.3 嵌套

```go
package main

import "fmt"

func main() {
	fmt.Println("欢迎致电10086，1.话费相关；2.业务办理；3.人工服务。")

	var number int
	fmt.Scanln(&number)

	if number == 1 {
		fmt.Println("话费服务，1.交话费；2.查询。")
		var n1 int
		fmt.Scanln(&n1)
		if n1 == 1 {
			fmt.Println("缴话费啦")
		} else if n1 == 2 {
			fmt.Println("查话费了")
		} else {
			fmt.Println("输入错误")
		}
	} else if number == 2 {
		fmt.Println("业务办理")
	} else if number == 3 {
		fmt.Println("人工服务")
	} else {
		fmt.Println("输入错误")
	}
	
	// 建议：条件的嵌套不要太多
}
```





## 今日作业

1. 提⽰⽤户输入⿇花藤. 判断⽤户输入的对不对。如果对, 提⽰真聪明, 如果不对, 提⽰你 是傻逼么。

2. 提示用户输入两个数字，计算两个数的和并输出。

3. 提示用户输入姓名、位置、行为三个值，然后做字符串的拼接 并输出，例如：xx 在 xx 做 xx 。

4. 设定一个理想数字比如：66，让用户输入数字，如果比66大，则显示猜测的结果大了；如果比66小，则显示猜测的结果小了;只有等于66，显示猜测结果正确。

5. 写程序，输出成绩等级。成绩有ABCDE5个等级，与分数的对应关系如下.

   ```
   A    90-100
   B    80-89
   C    60-79
   D    40-59
   E    0-39
   ```

   要求用户输入0-100的数字后，正确输出他的对应成绩等级。



老师微信：wupeiqi666















