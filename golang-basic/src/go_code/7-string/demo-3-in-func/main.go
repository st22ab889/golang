/*
为了编程方便Golang提了内置函数,这些函数可以直接使用. 文档: https://studygolang.com/pkgdoc  builtin章节
	len		用来求长度,比如string、array、slice、map、channel
	new		用来分配内存,主要用来分配值类型,比如int、float32, struct ... 返回的指针
	make	用来分配内存,主要用来分配引用类型,比如chan、map、slice
*/


package main
import(
	"fmt"
	_ "strconv"
	_ "strings"
)

func main() {

	 num1 := 100
	 fmt.Printf("num1的类型%T, num1的值=%v,num1的地址%v \n", num1, num1, &num1)		
	 //上面的代码执行结果: num1的类型int, num1的值=100,num1的地址0xc000014088

	 num2 := new(int)	// new返回了一个 *int 类型的指针类型
	 *num2 = 100
	 fmt.Printf("num2的类型%T, num2的值=%v,num2的地址=%v, num2这个指针指向的值%v \n", num2, num2, &num2, *num2) 
	 //上面的代码执行结果: num2的类型*int, num2的值=0xc0000aa090,num2的地址=0xc0000ce020, num2这个指针指向的值100 
}