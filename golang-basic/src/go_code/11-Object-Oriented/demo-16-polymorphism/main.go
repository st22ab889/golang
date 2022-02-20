/*
多态:变量(实例)具有多种形态.面向对象的第三大特征
	基本介绍:在Golang语言,多态特征是通过接口实现的.可以按照统一的接口来调用不同的实现.这时接口就呈现不同的形态.

接口体现多态特征:
	多态参数
	多态数组	
*/

package main

import (
	"fmt"
)

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	name string
}

func (p Phone) Start () {
	fmt.Println("Phone::Start()")
}

func (p Phone) Stop () {
	fmt.Println("Phone::Stop()")
}

func (p Phone) Call () {
	fmt.Println("Phone::Call()")
}

type Camera struct {
	name string
}

func (c Camera) Start(){
	fmt.Println("Camera::Start()")
} 

func (c Camera) Stop(){
	fmt.Println("Camera::Start()")
}


type Computer struct {
}
func (computer Computer) Working(usb Usb) {		// "(usb Usb)"就是多态参数
	usb.Start()
	// 使用类型断言实现: 如果usb是指向Phone结构体变量,则还需要调用Call方法
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}


func main() {

	// 定义多态数组,可以存放Phone和Camera的结构体变量
	var usbArr [3]Usb
	usbArr[0] = Phone{"HuaWei"}
	usbArr[1] = Phone{"iphone"}
	usbArr[2] = Camera{"Canon"}	

	var computer Computer
	for _, v := range usbArr{
		computer.Working(v)
		fmt.Println()
	}

	fmt.Println(usbArr)
}