/*
数组是一种数据类型,数组可以存放多个同一类型数据,在Golang中,数组是值类型

数组定义:
	var 数组名 [数组大小]数据类型
	var a [5]int
	a[0] = 88		// 赋初值
	a[1] = 99
	...

数组内存图

数组的使用:
	数组名[下标]

四种初始化数组的方式:
	var array01 [3]int = [3]int {1, 2, 3}
	var array02 = [3]int {1, 2, 3}
	var array03 = [...]int {1, 2, 3}
	var names = [3]string{1:"tom", 0:"jack", 2:"marry"}		// 可以指定元素值对应的下标 ...

数组的遍历:
	方式1:常规遍历
	方式1:for-range遍历:这时Golang语言一种独有的结构,可以用来遍历访问数组的元素
		for index, value := range array01{
			....
		}
		// 说明
		// 1.第一个返回值index是数组的下标
		// 2.第二个value是在该下标位置的值
		// 3.它们都是仅在for循环内部可见的局部变量
		// 4.遍历数组元素的时候,如果不想使用下标index,可以直接把下标index标为下划线 _
		// 5.index和value的名称不是固定的,可以自行指定,一般命名为 index 和 value
*/

package main
import(
	"fmt"
	"runtime"
	"strconv"
)


func testAarry_1() {
	// 定义一个数组,以及给每个元素赋值
	var hens [6]float64
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 2.0
	// 遍历数组求和
	total := 0.0
	for i := 0; i < len(hens); i++{
		total += hens[i]
	}
	fmt.Printf("total=%.2f \n", total)		// total=10.00
}


func testAarry_2() {
	
	var intArr [3]int
	
	// 当定义完数组后,数组的各个元素默认值都为0
	fmt.Println(intArr)	// [0 0 0]

	intArr[0] = 22
	intArr[1] = 11
	intArr[2] = 66

	fmt.Printf("intArr的地址=%p, intArr[0]的地址=%p, intArr[1]的地址=%p, intArr[2]的地址=%p \n", &intArr, &intArr[0], &intArr[1], &intArr[2])
	// intArr的地址=0xc0000ae078, intArr[0]的地址=0xc0000ae078, intArr[1]的地址=0xc0000ae080, intArr[2]的地址=0xc0000ae088
	// 0xc0000ae088 - 0xc0000ae080 = 8, 0xc0000ae080 - c0000ae078 = 8; 由此可见,int占8个字节

	// 由上面可知
	// 1.数组的地址可以通过数组名来获取
	// 2.数组的第一个元素的地址, 就是数组的首地址
	// 3.数组的各个元素的地址间隔是依据数组的类型决定, 比如int64间隔8, int32间隔4

	// golang中int占用多少个字节
	// 如果是32位CPU就是4个字节，如果是64位就是8个字节，由CPU的位数决定，然后按照公式1字节 = 8位计算
	fmt.Println("CPU型号 = ", runtime.GOARCH)	// CPU型号 =  amd64
	fmt.Println("int位数 = ", strconv.IntSize)	// int位数 =  64
}
 

func testAarry_3() {

	var array01 [3]int = [3]int {1, 2, 3}
	var array02 = [3]int {1, 2, 3}
	var array03 = [...]int {1, 2, 3}						// 这里的[...]是规定的写法
	var array04 = [...]int{1:88, 0:66, 2:99}
	// 类型推导
	names := [3]string{1:"tom", 0:"jack", 2:"marry"}		// 可以指定元素值对应的下标 ...
	

	fmt.Println(array01)	// [1 2 3]
	fmt.Println(array02)	// [1 2 3]
	fmt.Println(array03)	// [1 2 3]
	fmt.Println(array04)	// [66 88 99]
	fmt.Println(names)		// [jack tom marry]
}


// 数组的遍历
func testAarry_4(){

	names := [3]string{1:"tom", 0:"jack", 2:"marry"}
	for i, v := range names {
		fmt.Printf("index=%v, value=%v \n", i ,v)
		fmt.Printf("names[%d]=%v \n", i , names[i])
	}

	for _, v := range names {
		fmt.Printf("元素的值=%v \n" ,v)
	}
}


func main() {
	
	testAarry_1()

	testAarry_2()

	testAarry_3()

	testAarry_4()
}