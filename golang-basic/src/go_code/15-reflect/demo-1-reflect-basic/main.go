/*
1.反射的基本介绍:
	(1).反射可以在运行时动态获取变量的各种信息,比如变量的类型(type)、类别(kind)
	(2).如果是结构体变量,还可以获取到结构体本身的信息(包括结构体的字段、方法)
	(3).通过反射,可以修改变量的值,可以调用关联的方法
	(4).使用反射,需要 import("reflect")
2.反射的应用场景,反射常见应用场景有以下两种:
	(1).不知道接口调用哪个函数,根据传入参数在运行时确定调用的具体接口,这种需要对函数或方法反射.例如以下这种桥接模式 和 "反射1.jpg"中的问题.
			func bridge(funcPtr interface{}, args ...interface{})
		第一个参数funcPtr以接口的形式传入函数指针,函数参数args以可变参数的形式传入,bridge函数中可以用反射来动态执行funcPtr函数.
	(2).对结构体序列化时,如果结构体有指定Tag,也会使用到反射生成对应的字符串.
3.反射重要的函数和概念.
	(1). reflect.TypeOf(变量名),获取变量的类型,返回 reflect.Type 类型.
	(2). reflect.ValueOf(变量名),获取变量的值,返回 reflect.Value 类型, reflect.Value 是一个结构体类型. 
		通过 reflect.Value 可以获取到关于该变量的很多信息.
	(3). 变量、interface{} 和 reflect.Value 是可以相互转换的,这点在实际开发中,会经常使用到.
4.反射的快速入门-快速入门说明
	演示对 基本数据类型、interface{}、reflect.Value 进行反射的基本操作
	演示对 结构体类型、interface{}、reflect.Value 进行反射的基本操作
*/

package main

import (
	"fmt"
	"reflect"
)


// 反射案列1-通过反射获取的传入的变量的 type , kind, 值
func first (b interface{}) {
	// 1.先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType = ", rTyp)
	
	// 2.获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	k := rVal.Kind()
	fmt.Println("kind =", k)
	fmt.Println("value  =", rVal.Int())

	n2 := 2 + rVal.Int()
	fmt.Println("n2=", n2)
	fmt.Printf("rVal=%v rVal type=%T\n", rVal, rVal)

	//使用反射的方式来获取变量的值(并返回对应的类型),要求数据类型匹配,比如x是int,那么就应该使用 reflect.Value(x).Int(),而不能使用其它的,否则报panic
	//n3 := rVal.Float()
	//fmt.Println("n3=", n3)
	
	//下面将 rVal 转成 interface{}
	iV := rVal.Interface()
	//将 interface{} 通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Println("num2=", num2)
}


type Student struct {
	Name string
	Age int
}

type Monster struct {
	Name string
	Age int
} 

// 反射案列2-对结构体的反射-通过反射获取的传入的变量的 type , kind, 值
func second (b interface{}) {
	//1. 先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	//2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)

	//3. 获取变量对应的 Kind
	kind1 := rVal.Kind()
	kind2 := rTyp.Kind()
	fmt.Printf("kind1 = %v; kind2 = %v\n", kind1, kind2) 

	//4.下面将 rVal 转成 interface{}
	iV := rVal.Interface()
	fmt.Printf("iv=%v, iv type=%T \n", iV, iV)

	//将 interface{} 通过断言转成需要的类型.这里就简单使用了一个带检测的类型断言,但是可以使用 swtich 的断言形式来做的更加的灵活
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}

}


func main() {

	var num int = 100
	first (num)	// 案例:对(基本数据类型、interface{}、reflect.Value)进行反射的基本操作

	stu := Student {
		Name:"Tom",
		Age:20,
	}
	second (stu) // 案例:对结构体的反射

	//three ()
}