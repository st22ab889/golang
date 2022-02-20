
/*
Golang变量使用的三种方式(变量表示内存中的一个存储区域,该区域有自己的名称(变量名)和类型(数据类型)):
	(1)指定变量类型，声明后若不赋值,使用默认值, 比如 int 默认值为0, string默认值为空串, 小数默认为0
	(2)根据值自行判定变量类型(类型推导)
	(3)省略var, 注意 := 左侧的变量不应该是已经声明过的, 否则会导致编译错误

变量在同一个作用域内不能重名
变量 = 变量名 + 值 + 数据类型
	
多变量声明:一次性声明多个变量, Golang也提供这样的语法
*/
package main
import "fmt"
func main ()  {
	
	// 声明(定义)变量,声明后若不赋值,使用默认值,int的默认值是0
	var i int
	i = 10
	fmt.Println("i = ", i)

	// 类型推导
	var num = 10.01
	fmt.Println("num = ", num)

	// 省略var,不能省略 : , 下面的方式等价于声明变量(var name string)和赋值(name =  "Golang")
	name := "Golang"
	fmt.Println("name = ", name)

	//*********************************************************************

	// 多变量声明-方式1
	var n1, n2, n3 int
	fmt.Println("n1 = ", n1, "n2 = ", n2, "n3 = ", n3)

	// 多变量声明-方式2
	var age, user, height = 25, "json", 175
	fmt.Println("age = ", age, "user = ", user, "height = ", height)

	// 多变量声明-方式3-使用类型推导
	phone, num := "iphone", 8800
	fmt.Println("phone = ", phone, "num = ", num)

	//*********************************************************************

	fmt.Println("i1 = ", i1, "i2 = ", i2, "i3 = ", i3)

	fmt.Println("i4 = ", i4, "i5 = ", i5, "i6 = ", i6)

	//*********************************************************************
	
	// 该区域的数据值可以在同一类型范围内不断变化
	var aa int = 10
	aa = 20
	fmt.Println("aa = ", aa)
	aa = 30
	fmt.Println("aa = ", aa)
	// aa = 50.5  // 50.5不是int类型,不能改变数据类型,所以这里不能赋值给变量aa

}

// 定义全局变量 -- 规范写法是把全局变量定义在"func main ()"上面位置,这里为了方便查看
var i1 = 100
var i2 = 200
var i3 = "this is a global var"

// 上面的声明方式,也可以改成一次性声明
var(
	i4 = 100
	i5 = 200
	i6 = "this is a global var"
)

// 34节