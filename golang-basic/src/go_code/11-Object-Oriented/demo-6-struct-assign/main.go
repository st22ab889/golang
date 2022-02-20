/*
Golang在创建结构体实例(变量)时,可以直接指定字段的值
	方式1:
		var stu1 Student = Student{"Tom", 10}
		stu2 := Student{"tom", 10}
		var sut3 Student = Student{
			Name: "Tom",
			Age: 10,
		}
		stu4 := Student{
			Name: "Tom",
			Age: 10,
		}
	方式2:
		var stu1 *Student = &Student{"Tom", 10}
		var stu2 *Student = &Student{
			Name: "Tom",
			Age: 10,
		}
*/

package main

import (
	"fmt"
)

type Stu struct {
	Name string
	Age int
}

func first () {
	//方式1
	//在创建结构体变量时，就直接指定字段的值
	var stu1 = Stu{"小明", 19} 
	stu2 := Stu{"小明~", 20}

	//在创建结构体变量时，把字段名和字段值写在一起, 这种写法不依赖字段的定义顺序.
	var stu3 = Stu{
			Name :"jack",
			Age : 20,
		}
	stu4 := Stu{
		Age : 30,
		Name : "mary",
	}

	fmt.Println(stu1, stu2, stu3, stu4)
}


func second() {
	//方式2: 返回结构体的指针类型
	var stu5 *Stu = &Stu{"小王", 29} 
	stu6 := &Stu{"小王~", 39}

	//在创建结构体指针变量时，把字段名和字段值写在一起, 这种写法不依赖字段的定义顺序.
	var stu7 = &Stu{
		Name : "小李",
		Age :49,
	}
	stu8 := &Stu{
		Age :59,
		Name : "小李~",
	}
	
	fmt.Println(*stu5, *stu6, *stu7, *stu8) 
}

func main() {

	first ()

	second ()
}