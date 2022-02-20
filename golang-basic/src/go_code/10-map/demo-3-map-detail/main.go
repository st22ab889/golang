/*
map的使用细节:
	1) map是引用类型,遵守引用类型传递的机制,在一个函数接收map,修改后,会直接修改原来的map
	2) map的容量达到后,再想map增加元素,会自动扩容,并不会发生panic,也就是说map能动态的增长键值对
	3) map的value也经常使用struct类型,更适合管理复杂的数据(比value是一个map更好)
*/

package main

import (
	"fmt"
)

func modifyMap(map1 map[int]int) {
	//  map是引用类型,遵守引用类型传递的机制
	map1[10] = 99
}

func first () {
	map1 := make(map[int]int)
	map1[1] = 13
	map1[2] = 100
	map1[10] = 22
	fmt.Println("修改前 map1 = ", map1)		// map[1:13 2:100 10:22]
	modifyMap(map1)
	fmt.Println("修改后 map1 = ", map1)		// map[1:13 2:100 10:99]
}


func second () {
	//  map的容量达到后,再想map增加元素,会自动扩容,并不会发生panic
	map2 := make(map[int]int, 1)
	map2[1] = 13
	map2[2] = 100
	fmt.Println("map2 = ", map2)	
}

// 定义一个结构体
type people struct {
	Name string
	Age int
	Addr string
}

func three (){
	
	// map的value也经常使用struct类型,更适合管理复杂的数据
	peopleMap := make(map[string]people, 10)
	people1 := people{"Tom", 22, "北京"}
	people2 := people{"Lily", 23, "上海"}
	peopleMap["p1"] = people1
	peopleMap["p2"] = people2

	fmt.Println("peopleMap = ", peopleMap)	

	// 遍历
	for k,v := range peopleMap {
		fmt.Printf("编号是 %v \n", k)
		fmt.Printf("名字是 %v \n", v.Name)
		fmt.Printf("年龄是 %v \n", v.Age)
		fmt.Printf("地址是 %v \n", v.Addr)
	}
}


func main() {

	first ()

	second ()

	three ()
}