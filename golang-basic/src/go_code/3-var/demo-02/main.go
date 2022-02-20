/*
声明变量
	var a int  		// 表示声明一个变量,变量名是 a
	var b float32	// 表示声明一个单精度类型的小数,变量名是 b

初始化变量
	var a int =45	// 在声明的时候就给值,这就是初始化变量a. 注意:如果声明时就直接赋值,可省略数据类型
	var b = 45		// 声明时就直接赋值,可省略数据类型

给变量赋值
	var num int
	num = 99		// 先声明再赋值,就是给变量赋值	

程序中 + 号的的使用:
	当左右两边都是数值型,则做加法运算
	当左右两边都是字符串,则做字符串拼接
*/

package main
import "fmt"
func main()  {
	
	var i1, i2 = 3, 4
	var i3 = i1 + i2
	fmt.Println("i3 = ", i3)

	var s1, s2 = "hello ", "Golang"
	s3 := s1 + s2
	fmt.Println("s3 = ", s3)

}
