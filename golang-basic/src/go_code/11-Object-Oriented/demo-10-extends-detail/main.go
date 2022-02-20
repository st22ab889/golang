/*
深入讨论继承:
	1) 结构体可以使用嵌套匿名结构体所有的字段和方法,即:首字母大写或者小写的字段、方法都可以使用
	2) 匿名结构体字段访问可以简化
			b.a.name = "tom"	可以简化为  b.name = "tom"
			b.a.say()			可以简化为	b.say()
	3) 当结构体和匿名结构体有相同的字段或者方法时,编译器采用就近访问原则,如希望访问匿名结构体的字段和方法,可以通过匿名结构体名来区分
	4) 结构体嵌入两个(或多个)匿名结构体,如两个匿名结构体有相同的字段和方法(同时结构体本身没有同名的字段和方法),在访问时就必须指定匿名结构体名字,否则编译报错
	5) 如果一个struct嵌套了一个有名结构体,这种模式就是组合,如果是组合关系,那么在访问组合的结构体的字段或方法时,必须带上结构体的名字
			type A struct {
				Name string
			}
			type B struct {
				a A
			}
	6) 嵌套匿名结构体后,也可以在创建结构体变量(实例)时,直接指定各个匿名结构体字段的值
	7) 结构体的匿名字段是基本数据类型
	8) 多重继承说明:如果一个struct嵌套了多个匿名结构体,那么该结构体可以直接访问嵌套的匿名结构体的字段和方法,从而实现多重继承.
		a.如果嵌入的匿名结构体有相同的字段名或者方法名,则在访问时,需要通过匿名结构体类型名来区分
		b.为了保证代码的简洁性,建议尽量不使用多重继承
*/

package main

import (
	"fmt"
)

// 示例1 

type A struct {
	Name string
	age int
}

func (a *A) SayOk() {
	fmt.Println("A SayOk", a.Name)
}

func (a *A) hello() {
	fmt.Println("A hello", a.Name)
}

type B struct {
	A
	Name string 
}

func (b *B) SayOk() {
	fmt.Println("B SayOk", b.Name)
}

func first () {

	// 结构体A和B都有 Name 字段, 都绑定了一个名为 SayOk 的方法, 结构体A是结构体B的匿名结构体

	var b B
	b.Name = "jack" 	
	b.A.Name = "scott"
	b.age = 100 
	b.SayOk()  			// 	B SayOk  jack
	b.A.SayOk() 		//  A SayOk scott
	b.hello() 			//  A hello scott , 要注意这里为什么打印的是"A hello scott"
}

// 示例2

type AA struct {
	Name string
	age int
}

type BB struct {
	Name string
	Score float64
}

type CC struct {
	AA
	BB
	// 如果一个struct嵌套了多个匿名结构体,那么该结构体可以直接访问嵌套的匿名结构体的字段和方法,从而实现多重继承.
	// CC 这个结构体相当于实现了多重继承
}

func second () {

	var cc CC
	
	// 报编译错误,如果CC没有Name字段,而AA和BB有Name,这时就必须通过指定匿名结构体名字来区分,这个规则对方法也是一样的
	// cc.name = "abc"
	
	cc.AA.Name = "abc"
	fmt.Println("name = ", cc.AA.Name)
}

// 示例3

type DD struct {
	aa AA  
}

func three () {
	var dd DD
	// 如果DD中有一个有名结构体，则访问有名结构体的字段时，就必须带上有名结构体变量的名字	
	dd.aa.Name = "abcd"
	fmt.Println("name = ", dd.aa.Name)
}

// 示例4

type Goods struct {
	Name string
	Price float64
}

type Brand struct {
	Name string
	Address string
}

type TV struct {
	Goods
	Brand	
}

type TV2 struct {
	*Goods
	*Brand	
}

func four () {
	//嵌套匿名结构体后，也可以在创建结构体变量(实例)时，直接指定各个匿名结构体字段的值
	tv := TV{Goods{"电视机", 19999.99}, Brand{"长虹", "四川"}}
	fmt.Println(tv.Goods.Name)
	fmt.Println(tv.Price)
	
	tv2 := TV{
		Goods{
			Price: 1999.99,
			Name: "空调",
		},
		Brand{
			Name: "美的",
			Address: "广东",
		},
	}
	fmt.Println("tv", tv)
	fmt.Println("tv2", tv2)

	tv3 := TV2{&Goods{"电视机", 19999.99}, &Brand{"长虹", "四川"}}
	tv4 := TV2{
		&Goods{
			Price: 1999.99,
			Name: "空调",
		},
		&Brand{
			Name: "美的",
			Address: "广东",
		},
	}
	fmt.Println("tv3", *tv3.Goods, *tv3.Brand)
	fmt.Println("tv4", *tv4.Goods, *tv4.Brand)
}

// 示例5

type people struct {
	Name string
	Age int
}

type student struct {
	people	// 匿名字段是结构体
	int		// 匿名字段是基本数据类型
	n int

	// 如果一个结构体有int类型的匿名字段,就不能再有第二个int类型的匿名字段
	// 如果需要有多个int的字段,则必须给int字段指定名字
}

func five () {
	var stu student
	stu.Name = "Lily"
	stu.Age = 22
	stu.int = 66	// 给 基本数据类型 的匿名字段赋值
	stu.n =  99
	fmt.Println("stu = ", stu)
}

func main() {

	first ()	// 当结构体和匿名结构体有相同的字段或者方法时,编译器采用就近访问原则

	second ()	// 实现多重继承

	three ()	// 有名结构体

	four ()		// 嵌套匿名结构体后，也可以在创建结构体变量(实例)时，直接指定各个匿名结构体字段的值

	five ()		// 给 基本数据类型 的匿名字段赋值
}