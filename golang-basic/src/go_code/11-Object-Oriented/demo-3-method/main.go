/*
方法: Golang中的方法是作用在指定的数据类型上的(即:和指定的数据类型绑定),因此自定义类型都可以有方法,而不仅仅是struct
	1. 基本介绍: 结构体除了一些字段外,还可以有一些行为,这些行为就是需要声明(定义)的方法.
	2. 方法的声明和调用:
		type A struct {
			Num int
		}
		func (a A) test() {			// 表示A结构体有一个名为 test 的方法,"(a A)"体现了test方法和A类型绑定的关系
			fmt.Println(a. Num);
		}
		var t A
		t.test()
	3. 方法的调用和传参机制原理
		方法的调用和传参机制和函数基本一样,不一样的地方是方法调用时,会将调用方法的变量,当做实参也传递给方法(如果变量是值类型则进行值拷贝,如果变量是引用类型则进行地址拷贝).
*/

package main

import (
	"fmt"
)


type Person struct {
	Name string
}

// 表示Person结构体有一个名为 test 的方法,"(p Person)"体现了test方法和Person类型绑定的关系
// test方法只能通过Person类型的变量来调用,而不能直接调用,也不能使用其它类型变量来调用
// p表示哪个Person变量调用,这个p就是它的副本,这点和函数传参非常相似. 
// p这个名字由程序员指定,不是固定的.一个结构体如果有多个方法和它绑定,这个变量p可以一样,也可以不一样,但是最好保持一样.
func (p Person) test() {  	
	fmt.Println(p. Name);
}

func (p2 Person) speak () {
	fmt.Println(p2. Name, "is a good man.");
}

// 累加
func (p3 Person) calu (n int) {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	fmt.Println(p3. Name, " get a value is ", res);
}

// 求和
func (p3 Person) getSum(n1 int, n2 int) int {
	return n1 + n2
}

func first () {
	var p Person
	p.Name = "Tom"
	p.test()
	p.speak()
	p.calu(3)
	fmt.Println("getSum = ", p.getSum(90, 9));
}


type Circle struct {
	redius float64
}

// 当结构体变量调用这个方法时,这里是值传递(相当于把结构体拷贝了一份传递到这个方法中)
func (circle Circle) area1() float64 {
	return 3.14 * circle.redius * circle.redius
}

// 当结构体变量调用这个方法时,这里是引用传递(相当于把结构体的地址传递到这个方法中)
// 为了提高效率,通常方法和结构体的指针类型绑定
func (circle *Circle) area2() float64 {
	// 因为 circle 是指针,所以标准的访问其字段的方式是 (*circle).redius
	// return 3.14 * (*circle).redius * (*circle).redius	
	
	// 这种方式是编译器底层做了优化, circle.redius 等价于 (*circle).redius 
	return 3.14 * circle.redius * circle.redius
}

func second () {
	var c Circle
	c.redius = 7.0
	// res := (&c).area()	// 标准方式
	res := c.area2()	// 编译器底层做了优化 (&c).area() 等价于 c.area(),编译器会自动加上 &c
	fmt.Println("area = ", res)
}


func main() {

	first ()

	second ()
}