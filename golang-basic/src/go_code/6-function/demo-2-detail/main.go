/*
函数注意事项和细节讨论:
	1.函数的形参列表和返回值列表都可以时多个
	2.形参列表和返回值列表的数据类型可以时值类型和引用类型
	3.函数的命名遵守标识符命名规范,首字母不能时数字,首字母大写该函数可以被本包文件和其它包文件使用,类似public.
	  首字母小写,只能被本包文件使用,其他包文件不能使用,类似 private
	4.函数中的变量时局部的,函数外不能生效
	5.基本数据类型 和 数组 默认都是值传递的,既进行值拷贝.在函数内修改不会影响到原来的值
	6.如果希望函数内的变量能修改函数外的变量(指的是默认以值传递的方式的数据类型),可以传入变量的地址&,函数内以指针的方式操作变量.从效果上看类似引用
	7.Golang函数不支持函数重载
	8.在Golang中,函数也是一种数据类型,可以赋值给一个变量,则该变量就是一个函数类型的变量了.通过该变量可以对函数调用
	9.函数既然是一种数据类型,因此在Golang中,函数可以作为形参并且调用
	10.为了简化数据类型定义,Golang支持自定义数据类型(理解:相当于一个别名)
		基本语法: type 自定义数据类型名 数据类型
			type myInt int	//这时 myInt 就等价于 int 来使用了
			type mySum func(int, int) int //这时 mySum 就等价一个函数类型 func(int, int) int
	11.Golang支持对函数返回值命名
		func cal(n1 int, n2 int) (sum int, sub int){
			sum = n1 + n2
			sub = n1 - n2
			return
		}
	12.使用 _ 标识符,忽略返回值
	13.Golang支持可变参数
		支持0到多个参数:
			func sum(args... int) sum int {
			}
		支持1到多个参数:
			func sum(n1 int, args... int) sum int {
			}
		说明:
			args是slice切片,通过args[index]可以访问到各个值
			args只是个名称, 这个名称可以是任意的,只要满足命名规范即可
			如果一个函数的形参列表中有可变参数,则可变参数需要放在形参列表最后
*/

package main
import(
	"fmt"
)

/*
func updateValue {
	fmt.Printf("Golang函数不支持函数重载")
}
*/

func updateValue(n *int) {
	fmt.Printf("B -> n的地址 %v \n", &n)
	*n = *n + 9
	fmt.Println("C -> n = ", *n)
}


func getSum(n1 int, n2 int) int {
	return n1 + n2
}

// 函数既然是一种数据类型,因此在Golang中,函数可以作为形参并且调用
func myFun(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

// Golang自定义函数类型
type myFunType func(int, int) int  // 这时 myFunType 就是  func(int, int) int 类型 
func myFun2(funvar myFunType, num1 int, num2 int) int {
	return funvar(num1, num2)
}

// Golang支持对函数返回值命名
func getSunAndSub(n1 int, n2 int) (sum int, sub int){
	sum = n1 + n2
	sub = n1 - n2
	return
}

// Golang支持可变参数
func sum(n1 int, args... int) int {
	sum := n1
	//遍历args
	for i := 0; i < len(args); i++ {
		// args[0] 表示取出args切片的第i个元素,其它依次类推
		sum += args[i]
	}
	return sum
}


func main() {
	
	n := 90
	fmt.Printf("A -> n的地址 %v \n", &n)
	updateValue(&n)
	fmt.Println("D -> n = ", n)

	//在Golang中,函数也是一种数据类型,可以赋值给一个变量,则该变量就是一个函数类型的变量了.通过该变量可以对函数调用
	a := getSum
	fmt.Printf("a的类型%T, getSum类型是%T \n", a, getSum)
	res := a(90, 9)		// 等价 res := getSum(90, 9)
	fmt.Println("res = ", res)

	// 函数既然是一种数据类型,因此在Golang中,函数可以作为形参并且调用
	res2 := myFun(getSum, 90 , 9)
	fmt.Println("myFun -> res2 = ", res2)

	// Golang自定义数据类型
	// 给int取了别名, 在Golang中 myInt 和 int 虽然都是int类型,但是Golang从语法上认为 myInt 和 int还是两个不同的数据类型
	type myInt int
	var n1 myInt
	n1 = 40
	fmt.Println("n1 = ", n1)
	// var n2 int = n1  // 这里报错,因为Golang从语法上认为 myInt 和 int还是两个不同的数据类型
	var n2 int = int(n1) // 这里仍然需要显示转换
	fmt.Println("n2 = ", n2)

	res3 := myFun2(getSum, 90 , 9)
	fmt.Println("myFun2 -> res3 = ", res3)

	// Golang支持对函数返回值命名
	n3, n4 := getSunAndSub(9, 90)
	fmt.Printf("n3 = %v, n4 = %v \n", n3, n4)

	// Golang支持可变参数
	n5 := sum(10, 0, -5, 88, 6)
	fmt.Println("n5 = ", n5)

}
