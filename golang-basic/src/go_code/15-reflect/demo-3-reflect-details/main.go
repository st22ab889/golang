/*
反射注意事项和细节说明:
	(1) reflect.Value.Kind, 获取变量的类别,返回的是一个常量.
	(2) Type是类型,Kind是类别,Type和Kind可能是相同的,也可能是不同的.
		比如: var num int = 10, num的Type是int, Kind也是int
		比如: var stu Student, stu的Type是 包名.Student, Kind是struct
	(3) 通过反射可以让变量在 interface{} 和 Reflect.Value 之间相互转换
			变量 <----------> interface{} <----------> reflect.Value
	(4) 使用反射的方式来获取变量的值(并返回对应的类型),要求数据类型匹配,比如x是int,那么就应该使用 reflect.Value(x).Int(),而不能使用其它的,否则报panic
			func (v Value) Int() int64
			返回v持有的有符号整数(表示为int64),如果v的Kind不是Int、Int8、Int16、Int32、Int64会panic
	(5) 通过反射来修改变量
		注意:当使用SetXxx方法来设置,需要通过对应的指针类型来完成, 这样才能改变传入的变量的值,同时需要使用到 reflect.Value.Elem() 方法
			func (v Value) SetBool(x bool)
			func (v Value) SetInt(x int64)
			func (v Value) SetUint(x uint64)
			func (v Value) SetFloat(x float64)
			func (v Value) SetComplex(x complex128)
			func (v Value) SetBytes(x []byte)
			func (v Value) SetString(x string)
			func (v Value) SetPointer(x unsafe.Pointer)
			func (v Value) SetCap(n int)
			func (v Value) SetLen(n int)
			func (v Value) SetMapIndex(key, val Value)
			func (v Value) Set(x Value)

			func (Value) Elem
				func (v Value) Elem() Value
				Elem返回v持有的接口保管的值的Value封装,或者v持有的指针指向的值的Value封装.如果v的Kind不是Interface或Ptr会panic;如果v持有的值为nil,会返回Value零值
	(6) reflect.Value.Elem()应该如何理解?
			func main() {
				var num int = 100
				fn := reflect.ValueOf(&num)
				fn.Elem().SetInt(200)
				fmt.Printf("%v \n", num)

				// fn.Elem() 用于获取指针指向变量,类似:
				var num2 = 10
				var b *int = &num2
				*b = 3
			}
*/

package main

import (
	"fmt"
	"reflect"
)

func test (b interface{}) {
	//1.获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal kind=%v\n", rVal.Kind()) // 查看 rVal的Kind、
	fmt.Printf("rVal type=%T\n", rVal)

	//2.Elem返回v持有的接口保管的值的Value封装,或者v持有的指针指向的值的Value封装
	rVal.Elem().SetInt(20)
	fmt.Printf("rVal =%v\n", rVal)

}

//通过反射修改int类型变量的值
func first () {
	var num int = 10
	test(&num)
	fmt.Println("num=", num) 

	//可以这样理解rVal.Elem()
	num2 := 9
	var ptr *int = &num2
	num3 := *ptr  //=== 类似 rVal.Elem()
	*ptr = 22
	fmt.Println("num2=", num2)		// num2= 22
	fmt.Println("num3=", num3)		// num3= 9
}


func second () {
	var str string = "tom"

	// 下面这种写法是错误的
	//fs := reflect.ValueOf(str) 	// 这里不是传值,传值的话fs就是一个string类型(fs -> string), 而是传入一个地址
	//fs.SetString("jack") 			// 报错: panic: reflect: reflect.Value.SetString using unaddressable value

	fs := reflect.ValueOf(&str) 	// 传地址时这里的fs就是指针类型 	
	fs.Elem().SetString("jack") 	// 通过指针修改值
	fmt.Printf("%v\n", str) 		// jack
}


func main() {

	first ()	// 通过反射修改变量 以及 理解 func (Value) Elem

	second ()	// 通过反射修改变量-查找代码错误
}