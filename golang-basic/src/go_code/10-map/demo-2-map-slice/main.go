/*
map切片:
	基本介绍:切片的数据类型如果是map,则称为 slice of map, map切片,这样使用则map个数就可以动态变化了

map排序:
	1) Golang中没有一个专门的方法针对map的key进行排序
	2) Golang中的map默认时无序的,注意也不是按照添加的顺序存放的,每次遍历得到的输出可能不一样
	3) Golang中map的排序,是先将key进行排序,然后根据key值遍历输出即可	
*/

package main

import (
	"fmt"
	"sort"
)


func first () {
	// 声明一个map切片
	var peopleMap []map[string]string
	// 准备放入两个信息
	peopleMap = make([]map[string]string, 2)
	
	if peopleMap[0] == nil {
		peopleMap[0] = make(map[string]string, 2)
		peopleMap[0]["name"] = "Tom"
		peopleMap[0]["age"] = "20"
		peopleMap[0]["addr"] = "北京"
	}

	if peopleMap[1] == nil {
		peopleMap[1] = make(map[string]string, 2)
		peopleMap[1]["name"] = "Lily"
		peopleMap[1]["age"] = "18"
		peopleMap[1]["addr"] = "广州"
	}

	/*
	// 下面就会产生越界
	if peopleMap[2] == nil {
		peopleMap[2] = make(map[string]string, 2)
		peopleMap[2]["name"] = "Lucy"
		peopleMap[2]["age"] = "19"
		peopleMap[2]["addr"] = "上海"
	}
	*/

	// 使用切片的 append 函数, 可以动态的增加元素
	// 定义一个新的map
	newPeople := map[string]string {
		"name" : "Lucy",
		"age" : "19",
		"addr" : "上海",
	}
	peopleMap = append(peopleMap, newPeople)

	fmt.Println("peopleMap = ", peopleMap)
}


// map排序
func second () {

	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90
	// 多运行几次可以看到打印的结果都不一样
	fmt.Println(map1)

	// 按照map的key的顺序进行排序输出
	// 操作方法:先将map的key放入到切片中,对切片进行排序,遍历切片,然后按照key来输出map的值
	var keys []int
	for k,_ := range map1 {
		keys = append(keys, k)
	} 
	sort.Ints(keys)
	fmt.Println("keys = ", keys)

	for _, k := range keys {
		fmt.Printf("map1[%v] = %v \n", k, map1[k])
	}
}


func main() {

	first ()

	second ()
}