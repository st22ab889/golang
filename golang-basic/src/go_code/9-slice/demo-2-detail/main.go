/*
切片注意事项和细节说明:
	1.切片初始化 
		var slice = arr[startIndex:endIndex]
		说明: 从arr数组下标为startIndex取到下标为endIndex的元素(不包含arr[endIndex]这个元素)
	2.切片初始化时仍然不能越界。范围在 0 - len(arr) 之间,可以动态增长
		var slice = arr[0:end] 可以简写 var slice = arr[:end]
		var slice = arr[start:len(arr)] 可以简写 var slice = arr[start:]
		var slice = arr[0:len(arr)]	可以简写	var slice = arr[:]
	3.cap是一个内置函数,用于统计切片的容量,既最大可以存放多少个元素
	4.切片定义完后,还不能使用,因为本身是一个空的,需要让其引用到一个数组 或 make一个空间供切片使用
	5.切片可以继续切片
	6.用append内置函数,可以对切片进行动态追加,切片append操作的底层原理分析；
		1) 切片append操作的本质就是对数组扩容
		2) go底层会创建一个新的数组newArr(安装扩容后大小)
		3) 将slice原来包含的元素拷贝到新的数组newArr
		4) slice重新引用到newArr. 注意:newArr是在底层来维护的,程序员不可见
	7.切片的拷贝操作:切片使用copy内置函数完成拷贝
	8.切片是引用类型, 所以在传递时遵守引用传递机制
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
	var arr	[5]int = [...]int{10, 20, 30, 40, 50}
	slice := arr[:]

	slice2 := slice[1:2]
	fmt.Println("slice2 = ", slice2)
	slice2[0] = 100		//因为 slice、slice2指向的数据空间是同一个arr, 因此 slice2[0] = 100 会导致 slice 和 arr 中的元素都会被修改
	fmt.Println("slice = ", slice)
	fmt.Println("arr = ", arr)
}

func second() {
	var slice []int = []int{10, 20, 30}
	slice = append(slice, 40, 50, 60)
	fmt.Println("append函数对切片动态追加: slice = ", slice)

	// 通过append将切片slice追加给切片slice
	slice = append(slice, slice...);
	fmt.Println("切片追加给切片: slice = ", slice)
}

func three() {
	var slice1 []int = []int{1, 2, 3, 4, 5}
	var slice2 = make([]int, 10)
	// 切片使用copy内置函数完成拷贝
	copy(slice2, slice1)	// 表示把slice1拷贝到slice2,注意顺序
	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)
	/*
	注意:
		copy(para1, para2) 参数的数据类型是切片
		从上面的代码可知,slice1和slice2的数据空间是独立,互相不影响,也就说说 slice1[0] = 999, slice2[0]仍然是1
	*/

	// 注意:下面的代码没有错误
	var a []int = []int{1, 2, 3}
	var b = make([]int, 1)
	fmt.Println("b = ", b)	// b =  [0]
	copy(b, a)	// 这里能copy成功,只是copy一个元素而已
	fmt.Println("b = ", b)	// b =  [1]
}

func main() {

	first() 

	second()

	three()
}

