/*
channel(管道)-基本介绍:
	1.不同goroutine之间如何通讯
		(1) 全局变量加锁同步
		(2) 使用管道 channel 来解决
	2.使用全局变量加锁同步改进程序
		(1) 因为没有对全局变量 m 加锁,因此会出现资源争夺问题,代码会出现错误,提示 concurrent map writes. 解决方案: 加入互斥锁
		(2) 如果读比较多,则使用读写锁
	3.go有两种锁，读写锁和互斥锁. mutex:互斥; sync:同步(synchronized)
*/

package main

import (
	"fmt"
	"time"
	"sync"
)

var (
	myMap = make(map[int]int, 10)

	// 声明一个全局的互斥锁(变量)lock; sync是包(synchornized,同步的意思); Mutex:是互斥的意思
	lock sync.Mutex
)

// test 函数就是计算 n, 将这个结果放入到 myMap
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res += i	// 如果数的阶乘很大,结果会越界,可以将求阶乘改成 sum += uint64(i)
	}
	lock.Lock()		// 加锁
	myMap[n] = res 	// 这里将 res 放入到myMap
	lock.Unlock()	// 解锁
}

func first () {
	// 这里开启多个协程完成这个任务[200个]
	for i := 1; i <= 200; i++ {
		go test(i)
	}

	//休眠10秒钟
	time.Sleep(time.Second * 10)

	/*
	这里需要加互斥锁的原因是: 按理说10秒后上面的协程都应该执行完,后面不应该出现资源竞争问题,但是在实际运行中,还是可能在这里出现,
	因为程序从设计上可以知道10秒就执行完所有协程,但是主线程不知道,因此底层可能仍然出现资源争夺,因此加入互斥锁即可解决问题.
	这里就算出现 "Found X data race(s)" 这种资源竞争提示,但是也不影响什么,可能就是提示一个信息而已.
	*/ 
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)	// 这里输出变量这个结果
	}
	lock.Unlock()
}

func main() {
	
	first ()

	// 问题: 主线程具体等待协程多少时间?
}