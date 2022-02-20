/*
数据类型
	基本数据类型
		数值型(8、16、32、64表示多少位, int8表示一个1个字节,相当于一个byte)
			整数类型: int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,byte(uint8的别名),rune(int32的别名,和int32有点区别的是,rune存放的是Unicode码值)
					 complex64, complex128	// 复数
					 uintptr //整数类型的指针
			浮点类型: float3(单精度),float64(双精度,相当于java中的double)
		字符型: 没有专门的字符型(比如java中的char),使用byte来保存单个字母字符,byte不能存一个汉字,在go中都采用的时utf-8编码,一个byte存不下一个汉字，一个汉字占3个字节,至少要用int32
		布尔型: bool
		字符型: String [官方将String归属到基本数据类型]
	派生/复杂数据类型: 指针(Pointer),数组,结构体(struct),管道(Channel),函数(也是一种类型),切片(slice),接口(interface),map
*/

package main

// 导包方式1
// import "fmt"
// import "unsafe"

// 导包方式2
import(
	"fmt"
	"unsafe"
)

func main()  {
	
	// 整数类型
	// int8		有符号	1字节	-128 ~ 127
	// int16	有符号	2字节	-2的15次方 ~ 2的15次方 - 1 
	// int32	有符号	4字节	-2的31次方 ~ 2的31次方 - 1
	// int64	有符号	8字节	-2的63次方 ~ 2的63次方 - 1
	// uint8	无符号	1字节	0 ~ 255
	// uint16	无符号	2字节	0 ~ 2的16次方 - 1
	// uint32	无符号	4字节	0 ~ 2的32次方 - 1
	// uint64	无符号	8字节	0 ~ 2的64次方 - 1

	// int		有符号	[32位系统4个字节|64位系统8个字节]			[-2的31次方 ~ 2的31次方 - 1 | -2的63次方 ~ 2的63次方 - 1]
	// uint		无符号	[32位系统4个字节|64位系统8个字节]			[0 ~ 2的32次方 - 1 | 0 ~ 2的64次方 - 1]
	// rune		有符号	与int32一样(等价int32,表示一个Unicode码)	-2的31次方 ~ 2的31次方 - 1
	// byte		无符号	与uint8等价(当要存储字符时选用byte)			0 ~ 255

	var a int = 258369
	var b uint = 255255
	// 如果超出范围会报 overflows byte 错误. overflows表示溢出, 一个类型装一个超出它范围的值就会溢出
	var c byte = 255
	fmt.Println("a = ",a,"b = ",b,"c= ",c);



	// 整数类型使用细节:
	// 1. Golang各整数类型分有符号和无符号,int和uint的大小和系统有关
	// 2. Golang的整型默认声明为int型
	// 3. 如何在程序中查看某个变量的 占用字节大小 和 数据类型 (使用较多)
	// 4. Golang程序中整型变量在使用时,遵守报小不保大的原则.既在保证程序正确运行下,尽量使用占用空间小的数据类型.例如年龄
	// 5. bit是计算机中最小存储单位. byte是计算机中基本存储单元

	var d = 100		// 这里d默认是int类型
	fmt.Printf("d的类型是 %T", d)	// fmt.Printf() 可以用于做格式化输出

	var e int64 = 10
	fmt.Printf("e的类型是 %T, e占用的字节数是 %d", e, unsafe.Sizeof(e))	// unsafe.sizeof这个函数,可以返回变量占用的字节数

	var age byte = 20
	fmt.Println("年龄 = ", age)

	
	floatExample()
}

func floatExample()  {
	
	// 浮点类型可以表示一个小数
	var price float32 = 100.01
	fmt.Println("价格 = ", price)

	// 浮点型的分类
	// 单精度float32	4字节		-3.403E38 ~ 3.403E38
	// 双精度float64	8字节		-1.798E308 ~ 1.798E308

	// 浮点型说明 
	// 1.浮点数在机器中的存放形式：浮点数=符号位+指数位+尾数位
	//		3.56 ==> 11110000111.111111111111111111000
	// 2.尾数部分可能丢失,造成精度损失. -123.0000901
	// 3.浮点型的存储分为三部分：符号位+指数位+尾数位,在存储过程中,精度会有丢失
	// 4.看到在相同的占用字节大小的情况下,浮点数比整数能够存储的数据要大是以损失精度来的

	var num1 float32 = -0.00099
	var num2 float64 = -7809656.09
	fmt.Println("num1 = ", num1, "num2 = ", num2)

	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	// float64的精度比float32的要准确,如果要保存一个精度高的数,则应该选用float64	
	fmt.Println("num3 = ", num3, ", num4 = ", num4)		// num3 =  -123.00009 , num4 =  -123.0000901

	// 浮点型使用细节
	// 1.Golang浮点类型有固定的范围和字段长度,不受具体OS的影响
	// 2.Golang的浮点型默认声明为float64类型
	// 3.浮点型常量有两种表示方式
	//		十进制数形式(必须有小数点): 如 5.12 ，.512
	//		科学计数法形式: 如 5.1234e2 = 5.1234 * 10的2次方, 5.12E-2 = 5.12/10的2次方
	// 4.通常情况下应该使用float64,因为它比float32更精确

	var num5 = 1.1	// 默认是 float64
	fmt.Printf("num5的数据类型是 %T", num5)

	num6 := 5.12
	num7 := .123  // .123 ==> 0.123
	fmt.Println("num6 = ", num6, ", num7 = ", num7)

	num8 := 5.1234e2	// 5.1234e2 = 5.1234 * 10的2次方
	num9 := 5.1234E2	// 5.1234E2 = 5.1234 * 10的2次方
	num10 := 5.1234E-2	// 5.1234E-2 = 5.1234/10的2次方 = 0.051234
	fmt.Println("num8 = ", num8, ", num9 = ", num9, ", num10 = ", num10)
}
