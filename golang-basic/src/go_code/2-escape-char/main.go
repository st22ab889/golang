/*
Golang常用的转义字符(escape char)
\t	:一个制表位,实现对齐的功能
\n	:换行符
\\	:一个\
\"	:一个"
\r	:一个回车.  fmt.Println("Go语言\r性能好")
*/

package main

// fmt包中提供格式化、输出、输入的函数
import "fmt"

func main() {
	fmt.Println("GO\tJava")
	fmt.Println("GO\nJava")
	fmt.Println("d:\\go\\project")
	fmt.Println("GO \"性能好\"")
	// \r 回车,从当前行的最前面开始输出,覆盖掉以前内容
	fmt.Println("GO\rJava")

	fmt.Println("80个字符换行\n",
		"80个字符换行\n",
		"80个字符换行\n")
}

/*
Golang开发常见问题和解决方法
	找不到文件 -> 源文件名不存在或者写错，当当前路径错误
	语法错误 -> 看编译器报告的错误信息
*/

/*
Go语言支持C风格的块注释,也支持C++风格的行注释(//),行注释更通用,块注释主要用于针对包的详细说明或者疲敝大块的代码

正确的注释和注释风格:
	Go官方推荐使用行注释来注释整个方法和语句
	带来Go源码

正确的缩进和空白:
	a.使用tab实现整体向右边移动,使用 shift+tab 整体向左移动
	b.使用gofmt命令来进行格式化
		// 格式化后只是在命令窗口显示出格式化后的文本,不会更新到文件
		D:\GoProject\src\go_code\2-escape-char>gofmt main.go
		// 格式化后更新到文件
		D:\GoProject\src\go_code\2-escape-char>gofmt -w main.go
	c.运算符两边习惯各加一个空格,比如 2 + 4 * 5
	d. 如下格式是错误的：
			func main()
			{
				// ...
			}
		正确的格式如下(因为go的设计思想是一个问题尽量只有一个解决方法):
			func main() {
				// ...
			}
	e.行长约定:一行最长不超过80个字符,超过的使用换行展示,尽量保持格式优雅
			fmt.Println("80个字符换行\n",
				"80个字符换行\n",
				"80个字符换行\n")

Go官方编程指南: 
	官方地址: 	https://golang.google.cn/
	标准库API文档(可以查看Golang所有包下的函数和使用): https://golang.google.cn/pkg/				

术语解释:
	API(application program interface,应用程序编程接口):是Golang提供的基本编程接口,就是Go的各个包的各个函数

Go语言提供了大量的标准库,也为这些标准库提供了相应的API文档,用来告诉开发者如何使用这些标准库以及标准库中的方法
Golang中文网,在线标准库文档: https://studygolang.com/pkgdoc	

Go的安装包中提供了源码,路径:  D:\DevelopKit\go1.17.1.windows-amd64\go\src

DOS(Disk Operating System 磁盘操作系统)命令
	dir			--> 查看目录
	D:			--> 切换到D盘
	cd \		--> 进入根目录	
	md test		--> 新建一个名为test的目录
	rd test 	--> 如果test这个目录是空目录,可以直接删除
	rd /q/s test 	--> 删除test目录以及下面的子目录和文件,不带询问
	rd /s test	 	--> 删除test目录以及下面的子目录和文件,带询问
	echo hello > d:\abc.txt --> 新建或追加内容到文件,如果文件存在,里面的内容会被覆盖
	echo hello >> d:\abc.txt --> 新建或追加内容到文件,如果文件存在,不会覆盖原文件内容,新的内容追加到原文件内容的后面
	copy abc.txt d:\test  	--> 拷贝文件名不变
	copy abc.txt d:\test\ok.txt  --> 拷贝文件名改变
	del abc.txt		--> 删除指定文件
	del *.txt
	cls		--> 清屏
	exit	--> 退出
*/
