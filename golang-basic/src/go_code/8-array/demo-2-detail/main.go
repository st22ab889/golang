/*
数组使用注意事项和细节:
	1. 数组是多个相同类型数据的组合,一个数组一旦声明(定义了),其长度是固定的,不能动态变化
	2. var arr[]int 这时arr就是一个slice切片
	3. 数组中的元素可以是任何数据类型,包括值类型和引用类型,但是不能混用
	4. 数组创建后,如果没有赋值，有默认值
			数值类型数组,默认值为 0
			字符串数组,默认值为""
			bool数组,默认值为false
	5. 使用数组的步骤: 1.声明数组并开辟空间; 2.给数组各个元素赋值; 3.使用数组
	6. 数组的小标是从0开始的
	7. 数组下标必须在指定范围内使用,否则会报数组越界错误
	8. Golang的数组属值类型,在默认情况下是值传递,因此会进行值拷贝.数组间不会相互影响
	9. 如果要在其它函数中去修改原数组,可以使用引用传递(指针方式)
	10. 长度是数组类型的一部分,在传递函数参数时需要考虑数组的长度

数组排序
	1.内部排序:指将需要处理的所有数据都加载到内部存储器中进行排序.包括(交换式排序法、选择式排序法和插入式排序法)
		1.1 交换式排序属于内部排序法,是运用数据值比较后,依判断规则对数据位置进行交换,已达到数据位置进行交换,已达到排序的目的,交换式排序又可分为两种:
			1.1.1 冒泡排序法(Bubble sort):
				对待排序序列从后向前(从下标较大的元素开始),依次比较相邻元素的排序码，若发生逆序则交换,使排序码较小的元素逐渐从后部移向前部(从下标较大的单元移向下标较小的单元)
			1.1.2 快速排序法(Quick sort):
	2.外部排序:数据量过大,无法全部加载到内存中,需要借助外部存储进行排序.包括(合并排序法 和 直接合并排序法)

查找,在Golang中,常用的查找有两种
	顺序查找
	二分查找(前提是有序数组)
*/

package main
import(
	"fmt"
	"time"
	"math/rand"
	_ "runtime"
	_ "strconv"
)

// 值拷贝传递
func test01(arr [3]int){ // 这里不能写成"func test01(arr []int)",因为长度是数组类型的一部分,在传递函数参数时需要考虑数组的长度
	arr[0] = 99
}

// 指针传递
func test02(arr *[3]int){
	fmt.Printf("arr指针的地址=%p \n", &arr)
	(*arr)[0] = 99
}

func main() {

	var arr01 [3]float32
	var arr02 [3]string
	var arr03 [3]bool
	fmt.Printf("arr01 = %v, arr02 = %v, arr03 = %v \n", arr01, arr02, arr03)	// arr01 = [0 0 0], arr02 = [  ], arr03 = [false false false]

	arr := [3]int{88, 88, 66}
	test01(arr)
	fmt.Println(arr)	// [88 88 66], 结果没有变, 因为数组是值拷贝传递
	test02(&arr)
	fmt.Println(arr)	// [99 88 66]

	
	// 长度是数组类型的一部分,在传递函数参数时需要考虑数组的长度
	// arr2 := [2]int{88, 99}
	// test01(arr2)		// 错误,因为参数中数组的长度是3,这里的长度是2

	// arr3 := [2]int{66, 88, 99, 99}
	// test01(arr3)		// 错误,因为参数中数组的长度是3,这里的长度是4

	example1()
	example2()
	example3()
	example4() 
}

// 创建一个byte类型的26个元素的数组,分别放置"A-Z",使用for循环并打印出来
func example1() {
	var myChars [26]byte
	for i := 0; i < 26; i++ {
		myChars[i] = 'A' + byte(i) // 注意:这里需要将i转换byte
	}	
	for i := 0; i < 26; i++ {
		fmt.Printf("%c", myChars[i])
	}

	fmt.Printf("\n")
}

// 求一个数组最大值
func example2() {
	var intArr [5]int=[...]int {1, -1, 9, 90, 11}
	maxVal := intArr[0]
	maxValIndex := 0

	for i := 1; i < len(intArr); i++ {
		if maxVal < intArr[i] {
			maxVal = intArr[i]
			maxValIndex = i
		}
	}
	fmt.Printf("maxVal=%v, maxValIndex=%v \n", maxVal, maxValIndex)
}

// 求一个数组和的平平均值. 使用 for-range
func example3() {
	var intArr [5]int=[...]int {1, -1, 9, 90, 13}
	sum := 0
	for _, val := range intArr {
		sum += val
	}

	fmt.Printf("sum=%v, 平均值=%v \n", sum, float64(sum) / float64(len(intArr)))
}

// 数组复杂使用-数组反转
// 要求:随机生成5个数,并将其反转打印
func example4() {

	var intArr [5]int
	len := len(intArr)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		intArr[i] = rand.Intn(100) // 生成的随机数范围为 0 <= n < 100
	}
	fmt.Println("交换前:", intArr)

	// 反转打印
	temp := 0
	for i := 0; i < len / 2; i++ {
		temp = intArr[len - 1 -i]
		intArr[len - 1 -i] = intArr[i]
		intArr[i] = temp
	}
	fmt.Println("交换后:", intArr)
}