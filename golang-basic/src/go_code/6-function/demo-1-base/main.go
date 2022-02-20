/*
不使用函数就会出现 代码冗余 和 不利于代码维护

为完成某一功能的程序指令(语句)的集合,称为函数

在Golang中,函数分为:自定义函数、系统函数(查看Golang编程手册)

函数基本语法:
	func 函数名 (形参列表) (返回值类型列表){
		执行语句
		return 返回值列表
	}
	// 形参列表:表示函数的输入
	// 函数中的语句:为了实现某一功能代码块
	// 函数可以有返回值,也可以没有

Golang函数支持返回多个值,这一点是其它编程语言没有的.
	func 函数名 (形参列表) (返回值类型列表){
		语句...
		return 返回值列表
	}
	// 如果返回多个值时,在接收时,希望忽略某个返回值,则使用 _ 符号表示占位忽略
	// 如果返回值只有一个,  (返回值类型列表) 可以不写()

函数调用机制底层剖析:
	在调用一个函数时,会给该函数分配一个新的空间,编译器会通过自身的处理让这个新的空间和其它的栈空间区分开来
	在每个函数对应的栈中,数据空间时独立的, 不会混肴
	当一个函数调用完毕(执行完毕)后,程序会销毁这个函数对应的栈空间

递归调用:一个函数在函数体内又调用了本身,函数递归需要遵守的重要原则如下:
	执行一个函数时,就创建一个新的受保护的独立空间(新函数栈)
	函数的局部变量时独立的,不会相互影响
	递归必须向退出递归的条件逼近,否则就算无限递归
	当一个函数执行完毕,或者遇到 return 就会返回,遵守谁调用就将结果返回给谁,同时当函数执行完毕或者返回时,该函数本身也会被系统销毁
*/

package main
import(
	"fmt"
)

// 一个test函数
func test(n1 int){
	n1 += 1
	fmt.Println("test()函数执行结果", n1)
}

// 计算两个数的和
func getSum(n1 int, n2 int) int {
	sum := n1 + n2
	return sum
	// return n1 + n2 // 也可以这样写
}

// 计算两个数的和以及差
func getSumAndSub(n1 int, n2 int) (int, int) {
	//sum := n1 + n2
	//sub := n1 - n2
	//return sum ,sub
	return n1 + n2, n1 - n2	// 上面三句话也可以简写为一句话
}

// 递归1
func recursion1(n int){
	if n > 2 {
		n--
		recursion1(n)
	}
	fmt.Println("recursion1 n = ", n)
} 

// 递归2
func recursion2(n int){
	if n > 2 {
		n--
		recursion2(n)
	}else{
		fmt.Println("recursion2 n = ", n)
	}
	
} 

func cal (n1 float64, n2 float64, operator byte) float64 {
	var res float64
	switch operator {
		case '+':
			res = n1 + n2
		case '-':
			res = n1 - n2
		case '*':
			res = n1 * n2
		case '/':
			res = n1 / n2
		default:
			fmt.Println("操作符错误")
	}
	return res
}

func main() {
	
	// 示例1
	n1 := 10
	test(n1)

	// 示例2
	sum := getSum(9, 90)
	fmt.Println("sum = ", sum)

	// 示例3
	r1, r2 := getSumAndSub(99, 0)
	fmt.Println("sum = ", r1, ", sub = ", r2)

	// 示例4
	//如果希望忽略某本返回值, 则使用 _ 符号表示占位忽略
	_, r3 := getSumAndSub(99, 0)
	fmt.Println("sub = ", r3)
	_, _ = getSumAndSub(99, 0)	// 相当于 getSumAndSub(99, 0)
	
	// 示例5-测试递归
	n := 4
	recursion1(n)
	recursion2(n)

	// 示例6
	fmt.Println("90.0 + 9.0 = ", cal(90.0, 9.0, '+'))
	fmt.Println("108.0 - 9.0 = ", cal(108.0, 9.0, '-'))
	fmt.Println("90.0 * 9.0 = ", cal(90.0, 9.0, '*'))
	fmt.Println("90.0 / 9.0 = ", cal(90.0, 9.0, '/'))
	fmt.Println("90.0 & 9.0 = ", cal(90.0, 9.0, '&'))

}

