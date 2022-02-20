
package main
import(
	"fmt"
	"unsafe"
)

func main()  {
	
	var b = false
	fmt.Println("b",b)
	b = true
	fmt.Println("b",b)

	// 1.bool类型占用存储空间是1个字节
	fmt.Println("b占用空间 = ", unsafe.Sizeof(b))
	// 2.bool类型只能取true或者false,不能以0和非0替换, 这是和C语言的区别
	// b = 0

}

