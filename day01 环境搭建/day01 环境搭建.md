# day01 环境搭建

Go 和 C语言、C++、Python、Java 等一样都是编程语言。学习任何一门编程语言本质上都分3步走：

- 第一步：安装 `解释器 或 编译器`。
- 第二步：学相关编程语言语法，然后写代码。
- 第三步：用已安装`解释器 或 编译器` 去运行自己写的代码，这样代码就会去完成我们编写的功能了。

Go是编译型语言，所以我们一般称Go安装是`编译器`。

Go是开源且跨平台的一门编程语言，所以他支持window、linux、mac操作系统，同时也意味着我们可以在各种系统中安装Go的编译器。



## 1. Mac系统

### 1.1 下载Go编译器

https://golang.google.cn/

### 1.2 点击安装

<img src="assets/image-20200530225723360.png" alt="image-20200530225723360" style="zoom: 50%;float:left" />

默认的安装目录：`/usr/local/go/`

编译器启动文件：`/usr/local/go/bin/go`

### 1.3 配置环境PATH

```
export PATH=/usr/local/go/bin:$PATH
```

### 1.4 其他配置

#### 1.4.1 创建一个任意目录

此目录以后放你写的所有go代码。

```
/Users/wupeiqi/GolangProjects/
- bin,通过go install编译时候，生成的可执行文件。
- pkg,通过go install编译时候，生成的包文件。
- src,放我们以后编写的所有go代码和依赖。
	- crm
		- app.go
	- luffcity
		- xx.go
```

#### 1.4.2 配置环境变量

```
// Go安装目录
export GOROOT=/usr/local/go
// 代码和编译之后的文件相关代码
export GOPATH=/Users/wupeiqi/GolangProjects
// 存放编译之后的文件
export GOBIN=/Users/wupeiqi/GolangProjects/bin
```



### 1.5 环境变量“持久化”

vim ~/.bash_profile

```
export PATH=/usr/local/go/bin:$PATH
export GOROOT=/usr/local/go
export GOPATH=/Users/wupeiqi/GolangProjects
export GOBIN=/Users/wupeiqi/GolangProjects/bin
```





### 1.6 编写go代码

```
$GOPATH
├── bin
├── pkg
└── src
    └── crm
        └── app.go
```

```go
package main

import "fmt"

func main() {
    fmt.Println("叫爸爸")
}
```



### 1.7 运行

本质上就是让Go编译器去运行咱们编写的代码。



方式1：

```go
// 先进入项目目录
go run app.go
```



方式2（推荐）：

```go
// 先进入项目目录

// 编译代码
go build
// 运行
./crm
```



方式3：

```go
// 先进入项目目录
go install 

// 去bin运行
./crm
```

```
$GOPATH
├── bin
│   └── crm
├── pkg
└── src
    └── crm
        └── app.go
```





## 2. Linux系统

### 2.1 下载Go编译器

https://golang.google.cn/

### 2.2 安装

```
安装目录
/opt/go
```

启动Go编译器文件：`/opt/go/bin/go`

### 2.3 配置环境变量PATH

```
export PATH=/opt/go/bin:$PATH
```

### 2.4 其他配置

#### 2.4.1 创建一个任意目录

存放咱们编写的所有项目代码，编译文件之后存放编译后的文件。

```
/home/wupeiqi/GolangProjects/
- bin,在执行go install 命令，生成的可执行文件的目录。
- pkg,在执行go install 命令，存放生成的包文件。
- src,以后编写的所有Go代码都会放在这个目录。
	- crm
		- app.go
	- luffy
		- lk.go
```

#### 2.4.2 环境变量的配置

```
export GOROOT=/opt/go
export GOPATH=/home/wupeiqi/GolangProjects
export GOBIN=/home/wupeiqi/GolangProjects/bin
```



### 2.5 环境变量的“持久化”

vim ~/.bash_profile

```go
export PATH=/opt/go/bin:$PATH
export GOROOT=/opt/go
export GOPATH=/home/wupeiqi/GolangProjects
export GOBIN=/home/wupeiqi/GolangProjects/bin
```



### 2.6 编写go代码

```
/home/wupeiqi/GolangProjects（简写$GOPATH）
├── bin
├── pkg
└── src
    └── crm
        └── app.go
```

```go
package main
import "fmt"
func main() {
    // 调用Println函数在屏幕输出：叫爸爸
    fmt.Println("叫爸爸")
}
```



### 2.7 运行代码

本质上将写好的代码交给Go编译器去执行。

方式1：

```
// 进入项目目录
go run app.go
```



方式2（推荐）：

```
// 进入项目目录

// 编译代码并生成一个可执行文件
go build  

// 运行
./crm
```



方式3：

```
// 进入项目目录

// 编译代码，把编译之后的结果会放在 bin/pkg目录
go install 

// 运行
./crm
```

```
├── bin
│   └── crm
├── pkg
└── src
    └── crm
        └── app.go
```

Go程序员的项目：

- 产出一个可执行文件。
- 产出一个包文件。



## 3. Windows系统



### 3.1 下载Go编译器

https://golang.google.cn/



### 3.2 点击安装

建议安装：`C:\Go`



### 3.3 环境变量PATH

以便于以后运行GO编译器时，无需再写路径。



### 3.4 其他配置



#### 3.4.1 创建一个任意目录

以后咱们的go项目都要按照要求放在这个目录。

```
Y:\GolangProjects
 - bin,go install在编译项目时，生成的可执行文件会放到此目录。
 - pkg,go install在编译项目时，生成的包文件会放到此目录。
 - src,以后所有的项目都要放在这个目录。
 	- crm
 		- app.go
	- luffy
		- xx.go
```

#### 3.4.2 环境变量配置

![image-20200531115559729](assets/image-20200531115559729.png)



### 3.5 编写代码

```
Y:\GolangProjects
 - bin
 - pkg
 - src,以后所有的项目都要放在这个目录。
 	- crm
 		- app.go
```

```go
package main

import "fmt"

func main() {
    fmt.Println("叫爸爸")
}
```



### 3.6 运行

本质上就是把Go代码交给Go编译器去执行。

方式1：

```
// 进入项目目录
go run app.go
```

方式2（推荐）：

```
// 进入项目目录
go build

crm.exe
```

方式3：

```
// 进入项目目录
go install

执行 crm.exe
```

```
Y:\GolangProjects
 - bin
 	- crm.exe
 - pkg
 	- windows_amd64
 		- utils.a
 - src,以后所有的项目都要放在这个目录。
 	- crm
 		- app.go
 	- utils
 		- page.go
```

平时开发：

- 开发可执行文件，用来让用户使用。
- 开发一个包文件，其他项目来进行调用。



## 总结

首先要去下载Go编译器，然后进行安装，在安装目录下就是GO编译器相关的所有内容。

```
mac:     /etc/local/go/
linux:   /opt/go/
windows: C:\go\
```

在安装目录下有 bin目录中有一个go可执行文件，基于他来启动编译器。

- 直接通过路径找到可执行文件去运行（麻烦）
- 将 `/etc/local/go/bin` 目录添加PATH环境变量中。

那么在终端就可以执行执行`go version`，调用咱们安装的编译器。



如果想要正确的使用go编译器，还需做一些相关的配置（其他语言）。

- 创建目录，用于存放项目代码、编译后的可执行文件、编译后的包文件。

  ```
  xxxx
  - bin
  - pkg
  - src
  	- crm
  		app.go
  ```

- 环境变量

  ```
  GOROOT,GO编译器安装目录。
  GOPATH，用于存放项目代码、编译后的可执行文件、编译后的包文件（go 1.11版本后，go mod）。
  GOBIN，编译后的可执行文件存放的目录。
  ```

  

编写代码，然后进行。

写了两个项目：

- crm，编译后生成一个可执行文件。
- utils，编译后生成一个包文件。

运行项目

- go run，运行项目代码，内部会先编译并将编译后的文件放在系统的临时目录，然后再自动执行。
- go build，运行项目代码，手动编译并生成一个可执行文件，然后再自动执行。
- go install ，生成可执行文件 + 包文件，并且会将编译后的文件放在bin/pkg目录。



## 4.开发工具

- Goland，IDE（继承开发环境）【收费】
- vscode，编辑器 + 第三发组件  【免费】



### 4.1 下载Goland

https://www.jetbrains.com/go/



### 4.2 配置

- 字体
- 参数提示



### 4.3 项目开发

- 新项目
- 打开老项目

注意：项目放在 `$GOPATH/src`目录。







## 参考博客

https://pythonav.com/wiki/detail/4/42/







































































































