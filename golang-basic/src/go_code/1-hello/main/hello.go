/*
环境准备,三个环境变量:
GOROOT		D:\DevelopKit\go1.17.1.windows-amd64\go	  ==> 指定Go SDK安装目录
Path路径追加 %GOROOT%\bin							   ==> 指定sdk\bin目录,里面包含 go.exe、godoc.exe、gofmt.exe
GOPATH		D:\GoProject							  ==> 指定golang的工作目录,所有项目的源码都在这个目录下 
*/

// 表示该 hello.go 文件所在包是 main, 每个文件都必须归属一个包
// 这里包名不等同于文件夹名称,这里只是定义一个包名,"hello.go"这个源文件不一定非要在main这个文件夹下,文件夹名称和包名之间没有关系 
package main

// 表示引入一个名为 fmt 的包
import "fmt"

// func 是关键字,表示是个函数;
// main 是函数名,且是个主函数,也就是程序的入口
func main() {
	fmt.Println("hello golang!")
}

/*
C:\Users\WuJun>D:

D:\>cd D:\GoProject\src\go_code\1-hello\main

// 通过 go build 命令 对该 go 文件进行编译,生成 .exe 文件
D:\GoProject\src\go_code\1-hello\main>go build hello.go
// 运行 .exe 文件
D:\GoProject\src\go_code\1-hello\main>hello.exe
hello golang!
// 通过 go run 命令直接运行 hello.go 程序(其实还是会经过编译,这里只是表面隐藏了这一步), 这种方式类似执行一个脚本文件的形式
D:\GoProject\src\go_code\1-hello\main>go run hello.go
hello golang!

// 生成的可执行文件为"hello.exe"
D:\GoProject\src\go_code\1-hello\main>go build hello.go
// 生成的可执行文件为"myhello.exe"
D:\GoProject\src\go_code\1-hello\main>go build -o myhello.exe hello.go

在生产环境中,都是先编译再运行

两种执行流程的方式区别(面试可能文问到):
a.先编译成生成可执行文件,把该可执行文件拷贝到没有 go 开发环境的机器上,仍然可以运行
b.执行 go run go的源代码,如果要在另外一台机器上运行,需要go的开发环境
c.编译时,编译器会将程序运行依赖的库文件包含在可执行文件中,所以可执行文件变大了很多

go语言开发注意事项:
1. go源文件以"go"为扩展名
2. go应用程序的执行入口是 main() 方法
3. go语言严格区分大小写
4. go方法由一条条语句构成,每个语句后不需要分号(go语言会在每行自动加分号),这也体现了Golang的简洁性
5. go编译器是一行行进行编译的,因此一行就写一条语句,不要把多条语句写在同一行,否则报错
6. go语言定义的变量或者import的包, 如果没有使用到, 代码不能编译通过
7. 大括号都是成对出现的,缺一不可

*/



