/*
string和slice
1.string底层是一个byte数组,因此string也可以进行切片处理
2.string和切片在内存的形式
3.string是不可变的,也就是说不能通过 str[0]='z' 的方式来修改字符串
4.如果需要修改字符串,可以先将 string -> []byte, 或者 []rune -> 修改 -> 重写 转成 string
*/

package main
import(
	"fmt"
	_ "time"
	_ "math/rand"
	_ "runtime"
	_ "strconv"
)

func first() {
	str := "hello world"
	slice := str[4:]
	fmt.Println("slice = ", slice)

	// string是不可变的,不能通过 str[0]='z' 的方式来修改字符串
	// str[0] = 'z'		// [编译不会通过,报错,原因是string不可变]
}

func second() {

	str := "hello world"
	// 如果需要修改字符串,可以先将 string -> []byte, 转成[]byte后,可以处理英文和数字,但不能处理中文
	// 原因是 []byte 是按字节来处理,而一个汉字是3个字节,因此会出现乱码 
	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println("str = ", str)

	// []rune -> 修改 -> 重写 转成 string
	// 将 string 转成 []rune 即可,因为 []rune 是按字符处理,兼容汉字
	arr2 := []rune(str)
	arr2[0] = '中'
	str = string(arr2)
	fmt.Println("str = ", str)
}


func main() {

	first() 

	second()
}



