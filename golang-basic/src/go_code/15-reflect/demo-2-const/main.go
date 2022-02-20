/*
补充-常量
	(1)常量介绍
		常量使用 const 修改
		常量在定义的时候,必须初始化
		常量不能修改
		常量只能修饰bool、数值类型(int、float系列)、string类型
		语法:	const identifier [type] = value
		看下面的写法是否正确:
			const name = "tom"				// ok
			const tax float664 = 0.8		// ok
			const a int						// error
			const b = 9/3					// ok
			const v = getVal()				// error, 因为常量是在编译时就需要确定的值
			const b = num/3					// error, 因为num是个变量
	(2)常量使用注意事项
		2.1 比较简洁的写法:
				func main() {
					const (
						a = 1
						b = 2
					)
					fmt.Println(a, b)
				}
		2.2 还有一种专业的写法:
				func main() {
					// 表示给 a 赋值为 0, b在a的基础上+1, c在b的基础上+1,这种写法就比较专业
					const(
						a = iota
						b
						c
					)
					fmt.Println(a, b, c)
				}
		2.3 Golang中没有常量名必须字母大写的规范.比如 TAX_RATE
		2.4 仍然通过首字母的大小写来控制常量的访问范围
*/

package main

import (
	"fmt"
)

func main() {
	var num int
	num = 9 
	
	const tax int = 0 	//常量声明的时候，必须赋值。
	//tax = 10 			//常量不能被修改
	fmt.Println(num, tax)

	//常量只能修饰bool、数值类型(int, float系列)、string 类型
	const aa = 9/3
	fmt.Println(aa)

	//const bb = num/3	// 不能用这种方式定义常量,因为num是个变量
	//fmt.Println(bb)

	const (
		a = iota
		b 
		c
		d
	)
	fmt.Println(a, b, c, d)

	const (
		e = iota
		f = iota
		g = iota
		h = iota
	)
	fmt.Println(e, f, g, h)
}