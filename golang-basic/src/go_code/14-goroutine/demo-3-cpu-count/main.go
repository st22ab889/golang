/*
1.设置Golang运行的CPU数.介绍:为了充分利用多CPU的优势,在Golang程序中,设置运行的CPU数目.
	(1). Go1.8后,默认让程序运行在多个核上,可以不用设置
	(2). Go1.8前,还是要设置一下,可以更高效的利用CPU

2. 相关函数
	(1). NumCPU返回本地机器的逻辑CPU个数.
		func NumCPU() int 

	(2). GOMAXPROCS设置可同时执行的最大CPU数,并返回先前的设置.若 n < 1,它就不会更改当前设置.
		func GOMAXPROCS(n int) int	

*/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNum)

	//可以自己设置使用多个cpu
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}