/*
函数中的defer
	1.当Golang执行到一个defer时,不会立即执行defer后的语句,而是将defer后的语句压入一个栈(暂时称这个栈为defer栈),然后继续执行函数下一个语句
	2.当函数执行完毕后,在从defer栈中,依次从栈顶取出语句执行(注:遵守栈先入后出的机制)
	3.在defer将语句放入到栈时,也会将相关的值拷贝同时入栈

defer的最佳实践:defer主要用在当函数执行完毕后,可以及时的释放函数创建的资料
	1.在Golang编程中通常的做法是,创建资源后,比如(打开了文件、获取数据库的衔接、锁资源),可以执行 defer file.Close() 、defer connect Close()
	2.在defer后,可以继续使用创建的资源.当函数完毕后,系统会依次从defer栈中取出语句,关闭资源
	3.这种机制非常简洁
*/

package main
import(
	"fmt"
)

func sum(n1 int, n2 int) int {
	// 当执行到defer时,暂时不会执行,会将defer后面的语句压入到独立的栈(defer栈)
	defer fmt.Println("sum() n1 = ", n1)	// 这里仍然打印n1的值,并不会打印 n1++ 后的值
	defer fmt.Println("sum() n2 = ", n2)	// 这里仍然打印n2的值,并不会打印 n2++ 后的值

	n1++
	n2++

	r1 := n1 + n2
	fmt.Println("sum() r1 = ", r1)

	// 当执行完return前面一句后,再从defer栈,按照先入后出的方式出栈,然后再执行return语句

	return r1; 

	
}

func main() {

	res1 := sum(89, 8)
	fmt.Println("res1 = ", res1)

	/*
	执行结果如下:
		sum() r1 =  101
		sum() n2 =  9 
		sum() n1 =  90
		res1 =  101 
	*/

}