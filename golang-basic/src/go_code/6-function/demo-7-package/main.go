/*
包能解决的问题: 在实际开发中,往往需要调用不同文件中定义的函数,通过包就能实现
包的本质实际上就算创建不同的文件夹,来存放程序文件
Golang的每一个文件都是属于一个包,也就说说Golang是以包的形式来管理文件和项目目录结构的

包的三大作用:
	1.区分相同名字的函数、变量等标识符
	2.当程序文件很多时,可以很好的管理项目
	3.控制函数、变量等访问范围,既作用域

包的相关说明:
	打包基本语法		package util
	引入包的基本语法	import  "包的路径"

包的注意事项和细节说明:
	1.在给一个文件打包时,该包对应一个文件夹,文件的包名通常和文件所在的文件夹名一致,一般为小写字母
	2.当一个文件要使用其它函数或变量时,需要先引入对应的包
		引入方式1:
			import "包名"
		引入方式2:
			import{
				"包名"
				"包名"
			}
	3.package指令在文件第一行,然后是import指定
	4.在import包时,路径从 $GOPATH 的 src 开始,不用带 src,编译器会自动从 src 下开始引入
	5.为了让其它包的文件可以访问到本包的函数,则该函数名的首字母需要大写,类似其它语言的public,这样才能跨包访问
	6.在访问其它包函数时,其语法是 包名.函数名
	7.如果包名较长,Golang支持给包取别名. 注意细节：取别名后,原来的包就不能使用了
	8.同一包下不能有相同的函数名,也不能有相同的全局变量名,否则报重复定义
	9.如果你要编译成一个可执行程序文件,就需要将这个报声明为main, 即 package main. 这个就是一个语法规范,如果是写一个库,包名可以自定义
		// 使用"go build"编译指令时,编译的路径不需要带src,编译器会自动带
		// 编译时需要编译main报所在的文件夹
		// 编译后生成一个有默认名的可执行文件在 $GOPATH 目录下
		
		// 问题1: 生成了一个 demo-2-package.exe 文件在 C:\Users\WuJun 目录下,但是并没有在 $GOPATH 目录下 
		// 问题2: 在 GoProject\pkg 目录下,并没有生成"pkg\windows_amd64\go_code\6-function\demo-2-package\utils\util.a"这样一个文件
		C:\Users\WuJun>go build go_code\6-function\demo-2-package\
		//如下: 也可以指定名字和目录,如下表示生成一个"main.exe"文件放在D盘根目录下
		C:\Users\WuJun>go build -o D:\main.exe go_code\6-function\demo-2-package\

		// 这样生成的 demo-2-package.exe 文件在 $GOPATH 目录下
		D:\GoProject>go build go_code\6-function\demo-2-package\
		// 生成的 my.exe 文件在 D:\GoProject\bin 目录下
		D:\GoProject>go build -o bin\my.exe  go_code\6-function\demo-2-package\
*/

package main

import(
	"fmt"
	// util 就算包utils的别名,给包名取别名后,原来的包就不能使用了
	util "go_code/6-function/demo-7-package/utils"
)

func main() {
	fmt.Println("90.0 + 9.0 = ", util.Cal(90.0, 9.0, '+'))
	fmt.Println("108.0 - 9.0 = ", util.Cal(108.0, 9.0, '-'))
	fmt.Println("90.0 * 9.0 = ", util.Cal(90.0, 9.0, '*'))
	fmt.Println("90.0 / 9.0 = ", util.Cal(90.0, 9.0, '/'))
	fmt.Println("90.0 & 9.0 = ", util.Cal(90.0, 9.0, '&'))
}