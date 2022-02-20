/*
接口应用场景介绍:
	a.比如制造轰炸机武装直升机,专家只需把飞机需要的功能/规格定下来即可,然后让其它人具体实现
	b.项目经理管理三个程序员,开了一个软件, 为了控制和管理软件,项目经理可以定义一些接口,然后由程序员具体实现

注意事项和细节:
	1) 接口本身不能创建实例,但是可以指向一个实现了该接口的自定义类型的变量(实例)
	2) 接口中所有的方法都没有方法体(即都是没有实现的方法)
	3) 在Golang中,一个自定义类型需要将某个接口的所有方法都实现,才能说这个自定义类型实现了该接口
	4) 一个自定义类型只有实现了某个接口,才能将该自定义类型的实例(变量)赋给接口类型
	5) 只要是自定义数据类型,就可以实现接口,不仅仅是结构体类型
	6) 一个自定义类型可以实现多个接口
	7) Golang接口中不能有任何变量
	8) 一个接口(比如A接口)可以继承多个别的接口(比如B、C接口),这时如果要实现A接口,也必须将B、C接口的方法也全部实现
	9) interface默认是个指针(引用类型),如果没有对interface初始化就使用,那么会输出nil
	10) 空接口 interface{} 没有任何方法,所以所有类型都实现了空接口,即可以把任何一个变量赋给空接口
*/

package main

import (
	"fmt"
)

type People interface {
	Speak()
}

type Student struct {
	Name string
}

func (stu Student) Speak() {
	fmt.Println("stu::", stu.Name, "::Speak()")
}

func first () {
	var stu Student
	stu.Name = "Lily"
	// 接口本身不能创建实例,但是可以指向一个实现了该接口的自定义类型的变量(实例)
	var people People = stu
	people.Speak()
}


type integer int
//  只要是自定义数据类型,就可以实现接口,不仅仅是结构体类型
func (i integer) Speak() {
	fmt.Println("integer:: i = ", i)
}

func second () {
	var i integer = 99
	var p People = i
	p.Speak()
}


type Teacher interface {
	Teach()
}

type Professor struct {

} 

func(pro Professor) Teach() {
	fmt.Println("the professor is teaching")
}

func(pro Professor) Speak() {
	fmt.Println("the professor is speaking")
}

func three() {
	// Professor 实现了 People 和 Teacher 两个接口
	var professor Professor
	var people People = professor
	var teacher Teacher = professor

	people.Speak()
	teacher.Teach()
}


type Educator interface {
	People
	Teacher
	Research()
} 

type LiteraryEducator struct {
	
}

// 如果实现了 Educator 接口的方法, 那么 People 和 Teacher 接口的方法也都要实现
func (literaryEducator LiteraryEducator) Research() {
	fmt.Println("Researching Literary")
}

func(literaryEducator LiteraryEducator)  Speak() {
	fmt.Println("the literary educator is speaking")
}

func (literaryEducator LiteraryEducator) Teach() {
	fmt.Println("the literary educator is teaching")
}

func four() {
	var literaryEducator LiteraryEducator
	var educator Educator = literaryEducator
	educator.Research()
	educator.Teach()
	educator.Speak()

	// interface默认是个指针(引用类型),如果没有对interface初始化就使用,那么会输出nil
	var educator2 Educator
	fmt.Println("educator2 = ", educator2)	// educator2 =  <nil>
}

func main() {

	first ()

	second ()

	three ()

	four()

	five ()
}

type Empty interface {

}

func five () {
	
	// 空接口 interface{} 没有任何方法,所以所有类型都实现了空接口,即可以把任何一个变量赋给空接口
	
	// 用法一
	var stu Student
	var e Empty = stu
	fmt.Println("e = ", e)	// e =  {}

	// 用法二
	var e2 interface{} = stu	// 	var e2 interface{}  表示声明一个空接口变量
	fmt.Println("e2 = ", e2)	// e2 =  {}

	var e3 interface{}  
	e3 = integer(0)
	fmt.Println("e3 = ", e3)	// e3 =  0


	var  f float64 = 99.99
	e2 = f
	e = f
	fmt.Println("e2 = ", e2, ", e = ", e)	// e2 =  99.99 , e =  99.99
}