/*
匿名函数:如果希望某个函数只调用一次,可以考虑匿名函数,匿名函数也可以实现多次调用
	匿名函数使用方式1:在定义匿名函数时就直接调用
	匿名函数使用方式2:将匿名函数赋给一个变量(函数变量),再通过该变量来调用匿名函数
	全局匿名函数:如果将匿名函数赋给一个全局变量,那么这个匿名函数就成为一个全局匿名函数

*/

package main
import(
	"fmt"
)

var (
	// func1就是全局匿名函数
	func1 = func (n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	
	// 匿名函数使用方式1
	r1 := func (n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("r1 = ", r1)


	// 匿名函数使用方式2, r2的数据类型就是函数类型
	a := func (n1 int, n2 int) int {
		return n1 + n2
	}
	r2 := a(90, 9)
	fmt.Println("r2 = ", r2)
	
	// 全局匿名函数使用
	r3 := func1(11, 9)
	fmt.Println("r3 = ", r3)
}

