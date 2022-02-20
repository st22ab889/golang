/*
goroutine和channel结合:
1. 应用实例1: goroutine 和 channel 协同工作的案列,具体要求:
	(1).开启一个writeData协程,向管道intChan中写入50个整数.
	(2).开启一个readData协程,从管道intChan中读取writeData写入的数据.
	(3).注意:writeData和readData操作的是同一个管道,主线程需要等待writeData和readData协程都完成工作才能退出[管道]

2.管道阻塞的机制:应用实例2-阻塞
	如果编译器(运行)发现一个管道只有写而没有读,则该管道会阻塞.写管道和读管道的频率不一致无所谓.

3.应用实例3
	需求:要求统计1-8000的数字中,哪些是素数?
	分析思路
		传统方法:使用一个循环,循环的判断各个数是不是素数.
		使用并发/并行的方式:将统计素数的任务分配给多个(4个)goroutine去完成,完成任务时间短.
	说明:
		质数又称素数。一个大于1的自然数，除了1和它自身外，不能被其他自然数整除的数叫做质数；否则称为合数（规定1既不是质数也不是合数）
		使用goroutine完成后,可以在使用传统的方法来统计一下,看看完成这个任务,各个耗费的时间是多少?[用map保存primeNum]
	结论:使用go协程后,执行速度比普通方法提高至少4倍.
*/

package main

import (
	"fmt"
	"time"
)

// write data
func writeData (intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i	//放入数据
		fmt.Println("writeData ", i)
	}
	close(intChan)
}

// read Data
func readData (intChan chan int, exitChan chan bool) {
	for{
		v, ok := <- intChan
		if !ok {
			break
		}
		fmt.Printf("readData 读到数据=%v\n", v) 
	}

	//readData 读取完数据后，即任务完成
	exitChan <- true
	close(exitChan)
}


func first () {
	//创建两个管道
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)
	
	go writeData(intChan)
	go readData(intChan, exitChan)

	// 就是为了等待 readData 协程完成
	for{
		_, ok := <- exitChan
		if !ok {
			break
		}
	}
}


func second () {
	// 如果注销掉 first () 中的 go readData(intChan, exitChan),程序会怎样?
	// 答:如果只是向管道写入数据,而没有读取,就会出现阻塞而 dead lock,原因是 intChan容量是10,而代码 writeData 会写入50个数据,因此会阻塞在 writeData 的 ch <- i
}


//向 intChan 中放入1至naturalNumber个数
func putNum (intChan chan int, naturalNumber int){
	for i:= 0; i < naturalNumber; i++ {
		intChan <- i
	}
	close(intChan)		// 关闭 intChan
}

// 从 intChan取出数据，并判断是否为素数,如果是，就放入到primeChan
func primeNum (intChan chan int, primeChan chan int, exitChan chan bool){
	var flag bool
	for { 							//使用for 循环
		num, ok := <- intChan
		if !ok {					// intChan 取不到
			break
		}
		flag = true 				//假设是素数
		for i := 2; i < num; i++ {
			if num % i == 0 {		// 说明该num不是素数
				flag = false
				break
			}
		}

		if flag {
			primeChan <- num	// 将这个数就放入到 primeChan
		}
	}

	fmt.Println("有一个primeNum 协程已经完成工作,不能再取到数据,退出")
	exitChan <- true	//这里还不能关闭 primeChan, 向 exitChan 写入true
}


const naturalNumber int = 8000	// 自然数
const goroutineCount int = 4	// goroutine个数

func three () {

	intChan := make(chan int, 1000)				// 放自然数 channel
	primeChan := make(chan int, 1000)			// 放素数的channel
	exitChan := make(chan bool, goroutineCount)	// 标识退出的管道

	start := time.Now().Unix()

	//开启一个协程,向 intChan放入 1-8000个数
	go putNum(intChan, naturalNumber)

	//开启4个协程，从 intChan取出数据，并判断是否为素数,如果是，就放入到primeChan
	for i := 0; i < goroutineCount; i++{
		go primeNum(intChan, primeChan, exitChan)
	}

	//这里主线程进行处理,直接
	go func ()  {
		for i:= 0; i < goroutineCount; i++{
			<- exitChan
		}
		
		end := time.Now().Unix()
		fmt.Println("使用协程耗时=", end - start)

		//当从exitChan 取出了4个结果，就可以放心的关闭 prprimeChan
		close(primeChan)
	}()

	//遍历 primeChan ,把结果取出
	for {
		res, ok := <- primeChan
		if !ok {
			break
		}
		//将结果输出
		fmt.Printf("素数=%d\n", res)
	}	

	fmt.Println("main线程退出")
}

func main() {

	// first () 	// 应用实例1: goroutine 和 channel 协同工作的案列

	// second () 	// 应用实例2: 管道阻塞的机制

	three ()	// 应用实例3: 要求统计1-8000的数字中,哪些是素数?
}