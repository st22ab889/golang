/*
方法的声明(定义):
	func (receiver type) methodName (参数列表) (返回值类型) {
		方法主体
		return 返回值
	}
	1) 参数列表:表示方法输入
	2) receiver type: 表示这个方法和type这个类型绑定,或者说该方法作用于type类型
	3) receiver type: type可以是结构体,也可以是其它的自定义类型. 比如 type integer int
	4) receiver: 就是type类型的一个变量(实例)
	5) 参数列表: 表示方法输入
	6) 返回值类型: 表示返回的值,可以多个
	7) 方法主体: 实现某一功能的代码块
	8) return语句不是必须的


方法注意事项和细节讨论:
	1) 如果方法绑定的是值类型,如果结构体的变量调用这个方法,相当于是把这个结构体拷贝一份传入到这个方法中(也就是说在方法调用中遵守值类型的传递机制,是值拷贝传递方式).
	   如果方法绑定的是指针类型,如果结构体的变量调用这个方法,相当于是把这个结构体的地址传入到这个方法中.
	2) 如果程序员希望在方法中修改结构体变量的值,可以通过结构体指针的方式来处理
	3) Golang中的方法作用在指定的数据类型上的(即:和指定的数据类型绑定),因此自定义类型都可以有方法,而不仅仅是struct,比如int、float32等都可以有方法
	4) 方法的访问范围控制的规则,和函数一样.方法名首字母小写,只能在本包访问,方法首字母大写,可以在本包和其他包访问
	5) 如果一个类型实现了 String() 这个方法, 那么 fmt.Println 默认会调用这个变量的 String() 进行输出

方法和函数区别:
	1)调用方式不一样.
		函数的调用方式:		函数名(实参列表)
		方法的调用方式:		变量.方法名(实参列表)
	2)对于普通函数,接收者为值类型时,不能将指针类型的数据直接传递,反之亦然
	3)对于方法(如struct的方法),接收者为值类型时,可以直接用指针类型的变量调用方法,反过来也可以

*/

package main

import (
	"fmt"
)

type integer int
// (i integer) 说明这里是值传递
func (i integer) print () {
	fmt.Println("i = ", i)
}
// (i *integer) 说明是引用传递
func (i *integer) change () {
	*i = *i + 1
}

func first () {
	var i integer = 10
	i.print()
	i.change()
	fmt.Println("change i = ", i)
}


type Student struct {
	Name string
	Age int
}

func (stu *Student) String() string {
	str := fmt.Sprintf("Name=[%v], Age=[%v]", stu.Name, stu.Age)
	return str
}

func second () {
	stu := Student{
		Name : "Tom",
		Age : 22,
	}
	fmt. Println("stu = ", stu)		// stu =  {Tom 22}
	
	// 如果实现了 *Student 类型的 String 方法, 就会自动调用
	fmt. Println("&stu = ", &stu)	// &stu =  Name=[Tom], Age=[22]
}

// 普通函数
func test01(s Student) {
	fmt. Println(s.Name)	
}
func test02(s *Student) {
	fmt. Println(s.Name)	
}
// 方法.当结构体变量调用这个方法时, 不管时传入指针变量还是值变量,这里都是值拷贝
func (stu Student) test03() {
	fmt. Println(stu.Name)
}
// 方法.当结构体变量调用这个方法时, 不管时传入指针变量还是值变量,这里都是引用传递
func (stu *Student) test04() {
	fmt. Println(stu.Name)
}


func three() {
	stu := Student{
		Name : "Tom",
		Age : 22,
	}
	test01(stu)		// 必须传值类型
	test02(&stu)	// 必须传指针类型
	
	stu.test03()
	// (&stu).test03() 等价于 stu.test03()
	(&stu).test03()	// 形式上是传入地址,但本质仍然是值拷贝,因为方法的声明是: func (stu Student) test03()
	

	// stu.test04() 等价于 (&stu).test04() 
	stu.test04()	// 形式上是传入值类型,但本质仍然是地址传递(拷贝),因为方法的声明是: func (stu *Student) test04()
	(&stu).test04()	
	/*
	总结: 不管调用形式如何,真正决定是值拷贝还是地址拷贝,是要看这个方法和哪个类型绑定
	*/
}


func main() {

	first ()

	second ()

	three()
}