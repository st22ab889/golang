/*
闭包:一个函数和与其相关的引用环境组合的一个整体
*/

package main
import(
	"fmt"
	"strings"
)

/*
addUpper()是一个累加器函数,返回"func (int) int"这样一个函数类型.
实际上 addUpper() 返回的是一个匿名函数,但是这个匿名函数引用到函数外的n,因此这个匿名函数就和n形成一个整体,构成闭包.
可以这样理解: 闭包是类,函数是操作,n是字段.函数和它使用到的n构成闭包,方反复调用函数时,因为n时初始化一次,因此每调用一次就进行累加
闭包的关键:分析除返回的函数它使用(引用)到哪些变量,因为函数和它引用到的变量共同构成闭包.
*/
func addUpper() func (int) int { 
	var n int = 10
	var str = "hello"
	return func (x int) int {
		n = n + x
		str += string(36) 			//  36 对应的ASCII码是 '$'   
		fmt.Println("str=", str) 	//  第一次打印 str= hello$, 第二次打印 str= hello$$
		return n
	}
}

func main() {
	
	f := addUpper()

	//fmt.Println(f(2))	// 12

	// 为什么前面调用了"fmt.Println(f(1))", "fmt.Println(f(2))"结果就变成了13
	fmt.Println(f(1))	// 11 , 因为形成闭包后,n就变成11了
	fmt.Println(f(2))	// 13 ,  11 + 2 = 13 


	//测试makeSuffix,使用闭包的好处参数只需要传入一次.比如: ".jpg"传入一次后, 闭包就可以保留这个值,然后接下来就可以反复使用
	f2 := makeSuffix(".jpg")
	fmt.Println("文件名处理后=", f2("winter"))		// 文件名处理后= winter.jpg
	fmt.Println("文件名处理后=", f2("bird.jpg"))	// 文件名处理后= bird.jpg
}

/*
闭包的最佳实践:
	编写一个函数 makeSuffix(suffix string) 可以接收一个文件后缀名(比如.jpg),并返回一个闭包
	调用闭包,可以传入一个文件名,如果该文件名没有指定的后缀(比如.jpg),则返回"文件名.jpg",如果已经有".jpg"后缀,则返回原文件名
*/
func makeSuffix(suffix string) func (string) string {
	return func (name string) string{
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}