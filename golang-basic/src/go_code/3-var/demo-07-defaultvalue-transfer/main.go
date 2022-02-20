
/* 
在go中,数据类型都有一个默认值,默认值又叫 零值

整型  		默认值为	0
浮点型		默认值为	0
字符串		默认值为	""
布尔类型	默认值为	false
*/

/*
基本数据类型的转换:Golang和java/c不同,Go在不同类型的变量之间赋值时需要显示转换。也就是说Golang中数据类型不能自动转换

基本语法: 表达式 T(v) ，表示将值 v 转换为类型 T
	T: 就是数据类型, 比如 int32, int64, float32等等
	v: 就是需要转换的变量

基本数据类型的转换细节说明:
	a. Golang中,数据类型的转换可以是从 范围小 -> 范围大, 也可以 范围大 -> 范围小
	b. 被转换的是 变量存储的数据(既值), 变量本身的数据类型并没有变化
	c. 在转换中,比如将 int64 转成 int8,编译时不会报错,只是转换的结果是按溢出处理,和希望得到的结果不一样
	
*/


package main
import(
	"fmt"
	_ "unsafe"	// 如果没有使用到这个包,但又不想去掉,前面加一个 _ 表示忽略
	"strconv"
)

func main()  {

	var i int32 = 100
	var n1 float32 = float32(i)	// 整型 -> 浮点型
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)		// 低精度 -> 高精度

	fmt.Printf("i = %v, n1 = %v, n2 = %v, n3 = %v \n", i, n1, n2, n3) // %v 意思是值的默认格式表示
	// i = 100, n1 = 100, n2 = 100, n3 = 100

	baseTypeToString()

	stringToBaseType()
}

/*
在程序开发中,经常需要将基本数据类型和string类型之间互相转换.

基本类型转string类型
	方式1: fmt.Sprintf(format string, a ...interface{}) String	// Sprintf根据format参数生成格式化的字符串并返回该字符串
		fmt.Sprintf("%参数",表达式)	//这个方式灵活,函数会返回转换后的字符串
		参数需要和表达式的数据类型相匹配
	方式2: 使用strconv包的函数
		func FormatBool(b bool) String
		func FormatFloat(f float64, fmt byte, prec, bitSize int) String
		func FormatInt(i int64, base int) String
		func FormatUint(i uint64, base int) String

*/
func baseTypeToString(){

	var num1 int = 99
	var num2 float64 = 2.222
	var b bool = true
	var mychar byte = 'h'
	var str string	// 空的str

	// 方式一
	str = fmt.Sprintf("%d", num1)	// %d 表示十进制, %T 值的类型的Go语法表示
	fmt.Printf("str type %T, str = %v \n", str, str)	// str type string, str = 99

	str = fmt.Sprintf("%f", num2)	// %f 表示有小数部分但无指数部分
	fmt.Printf("str type %T, str = %v \n", str, str)	// str type string, str = 2.222000

	str = fmt.Sprintf("%t", b)	//  %q 该值对应的单引号括起来的Go语法字符字面值,必要时会采用安全的转义表示
	fmt.Printf("str type %T, str = %q \n", str, str)	// str type string, str = "true"

	str = fmt.Sprintf("%c", mychar)
	fmt.Printf("str type %T, str = %q \n", str, str)	// str type string, str = "h"
	

	// 方式二
	var num3 int = 99
	var num4 float64 = 2.222
	var b2 bool = true
	
	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T, str = %q \n", str, str)	// str type string, str = "99"

	// 'f'表示格式, 10表示小数位保留10位, 64表示这个小数是float64
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T, str = %q \n", str, str)	// str type string, str = "2.2220000000"

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T, str = %q \n", str, str)	// str type string, str = "true"

	// 方式三: strconv 包中有一个函数 Itoa
	var num5 int64 = 999
	str = strconv.Itoa(int(num5))
	fmt.Printf("str type %T, str = %q \n", str, str)	// str type string, str = "999"
}


/*
string类型转基本数据类型,使用 strconv 包的函数
	func ParseBool(str String) (value bool, err error)
	func ParseFloat(s string, bitSize int) (f float64, err error)
	func ParseInt(s string, base int, bitSize int) (i int64, err error)
		// 参数1 数字的字符串形式
		// 参数2 数字字符串的进制 比如二进制 八进制 十进制 十六进制
		// 参数3 返回结果的bit大小 也就是int8 int16 int32 int64
	func ParseUint(s string, b int, bitSize int) (n uint64, err error)

注意事项:
	在将string类型转成基本数据类型时,要确保string类型能够转成有效的数据.
	比如可以把"123"转成一个整数,但是不能把"hello"转成一个整数,如果这样做Golang直接将其转成0
	其它类型也是一样的道理,比如不能转成float就直接转为0,不能转换为bool就直接转换为false

*/
func stringToBaseType(){

	var str string = "true"
	var b bool

	// strconv.ParseBool(str) 会返回两个值 (value bool, err error), 因为只想获取 value bool,不想获取 err, 所以使用 _ 忽略
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T, str = %v \n", b, b)	// b type bool, str = true

	var str2 string = "1235689"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str2, 10, 64)
	// ParseInt 返回的是 int64, 如果希望得到int32可以如下处理
	n2 = int(n1)
	fmt.Printf("n1 type %T, n1 = %v \n", n1, n1)	// n1 type int64, n1 = 1235689
	fmt.Printf("n2 type %T, n2 = %v \n", n2, n2)	// n2 type int, n2 = 1235689

	var str3 string = "888.999"
	var f1 float64
	var f2 float32
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type %T, f1 = %v \n", f1, f1)	// f1 type float64, f1 = 888.999
	
	// ParseFloat 返回的是 float64, 希望得到float32可以如下处理
	f2 = float32(f1)
	fmt.Printf("f2 type %T, f2 = %v \n", f2, f2)	// f2 type float32, f2 = 888.999

	// 注意
	var str4 string = "hello"
	var n3 int64 = 99
	n3, _ = strconv.ParseInt(str4, 10, 64)
	// 如果没有转成功, n3就会被赋予默认值,也就是0
	fmt.Printf("n3 type %T, n3 = %v \n", n3, n3)	// n3 type int64, n3 = 0

}