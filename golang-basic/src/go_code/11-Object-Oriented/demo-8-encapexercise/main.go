/*
面向对象编程三大特性
	基本介绍:Golang仍然有面向对象编程的继、封装和多态的特性,只是实现的方式和其它OOP语言不一样

面向对象编程思想-抽象:把一类事物共有的属性和行为提取出来形成一个物理模型(模板),这种研究问题的方法称为抽象.

面向对象编程思想-封装(encapsulation):
	1) 把抽象出的字段和对字段的操作封装在一起,数据被保护在内部,程序的其他包只有通过被授权的操作(方法),才能对字段进行操作.
		例如:对电视机的操作就是典型封装
	2) 封装的理解和好处
		a. 隐藏实现细节
		b. 可以对数据进行验证,保证安全合理
	3) 如何体现封装
		a. 对结构体中的属性进行封装
		b. 通过方法、包实现封装
	4) 封装的实现步骤
		a. 将结构体、字段(属性)的首字母小写(不能导出,其它包不能使用,类似 private)
		b. 给结构体所在包提供一个工厂模式的函数,首字母大写.类似构造函数
		c. 提供一个首字母大写的Set方法(类似于其它语言的public),用于对属性判断并赋值
		d. 提供一个首字母大写的Get方法(类似于其它语言的public),获取属性的值
		特别注意: 在Golang开发中并没有特别强调封装,这点并不像java.所以不用总是用java的语法特性来看待Golang, Golang本身对面向对象的特性做了简化.
*/

package main

import (
	"fmt"
	"go_code/11-Object-Oriented/demo-8-encapexercise/model"
)


func first () {
	//创建一个account变量
	account := model.NewAccount("jzh11111", "000111", 40)
	if account != nil {
		fmt.Println("创建成功=", account)
		account.Deposite(100000, "000111")
		account.Query("000111")
		account.WithDraw(1000, "000111")
		account.Query("000111")
	} else {
		fmt.Println("创建失败")
	}
}

func main() {

	first ()	// 演示封装

}