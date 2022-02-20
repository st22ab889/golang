/*
指针
	基本数据类型的变量存的是值,也叫值类型
	使用 & 获取变量的地址
	指针类型的变量存的是一个地址,这个地址指向的空间存的是值
	使用 * 获取指针类型所指向的值

指针的使用陷进
	
*/

package main
import(
	"fmt"
	_ "unsafe"
	_ "strconv"
)

func main()  {

	var i int = 10
	fmt.Println("i的地址=", &i)		// i的地址= 0xc000128058

	// ptr是一个指针变量, ptr的类型是 *int, ptr 的值是 &i
	var prt *int =  &i	

	fmt.Printf("prt = %v \n", prt)			 // prt = 0xc0000aa058	
	fmt.Printf("prt 的地址 = %v \n", &prt)	 //	 prt 的地址 = 0xc0000ce020	
	fmt.Printf("prt 指向的值 = %v \n", *prt) //	 prt 指向的值 = 10

	
	var num int  = 9
	var ptr *int
	ptr = &num
	*ptr = 10 	// 这里修改会导致num的值变化
	fmt.Println("num = ", num)	//	num =  10


	/*
	// 错误案例1
	var a int = 300
	var ptr *int = a 	//这里直接赋值时错误的,因为这里接收的是个指针类型的值

	// 错误案例2
	var a int = 300
	var prt *float32 = &a	//这里错误的的原因时类型不匹配 

	// 指针的使用细节
	// 1.值类型都有对应的指针类型, 形式为 *数据类型 ,比如int对应的指针就是 *int, float32对应的指针类型就是 *float32, 以此类推
	// 2.值类型包括: 基本数据类型(int系列、float系列、bool、string、数组、结构体struct)

	*/

	valueType()
}


func valueType() {

	/*
	值类型和引用类型:
		值类型: 基本数据类型(int系列、float系列、bool、string、数组、结构体struct)
			值类型特点:变量直接存储值,内存通常在栈中分配
		引用类型: 指针、slice切片、map、管道Channel、interface等都是引用类型
			引用类型特点:变量存储的是一个地址,这个地址对应的空间才是真正存储数据(值),内存通常在堆上分配,当没有任何变量引用这个地址时,该地址对应的数据空间就成为一个垃圾,由GC来回收
	*/

}
