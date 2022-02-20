/*
结构体的注意事项和使用细节:
	1. 结构体的所有字段在内存中是连续的
	2. 结构体是用户单独定义的类型,和其它类型进行转换时需要有完全相同的字段(名字、个数和类型,以及顺序)
	3. 结构体使用type进行重新定义(相当于取别名),Golang认为是新的数据类型, 但是相互之间可以强转.
			例1:
				type Student struct {
					Name string
					Age int
				}
				type Stu Student
				func main () {
					var stu1 Student
					var stu2 Stu
					stu2 = stu1 	// 错误, 可以修改为 stu2 = Stu(stu1)
				}
			例2:
				type integer int
				func main () {
					var i integer = 10
					var j int = 20
					// j = i // 错误, 可以修改为 j = int(i)
					j = int(i)
					fmt.Println(i , j);
					// i = j //错误
					i = integer(j)
					fmt.Println(i , j);
				}
	4. struct的每个字段上可以写上一个tag,该tag可以通过反射机制获取,常见的使用场景就是序列号和反序列号
*/

package main

import (
	"fmt"
	"encoding/json"
)

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point
}

type Rect2 struct {
	leftUp, rightDown *Point
}

func first () {
	//r1有四个int, 在内存中是连续分布
	r1 := Rect{Point{1,2}, Point{3,4}} 

	//打印地址结果为: r1.leftUp.x 地址=0xc0000a8060 r1.leftUp.y 地址=0xc0000a8068 r1.rightDown.x 地址=0xc0000a8070 r1.rightDown.y 地址=0xc0000a8078 
	fmt.Printf("r1.leftUp.x 地址=%p r1.leftUp.y 地址=%p r1.rightDown.x 地址=%p r1.rightDown.y 地址=%p \n", &r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)

	//r2有两个 *Point类型，这个两个*Point类型的本身地址也是连续的，但是他们指向的地址不一定是连续
	r2 := Rect2{&Point{10,20}, &Point{30,40}} 

	//打印地址结果为: r2.leftUp 本身地址=0xc000088220 r2.rightDown 本身地址=0xc000088228 
	fmt.Printf("r2.leftUp 本身地址=%p r2.rightDown 本身地址=%p \n", &r2.leftUp, &r2.rightDown)

	//他们指向的地址不一定是连续(取决于系统在运行时是如何分配)
	fmt.Printf("r2.leftUp 指向地址=%p r2.rightDown 指向地址=%p \n", r2.leftUp, r2.rightDown)	// r2.leftUp 指向地址=0xc0000aa080 r2.rightDown 指向地址=0xc0000aa090 
}



type A struct {
	Num int
	Name string
}

type B struct {
	Num int
	Name string
}

func second () {
	var a A
	var b B
	a = A(b)	
	fmt.Println(a , b);	// 结构体之间可以转换,但是要求相互转换的两个结构体的字段要完全一样(包括:名字、个数、类型、顺序)
}



type Student struct {
	Name string
	Age int
}
type Stu Student

type integer int

func three () {
	var stu1 Student
	var stu2 Stu
	// stu2 = stu1 	// 错误, 可以修改为 stu2 = Stu(stu1)
	stu2 = Stu(stu1)
	fmt.Println(stu1 , stu2);
	// stu1 = stu2		// 错误
	stu1 = Student(stu2)
	fmt.Println(stu1 , stu2);


	var i integer = 10
	var j int = 20
	// j = i // 错误
	j = int(i)
	fmt.Println(i , j);
	// i = j //错误
	i = integer(j)
	fmt.Println(i , j);
}



type People struct {
	Name string `json:"name"`	// `json:"name"` 就是 struct tag
	Age int `json:"age"`
}

func four() {
	asian := People{"亚洲人", 20}
	// json.Marshal函数中使用反射
	jsonStr, err := json.Marshal(asian)
	if err != nil {
		fmt.Println("json 处理错误: ", err);
	}

	// 没有使用"struct tag",打印结果为: jsonStr =  {"Name":"亚洲人","Age":20}
	// 使用了"struct tag",打印结果为: jsonStr =  {"name":"亚洲人","age":20}
	fmt.Println("jsonStr = ", string(jsonStr));	

	// 如果不使用 "struct tag", json处理后的字段名也是首字母大写,这种命名方式会觉得不习惯
	// 如果将 struct 的首字母小写,会发现返回的是空字符串,因为 json.Marshal 相当于在其他包访问结构体,首字母小写就不能在其他包访问
	// 最终解决方案是使用 "struct tag"
}


func main() {

	first ()

	second ()

	three ()

	four()
}