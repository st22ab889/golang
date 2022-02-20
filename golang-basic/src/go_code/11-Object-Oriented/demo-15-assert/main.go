/*
类型断言:由于接口是一般类型,不知道具体类型,如果要转成具体类型,就需要使用到类型断言
		type Point struct {
			x int
			y int
		}
		func main() {
			var a interface{}
			var point Point = Point{1,2}
			a = point

			var b Point
			// b = a		// 将接口变量a赋给Point变量报错
			b = a.(Point)	// 不报错. "b = a.(Point)"就是类型断言,表示判断a是否指向Point类型的变量,如果是就转成Point类型并赋给b变量,否则报错
			fmt.Println("b = ", b )
		}
*/

package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}


func first () {
	var a interface{}
	var point Point = Point{1,2}
	a = point

	var b Point
	// b = a		// 不可以
	b = a.(Point)	// 可以,这里就是类型断言
	fmt.Println("b = ", b )
}


func second () {
	var t float32 = 99.99
	var x interface{}
	x = t				// 空接口,可以接收任意类型
	
	/*
	// 情况一: 类型断言,不带检查
	y := x.(float64) // 转成float64报错,在进行类型断言时,如果类型不匹配,就会报panic(panic发生后程序报错不再往下执行),因此进行类型断言时,要确保原来的空接口指向的就是断言的类型
	fmt.Printf("Convert Success :: y的类型是 %T, y = %v \n", y, y) 
	*/

	// 情况二: 类型断言,带检查
	y, ok := x.(float64) 	// 转成float,在进行断言时带上检测机制,无论成功还是失败程序都会往下执行,就算失败也不会报panic
	if ok == true {
		fmt.Printf("Convert Success :: y的类型是 %T, y = %v \n", y, y)
	} else {
		fmt.Println("Convert fail")
	}
}


func main() {

	first ()	// 类型断言

	second ()	// 类型断言带检查

	three ()	// 类型断言的最佳实践: 写一函数循环判断传入参数的类型

}

func three () {
	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int32 = 30
	var name string = "tom"
	address := "北京"
	n4 := 300

	stu1 := Student{}
	stu2 := &Student{}

	TypeJudge(n1, n2, n3, name, address, n4, stu1, stu2)
	
}



type Student struct {

}

//判断输入的参数是什么类型
func TypeJudge(items... interface{}) {
	for index, x := range items {
		switch x.(type) {	// 这里type是一个关键字, 固定写法
			case bool :
				fmt.Printf("第%v个参数是 bool 类型，值是%v\n", index, x)
			case float32 :
				fmt.Printf("第%v个参数是 float32 类型，值是%v\n", index, x)
			case float64 :
				fmt.Printf("第%v个参数是 float64 类型，值是%v\n", index, x)
			case int, int32, int64 :
				fmt.Printf("第%v个参数是 整数 类型，值是%v\n", index, x)
			case string :
				fmt.Printf("第%v个参数是 string 类型，值是%v\n", index, x)
			case Student :
				fmt.Printf("第%v个参数是 Student 类型，值是%v\n", index, x)
			case *Student :
				fmt.Printf("第%v个参数是 *Student 类型，值是%v\n", index, x)
			default :
				fmt.Printf("第%v个参数是  类型 不确定，值是%v\n", index, x)
		}
	}
}
