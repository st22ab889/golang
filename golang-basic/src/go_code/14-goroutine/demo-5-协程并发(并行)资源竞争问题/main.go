/*
channel(管道):

需求:计算1-100的各个数的阶乘,并且把各个数的阶乘放入到map中,最后显示出来.要求使用goroutine完成.
	分析思路:使用goroutine来完成,效率高,但是会出现并发/并行安全问题,要解决不同goroutine如何通信的问题.
	代码实现:
		(1).使用goroutine来完成(看看使用goroutine并发完成会出现什么问题?)
		(2).在运行某个程序时,如果知道是否存在资源竞争问题.方法很简单,在编译该程序时,增加一个参数 -race 即可.
				go build -race test.go  	// race就是竞争的意思,加上这个参数表示程序资源如果有竞争的时候,会把信息打印出来
*/

package main

import (
	_ "fmt"
	"time"
)

var (
	myMap = make(map[int]int, 10)
)

// test 函数就是计算 n, 将这个结果放入到 myMap
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	//这里将 res 放入到myMap
	myMap[n] = res // 问题1: 出现 fatal error: concurrent map writes
}

func first () {
	// 这里开启多个协程完成这个任务[200个]
	for i := 1; i <= 200; i++ {
		go test(i)
	}

	//休眠10秒钟(等协程运行结束)
	time.Sleep(time.Second * 10)
}


func main() {

	/*
	当运行  go run .\main.go 会出现两个问题
		1.当多个协程同时向map中写入数据时,没有保护机制,会出现互斥问题(可以并发读,但是不能并发写)
		2.主线程要等协程多长时间
	当运行 go build -race .\main.go 和  main.exe ,运行后会发现 "Found X data race(s)"(X代表数字)
		"Found X data race(s)"(X代表数字) 表示写的时候有 X 个数据产生了竞争关系,往里面同时写出现了问题.
		注意:路径中有中文,执行"main.exe"会出现 "cannot find package ..." 的问题
	*/
	first ()
}