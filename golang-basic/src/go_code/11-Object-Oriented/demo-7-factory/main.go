/*
Golang的结构体没有构造函数,通常可以使用工厂模式来解决这个问题

下面的定义的结构体类型首字母是小写,如果想在其他包创建 student 的实例, 这时不能直接创建, 需要用到工厂模式来解决
		type student struct {
			Name string
		}
	
*/

package main

import (
	"fmt"
	"go_code/11-Object-Oriented/demo-7-factory/model"
)

// 使用工厂模式实现跨包创建结构体实例(变量)
func main() {

	/*
	//创建student实例,因为student结构体是首字母小写,所以不能直接创建结构体实例
	var stu = model.student{
		Name :"tom",
		score : 78.9,
	}
	*/

	// 通过工厂模式来解决结构体首字母小写的问题
	var stu = model.NewStudent("tom", 98.8)
	fmt.Println(*stu)

	// 结构体中的score字段首字母是小写,所以在其它包不可以直接访问,但是可以通过一个支持跨包访问的方法进行访问
	fmt.Println("name=", stu.Name, " score=", stu.GetScore())


}
