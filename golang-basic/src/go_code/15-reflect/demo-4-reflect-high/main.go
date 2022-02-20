/*
反射最佳实践:
	(1) 使用反射来遍历结构体字段,调用结构体的方法,并获取结构体标签的值.
		func (Value) Method  ==>  func (v Value) Method(i int) Value
		func (Value) Call  ==>  func (v Value) Call(in []Value) []Value // 传入参数和返回结果(参数)是 []reflect.Value
		反射的 structtag 的核心代码:
			tag := tye.Elem().Field(0).Tag.Get("json")
			fmt.printf("tag=%s\n", tag) 
	(2) 使用反射的方式来获取结构体的tag标签,遍历字段的值,修改字段值,调用结构体方法(要求:通过传递地址的方式完成)
	(3) 定义两个函数test1和test2,定义一个适配器函数用作统一处理接口
	(4) 使用反射操作任意结构体类型
	(5) 使用反射创建并操作结构体

*/

package main

import (
	"fmt"
	"reflect"
)

// 定义一个Student结构体
type Student struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Score float32 `json:"score"`
	Sex string
}

// GetSum方法-返回两个数的和
func (s Student) GetSum(n1, n2 int) int {
	return n1 + n2
}

// Set方法,接收四个值,给s赋值
func (s Student) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
} 

// Print方法,显示s的值
func (s Student) Print() {
	fmt.Println("---start~----")
	fmt.Println(s)
	fmt.Println("---end~----")
}

func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)	// 获取reflect.Type 类型
	val := reflect.ValueOf(a)	// 获取reflect.Value 类型
	kd := val.Kind()			// 获取到a对应的类别
	if kd != reflect.Struct {	// 如果传入的不是struct，就退出
		fmt.Println("expect struct")
		return		
	}
	num := val.NumField()		// 获取到该结构体有几个字段
	fmt.Printf("struct has %d fields\n", num)

	// 遍历结构体所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: 值为=%v\n", i, val.Field(i))
		tagVal := typ.Field(i).Tag.Get("json")	//获取到struct标签, 注意需要通过reflect.Type来获取tag标签的值
		
		// 如果该字段于tag标签就显示，否则就不显示
		if tagVal != ""{
			fmt.Printf("Field %d: tag为=%v\n", i, tagVal)
		}
	}

	// 获取到该结构体有多少个方法
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)

	// 方法的排序默认是按照函数名的排序(ASCII码),也就说按字母顺序排序
	val.Method(1).Call(nil)		// 获取到第二个方法。调用它

	//调用结构体的第1个方法Method(0)
	var params []reflect.Value			// 声明 []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params)	 // 传入的参数是 []reflect.Value, 返回[]reflect.Value
	fmt.Println("res=", res[0].Int())	 // 返回结果, 返回的结果是 []reflect.Value*/
}

func first () {
	// 创建了一个Student实例
	var s Student = Student {
		Name: "小莉",
		Age: 20,
		Score: 30.8,
	}
	// 将Student实例传递给TestStruct函数
	TestStruct(s)
}

func main() {

	first ()	// 使用反射来遍历结构体字段,调用结构体的方法,并获取结构体标签的值
}