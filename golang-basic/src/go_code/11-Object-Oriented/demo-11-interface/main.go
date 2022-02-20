/*
在Golang中多态特性主要是通过接口来实现的.

接口
	基本介绍:接口(interface)类型可以定义一组方法,但是这些方法不需要实现.并且接口不能包含任何变量.某个自定义类型要使用接口的时候,根据情况把这些方法写出来.
	基本语法:
			type 接口名 interface {
				method1(参数列表) 返回值列表
				method2(参数列表) 返回值列表
				...
			}
		说明:
			a.接口里的所有方法都没有方法体,即接口的方法都是没有实现的方法.接口体现了程序设计的多态和高内聚低耦合的思想.
			b.Golang中的接口,不需要显示的实现.只要一个变量,含有接口类型中的所有方法,那么这个变量就实现了这个接口.因此Golang中没有implement这样的关键字.

*/

package main

import (
	"fmt"
)

// 声明(定义)一个接口,接口中的方法只声明不需要实现
type Usb interface {
	Start()
	Stop()
}

// 定义 Phone, 实现 Usb接口的方法
// 
type Phone struct {

}
func (p Phone) Start() {
	fmt.Println("Phone::Start()")
}
func (p Phone) Stop() {
	fmt.Println("Phone::Stop()")
}

// 定义 Camera, 实现 Usb接口的方法
type Camera struct {

}
func (c Camera) Start() {
	fmt.Println("Camera::Start()")
}
func (c Camera) Stop() {
	fmt.Println("Camera::Stop()")
}

// 定义 Computer
type Computer struct{

}
// Working 方法，接收一个Usb接口类型变量(所谓实现Usb接口，就是指实现了 Usb接口声明所有方法)
func (c Computer) Working(usb Usb){
	usb.Start()
	usb.Stop()
}


func first () {

	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	computer.Working(phone)		// interface默认是个指针(引用类型), 所以这里传递的phone实例,实际上是传递的是phone实例的地址
	computer.Working(camera)	// interface默认是个指针(引用类型), 所以这里传递的camera实例,实际上是传递的是camera实例的地址
}


func second () {

}


func main() {

	first ()

	second ()

	//three ()
}