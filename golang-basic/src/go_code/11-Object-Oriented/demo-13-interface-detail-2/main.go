/*
下面三种情况都能编译过:
		type AA interface {
			Test01()
			Test02()
		}
		type BB interface {
			Test01()
			Test03()
		}
	第一种情况:
		type CC struct {
		}
		func (cc CC) Test01() {
			fmt.Println("CC::Test01")
		}
		func (cc CC) Test02() {
			fmt.Println("CC::Test02")
		}
		func (cc CC) Test03() {
			fmt.Println("CC::Test03")
		}
	
	第二种情况:
		type DD interface {
			AA
			BB
		}
		type ABC struct{
		}
		func(abc ABC) Test01() {
			fmt.Println("ABC::Test01")
		}
		func(abc ABC) Test02() {
			fmt.Println("ABC::Test02")
		}
		func(abc ABC) Test03() {
			fmt.Println("ABC::Test03")
		}
	
	第三种情况:要注意
		type Usb interface {
			Record()
		}
		type Computer struct {
		}
		func (this *Computer) Record() {	// 注意"*Computer"是指针
			fmt.Println("Record()")
		}

		func mainhree() {
			var computer Computer = Computer{}
			// var usb Usb = computer	// error, 这里会报Computer没有实现Usb接口
			var usb Usb = &computer		// 改成这样可以编译过
			usb.Record()
		}

*/

package main

import (
	"fmt"
)

type AA interface {
	Test01()
	Test02()
}

type BB interface {
	Test01()
	Test03()
}

// 情况1:能编译
type CC struct {
}
func (cc CC) Test01() {
	fmt.Println("CC::Test01")
}
func (cc CC) Test02() {
	fmt.Println("CC::Test02")
}
func (cc CC) Test03() {
	fmt.Println("CC::Test03")
}

func first () {
	// 可以编译过,没有问题
	cc := CC{}
	var aa AA = cc
	fmt.Println(aa)

	var bb BB = cc
	fmt.Println(bb)
}


// 情况2:能编译
type DD interface {
	AA
	BB
}

type ABC struct{
}
func(abc ABC) Test01() {
	fmt.Println("ABC::Test01")
}
func(abc ABC) Test02() {
	fmt.Println("ABC::Test02")
}
func(abc ABC) Test03() {
	fmt.Println("ABC::Test03")
}

func second () {
	abc := ABC{}
	var dd DD = abc
	dd.Test01();
	fmt.Println(dd)
}


// 情况3: 能编译
type Usb interface {
	Record()
}
type Computer struct {
}
func (this *Computer) Record() {	// 注意"*Computer"是指针
	fmt.Println("Record()")
}

func three() {
	var computer Computer = Computer{}
	// var usb Usb = computer	// error, 这里会报Computer没有实现Usb接口
	var usb Usb = &computer		// 改成这样可以编译过
	usb.Record()
	fmt.Println(usb)
	
	// 下面这两种写法都可以
	computer.Record()
	(&computer).Record()
}


func main() {

	first ()	// 情况一

	second ()	// 情况二

	three ()	// 情况三:要注意
}

