/*
面向对象编程-继承:继承可以解决代码复用,让编程更加靠近人类思维
	当多个结构体存在相同的属性(字段)和方法时,可以从这些结构体中抽象出结构体,在该结构体中定义这些相同的属性和方法.其它的结构体不需要重新定义这些属性和方法,只需要嵌套一个匿名结构体即可.
	也就说说在Golang中,如果一个struct嵌套了另一个匿名结构体,那么这个结构体可以直接访问匿名结构体的字段和方法,从而实现了继承特性.
	嵌套匿名结构体的基本语法:
		type Goods struct {
			Name string
			Price int
		}

		type Book struct {
			Goods		// 这里就是嵌套匿名结构体Goods 
			Writer string
		}
	继承给编程带来的遍历:
		1) 代码的复用性提高了
		2) 代码的扩展性和维护性提高了
*/

package main

import (
	"fmt"
)

// 下面是定义学生
type Student struct {
	Name string
	Age int
	Score int
}

func (stu *Student) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score int) {
	stu.Score = score
}

func (stu *Student) GetSum(n1 int, n2 int) int {
	return n1 + n2
}

// 下面是定义小学生
type Pupil struct { 
	Student 	//嵌入了Student匿名结构体
}

func (p *Pupil) testing() {				// 这是Pupil结构体特有的方法		
	fmt.Println("小学生正在考试中.....")
}

// 下面是定义大学生
type Graduate struct {
	Student		//嵌入了Student匿名结构体
}

func (p *Graduate) testing() {			// 这是Graduate结构体特有的方法
	fmt.Println("大学生正在考试中.....")
}

func main() {

	// 当对结构体嵌入了匿名结构体,使用方法会发生变化
	pupil := &Pupil{}
	pupil.Student.Name = "小学生"
	pupil.Student.Age = 8
	pupil.testing();
	pupil.Student.SetScore(66)
	pupil.Student.ShowInfo()
	fmt.Println("pupil :: res=", pupil.Student.GetSum(1, 2))

	// 调用方法可以简化
	graduate := &Graduate{} 
	graduate.Name = "大学生"		// graduate.Student.Name = "大学生"	
	graduate.Age = 18				// graduate.Student.Age = 18
	graduate.testing();
	graduate.SetScore(99)			// graduate.Student.SetScore(99)
	graduate.ShowInfo()				// graduate.Student.ShowInfo()	
	fmt.Println("graduate :: res=", graduate.GetSum(66, 99))

	/*
	总结:
		1) 当直接通过 graduate 访问字段或方法时,编译器会先看 graduate 类型有没有这个字段或方法
		2) 如果有就直接调用 graduate 类型的字段或方法;如果没有就去看 graduate 中嵌入的匿名结构体 Student 中有没有相应的字段或方法
		3) 如果匿名结构体 Student 中有相应的字段或方法就调用,如果没有就继续查找,如果都找不到就报错
	*/

}