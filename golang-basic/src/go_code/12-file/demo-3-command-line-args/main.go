/*
命令行参数:
	1) 基本介绍: os.Args是一个string的切片,用来存储所有的命令行参数
		func ParseParameters () {
			fmt.Println("参数个数 =", len(os.Args))
			for i, v := range os.Args {
				fmt.Printf("args[%v] = %v \n",i , v)
			}
		}

	2) flag包用来解析命令行参数
		说明:前面的方式是比较原生的方式,对解析参数不是特别的方便,特别是带有指定参数形式的命令行.
		比如: cmd>main.exe -f c:/abc.txt -p 200 -u root 这样的形式命令行,Golang设计者们提供了flag包,可以方便的解析命令行参数,而且参数顺序可以随意.

*/

package main

import (
	"fmt"
	"os"
	"flag"
)

// 获取到命令行输入的各种参数
func first () {

	fmt.Println("参数个数 =", len(os.Args))
	// 遍历 os.Args 切片就可以得到所有的命令行输入的参数值
	for i, v := range os.Args {
		fmt.Printf("args[%v] = %v \n",i , v)
	}
	
}


func second () {
	
	//定义变量,用于接收命令行的参数值
	var user string
	var pwd string
	var host string
	var port int

	/*
	func StringVar(p *string, name string, value string, usage string)
	func IntVar(p *int, name string, value string, usage string)
	参数说明:
		p 是接收用户命令行中输入的 -u 后面的参数值
		name 就是 -u 指定参数
		value 表示默认值
		usage 表示说明
	*/
	flag.StringVar(&user, "u", "", "用户名,默认为空")	// 注册一个string类型flag	
	
	flag.StringVar(&pwd, "pwd", "", "密码,默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名,默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号,默认为3306")	// 注册一个int类型flag

	// func Parse() 表示从 os.Args 中解析注册的flag.必须在所有flag都注册好而未访问其值时执行. 未注册却使用flag时,会返回err
	flag.Parse()

	fmt.Printf("user=%v pwd=%v host=%v port=%v", user, pwd, host, port)
}


func main() {
	
	/*
	使用"go run .\main.go aa bb cc dd"命令运行结果如下:
		参数个数 = 5
		args[0] = C:\Users\WuJun\AppData\Local\Temp\go-build980242719\b001\exe\main.exe
		args[1] = aa
		args[2] = bb
		args[3] = cc
		args[4] = dd
	*/
	// first ()	// 获取到命令行输入的各种参数



	/*
	使用"go run .\main.go -u root -pwd 123456 -h 127.0.0.1 -port 3306" 命令运行结果如下:
		user=root pwd=123456 host=127.0.0.1 port=3306
	*/
	second ()	// flag包用来解析命令行参数

}