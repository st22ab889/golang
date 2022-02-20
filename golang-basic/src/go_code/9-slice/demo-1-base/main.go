/*
为什么需要切片:需要一个数组用于保存学生的成绩,但学生的个数不是确定的,就需要切片来解决

切片基本介绍:
	1.切片是数组的一个引用,因此切片是引用类型,在进行传递时遵守引用传递的机制
	2.切片的使用和数组类似,遍历切片、访问切片的元素和求切片长度len(slice)都一样
	3.切片的长度可以变化,因此切片是一个可以动态变化的数组
	4.切片定义的基本语法
		var 变量名 []类型

切片的内存布局:
	切片底层的数据结构可以理解成是一个结构体struct,这个结构体保存了以下信息：
		数组中起始元素的地址
		起始到结束的元素个数,也就说长度
		容量
	这个结构体类似如下形式:
		type slice struct {
			ptr *[2]int
			len int
			cap int
		}
*/

package main

import (
	"fmt"
)

// 切片的基本使用
// 切片的使用方式1: 定义一个切片,然后让切片去引用一个已经创建好的数组
func first () {

	var intArr [5]int=[...]int {1, 22, 66, 88, 99}

	// 声明/定义一个切片,slice就是切片名
	// intArr[1:3] 表示 slice 引用到 intArr 这个数组, 引用 intArr 数组的起始下标为1(包含1), 最后的下标为3(但是不包含3)
	slice := intArr[1:3]

	fmt.Println("slice的元素 = ", slice)
	fmt.Println("slice的元素个数 = ", len(slice))
	fmt.Println("slice的容量 = ", cap(slice))	// 切片的容量可以动态变化
}

// 切片的使用方式2: 通过make来创建切片  
func second () {
	/*
	基本语法: 
		var 切片名 []type = make([], len, [cap])
	参数说明: 
		type是数据类型
		len表示大小
		cap指定切片容量(可选参数,如果分配了cap,则要求 cap >= len)
	*/

	var sliceDemo []float64 = make([]float64, 5, 50)
	sliceDemo[1] = 10;
	sliceDemo[3] = 20;
	fmt.Println(sliceDemo)
	fmt.Println("sliceDemo的size = ", len(sliceDemo))
	fmt.Println("sliceDemo的cap = ", cap(sliceDemo))
	fmt.Println("slice[3] = ", sliceDemo[3])
	/*
	总结:
		1.通过make方式创建切片可以指定切片的大小和容量
		2.如果没有有个切片的各个元素赋值,就会使用默认值
		3.通过make方式创建的切片对应的数组是由make底层维护,对外不可见(既只能通过slice去访问各个元素)
	*/
}

/*
切片的使用方式1和方式2的区别(面试):
1.方式1是直接引用数组,这个数组事先存在,程序员可见
2.方式2是通过make来创建切片,make也会创建一个数组,是由切片在底层进行维护,程序员不可见
*/

// 切片的使用方式3: 定义一个切片,直接就指定具体数组,使用原理类似make方式 
func three () {
	var strSlice []string = []string{"tom", "jack", "mary"}
	fmt.Println("strSlice = ", strSlice)			// strSlice =  [tom jack mary]
	fmt.Println("strSlice size = ", len(strSlice))	// strSlice size =  3
	fmt.Println("strSlice cap = ", cap(strSlice))	// strSlice cap =  3
} 

// 切片的遍历:
func four() {
	// 方式1: for循环常规方式遍历
	var arr	[5]int = [...]int{10, 20, 30, 40, 50}
	slice := arr[1:4]
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v] = %v \n", i, slice[i])
	}

	// 方式2: for-range结构遍历切片
	for i, v := range slice {
		fmt.Printf("i=%v, v= %v \n" ,i , v)
	}
}	


func main() {

	first ()

	second ()

	three ()
	
	four()
}