/*
json基本介绍:JSON是在2001年开始推广使用的数据格式,目前以及称为主流的数据格式
	概述: JSON(JavaScript Object Notation)是一种轻量级的数据交换格式.易于人阅读和编写.同时也易于机器解析和生成.

	JSON易于机器解析和生成,并有效地提升网络传输速率,通常程序在网络传输时会先将数据(结构体、map等)序列化成json字符串,在反序列化恢复成原来的数据类型(结构体、map等).这种方式已然称为各个语言的标准.

	Golang ====序列化====> json字符串 ====网络传输====> 程序 ====反序列化====> 其它语言

json数据格式说明:
	在JS语言中,一切都是对象.因此,任何支持的类型都可以通过JSON来表示,例如字符串、数字、对象、数组、map、结构体等.
	JSON键值对是用来保存数据的一种方式.

json的序列化:
	介绍:json序列化是指,将有key-value结构的数据类型(比如结构体、map、切片)序列化成json字符串的操作.
	序列化方法；
		func Marshal(v interface{}) ([]byte, err)	// Marshal函数返回v的json编码
	注意事项: 对于结构体的序列化,如果希望重命名序列化后的key名字,那么可以给struct指定一个tag标签
		type People struct {
			Name string	`json:"name"`
			Age int `json:"age"`
			Birthday string	`json:"birthday"`
		}

json的反序列化:
	介绍:json反序列化是指,将json字符串反序列化成对应的数据类型(比如结构体、map、切片)的操作.
	注意事项:
		1) 在反序列化一个json字符串时,要确保反序列化后的数据类型和原来序列化前的数据类型一致.
		2) 如果直接定义json格式字符串,需要对字符串中的引号进行转义.
		3) 如果json字符串是通过网络传输或者程序获取到的,则不需要对json数据中的引号进行转义处理,因为内部已经转义.
*/

package main

import (
	"fmt"
	"encoding/json"
)

type Student struct {
	Name string
	Age int
	Birthday string
}

// 结构体的序列化
func first() {

	student := Student {
		Name : "Lily",
		Age : 19,
		Birthday : "2002-09-19",
	}
	
	data, err := json.Marshal(&student)		// 将结构体序列化
	if err != nil {
		fmt.Printf("序列化错误 err = %v \n", err)
	}
	fmt.Printf("student序列化结果= %v \n", string(data))
	
}

// map的序列化
func second() {

	var a map[string] interface{}
	a = make(map[string] interface{})		// 使用map需要先make
	a["name"] = "Lucy"
	a["age"] = "20"
	a["Birthday"] = "2001-09-19"
	
	data, err := json.Marshal(a)			// 将map序列化
	if err != nil {
		fmt.Printf("序列化错误 err = %v \n", err)
	}
	fmt.Printf("map数据序列化结果= %v \n", string(data))

}


// 切片的序列化
func three() {
	
	var m1 map[string] interface{}
	m1 = make(map[string] interface{})		// 使用map需要先make
	m1["name"] = "Lucy"
	m1["age"] = "20"
	m1["Birthday"] = [2]string{"2001-09-19","2001-08-16"}

	var m2 map[string] interface{}
	m2 = make(map[string] interface{})		// 使用map需要先make
	m2["name"] = "Lily"
	m2["age"] = "19"
	m2["Birthday"] = "2002-09-19"

	var slice []map[string] interface{}
	slice = append(slice, m1)

	data, err := json.Marshal(slice)			// 将切片序列化
	if err != nil {
		fmt.Printf("序列化错误 err = %v \n", err)
	}
	fmt.Printf("切片数据序列化结果= %v \n", string(data))	//  切片数据序列化结果= [{"Birthday":["2001-09-19","2001-08-16"],"age":"20","name":"Lucy"}]

}


// 基本数据类型的序列化
func four() {
	
	var price float64 = 99999.9999
	
	data, err := json.Marshal(price)			// 将切片序列化
	if err != nil {
		fmt.Printf("序列化错误 err = %v \n", err)
	}
	fmt.Printf("基本数据类型数据序列化结果= %v \n", string(data))	//  基本数据类型数据序列化结果= 99999.9999

	// 由结果可知: 对基本数据类型进行序列化意义不大
}


type People struct {
	Name string	`json:"name"`
	Age int `json:"age"`
	Birthday string	`json:"birthday"`
}

// 对于结构体的序列化,如果希望重命名序列化后的key名字,那么可以给struct指定一个tag标签
func five() {
	people := People {
		Name : "Lily",
		Age : 19,
		Birthday : "2002-09-19",
	}
	
	data, err := json.Marshal(&people)
	if err != nil {
		fmt.Printf("序列化错误 err = %v \n", err)
	}
	fmt.Printf("给结构体指定tag标签重命名key => people序列化结果= %v \n", string(data))
	
}


func main() {

	first()		// 结构体的序列化

	second()	// map的序列化

	three()		// 切片的序列化

	four()		// 基本数据类型的序列化(意义不大)
	
	five() 		//  对于结构体的序列化,如果希望重命名序列化后的key名字,那么可以给struct指定一个tag标签

	six()		// 将json字符串反序列化成结构体
	
	seven()		// 将json字符串反序列化成map

	eight()		// 将json字符串反序列化成切片
	
}


func six() {
	
	/*
	Student 结构体不带tag标签
	*/
	// 项目开发中,json字符串一般是通过网络传输或者读取文件获取.直接定义json格式字符串,需要对字符串中的引号进行转义
	str := "{\"Name\":\"Lily\",\"Age\":19,\"Birthday\":\"2002-09-19\"}"	
	var student Student
	err := json.Unmarshal([]byte(str), &student)
	if err != nil {
		fmt.Printf("Unmarshal err = %v \n", err)
	}
	fmt.Printf("反序列化后 student= %v \n", student)


	/*
	People 结构体带tag标签,对于带tag标签的结构体,结构体变量名不变或变了后的json字符串都可以反序列化为结构体,以下两个json字符串都可以:
		"{\"Name\":\"Lucy\",\"Age\":20,\"Birthday\":\"2001-09-19\"}"
		"{\"name\":\"Lucy\",\"age\":20,\"birthday\":\"2001-09-19\"}"
	*/
	str = "{\"name\":\"Lucy\",\"age\":20,\"birthday\":\"2001-09-19\"}"
	var people People
	err = json.Unmarshal([]byte(str), &people)
	if err != nil {
		fmt.Printf("Unmarshal err = %v \n", err)
	}
	fmt.Printf("反序列化后 people= %v \n", people)
}


func getJsonData() string {

	var a map[string] interface{}
	a = make(map[string] interface{})		// 使用map需要先make
	a["name"] = "Lucy"
	a["age"] = "20"
	a["Birthday"] = "2001-09-19"
	
	data, err := json.Marshal(a)			// 将map序列化
	if err != nil {
		fmt.Printf("序列化错误 err = %v \n", err)
	}
	
	return string(data)
}


func seven() {
	
	// str := "{\"Name\":\"Lily\",\"Age\":19,\"Birthday\":\"2002-09-19\"}"
	str := getJsonData()			// 通过程序获取到的,则不需要对json数据中的引号进行转义处理,因为内部已经转义.
	var a map[string] interface{}	// 注意:反序列化map不需要make,因为 Unmarshal 函数中已经封装了make操作

	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("Unmarshal err = %v \n", err)
	}
	fmt.Printf("反序列化后 map a= %v \n", a)
}


func eight() {
	
	str := "[{\"Birthday\":[\"2001-09-19\",\"2001-08-16\"],\"age\":\"20\",\"name\":\"Lucy\"}]"

	var slice []map[string] interface{}		// 定义一个切片,反序列化不需要make,因为 Unmarshal 函数中已经封装了make操作

	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("Unmarshal err = %v \n", err)
	}
	fmt.Printf("反序列化后 slice= %v \n", slice)
}