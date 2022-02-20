/*
channel(管道)-基本介绍:
1.为什么需要channel,前面使用全局变量加锁同步来解决goroutine的通讯不完美.
	(1)主线程在等待所有goroutine全部完成的时间很难确定,设置主线程的等待时间,这仅仅是估算.
	(2)如果主线程休眠时间长了,会加长等待时间,如果等待时间短了,可能还有goroutine处于工作状态,这时也会随主线程的退出而销毁.
	(3)通过全局变量加锁同步来是实现,也并不利用多个协程对全局变量的读写操作.
	(4)基于以上点所以要用channel这种通讯机制.
2.channel的介绍
	(1)channel本质就是一个数据结构--队列,数据是先进先出(FIFO:first in first out).
	(2)线程安全,多goroutine访问时,不需要加锁,就是说channel本身就是线程安全的.
	(3)channel是有类型的,一个string的channel只能存放string类型数据.
3.channel(管道)-基本使用
	3.1 定义/声明channel,格式: var 变量名 chan 数据类型 
		举例:
			var intChan chan int 			// intChan用于存放int数据
			var mapChan chan map[int]string	// mapChan用于存放map[int]string类型
			var perChan chan Person
			var perChan chan *Person
		说明:
			(1):channel是引用类型
			(2):channel必须初始化才能写入数据,既make后才能使用
			(3):管道是有类型的,intChan只能写入整数int
	3.2 channel初始化:
		说明:使用make进行初始化
			var intChan chan int
			intChan = make(chan int, 10)
		向channel中写入(存放)数据
			var intChan chan int
			intChan = make(chan int, 10)
			num := 999
			intChan <- 10
			intChan <- num
4.channel(管道)-基本使用-channel使用的注意事项
	(1).channel中只能存放指定的数据类型
	(2).channel的数据放满后,就不能再放入了
	(3).如果从channel取出数据后,可以继续放入
	(4).在没有使用协程的情况下,如果channel数据取完了再取,就会报 dead lock

5.channel的遍历和关闭
	(1).channel的关闭:使用内置函数close可以关闭channel,当channel关闭后,就不能再向channel写数据了,但是仍然可以从该channel读取数据.
	(2).channel的遍历:channel支持 for-range 的方式进行遍历,注意两个细节.
		在遍历时,如果channel没有关闭,会出现 dead lock 的错误.
		在遍历时,如果channel已经关闭,则会正常遍历数据,遍历完后,就会退出遍历.
*/

package main

import (
	"fmt"
)

// 演示管道的使用
func first () {
	// 创建一个可以存放2个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)

	// 看看intChan是什么,可以看出intChan是一个地址, 说明intChan是个引用类型.
	fmt.Printf("intChan 的值=%v intChan本身的地址=%p\n", intChan, &intChan)  // intChan 的值=0xc0000d4080 intChan本身的地址=0xc0000ce018

	// 向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num

	// 看看管道的长度和cap(容量)
	fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan))	// channel len= 2 cap=3 

	// 注意: 当给管写入数据时,不能超过其容量
	intChan <- 50
	// intChan <- 98  	// 超过其容量发生错误:  fatal error: all goroutines are asleep - deadlock!
	fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan))

	// 从管道中读取数据
	var num2 int
	num2 = <- intChan
	fmt.Println("num2=", num2)
	fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan))  // 2, 3

	// 在没有使用协程的情况下,如果管道数据已经全部取出，再取就会报告 deadlock
	num3 := <- intChan
	num4 := <- intChan
	// <- intChan				// 注意:取的时候也可以不用变量接收取出来的值
	// num5 := <-intChan		// fatal error: all goroutines are asleep - deadlock!
	fmt.Println("num3=", num3, "num4=", num4/*, "num5=", num5*/)
}

type Cat struct{
	Name string
	Age int
}

func second () {

	// 1.创建一个 mapChan,最多可以放10个map[string]string的key-val结构体变量
	// var mapChan chan map[string]string
	// mapChan = make(chan  map[string]string, 10)

	// 2.创建一个 catChan,最多可以放10个Cat结构体变量
	// var catChan chan Cat
	// catChan =  make(chan Cat, 10)

	// 3.创建一个 catChan2,最多可以放10个*Cat结构体变量
	// var catChan chan *Cat
	// catChan =  make(chan *Cat, 10)	

	// 4.创建一个allChan,最多可以放10个任意数据类型变量
	// var callChan chan interface{}
	// callChan = make(chan interface{}, 10)


	// 5.看下面的代码会输出什么?
	//var allChan chan interface{}			//定义一个存放任意数据类型的管道 3个数据
	allChan := make(chan interface{}, 3)

	allChan <- 10
	allChan <- "tom jack"
	cat := Cat{"小花猫", 6}
	allChan <- cat

	//希望获得到管道中的第三个元素，则先将前2个推出
	<-allChan
	<-allChan

	newCat := <- allChan //从管道中取出的Cat是什么?
	fmt.Printf("newCat=%T , newCat=%v\n", newCat, newCat)

	// 下面的写法是错误的!编译不通过
	// fmt.Printf("newCat.Name=%v", newCat.Name)	// .\main.go:122:37: newCat.Name undefined (type interface {} is interface with no methods)

	// 使用类型断言
	a := newCat.(Cat) 
	fmt.Printf("newCat.Name=%v", a.Name)
}


// channel的遍历和关闭案例
func three () {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan)		// 这时不能再写入数据到channel
	// intChan <- 300	// 写入数据会发生错误
	fmt.Println("<<<<<<<<<<<<<<<<<")

	n1 := <- intChan	//当管道关闭后,可以读取数据
	fmt.Println("n1=", n1)


	//遍历管道
	intChan2 := make(chan int, 100)
	for i:=0; i < cap(intChan2); i++ {
		intChan2 <- i * 2 	//放入100个数据到管道
	}

	/*
	//遍历管道不能使用普通的 for 循环
	for i := 0; i < len(intChan2); i++ {
	}
	*/

	//在遍历时,如果channel没有关闭,则会出现 deadlock 的错误
	//在遍历时,如果channel已经关闭,则会正常遍历数据,遍历完后,就会退出遍历
	close(intChan2)	// 如果不关闭会出现:  fatal error: all goroutines are asleep - deadlock!
	for v := range intChan2 {
		fmt.Println("v=", v)
	}

	/*
	channel遍历和关闭的案例小结:
	(1).遍历等价于从管道中取出数据，既: <- ch
	(2).注意需要close管道,否则会出现 deadlock
	(3).在 for range 管道时,当遍历到最后的时候,发现有管道关闭了,就结束从管道读取数据的遍历工作,正常退出.
	(4).在 for range 管道时,当遍历到最后的时候,发现有管道没有关闭,程序会认为可能有数据继续写入,因此就会等待,如果程序没有数据写入,就会出现死锁.
	*/
}

func main() {

	first ()

	second ()

	three ()
}