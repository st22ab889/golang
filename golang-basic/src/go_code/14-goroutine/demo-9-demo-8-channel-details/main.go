/*
channel使用细节和注意事项:
	(1).channel可以声明为只读,或者只写性质.
	(2).使用select可以解决从管道取数据的阻塞问题.
	(3).goroutine中使用recover,解决协程中出现panic,导致程序崩溃问题.
说明:
	如果起了一个协程,但这个协程出现了panic,如果没有捕获这个panic,就会造成整个程序崩溃.
	这时可以在goroutine中使用recover来捕获panic进行处理,这样即使这个协程发生问题,但是主线程仍然不受影响,可以继续执行.
*/

package main

import (
	"fmt"
	"time"
)


//管道可以声明为只读或者只写
func first () {

	// 1.在默认情况下,管道是双向
	// var chan1 chan int  // 可读可写

	// 2.声明为只写
	var chan2 chan <- int
	chan2 = make(chan int, 3)
	chan2 <- 20
	// num := <- chan2		// 读的话会发生error
	fmt.Println("chan2=", chan2)

	/*
	// 3.声明为只读
	var chan3 <- chan int
	num2 := <- chan3
	// chan3 <- 30 //err
	fmt.Println("num2", num2)
	*/
}


//使用select可以解决从管道取数据的阻塞问题
func second () {
	//1.定义一个管道 10个数据int
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	//2.定义一个管道 5个数据string
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	//传统方法在遍历管道时,如果不关闭会阻塞而导致 deadlock,问题是在实际开发中,可能不好确定什么时候关闭该管道,可以使用select 方式可以解决
	label:	// 标签尽量不要使用,但是没有办法也可以用
	for{
		select {
			//注意: 这里如果intChan一直没有关闭，不会一直阻塞而deadlock，会自动到下一个case匹配
		case v := <- intChan:
			fmt.Printf("从intChan读取的数据%d\n", v)
			time.Sleep(time.Second)
		case v := <- stringChan:
			fmt.Printf("从stringChan读取的数据%s\n", v)
			time.Sleep(time.Second)
		default:
			fmt.Printf("都取不到了，不玩了, 程序员可以加入逻辑\n")
			time.Sleep(time.Second)
			
			// return 			// 直接结束这个方法,return下面的代码就执行不到了
			// break			// 如果单独写一个break 没有用, 因为还是break在select这个区域,跳不出去
			break label			// break后,接下来的代码还可以执行到

			// 注意:这里推荐使用return
		}
	}

}


func sayHello() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world count = ", i)
	}
}

func test () {
	// 这里可以使用defer + recover
	defer func (){
		//捕获test抛出的panic
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()
	
	var myMap map[int]string	//定义了一个map
	myMap[0] = "Golang"
}

// goroutine中使用recover,解决协程中出现panic导致程序崩溃问题
func three() {
	go sayHello()
	go test()
	for i := 0; i < 10; i++{
		fmt.Println("main() ok=", i)
		time.Sleep(time.Second)		
	}
}

func main() {

	//first ()	// channel只读和只写的最佳实践案列

	//second ()	// 使用select可以解决从管道取数据的阻塞问题

	three()		// goroutine中使用recover,解决协程中出现panic导致程序崩溃问题
}