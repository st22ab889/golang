/*
结构体: 一个程序就是一个世界,有很多对象(变量)

结构体和结构体变量(实例)的区别和联系:
	1) 结构体是自定义的数据类型,代表一类事物
	2) 结构体变量(实例)是具体的、实际的,代表一个具体变量
	3) 注意事项和细节说明
		a.字段声明的语法和变量一样. 示例:	字段名 字段类型
		b.字段的类型可以为: 基本类型、数组或引用类型
		c.在创建一个结构体变量后,如果没有给字符按赋值,都对应一个零值(默认值),规则如下:\
			布尔类型是 false, 整型是 0, 字符串是""
			数组类型的默认值和它的元素类型相关
			指针、slice和map的零值都是 nil ,即还没有分配空间,如果需要使用这样的字段,需要先make才能使用
		d.不同结构体变量的字段是独立、互不影响的,一个结构体变量字段的更改,不影响另外一个

声明结构体:
	type 标识符 struct {
		field1 type
		field2 type
	}	

结构体的字段(属性):
	1) 从概念或叫法上看,结构体字段(field)也成为属性
	2) 字段是结构体的一个组成部分,一般是基本数据类型、数组,也可以是引用类型

创建结构体变量和访问结构体字段:
	方式1:直接声明
		var cat1 Cat
	方式2： {}
		cat2 := Cat{"小喵咪", 1, "白色"}
	方式3:	&
		var cat3 *Cat = new(Cat)
	方式4:	{}
		var cat4 *Cat = &Cat{}
	说明:
		a. 方式3和方式4返回的是 结构体指针
		b. 结构体指针访问字段的标准方式应该是: (*结构体指针).字段名
		c. 但Golang做了一个简化,也支持 结构体指针.字段名 ,更加符合程序员的使用习惯, go编译器底层把 "结构体指针.字段名" 转化为 "(*结构体指针).字段名"这种形式

结构体(struct)类型的分配机制:

Golang语言面向对象编程说明:
	1) Golang也支持面向对象编程(OOP),但是和传统的面向对象编程有区别,并不是纯粹的面向对象语言.所以说Golang支持面向对象编程特性是比较准确的.
	2) Golang没有类(class),Golang语言的结构体(struct)和其它编程语言的类(class)有同等的地位,可以理解Golang是基于struct来实现OOP特性的.
	3) Golang面向对象编程非常简洁,去掉了传统OOP语言的继承、方法重载、构造函数和析构函数、隐藏的this指针等等.
	4) Golang仍然有面向对象编程的继承、封装和多态的特性,只是实现的方式和其它OOP语言不一样,比如继承:Golang没有extends关键字,继续是通过匿名字段来实现.
	5) Golang面向对象(OOP)很优雅,OOP本身就是语言类型系统(type system)的一部分,通过接口(interface)关联,耦合性低,也非常灵活.在Golang中面向接口编程是非常重要的特性.
*/

package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Age int
	Color string
}

func first () {
	// 创建一个Cat的变量
	var cat1 Cat
	cat1.Name = "喵咪"
	cat1.Age = 2
	cat1.Color = "白色"
	fmt.Printf("cat1的地址 = %p \n", &cat1)
	fmt.Println("cat1 = ", cat1)

	fmt.Println("cat name = ", cat1.Name)
	fmt.Println("cat age = ", cat1.Age)
	fmt.Println("cat color = ", cat1.Color)
}

// 指针、slice和map的零值都是 nil ,即还没有分配空间,如果需要使用这样的字段,需要先make才能使用
type Person struct {
	Name string
	Age int
	Scores [5]float64
	ptr *int				// 指针
	slice []int				// 切片	
	map1 map[string]string 	// 切片
}

func second () {
	// 定义结构体变量
	var p1 Person
	fmt.Println("p1 = ", p1)
	if(p1.ptr == nil){
		fmt.Println("ptr为nil")
	}
	if(p1.slice == nil){
		fmt.Println("slice为nil")
	}
	if(p1.map1 == nil){
		fmt.Println("map1为nil")
	}

	// 使用slice一定要先make
	p1.slice = make([]int, 10)
	p1.slice[0] = 99

	// 使用map一定要先make
	p1.map1 = make(map[string]string)
	p1.map1["key1"] = "Lucy"

	fmt.Println("p1 = ", p1)
}

func three () {
	//创建结构体变量的四种方式

	// 方式1: 直接声明
	var cat1 Cat
	fmt.Println("cat1 = ", cat1)

	// 方式2:
	cat2 := Cat{"小喵咪", 1, "白色"}
	fmt.Println("cat2 = ", cat2)

	// 方式3:
	var cat3 *Cat = new(Cat)
	(*cat3).Name = "大喵咪"		// 因为cat3是一个指针,这是标准的给字段赋值的方式
	cat3.Age = 1	// 也可以这样赋值(不加*号), Golang为了程序员方便使用,底层会对  cat3.Age = 1 进行处理, 会给 cat3 加上取值运算,转化为 (*cat3).Age = 1
	fmt.Println("cat3 = ", *cat3)	// cat3 =  {大喵咪 1 }
	fmt.Println("cat3 = ", cat3)	// cat3 =  &{大喵咪 1 }

	// 方式4: 
	// var cat4 *Cat = &Cat{"黑猫", 1 ,"黑色"}	// 可以直接给字段赋值
	var cat4 *Cat = &Cat{}
	(*cat4).Name = "黑猫"
	cat4.Age = 1
	cat4.Color = "黑色"
	fmt.Println("cat4 = ", *cat4)

	// fmt.Println("cat4.Age = ", *cat4.Age)  // 这样写会报错,因为 . 运算符的优先级比 * 运算符高 

}

func four () {

	cat1 := Cat{"小喵咪", 1, "白色"}
	fmt.Printf("cat1的地址 = %p \n", &cat1)

	var cat2 Cat = cat1		// 这里是结构体拷贝(值拷贝), 也就是说 cat1 和 cat2 有各自的一块内存区间
	cat2.Name = "小猫"
	fmt.Printf("cat1.Name = %v, cat2.Name = %v \n", cat1.Name, cat2.Name) 	// cat1.Name = 小喵咪, cat2.Name = 小猫
	fmt.Printf("cat2的地址 = %p \n", &cat2)

	var cat3 *Cat = &cat1	// 这里是引用传递,也就是说 cat1 和 cat3 指的是同一块内存区间
	// 下面相当于是 (*cat3).Name = "小白描"
	cat3.Name = "小白描"
	fmt.Printf("cat1.Name = %v, cat3.Name = %v \n", cat1.Name, cat3.Name)	// cat1.Name = 小白描, cat3.Name = 小白描
	fmt.Printf("cat3的地址 = %p , cat3的值 = %p \n", &cat3, cat3)
}

func main() {

	first ()

	second ()

	three ()

	four () 
}